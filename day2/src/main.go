package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type gameInfo struct {
	id      int
	subsets []subsetInfo
}

type subsetInfo []cubeInfo

type cubeInfo map[string]int

func main() {
	log.Println("Iniciando leitura")
	file, err := os.Open("data/meu")
	if err != nil {
		log.Panic(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	powerSum := 0
	for scanner.Scan() {
		game := createGame(scanner.Text())
		if game.isPossible(12, 13, 14) {
			sum = sum + game.id
		}

		minimum := game.getMinimum()
		powerSum = powerSum + minimum["red"]*minimum["blue"]*minimum["green"]

	}

	//joao 2449
	log.Printf("Sum is: %d\n", sum)
	log.Printf("PowerSum is: %d\n", powerSum)
}

func (game gameInfo) getMinimum() cubeInfo {
	result := make(cubeInfo)
	result["red"] = 0
	result["green"] = 0
	result["blue"] = 0
	for _, subset := range game.subsets {

		for _, cube := range subset {
			if cube["red"] > result["red"] {
				result["red"] = cube["red"]
			}
			if cube["green"] > result["green"] {
				result["green"] = cube["green"]
			}
			if cube["blue"] > result["blue"] {
				result["blue"] = cube["blue"]
			}
		}
	}

	return result
}

func (game gameInfo) isPossible(rc, gc, bc int) bool {
	for _, subset := range game.subsets {
		for _, cube := range subset {
			if cube["red"] > rc {
				return false
			}
			if cube["green"] > gc {
				return false
			}
			if cube["blue"] > bc {
				return false
			}
		}
	}
	return true
}

func createGame(text string) gameInfo {
	data := strings.SplitN(text, ":", 2)
	gameData, subsetData := data[0], data[1]
	gameIdStr := strings.Replace(gameData, "Game ", "", 1)
	gameId, err := strconv.Atoi(gameIdStr)
	if err != nil {
		log.Panic(err)
	}
	someSubsets := strings.Split(subsetData, ";")
	subSet := make([]subsetInfo, len(someSubsets))
	for si, subset := range someSubsets {
		cubes := strings.Split(subset, ",")

		subSet[si] = make(subsetInfo, len(cubes))

		for ci, cube := range cubes {
			cubeValues := strings.Fields(cube)
			quantity, err := strconv.Atoi(cubeValues[0])
			if err != nil {
				log.Panic(err)
			}

			if subSet[si][ci] == nil {
				subSet[si][ci] = make(cubeInfo)
			}
			subSet[si][ci][cubeValues[1]] = quantity
		}
	}

	return gameInfo{
		id:      gameId,
		subsets: subSet,
	}
}
