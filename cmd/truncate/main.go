package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"truncate"
)

func main() {
	width := 20
	omit := "…"
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-width":
			if i+1 < len(args) {
				if w, err := strconv.Atoi(args[i+1]); err == nil {
					width = w
				}
				i++
			}
		case "-ellipsis":
			if i+1 < len(args) {
				omit = args[i+1]
				i++
			}
		}
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fmt.Println(truncate.Truncate(sc.Text(), width, omit))
	}
}
