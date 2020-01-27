package newbee_traps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func JsonEncoderAddNewline() (string, string) {
	data := map[string]int{"key": 1}

	var b bytes.Buffer
	_ = json.NewEncoder(&b).Encode(data)

	raw, _ := json.Marshal(data)

	byteString := b.String()
	rawString := string(raw)

	fmt.Println(byteString)
	fmt.Println(rawString)

	return byteString, rawString
}

func JsonEscapeHTML() (string, string, string) {
	data := "x < y"

	raw, _ := json.Marshal(data)
	fmt.Println(string(raw))

	var b1 bytes.Buffer
	_ = json.NewEncoder(&b1).Encode(data)
	fmt.Println(b1.String())

	var b2 bytes.Buffer
	enc := json.NewEncoder(&b2)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(data)
	fmt.Println(b2.String())

	return string(raw), b1.String(), b2.String()
}

func JsonUnmarshalNumberic(data []byte) (uint64, int64, uint64) {
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	var status1 = uint64(result["status"].(float64)) // 第一种方法，先转成uint64再使用

	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}

	var status2, _ = result["status"].(json.Number).Int64() // 第二种方法，使用Decoder明确指定数字类型

	var resultS struct {
		Status uint64 `json:"status"`
	}

	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&resultS); err != nil {
		log.Fatalln(err)
	}

	var status3 = resultS.Status // 第三种方法，使用结构体

	return status1, status2, status3
}

func JsonUnmarshalUncertainType(records [][]byte) {
	for _, record := range records {
		var result struct {
			StatusCode uint64          `json:"-"`
			StatusName string          `json:"-"`
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
			log.Fatalln(err)
		}

		var name string
		var code uint64
		if err := json.Unmarshal(result.Status, &name); err == nil {
			result.StatusName = name
		} else if err := json.Unmarshal(result.Status, &code); err == nil {
			result.StatusCode = code
		}

		fmt.Printf("result => %+v\n", result)
	}
}

func HiddenCapacityInSlice() ([]byte, []byte, []byte, []byte) {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	rawNew := raw[:3]
	fmt.Println(len(rawNew), cap(rawNew), &rawNew[0])
	rawCopy := make([]byte, 3)
	copy(rawCopy, raw)
	fmt.Println(len(rawCopy), cap(rawCopy), &rawCopy[0])
	rawFull := raw[:3:3]
	fmt.Println(len(rawFull), cap(rawFull), &rawFull[0])
	return raw, rawNew, rawCopy, rawFull
}

func DeferredExcutionTime() {
	a := []int{1, 2, 3}
	for _, v := range a {
		func(v int) {
			fmt.Println(v)
			defer fmt.Println("defer execution")
		}(v)
	}
}
