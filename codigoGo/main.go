package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	palavras := []string{"verdadeiro", "criança", "jogo", "caderno", "notebook"}
	
	var indice int
	var indiceTemporario int
	var indiceAleatorio int

	for {
		time.Sleep(500 * time.Millisecond)	

		indice = sorteiaIndiceAleatorio(palavras)
		indiceAleatorio = indice

		if comparaIndice(indiceAleatorio, indiceTemporario) {
			indice = sorteiaIndiceAleatorio(palavras)
			indiceAleatorio = indice
		} else {
		
			fmt.Println("indice Temp: ", indiceTemporario)

			fmt.Println("Indice Alea: ",  indiceAleatorio)
	
			fmt.Println(palavras[indiceAleatorio])
	
			indiceTemporario = indice
		}
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
