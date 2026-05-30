package parser

/*
	Types:
	* - array
	+ - simple string
	$ - bluk string
	: - number
	- - error
*/

// encode the tokens in tmDB serialized byte array
func Encode(tokens []string) []byte {
	var serializedTokens []byte
	for _, token := range tokens {
		parsedToken := encode(token)
		serializedTokens = append(serializedTokens, parsedToken...)
	}
	return serializedTokens
}

// encode individual token as byte slices
func encode(token string) []byte {
	return []byte{}
}
