package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var palavras = []string{
	"verdadeiro", "criança",
	//"verdadeiro", "criança", "jogo", "caderno", "notebook",
	//"cachorro", "presente", "festa", "Youtube", "motor",
}
var segundoSlice []string

var asPalavrasAcabaram = false

var indice int
var indiceAleatorio int

var contaPontos int

var contaTempo int
var tempoAcabou bool

func main() {

	partida()
}

func temporizador() {

}

func partida() {
loop:
	for {
		switch menu() {
		case 1:
			if asPalavrasAcabaram == true {
				fmt.Println("nao ha mais palavras")
				partida()

			}
			geraPalavras()
			verificaAcerto()

		case 2:
			fmt.Println("programa interrompido")
			break loop
		default:
			fmt.Println("número inválido | não aceito")
		}
	}
}

func verificaAcerto() {
	fmt.Println("Acertou a palavra??")
	fmt.Println("tecla [S] para sim\ntecla [N] para não")
	verifica := lerInput()
	switch verifica {
	case "S":
		contaPontos += 1
		fmt.Println("Parabéns")
	case "N":
		fmt.Println("Mais sorte da próxima vez")
	default:
		fmt.Println("Digito incorreto!")
	}
}

func menu() int {
	fmt.Println("\n=========================\n"+
		"Pontuação da rodada:", contaPontos, "\n"+
		"gerar nova palavra [1]\n"+
		"parar aplicação [2]")

	input, err := strconv.Atoi(lerInput())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n=========================")
	return input
}

func lerInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Println(err)
	}

	return scanner.Text()
}

func geraPalavras() {
	time.Sleep(200 * time.Millisecond)
	indice = sorteiaIndiceAleatorio(palavras)

	fmt.Println(indice)
	fmt.Println(palavras)
	fmt.Println(palavras[indice])

	segundoSlice = remove(palavras, indice)
	fmt.Println(segundoSlice)

	palavras = segundoSlice

	if len(segundoSlice) < 1 {
		asPalavrasAcabaram = true
	}

}

func sorteiaIndiceAleatorio(array []string) int {
	indiceAleatorio := rand.Intn(len(array))

	return indiceAleatorio
}

func remove[T comparable](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
