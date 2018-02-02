<p align="center">
  <img src="https://blog.golang.org/gopher/gopher.png" height="180"><img src="https://raw.githubusercontent.com/golang/dep/master/docs/assets/DigbyShadows.png" height="180"><img src="https://bazel.build/images/bazel-icon.svg" height="180"><img src="https://user-images.githubusercontent.com/49339/32078472-5053adea-baa7-11e7-9034-519002f12ac7.png" height="180">
</p>

## Example setup for go + dep + bazel + vscode

This project shows how to use [rules_go_dep](https://github.com/scele/rules_go_dep) to set up a development enviromnent.

### Clone this example project

```bash
go get -u github.com/scele/go-dep-bazel-vscode-example
cd $GOPATH/src/github.com/scele/go-dep-bazel-vscode-example
```

> **NOTE:** The build does not require the source to be under `$GOPATH`, but `dep ensure` does.

### Build and run with bazel

```bash
bazel build //...
bazel run //cmd/example
```

### Use vscode to debug

```bash
code -n .
```

In vscode, open `cmd/example/main.go` and start debugging using the pre-configured launch config.

### Update go dependencies

```bash
dep ensure
rm -rf ./vendor
```

The `vendor` directory should be manually deleted, since the dependencies will be pulled through bazel.

### Update BUILD files

```bash
bazel run //:gazelle
```
