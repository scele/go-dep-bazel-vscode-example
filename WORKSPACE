workspace(name="com_github_scele_go_dep_bazel_vscode_examples")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.9.0/rules_go-0.9.0.tar.gz"],
    sha256 = "4d8d6244320dd751590f9100cf39fd7a4b75cd901e1f3ffdfd6f048328883695",
)

http_archive(
    name = "com_github_scele_rules_go_dep",
    urls = ["https://github.com/scele/rules_go_dep/archive/70fa72816cae64f67634740bc6c0233f39b5d8c6.tar.gz"],
    strip_prefix = "rules_go_dep-70fa72816cae64f67634740bc6c0233f39b5d8c6",
    sha256 = "20eeac91a621af97a39e9e30848727d6c14c7d68fcf3fcc139e98e3c363b4661",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@com_github_scele_rules_go_dep//dep:dep.bzl", "dep_import")

go_register_toolchains(go_version="1.9")
go_rules_dependencies()

dep_import(
    name = "godeps",
    prefix = "github.com/scele/go-dep-bazel-vscode-examples",
    gopkg_lock = "//:Gopkg.lock",
)
load("@godeps//:Gopkg.bzl", "go_deps")
go_deps()
