# build go binary:
FROM golang:1.23-alpine AS builder

# workking directory container:
WORKDIR /app

# copy local files into /app directory
COPY go.mod ./
COPY main.go ./

# build/compile
RUN go build -o filehasher main.go



# run
FROM alpine:latest
# NOTE: supposedly, minimal image for just binary, not go or build tools

# set working directory
WORKDIR /root/

# copy compiled binary from filehasher builder stage
#   to this new lightweight container
COPY --from=builder /app/filehasher .


# set default command
#   apparently uses json array syntax
#   will run by default when someone uses the image
ENTRYPOINT ["./filehasher"]
