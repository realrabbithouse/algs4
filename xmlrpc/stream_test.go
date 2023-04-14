package xmlrpc

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var encodeTestcases = []struct {
	value any
	xml   string
}{
	{100, "<value><int>100</int></value>"},
	{"Once upon a time", "<value><string>Once upon a time</string></value>"},
	{"Mike & Mick <London, UK>", "<value><string>Mike &amp; Mick &lt;London, UK&gt;</string></value>"},
	{Base64("T25jZSB1cG9uIGEgdGltZQ=="), "<value><base64>T25jZSB1cG9uIGEgdGltZQ==</base64></value>"},
	{true, "<value><boolean>1</boolean></value>"},
	{false, "<value><boolean>0</boolean></value>"},
	{12.134, "<value><double>12.134</double></value>"},
	{-12.134, "<value><double>-12.134</double></value>"},
	{738777323.0, "<value><double>738777323</double></value>"},
	{time.Unix(1386622812, 0).UTC(), "<value><dateTime.iso8601>20131209T21:00:12</dateTime.iso8601></value>"},
	{[]any{1, "one"}, "<value><array><data><value><int>1</int></value><value><string>one</string></value></data></array></value>"},
	{&struct {
		Title  string
		Amount int
	}{"War and Peace", 20}, "<value><struct><member><name>Title</name><value><string>War and Peace</string></value></member><member><name>Amount</name><value><int>20</int></value></member></struct></value>"},
	{&struct {
		Value any `xml:"value"`
	}{}, "<value><struct><member><name>value</name><value/></member></struct></value>"},
	{
		map[string]any{"title": "War and Piece", "amount": 20},
		"<value><struct><member><name>amount</name><value><int>20</int></value></member><member><name>title</name><value><string>War and Piece</string></value></member></struct></value>",
	},
	{
		map[string]any{
			"Name":  "John Smith",
			"Age":   6,
			"Wight": []float32{66.67, 100.5},
			"Dates": map[string]any{"Birth": time.Date(1829, time.November, 10, 23, 0, 0, 0, time.UTC), "Death": time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)}},
		"<value><struct><member><name>Age</name><value><int>6</int></value></member><member><name>Dates</name><value><struct><member><name>Birth</name><value><dateTime.iso8601>18291110T23:00:00</dateTime.iso8601></value></member><member><name>Death</name><value><dateTime.iso8601>20091110T23:00:00</dateTime.iso8601></value></member></struct></value></member><member><name>Name</name><value><string>John Smith</string></value></member><member><name>Wight</name><value><array><data><value><double>66.67</double></value><value><double>100.5</double></value></data></array></value></member></struct></value>",
	},
	{&struct {
		Title  string
		Amount int
		Author string `xml:"author,omitempty"`
	}{
		Title: "War and Piece", Amount: 20,
	}, "<value><struct><member><name>Title</name><value><string>War and Piece</string></value></member><member><name>Amount</name><value><int>20</int></value></member></struct></value>"},
	{&struct {
		Title  string
		Amount int
		Author string `xml:"author,omitempty"`
	}{
		Title: "War and Piece", Amount: 20, Author: "Leo Tolstoy",
	}, "<value><struct><member><name>Title</name><value><string>War and Piece</string></value></member><member><name>Amount</name><value><int>20</int></value></member><member><name>author</name><value><string>Leo Tolstoy</string></value></member></struct></value>"},
	{&struct {
	}{}, "<value><struct></struct></value>"},
	{&struct {
		ID   int    `xml:"id"`
		Name string `xml:"-"`
	}{
		ID: 123, Name: "kolo",
	}, "<value><struct><member><name>id</name><value><int>123</int></value></member></struct></value>"},
}

func TestEncoder_Encode(t *testing.T) {
	enc := Encoder{
		sortMapKeys: true,
	}
	for _, tc := range encodeTestcases {
		b, err := enc.marshal(tc.value)
		require.NoError(t, err)
		require.Equal(t, tc.xml, string(b))
	}
}

type book struct {
	Title  string
	Amount int
}

type unexportedBook struct {
	title  string
	amount int
}

