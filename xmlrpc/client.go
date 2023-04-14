package xmlrpc

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/rpc"
	"net/url"
	"strconv"
	"sync"
)

type clientCodec struct {
	enc *Encoder
	dec *Decoder

	endpoint   *url.URL
	httpClient *http.Client

	// customHeaders presents some use specified headers.
	customHeaders map[string]string

	// cookies stores cookies received on the last request.
	cookies http.CookieJar

	mu      sync.Mutex
	pending map[uint64]response

	// In-flight raw response, need to be unmarshalled.
	response rawResponse

	// ready presents completed http request by sequence ID.
	ready chan uint64

	// close notifies clientCodec to close itself.
	close chan struct{}
}

type response struct {
	ServiceMethod string
	httpResponse  *http.Response
}

func (c *clientCodec) WriteRequest(req *rpc.Request, params any) (err error) {
	var body bytes.Buffer
	if err := c.enc.Encode(&body, req.ServiceMethod, params); err != nil {
		return err
	}

	httpRequest, err := http.NewRequest("POST", c.endpoint.String(), &body)
	if err != nil {
		return err
	}
	// Apply customer headers.
	for key, value := range c.customHeaders {
		httpRequest.Header.Set(key, value)
	}
	httpRequest.Header.Set("Content-Type", "text/xml")
	httpRequest.Header.Set("Content-Length", strconv.Itoa(body.Len()))

	if c.cookies != nil {
		for _, cookie := range c.cookies.Cookies(c.endpoint) {
			httpRequest.AddCookie(cookie)
		}
	}

	resp, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return err
	}

	if c.cookies != nil {
		c.cookies.SetCookies(c.endpoint, resp.Cookies())
	}

	c.mu.Lock()
	c.pending[req.Seq] = response{
		ServiceMethod: req.ServiceMethod,
		httpResponse:  resp,
	}
	c.mu.Unlock()

	c.ready <- req.Seq

	return nil
}

func (c *clientCodec) ReadResponseHeader(resp *rpc.Response) error {
	select {
	case seq := <-c.ready:
		c.mu.Lock()
		r := c.pending[seq]
		delete(c.pending, seq)
		c.mu.Unlock()

		resp.Seq = seq
		resp.ServiceMethod = r.ServiceMethod

		defer r.httpResponse.Body.Close()
		if r.httpResponse.StatusCode < 200 || r.httpResponse.StatusCode >= 300 {
			resp.Error = fmt.Sprintf("bad response: %s", r.httpResponse.Status)
			return nil
		}

		body, err := io.ReadAll(r.httpResponse.Body)
		if err != nil {
			resp.Error = err.Error()
			return nil
		}

		raw := rawResponse(body)
		// The raw response may wrap a corresponding response or an error.
		if err := raw.Fault(); err != nil {
			resp.Error = err.Error()
			return nil
		}

		// Record response for unmarshal later.
		c.response = raw
		return nil
	case <-c.close:
		// For network has been closed.
		return net.ErrClosed
	}
}

func (c *clientCodec) ReadResponseBody(x any) error {
	if x == nil {
		return nil
	}

	r := bytes.NewReader(c.response)
	if err := c.dec.Decode(r, x); err != nil {
		return err
	}

	return nil
}

func (c *clientCodec) Close() error {
	// Close all idle connections.
	c.httpClient.CloseIdleConnections()
	close(c.close)
	return nil
}

type Client struct {
	*rpc.Client
}

type clientOptions struct {
	httpClient    *http.Client
	customHeaders map[string]string
	charsetReader func(charset string, input io.Reader) (io.Reader, error)
}

// WithHttpClient uses c as http client in XML-RPC client.
func WithHttpClient(c *http.Client) Option {
	return func(o *clientOptions) {
		o.httpClient = c
	}
}

// WithCustomHeaders transforms additional headers into http request.
func WithCustomHeaders(headers map[string]string) Option {
	return func(o *clientOptions) {
		o.customHeaders = headers
	}
}

// WithCharsetReader uses a specified charset reader for decoding.
func WithCharsetReader(r func(charset string, input io.Reader) (io.Reader, error)) Option {
	return func(o *clientOptions) {
		o.charsetReader = r
	}
}

// Option represents an option applied on XML-RPC client.
type Option func(*clientOptions)

// NewClient create an instance of rpc.Client, which is used to send request to XML-RPC services.
func NewClient(requestURL string, opts ...Option) (*Client, error) {
	var options clientOptions
	for _, opt := range opts {
		opt(&options)
	}
	if options.httpClient == nil {
		options.httpClient = http.DefaultClient
	}

	u, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	codec := &clientCodec{
		enc:           &Encoder{},
		dec:           &Decoder{CharsetReader: options.charsetReader},
		endpoint:      u,
		httpClient:    options.httpClient,
		customHeaders: options.customHeaders,
		cookies:       jar,
		pending:       make(map[uint64]response),
		ready:         make(chan uint64),
		close:         make(chan struct{}),
	}

	return &Client{rpc.NewClientWithCodec(codec)}, nil
}
