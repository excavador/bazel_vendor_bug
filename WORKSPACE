git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "ad076c97335c03319280fa438c9f958e7c81597a",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains(go_version="1.9.1")

#load("@io_bazel_rules_go//proto:def.bzl", "proto_register_toolchains")
#proto_register_toolchains()
