package utility

import (
	b64 "encoding/base64"
	"strconv"

	"github.com/bjarneh/latinx"
)

func DecodeStr(value string) string {

	encodeValue, _ := b64.StdEncoding.DecodeString(value)
	decodeValue := string(encodeValue)
	return decodeValue
}

func EncodeBase64Str(value string) string {
	encodValue := b64.StdEncoding.EncodeToString([]byte(value))
	return encodValue
}

func ToISO88591(utf8 string) (bool, string) {

	converter := latinx.Get(latinx.ISO_8859_1)

	latin1bytes, _, err := converter.Encode([]byte(utf8))

	if err != nil {
		return false, ""
	}
	return true, string(latin1bytes)
}

func ToUnicode(origin string) string {
	rs := []rune(origin)
	json := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	//fmt.Printf("JSON: %s\n", json)
	return json
}

func FromISO88591(iso88591 string) (bool, string) {

	converter := latinx.Get(latinx.ISO_8859_1)

	utf8bytes, err := converter.Decode([]byte(iso88591))

	if err != nil {
		return false, ""
	}

	return true, string(utf8bytes)
}
