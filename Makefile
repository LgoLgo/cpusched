.PHONY: build
build:
	go build -o cpusched

.PHONY: delete
delete:
	rm cpusched

.PHONY: run
run:
	./cpusched

.PHONY: debug
debug:
	rm cpusched
	go build -o cpusched
	./cpusched -n 10 -total 1000 -resol 100 > res1.txt
	python3 print.py 1
	./cpusched -n 20 -total 1000 -resol 100 > res2.txt
	python3 print.py 2
	./cpusched -n 30 -total 1000 -resol 100 > res3.txt
	python3 print.py 3
