package mft
import (
	"../system"
	"io"
	"fmt"
	"os"
)

func ReadDevice(device *os.File) (*MFTDevice, error) {
	var result = &MFTDevice{};
	var buffer = make([]byte, 512);
	var _, err = io.ReadFull(device, buffer);
	if err != nil {
		return nil, err;
	}
	var reader                   = system.NewLittleEndianReader(buffer);
	result.content			     = system.CleanASCII(string(buffer));
	result.Platform			     = string(reader.Data(0x03, 4));
	result.Version			     = string(reader.Data(0x03 + 4, 4));
	result.Descriptor			 = fmt.Sprintf("%v", buffer[0x15]);
	result.TrackSectors		     = int(reader.UInt16(0x18));
	result.FileRecordSegmentSize = int(buffer[0x40]);
	result.TotalSectorCount	     = int(reader.UInt16(0x28));
	result.ClusterSectors		 = int(int8(buffer[0x0D]));
	return result, nil;
}

