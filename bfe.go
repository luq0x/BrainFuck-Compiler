package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const memorySize = 30000
	mem := make([]byte, memorySize)
	ptr := 0

	code, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Erro ao ler entrada:", err)
		return
	}

	loopStack := []int{}
	for pc := 0; pc < len(code); pc++ {
		switch code[pc] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '.':
			os.Stdout.Write([]byte{mem[ptr]})
		case ',':
			b := make([]byte, 1)
			os.Stdin.Read(b)
			mem[ptr] = b[0]
		case '[':
			if mem[ptr] == 0 {
				loop := 1
				for loop > 0 {
					pc++
					if code[pc] == '[' {
						loop++
					} else if code[pc] == ']' {
						loop--
					}
				}
			} else {
				loopStack = append(loopStack, pc)
			}
		case ']':
			if mem[ptr] != 0 {
				pc = loopStack[len(loopStack)-1]
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		}
	}
}
