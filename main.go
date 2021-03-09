package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/tarrows/csdb/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		print(err)
	}

	fmt.Printf("Hi, %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
