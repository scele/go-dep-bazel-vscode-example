<p align="center">
  <img src="https://blog.golang.org/gopher/gopher.png" height="180"><img src="https://raw.githubusercontent.com/golang/dep/master/docs/assets/DigbyShadows.png" height="180"><img src="https://bazel.build/images/bazel-icon.svg" height="180"><img src="https://user-images.githubusercontent.com/49339/32078472-5053adea-baa7-11e7-9034-519002f12ac7.png" height="180">
</p>

## Example setup for go + dep + bazel + vscode

This project shows how to use [rules_go_dep](https://github.com/scele/rules_go_dep) to set up a development enviromnent.

### Prerequisites

It is assumed that dep, bazel and vscode are already installed.

Currently the setup requires a custom build of the `vscode-go` extension, which can be installed with the following commands:

```bash
wget https://github.com/scele/vscode-go/releases/download/0.6.72-lpeltonen.2/Go-0.6.72-lpeltonen.2.vsix
code --install-extension ./Go-0.6.72-lpeltonen.2.vsix
rm ./Go-0.6.72-lpeltonen.2.vsix
```

#### MacOS

- [Install delve](https://github.com/derekparker/delve/blob/master/Documentation/installation/osx/install.md) (you may need to apply [this workaround](https://github.com/go-delve/homebrew-delve/issues/19#issuecomment-330442033) if the installation fails)

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
