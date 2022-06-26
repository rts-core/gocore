FROM golang:1.18-buster

ENV GOPATH /
RUN go install golang.org/x/tools/cmd/godoc@latest
COPY ./script.sh /bin/script.sh

CMD ["/bin/script.sh"]