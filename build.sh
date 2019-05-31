echo "cleaning the code..."
goimports -w .
echo "building the ginsight cli..."
go build -o ginsight
