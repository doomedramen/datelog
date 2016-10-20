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
		fmt.Printf("[%v] %v \n",time.Now().Format(time.RFC850),scannerStdIn.Text())
	}
}