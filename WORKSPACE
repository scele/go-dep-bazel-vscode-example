workspace(name="com_github_scele_go_dep_bazel_vscode_examples")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.9.0/rules_go-0.9.0.tar.gz",
    sha256 = "4d8d6244320dd751590f9100cf39fd7a4b75cd901e1f3ffdfd6f048328883695",
)

local_repository(
    name = "com_github_scele_rules_go_dep",
    # TODO: Use http_archive and SHA
    path = "/home/lpeltonen/go/src/github.com/scele/rules_go_dep",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@com_github_scele_rules_go_dep//dep:dep.bzl", "dep_import")

go_register_toolchains(go_version="1.9")
go_rules_dependencies()

dep_import(
    name = "godeps",
    gopkg_lock = "//:Gopkg.lock",
)
load("@godeps//:Gopkg.bzl", "go_deps")
go_deps()
