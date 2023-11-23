package calc

import (
	"hello-build-with-golang2/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSub(t *testing.T) {
	num1 := util.RandomInt(1, 1000)
	num2 := util.RandomInt(1, 1000)

	result := Sub(num1, num2)
	require.Equal(t, result, num1-num2)
}
