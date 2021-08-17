#!/usr/bin/env python3
import os, sys, glob, shutil;

class Compiler:
	base_path=None;clean=None;component=None;no_tar=None;skip_depends=None;
	def __init__(self, params=dict()):
		self.base_path = params["-b"] if "-b" in params.keys() else "{0}/../".format(os.path.realpath(os.path.dirname(__file__)).replace("\\", "/"));
		self.clean = True if "--clean" in params.keys() and int(params["--clean"]) > 0 else False;
		self.no_tar = True if "--no-tar" in params.keys() and int(params["--no-tar"]) > 0 else False;
		self.component = "-c" if "-c" in params.keys() else "*";
		self.skip_depends = True if "--skip-depends" in params.keys() and int(params["--skip-depends"]) > 0 else False;
		pass;
	
	## Purges old builds to ensure binaries are up to date
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
		os.system("{0}dist/rpc1service &".format(self.base_path));				   ## Start the server as a job
		
		os.system("pkill -f rpc1server")										   ## Kill the server so that it does not become a runaway process
		pass;

	## Installs the following dependencies if they don't exists:
	## 1. google.golang.org/grpc
	## 2. google.golang.org/protobuf
	## 3. google.golang.org/genproto
	## Natively, it will also install:
	## 1. github.com/golang/protobuf
	## 2. golang.org/x/net
	## 3. golang.org/x/text
	def __install_golang_dependencies(self):
		packages = {"google.golang.org/grpc", "google.golang.org/protobuf", "google.golang.org/genproto"};
		assert os.path.exists("{0}".format(os.environ.get("HOME"))), "Missing Go directory...";
		if not os.path.exists("{0}/src".format(os.environ.get("HOME"))): os.mkdir("{0}/src".format(os.environ.get("HOME")));
		for package in packages:
			if not os.path.exists("{0}/src/{1}".format(os.environ.get("HOME"), package)):
				## Check if the package exists in the cache already
				if len(glob.glob("{0}/pkg/mod/{1}@*".format(os.environ.get("HOME"), package))) > 0:
					cache = glob.glob("{0}/pkg/mod/{1}@*".format(os.environ.get("HOME"), package));
					for item in cache: shutil.copytree(item, item.split("@")[0].replace("pkg/mod", "src"));
					pass;
				else:
					## Install using go install
					os.system("go install {0}@latest".format(package));
					
					## Copy from mod to src
					cache = None;
					if len(glob.glob("{0}/pkg/mod/{1}@*".format(os.environ.get("HOME"), package))) > 0:
						cache = glob.glob("{0}/pkg/mod/{1}@*".format(os.environ.get("HOME"), package));
						for item in cache: shutil.copytree(item, item.split("@")[0].replace("pkg/mod", "src"));
					pass;
				pass;
		pass;

	def invoke(self):
		assert self.base_path != "", "Base path cannot be empty...";
		assert self.component != "", "Component cannot be empty...";
		if not os.path.exists("{0}dist".format(self.base_path)): os.mkdir("{0}dist".format(self.base_path));
		self.__clean("{0}dist".format(self.base_path));
		if not self.clean:
			print("Installing Go dependencies...");
			if not self.skip_depends: self.__install_golang_dependencies();
			if (self.component == "*"):
				components = glob.glob("./main_*.go");
				for c in components:
					comps = c.split("_");
					component_name = comps[1].replace(".go", "");
					os.system("cd {0}src && go build -o {0}dist/rpc1{1} main_{1}.go".format(self.base_path, component_name));
				pass;
			else:
				os.system("cd {0}src && go build -o {0}dist/rpc1{1} main_{1}.go".format(self.base_path, self.component));
				pass;
			self.__run_unit_tests();
			pass;
		print("Completed!");
	pass;

if __name__ == "__main__":
	param = dict();
	for x in range(1, len(sys.argv[1:]), 2) : param[sys.argv[x]] = sys.argv[x + 1];
	session = Compiler(param);
	session.invoke();
	pass;

