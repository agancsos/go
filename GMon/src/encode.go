package main
import (
	"os"
	"./classes/incryptEncode"
)

func main() {
	var str = "";
	var encoder = &incryptEncode.IncryptEncodeFullBinary{};
	var isDecoded = false;

	for i := 0; i < len(os.Args); i++ {
		switch (os.Args[i]) {
			case "--str":
				str = os.Args[i + 1];
				break;
			case "--decode":
				isDecoded = true;
				break;
		}
	}
	if !isDecoded {
		println(encoder.GetEncoded(str));
	} else {
		println(encoder.GetDecoded(str));
	}
	os.Exit(0);
}

