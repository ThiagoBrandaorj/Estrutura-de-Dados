package main

import("fmt")

// 1
func torreHanoi(n int , A[]int , B[]int , C[]int) int{
	movimentos := make([]int,0)
	j:= 0
	if n > 0 && A == nil{
		for i:= 0 ;i<n; i++{
			A[i] = n - i
		}
		return movimentos[j] + torreHanoi(n - 1, A , C , B)
	}else if n == 2{
		for i:= 0 ; i < 2; i++{
			A[i] = 
			movimentos[j] ++
		}  
	}else if n == 0{
		return 0
	} else{
		fmt.Println("não deveria chegar aqui")
	}
}

// 2
func aleatoria(a []int , valor int){
	var soma int
	for i:= 0 ; i < len(a); i++{
		for j:= 0 ; j < len(a) - 1 ; j++
	}
}

func fatorial(n int) int{
	if n > 0{
		return n * fatorial(n - 1)
	}else if n == 0{
		return 1
	}else{
		fmt.Println("Digite um numero inteiro")
		return -1
	}
}

func main() {
	var x,n, tam int
	var a []int
	var b []int
	var c []int
	fmt.Println("Digite um número inteiro:")
	_, err := fmt.Scanln(&x)

	if err != nil {
		fmt.Println("Erro ao ler o número:", err)
		return
	}

	result := fatorial(x)
	if result != -1 {
		fmt.Println("Fatorial de", x, "é:", result)
	}

	fmt.Println("Digite o número de peças para o algoritimo de hanoi: ")
	fmt.Scanln(&n)
	torreHanoi(n,a,b,c)
	
	fmt.Println("Digite o tamanho do array ordenado que deseja: ")
	fmt.Scanln(&tam)
	array := [tam] int{for i:=0;i<}
	for i:= 0 
}