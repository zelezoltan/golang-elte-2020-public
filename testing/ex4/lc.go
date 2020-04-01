// Binary lc counts the number of lines in a file and also determines the
// lenght of the shortest and the longest lines.
package main

import (
	"flag"
	"fmt"
	"github.com/hu-univ-golang/golang-elte-2020-public/testing/ex4/lines"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		mmc, err := lines.Count(path)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			continue
		}
		fmt.Printf("%+v\t%s\n", mmc, path)
	}
}
