This is a sample repo for trying out go fuzz that was added with go1.17

See more info here https://go.dev/blog/fuzz-beta

Fuzzing is a beta feature in go1.17, so you'll need to install gotip

```
$ go get golang.org/dl/gotip
$ gotip download dev.fuzz
```

After that you should be able to run fuzz tests locally:
```
make fuzz
```
See Makefile for other stuff