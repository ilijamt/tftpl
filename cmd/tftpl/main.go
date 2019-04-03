package main

import (
	"fmt"
	"github.com/ilijamt/tftpl/internal/cmd"
	"os"
)

func main() {
	var err error
	if err = cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
