# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.11.1

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/ildarusmanov/msofficepreview

# setup dependencies
WORKDIR /go/src/github.com/ildarusmanov/msofficepreview
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure


RUN go install github.com/ildarusmanov/msofficepreview

# Run the command by default when the container starts.
ENTRYPOINT /go/bin/msofficepreview -configfile "/go/src/github.com/ildarusmanov/msofficepreview/config.yml"

# Document that the service listens on port 8001.
EXPOSE 8001