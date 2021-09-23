package main

import (
	"entregable/codigo"
	"fmt"
)

func main() {
	encriptador := codigo.NuevoCodigo(2, 2, "NN", "TX")
	//codigoResultante, err := encriptador.Formatear("NN040001")
	//codigoResultante, err := encriptador.Formatear("TX03ABC")
	//codigoResultante, err := encriptador.Formatear("NN100987654321")
	//codigoResultante, err := encriptador.Formatear("TX02AB")
	//codigoResultante, err := encriptador.Formatear("TX06ABCDE")
	//codigoResultante, err := encriptador.Formatear("NN04000A")
	codigoResultante, err := encriptador.Formatear("NN040001")
	if err != nil {
		panic(err)
	}
	fmt.Println(codigoResultante)

}

//go run entregable.go
