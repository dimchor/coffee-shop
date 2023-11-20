FROM golang:1.21

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only 
# redownloading them in subsequent builds if they change
COPY . ./
RUN go mod download 

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

WORKDIR /usr/local/bin/

CMD ["app"]
