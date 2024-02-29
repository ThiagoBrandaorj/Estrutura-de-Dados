package main

import (
	"fmt"
	"math"
	m "projeto/geometria"
)

func main() {
	// 1
	vetor := [10]int{1, 34, 5, 70, 99, 100, 2, 23, 88, 65}
	for i := 0; i < len(vetor); i++ {
		println(vetor[i])
	}

	// 2
	var palavra string
	fmt.Print("Insira uma palavara , frase ou qualquer string: ")
	fmt.Scanln(&palavra)
	palavra_invertida := inverteString(palavra)
	fmt.Println(palavra_invertida)

	//3
	obj := Ponto{}
	var x , y float64
	fmt.Println("Insira o X do ponto: ")
	fmt.Scanln(&x)
	fmt.Println("Insira o Y do ponto: ")
	fmt.Scanln(&y)
	dist:= obj.DistanciaOrigem(&x,&y)
	fmt.Println("Distancia do ponto de (0,0): ",dist)

	//4
	var comprimento, largura float64 
	fmt.Println("Insira o comprimento do retângulo: ")
	fmt.Scanln(&comprimento) 
	fmt.Println("Insira o largura do retângulo: ")
	fmt.Scanln(&largura)
	area := m.CalculaAreaRetangulo(&comprimento, &largura)
	perimetro := m.CalculaPerimetroRetangulo(&comprimento, &largura)
	fmt.Println("A área do retângulo: " , area , "o perímetro: ", perimetro)
}

type Ponto struct {
	X float64
	Y float64
}

func (p Ponto) DistanciaOrigem(x, y *float64) float64 {
	calculo := math.Sqrt(math.Pow(*x-0, 2) + math.Pow(*y-0, 2))
	return calculo
}

func inverteString(str string) string {
	slice := []rune(str)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return string(slice)
}
