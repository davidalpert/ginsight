#!/bin/bash
echo "getting goimports..."
go get golang.org/x/tools/cmd/goimports

echo "installing goimports..."
go install golang.org/x/tools/cmd/goimports

echo "cleaning the code..."
goimports -w .

echo "building the ginsight cli..."
go build -o ginsight
