FROM golang:1.23.7 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0 

RUN go build -o skillsRest cmd/skillsRest/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/skillsRest /app/skillsRest
RUN chmod +x /app/skillsRest

EXPOSE 3000

CMD sh -c "/app/skillsRest"
# FROM alpine:latest

# WORKDIR /app

# COPY --from=builder /app/skillsRest /app/skillsRest

# EXPOSE 3000

# CMD ["./skillsRest"]
