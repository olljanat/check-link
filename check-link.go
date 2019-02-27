package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 1 {
		fmt.Println("Link URL:")
		fmt.Println(argsWithoutProg[1])
		fmt.Println("\nDo you want open it? [Y/N]")
		var input string
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "Y" {
			cmd := exec.Command(`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`, argsWithoutProg[1])
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Println("No link provided")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
