package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Bem-vindo ao jogo da adivinhação!")
	var todasTentativas [][]int

	for {
		resposta := numAleatorio()
		tentativas := make([]int, 0)
		encontrou := false

		for !encontrou {
			chute := numUsuario()
			tentativas = append(tentativas, chute)

			switch {
			case chute == resposta:
				fmt.Println("Parabéns, você acertou!")
				encontrou = true
				fmt.Printf("Você usou %d tentativas\n", len(tentativas))
			case chute < resposta:
				fmt.Printf("O número é maior que %d\n", chute)
			case chute <= 0 || chute > 100:
				fmt.Println("Digite um número válido")
			default:
				fmt.Printf("O número é menor que %d\n", chute)
			}
		}

		todasTentativas = append(todasTentativas, tentativas)

		if !pedirParaJogar() {
			break
		}
	}

	fmt.Println("\nHistórico de tentativas:")

	for i, tentativas := range todasTentativas {
		fmt.Printf("Jogada: %d  -  Tentativas: %d\n", i+1, len(tentativas))
	}
}

func numAleatorio() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

func numUsuario() int {
	var chute int
	fmt.Print("Digite um número entre 1 e 100: ")
	fmt.Scanln(&chute)
	return chute
}

func pedirParaJogar() bool {
	var resposta string
	fmt.Print("Você deseja jogar novamente? (s/n): ")
	fmt.Scanln(&resposta)
	return resposta == "s" || resposta == "S"
}
