package main

import (
	"fmt"
	"math/rand"
)
var palavras = []string{"verdadeiro", "criança", "jogo", "caderno", "notebook"}

var indice int
var indiceAleatorio int
var indiceTemporario int

func main(){
		geradorDePalavras()
}

func geradorDePalavras(){	

	indiceAleatorio = sorteiaIndiceAleatorio(palavras)

	if comparaIndice(indiceAleatorio, indiceTemporario) {
		indiceAleatorio = sorteiaIndiceAleatorio(palavras)
	} else {
		fmt.Println("indice Temp: ", indiceTemporario)
		fmt.Println("Indice Alea: ",  indiceAleatorio)
		fmt.Println(palavras[indiceAleatorio])
		indiceTemporario = indiceAleatorio
	}
}

func sorteiaIndiceAleatorio(array []string) int {
	indiceSorteado := rand.Intn(len(array))

	return indiceSorteado
}

func comparaIndice(a, b int) bool {
	if a == b{
		return true
	}else{
		return false
	}
}

