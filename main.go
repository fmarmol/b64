package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

const usageMsg = `	b64 encode or decode in base64 a message to standard output
	Usage: b64 [encode|decode] arg
`

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "wrong usage:"+usageMsg)
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "encode":
		res := base64.RawStdEncoding.EncodeToString([]byte(os.Args[2]))
		fmt.Println(res)
	case "decode":
		res, err := base64.RawStdEncoding.DecodeString(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: could not decode msg:%v\n", err)
		} else {
			fmt.Println(string(res))
		}
	default:
		fmt.Fprintf(os.Stderr, "wrong usage\n"+usageMsg)
		os.Exit(1)
	}
}
