# goImportVarShadowLint

## What is this??

`analyzer.go` is a linter to detect and fix variable shadowing against imported packages in Go code.\
It is generally not good practice to name a variable as it shadows an imported package. This linter will help you to
detect such cases.

> [!NOTE]
> I did not find any linters available as OSS to detect this at first, so I created this linter.
> However it turned out that [go-critic](https://github.com/go-critic/go-critic) does support this in `importShadow` check.
> 
> `go-critic` does not support suggesting fixes for this, but this linter does, so this can still be useful (Support is limited)
> 
> It is recommended to use `go-critic` instead of this linter, as it is available as a part of `golangci-lint`. \
> You can write config like below to enable `importShadow` check in `golangci-lint`:
> ```yaml
> version: 2
> linters:
>   enable:
>     - gocritic
>   settings:
>     gocritic:
>       disable-all: true
>       enabled-checks:
>         - importShadow
> ```

## Usage

1. Clone this repository
2. Run `go install ./cmd/goImportVarShadowLint`
3. Run `goImportVarShadowLint ./...` in your project directory
