# Stop Words

[![Build Status](https://travis-ci.org/zoomio/stopwords.svg?branch=master)](https://travis-ci.org/zoomio/stopwords)
[![Go Report Card](https://goreportcard.com/badge/github.com/zoomio/stopwords)](https://goreportcard.com/report/github.com/zoomio/stopwords)
[![Coverage](https://codecov.io/gh/zoomio/tagify/branch/master/graph/badge.svg)](https://codecov.io/gh/zoomio/stopwords)
[![GoDoc](https://godoc.org/github.com/zoomio/tagify?status.svg)](https://godoc.org/github.com/zoomio/stopwords)

List of stop-words for convenient consumption in Go projects.

## Usage

* import `github.com/zoomio/stopwords` and it will have stop-words registered;
* use `stopwords#IsStopWord` to check whether given string is a stop-word;
* use `stopwords#Slice` to retrieve a list of registered stop-words.

## Changelog

See [CHANGELOG.md](https://raw.githubusercontent.com/zoomio/stopwords/master/CHANGELOG.md)

## Contributing

See [CONTRIBUTING.md](https://raw.githubusercontent.com/zoomio/stopwords/master/CONTRIBUTING.md)

## License

Released under the [Apache License 2.0](https://raw.githubusercontent.com/zoomio/stopwords/master/LICENSE).
