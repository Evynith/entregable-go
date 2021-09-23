package codigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatear(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{Input: "TX02AB", Success: true, Type: "TX", Value: "AB", Length: 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	encriptador := NuevoCodigo(2, 2, "NN", "TX")

	for _, testData := range cases {
		_, err := encriptador.Formatear(testData.Input)
		assert.Equal(t, err == nil, testData.Success)
	}
}

// go test ./... -count=1 -cover
