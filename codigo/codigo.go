package codigo

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Resultado struct {
	Tipo  string
	Largo int
	Valor string
}

type CodigoTipo struct {
	id               string
	serie            string
	valoresAsociados string
	extension        int
}

type codigo struct {
	largoValor int
	tipos      []CodigoTipo
}

/*codigo
 * NuevoCodigo inicializa el tipo codigo con el valor pasado por parámetro, dentro se definen el largo del tipo como 2 (ocupando los 2 primeros lugares) y la cantidad de cifras de la definición de la cantidad de valores del código (ocupando tantos lugares como se indica en la inicializacion luego de la definicion de tipo)
 */
func NuevoCodigo(lrt int, tipos []CodigoTipo) codigo {
	return codigo{lrt, tipos}
}

/*
 * NuevoTipo inicializa un nvo. tipo con sus valores predefinidos y su extensión
 */
func NuevoTipo(id string, serie string, valoresAsociados string, extension int) CodigoTipo {
	return CodigoTipo{id, serie, valoresAsociados, extension}
}

/*
 * es se retorna a sí mismo si se contiene en el strig recibido
 */
func (c *CodigoTipo) es(elemento string) (bool, error) {
	fin := c.extension
	if len(elemento) < fin {
		return false, errors.New("El código ingresado no tiene la longitud requerida")
	}
	tipo := elemento[:fin]
	if strings.Compare(c.serie, tipo) == 0 {
		return true, nil
	}
	return false, errors.New("El código ingresado no es válido")
}

/*
 * Formatear en caso de ser un string válido le dá un formato definido y lo devuelve
 */
func (c *codigo) Formatear(contenido string) (Resultado, error) {
	caracteresValidos := `^[A-Z0-9]+$`
	if valido, err := contieneCaracteresValidos(contenido, caracteresValidos); err != nil || valido == false {
		return Resultado{}, err
	}
	tipo, err := c.obtenerTipo(contenido)
	if err != nil {
		return Resultado{}, err
	}
	largo, err := c.obtenerLargo(tipo, contenido)
	if err != nil {
		return Resultado{}, err
	}
	valor, err := c.obtenerValor(contenido, largo, tipo)
	if err != nil {
		return Resultado{}, err
	}
	return Resultado{tipo.serie, largo, valor}, nil
}

/*
 * contieneCaracteresValidos verifica que el string contenga caracteres válidos
 */
func contieneCaracteresValidos(cadena, caracteres string) (bool, error) {
	caracteresValidos := regexp.MustCompile(caracteres) //suponiendo que deba tener sólo mayúsculas
	resultadoComparacion := caracteresValidos.MatchString(cadena)
	if resultadoComparacion == false {
		return false, errors.New("El código ingresado no es válido")
	}
	return resultadoComparacion, nil
}

/*
 * devolverTipos busca entre lso tipos definidos si exste una coincidencia y la devuelve
 */
func (c *codigo) obtenerTipo(s string) (CodigoTipo, error) {
	err := errors.New("El codigo ingresado no contiene un tipo válido")
	for _, tipo := range c.tipos {
		if rta, e := tipo.es(s); rta != true {
			err = e
		} else {
			return tipo, nil
		}
	}
	return CodigoTipo{}, err
}

/*
 * obtenerLargo retorna el contenido de los 2 siguientes caracteres representativos al tipo que representan el largo del código resultante
 */
func (c *codigo) obtenerLargo(tipo CodigoTipo, contenido string) (int, error) {
	inicio := tipo.extension
	fin := inicio + c.largoValor
	if len(contenido) < fin {
		return 0, errors.New("El código ingresado no tiene la longitud requerida")
	}
	largo := contenido[inicio:fin]
	sv, err := strconv.Atoi(largo)
	return sv, err
}

/*
 * obtenerValor retorna el contenido de los siguientes n caracteres del representativo al largo que representa el código resultante
 */
func (c *codigo) obtenerValor(contenido string, cantidad int, tipo CodigoTipo) (string, error) {
	inicio := tipo.extension + c.largoValor
	fin := inicio + cantidad
	fmt.Println(cantidad, inicio, fin)
	if len(contenido) < fin {
		return "", errors.New("El código ingresado no tiene la longitud requerida")
	}
	valor := contenido[inicio:fin]
	if resultado, err := contieneCaracteresValidos(valor, tipo.valoresAsociados); resultado == false || err != nil {
		return "", errors.New("El código ingresado no contiene el formato requerido")
	}
	return valor, nil
}
