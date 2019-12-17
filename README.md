# build

```bash
$ go build ./...     # build recursively (NO OUTPUT - just verifies build works). -n => dry run print of build steps
$ go build -o path/to/output ./path/to/cmd/dir  # build a command/executable.  w/o -o writes to pwd
$ env GOOS=windows GOARCH=amd64 go build ./path/to/cmd/dir  # build a command/executable for a target OS/CPU (here Windows/AMD64)
$ go install ./...   # build and copy results to GOPATH - includes executables and packages
$ go clean -i ./...  # cleans out results of install (-i).  -n => dry run print of files affected
$ git tag -a v0.1.0 -m "tag comment"  # set module version
```