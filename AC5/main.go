package main

import (
	"fmt"
	"math"
)

func Triangulo(){
	var a,b,c float64
	array := [3] *float64{&a,&b,&c}
	var mensagem = ""
	fmt.Scan(&a,&b,&c)
	for i:= 0; i < len(array) - 1; i++{
		for j:= 0; j < len(array) - i - 1; j++{
			if *array[j] < *array[j + 1]{
				var temp = *array[j]
				array[j] = array[j+1]
				array[j+1] = &temp
			}
		}
	}
	if *array[0] >= *array[1] + *array[2]{
		mensagem += "NAO FORMA TRIANGULO \n"
	}else if math.Pow(*array[0],2) == math.Pow(*array[1],2) + math.Pow(*array[2],2){
		mensagem += "TRIANGULO RETANGULO \n"
	} else if math.Pow(*array[0],2) > math.Pow(*array[1],2) + math.Pow(*array[2],2){
		mensagem += "TRIANGULO OBTUSANGULO \n" 
	}else if math.Pow(*array[0],2) < math.Pow(*array[1],2) + math.Pow(*array[2],2){
		mensagem += "TRIANGULO ACUTANGULO \n"
	}

	if *array[0] == *array[1] && *array[1] == *array[2]{
		mensagem += "TRIANGULO EQUILATERO \n"
	} else if *array[0]==*array[1] || *array[0]==*array[2] || *array[1]==*array[2]{
		mensagem += "TRIANGULO ISOSCELES \n"
	}
	fmt.Println(mensagem)
}

func CrescimentoPopulacional() {
	var T int
	fmt.Scanln(&T)

	if T >= 1 && T <= 3000 {
		for i := 0; i < T; i++ {
			var PA, PB int
			var G1, G2 float64
			fmt.Scan(&PA, &PB, &G1, &G2)
			if PA >= 100 && PA <= 1000000 && PB >= 100 && PB <= 1000000 &&
				G1 >= 0.1 && G1 <= 10.0 && G2 >= 0.0 && G2 <= 10.0 {
				tempo := 0
				for PA <= PB {
					PA = int(math.Round(float64(PA) * (1 + G1/100)))
					PB = int(math.Round(float64(PB) * (1 + G2/100)))

					tempo++
					if tempo > 100 {
						fmt.Println("Mais de 1 seculo")
						break
					}
				}
				if tempo <= 100 {
					fmt.Println(tempo, " Anos")
				}
			}
		}
	}
}

func main() {
	// 1
	fmt.Println("Hello World!")
	// 2
	Triangulo()
	// 3
	CrescimentoPopulacional()
}
