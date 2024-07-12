package main

import (
	"fmt"

	helper "canonical.com/my_app/v2/pkg"
)

var commit string

func main() {
	fmt.Printf("Version %s\n", commit)
	helper.Help()
}
