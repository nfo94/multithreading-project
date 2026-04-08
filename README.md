### Go multithreading

This sample project is an example of a multithreaded application. To run this project you need
[Go](https://go.dev/) installed. After setting up Go, go to the root folder of the project
and install and tidy dependencies:

```bash
go mod tidy
```

Run the app searching for a CEP:

```bash
go run cmd/main.go {CEP}
```

This will call two different APIs, the fastest to respond will be shown in the terminal.
