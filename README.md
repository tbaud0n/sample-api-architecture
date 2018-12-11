# Sample architecture for Golang API

The following api architecture is based on the [standard package layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) proposed by Ben Johnson.

*The following api must*

## Explanation

The `api.go` file at the root of this project defines the structs and the services shared by the application subpackages.

Each subpackage is independant of the others. It consumes and produce only mocks of the interfaced services defined in the `api.go` file.

The `main` package loads the configuration file and launch the http server using the required mocks.

Using this kind of architecture, each subpackage can be easily replaced by another implementation without having to modify the whole application.

```
.
├── api.go // Definition of the structs and services shared by the subpackages
├── cmd
│   └── main.go // Our main package to configure and run the api
├── config.yaml
├── pkg
│   ├── config
│   │   └── config.go // A subpackage to load the config from a yaml file
│   ├── http // The HTTP handlers
│   │   ├── http.go
│   │   └── user.go
│   ├── logger
│   │   └── logger.go // A simple logger
│   └── storage
│       └── postgresql // The postgresql mocks implementation
│           ├── postgresql.go
│           └── user.go
```
