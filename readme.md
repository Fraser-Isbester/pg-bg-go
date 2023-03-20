# Blue-Green PostgreSQL Cutover with Logical Replication and pgcat

This project sets up a blue-green PostgreSQL database cutover using logical replication and the pgcat protocol-aware load balancer.

## Prerequisites

- Docker
- Docker Compose

## Project Structure
```
pg-bg-go/
├── cmd
│ └── main.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── internal
│ ├── db
│ │ ├── db.go
│ │ └── replication.go
│ └── utils
│   └── utils.go
└── README.md
```

## Setup

1. Clone the repository:
```shell
git clone https://github.com/fraser-isbester/pg-bg-go.git
cd pg-bg-go
```

2. Start the PostgreSQL databases and pgcat load balancer:

```shell
docker-compose up -d
```

3. Build and run the Go application:

```shell
docker build -t blue-green-postgres .
docker run --rm --network blue-green-postgres_db_network blue-green-postgres
```

## Usage

After running the Go application, it will set up the logical replication between the blue and green PostgreSQL databases and perform the read-write cutover using the pgcat load balancer. Connect your application to the pgcat load balancer at `localhost:7432`.

## License

MIT License

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
