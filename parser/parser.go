package parser

import (
	"strconv"
)

/*
Types:
* - array
$ - bluk string
: - number
*/
const (
	Array  = "*"
	String = "$"
	Number = ":"
	Error  = "-"
)

/*
EOL - End Of Line
*/
const (
	EOL = "\r\n"
)

// encode the tokens in tmDB serialized byte array
func Encode(tokens any) []byte {

	switch v := tokens.(type) {
	case []string:
		return handleString(v)
	case error:
		return handleError(v)
	default:
		return []byte{}
	}

}

// encode individual token as byte slices
func encode(token any) []byte {
	var prepare string
	var serialized = []byte{}

	switch v := token.(type) {
	case []string:
		prepare = Array + strconv.Itoa(len(v)) + EOL
	case string:
		prepare = String + strconv.Itoa(len(v)) + EOL + v + EOL
	case int:
		prepare = Number + strconv.Itoa(v) + EOL
	case error:
		prepare = Error + v.Error() + EOL
	}

	serialized = append(serialized, []byte(prepare)...)

	return serialized
}
