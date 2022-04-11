package main

import (
	"fmt"
	"reflect"
	"testing"
)

var TestMessages []Message = []Message{
	{
		Content: []byte("Some_content\nto test"),
		From:    []byte("Myself"),
		To:      []byte("Himself"),
	},
	{
		Content: []byte("Some another content!@#$%&*()+"),
		From:    []byte("Himself"),
		To:      []byte("Myself"),
	},
}

func TestJsonEncoding(t *testing.T) {
	for _, msg := range TestMessages {
		var asStruct []Message
		asBytes := ToJson(msg)
		fmt.Println(asBytes)

		res := FromJson(asBytes, asStruct)
		fmt.Println(res)

		if !reflect.DeepEqual(msg, res) {
			t.Fail()
		}
	}
}

func TestByteEncoding(t *testing.T) {
	for _, msg := range TestMessages {
		var asStruct []Message
		asBytes := ToBytes(msg)
		fmt.Println(asBytes)

		res := FromBytes(asBytes, asStruct)
		fmt.Println(res)

		if !reflect.DeepEqual(msg, res) {
			t.Fail()
		}
	}
}

func TestBase64Encoding(t *testing.T) {
	for _, msg := range TestMessages {
		var asStruct []Message
		asBytes := ToBase64(ToBytes(msg))
		fmt.Println(asBytes)

		res := FromBytes(FromBase64(asBytes), asStruct)
		fmt.Println(res)

		if !reflect.DeepEqual(msg, res) {
			t.Fail()
		}
	}
}

func TestByteStreamEncoding(t *testing.T) {
	var msgByteSlice [][]byte
	for i, msg := range TestMessages {
		msgByteSlice := append(msgByteSlice, ToBytes(msg))
		fmt.Println(msgByteSlice[i])
	}
	msgByteStream := ByteStreamJoin(msgByteSlice, StreamSep)

	res := ByteStreamSplit(msgByteStream, StreamSep)
	fmt.Println(res)

	if !reflect.DeepEqual(msgByteSlice, res) {
		t.Fail()
	}

}
