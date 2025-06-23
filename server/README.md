# Getting started

First you should have the installed Go in your system and exported all the requried path. To getting started with the setup of the server your present working directory should be `/server`.

Now run the following commands to get started:

```
go mod tidy
```

To start the server there are two options:

1. Run only server - without hot reloading
2. Run server with hot reloading (using `air`)

Commands:

1. for first one

```
go run main.go
```

2. for second one

```
air -c .air.toml
```
