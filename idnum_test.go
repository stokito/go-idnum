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

func Test_IdNums(t *testing.T) {
	assert.Empty(t, ParseIdNums(""))
	assert.Equal(t, "42", ParseIdNums(" 42 \n").String())
	idsStr := "1, 2, 3 ,"
	idNums := ParseIdNums(idsStr)
	assert.Equal(t, int64(1), idNums[0].Num)
	assert.Equal(t, int64(2), idNums[1].Num)
	assert.Equal(t, int64(3), idNums[2].Num)
	assert.Equal(t, "1,2,3", idNums.String())
	assert.Equal(t, []int64{1, 2, 3}, idNums.ToInt64s())
}
