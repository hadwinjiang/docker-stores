// https://w.amazon.com/bin/view/AWS/Mobile/API_Gateway/NewHire/GoLangDevSetUp/#HGoLang
// go env -w GOPROXY=direct

// Build Docker image
docker build -t booklibrary .

// Run Docker
docker run -p 8080:8080 -it booklibrary

// Run Docker with environment
docker run -d --env PORT=8082 -p 8080:8082 -it booklibrary
docker run --env PORT=8082 -p 8080:8082 -it booklibrary
docker run --env PORT=8082 --env RUNTIME_SETUP=dev -p 8080:8082 -it booklibrary
docker run --env-file envlist.txt -p 8080:8082 -it booklibrary

// *** GO *** 
// init go module
go mod init github.com/booklibrary

// download package
go mod tidy

// build go project
go build main.go
