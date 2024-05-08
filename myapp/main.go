package main

import (
	"fmt"

	"github.com/polyglotdev/celeritas"
)

func main() {
	result := celeritas.TestFunc(10, 10)
	fmt.Println(result)

	result = celeritas.TestFunc2(10, 20)
	fmt.Println(result)
}
