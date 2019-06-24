package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"ch3/words"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)
	fmt.Printf("%s\n", text)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text.\n", count)
}
