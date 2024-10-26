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
