# Leaves

## If this is a whole new world for you

Check out [SETUP.md](https://github.com/revolting/leaves/blob/master/SETUP.md) and change into the project root.

## Install dependencies

```
govendor sync
```

## Build

```
go build
```

## Start local server

```
./leaves
```

If you need to change the default flags for the port and whether it is a dev/prod server, where isDev defaults to true:

```
./leaves -port=:8080 -isDev=false -twilioSid=111 -twilioToken=111 -twilioPhone=5555555
```
Visit http://localhost:8080 (or whatever your chose your port as)
