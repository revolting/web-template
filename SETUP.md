# Setting up your Go environment (for people who don't want to RTFM)

## Change soft tabs to hard tabs

I know ... I know.

You can customize for your editor of course https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

## Install and set up initial subdirectories

Install go, e.g. `brew install go`

Make a directory somewhere where you want all your Go projects to be, e.g. `$HOME/code/go`. Make two subdirectories, `bin` and `src` directly under like so:

```
-> go
  -> bin
  -> src
```

## Set your paths

Edit your .bashrc or .profile or whatever you use and add this:

```
export GOPATH=$HOME/code/go
PATH=$PATH:$GOPATH/bin
```

Save the file and open a new terminal window to get the latest .bashrc/.profile changes or just type something like this:

```
source ~/.profile
```

## Install prerequisite dependencies

Install govendor:

```
go get -u github.com/kardianos/govendor
```

That will put govendor into `$HOME/code/go/bin`. All your Go binaries will also be in the same subdirectory.

## Clone the Github repository

Go into `$HOME/code/go/src` and add a subdirectory for github.com repos or wherever you get your repositories:

```
-> go
  -> bin
  -> src
    -> github.com
      -> revolting
```

Change into the src/github.com/revolting directory and `git clone https://github.com/revolting/leaves.git`
