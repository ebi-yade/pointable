# Pointable
The code-generator used to get the pointer to scalar literals in Go

## About

Go does not currently provide the way to directly get the pointer to literal values as though `&"foo"` or `&3.14`.
You can do it by defining trivial functions like below,

```go
package ptr

func String(v string) *string {
	return &v
}
```

and calling `pkg.String("foo")` will be satisfying.

However, writing them everytime when you need is really tedious thing.
Although it is a good idea to create a convenient package, fewer dependencies are preferred generally.
So the best way is to generate the code via a command line tool: `pointable`.

## Usage

If you execute the `pointable` command without any flags, it will create a `./pkg` directory and Go files inside it.
You can specify the path and the name of package, but creating code inside the directory that has already existed is not supported.
This is because Pointable is oriented to keep its simplicity, and you easily do a workaround by creating temporally directory and merging them manually.

```shell
$ pointable -help
Usage: pointable [FLAGS]

Flags:
  -name string
    	the name of package (default: automatically specified via -path)
  -path string
    	the path to package (default "pkg")
  -version
    	show the version
```
