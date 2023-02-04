clean:
	rm -rvf .data > /dev/null 2>&1

make run:
	make clean
	go run .
