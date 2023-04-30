package main

import (
	"fmt"
	"math/rand"
)

func main(){
	palavras := []string{"verdadeiro", "criança", "jogo", "caderno", "notebook"}
	
	indiceAleatorio := sorteiaIndiceAleatorio(palavras)

	fmt.Println(indiceAleatorio)
	fmt.Println(palavras[indiceAleatorio])
}

func sorteiaIndiceAleatorio(array []string) int {
	indiceSorteado := rand.Intn(len(array))

	return indiceSorteado
}










