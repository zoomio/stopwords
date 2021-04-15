# Contributing

## Guidelines for pull requests

- Write tests for any changes.
- Separate unrelated changes into multiple pull requests.
- For bigger changes, make sure you start a discussion first by creating an issue and explaining the intended change.
- Ensure the build is green before you open your PR. The Pipelines build won't run by default on a remote branch, so enable Pipelines.

## Release

* All notable changes comming with the new version should be documented in [CHANGELOG.md](https://raw.githubusercontent.com/zoomio/stopwords/master/CHANGELOG.md).
* Use `./_bin/test.sh` to run tests. 
* Use `./_bin/tag.sh <x.y.z>` to tag, push and trigger new release. 
