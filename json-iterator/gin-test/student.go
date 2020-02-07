package main

import (
	jsoniter "github.com/json-iterator/go"
	"strings"
	"unsafe"
)

func init() {
	InitRegisterFiledEncoders()
}

func InitRegisterFiledEncoders() {
	jsoniter.RegisterFieldEncoder("main.Student", "Location", &locationAsStringCodec{})
}

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

type locationAsStringCodec struct{}

func (codec *locationAsStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	lc := *((*Location)(ptr))

	return lc.Country == "" && lc.Province == "" && lc.District == "" && lc.City == ""
}

func (codec *locationAsStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	lc := *((*Location)(ptr))

	stream.WriteString(strings.Join([]string{lc.Country, lc.Province, lc.City, lc.District}, " "))
}
