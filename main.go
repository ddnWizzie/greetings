package main

import (
	"log"

	greeting "wizzie.io/golinuxcloud/greetings"
)

func main() {
	name := "Dayana"
	res := greeting.Hello(name)

	log.Println(res)
}
