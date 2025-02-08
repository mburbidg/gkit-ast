.DEFAULT_GOAL := generate

generate:
	rm -fr java
	mkdir ./java
	protoc --java_out=./java *.proto
	rm -fr go
	mkdir ./go
	protoc --go_out=./go *.proto
.PHONY:generate
