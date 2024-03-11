# TODO: current app goes into crash loop -- no user input handling -- however, webassembly changes will likely adjust this

FROM golang:1.22.1

WORKDIR /app

# COPY go.mod ./
# RUN go mod download

# COPY *.go ./

COPY . .

# RUN go build -o /smoke-or-fire

# For Gitlab we need the following:
RUN CGO_ENABLED=0 GOOS=linux go build -o /smoke-or-fire

# EXPOSE 8080

CMD [ "/smoke-or-fire" ]

# docker build -t smoke-or-fire .
# docker run smoke-or-fire