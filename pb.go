package encodingtools

/*
 Base package for tools used to manipulate data in different format : json, yaml, binary (from protobuf def)
*/
import (
	"fmt"
	"io/ioutil"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	// EncodingTypeJSON : JSON format
	EncodingTypeJSON EncodingType = iota

	// EncodingTypeYaml : Yaml format
	EncodingTypeYaml

	// EncodingTypeBinaryPB : Protocol buffer binary encoded format
	EncodingTypeBinaryPB
)

// EncodingType format code
type EncodingType int

// EncodingTypeFromString Convert string JSON, Yaml, or any to EncodingType Code
func EncodingTypeFromString(str string) EncodingType {
	switch strings.ToUpper(str) {
	case "JSON":
		return EncodingTypeJSON
	case "YAML":
		return EncodingTypeYaml
	case "YML":
		return EncodingTypeYaml
	default:
		return EncodingTypeBinaryPB
	}
}

// Bytes2Object Convert object which has proto, json, yaml interface to data
func Bytes2Object(obj proto.Message, data []byte, intype EncodingType) error {
	var err error
	switch intype {
	case EncodingTypeJSON:
		err = protojson.Unmarshal(data, obj)
	case EncodingTypeYaml:
		var jsonFile []byte
		jsonFile, err = Yaml2JSON(data)
		if err != nil {
			return err
		}
		err = protojson.Unmarshal(jsonFile, obj)
	case EncodingTypeBinaryPB:
		err = proto.Unmarshal(data, obj)
	default:
		err = fmt.Errorf("Bytes2Object: Unknown source type %d", intype)
	}
	return err
}

// Objet2Bytes marshal an object from type proto.Message to byte array depending on EncodingType code
func Objet2Bytes(obj proto.Message, outtype EncodingType) ([]byte, error) {
	switch outtype {
	case EncodingTypeJSON:
		marshaller := protojson.MarshalOptions{ EmitUnpopulated: true}
		return marshaller.Marshal(obj)
	case EncodingTypeYaml:
		data, err := protojson.Marshal(obj)
		if err != nil {
			return nil, err
		}
		return JSON2Yaml(data)
	case EncodingTypeBinaryPB:
		return proto.Marshal(obj)
	default:
		return nil, fmt.Errorf("Objet2Bytes: Unknown destination type %d", outtype)
	}
}

// Load : this function load a file and unmarshall file into obj. File can have different formats depending on its extension
// .json -> json file
// .yaml, .yml -> yaml file
// any -> binary protobuf encoded
func Load(filename string, obj proto.Message) error {
	var data []byte
	var err error

	sl := strings.Split(filename, ".")
	intype := sl[len(sl)-1]

	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = Bytes2Object(obj, data, EncodingTypeFromString(intype))
	if err != nil {
		return err
	}
	return nil
}
