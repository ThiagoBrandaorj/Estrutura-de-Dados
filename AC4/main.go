package main

import("fmt")

// 1
func torreHanoi(n int , A , B , C string){
	if n > 0{
		torreHanoi(n - 1 , A , C , B)
		fmt.Println("Mover o disco do topo de " , A ," para " , B)
		torreHanoi(n - 1 , C , B , A)
	}
}

// 2
func encontraPar(a []int , valor int) (int,int){
	for i:= 0 ; i < len(a); i++{
		for j:= 1 ; j < len(a) ; j++{
			if a[i] + a[j] == valor{
				return a[i] , a[j]
			}
		}
	}
	return -1,-1
}

func main() {
	A , B , C := "A" , "B" , "C"
	n := 3
	torreHanoi(n , A , B , C)
	array := []int{1,2,4,5,8,11,12}
	alvo := 16
	n1 , n2 := encontraPar(array, alvo)
	fmt.Println(n1 , n2)
}