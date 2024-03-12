# docker build -t smoke-or-fire .
# INTERACTIVE MODE FOR USER INPUT: 
# docker run -it smoke-or-fire

###################################################
# Use go image for builder stage
FROM golang:1.22.1 AS builder

WORKDIR /app

# COPY go.mod ./
# RUN go mod download

# COPY *.go ./
COPY . .

# Version for builder stage means no preceeding /
# RUN go build -o /smoke-or-fire
RUN go build -o smoke-or-fire

# Use a smaller base image for the final stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the binary from the builder stage
COPY --from=builder /app/smoke-or-fire .

# Command to run the binary
# EXPOSE 8080
CMD ["./smoke-or-fire"]