#include "encoder.h"

char *getHashEncode(char *str) {
	return "";
}

char *getBinaryEncode(char *str) {
	return "";
}

char *getBinaryDecode(char *str) {
	return "";
}

char *getFullBinaryEncode(char *str) {
	return "";
}

char *getFullBinaryDecode(char *str) {
	return "";
}

char *getEncoded(char *str, int algorithm) {
	char result[400000];
	switch (algorithm) {
		case 1:
			sprintf(result, "%s", getHashEncode(str));
			break;
		case 2:
			sprintf(result, "%s", getBinaryEncode(str));
			break;
		case 3:
			sprintf(result, "%s", getFullBinaryEncode(str));
			break;
		default: 
			sprintf(result, "Invalid algorithm %d", algorithm);
			break;
	}
	return result;
}

char *getDecoded(char *str, int algorithm) {
	char result[400000];
	switch (algorithm) {
		case 2:
			sprintf(result, "%s", getBinaryDecode(str));
			break;
		case 3:
			sprintf(result, "%s", getFullBinaryDecode(str));
			break;
		default: 
			sprintf(result, "Invalid algorithm %d", algorithm);
			break;
	}
	return result;
}

