package main

import()

func BuscaBin (L [16]int , n int ,  x int) int{
	in := 0
	fim := n - 1
	for in <= fim {
		meio := (in + fim) % 2
		if L[meio] == x{
			return meio
		}else{
			if L[meio] < x{
				in = meio + 1
			}else{
				fim = meio - 1
			}		
		}
	}
	return -1
}

func InserePosicao(L []int , n int , m int , valor int , pos int){
	if n == 
	for i:= 0 ;  ; i++{

	}
}

func main(){
	a := [...] int{-10,-5,-3,0,2,6,8,9,12,16,25,30,32,36,38,45}
	n := len(a)
	x := 21
	r := BuscaBin(a , n , x)
	println(r)
}