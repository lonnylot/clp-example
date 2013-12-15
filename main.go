package main

import (
	"fmt"
	"github.com/lonnylot/clp"
	"os"
	"time"
)

func main() {
	fmt.Println("Bar: Reading Large File")
	largeFileExample()
	fmt.Println("Dots: Waiting for a process to complete")
	dotsExample()
}

func largeFileExample() {
	file, _ := os.Open("largeFile")
	fInfo, _ := file.Stat()

	// Create our clp bar
	clp := clp.NewBar(fInfo.Size())
	// Start the output
	clp.Start()

	b := make([]byte, 1024)
	for {
		n, err := file.Read(b)

		// Increment clp
		clp.Inc(int64(n))

		// Catch our EOL
		if err != nil {
			break
		}
	}

	// Stop our clp bar
	clp.Stop()
}

func dotsExample() {
	// Create our clp dots
	clp := clp.NewDots()
	// Start the output
	clp.Start()

	t := time.After(time.Duration(time.Second * 10))
	<-t // Wait for the timer...

	// Stop our clp dots
	clp.Stop()
}
