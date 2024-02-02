package calc

import (
	"testing"

	"github.com/TeamTestCodeowner/calc"
	rand "github.com/TeamTestCodeowner/rand"
	"github.com/stretchr/testify/require"
)

func TestSub(t *testing.T) {
	num1 := rand.RandomInt(1, 1000)
	num2 := rand.RandomInt(1, 1000)

	result := calc.Sub(num1, num2)
	require.Equal(t, result, num1-num2)
}
