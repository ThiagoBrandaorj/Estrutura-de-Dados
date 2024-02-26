package main

import (
)

type Contato struct {
	Nome  string
	email string
}

func (cont *Contato) AlterarEmail(nEmail string) { 
	cont.email = nEmail
}