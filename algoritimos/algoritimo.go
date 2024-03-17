package main

import "fmt"

func tomadas() {
	var t1, t2, t3, t4 int8
	fmt.Scanln(&t1, &t2, &t3, &t4)
	fmt.Println(t1 + t2 + t3 + t4 - 3)
}

func meteoritos() {
	var x1, x2, y1, y2, x, y, n int
	num := make([]int, 0)
	i := 0
	for {
		fmt.Scanln(&x1, &y1, &x2, &y2)
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 { break }
		fmt.Scanln(&n)
		num = append(num, 0)
		for j := 1; j <= n; j++ {
			fmt.Scanln(&x, &y)
			if x1 <= x && x <= x2 && y2 <= y && y <= y1 {
				num[i]++
			}
		}
		i++
	}
	for j, num_meteoritos := range num {
		fmt.Println("Teste", j + 1)
		fmt.Println(num_meteoritos)
		fmt.Println("")
	}
}

func maiorNumero() {
	var numero int
	maior := 0

	for {
		fmt.Scan(&numero)
		if numero == 0 { break }
		if numero > maior {
			maior = numero
		}
	}
	fmt.Println(maior)
}

func CartasOrdem() {
    var n1, n2 int
    classificacao := ""

    for {
        fmt.Scan(&n2)

        if n1 == 0 {
            n1 = n2
            continue
        }

        if n2 > n1 {
            if classificacao == "" {
                classificacao = "C"
            } else if classificacao == "D" {
                classificacao = "N"
                break
            }
        }

        if n2 < n1 {
            if classificacao == "" {
                classificacao = "D"
            } else if classificacao == "C" {
                classificacao = "N"
                break
            }
        }

        n1 = n2
    }

    fmt.Println(classificacao)
}

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

/*func InserePosicao(L []int , n int , m int , valor int , pos int){
	if n == 
	for i:= 0 ;  ; i++{

	}
}
*/

func main() {
	// tomadas()
	//meteoritos()
	CartasOrdem()
	//maiorNumero()
	a := [...] int{-10,-5,-3,0,2,6,8,9,12,16,25,30,32,36,38,45}
	n := len(a)
	x := 21
	r := BuscaBin(a , n , x)
	println(r)
}