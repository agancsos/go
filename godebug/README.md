# GoDebug

# Synopsis

# Assumptions
* There is a set of scripts that should be kept in a central location.
* There is a need for these scripts.
* These scripts will be used on Windows systems.
* These scripts might also be used on non-Windows systems.
* The systems will have the .Net Framework installed prior to using the toolkit.
* The scripts can be ran on any system (no dependencies).

# Implementation Details
At a very high level, this program reads in the name of the command that would normally be ran then creates a concrete instance of the AMGTool class using a Factory Method design pattern.  This AMGTool abstract class is simply a Task or Command class that has an Invoke method to perform the actual operations and a GetName as a florish.  Each operation, which was previously in a separate script form, has been simplified and refactored in this toolkit for scalability and for mantainability.

## Supported Operating Systems
* Windows 10
* Windows 7
* Ubuntu 16.0
* OpenSUSE 15.0+
* SLES 12.3+

# Prerequsites 

# Execution
<path-to-binary>/godebug -d <debug-name> [arguments]

# Default error message
Debugger not found...

