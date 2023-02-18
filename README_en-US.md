[中文简体](README.md) | English

[![Build Status](https://github.com/axetroy/chatgpt/workflows/ci/badge.svg)](https://github.com/axetroy/chatgpt/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/chatgpt)](https://goreportcard.com/report/github.com/axetroy/chatgpt)
![Latest Version](https://img.shields.io/github/v/release/axetroy/chatgpt.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/chatgpt.svg)

## open.ai SDK

Open.ai's unofficial Golang SDK and chatGPT's command line tool

### Install

1. Shell (Mac/Linux)

```bash
curl -fsSL https://github.com/release-lab/install/raw/v1/install.sh | bash -s -- -r=axetroy/chatgpt
```

2. PowerShell (Windows):

```bash
$r="axetroy/chatgpt";iwr https://github.com/release-lab/install/raw/v1/install.ps1 -useb | iex
```

3. [Github release page](https://github.com/axetroy/chatgpt/releases) (All platforms)

> download the executable file and put the executable file to `$PATH`

4. Build and install from source using [Golang](https://golang.org) (All platforms)

```bash
go install github.com/axetroy/chatgpt/cmd/chatgpt
```

5. Install via npm

```sh
npm install @axetroy/chatgpt -g
```

### Test

```bash
$ make test
```

### License

The [Anti-996 License](LICENSE)
