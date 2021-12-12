package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
)

var helpstr string = `bits - how many bits are required to represent some number
Usage: 
	bits [-h] [-help] <number> 
`

func main() {

	h := flag.Bool("h", false, "show help")
	help := flag.Bool("help", false, "show help")
	flag.Parse()

	if *h == true || *help == true {
		fmt.Printf(helpstr)
		return
	}

	if len(os.Args) < 2 {
		fmt.Printf(helpstr)
		return
	}

	num := new(big.Int)
	_, err := fmt.Sscan(os.Args[1], num)
	if err != nil {
		fmt.Printf("failed to parse input: %v\n", err)
		return
	}
	fmt.Printf("%d\n", num.BitLen())
	return
}
