load("@io_bazel_rules_go//go:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    command = "fix",
    args = [
        "-build_file_name",
        "BUILD,BUILD.bazel",
    ],
    prefix = "github.com/scele/go-dep-bazel-vscode-example",
)
