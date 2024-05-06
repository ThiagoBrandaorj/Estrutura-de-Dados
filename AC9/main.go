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
	var alfabeto = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Scan(&N)
	if N > 0{
		for i:= 0; i < N; i++{
			fmt.Scan(&sentenca)
			if len(sentenca) >= 1 && len(sentenca) <= 50{
				fmt.Scan(&deslocamento)
				if deslocamento >= 0{
					novastr := ""
					for i:= 0; i< len(sentenca); i++{
						char := sentenca[i]
						if char >= 'A' && char <= 'Z'{
							index := strings.IndexByte(alfabeto,char)
							novoindex := (index - deslocamento + 26) % 26
							novastr += string(alfabeto[novoindex])
						}
					}
					fmt.Println(novastr)
				}
			}
		}
	}
}
//4
func PapaiNoel(){
	var entrada string
	paises:= map[string]string{
		"brasil":     "Feliz Natal!",
        "alemanha":      "Frohliche Weihnachten!",
        "austria":      "Frohe Weihnacht!",
        "coreia":    "Chuk Sung Tan!",
        "espanha":      "Feliz Navidad!",
        "greci":  "Kala Christougena!",
        "estados-unidos":      "Merry Christmas!",
        "inglaterra":      "Merry Christmas!",
        "australia":      "Merry Christmas!",
        "portugal":      "Feliz Natal!",
        "suecia":      "God Jul!",
        "turquia":      "Mutlu Noeller",
        "argentina":      "Feliz Navidad!",
        "chile":      "Feliz Navidad!",
        "mexico":      "Feliz Navidad!",
        "antardida":      "Merry Christmas!",
        "canada":      "Merry Christmas!",
        "irlanda":      "Nollaig Shona Dhuit!",
        "belgica":      "Zalig Kerstfeest!",
        "italia":      "Buon Natale!",
        "libia":      "Buon Natale!",
        "siria":      "Milad Mubarak!",
        "marrocos":      "Milad Mubarak!",
        "japao":     "Merii Kurisumasu!",
	}
	for{
	fmt.Scan(&entrada)
	if cumprimento , ok:= paises[entrada]; ok{
		fmt.Println(cumprimento)
	}else{
		fmt.Println("--- NOT FOUND ---")
	}
   } 
}
func main(){
	//RevisaoContrato()
	//Led()
	//CifraCesar()
	//PapaiNoel()
}