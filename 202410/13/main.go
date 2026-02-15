package main

import (
	"fmt"
	"log"
	"os/user"
	"strings"
)

func main() {
	log.SetFlags(19)

	info, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	username := info.Username
	tmp := strings.Split(username, "\\")

	fmt.Println("Hello, " + tmp[1] + "!")
}
