package util

import (
	"fmt"

	tcolor "github.com/munnaMia/TColor"
)

// print any massage as bold-white text
func Print(msg string) {
	fmt.Print(tcolor.Sprintf(tcolor.BlodWhite, tcolor.None, "\n %s >> ", msg))
}

// print error massage as bold-red text
func PrintError(err error) {
	fmt.Print(tcolor.Sprintf(tcolor.BlodRed, tcolor.None, "%s", err.Error()))
}
