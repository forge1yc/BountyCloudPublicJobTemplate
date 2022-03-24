FROM ubuntu:16.04
RUN apt-get update
RUN apt-get install -y ca-certificates
RUN update-ca-certificates
COPY bin/linux_bounty_email /bin/
CMD ["/bin/linux_bounty_email"]