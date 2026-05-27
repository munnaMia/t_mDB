package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	tcolor "github.com/munnaMia/TColor"
	"github.com/munnaMia/t_mDB/config"
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
		fmt.Print(tcolor.Sprintf(tcolor.BlodWhite, tcolor.None, "\n %s:%d >> ", cnf.IP, cnf.PORT))
		if scanner.Scan() {
			input := scanner.Text()
			tokens := strings.Split(input, " ")
			command, err := extractCommand(tokens)
			if err != nil {
				fmt.Print(tcolor.Sprintf(tcolor.BlodRed, tcolor.None, "%s", err.Error()))
				continue 
			}

			// handle the command
			CommandRegistry[command]()
		}

	}
}

// extract the main command from an input stream.
func extractCommand(tokens []string) (string, error) {
	command := strings.ToUpper(tokens[0])

	// check the command exist or not
	if _, ok := CommandRegistry[command]; ok {
		return command, nil
	}

	return "", errors.New("unknown command.")
}
