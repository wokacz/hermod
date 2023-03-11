# hermod
## REST Microservices in Go

### Project Layout

* **cmd** - This will contain the entry point (main.go) files for all the services and also any other container images if any
* **docs** - This will contain the documentation for the project
* **config** - All the sample files or any specific configuration files should be stored here
* **internal** - This package is the conventional internal package identified by the Go compiler. It contains all the packages which need to be private and imported by its child directories and immediate parent directory. All the packages from this directory are common across the project
* **pkg** - This directory will have the complete executing code of all the services in separate packages.
* **tests** - It will have all the integration and E2E tests
* **vendor** - This directory stores all the third-party dependencies locally so that the version doesn't mismatch later