package xmlrpc

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidXML = errors.New("invalid XML")

type TypeMismatchError struct {
	source, target string
}

func (e TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch: cannot decode %s to %s", e.source, e.target)
}

const (
	TimeLayoutISO8601        = "20060102T15:04:05"
	TimeLayoutISO8601Z       = "20060102T15:04:05Z07:00"
	TimeLayoutISO8601Hyphen  = "2006-01-02T15:04:05"
	TimeLayoutISO8601HyphenZ = "2006-01-02T15:04:05Z07:00"
	debug                    = false
)

// Base64 represents string values in base64 encoding.
type Base64 string

type Encoder struct {
	indentPrefix string
	indentValue  string
	sortMapKeys  bool
}

func (enc *Encoder) Encode(w io.Writer, methodName string, params any) error {
	if methodName == "" {
		return errors.New("empty method name")
	}

	buf := bufio.NewWriter(w)
	ignoreWriteError(buf, xml.Header)
	ignoreWriteError(buf, fmt.Sprintf("<methodCall><methodName>%s</methodName>", methodName))

	var (
		args []any
		ok   bool
	)
	if args, ok = params.([]any); !ok {
		if params != nil {
			args = []any{params}
		}
	}

	if args != nil {
		ignoreWriteError(buf, "<params>")
		for _, arg := range args {
			b, err := enc.marshal(arg)
			if err != nil {
				return err
			}

			ignoreWriteError(buf, fmt.Sprintf("<param>%s</param>", b))
		}
		ignoreWriteError(buf, "</params>")
	}

	ignoreWriteError(buf, "</methodCall>")

	return buf.Flush()
}

func (enc *Encoder) marshal(v any) ([]byte, error) {
	if v == nil {
		return nil, nil
	}

	return enc.encodeValue(reflect.ValueOf(v))
}

func (enc *Encoder) encodeValue(v reflect.Value) ([]byte, error) {
	var (
		p   []byte
		err error
	)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return []byte("<value/>"), nil
		}

		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		if t, ok := v.Interface().(time.Time); ok {
			p = []byte(fmt.Sprintf("<dateTime.iso8601>%s</dateTime.iso8601>", t.Format(TimeLayoutISO8601)))
		} else {
			p, err = enc.encodeStruct(v)
		}
	case reflect.Map:
		p, err = enc.encodeMap(v)
	case reflect.Slice:
		p, err = enc.encodeSlice(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p = []byte(fmt.Sprintf("<int>%s</int>", strconv.FormatInt(v.Int(), 10)))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		p = []byte(fmt.Sprintf("<i4>%s</i4>", strconv.FormatUint(v.Uint(), 10)))
	case reflect.Float32, reflect.Float64:
		p = []byte(fmt.Sprintf("<double>%s</double>", strconv.FormatFloat(v.Float(), 'f', -1, v.Type().Bits())))
	case reflect.String:
		var buf bytes.Buffer
		xml.Escape(&buf, []byte(v.String()))
		if _, ok := v.Interface().(Base64); ok {
			p = []byte(fmt.Sprintf("<base64>%s</base64>", buf.String()))
		} else {
			p = []byte(fmt.Sprintf("<string>%s</string>", buf.String()))
		}
	case reflect.Bool:
		if v.Bool() {
			p = []byte("<boolean>1</boolean>")
		} else {
			p = []byte("<boolean>0</boolean>")
		}
	default:
		return nil, errors.New("unsupported type for encoding")
	}
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("<value>%s</value>", p)), nil
}

func (enc *Encoder) encodeStruct(v reflect.Value) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("<struct>")

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)

		name := ft.Tag.Get("xml")
		// Skip ignored field.
		if name == "-" {
			continue
		}

		// If omitempty and field value is zero, skip it.
		if (strings.Contains(name, ",omitempty,") || strings.HasSuffix(name, ",omitempty")) && fv.IsZero() {
			continue
		}
		parts := strings.SplitN(name, ",", 2)
		name = parts[0]
		if name == "" {
			// Use the field name instead.
			name = ft.Name
		}

		data, err := enc.encodeValue(fv)
		if err != nil {
			return nil, err
		}

		b.WriteString("<member>")
		b.WriteString(fmt.Sprintf("<name>%s</name>", name))
		b.Write(data)
		b.WriteString("</member>")
	}

	b.WriteString("</struct>")

	return b.Bytes(), nil
}

