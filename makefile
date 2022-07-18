BINARY=cyber-waffle

build_go:
	go build -o ./go/${BINARY}-go ./go/main.go ./go/LazarusUI.go ./go/configWaffle.go ./go/mix_scopeWaffle.go
	
run_go:
	./go/${BINARY}-go

go: build_go run_go

clean_go:
	rm -f ./go/${BINARY}-go

build_c:
	gcc -o ./c/${BINARY}-c ./c/main.c -lcurl
	
run_c:
	./c/${BINARY}-c

c: build_c run_c

clean_c:
	rm -f ./c/${BINARY}-c