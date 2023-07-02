# NATS - STREAMING service

Nats-streaming service using ```gorilla/mux```, ```nats-streaming```, ```postgreSQL```.

## Video demonstration

https://github.com/errrov/nats-streaming-service/assets/79909234/73ed3fd2-c63d-4b0e-914f-1c81576605b1

## Running nats-streaming and postgreSQL

To run nats-streaming and postgreSQL

- ```docker-compose up```

## Running subscriber

- ``` go run ./cmd/subscriber/main.go```

    For publisher you can run

- ``` go run ./cmd/publisher/main.go```

