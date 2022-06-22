package main

import (
	"fmt"
	"io"
	"os"
)

type args interface{}

func stringPrinter(w io.Writer, a []args) {
	for _, e := range a {
		if str, ok := e.(string); ok {
			str = fmt.Sprintf("%s\n", str)
			w.Write([]byte(str))
		}
	}
}

func main() {
	f, err := os.Create("./only_strings.txt")
	if err != nil {
		fmt.Println("File not open", err)
		return
	}
	defer f.Close()

	stringPrinter(f, []args{"qwe", 2, "asd", 123})
}
