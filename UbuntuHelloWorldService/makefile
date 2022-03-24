# linux文件
#build_mac:
	#go build -o bin/mac_bounty_email cmd/main_email.go

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux_bounty_email cmd/main_email.go

#build_win:
#	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/win_bounty_eamil cmd/main_email.go

buildHelloServiceAll:
	go build -o bin/mac_bounty_hello_service cmd/bounty_hello_service.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux_bounty_hello_service cmd/bounty_hello_service.go
#	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/win_bounty_eamil cmd/main_email.go

dockerBuild:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bounty_hello_service cmd/bounty_hello_service.go
	docker build -t bountycloud/bounty_hello_service:v$(version) .

