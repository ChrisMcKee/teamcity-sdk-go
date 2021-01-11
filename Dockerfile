FROM golang:1.15
RUN go get -t -v "github.com/stretchr/testify"
COPY . $GOPATH/src/github.com/ChrisMcKee/teamcity-sdk-go
WORKDIR $GOPATH/src/github.com/ChrisMcKee/teamcity-sdk-go
RUN go test -v -c ./types && \
    go test -v -c ./teamcity
CMD ./run-test.sh
