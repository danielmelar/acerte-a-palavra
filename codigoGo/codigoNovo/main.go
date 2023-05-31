package main

import (
	"fmt"
	"math/rand"
	"time"
)

var palavras = []string{
	"verdadeiro", "criança", "jogo", "caderno", "notebook",
	"cachorro", "presente", "festa", "Youtube", "motor",
}
var segundoSlice []string

var indice int
var indiceAleatorio int
var indiceTemporario int

func main() {
	geraPalavras()
}

func geraPalavras() {

loop:
	for {
		time.Sleep(200 * time.Millisecond)
		indice = sorteiaIndiceAleatorio(palavras)

		fmt.Println(indice)
		fmt.Println(palavras)
		fmt.Println(palavras[indice])

		segundoSlice = remove(palavras, indice)
		fmt.Println(segundoSlice)

		palavras = segundoSlice

		if len(palavras) == 0 {
			break loop
		}
	}

}

func sorteiaIndiceAleatorio(array []string) int {
	indiceAleatorio := rand.Intn(len(array))

	return indiceAleatorio
}

func remove[T comparable](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
