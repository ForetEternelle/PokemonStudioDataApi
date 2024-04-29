# psapi - Pokémon Studio Data API

`psapi` is a Go-based API designed for accessing data related to a Pokémon studio project. This tool provides both a command-line interface (CLI) and a Go library, allowing users to efficiently query and interact with various datasets. With `psapi`, you can seamlessly integrate data access into your applications or scripts.

## Features

- **Fast and lightweight:** Efficient data querying with minimal setup.
- **Flexible:** Customizable CORS settings, data directories, and log levels.
- **Go Library:** Easily integrate with Go-based applications.
- **OpenAPI Documentation:** Comprehensive documentation for the API and its endpoints.

---

## Installation

### Download Executable and OpenAPI Documentation

You can download the latest executable from the [releases page](https://github.com/rcharre/psapi/releases).

### Go install

Alternatively, you can install via the go packeage manager.

```
go install github.com/rcharre/psapi
```

---

## CLI Usage

You can use `psapi` through the command line to start a local server and interact with the Pokémon studio data.

### Basic Usage

```bash
psapi [flags]
```

### Available Flags

| Flag                     | Description                                     | Default |
|--------------------------|-------------------------------------------------|---------|
| `-cors=<string>`      | Set the CORS headers for the API                | "*"     |
| `-data=<string>`      | Specify the directory containing the data files | "data"  |
| `-log-level=<string>` | Set the logging level (e.g., DEBUG, INFO, WARN) | "INFO"  |
| `-port=<int>`         | Specify the port to run the API on              | 8000    |

### Example

Start the API server on port 8080 with DEBUG log level:

```bash
psapi -port=8080 -log-level=DEBUG
```

## Library Usage

You can also use `psapi` as a Go library to integrate Pokémon studio data access into your own applications.

For now the only dependency of the project is th Chi router implied by the openapi generator `go-server`.

### Install the Library

To include the library in your Go project, install it via `go get`:

```bash
go get github.com/rcharre/psapi
```

### Basic Example

Here is a basic example where the variable `dataFolder` contains the path to the folder containing Pokemon studio data.

```go
	store := studio.NewStore()

	if err := studio.Import(dataFolder, store); err != nil {
		return err
	}
	psapiRouter := psapi.MakeDefaultRouter(store)

	r := chi.NewRouter()
	r.Mount("/", psapiRouter)


	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return server.ListenAndServe()
```

## Contact

If you have any questions, suggestions, or feedback, please feel free to reach out to us:

- **Email:** charre.raphael@gmail.com
- **GitHub Issues:** [rcharre/psapi Issues](https://github.com/rcharre/psapi/issues)

We value your feedback and are always happy to assist with any issues or questions.

---

## License

This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for more details.

---

## Future Plans
You can find the what's next for this project in the [Roadmap](./ROADMAP.md) section of the repository.

We are always working on improving the project. If you have any ideas or suggestions, please let us know!

Stay tuned for future updates and enhancements!
