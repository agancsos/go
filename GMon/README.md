# Gancsos Monitor

## Synopsis
Gancsos Monitor, also know as GMON or the binary, is a system monitoring project influenced by Microsoft's SCOM product.  

### Server
The server is simply used to server th gRPC API and does nothing else.  The server exposes the following methods to the client.

#### Methods
* version            : Gets the version of the API

### Client
The client is used to interact with the methods in the server.

#### Flags
* -h|--help          : Print the help menu

## Assumptions
* The binary can be built and ran either on Linux or Windows.
* The binary can be built and ran on a container or local system.
* The provided Docker image will not be modified.
* Some API methods will be useful while others might not be.
* Some API methods will require arguments while others will not.

## Requirements
* The binary will implement a gRPC server.
* The binary will expose usable gRPC methods.
* The binary will implement a sample gRPC client for the server.
* The binary will implement a client call to each server endpoint.

## Constraints
* The binary must compile with Go 1.16.5 and higher.
* The binary must run on macOS 10.14 and higher.
* The binary must run on Ubuntu 18.04 and higher.
* The binary must compile and run under Docker 20.10.7 and higher.
* The binary must be compatible with Chef Workstation version: 21.7.545 and higher.
* The binary must be compatible with Chef Infra Client version: 17.3.48 and higher.
* The binary must be compatible with Chef InSpec version: 4.38.9 and higher.
* The binary must be compatible with Chef CLI version: 5.3.1 and higher.
* The binary must be compatible with Chef Habitat version: 1.6.351 and higher.
* The binary must be compatible with Test Kitchen version: 3.0.0 and higher.
* The binary must be compatible with Cookstyle version: 7.15.2 and higher.

## Contents
### SRC
The source code is structured in a way where at the root, there're the main modules for the components and a folder called classes that contains additional modules used to build out the components.  Under the classes folder, there's a common package I use in all of my projects for operations that I might use in the project.  There's also an sr package used for the project's static resources (globals, constants, command-line helpers, etc).  Then there's a rpcapi package, which in this case contains the gRPC service implementation.  Then there's a section for the protocol, which is just the service, calls, and message defintions.


### Compile script
The project comes with a compile script, compile.py that when the compile.py ran, it searches for main_X.go files, where X is the component to compile, which then gets built to the distribution directory.  The compile script supports the following switches.
* --clean       : Don't build the code, just clean the distribution directory
* -b            : The base path to search and compile from; default is based off of the location of the compile script
* -c            : Name of the component to compile

### Bootstrap
The bootstrap script is being provided to demonstrate knowledge of Ruby as well as to assist in running the source code.  The bootstrap script can be ran in two primary modes and depending on which mode is detected, a container will be attempted or the build will be created locally.  When the build is ran through the bootstrap script, the working directory will be driven by the "$HOME/stuff/go/gmon" value.  This is somewhat of an abstract path, in that $GOPATH is explicitly set to "$HOME/stuff/go" due to custom packages.  The "gmon" directory is then used simply as the workspace for the project.

#### QA (--QA)
When in this mode, the script will attempt to run the build and Unit Tests via a Docker container.  It first checks if Docker is installed, if not, the script will raise an exception.  If Docker is found, it continues it's standard workflow described in more detail under the "Container" section.  

#### Build (--build)
When in this mode the bootstrapper will simply generate the binaries locally in the dist directory, ready for publication.

#### Generate protocol (--gen-proto)
When in this mode, the script will help generate a fresh copy of the stubs that can be used to update the rpcapi package modules.  Once the protocol is generated, all other steps will require manual intervention.  As this mode requires additional development steps, it is a local only operation.

### Container
A Docker container is being provided for use within a QA test bed and to demonstrate knowledge of configuring a container.  The Dockerfile takes the base image, in this case the latest OpenSUSE, installs the dependencies, creates required directories ($HOME/stuff/go/gmon), then copies the initial source files.  The bootstrap script then checks if the container is still up, the following steps are taken:
1. Performs needed steps to copy the latest source files (stops the container, copies the files, and starts the container).
2. Compiles the code and runs the Unit Tests via the compile.py script

If the container is not up, the following steps are taken:
1. Copies the source files to a staging directory
2. Builds a new image to ensure latest configurations
3. Uploads the latest source files (purely as a safety check)
4. Compiles the code and runs the Unit Tests via the compile.py script

If there were any errors during the build or Unit Tests, it keeps the container live for debugging, otherwise, the container is removed.  If the user prefers to keep the container in the current state, an ovverride command-line argument, "--no-shutdown" can be provided.  If the container failed to get removed despite not passing in the override, the script will throw a final expection as it wasn't expected.  If the base image of the container is replaced with another platform, the Dockerfile along with the bootstrap script may need to be updated to match the proper package management system and build process.

### Cookbook
A Chef Cookbook is being provided to demonstrate knowledge of Chef software as well as to show the difference in needed effort to run the full test when compared to the bootstrap script. The recipe attempts to take most if not all platforms into consideration; however, the intended platform is anything Unix based with the build dependencies installed. Please note that the cookbook does not create a container. The following steps will be taken when started by the chef-client CLI.

1. Detects the base path, similarly to the bootstrap script, via the node object. This is needed in order to know where to copy the source files and run the build steps.
2. Copies the source files from the cookbook to the target node.
3. Compiles the source and runs the Unit Tests via the compile.py script.
4. Removes the project directory.

## Retrospective

## References

