package system
import (
	"testing"
	"fmt"
)

func CleanASCIITest(t *testing.T) {
	var var1 = ")()))t*e$%st)))";
	if CleanASCII(var1) != "test" {
		t.Fail();
	}
}

func TestLittleUInt16(t *testing.T) {
	var var1   = uint16(244);
	var reader = NewLittleEndianReader([]byte(fmt.Sprintf("%d", var1)));
	if reader.UInt16(0) != 13362 {
		t.Log(reader.UInt16(0));
		t.Fail();
	}
}

func TestLittleUInt32(t *testing.T) {
	var var1   = uint32(244);
	var reader = NewLittleEndianReader([]byte(fmt.Sprintf("%d", var1)));
	if reader.UInt32(0) != 3421234 {
		t.Log(reader.UInt32(0));
		t.Fail();
	}
}

func TestLittleUInt64(t *testing.T) {
	var var1   = uint64(244);
	var reader = NewLittleEndianReader([]byte(fmt.Sprintf("%d", var1)));
	if reader.UInt64(0) != 3421234 {
		t.Log(reader.UInt64(0));
		t.Fail();
	}
}

func TestBigUInt16(t *testing.T) {
	var var1   = uint16(244);
	var reader = NewBigEndianReader([]byte(fmt.Sprintf("%d", var1)));
	if reader.UInt16(0) != 12852 {
		t.Log(reader.UInt16(0));
		t.Fail();
	}
}

func TestBigUInt32(t *testing.T) {
	var var1   = uint32(244);
	var reader = NewBigEndianReader([]byte(fmt.Sprintf("%d", var1)));
	if reader.UInt32(0) != 842281984 {
		t.Log(reader.UInt32(0));
		t.Fail();
	}
}

func TestBigUInt64(t *testing.T) {
	var var1   = uint64(244);
	var reader = NewBigEndianReader([]byte(fmt.Sprintf("%d", var1)));
	if reader.UInt64(0) != 3617573575289995264 {
		t.Log(reader.UInt64(0));
		t.Fail();
	}
}

