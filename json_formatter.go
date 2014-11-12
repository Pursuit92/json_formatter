package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func errExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	in := make(map[string]interface{})
	err := json.NewDecoder(os.Stdin).Decode(&in)
	errExit(err)
	out, err := json.MarshalIndent(in, "", "  ")
	errExit(err)
	_, err = os.Stdout.Write(out)
	errExit(err)
	fmt.Fprintf(os.Stdout, "\n")
}
