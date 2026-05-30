package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

	// dialing a tcp connection.
	conn, err := net.Dial("tcp", cnf.IP+":"+strconv.Itoa(cnf.PORT))
	if err != nil {
		fmt.Println(cnf.PORT)
		log.Fatalln(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	tcolor.Println(tcolor.BlodWhite, tcolor.BgBlue, "\n t_mDB CLI started ")

	for {
		util.Print(cnf.IP + ":" + strconv.Itoa(cnf.PORT))
		if scanner.Scan() {
			input := scanner.Text()
			tokens := strings.Split(input, " ")

			if err := cmd.ValidateCommand(tokens, cnf); err != nil {
				util.PrintError(string(parser.Encode(err)))
				continue
			}

			serializedEncodedTokes := parser.Encode(tokens)

			sendData(conn, serializedEncodedTokes)
		}

	}
}

// sending the data through the TCP.
func sendData(conn net.Conn, data []byte) {
	defer conn.Close()

	_, err := conn.Write(data)
	if err != nil {
		util.PrintError(string(parser.Encode(err)))
	}
}
