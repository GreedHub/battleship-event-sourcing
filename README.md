### Force ssh github use in `~/.gitconfig`
```sh
[url "ssh://git@github.com/"]
        insteadof = https://github.com/
```

### envs
```sh
export GO111MODULE=on
export GOPATH="$HOME/go"
export GOPRIVATE="github.com/GreedHub"
export PATH="$GOPATH/bin:$PATH"
```