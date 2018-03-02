# review-tools

## Introduction
This is a simple tool that collects label information from Github Pull Requests. In the future it may become something more, hence a more generic name... :)

## Installation
1. Clone the repository
2. Run `dep ensure`
3. Run `go install`

Make sure to have your `$GOPATH` added as part of `$PATH`!!!

## How to run
### Basic
```
$ review-tools label-stats -o golang -r go
```

### Private repositories
```
$ review-tools label-stats -o golang -r go -t <OAuth Token>
```

### Help
```
$ review-tools                  # Basic help
$ review-tools -h               # The same as above
$ review-tools label-stats -h   # Sub-command help
```

## Env
You can omit passing the options into the command by defining env variables. Please refer to the in-command help to identify the right env variables.
