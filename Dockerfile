# multitstage build since Go apps produce a standalone executable (recommended here: https://docs.docker.com/build/building/multi-stage/)

FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
# will only download again if go mod or sum changes
RUN go mod download
COPY . .
# will copy executable into production image
RUN go build -o /out/api ./cmd/go-acis

# now image doesnt even need go complier or source code
FROM alpine:3.23
WORKDIR /app
COPY --from=build /out/api ./api
USER nobody
EXPOSE 8080
CMD ["./api"]
