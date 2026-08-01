default:
	go build -ldflags "-s -w" -o busy-bar-cli ./cli/main.go