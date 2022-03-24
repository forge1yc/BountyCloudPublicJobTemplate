FROM ubuntu:16.04
RUN apt-get update

RUN apt-get install -y ca-certificates
RUN update-ca-certificates

COPY bin/bounty_cpu_ram_service /bin/
CMD ["/bin/bounty_cpu_ram_service"]