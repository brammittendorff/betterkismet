package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var (
	myinterface string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func init() {
	flag.StringVarP(&myinterface, "interface", "i", "", "Select interface")
}
