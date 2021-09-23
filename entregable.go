package main

import (
	"entregable/codigo"
	"fmt"
)

func main() {
	codigoTipo1 := codigo.NuevoTipo("numerico", "NN", `^[0-9]+$`, 2)
	codigoTipo2 := codigo.NuevoTipo("alfabetico", "TX", `^[A-Z]+$`, 2)
	codTipos := make([]codigo.CodigoTipo, 0)
	codTipos = append(codTipos, codigoTipo1)
	codTipos = append(codTipos, codigoTipo2)
	encriptador := codigo.NuevoCodigo(2, codTipos)
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
