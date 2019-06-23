package main

import (
	"log"
	"os"

	_ "ch2/sample/matchers"
	"ch2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
