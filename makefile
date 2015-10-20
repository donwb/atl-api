all:
	go run *.go

linux:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o atl-api.linux *.go