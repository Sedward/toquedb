package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ErrUnrecognizedCommand struct{}

func (u ErrUnrecognizedCommand) Error() string {
	return "unrecognized command"
}

func handleMetaCommand(input string) error {
	if input == ".exit" {
		fmt.Println("Closing database")
		os.Exit(0)
	} else {
		return new(ErrUnrecognizedCommand)
	}

	return nil

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("ToqueDB version 0.1.0\n")

	for {
		fmt.Printf("toquedb > ")
		in, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("%T", err)
		}

		in = strings.TrimSpace(in)
		if in[0] == '.' {
			err := handleMetaCommand(in)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
