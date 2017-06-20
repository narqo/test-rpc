package runtime

import (
	"reflect"
	"testing"
)

func TestURLReader(t *testing.T) {
	type KV struct {
		Key, Value string
	}

	tests := []struct {
		In  string
		Out []KV
	}{
		{
			"",
			nil,
		},
		{
			"key=value",
			[]KV{{"key", "value"}},
		},
		{
			"key1=value1&key2=value2",
			[]KV{{"key1", "value1"}, {"key2", "value2"}},
		},
		{
			"key1=1&key2=",
			[]KV{{"key1", "1"}, {"key2", ""}},
		},
	}

	for n, test := range tests {
		r := URLReader{query: test.In}

		var values []KV
		for !r.EOF() {
			values = append(values, KV{r.Key(), r.String()})
		}
		if !reflect.DeepEqual(values, test.Out) {
			t.Errorf("case %v, got %v; want %v", n, values, test.Out)
		}
	}
}

func TestURLWriter(t *testing.T) {
	type KV struct {
		Key, Value string
	}

	tests := []struct {
		In  []string
		Out string
	}{
		{
			nil,
			"",
		},
		{
			[]string{"key", "value"},
			"key=value",
		},
		{
			[]string{"key1", "value1", "key2", "value2"},
			"key1=value1&key2=value2",
		},
		{
			[]string{"key1", "val ue1"},
			"key1=val+ue1",
		},
	}

	for n, test := range tests {
		w := URLWriter{}
		for i := 0; i < len(test.In); i += 2 {
			key, value := test.In[i], test.In[i+1]
			w.WriteString(key, value)
		}
		res := w.String()
		if res != test.Out {
			t.Errorf("case %v, got %v; want %v", n, res, test.Out)
		}
	}
}

func TestURLWriter_WriteBool(t *testing.T) {
	w := URLWriter{}
	w.WriteBool("key1", true)
	w.WriteBool("key2", false)
	res := w.String()
	if res != "key1=1&key2=0" {
		t.Errorf("WriteBool: got %v, want key1=1&key2=0", res)
	}
}

func TestURLWriter_WriteFlag(t *testing.T) {
	w := URLWriter{}
	w.WriteFlag("key1")
	w.WriteFlag("key2")
	res := w.String()
	if res != "key1=1&key2=1" {
		t.Errorf("WriteFlag: got %v, want key1=1&key2=1", res)
	}
}

func TestURLWriter_WriteXIntx(t *testing.T) {
	w := URLWriter{}
	w.WriteInt("key1", 64)
	w.WriteInt8("key2", 8)
	w.WriteInt64("key3", 64)
	res := w.String()
	expect := "key1=64&key2=8&key3=64"
	if res != expect {
		t.Errorf("WriteFlag: got %v, want %v", res, expect)
	}
}
