# Go Programming Learning

## Note

* go is a tool for managing Go source code. 
* `go run` - It compiles and runs the application. It doesn't produce an executable. 
* `go run file.go` - compiles and immediately runs Go programs.
* `go build` - It just compiles the application. It produces an executable. 
* `go build file.go` - compiles a bunch of Go source files. It compiles packages and dependencies. If I run `go build` it will compile the files in the current directory and will produce an executable file with the name of the current working directory. 
* `go build -o app.exe` - will produce an executable file called `app`. (Run `go mod init` before running `go build` and if I am not `GOPATH/src` folder then it will not work.)

### go install

* Both `go install` and `go build` will compile the package in the current directory. 
* If the package is `main`, `go build` will place the resulting executable in the current directory and `go install` will remove the executable to `GOPATH/bin`. 
* When running `go install` I use paths relative to `GOPATH/src`.

### Compiling (go build) and Running Go Applications (go run)

* Compiling for Windows: `GOOS=windows GOARCH=amd64 go build -o winapp.exe`  
* Compiling for Linux: `GOOS=linux GOARCH=amd64 go build -o linuxapp`  
* Compiling for Mac: `GOOS=darwin GOARCH=amd64 go build -o macapp`

### Formatting Go Source Code (gofmt)

* Go strongly suggests certain styles. 
* `gofmt` which comes from `golang formatter` will format a program's source code in an idiomatic way that is easy to read and understand. 
* This tool, `gofmt` is build into the language runtime, and it formats Go code according to a set of stable, well-understood language rules. I can run it manually at the command line or I can configure my IDE, for example: VS Code, to run `gofmt` each time a file is saved. `gofmt -w main.go` or for folder `gofmt -w -l folder_name`, `-l` for showing the changed files.
* Instead of `gofmt -w main.go`, I can use `go fmt`.
