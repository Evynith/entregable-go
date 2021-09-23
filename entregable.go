package main

import (
	"entregable/codigo"
	"fmt"
)

func main() {

	//codigo := codigo.NuevoCodigo("NN040001")
	//codigo := codigo.NuevoCodigo("TX03ABC")
	//codigo := codigo.NuevoCodigo("NN100987654321")
	//codigo := codigo.NuevoCodigo("TX02AB")
	//codigo := codigo.NuevoCodigo("TX06ABCDE")
	codigo := codigo.NuevoCodigo("NN04000A")
	//codigo := codigo.NuevoCodigo("NN040001")
	codigo.Imprimir()
	codigoResultante, err := codigo.Formatear()
	if err != nil {
		panic(err)
	}
	fmt.Println(codigoResultante)

}

//go run entregable.go
