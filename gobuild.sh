echo "Formatting Go source code in project using go fmt..."
go fmt $(go list ./...)
echo "Running go vet on Go source code in project. Any reported issues may lead to issues building..."
go vet $(go list ./...)
echo "Running golint on Go source code in project. Please try to address reported issues..."
#golint $(go list ./...)
echo "Running tests with coverage on Go source code in project..."
go test -cover $(go list ./...)
echo "Building the executable for the project..."
go build -o applecrud applecrud/cmd/svr
