package common
import (
	"testing"
)

func TestEmptyPropertyStore(t *testing.T) {
	var var1 = &EmptyPropertyStore{};
	var1.GetKeys();
}

func TestJsonPropertyStore(t *testing.T) {
	var var1 = &JsonPropertyStore{};
	var1.GetKeys();
}

func TestPlainPropertyStore(t *testing.T) {
	var var1 = &PlainPropertyStore{};
	var1.GetKeys();
}

