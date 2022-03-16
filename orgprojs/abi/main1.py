#!/usr/bin/env python3
###############################################################################
# Name        : clib1.py                                                      #
# Version     : v. 1.0.0.0                                                    #
# Author      : Abel Gancsos                                                  #
# Description : Researching ctypes package.                                   #
###############################################################################
import os, sys, ctypes;

class CIncryptEncode:
	lib=None;lib_path=None;encoder=None;
	def __init__(self, path="./resources/libencoder.dylib"):
		assert os.path.exists(path), "Library path must exist...";
		self.lib_path = path; self.lib = ctypes.cdll.LoadLibrary(self.lib_path);
		self.lib.getEncoded.restype  = ctypes.c_char_p;
		self.lib.getEncoded.argtypes = [ctypes.c_char_p, ctypes.c_int];
		self.lib.getDecoded.restype  = ctypes.c_char_p;
		self.lib.getDecoded.argtypes = [ctypes.c_char_p, ctypes.c_int];
	def get_encoded(self, string, alg=3, force=False):
		return self.lib.getEncoded(string.encode("utf-8"), alg).decode("utf-8");
	def get_decoded(self, string, alg=3, force=False):
		assert alg > 1, "Invalid algorithm for method ({0})".format(alg);
		return self.lib.getDecoded(string.encode("utf-8"), alg).decode("utf-8");
	pass;

if __name__ == "__main__":
	params = dict();
	for i in range(0, len(sys.argv) - 1): params[sys.argv[i]] = sys.argv[i + 1];
	incrypt_encode = CIncryptEncode(params["-p"]) if "-p" in params.keys() else CIncryptEncode();
	assert "--str" in params.keys(), "String to encode/decode not found...";
	operation = params["-o"] if "-o" in params.keys() else "encode";
	alg       = int(params["--alg"]) if "--alg" in params.keys() else 3;
	if operation == "encode": print(incrypt_encode.get_encoded(params["--str"], alg));
	elif operation == "decode": print(incrypt_encode.get_decoded(params["--str"], alg));
	else: raise Exception("Invalid operation ({0})".format(operation));

