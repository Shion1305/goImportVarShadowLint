# goImportVarShadowLint

## What is this??

`analyzer.go` is a linter to detect and fix variable shadowing against imported packages in Go code.\
It is generally not good practice to name a variable in a way that shadows an imported package. This linter helps you
detect such cases.

> [!NOTE]
> I did not find any linters available as OSS to detect this issue initially, so I created this linter.
> However, it turned out that [go-critic](https://github.com/go-critic/go-critic) does support this via the `importShadow` check.
> 
> While `go-critic` does not support suggesting fixes for this issue, this linter does (though with limited support).
> 
> It is recommended to use `go-critic` instead of this linter, as it is available as part of `golangci-lint`. \
> You can use the following configuration to enable the `importShadow` check in `golangci-lint`:
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
