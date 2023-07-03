# NATS - STREAMING service

Nats-streaming service using ```gorilla/mux```, ```nats-streaming```, ```postgreSQL```.

## Video demonstration



https://github.com/errrov/nats-streaming-service/assets/79909234/af4da905-c466-410c-9804-9300daad51e7


## Running nats-streaming and postgreSQL

To run nats-streaming and postgreSQL

- ```docker-compose up```

## Running subscriber

- ```go run ./cmd/subscriber/main.go```

    For publisher you can run

- ```go run ./cmd/publisher/main.go```

