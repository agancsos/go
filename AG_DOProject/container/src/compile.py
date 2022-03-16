#!/usr/bin/env python3
import os, sys, glob, shutil;
from os.path import expanduser;
class Compiler:
	base_path=None;should_clean=None;component=None;
	def __init__(self, params=dict()):
		self.base_path = params["-b"] if "-b" in params.keys() else "{0}/../".format(os.path.realpath(os.path.dirname(__file__)).replace("\\", "/"));
		self.should_clean = True if "--clean" in params.keys() and int(params["--clean"]) > 0 else False;
		self.component = params["-c"] if "-c" in params.keys() else "*";
		pass;
	def clean(self, path=""):
		files = os.listdir(path);
		for f in files:
			print("Purging: {0}".format(f));
			try:
				if os.path.isdir("{0}/{1}".format(path, f)):
					self.clean("{0}/{1}".format(path, f));
					os.rmdir("{0}/{1}".format(path, f));
				else:
					os.remove("{0}/{1}".format(path, f));
			except Exception as ex:
				print("{0}".format(ex));
		pass;
	def run_unit_tests(self):
		os.system("{0}dist/agdoapi &".format(self.base_path));                         ## Start REST API

		## Create local cache
		os.system("mkdir -p  {0}/agdo/agdo/test1".format(expanduser("~")));
		json_obj = "{\"Name\":\"test1\",\"Author\":\"Abel Gancsos\",\"Version\":\"1.0.0.0\", \"Dependencies\":[]}";
		os.system("echo '{0}' >> {1}/agdo/agdo/test1/package.json".format(json_obj, expanduser("~")));
		json_obj = "{\"Name\":\"agdo\",\"Url\":\"http://localhost:4455\", \"Packages\":[]}";
		os.system("echo '{0}' >> {1}/agdo/agdo/repo.json".format(json_obj, expanduser("~")));
		os.system("export HOME={0}; {1}dist/agdocli --op get".format(expanduser("~"), self.base_path));                  ## Test get operation
		os.system("{0}dist/agdocli --op upload --package {1}/agdo/agdo/test1".format(self.base_path, expanduser("~"))); ## Test upload operation
		os.system("export HOME={0}; {1}dist/agdocli --op update".format(expanduser("~"), self.base_path));               ## Test update operation
		os.system("export HOME={0}; {1}dist/agdocli --op get | grep test1".format(expanduser("~"), self.base_path));     ## Test get operation after update
		os.system("export HOME={0}; {1}dist/agdocli --op install --package test1".format(expanduser("~"), self.base_path));
		os.system("pkill -f agdoapi");                                                 ## Kill REST API
		pass;
	def invoke(self):
		assert self.base_path != "", "Base path cannot be empty...";
		if os.path.exists("{0}dist".format(self.base_path)): self.clean("{0}dist".format(self.base_path));
		if not self.should_clean:
			if not os.path.exists("{0}dist".format(self.base_path)): os.mkdir("{0}dist".format(self.base_path));
			if self.component != "*" and self.component != "":
				os.system("cd {0}src && go build -o {0}dist/agdo{1} main_{1}.go".format(self.base_path, component));
			else:
				components = os.listdir("{0}src".format(self.base_path));
				for c in components:
					if c == "." or c == ".." or c == "classes" or c == "main.go": continue;
					if "main" in c:
						print("Building component: {0}".format(c.replace("main_", "").replace(".go", "")));
						os.system("cd {0}src && go build -o {0}dist/agdo{1} main_{1}.go".format(self.base_path, c.replace("main_", "").replace(".go", "")));
			shutil.copyfile("{0}src/index.htm".format(self.base_path), "{0}dist/index.htm".format(self.base_path));
			if not os.path.exists("{0}dist/agdo".format(self.base_path)): os.mkdir("{0}dist/agdo".format(self.base_path));
			shutil.copyfile("{0}test1.agdo".format(self.base_path), "{0}dist/agdo/test1.agdo".format(self.base_path));
			self.run_unit_tests();
		print("Completed!");
	pass;

if __name__ == "__main__":
	param = dict();
	for x in range(1, len(sys.argv[1:]), 2) : param[sys.argv[x]] = sys.argv[x + 1];
	session = Compiler(param);
	session.invoke();
	pass;

