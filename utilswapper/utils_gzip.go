package utilswapper

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
)

func GzlibUncompress(compressSrc string) (string, error) {
	compressedData, _ := base64.StdEncoding.DecodeString(compressSrc)
	b := bytes.NewReader([]byte(compressedData))
	var out bytes.Buffer
	r, err := gzip.NewReader(b)
	if err != nil {
		return "", err
	}
	io.Copy(&out, r)
	return string(out.Bytes()), nil
}
func GzlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := gzip.NewWriter(&in)
	w.Write(src)
	w.Close()
	return ([]byte)(base64.StdEncoding.EncodeToString(in.Bytes()))
}

/*
func GzlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := gzip.NewWriter(&in)
	w.Write(src)
	w.Close()
	return ([]byte)(base64.StdEncoding.EncodeToString(in.Bytes()))
}*/
