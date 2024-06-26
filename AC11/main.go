package main

import "fmt"



//1
func Distancia(){
	var entrada int
	fmt.Scan(&entrada)
	tempo := entrada * 2
	fmt.Print( tempo," minutos ")
}
//2
func Feira(){
	var N int
	fmt.Scanln(&N)
	for i:=0; i < N; i++{
		var M, P int
		fmt.Scanln(&M)
		produtos:= make(map[string]float64)
		for i:= 0; i < M; i++{
			var produto string
			var valor float64
			fmt.Scanf("%s %f\n", &produto, &valor)
			produtos[produto] = valor
		}
		fmt.Scanln(&P)
		compras := make(map[string]int)
		for i:=0; i< P ; i++{
			var produto string
			var qtd int
			fmt.Scanf("%s %d\n", &produto, &qtd)
			compras[produto] = qtd
		}
		var custo float64
		for produto, quantidade := range compras{
			custo += produtos[produto] * float64(quantidade)
		}
		fmt.Printf("R$ %.2f\n", custo)
	}
}
//3
func Aviao(){
	var C,P,F int
	fmt.Scanf("%d %d %d\n", &C, &P, &F)
	if C * F <= P{
		fmt.Println("S")
	} else{
		fmt.Println("N")
	}
}
//4
func SequenciaSecreta(){
	var N, V, qtdMaxima int
	fmt.Scanf("%d\n",&N)
	if N >= 3 && N <= 300{
		sequencia := make([]int, N)
		for i:=0; i< N; i++{
			fmt.Scanf("%d\n", &V)
			if V == 1 || V == 2{
				sequencia[i] = V
			}
		}
		for i:=0; i < N ; i++{
			if i == N-1 && sequencia[i] != sequencia[i-1]{
				qtdMaxima++
			}
			if i < N - 1{
				if sequencia[i] != sequencia[i+1]{
					qtdMaxima++
				}
			}
		}
	}
	fmt.Printf("%d\n", qtdMaxima)
}

func main(){
	//Distancia()
	//Feira()
	//Aviao()
	SequenciaSecreta()
}