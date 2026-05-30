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
)

/*
EOL - End Of Line
*/
const (
	EOL = "\r\n"
)

// encode the tokens in tmDB serialized byte array
func Encode(tokens []string) []byte {
	var serializedTokens = []byte{}

	// parsed the array first
	serializedTokens = append(serializedTokens, encode(tokens)...)

	for _, token := range tokens {
		parsedToken := encode(token)
		serializedTokens = append(serializedTokens, parsedToken...)
	}
	return serializedTokens
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
	}

	serialized = append(serialized, []byte(prepare)...)

	return serialized
}
