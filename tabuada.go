package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Seja bem vinda ao quiz de tabuada :D")
	finalResult := 0
	index := 0
	for index < 10 {
		finalResult += exercise()
		index++
	}

	fmt.Println("Sua nota final foi: ", finalResult)
}

func exercise() int {
	reader := bufio.NewReader(os.Stdin)

	rand.Seed(time.Now().UnixNano())
	multiplier := rand.Intn(11)
	by := rand.Intn(11)

	fmt.Print("Quanto é ", multiplier, " x ", by, " = ")
	value, _ := reader.ReadString('\n')

	newValue := value[:len(value)-1]

	x, err := strconv.Atoi(newValue)
	if err != nil {
		fmt.Println("Atenção, pressione a tecla certa!")

	}
	if x == multiplier*by {
		fmt.Println("Hmmmmm.... está estudando")
		return 1
	} else {
		fmt.Print("A resposta certa é ", multiplier, " x ", by, " = ", multiplier*by, "\n")
		return 0
	}
}
