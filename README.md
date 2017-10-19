Introduction
============

```
➜ bazel version
..........
Build label: 0.7.0-homebrew
Build target: bazel-out/darwin_x86_64-opt/bin/src/main/java/com/google/devtools/build/lib/bazel/BazelServer_deploy.jar
Build time: Thu Oct 19 09:12:48 2017 (1508404368)
Build timestamp: 1508404368
Build timestamp as int: 1508404368
```


First problem
=============
https://github.com/bazelbuild/bazel/issues/3928

```
➜ bazel build -s //vendor/github.com/openzipkin/zipkin-go-opentracing/wire:go_default_library
..........
Unhandled exception thrown during build; message: CONFIGURED_TARGET://vendor/github.com/openzipkin/zipkin-go-opentracing/wire:go_default_library 558e55e56229d613c257787a2509ec34 (354423329 1308709877) -> <ErrorInfo exception=com.google.devtools.build.lib.skyframe.ToolchainResolutionFunction$NoToolchainFoundException: no matching toolchain found for @io_bazel_rules_go//proto:proto rootCauses={ToolchainResolutionKey{configuration=558e55e56229d613c257787a2509ec34, toolchainType=@io_bazel_rules_go//proto:proto, targetPlatform=struct(constraints = [struct(constraint = struct(label = Label("@bazel_tools//platforms:cpu")), label = Label("@bazel_tools//platforms:x86_64")), struct(constraint = struct(label = Label("@bazel_tools//platforms:os")), label = Label("@bazel_tools//platforms:osx"))], label = Label("@bazel_tools//platforms:target_platform")), execPlatform=struct(constraints = [struct(constraint = struct(label = Label("@bazel_tools//platforms:cpu")), label = Label("@bazel_tools//platforms:x86_64")), struct(constraint = struct(label = Label("@bazel_tools//platforms:os")), label = Label("@bazel_tools//platforms:osx"))], label = Label("@bazel_tools//platforms:host_platform"))}} cycles=[]>
INFO: Elapsed time: 1.810s
java.lang.IllegalStateException: CONFIGURED_TARGET://vendor/github.com/openzipkin/zipkin-go-opentracing/wire:go_default_library 558e55e56229d613c257787a2509ec34 (354423329 1308709877) -> <ErrorInfo exception=com.google.devtools.build.lib.skyframe.ToolchainResolutionFunction$NoToolchainFoundException: no matching toolchain found for @io_bazel_rules_go//proto:proto rootCauses={ToolchainResolutionKey{configuration=558e55e56229d613c257787a2509ec34, toolchainType=@io_bazel_rules_go//proto:proto, targetPlatform=struct(constraints = [struct(constraint = struct(label = Label("@bazel_tools//platforms:cpu")), label = Label("@bazel_tools//platforms:x86_64")), struct(constraint = struct(label = Label("@bazel_tools//platforms:os")), label = Label("@bazel_tools//platforms:osx"))], label = Label("@bazel_tools//platforms:target_platform")), execPlatform=struct(constraints = [struct(constraint = struct(label = Label("@bazel_tools//platforms:cpu")), label = Label("@bazel_tools//platforms:x86_64")), struct(constraint = struct(label = Label("@bazel_tools//platforms:os")), label = Label("@bazel_tools//platforms:osx"))], label = Label("@bazel_tools//platforms:host_platform"))}} cycles=[]>
	at com.google.common.base.Preconditions.checkState(Preconditions.java:721)
	at com.google.devtools.build.lib.util.Preconditions.checkState(Preconditions.java:238)
	at com.google.devtools.build.lib.skyframe.SkyframeBuildView.assertSaneAnalysisError(SkyframeBuildView.java:412)
	at com.google.devtools.build.lib.skyframe.SkyframeBuildView.configureTargets(SkyframeBuildView.java:278)
	at com.google.devtools.build.lib.analysis.BuildView.update(BuildView.java:600)
	at com.google.devtools.build.lib.buildtool.BuildTool.runAnalysisPhase(BuildTool.java:505)
	at com.google.devtools.build.lib.buildtool.BuildTool.buildTargets(BuildTool.java:214)
	at com.google.devtools.build.lib.buildtool.BuildTool.processRequest(BuildTool.java:341)
	at com.google.devtools.build.lib.runtime.commands.BuildCommand.exec(BuildCommand.java:67)
	at com.google.devtools.build.lib.runtime.BlazeCommandDispatcher.execExclusively(BlazeCommandDispatcher.java:600)
	at com.google.devtools.build.lib.runtime.BlazeCommandDispatcher.exec(BlazeCommandDispatcher.java:351)
	at com.google.devtools.build.lib.runtime.CommandExecutor.exec(CommandExecutor.java:58)
	at com.google.devtools.build.lib.server.GrpcServerImpl.executeCommand(GrpcServerImpl.java:850)
	at com.google.devtools.build.lib.server.GrpcServerImpl.access$2100(GrpcServerImpl.java:109)
	at com.google.devtools.build.lib.server.GrpcServerImpl$2.lambda$run$0(GrpcServerImpl.java:915)
	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)
	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)
	at java.lang.Thread.run(Thread.java:745)
java.lang.IllegalStateException: CONFIGURED_TARGET://vendor/github.com/openzipkin/zipkin-go-opentracing/wire:go_default_library 558e55e56229d613c257787a2509ec34 (354423329 1308709877) -> <ErrorInfo exception=com.google.devtools.build.lib.skyframe.ToolchainResolutionFunction$NoToolchainFoundException: no matching toolchain found for @io_bazel_rules_go//proto:proto rootCauses={ToolchainResolutionKey{configuration=558e55e56229d613c257787a2509ec34, toolchainType=@io_bazel_rules_go//proto:proto, targetPlatform=struct(constraints = [struct(constraint = struct(label = Label("@bazel_tools//platforms:cpu")), label = Label("@bazel_tools//platforms:x86_64")), struct(constraint = struct(label = Label("@bazel_tools//platforms:os")), label = Label("@bazel_tools//platforms:osx"))], label = Label("@bazel_tools//platforms:target_platform")), execPlatform=struct(constraints = [struct(constraint = struct(label = Label("@bazel_tools//platforms:cpu")), label = Label("@bazel_tools//platforms:x86_64")), struct(constraint = struct(label = Label("@bazel_tools//platforms:os")), label = Label("@bazel_tools//platforms:osx"))], label = Label("@bazel_tools//platforms:host_platform"))}} cycles=[]>
	at com.google.common.base.Preconditions.checkState(Preconditions.java:721)
	at com.google.devtools.build.lib.util.Preconditions.checkState(Preconditions.java:238)
	at com.google.devtools.build.lib.skyframe.SkyframeBuildView.assertSaneAnalysisError(SkyframeBuildView.java:412)
	at com.google.devtools.build.lib.skyframe.SkyframeBuildView.configureTargets(SkyframeBuildView.java:278)
	at com.google.devtools.build.lib.analysis.BuildView.update(BuildView.java:600)
	at com.google.devtools.build.lib.buildtool.BuildTool.runAnalysisPhase(BuildTool.java:505)
	at com.google.devtools.build.lib.buildtool.BuildTool.buildTargets(BuildTool.java:214)
	at com.google.devtools.build.lib.buildtool.BuildTool.processRequest(BuildTool.java:341)
	at com.google.devtools.build.lib.runtime.commands.BuildCommand.exec(BuildCommand.java:67)
	at com.google.devtools.build.lib.runtime.BlazeCommandDispatcher.execExclusively(BlazeCommandDispatcher.java:600)
	at com.google.devtools.build.lib.runtime.BlazeCommandDispatcher.exec(BlazeCommandDispatcher.java:351)
	at com.google.devtools.build.lib.runtime.CommandExecutor.exec(CommandExecutor.java:58)
	at com.google.devtools.build.lib.server.GrpcServerImpl.executeCommand(GrpcServerImpl.java:850)
	at com.google.devtools.build.lib.server.GrpcServerImpl.access$2100(GrpcServerImpl.java:109)
	at com.google.devtools.build.lib.server.GrpcServerImpl$2.lambda$run$0(GrpcServerImpl.java:915)
	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)
	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)
	at java.lang.Thread.run(Thread.java:745)
```

Second problem
==============

Uncomment in WORKSPACE
````
#load("@io_bazel_rules_go//proto:def.bzl", "proto_register_toolchains")
#proto_register_toolchains()
```

