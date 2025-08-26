
# 13. Packages and Modules in Go

This section covers how to organize code into packages and use modules to manage dependencies.

## Packages

A package is a way to group related Go source files together. Every Go program is made up of one or more packages.

The `main` package is the entry point of a program.

## Modules

A module is a collection of related Go packages that are versioned together as a single unit. Modules are the standard way to manage dependencies in Go.

A module is identified by a module path, which is declared in a `go.mod` file.

### Creating a Module

To create a module, you use the `go mod init` command.

```bash
go mod init example.com/greetings
```

This creates a `go.mod` file that contains the module path.

### `go.mod` file

The `go.mod` file defines the module's path and its dependencies.

```
module example.com/greetings

go 1.18
```

## Importing Packages

You can import a package from another module by using its module path.

```go
import "example.com/greetings"
```

## Local Modules

When developing code locally, you can use the `replace` directive in the `go.mod` file to use a local copy of a module instead of a remote one.

```
replace example.com/greetings => ./greetings
```

This tells Go to use the `greetings` module from the local `greetings` directory.

## Running the Code

To run the code in `13_Packages_and_Modules.go`, you first need to tidy the module dependencies.

```bash
cd "13 Packages and Modules"
go mod tidy
```

Then you can run the code.

```bash
go run .
```
