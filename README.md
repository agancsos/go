# Cellebrite

![Alt ""](https://edrm.net/wp-content/uploads/2021/10/cellebrite-logo.png "Cellebrite")

## Synopsis
A commulative tool based on projects built in preparation for Cellebrite interviews.  Meant to demonstrate current understand as well continuous learning of Go and low-level technologies.

## Requirements
* The tool will read a Master Boot Record from disk.
* The tool will generate an $MFT record from disk.
* The tool will read an $MFT dump.
* The tool will demonstrate low-level system calls.

## Assumptions
* The $MBR and $MFT portions of the tool will be ran as root or Administrator.
* The $MBR and $MFT portions of the tool may be ran against an NTFS or GUID disks. 
* The tool as a whole may or may not be ran on Windows, Linux, or macOS.
* The tool may or may not need a database.
* Go 1.17.6+ is installed on the system that is to build or run the tool.

## Implementation Details
The tool is made up of a single CLI, which is driven by a tool parameter to drive the following features.

### REST API
A REST service is by default generated at port 3434 (configurable with the --port parameter), which exposes basic endpoints to demonstrate REST API design.

### MFT
The "MFT" feature is the bulk of the tool as this appeared to be a strong component of the role applying to.  This feature uses a path parameter to either read the $MBR record from the disk or $MFT from a snapshot.

### BOMB
The "BOMB" feature is a simple goroutine that waits for 300 seconds or until a keyboard termination signal is sent.

### Unit Tests
Unit Tests were added to the tool to demonstrate TDD practices and ability to test features.

### Dockerfile
A Dockerfile was added to demonstrate knowledge of containers and also to ensure that there's a guarenteed way to test the tool.  The image is based off of a clean Debian 10 image and then configures basic software.  Upon running the container, the Unit Tests will be invoked through a Bash script, printing out the results of the tests.

### Note about ABI and binary compatibility
Instead of architecting a full database and a custom DataService to demonstrate ABI and binary compatibility, I will reference a personal project, [GMon](https://github.com/agancsos/go/tree/main/GMon), which uses a data package through the unixODBC API, SQLite API, and corresponding ABI artifacts.  Upon compiling, CGO will generate the proper ABI for the SQLite API that would be appropriaate for the platform to use the SQLite DataConnection.  Upon compiling, CGO will use an installed ODBC API and then link against the installed ODBC ABI artifacts.  The project also includes a gRpc protocol to demonstrate further knowledge of Go.

I ended up adding a smaller ABI project which simulates a personal API to encode and decode strings using different algorithms.  Note that due to security reasons, the API implementations have been redacted to the bare minimum in order to compile and run.

#### Build
```bash
docker build -t cellebrite_prep .
```

#### Run
```bash
docker run --rm --name test1 -t cellebrite_prep
```

#### Flags
|Flag                           | Description                                                                                              |
|--|--|
|-t                             | Tool to run.                                                                                             |
|-o                             | Operation of the tool to run.                                                                            |
|-p                             | Full path of resource for tool operation.                                                                |
|--port                         | Port for the REST API.  Default is 3434.                                                                 |

#### Tools
|Tool                           | Description                                                                                              |
|--|--|
|mft                            | Tool that works with $MBR and $MFT records.                                                              |
|bomb                           | Tool to demonstrate low-level system calls.                                                              |


## Retrospective
* Most of the retrospective has to do more with the process and the people met.
    * All interviewers/screeners seemed to be down to earth and cool to work with.
    * I feel I had a good connection with those I met and would have fit into the company culture.
* In terms of the project, as already mentioned, it's a combination of other projects built in preparation for any technical interviews.
    * The "BOMB" tool is based off a project built to look into system calls using Go, while C++ implementation had already been done before.
    * The MFT tool is based off a project built to look into low-level file systems after screening with an architect.  I enjoyed the discussion as well as the research.
* Although C++ is my "reasoning" language, Go is a beautiful programming language and I plan to continue my education of it.
* Although the $MBR and $MFT portion of the tool are very much incomplete (and to be honest, pretty crappy), this should demonstrate a desire to learn and a lack of fear of jumping into new topics.
    * There are already packages and tools to extract the $MBR/$MFT records, but it was an interesting topic to dive into in order to get a better understanding of the underlying concepts for said packages/tools.

## Closing thoughts
Thank you again for all the time and effort put into looking over my application as well as the great technical questions during technical screenings.  Unfortunately, I started the process too late and another offer came in, but I will definetely keep Cellebrite on the top of my list as well as continue to develop my skills to potentially give you my best (if you'd still have me).  Thank you so much for your time.

## References
* https://youtu.be/xW5UwDztkX4 
* https://en.wikipedia.org/wiki/Master_boot_record 
* https://writeblocked.org/resources/NTFS_CHEAT_SHEETS.pdf 
* https://andreafortuna.org/2017/07/18/how-to-extract-data-and-timeline-from-master-file-table-on-ntfs-filesystem/ 

