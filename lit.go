package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const TAB string = "    "

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.Compare(arg, "--usage") == 0 {
			fmt.Println("Usage: lit [opts] [file]")
			fmt.Println("Opts:")
			fmt.Println(TAB + "--usage")
			fmt.Println(TAB + "--stdin")
			return
			fmt.Println("shouldnt happen")
		}
		fmt.Println("for loop")
	}
	//buffer := make([]string, 32)
	if strings.Compare(args[0], "--stdin") != 0 {
		path := args[0]
		data, err := ioutil.ReadFile(path)
		check(err)
		fmt.Println("here if")
		fmt.Print(string(data))
	}
	fmt.Println("prob not")
	return
}
