package mft
import (
	"os"
	"errors"
	"os/user"
)

func ParseMFT(path string, bufferSize int) (MFTItem, error) {
	var user, _ = user.Current();
	if user.Username != "root" {
		return nil, errors.New("User must be root/Administrator");
	}
	var d, err = os.Open(path);
	if err != nil {
		return nil, err;
	}
	defer d.Close();
	var entry MFTItem;
	var fileInfo, _ = os.Stat(path);
	if fileInfo.Size() <= 512 {
		entry, err = ReadDevice(d);
	} else {
		entry, err = ReadDump(d);
	}
	if err != nil {
		return nil, err;
	}
	return entry, nil;
}

