# build

```bash
$ go build ./...     # build recursively (NO OUTPUT - just verifies build works). -n => dry run print of build steps
$ go build -o path/to/output ./path/to/cmd/dir  # build a command/executable.  w/o -o writes to pwd
$ env GOOS=windows GOARCH=amd64 go build ./path/to/cmd/dir  # build a command/executable for a target OS/CPU (here Windows/AMD64)
$ go install ./...   # build and copy results to GOPATH - includes executables and packages
$ go clean -i ./...  # cleans out results of install (-i).  -n => dry run print of files affected
$ git tag -a v0.1.0 -m "tag comment"  # set module version
```

# usage

```bash
$ cat > mycmd.go << EOF
package main

// // The following import fails because myintpkg visible only by module members
// import "github.com/heck/gomoduleexample/internal/myintpkg"
import "github.com/heck/gomoduleexample/myextpkg"

func main() {
    // // myintpkg isn't visible outside of module, so this isn't possible:
    // myintpkg.Run()
    myextpkg.Run()
}
EOF
$ go build ./mycmd.go
$ ./mycmd    # prints "Hello, external example!"
```

# how to recreate this example

```bash
$ mkdir path/to/gomoduleexample
$ cd path/to/gomoduleexample
$ git init
$ go mod init github.com/heck/gomoduleexample
$ cat > gomoduleexample.go << EOF
package gomoduleexample
EOF
$ mkdir cmd internal
$ mkdir cmd/mycmd internal/myintpkg myextpkg
# create the source files
$ cat > myintpkg/myintpkg.go << EOF
package myintpkg

import "fmt"

// Run func
func Run() {
    fmt.Printf("Hello, internal example!\n")
}
EOF
$ cat > myextpkg/myextpkg.go << EOF
package myextpkg

import "fmt"

// Run func
func Run() {
    fmt.Printf("Hello, external example!\n")
}
EOF
$ cat > cmd/mycmd/main.go << EOF
package main

import "github.com/heck/gomoduleexample/internal/myintpkg"
import "github.com/heck/gomoduleexample/myextpkg"

// Main func
func main() {
    myintpkg.Run()
    myextpkg.Run()
}
EOF
```

# how to setup a Mac

```bash
$ brew install go
# one time setup of the Go workspace (used for all projects)
$ cat >> ~/.zshrc << EOF
# Set the Go workspace to $HOME/go
export GOPATH="$HOME/go"
EOF
$ mkdir ~/go
$ cd ~/go
$ mkdir src pkg bin  # create dirs used by go commands
```