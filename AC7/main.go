package main

import (
    "fmt"
)
func mdc(num, den int) int {
    for den != 0 {
        num, den = den, num%den
    }
    return num
}
// 1
func tdaracional(){
	var N int
	fmt.Scanln(&N)
	if N >= 1 && N <= 1000{
		for i:= 0; i < N; i++{
			var n1,n2,d1,d2 int
			var op rune
			var res_numerador , res_denominador int
			fmt.Scanln(&n1,&d1,&op,&n2,&d2)
            if n1 >= 1 && n1 <= 1000 &&      
            n2 >= 1 && n2 <= 1000 &&        
            d1 >= 1 && d1 <= 1000 &&        
            d2 >= 1 && d2 <= 1000 &&        
            (op == '/' || op == '+' ||      
             op == '-' || op == '*') {
				if op == '+'{
					res_numerador = ((n1 * d2) + (n2 * d1))
					res_denominador = d1 * d2
				} else if op == '-'{
					res_numerador = ((n1*d2)-(n2*d1))
					res_denominador = d1*d2
				} else if op == '*'{
					res_numerador = (n1 * n2)
					res_denominador = d1 * d2
				} else if op == '/'{
					res_numerador = n1/d1
					res_denominador = n2/d2
				}
				mdc := mdc(res_numerador,res_denominador)
				res_numerador /= mdc
                res_denominador /= mdc
        
                fmt.Printf("%d/%d = %d/%d\n", res_numerador,res_denominador,res_numerador,res_denominador)
			}
		}
	}
}

// 2
func ricoDoMate(){
    var n int
	var l, q float32

	fmt.Scanln(&n, &l, &q)

	participantes := make([]string,n)
	for i := 0; i < n; i++ {
		fmt.Scan(&participantes[i])
	}

	i := 0
	for l - q > 0{
		l -= q
		i++
	}
	i %= n
	fmt.Printf("%s %.1f\n", participantes[i], l)
}

// 3
func calculoSimples(){
    var cod1,cod2, n1,n2 int
    var valor1,valor2 float32
    fmt.Scanln(&cod1,&n1,&valor1)
    fmt.Scanln(&cod2,&n2,&valor2)
    total := (float32(n1) * valor1) + (float32(n2)*valor2)
    fmt.Printf("VALOR A PAGAR: R$ %.2f", total)
}

func main() {
	
	tdaracional()
    //ricoDoMate()
    //calculoSimples()
}
