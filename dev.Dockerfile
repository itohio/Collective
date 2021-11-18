FROM golang:1.17-alpine as backend
#RUN go install github.com/cespare/reflex@latest
RUN go install github.com/itohio/xnotify@v0.3.1
#RUN go install github.com/makiuchi-d/arelo@latest
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
# ENTRYPOINT [ "arelo", "-p", "'**/*.go'", "-i", "'**/.*'", "-i", "'**/*_test.go'" ]
ENTRYPOINT [ "xnotify", "-i", ".", "--batch", "100" ]