package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	tcolor "github.com/munnaMia/TColor"
	"github.com/munnaMia/t_mDB/config"
	cmd "github.com/munnaMia/t_mDB/internals/command"
	"github.com/munnaMia/t_mDB/internals/util"
	"github.com/munnaMia/t_mDB/parser"
)

func Run() {
	// load configurations
	cnf := config.GetConfig()

	// _, err := net.Dial("tcp", strconv.Itoa(cnf.PORT))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	scanner := bufio.NewScanner(os.Stdin)

	tcolor.Println(tcolor.BlodWhite, tcolor.BgBlue, "\n t_mDB CLI started ")

	for {
		util.Print(cnf.IP + ":" + strconv.Itoa(cnf.PORT))
		if scanner.Scan() {
			input := scanner.Text()
			tokens := strings.Split(input, " ")
			command, err := extractCommand(tokens)
			if err != nil {
				util.PrintError(err)
				continue
			}

			if err = cmd.CommandValidation(command, tokens); err != nil {
				util.PrintError(err)
				continue
			}

			serializedEncodedTokes := parser.Encode(tokens)
			for _, v := range serializedEncodedTokes {
				fmt.Printf("%s", string(v))
			}
			// send seriallized tokens to backend.
		}

	}
}

// extract the main command from an input stream.
func extractCommand(tokens []string) (string, error) {
	command := strings.ToUpper(tokens[0])

	// check the command exist or not
	if _, ok := cmd.CommandRegistry[command]; ok {
		return command, nil
	}

	return "", errors.New("unknown command.")
}
