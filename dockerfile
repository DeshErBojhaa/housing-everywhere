# Build the Go Binary.
FROM golang:1.13 as dns
ENV CGO_ENABLED 0
#ARG VCS_REF
#ARG PACKAGE_NAME
#ARG PACKAGE_PREFIX

# Create a location in the container for the source code. Using the
# default GOPATH location.
RUN mkdir -p /go/src/github.com/DeshErBojhaa/service

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
COPY go.* /go/src/github.com/DeshErBojhaa/service/
WORKDIR /go/src/github.com/DeshErBojhaa/service
RUN go mod download

# Copy the source code into the container.
COPY cmd cmd
COPY handlers handlers
COPY web web

# Build the service binary. We are doing this last since this will be different
# every time we run through this process.
WORKDIR /go/src/github.com/DeshErBojhaa/service/cmd
RUN set GO111MODULE=on 
RUN go build -mod=readonly -o=main


# Run the Go Binary in Alpine.
FROM alpine:3.7
COPY --from=dns /go/src/github.com/DeshErBojhaa/service/cmd/main /app/main
WORKDIR /app
CMD /app/main