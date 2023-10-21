## url-probe
### that the tool is guarding or watching over URLs to ensure their availability and performance

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###


- a simple command-line tool in golang which is checking url health and performance 
- provides a command-line interface for checking url health and performance

## project structure:

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

## installation

```shell
git clone https://github.com/kenjitheman/url-probe
```

## usage

```
go build
```

- using command line arguments:

```sh
./url-probe https://example.com https://example2.com
```

- using csv as a source:

```sh
./url-probe -source=csv -file=urls.csv
```

- using json as a source:

```sh
./url-probe -source=json -file=urls.json
```

- using txt as a source:

```sh
./url-probe -source=txt -file=urls.txt
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

- please make sure to update tests as appropriate

## license

- [MIT](https://choosealicense.com/licenses/mit/)
