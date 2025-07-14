FROM golang:1.24
WORKDIR /src/ChrisMcKee/teamcity-sdk-go
COPY . .
RUN go mod download
RUN go build
RUN go test -v -c ./types && \
    go test -v -c ./teamcity
CMD ./run-test.sh
