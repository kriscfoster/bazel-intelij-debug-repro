package main

import (
	"log"
	"os"
)

func main() {
	log.Println("MY_VARIABLE: " + os.Getenv("MY_VARIABLE"))
}
