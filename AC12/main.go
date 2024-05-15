package main

import (
	"fmt"
)

//1
func Alarme(){
    for {
        var H1, M1, H2, M2 int
        fmt.Scan(&H1, &M1, &H2, &M2)
        if H1 == 0 && M1 == 0 && H2 == 0 && M2 == 0 {
            break
        }
        minutos := H1*60 + M1
        alarme := H2*60 + M2
        tempoMinutos := alarme - minutos

        if tempoMinutos <= 0 {
          tempoMinutos += 24 * 60
        }
        fmt.Println(tempoMinutos)
    }
}
//2
func TopN() {
	var K int
	fmt.Scan(&K)
	switch {
	case K == 1:
		fmt.Println("Top 1")
	case K <= 3:
		fmt.Println("Top 3")
	case K <= 5:
		fmt.Println("Top 5")
	case K <= 10:
		fmt.Println("Top 10")
	case K <= 25:
		fmt.Println("Top 25")
	case K <= 50:
		fmt.Println("Top 50")
	case K <= 100:
		fmt.Println("Top 100")
	default:
		fmt.Println("Erro")
	}
}
//3
func Escada() {
	var N, H, C, L int

    for {
        fmt.Scan(&N)
        
        fmt.Scan(&H, &C, &L)
        
        // Área que cobre os degraus
        areaDegraus := N * (C * L)
        
        // Área de transição
        alturaTransicao := H * (N - 1)
        areaTransicao := alturaTransicao * C
        
        // Área total
        areaTotal := areaDegraus + areaTransicao
        
        // Convertendo para metros quadrados
        areaTotalMetrosQuadrados := float64(areaTotal) / 10000 // 100 cm² = 1 m²
        
        // Imprimir com 4 casas decimais
        fmt.Printf("%.4f\n", areaTotalMetrosQuadrados)
    }
}
func main() {
	//Alarme()
	//TopN()
	Escada()
}