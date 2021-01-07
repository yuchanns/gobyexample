package common

import "encoding/json"

func MustConvert(from, to interface{}) {
	jsonByte, err := json.Marshal(from)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(jsonByte, to); err != nil {
		panic(err)
	}
}
