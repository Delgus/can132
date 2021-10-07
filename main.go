package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	for i := 0; i < 10000; i++ {
		output, err := exec.Command("go", "run", "task/main.go").Output()

		if err != nil {
			log.Fatal(err)
		}

		if string(output) != "13" {
			fmt.Printf("unexpected - %s \n", string(output))
		}
	}
}
