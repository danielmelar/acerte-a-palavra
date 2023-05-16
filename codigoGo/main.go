package main

import (
	"fmt"
	"math/rand"
	"bufio"
	"os"
	"strconv"
)
var palavras = []string{"verdadeiro", "criança", "jogo", "caderno", "notebook"}

var indiceJaSorteado []int

var indice int
var indiceAleatorio int
var indiceTemporario int

var inputText string

var inputConvertidoParaInt int

func main(){

	geraMenu()

	geraSwitch()
}

func geraSwitch() {
	loop:
	for {
		switch  inputConvertidoParaInt {
		case 1:
			geradorDePalavras()
			indiceJaSorteado = append(indiceJaSorteado, indiceAleatorio)	
			fmt.Println(indiceJaSorteado)
			geraMenu()

		case 2: 
			fmt.Println("Programa Encerrado")
			break loop
		default:
			fmt.Println("Número inválido|Não aceito")
			geraMenu()
		}
	}
}

func convertStringParaInt(input string) int{
	inputConvertidoParaInt, err := strconv.Atoi(input)
	if err != nil{
		fmt.Println(err)
	}

	return inputConvertidoParaInt
}

func geraMenu() {
	fmt.Print(`
	========================
	O que desejas?
	1 - Gerar nova palavra
	2 - Para programa
	========================
	Input: `)
	inputText = lerInput()

	fmt.Println("\n")
	inputConvertidoParaInt = convertStringParaInt(inputText)
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