var decodeTestcases = []struct {
	desc     string
	expected any
	ptr      any
	input    string
}{
	// int, i4, i8
	{"int - default", 0, new(*int), "<value><int></int></value>"},
	{"int", 100, new(*int), "<value><int>100</int></value>"},
	{"int - i4", 314159, new(*int), "<value><i4>314159</i4></value>"},
	{"int - i8", int64(31415926), new(*int64), "<value><i8>31415926</i8></value>"},

	// string
	{"string", "Once upon a time", new(*string), "<value><string>Once upon a time</string></value>"},
	{"string - escape special characters", "Mick & Mike <London, UK>", new(*string), "<value><string>Mick &amp; Mike &lt;London, UK&gt;</string></value>"},
	{"string - UTF-8", "When rabbit rules the world ,.<>?/;':\"[]{}\\~!@#$%^&*()_+|", new(*string), "<value>When rabbit rules the world ,.&lt;&gt;?/;&apos;:&quot;[]{}\\~!@#$%^&amp;*()_+|</value>"},

	// base64
	{"base64", "T25jZSB1cG9uIGEgdGltZQ==", new(*string), "<value><base64>T25jZSB1cG9uIGEgdGltZQ==</base64></value>"},

	// boolean
	{"boolean - true", true, new(*bool), "<value><boolean>1</boolean></value>"},
	{"boolean - false", false, new(*bool), "<value><boolean>0</boolean></value>"},

	// double
	{"", 12.34, new(*float32), "<value><double>12.34</double></value>"},
	{"", -9999.83, new(*float32), "<value><double>-9999.83</double></value>"},

	// datetime.iso8601
	{"datetime.iso8601 - 1", parseTime("2023-12-09T21:00:12Z"), new(*time.Time), "<value><dateTime.iso8601>20231209T21:00:12</dateTime.iso8601></value>"},
	{"datetime.iso8601 - 2", parseTime("2023-12-09T21:00:12Z"), new(*time.Time), "<value><dateTime.iso8601>20231209T21:00:12Z</dateTime.iso8601></value>"},
	{"datetime.iso8601 - 3", parseTime("2023-12-09T21:00:12-01:00"), new(*time.Time), "<value><dateTime.iso8601>20231209T21:00:12-01:00</dateTime.iso8601></value>"},
	{"datetime.iso8601 - 4", parseTime("2023-12-09T21:00:12+01:00"), new(*time.Time), "<value><dateTime.iso8601>20231209T21:00:12+01:00</dateTime.iso8601></value>"},
	{"datetime.iso8601 - 5", parseTime("2023-12-09T21:00:12Z"), new(*time.Time), "<value><dateTime.iso8601>2023-12-09T21:00:12</dateTime.iso8601></value>"},
	{"datetime.iso8601 - 6", parseTime("2023-12-09T21:00:12Z"), new(*time.Time), "<value><dateTime.iso8601>2023-12-09T21:00:12Z</dateTime.iso8601></value>"},
	{"datetime.iso8601 - 7", parseTime("2023-12-09T21:00:12-01:00"), new(*time.Time), "<value><dateTime.iso8601>2023-12-09T21:00:12-01:00</dateTime.iso8601></value>"},
	{"datetime.iso8601 - 8", parseTime("2023-12-09T21:00:12+01:00"), new(*time.Time), "<value><dateTime.iso8601>2023-12-09T21:00:12+01:00</dateTime.iso8601></value>"},

	// array
	{"int slice", []int{1, 2, 3}, new(*[]int), "<value><array><data><value><int>1</int></value><value><int>2</int></value><value><int>3</int></value></data></array></value>"},
	{"string slice", []any{"Hello", "World"}, new(any), "<value><array><data><value><string>Hello</string></value><value><string>World</string></value></data></array></value>"},
	{"any slice", []any{"A", int64(1)}, new(any), "<value><array><data><value><string>A</string></value><value><int>1</int></value></data></array></value>"},

	// struct
	{"exported struct", book{"War and Peace", 9973}, new(*book), "<value><struct><member><name>Title</name><value><string>War and Peace</string></value></member><member><name>Amount</name><value><int>9973</int></value></member></struct></value>"},
	{"unexported struct", unexportedBook{}, new(*unexportedBook), "<value><struct><member><name>title</name><value><string>War and Piece</string></value></member><member><name>amount</name><value><int>20</int></value></member></struct></value>"},
	{"map", map[string]string{"Name": "John Smith"}, new(any), "<value><struct><member><name>Name</name><value><string>John Smith</string></value></member></struct></value>"},
	{"empty map", map[string]any{}, new(any), "<value><struct></struct></value>"},
}

