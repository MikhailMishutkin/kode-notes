FROM golang:1.23-alpine AS base

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o server ./cmd/main.go

# FROM scratch
# COPY --from=base /app/server /server
# ENTRYPOINT ["/server"]


# RUN apk update && apk add --no-cache git

# WORKDIR /usr/src/app
# COPY go.mod go.sum ./

# RUN go mod tidy
# RUN go mod download && go mod verify

# COPY . .

# FROM base AS build-http
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -o http ./cmd/httpserver/main.go
# #RUN go build -o /bin/http ./cmd/httpserver/main.go
# #WORKDIR /app
# #ENTRYPOINT [ "./http" ]




