package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func countLetters(filename io.Reader) (int, error) {
	content, _ := ioutil.ReadAll(filename)

	return len(content), nil

}

func countLines(filename io.Reader) (int, error) {
	sumLine := 0
	content, _ := ioutil.ReadAll(filename)

	for _, v := range content {
		if string(v) == "\n" {
			sumLine += 1
		}
	}
	return sumLine, nil

}

var (
	// c = flag.String("c", "-c", "count letter")
	// l = flag.String("l", "-l", "count lines")
	lines   = flag.Bool("l", false, "print lines")
	letters = flag.Bool("c", false, "print letters")
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if *lines {
		fmt.Println(countLines(f))

	} else if *letters {
		fmt.Println(countLetters(f))
	}

}
