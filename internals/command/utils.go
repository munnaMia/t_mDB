package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/munnaMia/t_mDB/config"
)

// extract the main command from an input stream.
func ValidateCommand(tokens []string, cnf *config.Config) error {
	command := strings.ToUpper(tokens[0])

	// check the command exist or not
	if _, ok := CommandRegistry[command]; !ok {
		return errors.New("unknown command.")
	}

	// validate commands
	if err := CommandValidation(command, tokens); err != nil {
		return err
	}

	//validate key length defaile key size 64
	if err := ValidateKeySize(tokens[1], cnf.KEY_SIZE); err != nil {
		return err
	}

	//validate key length defaile key size 1024
	if err := ValidatePayloadSize(tokens[1], cnf.PAYLOAD_SIZE); err != nil {
		return err
	}

	return nil

}

// validate an user given token length
func CommandValidation(cmd string, tokens []string) error {
	if CommandSize[cmd] != len(tokens) {
		return errors.New("Unwanted arguments.")
	}
	return nil
}

// validate key size
func ValidateKeySize(key string, size int) error {
	if len(key) > size {
		return fmt.Errorf("key is too long, max size is : %d", size)
	}
	return nil
}

// validate payload size
func ValidatePayloadSize(payload string, size int) error {
	if len(payload) > size {
		return fmt.Errorf("Payload is too long, max size is : %d", size)
	}
	return nil
}
