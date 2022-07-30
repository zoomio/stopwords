# Stop Words

[![Go Report Card](https://goreportcard.com/badge/github.com/zoomio/stopwords)](https://goreportcard.com/report/github.com/zoomio/stopwords)
[![Coverage](https://codecov.io/gh/zoomio/stopwords/branch/master/graph/badge.svg)](https://codecov.io/gh/zoomio/stopwords)
[![GoDoc](https://godoc.org/github.com/zoomio/stopwords?status.svg)](https://godoc.org/github.com/zoomio/stopwords)

List of stop-words for convenient consumption in Go projects. 

Supported languages:

* English - [StopWordsEN](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_en.go)
* Russian - [StopWordsRu](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_ru.go)
* Chinese - [StopWordsZh](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_zh.go)
* Hindi - [StopWordsHi](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_hi.go)
* Spanish - [StopWordsEs](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_es.go)
* Arabic - [StopWordsAr](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_ar.go)
* Japanese - [StopWordsJa](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_ja.go)
* German - [StopWordsDe](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_de.go)
* Hebrew - [StopWordsHe](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_he.go)
* French - [StopWordsFr](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_fr.go)
* Korean - [StopWordsKo](https://raw.githubusercontent.com/zoomio/stopwords/master/stopwords_ko.go)

## Usage

* import `github.com/zoomio/stopwords` and it will have stop-words registered;
* use `stopwords.Setup().IsStopWord("a")` to check whether given string is a stop-word;
* use `stopwords.Setup().Slice()` to retrieve a list of registered stop-words.

## Changelog

See [CHANGELOG.md](https://raw.githubusercontent.com/zoomio/stopwords/master/CHANGELOG.md)

## Contributing

See [CONTRIBUTING.md](https://raw.githubusercontent.com/zoomio/stopwords/master/CONTRIBUTING.md)

## License

Released under the [Apache License 2.0](https://raw.githubusercontent.com/zoomio/stopwords/master/LICENSE).
