workspace(name="com_github_scele_go_dep_bazel_vscode_example")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/scele/rules_go/archive/5c5051105c833287138b5e8c8361b82dc6e25721.tar.gz"],
    strip_prefix = "rules_go-5c5051105c833287138b5e8c8361b82dc6e25721",
    sha256 = "ce9a65847d38ace1ac1a43f3922f0671560ec63ea7011158822f3388a9f5de81",
)

http_archive(
    name = "com_github_scele_rules_go_dep",
    urls = ["https://github.com/scele/rules_go_dep/archive/a814fad7f886bbe0fbd79df8549029708d570a37.tar.gz"],
    strip_prefix = "rules_go_dep-a814fad7f886bbe0fbd79df8549029708d570a37",
    sha256 = "3fe2c7cd2a629e8e55c5f6fef778ec753785a6da8f3435f45d7fc7b7ad78dfdb",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@com_github_scele_rules_go_dep//dep:dep.bzl", "dep_import")

go_register_toolchains(go_version="1.9")
go_rules_dependencies()

dep_import(
    name = "godeps",
    prefix = "github.com/scele/go-dep-bazel-vscode-example",
    gopkg_lock = "//:Gopkg.lock",
)
load("@godeps//:Gopkg.bzl", "go_deps")
go_deps()
