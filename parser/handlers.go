package parser

// handle string type massages
func handleString(tokens []string) []byte {
	var serializedTokens []byte
	// parsed the array first
	serializedTokens = append(serializedTokens, encode(tokens)...)

	for _, token := range tokens {
		parsedToken := encode(token)
		serializedTokens = append(serializedTokens, parsedToken...)
	}

	return serializedTokens
}

// handle error type massage
func handleError(err error) []byte {
	return encode(err)
}
