# Abel Gancsos DigitalOcean Project

![Alt ""](https://logonoid.com/images/digitalocean-logo.png "DigitalOcean")

## Synopsis
Abel Gancsos DigitalOcean Project, agdo for short, also known as the binary, is a package management system written in Go, but isn't specific to the Go programming language.  The system is made up of two components, the REST service and the CLI.

### API
The API component is used to interact with the package registry.  Supported endpoints include:
* /api/get/      : List packages                         
* /api/update/   : Update local cache               
* /api/upgrade/  : Upgrade packages                
* /api/upload/   : Upload a package to the registry
* /api/version/  : Print the version of the REST API

The endpoints will then invoke the appropriate handler, which include: Get, List, and Upload.

### CLI
The CLI component is used as the client that uses the REST service to interact with the package registry.  Supported flags include:
* -p|--package  : Package name     
* -r|--registry : URL to the package registry 
* -h|--help     : Print the help menu 
* -v|--version  : Print the version of the CLI
* --op          : Operation to perform
* --repo-name   : Name for the new repository to create
* --repo-url    : Url for the new repository to create

CLI operations map directly to the REST endpoint.  The "purge" operation can also be used to remove a package.

## Assumptions
* The binary can be built and ran either on Linux or Windows.
* The binary can be built and ran on a container or local system.
* The provided Docker image will not be modified.

## Requirements
* The binary will be able to list packages.
* The bianry will be able to store a package cache.
* The binary will be able to upload the package cache.
* The binary will be able to update the package cache.
* The binary will be able to upload to the package registry.
* The binary will be able to use a different package registry.

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
The source code is structured in a way where at the root, there're the main modules for the components and a folder called classes that contains additional modules used to build out the components.  Under the classes folder, there's a common package I use in all of my projects for operations that I might use in the project.  There's also an sr package used for the project's static resources (globals, constants, command-line helpers, etc).  Then there's a helpers package, which in this case contains the modules used specifically for the project.  Typically, I would organize these under a specific folder with the project name with child folders, models and services; however, I felt it wasn't needed in this case, so I just bundled them together in a helpers package.

So, how do the modules interact with each other?  When the binaries run, they run from the appropriate main module, which is used to setup the environment based on the command-line input using common functions in the common package. The main modules are also where we intiantiate the instance of the RegistryCache as the project uses the Singlton programming pattern. Then when it's time to run the payload (invoke the operation), it goes to the helpers.Registry module and uses the RegistryCache struct to invoke the operation.  When we enter into the RegistryCache, it should already have the repositories loaded, which is done when the instance is instantiated.  For purposes of the project, a Repository is considered a directory where all the packages will be held and installed.  The repository contains a JSON manifest that holds basic information on the repository, the critical piece being the URL to that repository's registry, as well as a cached map of the packages contained in the repository.  

Now, when we go to invoke the operation, depending on the operation, it will either search the local cache (install, get, purge) or it will try to reach out to the REST API (update, upgrade, upload).  Despite the performance benefits of gRPC, REST was used instead due to time constraints and REST still being an acceptable industry standard.  We then perform the steps needed to fulfil the operation.  I tried to handle as many of the errors as possible by printing out a message of the issue and skipping the rest of the operation.  So if the binary does end up crashing, this would be considered an unhandled exception.  In some cases, an error might not be considered at all, in these cases the action isn't mandatory or it was done as a check.

#### Get
On the get operation, we simply iterate through the local cache and siplay the package name and version.

#### Install
On the install operation, we try to lookup the package name, reach out to the REST api to download the package, and then download it to the repository directory.  Additional steps would normally be taken.

### Update
On the update operation, we try to reach out to the REST APi to grab the list of packages and update the repository manifest.

#### Upgrade
On the upgrade operation, we try to reach out to the REST API to grab the list of packages, then using the local cache, we check if the repository manifest would need to be updated with the new package information.  We then also run the same install steps to update the contents.

#### Purge
On the purge operation, we try to lookup the package name in the local cache and if found, we try to delete the extracted package along with the downloaded package file.

### Compile script
The project comes with a compile script, compile.py that when the compile.py ran, it searches for main_X.go files, where X is the component to compile, which then gets built to the distribution directory.  The compile script supports the following switches.
* --clean       : Don't build the code, just clean the distribution directory
* -b            : The base path to search and compile from; default is based off of the location of the compile script
* -c            : Name of the component to compile

### Example package
A sample package is provided for the Unit Tests as well as a description of what can be included in a package.  A package is constructed in the following folder structure.
- <base-path> (Folder; typically agdo)
    - <repo-name> (Folder)
        - repo.json (File)
		- <package-name>.agdo (File; downloaded)
        - <package-name> (Folder; installed)
            - package.json (File)

The repository manifest is structured with the following properties.
{
	"name":"",
	"url":""
}

The package manifest is structured with the following properties.
{
	"name":"",
	"version":"",
	"author":"",
	"dependencies":
	[
		{
			"name":"",
			"version":""
		}
	]
}

### Bootstrap
The bootstrap script is being provided to demonstrate knowledge of Ruby as well as to assist in running the source code.  The bootstrap script can be ran in two primary modes and depending on which mode is detected, a container will be attempted or the build will be created locally.  When the build is ran through the bootstrap script, the working directory will be driven by the "$HOME/stuff/go/agdo" value.  This is somewhat of an abstract path, in that $GOPATH is explicitly set to "$HOME/stuff/go" due to custom packages.  The "agdo" directory is then used simply as the workspace for the project.

#### QA (--QA)
When in this mode, the script will attempt to run the build and Unit Tests via a Docker container.  It first checks if Docker is installed, if not, the script will raise an exception.  If Docker is found, it continues it's standard workflow described in more detail under the "Container" section.  

#### Build (--build)
When in this mode the bootstrapper will simply generate the binaries locally in the dist directory, ready for publication.

### Container
A Docker container is being provided for use within a QA test bed and to demonstrate knowledge of configuring a container.  The Dockerfile takes the base image, in this case a Ubuntu 18.04, installs the dependencies, creates required directories ($HOME/stuff/go/agdo), then copies the initial source files.  The bootstrap script then checks if the container is still up, the following steps are taken:
1. Performs needed steps to copy the latest source files (stops the container, copies the files, and starts the container).
2. Compiles the code and runs the Unit Tests via the compile.py script

If the container is not up, the following steps are taken:
1. Copies the source files to a staging directory
2. Builds a new image to ensure latest configurations
3. Uploads the latest source files (purely as a safety check)
4. Compiles the code and runs the Unit Tests via the compile.py script

If there were any errors during the build or Unit Tests, it keeps the container live for debugging, otherwise, the container is removed.  If the user prefers to keep the container in the current state, an ovverride command-line argument, "--no-shutdown" can be provided.  If the container failed to get removed despite not passing in the override, the script will throw a final expection as it wasn't expected.  If the base image of the container is replaced with another platform, the Dockerfile along with the bootstrap script may need to be updated to match the proper package management system and build process.

### Cookbook
A Chef Cookbook is being provided to demonstrate knowledge of Chef software as well as to show the difference in needed effort to run the full test when compared to the bootstrap script.  The recipe attempts to take most if not all platforms into consideration; however, the intended platform is anything Unix based **with the build dependencies installed**. Please note that the cookbook does not create a container.   The following steps will be taken when started by the chef-client CLI.
1. Detects the base path, similarly to the bootstrap script, via the node object.  This is needed in order to know where to copy the source files and run the build steps.
2. Copies the source files from the cookbook to the target node.
3. Compiles the source and runs the Unit Tests via the compile.py script.
4. Removes the project directory.

## Thank You
I would like to thank you again for your time to discuss my background and experience.  I look forward to hearing any feedback on the assignment and discussing next steps.

## Disclaimer
As I didn't get past the Hiring Manager Screen (no fault, just not the right role), I didn't get a chance to actually provide my solution.  I am posting the code purely for education purposes and due to the time spent on preparing this project.

## Retrospective
* The main thing I would change is the implementation of the API.  For scaling purposes, I would use gRPC instead of REST to take advantage of the performance improvements in HTTP/2.
* I would also add a SQLite database file to hold the repository manifests, so that instead of walking the file system, I can retrieve everything at once with a lightwight query.
* Design and implementation time was about 7 days, which I'm not really thrilled with as if I would have continued the rest of the process, this would have shown insufficient experience.
* I would have probably scoped the project better as I included other items (Chief recipe, bootstrap, etc) that may not have mattered.  I included these to show my range of skills.  This would have most likely helped with implementation time.
* If the source was put out in a PR, I'm sure feedback would include adding a generic error handler, which I agree with.
* The install operation supports only 1 package at a time, I would change the implementation to support multiple.

## References
The following referenced were used to help build some of the aspects of this tool.  Although slightly rewritten to match the style, they are still close enough where I can't take all the credit.
* https://gist.github.com/mattetti/5914158/f4d1393d83ebedc682a3c8e7bdc6b49670083b84
* https://tutorialedge.net/golang/go-file-upload-tutorial/
* https://github.com/agancsos/go/tree/main/testsuite/common
