version=$1
make buildHelloServiceAll
docker build -t bountycloud/bounty_hello_service:"${version}" .
