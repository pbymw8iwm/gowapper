package utilswapper
import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
func Marshal(v interface{}) (bt []byte, err error) {
	_, err = json.Marshal(v)
	return
}
