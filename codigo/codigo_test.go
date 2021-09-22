package codigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatear(t *testing.T) {

	codigo := NuevoCodigo("TX03ABC")
	resultado := Resultado{"TX", 03, "ABC"}
	codigoResultante := codigo.Formatear()
	assert.Equal(t, codigoResultante, resultado)
}
