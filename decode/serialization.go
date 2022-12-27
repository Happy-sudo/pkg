package decode

import (
	"bytes"
	"encoding/json"
	"reflect"
)

// ECoding 编码
func ECoding(str interface{}) ([]byte, error) {
	return json.Marshal(str)
}

// DeCoding 解码
func DeCoding(byte []byte) (interface{}, error) {
	var data interface{}
	d := json.NewDecoder(bytes.NewReader(byte))
	d.UseNumber()
	err := d.Decode(&data)
	return data, err
}

// DeCodingToMap to interface 解码
func DeCodingToMap(byte []byte) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	d := json.NewDecoder(bytes.NewReader(byte))
	d.UseNumber()
	err := d.Decode(&data)
	return data, err
}

// DeCodingToTwoMap to interface 解码
func DeCodingToTwoMap(byte []byte) (map[string]map[string]interface{}, error) {
	data := make(map[string]map[string]interface{})
	d := json.NewDecoder(bytes.NewReader(byte))
	d.UseNumber()
	err := d.Decode(&data)
	return data, err
}

// CoventArray 对象转数组
func CoventArray(value interface{}) []string {
	var prop []string
	refType := reflect.TypeOf(value)
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		if field.Anonymous {
			prop = append(prop, field.Name)
			for j := 0; j < field.Type.NumField(); j++ {
				prop = append(prop, field.Type.Field(j).Name)
			}
			continue
		}
		prop = append(prop, field.Name)
	}
	return prop
}
