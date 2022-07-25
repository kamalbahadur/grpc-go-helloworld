FROM golang:1.16

ENV GO111MODULE=on
ENV APP_HOME /grpc-go-helloworld-server

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

# Copy files into app home
COPY . .

RUN go mod download

WORKDIR $APP_HOME/cmd

# Build the application
#RUN go mod tidy
#RUN go mod vendor
RUN go build -o main

# Export necessary port
EXPOSE 53080
EXPOSE 51080

# Command to run when starting the container
ENTRYPOINT ["sh", "-c", "${APP_HOME}/cmd/main"]