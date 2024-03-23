package main

import "fmt"

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
				num := &res_numerador
				den := &res_denominador
				for *den != 0{
					*num,*den = *den, *num % *den
				}
				mdc := &num
				res_numerador_simplificado := res_numerador / mdc
				res_denominador_simplificado := res_denominador / mdc
				fmt.Print(res_numerador + )

			 }
		}
	}
}
func main(){

}
