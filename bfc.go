package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func repeat(char rune, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += string(char)
	}
	return result
}

func encodeByte(b byte) string {
	loop := b / 10
	rest := b % 10

	if loop == 0 {
		return ">" + repeat('+', int(rest)) + ".[-]<"
	}
	bf := ""
	bf += repeat('+', 10)
	bf += "[" + ">" + repeat('+', int(loop)) + "<" + "-]"
	bf += ">" + repeat('+', int(rest)) + ".[-]<"
	return bf
}

func encodeString(s string) string {
	out := ""
	for _, b := range []byte(s) { // itera corretamente sobre os bytes UTF-8
		out += encodeByte(b)
	}
	return out
}

func encodeInt(n int) string {
	return encodeString(fmt.Sprintf("%d", n))
}


func compileExpression(expr string) int {
	expr = strings.TrimSpace(expr)
	stack := []int{}
	op := '+'
	num := 0
	for i := 0; i < len(expr); i++ {
		c := expr[i]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}
		if c < '0' || c > '9' || i == len(expr)-1 {
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			op = rune(c)
			num = 0
		}
	}
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}

func encodeNumber(n int) string {
	return encodeByte(byte(n))
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Erro ao ler entrada:", err)
		return
	}
	entrada := strings.TrimSpace(string(data))

	eq := strings.Index(entrada, "=")
	if eq == -1 {
		fmt.Println("Entrada inválida")
		return
	}
	prefix := entrada[:eq+1]
	expr := entrada[eq+1:]
	result := compileExpression(expr)

	fmt.Print(encodeString(prefix))
	fmt.Print("[-]") 
	fmt.Print(encodeInt(result))

}
