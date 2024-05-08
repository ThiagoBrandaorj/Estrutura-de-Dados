package main

import "fmt"

type No struct {
    valor int  
    prox *No 
}

type Fila struct {
    tam int
    inicio *No  
    fim  *No
}

func (f *Fila) Inserir(nvalor int) {
    novoNo := &No{valor: nvalor}

    if f.inicio == nil {
        f.inicio = novoNo
        f.fim = f.inicio
    } else {
        f.fim.prox = novoNo
        f.fim = novoNo
    }
    f.tam++
}

func (f *Fila) Remover() error {
    if f.inicio == nil {
        return fmt.Errorf("A Fila está vazia")
    }

    if f.tam == 1{
        f.inicio = nil
        f.fim = nil
    } else{
        aux := f.inicio
        f.inicio = f.inicio.prox
        aux.prox = nil

        if f.inicio.prox == nil{
            f.fim = f.inicio
        }
    }

    f.tam--
    return nil
}

func (f *Fila) Imprimir() {
    if f.inicio == nil {
        fmt.Println("A fila está vazia")
    } else{
        no := f.inicio
        for no.prox != nil{
            fmt.Print(no.valor, " / ")
            no = no.prox
        }
        fmt.Println(no.valor)
    }
}

func main() {
    f := Fila{}
    f.Inserir(10)
    f.Inserir(20)
    f.Inserir(30)
    f.Inserir(40)
    fmt.Println("Elementos da fila:")
    f.Imprimir()
    fmt.Println("Removendo um elemento da fila:")
    f.Remover()
    fmt.Println("Elementos da fila após a remoção:")
    f.Imprimir()
}
