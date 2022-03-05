package common
import (
	"testing"
	"fmt"
)

func TestBoolToInt(t *testing.T) {
	var var1 = BoolToInt(false);
	if var1 > 0 {
		t.Fail();
	}
}

func TestPadRight(t *testing.T) {
	var var1 = "test";
	var var2 = PadRight(var1, len(var1) + 1, "#");
	if len(var2) != len(var1) + 1 {
		t.Fail();
	}
}

func TestCleanString(t *testing.T) {
	var var1 = "test\x00";
	var var2 = CleanString(var1);
	if var2 == var1 {
		t.Fail();
	}
}

func TestPadLeft(t *testing.T) {
	var var1 = "test";
	var var2 = PadLeft(var1, len(var1) + 2, "#");
	if len(var2) != len(var1) + 2 {
		t.Fail();
	}
}

func TestCStr(t *testing.T) {
	var var1 = "test";
	var var2 = CStr(var1);
	if var1 == fmt.Sprintf("%v", *var2) {
		t.Log(fmt.Sprintf("%v", *var2));
		t.Fail();
	}
}

func TestRunCmdNoWait(t *testing.T) {
	// Do not implement as this is a runaway invocation
}

func TestRunCmd(t *testing.T) {
	var var1 = "test";
	var var2 = RunCmd("echo " + var1);
	if var2 != var1 + "\n" {
		t.Log(var2);
		t.Fail();
	}
}

func TestArgsToDictionary(t *testing.T) {
	var var1 = []string{"-p", "test"};
	var var2 = ArgsToDictionary(var1);
	if var2["-p"] == "" {
		t.Fail();
	}
}

func TestStrToDictionary(t *testing.T) {
	var var1 = `{"var":"test"}`;
	var var2 = StrToDictionary([]byte(var1));
	if var2["var"] == nil {
		t.Fail();
	}
}

func TestDictionaryToJsonString(t *testing.T) {
	var var1 = map[string]interface{}{"var":"test",};
	var var2 = DictionaryToJsonString(var1);
	if var2 != `{"var":"test"}` {
		t.Log(var2);
		t.Fail();
	}
}

func TestStrDictionaryToJsonString(t *testing.T) {
	var var1 = map[string]string{"var":"test", };
	var var2 = StrDictionaryToJsonString(var1);
	if var2 != `{"var":"test"}` {
		t.Log(var2);
		t.Fail();
	}
}

func TestToConstStr(t *testing.T) {
	var var1 = "test";
	var var2 = ToConstStr(var1);
	if var1 == fmt.Sprintf("%v", var2) {
		t.Fail();
	}
}

func TestStrToStrDictionary(t *testing.T) {
	var var1 = `{"var":2}`;
	var var2 = StrToStrDictionary(var1);
	if var2["var"] != "2" {
		t.Fail();
	}
}

func TestDictionaryToStrDictionary(t *testing.T) {
	var var1 = map[string]interface{}{"var":2,};
	var var2 = DictionaryToStrDictionary(var1);
	if var2["var"] != "2" {
		t.Fail();
	}
}

func TestStrToStrPtr(t *testing.T) {
	var var1 = "test";
	var var2 = StrToStrPtr(var1);
	if var1 == fmt.Sprintf("%v", var2) {
		t.Fail();
	}
}

func TestIntToIntPtr(t *testing.T) {
	var var1 int32;
	var1 = 2;
	var var2 = IntToIntPtr(var1);
	if fmt.Sprintf("%v", var1) == fmt.Sprintf("%v", var2) {
		t.Fail();
	}
}

func TestBoolToBoolPtr(t *testing.T) {
	var var1 = false;
	var var2 = BoolToBoolPtr(var1);
	if fmt.Sprintf("%v", var1) == fmt.Sprintf("%v", var2) {
		t.Fail();
	}
}

