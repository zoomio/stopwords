# Changelog

## v0.9.0

- added `WordsSlice` option to accept a slice of strings;
- added `Text` option, which accepts two strings, first is the string containing words divided by some separator and second argument is `sep`, to specify which separator devides the words.

## v0.7.0

- added domain names option to the stop-words, enabled via `WithDomains`.

## v0.6.0

- BREAKING: from now on stopwords require `Setup()` in order to create an instance and use it.

## v0.5.0

- added Chinese, Hindi, Spanish, Arabic, Japanese, German, Hebrew, French and Korean stop-words.

## v0.4.0

- added Russian stop-words.

## v0.3.0

- bumped Go to 1.13;
- avoid duplicates in the slice of stop-words.

## v0.2.0

- moved onto Go modules;
- `stopwords#registerStopWords` made private.

## v0.1.0

- first release, and it will have stop-words registered.