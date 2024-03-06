package main

import("fmt")

func torreHanoi(n int , A[]int , B[]int , C[]int){
	if n > 0{
		for i:= 0 ;i<n; i++{
			A[i] == n - i
		}
		return torreHanoi(n - 1, A , C , B)

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
	var x int
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
}