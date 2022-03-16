#!/usr/bin/env python3
import os, sys, glob;

class Compiler:
	base_path=None;clean=None;
	def __init__(self, params=dict()):
		self.base_path = params["-b"] if "-b" in params.keys() else "{0}/../".format(os.path.realpath(os.path.dirname(__file__)).replace("\\", "/"));
		self.clean = True if "--clean" in params.keys() and int(params["--clean"]) > 0 else False;
		pass;
	def __clean(self, path=""):
		files = os.listdir(path);
		for f in files:
			print("Purging: {0}".format(f));
			try:
				if os.path.isdir("{0}/{1}".format(path, f)):
					self.__clean("{0}/{1}".format(path, f));
					os.rmdir("{0}/{1}".format(path, f));
				else:
					os.remove("{0}/{1}".format(path, f));
			except Exception as ex:
				print("{0}".format(ex));
		pass;
	def __run_unit_tests(self):
		pass;
	def invoke(self):
		assert self.base_path != "", "Base path cannot be empty...";
		if os.path.exists("{0}dist".format(self.base_path)): self.__clean("{0}dist".format(self.base_path));
		if not self.clean:
			if not os.path.exists("{0}dist".format(self.base_path)): os.mkdir("{0}dist".format(self.base_path));
			os.system("go build -o {0}dist/agdo {0}src/main.go".format(self.base_path));
			self.__run_unit_tests();
		print("Completed!");
	pass;

if __name__ == "__main__":
	param = dict();
	for x in range(1, len(sys.argv[1:]), 2) : param[sys.argv[x]] = sys.argv[x + 1];
	session = Compiler(param);
	session.invoke();
	pass;

