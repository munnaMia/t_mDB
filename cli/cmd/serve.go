package cmd

import (
	"fmt"
	"time"

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
	tcolor.Println(tcolor.BlodYellow, tcolor.None, "t_mDB CLI started")

	for {
		fmt.Print(tcolor.Sprintf(tcolor.Blue, tcolor.None, "\n %s:%d >>", cnf.IP, cnf.PORT))

		time.Sleep(time.Second*2)
	}
}
