# Java with Go interoperability?

Say, we have a java classes with a lot of code that can not reasonably be ported to golang. What do we do?

We call java from golang 🤪

In this project we compile java to native ahead of time
https://www.graalvm.org/native-image/

Then, import it in go with cgo
https://pkg.go.dev/cmd/cgo

And run. Hello-world level.

# Requirements

Install graal, export JAVA_HOME, GAALVM_HOME, PATH.
https://www.graalvm.org/docs/getting-started/
```shell
export GRAALVM_HOME=/Library/Java/JavaVirtualMachines/graalvm-ce-java17-21.3.0/Contents/Home
export JAVA_HOME="$GRAALVM_HOME"
export PATH="$JAVA_HOME/bin:$PATH"
```

Install graal's native-image https://www.graalvm.org/reference-manual/native-image/#install-native-image

Download and install go https://go.dev/doc/install

# Build

One console for java, `cd java`
```shell
mvn -Pnative -DskipTests package
```
And one for go, `cd go`
```shell
go run main.go
```

# TODO

* create stateful java class instance from go
* call methods
* convert call result to go types
* profile to see how slow it is and how much memory consumed
* use make tool (Makefile?)
