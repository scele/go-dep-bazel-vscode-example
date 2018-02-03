load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/scele/go-dep-bazel-vscode-example/cmd/example",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/something:go_default_library",
        "@com_github_sirupsen_logrus//:go_default_library",
    ],
)

go_binary(
    name = "example",
    embed = [":go_default_library"],
    importpath = "github.com/scele/go-dep-bazel-vscode-example/cmd/example",
    visibility = ["//visibility:public"],
    gc_linkopts = [
        "-linkmode",
        "internal",
    ],
)
