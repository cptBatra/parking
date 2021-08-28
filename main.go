package main

import (
	"flag"
	"fmt"
	"os"
	"parking/controller"
	"parking/model"
	"strings"
)

func main() {
	inputFileName := flag.String("input", "input.txt", "a txt filename to execute commands")
	flag.Parse()
	insSet, err := readInstructionFile(*inputFileName)
	if err != nil {
		exit(err.Error())
	}
	var o model.Response
	for _, v := range insSet {
		o, err = controller.RouteInstruction(v)
		if err != nil && strings.Contains(err.Error(), "Could not create Parking Lot") {
			exit(err.Error())
		}
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(o.Message)
	}

}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
