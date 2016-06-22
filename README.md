# Leaves

## If this is a whole new world for you

Check out [SETUP.md](https://github.com/revolting/leaves/blob/master/SETUP.md)

## Install dependencies

```
govendor sync
```

## Start local server

```
go run main.go
```

If you need to change the default flags for the port and whether it is a dev/prod server, where isDev defaults to true:

```
go run main.go -port=:8000 -isDev=false
```
