package main

import (
	"fmt"
	"math/big"
	"os"
)

type IsIntegerConf struct {
	KeepFloat bool
	Input     string
}

func parseConfig() *IsIntegerConf {
	var conf IsIntegerConf
	// parse args manually to avoid flag package from misinterpreting negative numbers as flags
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-d" {
			conf.KeepFloat = true
		} else {
			conf.Input = arg
		}
	}
	if conf.Input == "-h" || conf.Input == "--help" {
		fmt.Println(`Usage: is-integer [-h] [-d] <number>
	-d: keep decimal part of number
	<number>: number to check if it is an integer

If it is an integer, it will be printed as such
If its a float and -d is not specified, the decimal part will be truncated

If the number is not a valid number, this will exit with a exit code of 1

To use this in a script, you can do something like this:

read -r somevar  # ask user for input
if parsed="$(is-integer "$somevar")"; then
	some_other_program "$parsed"
fi`)
		os.Exit(1)
	}
	return &conf
}

func isInteger() int {
	conf := parseConfig()
	if conf.KeepFloat {
		// try to parse as big.NewFloat
		bf, ok := new(big.Float).SetString(conf.Input)
		if ok {
			// print as big.Float
			fmt.Println(bf)
			return 0
		} else {
			return 1
		}
	} else {
		// try to parse as big.NewInt
		bi, ok := new(big.Int).SetString(conf.Input, 10)
		if ok {
			// print as big.Int
			fmt.Println(bi)
			return 0
		} else {
			// if that fails, try to parse as big.NewFloat
			bf, ok := new(big.Float).SetString(conf.Input)
			if ok {
				// and then print without fractional part
				bi, _ := bf.Int(nil)
				fmt.Println(bi)
				return 0
			} else {
				return 1
			}
		}
	}
}

func main() {
	os.Exit(isInteger())
}
