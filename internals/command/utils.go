package command

import "errors"

// validate an user given token length
func CommandValidation(cmd string, tokens []string) error {
	if CommandSize[cmd] != len(tokens) {
		return errors.New("Unwanted arguments.")
	}
	return nil
}
