# betterkismet

WARNING: Alpha do not use

## Golang instructions

### Creating folders and spaces

First install golang after that make sure you have the workspace exports in your `.bashrc` in our case

```
# Golang
export GOPATH="/home/YOURUSERNAME/workspace"
export PATH="$PATH:$GOPATH/bin"
```

And reload the `.bashrc`

```
source ~/.bashrc
```

In the workspace dir create the following folders and clone this repository

```
mkdir -p $GOPATH/src/github.com/brammittendorff
cd $GOPATH/src/github.com/brammittendorff
git clone git@github.com:brammittendorff/betterkismet.git
cd betterkismet
```

### Run app

```
go run main.go
```

### Build app

You can build normally with

```
go build
```

When you want to use less symbols you can compile with

```
go build -ldflags "-s -w"
```

And when you really want you binary to be small you can pack it easily with upx. Make sure you install this package before you try it

```
upx --brute betterkismet
```

### Run app binary

```
./betterkismet
```
