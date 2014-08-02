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
		receivedSignal := <-sigChannel
		os.Exit(1)
	}()

	hjalp := flag.Bool("help", false, "I have no idea what I am doing")
	flag.Parse()

	if *hjalp == true {
		fmt.Println("Usage: ls -al | rnbw")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)
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
