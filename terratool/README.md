# Terratool

## Synopsis
Lightweight solution to generate Terraform templates based on standard specifications and user input.  The solution doesn't actually generate the cloud resources, but only packages the different Terraform modules that will be used to create the resources.

## Assumptions
* There is a need for the solution.
* The solution will not create the actual cloud resources.
* The solution will create all neccessary modules.
* The solution may or may not need to generate templates for federated authentication used for local testing.
* The solution may or may not run on Windows, Linux, or macOS.
* The solution itself will not have a dependency on the Terraform CLI.
* We will generate the needed Terraform tasks with defined variables.
* The defined variables will be updated at deployment time.
* The solution may or may not need to support Amazon resources.
* The solution may or may not need to support Azure resources.

## Requirements
* The solution will be able to read in user specifications from the command-line.
* The solution will be able to read in user specifications from a JSON file.
* The solution will be able to create the neccessary folder structure.
* The solution will be able to generate the needed Terraform templates in an organized format.
* The solution will verify if all required properties have been passed.
* The solution will be able ot read a JSON formatted list of resources.

## Implementation Details
Implementation of the utility was done via a primary singlet service, TerraToolService, which upon invoking the CLI, will prepare the environment with the needed properties and then validate if any required field are missing.  Once the environment is bootstrapped and verified, we start to build out the package using a root package folder and add the different Terraform template files as defined by standardized specifications and Hashicorp best pracrices.

### Operations
| Operation                | Description                                                     |
| --                       | --                                                              |
| generate                 | Genreate the Terraform template package.                        |
| purge                    | Remove the Terraform template package (not the cloud resources. |

### Resources
```json
{
	"service": "ec2",
	"type": "instance",
	"count": 300
}
```

### Flags
| Flag                     | Description                                            |
| --                       | --                                                     |
| -h                       | Show help menu                                         |
| --version                | Show version                                           |
| --dry                    | No system changes                                      |
| --federated              | Generate federated authenticated templates             |
| --provider               | Provider used for the template generation              |
| --sysid                  | Application symbol used for tagging                    |
| -o                       | Full path for the base output                          |
| -f                       | Full path for input file                               |
| --op                     | Name of the operation to perform                       |

## References

