package mft
import (
	"../system"
	"io"
	"os"
)

func ReadDump(dump *os.File) (*MFTEntry, error) {
	var result = &MFTEntry{};
	result.content = "";
	for {
		var buffer = make([]byte, 1024);
		var _, err = io.ReadFull(dump, buffer);
		if err != nil { break; }
		result.content += system.CleanASCII(string(buffer));
	}
	return result, nil;
}

