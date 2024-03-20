FROM golang:1.22

# set current working directory
WORKDIR /app

# download modules
COPY go.mod go.sum ./
RUN go mod download

# copy source code
WORKDIR /app/src
COPY ./src .

# also copy the environment file
COPY .env .

# build the app
RUN go build -o /groundhog

# run the executable
CMD ["/groundhog"]