## All the code and notes are derivatives of the book let's go by Alex Edwards : https://lets-go.alexedwards.net/
-------------------------------------------------------------------------

To run the go program on a different port  (4000 is the default)

```sh
go run ./cmd/web -addr=":9999"
```

To get automated help that shows the available command line flags use the -help flag

```sh
go run ./cmd/web -help
```
Getting the cmd flags from the environment

```sh
export SNIPPETBOX_ADDR=":9999"
go run ./cmd/web -addr=$SNIPPETBOX_ADDR

```

If you needed to get it directly in code (os has to be imported):

```go
addr := os.Getenv("SNIPPETBOX_ADDR")
```

To redirect info and error logs. Note: Using the double arrow >> will append to an existing file, instead of truncating it when starting the application.

```sh
go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
```
To install dependencies

```sh
go mod download
```
To upgrade to latest available minor or patch release of a package, you can simply run go get with the -u flag like so:
```sh
go get -u github.com/foo/bar
```

Or alternatively, if you want to upgrade to a specific version then you should run the same command but with the appropriate @version suffix. For example:

```sh
go get -u github.com/foo/bar@v2.0.0
```

Sometimes you might go get a package only to realize later that you don’t need it anymore. When this happens you’ve got two choices.

You could either run go get and postfix the package path with @none, like so:

```sh
go get github.com/foo/bar@none

```
Or if you’ve removed all references to the package in your code, you could run go mod tidy, which will automatically remove any unused packages from your go.mod and go.sum files.

```sh
go mod tidy -v
```