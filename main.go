package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		index := strings.Index(text, ":")
		text = text[index+2:]

		parts := strings.Split(text, " | ")
		winning := strings.Split(parts[0], " ")
		mine := strings.Split(parts[1], " ")
		numWin := 1
		for _, win := range winning {
			if win == "" {
				continue
			}
			for _, m := range mine {
				if m == "" {
					continue
				}
				if win == m {
					numWin *= 2
					break
				}
			}
		}
		if numWin != 1 {
			sum += numWin / 2
		}

	}

	fmt.Println(sum)
}
