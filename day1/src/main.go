package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	ls := readFile("../input/input")
	start := time.Now()
	sum := 0

	for _, l := range ls {
		sum += retornaValor(l)
	}
	elapsed := time.Since(start)

	log.Printf("Resultado %d", sum)
	log.Printf("tempo de execução: %s", elapsed)
}

func readFile(s string) []string {
	var r []string
	f, err := os.Open(s)
	if err != nil {
		log.Panic(err)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()
		if len(t) > 0 {
			r = append(r, t)
		}
	}
	return r
}

func getValue(s string) int {
	if len(s) == 0 {
		return 0
	}
	var f, l rune
	fb := false

	for _, r := range s {
		if isNum(r) {
			l = r
			if !fb {
				fb = true
				f = r
			}
		}

	}
	sv := string(f) + string(l)
	v, err := strconv.Atoi(sv)
	if err != nil {
		log.Panic(err)
		return 0
	}
	return v
}

func runeMagic(r rune) int {
	return int(r - '0')
}

func isNum(b rune) bool {
	rs := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	return slices.Contains(rs, b)
}

func retornaValor(s string) int {
	vp := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	pi, ui := -1, -1
	pv, uv := "", ""

	log.Print(s)
	for _, v := range vp {
		fix, lix := strings.Index(s, v), strings.LastIndex(s, v)
		if fix == -1 {
			continue
		}

		if pi == -1 {
			pi = fix
			ui = lix
			pv, uv = v, v
		} else {
			if fix < pi {
				pi = fix
				pv = v
			}

			if lix > ui {
				ui = lix
				uv = v
			}

		}
	}

	vr := literalParaNumerico(pv) + literalParaNumerico(uv)
	c, err := strconv.Atoi(vr)
	if err != nil {
		log.Panic(err)
	}
	return c
}

func literalParaNumerico(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	case "zero":
		return "0"
	default:
		return s
	}
}
