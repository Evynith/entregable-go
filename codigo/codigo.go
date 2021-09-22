package codigo

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Resultado struct {
	Tipo  string
	Largo int
	Valor string
}

type codigo struct {
	contenido  string
	largoTipo  int
	largoValor int
}

/*
 * NuevoCodigo inicializa el tipo codigo con el valor pasado por parámetro, dentro se definen el largo del tipo como 2 (ocupando los 2 primeros lugares) y la cantidad de cifras de la definición de la cantidad de valores del código (ocupando tantos lugares como se indica en la inicializacion luego de la definicion de tipo)
 */
func NuevoCodigo(c string) codigo {
	return codigo{c, 2, 2} //ambos tamaños podrían ser pasados por parámetro para que no sea algo fijo, en caso de que a futuro cambie ¿?
}

/*
 * Imprimir imprime por patalla el código contenido
 */
func (c *codigo) Imprimir() {
	fmt.Println(c.contenido)
}

/*
 * Formatear en caso de ser un string válido le dá un formato definido y lo devuelve
 */
func (c *codigo) Formatear() (Resultado, error) {
	if valido, err := c.contieneCaracteresValidos(); err != nil || valido == false {
		return Resultado{}, err
	}
	tipo, err := c.obtenerTipo()
	if err != nil {
		return Resultado{}, err
	}
	largo, err := c.obtenerLargo()
	if err != nil {
		return Resultado{}, err
	}
	valor, err := c.obtenerValor(largo)
	if err != nil {
		return Resultado{}, err
	}
	return Resultado{tipo, largo, valor}, nil
}

/*
 * contieneCaracteresValidos verifica que el string contenga caracteres válidos
 */
func (c *codigo) contieneCaracteresValidos() (bool, error) {
	caracteresValidos := regexp.MustCompile(`^[A-Z0-9]+$`) //suponiendo que deba tener sólo mayúsculas
	resultadoComparacion := caracteresValidos.MatchString(c.contenido)
	if resultadoComparacion == false {
		return false, errors.New("El código ingresado no es válido")
	}
	return resultadoComparacion, nil
}

/*
 * obtenerTipo retorna los 2 primeros caracteres de la secuencia de código dado que representan el tipo del código resultante
 */
func (c *codigo) obtenerTipo() (string, error) {
	fin := c.largoTipo
	if len(c.contenido) < fin {
		return "", errors.New("El código ingresado no tiene la longitud requerida")
	}
	return c.contenido[:fin], nil
}

/*
 * obtenerLargo retorna el contenido de los 2 siguientes caracteres representativos al tipo que representan el largo del código resultante
 */
func (c *codigo) obtenerLargo() (int, error) {
	inicio := c.largoTipo
	fin := c.largoTipo + c.largoValor
	if len(c.contenido) < fin {
		return 0, errors.New("El código ingresado no tiene la longitud requerida")
	}
	largo := c.contenido[inicio:fin]
	sv, err := strconv.Atoi(largo)
	return sv, err
}

/*
 * obtenerValor retorna el contenido de los siguientes n caracteres del representativo al largo que representa el código resultante
 */
func (c *codigo) obtenerValor(cantidad int) (string, error) {
	inicio := c.largoTipo + c.largoValor
	fin := c.largoTipo + c.largoValor + cantidad
	if len(c.contenido) < fin {
		return "", errors.New("El código ingresado no tiene la longitud requerida")
	}
	return c.contenido[inicio:fin], nil
}
