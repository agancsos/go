package mft
import "testing"

func TestDeviceTypeName(t *testing.T) {
	var var1 = &MFTDevice{};
	if var1.TypeName() != "Device" {
		t.Fail();
	}
}

func TestDeviceContent(t *testing.T) {
	var var1 = &MFTDevice{};
	var1.content = "TEST";
	if var1.Content() != var1.content {
		t.Fail();
	}
}

func TestEntryTypeName(t *testing.T) {
	var var1 = &MFTEntry{};
	if var1.TypeName() != "Entry" {
		t.Fail();
	}
}

func TestEntryContent(t *testing.T) {
	var var1 = &MFTEntry{};
	var1.content = "TEST";
	if var1.Content() != var1.content {
		t.Fail();
	}
}

