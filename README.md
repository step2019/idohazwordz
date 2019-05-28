# A set of example solutions for icanhazwordz

There are 3 executable programs in here:

* `simplified.go` - a self-contained single program example solution. You can run this with `go run simplified.go` directly.
* `anysolver/solve.go` - a generic interactive solver that runs one of the solvers from the solver package.
* `benchmark/bench.go` - a benchmarking program to compare all of the solvers.

## How to run these

### Installing Go

If you are using the STEP virtual machine (仮想マシン) then you already have Go
installed and can skip ahead.

If you're using Mac OS X you'll have to download and install Go on your machine
first. You can [download](https://golang.org/dl/) it from https://golang.org .

The default installer on Mac OS X doesn't set up a GOPATH in your environment,
so it's using the default directory: `~/go` .

### Running individual programs.

solve.go and bench.go require that you have the `step2019/idohazwordz` packages
available in your [$GOPATH](https://golang.org/doc/code.html). The easiest way
to do this is to run

```sh
go get github.com/step2019/idohazwordz
```

This will clone this repository into
`$GOPATH/src/github.com/step2019/idohazwordz`
(`~/go/src/github.com/step2019/idohazwordz` by default).  You can then just cd
to this directory, and run these programs directly:

```sh
cd ${GOPATH:-~/go}/src/github.com/step2019/idohazwordz
go run anysolver/solve.go
```

or just run them directly:

```sh
go run ${GOPATH:-~/go}/src/github.com/step2019/idohazwordz/benchmark/bench.go
```

## Learning Go

A great way to start is the tutorial at https://tour.golang.org/ ([日本語版もある](https://go-tour-jp.appspot.com/)). The "basics" section covers enough to be
able to read
[simplified.go](https://github.com/step2019/idohazwordz/blob/public/simplified.go).
