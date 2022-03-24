FROM ubuntu:16.04
RUN apt-get update
RUN apt-get install -y ca-certificates
RUN update-ca-certificates
COPY bin/linux_bounty_hello_service /bin/
EXPOSE 8080
CMD ["/bin/linux_bounty_hello_service"]