package main

import (
	"fmt"
	"os"
	"bufio"
	"time"
)

func main() {
	scannerStdIn := bufio.NewScanner(os.Stdin)
	for scannerStdIn.Scan() {
		fmt.Printf("\x1b[32;1m[%v]\x1b[0m %v \n", time.Now().Format(time.Stamp), scannerStdIn.Text())
	}
}