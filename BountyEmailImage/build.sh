version=$1
make buildEmailAll
docker build -t bountycloud/bounty_cloud_email:"${version}" .
