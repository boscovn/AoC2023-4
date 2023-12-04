package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	num := 0
	wins := make(map[int]int)

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
		numWin := 0
		for _, win := range winning {
			if win == "" {
				continue
			}
			for _, m := range mine {
				if m == "" {
					continue
				}
				if win == m {
					numWin++
					break
				}
			}
		}
		wins[num] = numWin
		num += 1

	}
	var pending []int
	for k, v := range wins {
		if v > 0 {
			for j := 1; j <= v; j++ {
				index := j + k
				if index < num {
					pending = append(pending, index)
				}
			}
		}
	}
	sum := num
	for len(pending) > 0 {
		k := pending[0]
		pending = pending[1:]
		v := wins[k]
		if v > 0 {
			for j := 1; j <= v; j++ {
				index := j + k
				if index < num {
					pending = append(pending, index)
				}
			}
		}
		sum += 1

	}

	fmt.Println(sum)
}
