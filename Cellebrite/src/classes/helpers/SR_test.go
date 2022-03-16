package helpers
import "testing"

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

func TestArgsToDictionary(t *testing.T) {
	var var1 = []string{"-p", "test"};
	var var2 = ArgsToDictionary(var1);
	if var2["-p"] == "" {
		t.Fail();
	}
}

