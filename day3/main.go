package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	anterior, atual, proxima := "", "", ""
	soma := 0
	for {
		anterior, atual = atual, proxima
		switch linha, err := reader.ReadString('\n'); err {
		case nil:
			proxima = strings.Trim(linha, "\n")
			soma += somarLinha(anterior, atual, proxima)
		case io.EOF:
			fmt.Println("Fim do arquivo")
			soma += somarLinha(anterior, atual, "")
			fmt.Printf("Resposta: %d\n", soma)
			os.Exit(0)
		default:
			log.Panic(err)
		}
	}
}

func somarLinha(anterior, atual, proxima string) int {
	if atual == "" {
		return 0
	}
	var posSimbolos []int
	var r rune
	soma := 0
    i := 0
	if anterior != "" {
		for i, r = range anterior {
			if isSymbol(r) {
				posSimbolos = append(posSimbolos, i)
			}
		}
	}

	if proxima != "" {
		for i, r = range proxima {
			if isSymbol(r) {
				posSimbolos = append(posSimbolos, i)
			}
		}
	}

	p0, p1 := -1, -1
	strNum := ""
	r_atual := []rune(atual)
    tam := len(r_atual) 
    for i = 0; i < tam+1; i++{ 
		if i < tam && isNumber(r_atual[i]) {
			if p0 == -1 {
				p0 = i
				p1 = i
			} else {
				p1 = i
			}
			strNum += string(r_atual[i])
			fmt.Printf("numero eh %s\n", strNum)
		} else if p0 != -1 {
			fmt.Printf("checando valor: %s ", strNum)
			if checarMesmaLinha(p0, p1, r_atual) || checarOutras(p0, p1, posSimbolos) {
				fmt.Println("aceito")
				n, err := strconv.Atoi(strNum)
				if err != nil {
					log.Panic(err)
				}

				soma += n
			} else {
				fmt.Println("não aceito")
			}

			strNum = ""
			p0, p1 = -1, -1

		}
	}

	return soma
}

func checarMesmaLinha(pos0 int, pos1 int, linha []rune) bool {
	if pos0-1 >= 0 && isSymbol(linha[pos0-1]) {
		return true
	}
	if pos1+1 < len(linha) && isSymbol(linha[pos1+1]) {
		return true
	}
	return false
}

func checarOutras(pos0 int, pos1 int, posSimbolos []int) bool {
	for i := pos0 - 1; i <= pos1+1; i++ {
		if slices.Contains(posSimbolos, i) {
			return true
		}
	}
	return false
}
func isNumber(r rune) bool {
	numbers := "0123456789"
	c := strings.ContainsRune(numbers, r)
	return c
}

func isDot(r rune) bool {
	return r == '.'
}

func isSymbol(r rune) bool {
	return !isNumber(r) && !isDot(r)
}
