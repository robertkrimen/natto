.PHONY: test test-race test-release release

test: test-release

test-release:
	go test -i
	go test

test-race:
	go test -race -i
	go test -race

release: test-race test-release
	for package in . ; do (cd $$package && godocdown --signature > README.markdown); done
	@echo PASS
