package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello GO")

	godotenv.Load()

	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("PORT:", portstring)
}
