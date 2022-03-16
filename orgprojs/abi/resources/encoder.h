#ifndef __ENCODER_API_H_INCLUDED__
#define __ENCODER_API_H_INCLUDED__
#include <stdio.h>
#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif
	char *getEncoded(char *str, int algorithm);
	char *getDecoded(char *str, int algorithm);
#ifdef __cplusplus
}
#endif

#endif
