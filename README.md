# A set of example solutions for icanhazwordz

There are 3 executable programs in here:

* simplified.go - a self-contained single program example solution. You can run this with `go run simplified.go` directly.
* anysolver/solve.go - a generic interactive solver that runs one of the solvers from the solver package.
* benchmark/bench.go - a benchmarking program to compare all of the solvers.

## How to run these

solve.go and bench.go require that you have the `step2018/idohazwordz` packages
available in your [$GOPATH](https://golang.org/doc/code.html). The easiest way
to do this is to run

```sh
go get github.com/step2018/idohazwordz
```

This will clone this repository into `$GOPATH/src/github.com/step2018/idohazwordz`.
You can then just cd to this directory, and run these programs directly:

```sh
cd $GOPATH/src/github.com/step2018/idohazwordz
go run anysolver/solve.go
```

or just run them directly:

```sh
go run $GOPATH/src/github.com/step2018/idohazwordz/benchmark/bench.go
```
