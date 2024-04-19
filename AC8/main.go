package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//1
func RevisaoContrato(){
	var D , N int
	for{
		fmt.Scan(&D, &N)
		if D >= 1 && D <= 9 && N >= 1 && N <= int(math.Pow(10,100)){
			Nstring := fmt.Sprintf("%d", N)
			Dstring := fmt.Sprintf("%d", D)
			Vstring := strings.ReplaceAll(Nstring,Dstring,"")
			V, err := strconv.Atoi(Vstring)
			if err != nil{
				V = 0
			}
			fmt.Println(V)
		}
		if D == 0 && N == 0{
			break
		}
	}
}
//2
// Na condição do V ser menor do que 10**100 , o número em questao era muito grande
// em troca usei o tamanho de V <= 100
func Led(){
	var N ,V int
	var qleds = []int{6,2,5,5,4,5,6,3,7,6}
	fmt.Scan(&N)
	if N >= 0 && N <= 1000{
		for i:=0;i<N;i++{
			fmt.Scan(&V)
			if V >= 1 && len(fmt.Sprintf("%d", V)) <= 100{
				totleds := 0
				for V > 0{
					x := V % 10
					totleds += qleds[x]
					V /= 10
				}
				fmt.Println( totleds , " Leds")
			}
		}
	}
}
//3
func CifraCesar(){
	var N, deslocamento int
	var sentenca string
	var alfabeto = []string{"A","B","C","D","E","F","G","H","I","J","K",
	"L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	fmt.Scan(&N)
	if N > 0{
		for i:= 0; i < N; i++{
			fmt.Scan(&sentenca)
			if len(sentenca) >= 1 && len(sentenca) <= 50{
				fmt.Scan(&deslocamento)
				if deslocamento >= 0{
					
				}
			}
		}
	}
}
func main(){
	//RevisaoContrato()
	Led()
}