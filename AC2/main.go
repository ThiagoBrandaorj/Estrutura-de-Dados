package main

import (
	"fmt"
)

func main(){
	vetor := [10]int{1,34,5,70,99,100,2,23,88,65}
	for i:= 0; i < len(vetor); i++{
		println(vetor[i])
	}

	var palavra string
	fmt.Print("Insira uma palavara , frase ou qualquer string: ")
	fmt.Scanln(&palavra)
	palavra_invertida := inverteString(palavra)
	fmt.Println(palavra_invertida)
}

type Ponto struct{
	X float64
	Y float64
}

func (p Ponto) DistanciaOrigem(){
	
}

func inverteString(str string) string{
	runes := []rune(str)
	tamanho := len(runes)
	for i,j := 0, tamanho - 1; i < j; i,j = i+1,j+1{
		runes[i],runes[j] = runes[j], runes[i]
	}

	return string(runes)
}