func (enc *Encoder) encodeMap(v reflect.Value) ([]byte, error) {
	var (
		t = v.Type()
		b bytes.Buffer
	)
	if t.Key().Kind() != reflect.String {
		return nil, errors.New("cannot encode none string map keys")
	}

	b.WriteString("<struct>")

	keys := v.MapKeys()
	if enc.sortMapKeys {
		sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
	}

	N := v.Len()
	for i := 0; i < N; i++ {
		key := keys[i]
		b.WriteString("<member>")
		b.WriteString(fmt.Sprintf("<name>%s</name>", key.String()))

		value := v.MapIndex(key)
		p, err := enc.encodeValue(value)
		if err != nil {
			return nil, err
		}
		b.Write(p)
		b.WriteString("</member>")
	}

	b.WriteString("</struct>")

	return b.Bytes(), nil
}

func (enc *Encoder) encodeSlice(v reflect.Value) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("<array><data>")

	for i := 0; i < v.Len(); i++ {
		data, err := enc.encodeValue(v.Index(i))
		if err != nil {
			return nil, err
		}

		b.Write(data)
	}

	b.WriteString("</data></array>")

	return b.Bytes(), nil
}

type Decoder struct {
	*xml.Decoder
	// CharsetReader, if non-nil, defines a function to generate
	// charset-conversion readers, converting from the provided
	// non-UTF-8 charset into UTF-8. If CharsetReader is nil or
	// returns an error, parsing stops with an error. One of the
	// CharsetReader's result values must be non-nil.
	CharsetReader func(charset string, input io.Reader) (io.Reader, error)
}

func (dec *Decoder) Decode(r io.Reader, v any) error {
	dec.Decoder = xml.NewDecoder(r)

	var (
		token xml.Token
		err   error
	)
	for {
		if token, err = dec.Token(); err != nil {
			return err
		}

		if elem, ok := token.(xml.StartElement); ok {
			if elem.Name.Local == "value" {
				val := reflect.ValueOf(v)
				if val.Kind() != reflect.Ptr {
					return errors.New("cannot pass non-pointer value to decode")
				}

				if err = dec.decodeValue(val.Elem()); err != nil {
					return err
				}
				break
			}
		}
	}

	// Read until the end of document.
	if err = dec.Skip(); err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return nil
}

