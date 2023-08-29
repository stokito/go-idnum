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

// UnmarshalJSON implements json.Unmarshaler
// We can deserialize both string form "42" and int 42.
func (idn *IdNum) UnmarshalJSON(data []byte) error {
	bytesNumOnly := data
	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
		bytesNumOnly = data[1 : len(data)-1]
	}
	idStr := string(bytesNumOnly)
	idNum, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}
	idn.Num = idNum
	idn.Str = idStr
	return nil
}

// MarshalJSON implements json.Marshaler
// The IdNum will always be serialized as an integer
func (idn *IdNum) MarshalJSON() ([]byte, error) {
	return []byte(idn.Str), nil
}

type IdNums []IdNum

func ParseIdNums(idNumsStr string) IdNums {
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

func (ids IdNums) ToInt64s() []int64 {
	idNums := make([]int64, 0, len(ids))
	for _, id := range ids {
		idNums = append(idNums, id.Num)
	}
	return idNums
}

func (ids IdNums) String() string {
	lenIds := len(ids)
	if lenIds == 0 {
		return ""
	}
	idNumsStr := ids[0].String()
	for i := 1; i < lenIds; i++ {
		idNumsStr += "," + ids[i].String()

	}
	return idNumsStr
}
