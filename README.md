# go-open

`go-open` provides a unified API to stream files from URI (file, HTTP, and HTTPS) in Go.

`go-open` currently supports:

- Text or binary files
- Streaming from the local filesystem
    - relative paths
    - absolute paths (with or without `file://`)
- Streaming from the Internet
    - `http://` or `https://` URLs


In the future...

- Windows support?
- S3 extension


## Get started

To install the `go-open` library for your codebase, simply run:

```console
go get -u github.com/w-gao/go-open@latest
```

## Usage

Below is some boilerplate code to get started with `go-open`.

```go
func processFile(uri string) error {
	// f is an io.ReadCloser
	f, err := SmartOpen(uri)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	// Make sure to close the file!
	defer f.Close()

	// You can do anything with f
	// For example, to iterate through the lines, you can do:
	var line string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line = sc.Text()
		fmt.Printf("line: %s\n", line)
	}

	if err := sc.Err(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
```

```go
// relative paths
processFile("README.md")

// absolute paths
processFile("/Users/wlgao/code/go-open/README.md")

// file from the Internet
processFile("https://raw.githubusercontent.com/w-gao/wdl-viewer/main/index.html")
```


## LICENSE

MIT License. Copyright (c) 2022 William Gao
