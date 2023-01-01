build:
	go build ./cmd/todo 

new:
	go run ./cmd/todo -add 

list:
	go run ./cmd/todo -list
