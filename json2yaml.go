package encodingtools

/*
 Base package for tools used to manipulate data in different format : json, yaml, binary (from protobuf def)
*/
import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

// JSON to YAML Converter
// j2yConvert : internal function, used by json2Yaml
func j2yConvert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[string]interface{}:
		m2 := map[interface{}]interface{}{}
		for k, v := range x {
			m2[j2yConvert(k)] = j2yConvert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = j2yConvert(v)
		}
		return x
	case string:
		return i
	case bool:
		return i
	case float64:
		return i
	default:
		fmt.Fprintf(os.Stderr, "j2yConvert : Unmatched type %v\n", reflect.TypeOf(i))
		return i
	}
}

// JSON2Yaml Convert any JSON formatted data-block to YAML data-block
func JSON2Yaml(buf []byte) ([]byte, error) {
	var body interface{}
	if err := json.Unmarshal(buf, &body); err != nil {
		return nil, err
	}

	body = j2yConvert(body)

	b, err := yaml.Marshal(body)
	return b, err

}
