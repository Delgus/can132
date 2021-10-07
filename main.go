package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	cnt := flag.Int("cnt", 100000, "count of attempts")
	flag.Parse()


	fmt.Print(`
			Машина смерти сошла с ума
			Она летит, сметая всех
			Мы увернулись - на этот раз
			Ушли по белой полосе`, "\n\n",
	)

	result := make(map[string]int)

	bar := pb.StartNew(*cnt)

	for i := 0; i < *cnt; i++ {
		bar.Increment()

		output, err := exec.Command("./task/task").Output()

		if err != nil {
			log.Fatal(err)
		}

		result[string(output)]++
	}

	bar.Finish()

	for key, val := range result {
		fmt.Printf("\"%s\" - %d times \n", key, val)
	}

	fmt.Println(`
			Мы здесь сегодня
			А завтра будем там
			Где тошно от огня чертям!`,
	)
}
