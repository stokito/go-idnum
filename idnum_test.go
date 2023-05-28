package idnum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewIdNumFromStr(t *testing.T) {
	idStr := " 42 " // spaces should be trimmed
	idNum := NewIdNumFromStr(idStr)
	exp := NewIdNum(42)
	assert.Equal(t, exp, idNum)
	assert.Equal(t, int64(42), idNum.Num)
	assert.Equal(t, "42", idNum.Str)
	assert.Equal(t, "42", idNum.String())
}

func Test_NewIdNumFromStr_invalid(t *testing.T) {
	idStr := " 42invalid000 "
	idNum := NewIdNumFromStr(idStr)
	assert.Equal(t, int64(0), idNum.Num)
	assert.Equal(t, "", idNum.Str)
	assert.Equal(t, "", idNum.String())
}
