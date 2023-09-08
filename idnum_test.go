package idnum

import (
	"encoding/json"
	"testing"
)

func Test_NewIdNumFromStr(t *testing.T) {
	idStr := " 42 " // spaces should be trimmed
	idNum := NewStrInt64FromStr(idStr)
	exp := NewStrInt64(42)
	if exp != idNum {
		t.Fail()
	}
	if int64(42) != idNum.Num {
		t.Fail()
	}
	if "42" != idNum.Str {
		t.Fail()
	}
	if "42" != idNum.String() {
		t.Fail()
	}
}

func Test_NewIdNumFromStr_invalid(t *testing.T) {
	idStr := " 42invalid000 "
	idNum := NewStrInt64FromStr(idStr)
	if int64(0) != idNum.Num {
		t.Fail()
	}
	if "" != idNum.Str {
		t.Fail()
	}
	if "" != idNum.String() {
		t.Fail()
	}
}

type User struct {
	Id StrInt64
}

func Test_UnmarshalJSON_from_int(t *testing.T) {
	u := &User{}
	err := json.Unmarshal([]byte(`{"Id":42}`), u)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	if int64(42) != u.Id.Num {
		t.Fail()
	}
	if "42" != u.Id.Str {
		t.Fail()
	}
}

func Test_UnmarshalJSON_from_str(t *testing.T) {
	u := &User{}
	err := json.Unmarshal([]byte(`{"Id":"42"}`), u)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	if int64(42) != u.Id.Num {
		t.Fail()
	}
	if "42" != u.Id.Str {
		t.Fail()
	}
}

func Test_MarshalJSON(t *testing.T) {
	u := &User{
		Id: NewStrInt64(42),
	}
	body, err := json.Marshal(u)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}

	if `{"Id":42}` != string(body) {
		t.Fail()
	}
}

func Test_IdNums(t *testing.T) {
	if len(ParseIdNums("")) != 0 {
		t.Fail()
	}

	if "42" != ParseIdNums(" 42 \n").String() {
		t.Fail()
	}

	idsStr := "1, 2, 3 ,"
	idNums := ParseIdNums(idsStr)
	if int64(1) != idNums[0].Num {
		t.Fail()
	}
	if int64(2) != idNums[1].Num {
		t.Fail()
	}
	if int64(3) != idNums[2].Num {
		t.Fail()
	}

	if "1,2,3" != idNums.String() {
		t.Fail()
	}

	int64s := idNums.ToInt64s()
	if int64s[0] != 1 {
		t.Fail()
	}
	if int64s[1] != 2 {
		t.Fail()
	}
	if int64s[2] != 3 {
		t.Fail()
	}
}
