package idnum

import (
	"strconv"
	"strings"
)

type IdNum struct {
	Num int64
	Str string
}

var IdNumZero = IdNum{0, ""}

func NewIdNumFromStr(idStr string) IdNum {
	idStr = strings.TrimSpace(idStr)
	idNum, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return IdNumZero
	}
	return IdNum{idNum, idStr}
}

func NewIdNumFromBytes(idBytes []byte) IdNum {
	idStr := string(idBytes) // copy and create a string
	return NewIdNumFromStr(idStr)
}

func NewIdNum(idNum int64) IdNum {
	idStr := strconv.FormatInt(idNum, 10)
	return IdNum{idNum, idStr}
}

func (idn *IdNum) String() string {
	return idn.Str
}

func ParseIdNums(idNumsStr string) []IdNum {
	idNumStrArr := strings.Split(idNumsStr, ",")
	idNums := make([]IdNum, 0, len(idNumStrArr))
	for _, idNumStr := range idNumStrArr {
		if idNumStr == "" {
			continue
		}
		idNums = append(idNums, NewIdNumFromStr(idNumStr))
	}
	return idNums
}
