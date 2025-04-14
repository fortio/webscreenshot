# webscreenshot

[![Go Report Card](https://goreportcard.com/badge/github.com/fortio/webscreenshot)](https://goreportcard.com/report/github.com/fortio/webscreenshot)
[![GitHub Release](https://img.shields.io/github/release/fortio/webscreenshot.svg?style=flat)](https://github.com/fortio/webscreenshot/releases/)
[![govulncheck](https://img.shields.io/badge/govulncheck-No%20vulnerabilities-success)](https://github.com/fortio/webscreenshot/actions/workflows/gochecks.yml)
[![golangci-lint](https://img.shields.io/badge/golangci%20lint-No%20issue-success)](https://github.com/fortio/webscreenshot/actions/workflows/gochecks.yml)

Take a (headless) screenshot of a Web page after giving it a delay for rendering

## Install

using golang 1.21+

```shell
go install github.com/fortio/webscreenshot@latest
```

You can also download one of the many binary [releases](https://github.com/fortio/webscreenshot/releases)

We publish a multi architecture docker image (linux/amd64, linux/arm64) `docker run fortio/webscreenshot`

Or brew custom tap `brew install fortio/tap/webscreenshot`

## Usage
```bash
webscreenshot [-delay duration -width w -height h] URL > screenshot.png
```
(`webscreenshot help` for all the options)

e.g
```bash
webscreenshot -delay 5s \
    "https://demo.fortio.org/browse?url=qps_max-s1_to_s2-0.7.1-2018-04-05-22-04.json" \
    > screenshot.png
open screenshot.png
```
