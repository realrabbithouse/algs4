package xmlrpc

import (
	"bytes"
	"fmt"
	"regexp"
)

var faultRegex = regexp.MustCompile(`<fault>(\s|\S)+</fault>`)

type FaultError struct {
	Code   int    `xml:"faultCode"`
	String string `xml:"faultString"`
}

func (f FaultError) Error() string {
	return fmt.Sprintf("Code %d: %s", f.Code, f.String)
}

type rawResponse []byte

func (r rawResponse) Fault() error {
	if !faultRegex.Match(r) {
		return nil
	}

	var (
		dec   Decoder
		fault FaultError
	)
	if err := dec.Decode(bytes.NewReader(r), &fault); err != nil {
		return err
	}

	return &fault
}

func (r rawResponse) Unmarshal(v any) error {
	var dec Decoder
	return dec.Decode(bytes.NewReader(r), v)
}
