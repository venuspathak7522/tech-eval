FROM golang:1.15-alpine

WORKDIR /tech-eval/go-service-skeleton

COPY ./go.mod ./
COPY ./go.sum ./

# download all the dependencies
RUN go mod download

# copy all the files into the container image filesnapshot
COPY . .

RUN go build -o main .

CMD ["/tech-eval/go-service-skeleton/main"]
