package json_iterator

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"unsafe"
)

//Reference http://jsoniter.com/go-tips.html
type Student struct {
	ID       uint     `json:"id"`
	Age      uint8    `json:"age"`
	Gender   uint8    `json:"gender"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

type Location struct {
	Country  string
	Province string
	City     string
	District string
}

func Marshal() ([]byte, error) {
	s := Student{
		ID:     1,
		Age:    27,
		Gender: 1,
		Name:   "yuchanns",
		Location: Location{
			Country:  "China",
			Province: "Guangdong",
			City:     "Shenzhen",
			District: "Nanshan",
		},
	}

	//json := jsoniter.ConfigCompatibleWithStandardLibrary
	//sjson, err := json.Marshal(s)
	sjson, err := jsoniter.Marshal(s)

	if err == nil {
		fmt.Println(string(sjson))
	}

	return sjson, err
}

type locationAsStringCodec struct{}

func (codec *locationAsStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	lc := *((*Location)(ptr))

	return lc.Country == "" && lc.Province == "" && lc.District == "" && lc.City == ""
}

func (codec *locationAsStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	lc := *((*Location)(ptr))

	stream.WriteString(strings.Join([]string{lc.Country, lc.Province, lc.City, lc.District}, " "))
}

func RegisterEncoder() ([]byte, error) {
	jsoniter.RegisterTypeEncoder("json_iterator.Location", &locationAsStringCodec{})
	s := Student{
		ID:     1,
		Age:    27,
		Gender: 1,
		Name:   "yuchanns",
		Location: Location{
			Country:  "China",
			Province: "Guangdong",
			City:     "Shenzhen",
			District: "Nanshan",
		},
	}
	sjson, err := jsoniter.Marshal(s)

	if err == nil {
		fmt.Println(string(sjson))
	}

	return sjson, err
}
