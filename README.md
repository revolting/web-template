# Leaves

## Install dependencies

```
govendor sync
```

## Start local server

```
go run server.go
```

If you need to change the default flags for the port and whether it is a dev/prod server, where isDev defaults to true:

```
go run server.go -port=:8000 -isDev=false
```
