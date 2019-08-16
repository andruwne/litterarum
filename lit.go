package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
)

const TAB string = "    "

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func print_usage() {
	fmt.Println("Usage: lit.go [opts] [file]")
	fmt.Println("Opts:")
	fmt.Println(TAB + "--stdin")
}

func contains(s []byte, e uint8) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	var args []string
	if len(os.Args) < 2 {
		return
	}
	args = os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.Compare(arg, "--usage") == 0 {
			print_usage()
			return
		}
	}
	//buffer := make([]string, 32)
	var data []byte
	if strings.Compare(args[0], "--stdin") == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data = []byte(scanner.Text())
			if contains(data, uint8(4)) {
				return
			}
			fmt.Println(string(data)) // Println will add back the final '\n'
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	} else {
		path := args[0]
		data, err := ioutil.ReadFile(path)
		a := &data
		fmt.Println(a)
		check(err)
	}
	fmt.Print(string(data))
	return
}
