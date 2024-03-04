package ac4

import()

func algoritimo_Hanoi(n int , A , B , C rune){

}

func buscaMatriz(m [][]int, n , k int)bool{
	i:= 0
	var j int 
	for i < n {
		j = 0
		for j < n{
			if m[i][j] == k{
				return true
			}
			j ++
		}
		i++
	}
	return false
}

func somaPar(a array, tamanho int,alvo int){
	i:= 0
	for i < tamanho - 1{
		j:= i + 1
		for j <  tamanho{
			if a[i] + a[j] = alvo{
				return a[i], a[j]
			}
			j++
		}
		i++
	}
	return (-1,-1)
}

func main(){
	matriz := [][]int := {[1,2],[2,4]}
	buscaMatriz([][],4,2)
}