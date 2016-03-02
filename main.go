package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/atotto/clipboard"
)

func out() {
	if cli, err := clipboard.ReadAll(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	} else {
		fmt.Print(cli)
	}
}

func scan() (input string, err error) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	input = string(bytes)
	return
}

func in() {
	if d, err := scan(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	} else if err := clipboard.WriteAll(d); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}

func help() {
	fmt.Println("\t\tgoclip")
	fmt.Println("\ncross platform clipboard command like the xclip.")
	fmt.Println("")
	fmt.Printf("usage: %s [-i | -o | -h]\n", os.Args[0])
	fmt.Println("  -i  write stdin to clipboard")
	fmt.Println("  -o  read from clipboard.")
	fmt.Println("  -h  show this message and exit.")
}

func main() {
	switch len(os.Args) {
	case 1:
		out()
	case 2:
		switch os.Args[1] {
		case "-i":
			in()
		case "-o":
			out()
		case "-h":
			help()
		default:
			fmt.Fprintln(os.Stderr, "unknown option: %s\nplease see help(%s -h)\n", os.Args[1])
		}
	default:
		fmt.Fprintf(os.Stderr, "too much argument. can't do multi acction in one command.\nplease see help(%s -h)\n", os.Args[0])
		os.Exit(1)
	}
}
