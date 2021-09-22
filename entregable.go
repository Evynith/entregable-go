package main

import (
	"entregable/codigo"
	"fmt"
)

func main() {

	//codigo := codigo.NuevoCodigo("NN040001")
	//codigo := codigo.NuevoCodigo("TX03ABC")
	codigo := codigo.NuevoCodigo("TX03ABC")
	codigo.Imprimir()
	codigoResultante := codigo.Formatear()
	fmt.Println(codigoResultante)

}

//go run entregable.go
