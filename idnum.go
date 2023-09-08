package idnum

import (
	"strconv"
	"strings"
)

type StrInt64 struct {
	Num int64
	Str string
}

var StrInt64Zero = StrInt64{0, ""}

func NewStrInt64FromStr(idStr string) StrInt64 {
	idStr = strings.TrimSpace(idStr)
	idNum, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return StrInt64Zero
	}
	return StrInt64{idNum, idStr}
}

func NewStrInt64FromBytes(idBytes []byte) StrInt64 {
	idStr := string(idBytes) // copy and create a string
	return NewStrInt64FromStr(idStr)
}

func NewStrInt64(idNum int64) StrInt64 {
	idStr := strconv.FormatInt(idNum, 10)
	return StrInt64{idNum, idStr}
}

func (si *StrInt64) String() string {
	return si.Str
}

// UnmarshalJSON implements json.Unmarshaler
// We can deserialize both string form "42" and int 42.
func (si *StrInt64) UnmarshalJSON(data []byte) error {
	bytesNumOnly := data
	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
		bytesNumOnly = data[1 : len(data)-1]
	}
	idStr := string(bytesNumOnly)
	idNum, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}
	si.Num = idNum
	si.Str = idStr
	return nil
}

// MarshalJSON implements json.Marshaler
// The StrInt64 will always be serialized as an integer
func (si *StrInt64) MarshalJSON() ([]byte, error) {
	return []byte(si.Str), nil
}

type StrInt64s []StrInt64

func ParseIdNums(idNumsStr string) StrInt64s {
	idNumStrArr := strings.Split(idNumsStr, ",")
	idNums := make([]StrInt64, 0, len(idNumStrArr))
	for _, idNumStr := range idNumStrArr {
		if idNumStr == "" {
			continue
		}
		idNums = append(idNums, NewStrInt64FromStr(idNumStr))
	}
	return idNums
}

func (sis StrInt64s) ToInt64s() []int64 {
	idNums := make([]int64, 0, len(sis))
	for _, id := range sis {
		idNums = append(idNums, id.Num)
	}
	return idNums
}

func (sis StrInt64s) String() string {
	lenIds := len(sis)
	if lenIds == 0 {
		return ""
	}
	idNumsStr := sis[0].String()
	for i := 1; i < lenIds; i++ {
		idNumsStr += "," + sis[i].String()

	}
	return idNumsStr
}
