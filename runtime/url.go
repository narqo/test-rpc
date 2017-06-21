package runtime

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"
)

type Message interface {
	URLUnmarshaler
}

type URLUnmarshaler interface {
	UnmarshalURL(r *URLReader)
}

type URLReader struct {
	query string
	key   string
	value string
}

func NewURLReader(query string) *URLReader {
	return &URLReader{query: query}
}

func (r *URLReader) EOF() bool {
	return len(r.query) == 0
}

func (r *URLReader) Err() error {
	return nil
}

func (r *URLReader) addError() {
	panic("not implemented")
}

func (r *URLReader) Key() string {
	var chunk, key string

	chunk, r.query = split2(r.query, '&')
	key, r.value = split2(chunk, '=')

	var err error
	r.key, err = url.QueryUnescape(key)
	if err != nil {
		r.addError()
	}
	return r.key
}

func (r *URLReader) String() string {
	return r.value
}

func (r *URLReader) parseInt(bits int) int64 {
	ret, err := strconv.ParseInt(r.String(), 10, bits)
	if err != nil {
		r.addError()
	}
	return ret
}

func (r *URLReader) parseUint(bits int) uint64 {
	ret, err := strconv.ParseUint(r.String(), 10, bits)
	if err != nil {
		r.addError()
	}
	return ret
}

func (r *URLReader) parseFloat(bits int) float64 {
	ret, err := strconv.ParseFloat(r.String(), bits)
	if err != nil {
		r.addError()
	}
	return ret
}

func (r *URLReader) Bool() bool {
	switch r.String() {
	case "true":
		return true
	case "1":
		return true
	case "false":
		return false
	case "0":
		return false
	default:
		r.addError()
		return false
	}
}

func (r *URLReader) Int() int     { return int(r.parseInt(64)) }
func (r *URLReader) Int8() int8   { return int8(r.parseInt(8)) }
func (r *URLReader) Int16() int16 { return int16(r.parseInt(16)) }
func (r *URLReader) Int32() int32 { return int32(r.parseInt(32)) }
func (r *URLReader) Int64() int64 { return int64(r.parseInt(64)) }

func (r *URLReader) Uint() uint       { return uint(r.parseUint(64)) }
func (r *URLReader) Uint8() uint8     { return uint8(r.parseUint(8)) }
func (r *URLReader) Uint16() uint16   { return uint16(r.parseUint(16)) }
func (r *URLReader) Uint32() uint32   { return uint32(r.parseUint(32)) }
func (r *URLReader) Uint64() uint64   { return uint64(r.parseUint(64)) }
func (r *URLReader) Float64() float64 { return float64(r.parseFloat(64)) }
func (r *URLReader) Float32() float32 { return float32(r.parseFloat(32)) }

type URLMarshaler interface {
	MarshalURL(w *URLWriter) string
}

type URLWriter struct {
	buf bytes.Buffer
}

func (w *URLWriter) Reset() {
	w.buf.Reset()
}

func (w *URLWriter) Bytes() []byte {
	return w.buf.Bytes()
}

func (w *URLWriter) String() string {
	return w.buf.String()
}

func (w *URLWriter) writeString(key, value string) {
	if w.buf.Len() > 0 {
		w.buf.Grow(len(key) + len(value) + 2)
		w.buf.WriteByte('&')
	} else {
		w.buf.Grow(len(key) + len(value) + 1)
	}

	w.buf.WriteString(key)
	w.buf.WriteByte('=')
	w.buf.WriteString(value)
}

func (w *URLWriter) WriteString(key, value string) {
	w.writeString(key, url.QueryEscape(value))
}

func (w *URLWriter) WriteBool(key string, value bool) {
	if w.buf.Len() > 0 {
		w.buf.Grow(len(key) + 3)
		w.buf.WriteByte('&')
	} else {
		w.buf.Grow(len(key) + 2)
	}
	w.buf.WriteString(key)
	w.buf.WriteByte('=')
	if value {
		w.buf.WriteByte('1')
	} else {
		w.buf.WriteByte('0')
	}
}

func (w *URLWriter) WriteFlag(key string) {
	w.WriteBool(key, true)
}

func (w *URLWriter) WritePairs(pairs ...interface{}) error {
	var key string
	for i := 0; i < len(pairs); i++ {
		if key, _ = pairs[i].(string); key == "" {
			return fmt.Errorf("key must be a string: got %T", pairs[i])
		}
		if i+1 >= len(pairs) {
			w.WriteFlag(key)
			return nil
		}
		i++
		switch value := pairs[i].(type) {
		case string:
			w.WriteString(key, value)
		case bool:
			w.WriteBool(key, value)
		case int:
			w.writeInt(key, int64(value))
		case int8:
			w.writeInt(key, int64(value))
		case int16:
			w.writeInt(key, int64(value))
		case int64:
			w.writeInt(key, value)
		case uint:
			w.writeUint(key, uint64(value))
		case uint8:
			w.writeUint(key, uint64(value))
		case uint16:
			w.writeUint(key, uint64(value))
		case uint64:
			w.writeUint(key, value)
		case fmt.Stringer:
			w.WriteString(key, value.String())
		}
	}
	return nil
}

func (w *URLWriter) writeInt(key string, value int64) {
	w.writeString(key, strconv.FormatInt(value, 10))
}

func (w *URLWriter) writeUint(key string, value uint64) {
	w.writeString(key, strconv.FormatUint(value, 10))
}

func (w *URLWriter) WriteInt(key string, value int)     { w.writeInt(key, int64(value)) }
func (w *URLWriter) WriteInt8(key string, value int8)   { w.writeInt(key, int64(value)) }
func (w *URLWriter) WriteInt16(key string, value int16) { w.writeInt(key, int64(value)) }
func (w *URLWriter) WriteInt32(key string, value int32) { w.writeInt(key, int64(value)) }
func (w *URLWriter) WriteInt64(key string, value int64) { w.writeInt(key, value) }

func (w *URLWriter) WriteUint(key string, value uint)     { w.writeUint(key, uint64(value)) }
func (w *URLWriter) WriteUint8(key string, value uint8)   { w.writeUint(key, uint64(value)) }
func (w *URLWriter) WriteUint16(key string, value uint16) { w.writeUint(key, uint64(value)) }
func (w *URLWriter) WriteUint32(key string, value uint32) { w.writeUint(key, uint64(value)) }
func (w *URLWriter) WriteUint64(key string, value uint64) { w.writeUint(key, value) }
