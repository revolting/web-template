# Leaves

## If this is a whole new world for you

Check out [SETUP.md](https://github.com/revolting/leaves/blob/master/SETUP.md) and change into the project root.

## Install dependencies

```
govendor sync
```

## Start local server

```
go run *.go
```

If you need to change the default flags for the port and whether it is a dev/prod server, where isDev defaults to true:

```
go run *.go -port=:8000 -isDev=false
```
Visit http://localhost:8080 (or whatever your chose your port as)