# feedks-api

## Pre Setup

- Create a copy of the .env.example file and rename it to .env OR run `make env` IF you are in Unix

## Run

- Install all modules

```sh
go mod download
```

- install all modules (IF you are in UNIX)

```sh
make install
```

- Build application

```sh
go build -o bin/feedks-api ./cmd/main.go
```

- Build application (if you are in Unix)

```sh
make build
```

- Start application 

```sh
./bin/feedks-api.exe
```

- Start application (IF you are in Unix)

```sh
make run
```

## Run w/ Docker