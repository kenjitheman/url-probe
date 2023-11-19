## url-probe

### That the tool is guarding or watching over URLs to ensure their availability and performance.

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

- A simple command-line tool in golang which is checking url health and performance.
- Provides a command-line interface for checking url health and performance.

## Project structure:

```go
.
├── core
│   ├── args.go
│   ├── core.go
│   └── reader.go
├── go.mod
├── go.sum
├── main.go
├── README.md
└── tests
    ├── args_test.go
    ├── core_test.go
    ├── probe_test.go
    └── reader_test.go
```

## Installation

```shell
git clone https://github.com/kenjitheman/url-probe
```

## Usage

```shell
go build
```

- Using command line arguments:

```shell
./url-probe https://example.com https://example2.com
```

- Using csv as a source:

```shell
./url-probe -source=csv -file=urls.csv
```

- Using json as a source:

```shell
./url-probe -source=json -file=urls.json
```

- Using txt as a source:

```shell
./url-probe -source=txt -file=urls.txt
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

- Please make sure to update tests as appropriate.

## License

- [MIT](https://choosealicense.com/licenses/mit/)
