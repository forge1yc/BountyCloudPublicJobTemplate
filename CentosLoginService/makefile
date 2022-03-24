# linux文件
#build_mac:
	#go build -o bin/mac_bounty_email cmd/main_email.go

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux_bounty_email cmd/main_email.go

#build_win:
#	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/win_bounty_eamil cmd/main_email.go

buildEmailAll:
#	go build -o bin/mac_bounty_email cmd/main_email.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/bounty_hello_service cmd/main_hello_service.go
#	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/win_bounty_eamil cmd/main_email.go

#buildCentosAll:
#	go build -o bin/mac_bounty_centos cmd/main_centos.go