func (dec *Decoder) decodeValue(val reflect.Value) error {
	var (
		token    xml.Token
		typeName string
		err      error
	)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		val = val.Elem()
	}

	for {
		if token, err = dec.Token(); err != nil {
			return err
		}

		if elem, ok := token.(xml.EndElement); ok {
			if elem.Name.Local == "value" {
				return nil
			}
			return ErrInvalidXML
		}

		if elem, ok := token.(xml.StartElement); ok {
			typeName = elem.Name.Local
			break
		}

		// Treat xml.CharData without type identifier as string.
		if chars, ok := token.(xml.CharData); ok {
			if value := strings.TrimSpace(string(chars)); value != "" {
				if err = checkTypeMatch(val, reflect.String); err != nil {
					return err
				}

				val.SetString(value)
				return nil
			}
		}
	}

	switch typeName {
	case "struct":
		/*
		   <struct>
		      <member>
		         <name>name</name>
		         <value><string>Pitt</string></value>
		      </member>
		      <member>
		         <name>age</name>
		         <value><int>28</int></value>
		      </member>
		   </struct>
		*/
		var (
			fields     = make(map[string]reflect.Value)
			mapPtr     = val
			structType = val.Type()
			isMap      bool
			dummy      map[string]any
		)
		if err = checkTypeMatch(val, reflect.Struct); err != nil {
			if checkTypeMatch(val, reflect.Map) == nil {
				if structType.Key().Kind() != reflect.String {
					return fmt.Errorf("unsupported map key type %q: can only decode map of string key", structType.Key().Kind())
				}
				isMap = true
			} else if checkTypeMatch(val, reflect.Interface) == nil && val.IsNil() {
				structType = reflect.TypeOf(dummy)
				mapPtr = reflect.New(structType).Elem()
				val.Set(mapPtr)
				isMap = true
			} else {
				return err
			}
		}

		if !isMap {
			for i := 0; i < structType.NumField(); i++ {
				key := structType.Field(i)
				value := val.FieldByName(key.Name)

				if value.CanSet() {
					name := key.Tag.Get("xml")
					parts := strings.SplitN(name, ",", 2)
					name = parts[0]
					if name == "-" {
						continue
					}
					if name == "" {
						name = key.Name
					}
					fields[name] = value
				}
			}
		} else {
			// Create initial empty map.
			mapPtr.Set(reflect.MakeMap(structType))
		}

	StructLoop:
		for {
			if token, err = dec.Token(); err != nil {
				return err
			}
			switch t := token.(type) {
			case xml.StartElement:
				if t.Name.Local != "member" {
					return ErrInvalidXML
				}

				tagName, key, err := dec.readTag()
				if err != nil {
					return err
				}
				if tagName != "name" {
					return ErrInvalidXML
				}

				var (
					value reflect.Value
					ok    = true
				)
				if !isMap {
					value, ok = fields[string(key)]
				} else {
					value = reflect.New(structType.Elem())
				}

				if ok {
					for {
						if token, err = dec.Token(); err != nil {
							return err
						}

						if elem, ok := token.(xml.StartElement); ok && elem.Name.Local == "value" {
							if err = dec.decodeValue(value); err != nil {
								return err
							}

							if err = dec.Skip(); err != nil {
								return err
							}
							break
						}
					}
				}

				if err = dec.Skip(); err != nil {
					return err
				}

				if isMap {
					mapPtr.SetMapIndex(reflect.ValueOf(string(key)), reflect.Indirect(value))
					val.Set(mapPtr)
				}
			case xml.EndElement:
				break StructLoop
			}
		}
	case "array":
		/*
		   <array>
		      <data>
		         <value><string>Are</string></value>
		         <value><string>you</string></value>
		         <value><string>OK?</string></value>
		      </data>
		   </array>
		*/
		slice := val
		if checkTypeMatch(val, reflect.Interface) == nil && val.IsNil() {
			slice = reflect.ValueOf([]any{})
		} else if err = checkTypeMatch(val, reflect.Slice); err != nil {
			return err
		}

	ArrayLoop:
		for {
			if token, err = dec.Token(); err != nil {
				return err
			}

			switch t := token.(type) {
			case xml.StartElement:
				var index int
				if t.Name.Local != "data" {
					return ErrInvalidXML
				}
			DataLoop:
				for {
					if token, err = dec.Token(); err != nil {
						return err
					}

					switch tt := token.(type) {
					case xml.StartElement:
						if tt.Name.Local != "value" {
							return ErrInvalidXML
						}

						if index < slice.Len() {
							elem := slice.Index(index)
							if elem.Kind() == reflect.Interface {
								elem = elem.Elem()
							}
							if elem.Kind() != reflect.Ptr {
								return errors.New("cannot write to non-pointer array element")
							}

							if err = dec.decodeValue(elem); err != nil {
								return err
							}
						} else {
							elem := reflect.New(slice.Type().Elem())
							if err = dec.decodeValue(elem); err != nil {
								return err
							}
							slice = reflect.Append(slice, elem.Elem())
						}

						if err = dec.Skip(); err != nil {
							return err
						}

						index++
					case xml.EndElement:
						val.Set(slice)
						break DataLoop
					}
				}
			case xml.EndElement:
				break ArrayLoop
			}
		}
	default:
		if token, err = dec.Token(); err != nil {
			return err
		}

		var data []byte
		switch t := token.(type) {
		case xml.EndElement:
			return nil
		case xml.CharData:
			data = t.Copy()
		default:
			return ErrInvalidXML
		}

		switch typeName {
		case "int", "i4", "i8":
			if checkTypeMatch(val, reflect.Interface) == nil && val.IsNil() {
				i64, err := strconv.ParseInt(string(data), 10, 64)
				if err != nil {
					return err
				}

				pi64 := reflect.New(reflect.TypeOf(i64)).Elem()
				pi64.SetInt(i64)
				val.Set(pi64)
			} else if err = checkTypeMatch(val, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64); err != nil {
				return err
			} else {
				i64, err := strconv.ParseInt(string(data), 10, val.Type().Bits())
				if err != nil {
					return err
				}

				val.SetInt(i64)
			}
		case "string", "base64":
			s := string(data)
			if checkTypeMatch(val, reflect.Interface) == nil && val.IsNil() {
				ps := reflect.New(reflect.TypeOf(s)).Elem()
				ps.SetString(s)
				val.Set(ps)
			} else if err = checkTypeMatch(val, reflect.String); err != nil {
				return err
			} else {
				val.SetString(s)
			}
		case "dateTime.iso8601":
			var t0 time.Time
			for _, layout := range []string{TimeLayoutISO8601, TimeLayoutISO8601Z, TimeLayoutISO8601Hyphen, TimeLayoutISO8601HyphenZ} {
				t0, err = time.Parse(layout, string(data))
				if err == nil {
					break
				}
			}
			if err != nil {
				return err
			}

			if checkTypeMatch(val, reflect.Interface) == nil && val.IsNil() {
				pt0 := reflect.New(reflect.TypeOf(t0)).Elem()
				pt0.Set(reflect.ValueOf(t0))
				val.Set(pt0)
			} else if _, ok := val.Interface().(time.Time); !ok {
				return &TypeMismatchError{
					source: val.Elem().String(),
					target: "time",
				}
			} else {
				val.Set(reflect.ValueOf(t0))
			}
		case "boolean":
			var tf bool
			tf, err = strconv.ParseBool(string(data))
			if err != nil {
				return err
			}

			if checkTypeMatch(val, reflect.Interface) == nil && val.IsNil() {
				ptf := reflect.New(reflect.TypeOf(tf)).Elem()
				ptf.SetBool(tf)
				val.Set(ptf)
			} else if err = checkTypeMatch(val, reflect.Bool); err != nil {
				return err
			} else {
				val.SetBool(tf)
			}
		case "double":
			var f float64
			if checkTypeMatch(val, reflect.Interface) == nil && val.IsNil() {
				f, err = strconv.ParseFloat(string(data), 64)
				if err != nil {
					return err
				}

				pf := reflect.New(reflect.TypeOf(f)).Elem()
				pf.SetFloat(f)
				val.Set(pf)
			} else if err = checkTypeMatch(val, reflect.Float32, reflect.Float64); err != nil {
				return err
			} else {
				f, err = strconv.ParseFloat(string(data), val.Type().Bits())
				if err != nil {
					return err
				}

				val.SetFloat(f)
			}
		default:
			return fmt.Errorf("unsupported type: %s", typeName)
		}

		if err = dec.Skip(); err != nil {
			return err
		}
	}

	return nil
}

