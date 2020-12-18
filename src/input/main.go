package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%#v\n", os.Args[0])
	fmt.Printf("%#v\n", os.Args[1])
	fmt.Printf("%#v\n", os.Args[2])
	fmt.Printf("%#v\n", len(os.Args))
}
