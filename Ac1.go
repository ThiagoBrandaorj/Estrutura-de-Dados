package main

import (
	"fmt"
)
func calculaMedia(args...float64) {
	//var sum float64
}

func verificaParidade(n int) {
	if n % 2 == 0{
		fmt.Println("o número" , n , "é par")
	} else{
		fmt.Println("o número" , n , "é Ímpar")
	}
}
func main(){
	verificaParidade(5)
	verificaParidade(12)

}