func (dec *Decoder) readTag() (string, []byte, error) {
	var (
		token xml.Token
		name  string
		err   error
	)
	for {
		if token, err = dec.Token(); err != nil {
			return "", nil, err
		}

		if elem, ok := token.(xml.StartElement); ok {
			name = elem.Name.Local
			break
		}
	}

	chars, err := dec.readCharData()
	if err != nil {
		return "", nil, err
	}

	return name, chars, dec.Skip()
}

func (dec *Decoder) readCharData() ([]byte, error) {
	var (
		token xml.Token
		err   error
	)
	if token, err = dec.Token(); err != nil {
		return nil, err
	}

	if chars, ok := token.(xml.CharData); ok {
		return chars.Copy(), nil
	}

	return nil, ErrInvalidXML
}

func checkTypeMatch(v reflect.Value, kinds ...reflect.Kind) error {
	if len(kinds) == 0 {
		return nil
	}

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	var match bool
	for _, kind := range kinds {
		if v.Kind() == kind {
			match = true
			break
		}
	}

	if !match {
		return &TypeMismatchError{
			source: v.Kind().String(),
			target: kinds[len(kinds)-1].String(),
		}
	}

	return nil
}

func ignoreWriteError(w *bufio.Writer, s string) {
	_, err := w.WriteString(s)
	if debug && err != nil {
		panic(err)
	}
}
