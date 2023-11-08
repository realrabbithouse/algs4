package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

var httpVersion = flag.Int("version", 1, "The http version HTTP/1 or HTTP/2")

func main() {
	flag.Parse()

	var t http.RoundTripper
	switch *httpVersion {
	case 1:
		t = transport1()
	case 2:
		t = transport2()
	default:
		logrus.Fatal("unknown http version")
	}

	client := &http.Client{Transport: t}

	res, err := client.Get("https://" + "localhost:9973")
	if err != nil {
		logrus.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.WithFields(logrus.Fields{
		"code": res.StatusCode,
		"body": string(body),
	}).Info("server respond")
}

func transport1() *http.Transport {
	return &http.Transport{
		// Original configurations from `http.DefaultTransport` variable.
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		// Set it false to enforce HTTP/1.
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,

		// Our custom configurations.
		ResponseHeaderTimeout: 10 * time.Second,
		DisableCompression:    true,
		// Set DisableKeepAlives to true when using HTTP/1 otherwise it will
		// cause error: "dial tcp [::1]:8090: socket: too many open files".
		DisableKeepAlives: false,
		TLSClientConfig:   tlsConfig(),
	}
}

func transport2() *http2.Transport {
	return &http2.Transport{
		TLSClientConfig:    tlsConfig(),
		DisableCompression: true,
		AllowHTTP:          false,
	}
}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("./cert/ca.crt")
	if err != nil {
		logrus.Fatal(err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}
}
