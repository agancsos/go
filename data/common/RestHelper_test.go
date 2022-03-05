package common
import (
	"testing"
)


func TestInvokeGet(t *testing.T) {
	var var1 = InvokeGet("http://localhost:4445/version", map[string]string{});
	if var1 == nil {
		t.Fail();
	}
}

func TestInvokePost(t *testing.T) {
	var var1 = InvokePost("http://localhost:4445/version", map[string]string{}, map[string]string{});
	if var1 == nil {
		t.Fail();
	}
}

func TestEnsureRestMethod(t *testing.T) {
}
