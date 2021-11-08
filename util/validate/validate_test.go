package validate

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValuesHasZero(t *testing.T) {
	vi1 := 1
	var vi2 int
	r1 := ValuesHasZero(vi1)
	r2 := ValuesHasZero(vi2)
	require.True(t, r2)
	require.False(t, r1)
	vs1 := "1"
	vs2 := ""
	r1 = ValuesHasZero(vs1)
	r2 = ValuesHasZero(vs2)
	require.True(t, r2)
	require.False(t, r1)
	type s1 struct {
		i int
	}
	vb1 := s1{
		i: 1,
	}
	vb2 := s1{}
	r1 = ValuesHasZero(vb1)
	r2 = ValuesHasZero(vb2)
	require.True(t, r2)
	require.False(t, r1)
}

func TestStringToDate(t *testing.T) {
	res := StringToDate("11")
	require.True(t, res.IsZero())
}
