package system
import "encoding/binary"

type binaryReader struct {
	content           []byte
	byteOrder         binary.ByteOrder
}

func NewLittleEndianReader(content []byte) *binaryReader {
	return &binaryReader{content:content, byteOrder:binary.LittleEndian};
}

func NewBigEndianReader(content []byte) *binaryReader {
	return &binaryReader{content:content, byteOrder:binary.BigEndian};
}

func (x binaryReader) Data(offset int, length int) []byte {
	return x.content[offset:offset + length];
}

func CleanASCII(str string) string {
	var result = "";
	for _, c := range str {
		var asciiCode = int(c);
		if c != '.' && c != '\n' && c != ' ' && !(asciiCode >= 48 && asciiCode <= 57) &&
			!(asciiCode >= 65 && asciiCode <= 90) &&
			!(asciiCode >= 97 && asciiCode <= 122)  {
			continue;
		}
		result += string(c);
	}
	return result;
}

func (x binaryReader) UInt16(offset int) uint16 {
	return x.byteOrder.Uint16(x.content[offset:offset + 2]);
}

func (x binaryReader) UInt32(offset int) uint32 {
	return x.byteOrder.Uint32(x.content[offset:offset + 4]);
}

func (x binaryReader) UInt64(offset int) uint64 {
	return x.byteOrder.Uint64(x.content[offset:offset + 8]);
}

