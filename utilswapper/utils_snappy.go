package utilswapper

import (
	"github.com/golang/snappy"
)

func SnappyEncode(src []byte) (encode []byte) {
	encode = snappy.Encode(nil, src)
	return encode
}
func SnappyDecode(src []byte) (decode []byte, err error) {
	decode, err = snappy.Decode(nil, src)
	return
}
func SnappyExample() {
	marshal := "123"
	encode := SnappyEncode(([]byte)(marshal))
	SnappyDecode(encode)
}
