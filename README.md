# Backoff

[![Test](https://github.com/aofei/backoff/actions/workflows/test.yaml/badge.svg)](https://github.com/aofei/backoff/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/aofei/backoff/branch/master/graph/badge.svg)](https://codecov.io/gh/aofei/backoff)
[![Go Report Card](https://goreportcard.com/badge/github.com/aofei/backoff)](https://goreportcard.com/report/github.com/aofei/backoff)
[![Go Reference](https://pkg.go.dev/badge/github.com/aofei/backoff.svg)](https://pkg.go.dev/github.com/aofei/backoff)

A Full-Jitter exponential backoff helper for Go.

The algorithm used to compute the randomized delay mainly comes from the Full-Jitter exponential backoff strategy
described in the AWS Architecture Blog post
[Exponential Backoff and Jitter](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/), ensuring
evenly distributed retries without synchronized bursts.

## Features

- Extremely easy to use
- Zero third-party dependencies

## Installation

To use this project programmatically, `go get` it:

```bash
go get github.com/aofei/backoff
```

## Community

If you have any questions or ideas about this project, feel free to discuss them
[here](https://github.com/aofei/backoff/discussions).

## Contributing

If you would like to contribute to this project, please submit issues [here](https://github.com/aofei/backoff/issues)
or pull requests [here](https://github.com/aofei/backoff/pulls).

When submitting a pull request, please make sure its commit messages adhere to
[Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/).

## License

This project is licensed under the [MIT License](LICENSE).
