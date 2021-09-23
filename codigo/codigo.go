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
	caracteresValidos := `^[A-Z0-9]+$`
	if valido, err := contieneCaracteresValidos(c.contenido, caracteresValidos); err != nil || valido == false {
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
	if tipo == "NN" {
		caracteresValidos = `^[0-9]+$`
	} else { //tipo TX
		caracteresValidos = `^[A-Z]+$`
	}
	valor, err := c.obtenerValor(largo, caracteresValidos)
	if err != nil {
		return Resultado{}, err
	}
	return Resultado{tipo, largo, valor}, nil
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
 * obtenerTipo retorna los 2 primeros caracteres de la secuencia de código dado que representan el tipo del código resultante
 */
func (c *codigo) obtenerTipo() (string, error) {
	fin := c.largoTipo
	if len(c.contenido) < fin {
		return "", errors.New("El código ingresado no tiene la longitud requerida")
	}
	tipo := c.contenido[:fin]
	if strings.Compare(tipo, "NN") != 0 && strings.Compare(tipo, "TX") != 0 {
		return "", errors.New("El codigo ingresado no contiene un tipo válido")
	}
	return tipo, nil
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
func (c *codigo) obtenerValor(cantidad int, caracteresValidos string) (string, error) {
	inicio := c.largoTipo + c.largoValor
	fin := c.largoTipo + c.largoValor + cantidad
	if len(c.contenido) < fin {
		return "", errors.New("El código ingresado no tiene la longitud requerida")
	}
	valor := c.contenido[inicio:fin]
	if resultado, err := contieneCaracteresValidos(valor, caracteresValidos); resultado == false || err != nil {
		return "", errors.New("El código ingresado no contiene el formato requerido")
	}
	return valor, nil
}