func parseTime(s string) time.Time {
	t0, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t0
}

func TestDecoder_Decode(t *testing.T) {
	var dec Decoder
	for _, tc := range decodeTestcases {
		t.Run(tc.desc, func(t *testing.T) {
			r := bytes.NewReader([]byte(tc.input))

			obj := reflect.New(reflect.TypeOf(tc.expected))
			err := dec.Decode(r, obj.Interface())
			require.NoError(t, err)

			obj = obj.Elem()
			if obj.Kind() == reflect.Slice {
				vv := reflect.ValueOf(tc.expected)
				require.Equal(t, vv.Len(), obj.Len())

				N := obj.Len()
				for i := 0; i < N; i++ {
					require.Equal(t, vv.Index(i).Interface(), obj.Index(i).Interface())
				}
			} else {
				require.Equal(t, tc.expected, obj.Interface())
			}
		})
	}
}

func TestDecoder_Decode_pointer(t *testing.T) {
	var dec Decoder
	for _, tc := range decodeTestcases {
		t.Run(tc.desc, func(t *testing.T) {
			r := bytes.NewReader([]byte(tc.input))

			err := dec.Decode(r, tc.ptr)
			require.NoError(t, err)
		})
	}
}

func TestDecoder_Decode_typeMismatch(t *testing.T) {
	var (
		input = "<value><int>100</int></value>"
		dec   Decoder
		s     string
		err   error
	)
	err = dec.Decode(bytes.NewReader([]byte(input)), &s)
	require.Error(t, err)
	var typeMismatchError *TypeMismatchError
	ok := errors.As(err, &typeMismatchError)
	require.True(t, ok)
}

func TestDecoder_Decode_emptyValue(t *testing.T) {
	var (
		dec Decoder
		x   any
	)
	err := dec.Decode(bytes.NewReader([]byte("<value/>")), &x)
	require.NoError(t, err)
}

const emptyStructXML = `
<value>
  <struct>
  </struct>
</value>
`

func TestDecoder_Decode_emptyStruct(t *testing.T) {
	var (
		dec Decoder
		x   any
	)
	err := dec.Decode(bytes.NewReader([]byte(emptyStructXML)), &x)
	require.NoError(t, err)
	require.True(t, x != nil, "should get nil map")
}

const arrayXML = `
<value>
  <array>
    <data>
      <value><int>9973</int></value>
      <value><boolean>1</boolean></value>
      <value><string>Hello World</string></value>
      <value><string>Extra Value</string></value>
    </data>
  </array>
</value>
`

func TestDecoder_Decode_array(t *testing.T) {
	var (
		v1  int
		v2  bool
		v3  string
		dec Decoder
		err error
		v   = []any{&v1, &v2, &v3}
	)
	err = dec.Decode(bytes.NewReader([]byte(arrayXML)), &v)
	require.NoError(t, err)
	require.Equalf(t, 9973, v1, "want 9973, got %d", v1)
	require.Equalf(t, true, v2, "want true, got %t", v2)
	require.Equalf(t, "Hello World", v3, "want 'Hello World', got '%s'", v3)
	require.Len(t, v, 4, "missing appended value")

	got, ok := v[3].(string)
	require.True(t, ok)
	require.Equalf(t, "Extra Value", got, "want 'Extra Value', got '%s'", got)
}

const emptyValueStructXML = `
<?xml version="1.0" encoding="UTF-8"?>
<methodResponse>
	<params>
		<param>
			<value>
				<struct>
					<member>
						<name>user</name>
						<value><string>Brad Pitt</string></value>
					</member>
					<member>
						<name>token</name>
						<value/>
					</member>
				</struct>
			</value>
		</param>
	</params>
</methodResponse>`

func TestDecoder_Decode_emptyValueStruct(t *testing.T) {
	var (
		dec    Decoder
		result struct {
			User  string `xml:"user"`
			Token string `xml:"token"`
		}
	)
	err := dec.Decode(bytes.NewReader([]byte(emptyValueStructXML)), &result)
	require.NoError(t, err)
	require.Equal(t, "Brad Pitt", result.User)
	require.Empty(t, result.Token)
}
