# Leaves

## If this is a whole new world for you

Check out [SETUP.md](https://github.com/revolting/leaves/blob/master/SETUP.md) and change into the project root.

## Install dependencies

```
go get github.com/pilu/fresh
govendor sync
```

## Build after making changes

```
go build
```

## Start local server

```
fresh
```

If you need to change the default flags for the port and whether it is a dev/prod server, where isDev defaults to true:

```
fresh -port=:8080 -isDev=true -twilioSid=111 -twilioToken=111 -twilioPhone=5555555
```
Visit http://localhost:8080 (or whatever your chose your port as)
