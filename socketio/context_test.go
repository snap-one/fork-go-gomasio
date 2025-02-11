package socketio

import (
	"bytes"
	"io"
	"testing"

	"github.com/snap-one/fork-go-gomasio"
)

type testWriterFactory struct {
	w io.Writer
}

func (f *testWriterFactory) NewWriter() gomasio.WriteFlusher {
	return gomasio.NopFlusher(f.w)
}

func TestContext_Emit(t *testing.T) {
	ts := []struct {
		event    string
		args     []interface{}
		expected string
	}{
		{
			event:    "hello",
			args:     nil,
			expected: `2["hello"]` + "\n",
		},
		{
			event:    "string",
			args:     []interface{}{"hoge"},
			expected: `2["string","hoge"]` + "\n",
		},
		{
			event:    "number",
			args:     []interface{}{1},
			expected: `2["number",1]` + "\n",
		},
		{
			event: "custom",
			args: []interface{}{&struct {
				Id  int    `json:"id"`
				Msg string `json:"msg"`
			}{
				Id:  15,
				Msg: "hello",
			}},
			expected: `2["custom",{"id":15,"msg":"hello"}]` + "\n",
		},
	}
	for _, tc := range ts {
		var b bytes.Buffer
		ctx, err := NewContext(&testWriterFactory{&b}, &Packet{})
		if err != nil {
			t.Error(err)
			continue
		}
		if err := ctx.Emit(tc.event, tc.args...); err != nil {
			t.Error(err)
			continue
		}
		if got := b.String(); got != tc.expected {
			t.Errorf("unexpected emit event. expected: %v, but got: %v", tc.expected, got)
		}
	}
}

func TestContext_Event(t *testing.T) {
	b := new(bytes.Buffer)
	ts := []struct {
		body     string
		expected string
	}{
		{
			body:     `["message"]`,
			expected: "message",
		},
		{
			body:     `["reply",1,2,3,4,5,6,7]`,
			expected: "reply",
		},
	}
	for _, tc := range ts {
		ctx, err := NewContext(&testWriterFactory{b}, &Packet{
			Type: EVENT,
			Body: bytes.NewBufferString(tc.body),
		})
		if err != nil {
			t.Error(err)
			continue
		}
		if got := ctx.Event(); got != tc.expected {
			t.Errorf("unexpected event name. expected: %v, but got: %v", tc.expected, got)
		}
	}

}

func TestContext_Args(t *testing.T) {
	b := new(bytes.Buffer)
	ctx, err := NewContext(&testWriterFactory{b}, &Packet{
		Type: EVENT,
		Body: bytes.NewBufferString(`["sample",1,"test",{"dict":1}]`),
	})
	if err != nil {
		t.Fatal(err)
	}
	var i int
	var s string
	var d map[string]int
	if err := ctx.Args(&i, &s, &d); err != nil {
		t.Fatal(err)
	}
	if i != 1 {
		t.Errorf("unexpected number. expected: 1, but got: %v", i)
	}
	if s != "test" {
		t.Errorf("unexpected string. expected: test, but got: %v", s)
	}
	if got := d["dict"]; got != 1 {
		t.Errorf("unexpected d['dict']. expected: 1, but got: %v", got)
	}
}
