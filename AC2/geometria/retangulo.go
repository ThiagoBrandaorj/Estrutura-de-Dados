package geometria


func CalculaAreaRetangulo(comprimento, largura *float64) float64{
	return (*comprimento) * (*largura)
}

func CalculaPerimetroRetangulo(comprimento, largura *float64) float64{
	return (*comprimento*2) + (*largura*2)
}
