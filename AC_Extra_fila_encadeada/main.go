package main

import "fmt"

type No struct {
    data int  
    next *Node // Referência para o próximo nó na fila
}

type Fila struct {
    front *No // Referência para o primeiro nó da fila
    rear  *No // Referência para o último nó da fila
}

// Função para criar uma nova fila vazia
func NovaFila() *Fila {
    return &Fila{}
}

// Função para verificar se a fila está vazia
func (f *Fila) IsEmpty() bool {
    return f.front == nil
}

// Função para inserir um elemento na fila
func (q *Fila) Enqueue(data int) {
    newNode := &No{data, nil}

    if q.rear == nil {
        q.front = newNode
        q.rear = newNode
    } else {
        q.rear.next = newNode
        q.rear = newNode
    }
}

// Função para remover um elemento da fila
func (q *Fila) Dequeue() int {
    if q.IsEmpty() {
        panic("A fila está vazia")
    }

    data := q.front.data
    q.front = q.front.next

    if q.front == nil {
        q.rear = nil
    }

    return data
}

// Função para percorrer e imprimir os elementos da fila
func (q *Fila) Traverse() {
    if q.IsEmpty() {
        fmt.Println("A fila está vazia")
        return
    }

    current := q.front
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}

func main() {
    q := NewQueue()
    q.Enqueue(10)
    q.Enqueue(20)
    q.Enqueue(30)
    q.Enqueue(40)
    fmt.Println("Elementos da fila:")
    q.Traverse()
    fmt.Println("Removendo um elemento da fila:", q.Dequeue())
    fmt.Println("Elementos da fila após a remoção:")
    q.Traverse()
}
