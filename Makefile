.PHONY: gen
gen:
	$(MAKE) -C testdata

.PHONY: test
test:
	go test -v ./...

.PHONY: bench
bench: $(wildcard testdata/*.json)
	go test -bench . -args $^
