install:
	go install

test:
	go test -timeout 30s

fuzz: # default fuzzing
	gotip test -fuzz=FuzzHostsfileToString

clean-fuzz: # reruns fuzzing without cache and previously found errors
	rm -rf ./testdata
	gotip clean -fuzzcache
	gotip test -fuzz=FuzzHostsfileToString