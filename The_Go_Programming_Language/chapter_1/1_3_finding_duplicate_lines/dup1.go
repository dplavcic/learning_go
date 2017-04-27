// Dip1 prints the text of each line that appears more than
// once in the std IN, preceded by its count..

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)

// 	for input.Scan() {
// 		if len(input.Text()) > 0 {
// 			counts[input.Text()]++
// 			fmt.Println(counts)
// 		} else {
// 			break
// 		}

// 	}

// 	// ignore errors
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin, &counts)
// 	} else {
// 		for _, args := range files {
// 			f, err := os.Open(args)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countLines(f, &counts)
// 			f.Close()
// 		}
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func countLines(f *os.File, counts *map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		if len(input.Text()) > 0 {
// 			(*counts)[input.Text()]++
// 			fmt.Println(counts)
// 		} else {
// 			break
// 		}
// 	}
// 	//ignore error
// }

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// returns byte slice
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
