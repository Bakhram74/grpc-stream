gen:
	protoc -I proto \
	proto/*.proto \
	--go_out=./pb --go_opt=paths=source_relative \
	--go-grpc_out=./pb \
	--go-grpc_opt=paths=source_relative

clean:
	rm pb/*.go

run:
	go run main.go



	# protoc --go_out=. --go-grpc_out=. proto/*.proto  