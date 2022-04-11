package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

var (
	StreamSep []byte = []byte(";")
)

func FromJson(data []byte, schema interface{}) interface{} {
	err := json.Unmarshal(data, &schema)
	if err != nil {
		fmt.Println(err.Error())
	}
	return schema
}

func ToJson(schema interface{}) []byte {
	bytes, err := json.Marshal(schema)
	if err != nil {
		fmt.Println(err.Error())
	}
	return bytes
}

func ToBytes(schema interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(schema)
	if err != nil {
		fmt.Println(err.Error())
	}
	return buf.Bytes()
}

func FromBytes(data []byte, schema interface{}) interface{} {
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&schema)
	if err != nil {
		fmt.Println(err.Error())
	}
	return schema
}

func ToBase64(data []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	return dst
}

func FromBase64(data []byte) []byte {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dst, data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return dst[:n]
}

func ByteStreamJoin(data [][]byte, sep []byte) []byte {
	return bytes.Join(data, sep)
}

func ByteStreamSplit(stream []byte, sep []byte) [][]byte {
	return bytes.Split(stream, sep)
}
