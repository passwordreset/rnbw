package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var Rainbow = [24]int{
	154, 184, 214, 208, 209, 203, 204, 198, 199, 164, 129,
	93, 99, 63, 69, 33, 39, 44, 49, 48, 84, 83, 119, 118}

var AnsiColourFormat = "\033[38;5;%d;m%s\033[0m"

func main() {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM)

	go func() {
		<-sigChannel
		os.Exit(1)
	}()

    file := flag.String("f", "", "File to rnbw up")
	flag.Parse()

    if *file != "" {
        f, err := os.Open(*file)
        if err != nil {
            panic(fmt.Sprintf("Can't open file %s", *file))
        }
        reader := bufio.NewReader(f)
        scanner := bufio.NewScanner(reader)
        readIn(scanner)
    } else {
        scanner := bufio.NewScanner(os.Stdin)
        readIn(scanner)
    }

}

func readIn(scanner *bufio.Scanner) {
	for scanner.Scan() {
		// TODO: Instead of going new-line splitting, go by bytes here, and then
		// split by newlines later
		makeItPretty(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		os.Exit(0)
	}

}

func makeItPretty(str string) {
	chars := strings.Split(str, "")
	pos := 0
    escape := false
	for _, chr := range chars {
        // Strip out other colours
        if chr == "\033" {
            escape = true
            continue
        } else if escape == true && chr == "m"{
            escape = false
            continue
        } else if escape == true {
            continue
        }

		fmt.Printf(AnsiColourFormat, Rainbow[pos], chr)

		if pos < 23 {
			pos++
		} else {
			pos = 0
		}
	}
	fmt.Println("")
}
