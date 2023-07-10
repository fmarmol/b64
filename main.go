package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

const usageMsg = `
b64 encode or decode in base64 a message to standard output
  Usage: b64 [encode|decode] arg
`

const usageMsgPipe = `
b64 encode or decode in base64 a message to standard output
  Usage: echo "xxx" | b64 [encode|decode]
`

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot stat exectuable")
		os.Exit(1)
	}

	var argument string

	checkPipeMode := fi.Mode()&os.ModeNamedPipe == os.ModeNamedPipe
	if checkPipeMode {

		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "wrong usage:"+usageMsgPipe)
			os.Exit(1)
		}
		argumentRaw, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot read stdin: %v\n", err)
			os.Exit(1)
		}
		argument = string(argumentRaw)
	} else {
		if len(os.Args) != 3 {
			fmt.Fprintf(os.Stderr, "wrong usage:"+usageMsg)
			os.Exit(1)
		}
		argument = os.Args[2]

	}

	cmd := os.Args[1]
	switch cmd {
	case "encode":
		res := base64.StdEncoding.EncodeToString([]byte(argument))
		fmt.Println(res)
	case "decode":
		res, err := base64.StdEncoding.DecodeString(argument)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: could not decode msg:%v\n", err)
		} else {
			fmt.Println(string(res))
		}
	default:
		if !checkPipeMode {
			fmt.Fprintf(os.Stderr, "wrong usage\n"+usageMsg)
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "wrong usage\n"+usageMsgPipe)
			os.Exit(1)

		}
	}
}
