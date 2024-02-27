package main

import(
	"fmt"
)

type Contato struct {
	Nome  string
	Email string
}

func (cont *Contato) AlterarEmail(nEmail string) { 
	cont.Email = nEmail
}

func main(){
	var arraycontatos [5]Contato

	for{
		var opcao int
		fmt.Println("clique 1 para Adicionar um contato, 2 para excluir ou 3 para sair: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var contato Contato
			fmt.Println("Digite o nome do contato: ")
			fmt.Scan(&contato.Nome)
			fmt.Println("Digite o email do contato: ")
			fmt.Scan(&contato.Email)
			adicionarContato(contato,&arraycontatos)
			fmt.Printf("arraycontatos: %v\n", arraycontatos)
		
		case 2:
			excluirContato(&arraycontatos)
			fmt.Printf("arraycontatos: %v\n", arraycontatos)

		case 3:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida , digite 1 , 2 ou 3 ..")
		}

	}

}

func adicionarContato(cont Contato, a*[5]Contato){
	for i:= 0; i < len(a) ; i++{
		if a[i].Nome == "" && a[i].Email == ""{
			a[i] = cont
			break
		}
	}
}

func excluirContato(a*[5]Contato){
	for i:= len(a) - 1; i >= 0 ; i--{
		if a[i].Nome != "" && a[i].Email != ""{
			a[i] = Contato{}
			break
		}
	}
}