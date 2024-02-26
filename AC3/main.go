package main

import(
	"fmt"
)

func main(){
	fmt.Println("Interface")
	var arraycontato [5]Contato
	fmt.Println("Você deseja adicionar ou excluir ?")
	contato1 := Contato{}
	contato1.Nome = "Thiago"
	contato1.email = "seila@gmail.com"
	contato1.AlterarEmail(contato1.email,"blablabla@gmail.com")
}

func adicionarContato(cont *Contato, a[] string){
	for i:= 0; i < len(a) ; i++{
		if a[i] == Contato{}{
			a[i] = cont
		}
	}
}

func excluirContato(a[] string){
	for i:= len(a); i > 0 ; i--{
		if a[i] =! Contato{}{
			
		}
	}
}