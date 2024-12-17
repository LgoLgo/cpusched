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
	./cpusched -n 15 -total 10000 -resol 1000