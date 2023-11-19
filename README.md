## Installation

```sh
git clone https://github.com/kenjitheman/url-probe
```

## Usage

```sh
go build
```

- Using command line arguments:

```sh
./url-probe https://example.com https://example2.com
```

- Using csv as a source:

```sh
./url-probe -source=csv -file=urls.csv
```

- Using json as a source:

```sh
./url-probe -source=json -file=urls.json
```

- Using txt as a source:

```sh
./url-probe -source=txt -file=urls.txt
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

- Please make sure to update tests as appropriate.

## License

- [MIT](https://choosealicense.com/licenses/mit/)
