# Trabajo Práctico
## TUDAI 2021, Seminario GO

## Solución :thinking: :
El ejercicio requería que se ingrese un código en string del cual hay que sacar un tipo de formato específico, siendo:
- 2 primeros caracteres para el tipo
- 2 siguientes caracteres para el largo del código
- n siguientes caracteres, a partir del largo, para el código resultante
- debe revisarse que el valor cumpla con el tipo dado, por ejemplo para TX el valor debe ser numérico y para NN el valor debe ser alfabético

Por ejemplo:

````
TX03ABC
````

Deberá generar una estructura con los siguientes valores:

````
{TX 3 ABC}
````

Para ello creé un módulo específico para las funcionalidades pertinentes al manejo del código recibido. En él hay dos tipos, `Resultado` que es quién se pasa devuelto con el contenido ya fraccionado según corresponde y `Codigo` que contiene el string a procesar, las variables que definen la estructura del código resultante y las funciones requeridas para esta tarea.

Las funciones existentes en el módulo son las siguientes:

- `NuevoCodigo(c string)` : inicializa el tipo con los datos a utilizar que recibe desde el main, recibe un string y define la estructura que deberá tener el resultado. Es público ya que debe ser llamado desde afuera
- `Imprimir()` : imprime el valor asignado. Es público ya que debe ser llamado desde afuera
- `Formatear()` : obtiene los elementos del tipo `Resultado` uno por uno y verifica errores. Es público ya que necesita ser llamado desde afuera
- `contieneCaracteresValidos()` : verifica que los caracteres que se hayan ingresado entén entre los válidos (números del 0 al 9 y letras mayúsculas, según sea el caso). Es privada ya que la funcion que lo utiliza está en el módulo
- `obtenerTipo()` : devuelve los caracteres correspondientes al tipo del código, si es correcto. Es privada ya que la funcion que lo utiliza está en el módulo
- `obtenerLargo()` : decodigovuelve la cantidad de caracteres que deberá tener el código, si es correcto. Es privada ya que la funcion que lo utiliza está en el módulo
- `obtenerValor(cantidad int)` : recibe la cantidad de caracteres que deberá tener el valor del código resultante y lo devuelve, si es correcto. Es privada ya que la funcion que lo utiliza está en el módulo

En cada una de las funciones privadas específicas de obtención se accede a los sectores del string pasandole los índices de inicio y fin para cada sección, se verifica que no se quiera obtener algo que sobrepase el límite del tamaño del código dado y que cumpla con características del mismo (como el tipo y que los valores correspondan a este)

Para la resolución de este TP se debieron usar los paquetes `"strconv"`, para convertir el string del largo del contenido a un int, `"fmt"` para hacer las impresiones en consola, `"regexp"` para simplificar la comparación del string ingresado con un conjunto de caracteres dados, `"testing"` y `"github.com/stretchr/testify/assert"` para realizar los testeos de unidad.

En el test unitario lo que se hace es tomar una función que se quiera probar, creamos los datos necesarios y preparamos los resultados de deben dar, se la hace correr y luego se comparan los resultados obtenidos con los que se deberían obtener.
Para este TP se espera una cobertura del 80% o más.

## Autora :writing_hand: :
**Evelyn Vega** [EvelynVega](https://github.com/Evynith)