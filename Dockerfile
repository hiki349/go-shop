FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src/app

RUN apk --no-cache add bash git make gcc musl-dev

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# copy source
COPY cmd ./cmd
COPY configuration ./configuration
COPY internal ./internal
COPY migrations ./migrations
COPY pkg ./pkg

# build
RUN go build -o bin/go-shop -ldflags "-s -w" -a cmd/main/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/app/bin/go-shop /bin
# copied default enviroments
COPY .env ./.env

# run app
CMD [ "./bin/go-shop" ]