.PHONY: build
build:
	go build -o cpusched

.PHONY: delete
delete:
	rm cpusched

.PHONY: run
run: build
	./cpusched

.PHONY: debug
debug: build
	rm cpusched
	go build -o cpusched
	./cpusched -n 24 -total 1000 -resol 10