package main

import (
	"fmt"

	modhelloworld "github.com/SutanBahtiar/go-mod-hello-world"
)

func main() {
	modhelloworld.PrintHelloWorld()
	fmt.Println(modhelloworld.GetHelloWorld())
}
