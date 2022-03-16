#!/usr/bin/ruby
###############################################################################
# Name        : run.rb                                                        #
# Author      : Abel Gancso                                                   #
# Version     : v. 1.0.0.0                                                    #
# Description : Bootstrapper for running the test code                        #
###############################################################################

class Bootstrapper
	@mode=nil;@working_directory=nil;@go_path=nil;@docker_path=nil;@no_shutdown=nil;
	@container_name=nil;@container_tag;@remote_base_path=nil;@where_command=nil;@grep_command=nil;
	def initialize(m, shutdown=false)
		@mode = m;
		@working_directory = File.expand_path(File.dirname(__FILE__));
		@where_command = RUBY_PLATFORM.include?("mingw") ? "where" : "which";
		@grep_command = RUBY_PLATFORM.include?("mingw") ? "findstr" : "grep";
		@where_command = "" if RUBY_PLATFORM.include?("mingw");
		@remote_base_path = "#{ENV["HOME"]}/stuff/go/rpc1"
		@go_path = `#{@where_command} go`.strip();

		## Since we try to print the Docker executable path and Windows might have multiple entries, this needs to be handle differently.
		@docker_path = RUBY_PLATFORM.include?("mingw") ? `#{@where_command} docker.exe`.strip() : `#{@where_command} docker`.strip();
		@no_shutdown = shutdown;
		@container_name = "rpcapi1";
		@container_tag = "rpcapi1:latest";
		puts("".ljust(80, "#"));
		puts("# Working Directory : #{@working_directory}".ljust(79, " ") + "#");
		puts("# GOPATH            : #{@go_path}".ljust(79, " ") + "#");
		puts("# Docker Path       : #{@docker_path}".ljust(79, " ") + "#");
		puts("# Shutdown          : #{!@no_shutdown}".ljust(79, " ") + "#");
		puts("# Mode              : #{@mode}".ljust(79, " ") + "#");
		puts("".ljust(80, "#"));
	end

	## Prepare the docker container by taking the base image, installing dependencies, and copying the source files
	def prepare_container()
		throw "Docker is not installed...." if @docker_path == nil or @docker_path == "";
		`cp -fr #{@working_directory}/src #{@working_directory}/container`;
		throw "Failed to copy latest source to container staging directory...." if !$?.success?
		## Check if container already exists
		check = `docker ps | #{@grep_command} #{@container_name}`;
		if check != nil and check != ""
			## Ensure that the latest souce is on the container
			`docker stop #{@container_name}`;
			throw "Failed to stop container..." if !$?.success?;
			`docker cp #{@working_directory}/container/src #{@container_name}:#{@remote_base_path}`;
			throw "Failed to copy latest source..." if !$?.success?;
			`docker start #{@container_name}`;
			throw "Failed to start container after updating sources..." if !$?.success?;
		else
			## Build container
			`cd #{@working_directory}/container; docker build -t #{@container_tag} .`
			throw "Failed to build container..." if !$?.success?
			## Start container
			`docker run -d -p 8888:80 -p 2222:22 -p 36066:3606 -p 4441:4443 --name #{@container_name} -t #{@container_tag}`;
			throw "Failed to start container..." if !$?.success?
		end
	end

	## Check the status of the run; if it failed, keep the container for debugging, else stop and remove the container
	def ensure_no_errors(result)
		if !result
			puts("Errors were found in the build or Unit Tests.  Not removing container...");
			@no_shutdown = true;
		else
			if not @no_shutdown
				puts("No errors were found, removing container...");
				`docker stop #{@container_name}`;
        		throw "Failed to stop container..." if !$?.success?;
				`docker rm #{@container_name}`;
				throw "Failed to remove container..." if !$?.success?;
			end
		end
	end

	## Invoke the QA run, where it will build the source in a container then run the Unit Tests
	def invoke_qa()
		@remote_base_path = "/root/stuff/go/rpc1";
		self.prepare_container();
		result = `docker exec #{@container_name} python3 #{@remote_base_path}/src/compile.py`;
		puts(result);
		self.ensure_no_errors($?.success?);
	end

	## Invoke the standard build run, where it will compile the source locally and then run the Unit Tests
	def invoke_build()
		throw "Python is not installed..." if `#{@where_command} python3` == "";
		`python3 #{@working_directory}/src/compile.py`;
	end

	## Main handler
	def invoke()
		if @mode == "--QA"
			self.invoke_qa();
		elsif @mode == "--build"
			self.invoke_build();
		elsif @mode == "--gen-proto"
			if `#{@where_command} protoc` == ""
				throw "Protoc is not installed...";
			end
			`protoc --proto_path=#{@working_directory}/src/rpcapi --go_out=#{@working_directory}/src --go_opt=paths=source_relative \
				--go-grpc_out=#{@working_directory}/src --go-grpc_opt=paths=source_relative #{@working_directory}/src/rpcapi/rpcapi.proto`
		else
			throw "Mode not supported at this time...";
		end
		throw "Failed to remove container..." if !@no_shutdown and `docker ps | #{@grep_command} #{@container_name}` != "";
	end
end

## Main
mode = ARGV[0] != nil ? ARGV[0] : "--QA";
no_shutdown = false;
for i in 1..ARGV.length + 1
	no_shutdown = true if ARGV[i] == "--no-shutdown";
end
session = Bootstrapper::new(mode, no_shutdown);
session.invoke();

