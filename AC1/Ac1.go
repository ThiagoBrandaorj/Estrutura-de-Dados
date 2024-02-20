package main

import (
	"fmt"
)

func calculaMedia(args...float64) float64{
	var sum = 0.0
	for i:= 0; i < len(args); i++{
		sum = sum + args[i]
	}
	return sum/float64(len(args))
}

func verificaParidade(n int) {
	if n % 2 == 0{
		fmt.Println("o número" , n , "é par")
	} else{
		fmt.Println("o número" , n , "é Ímpar")
	}
}

func minhaPotencia(base , expoente int) int {
	var res = 1
	for i:= 0; i < expoente; i++ {
		res = res * base
	}
	return res
}

func converteCelsiusParaFahrenheit(c float64) float64{
	var f = (9*c + 160)/5
	return f
}
func main(){
	fmt.Println(converteCelsiusParaFahrenheit(28))
	fmt.Println(calculaMedia(9.3, 7.8, 7.2))
	fmt.Println(minhaPotencia(4,3))
	verificaParidade(17)
	verificaParidade(30)
}