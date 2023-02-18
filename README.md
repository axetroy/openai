中文简体 | [English](README_en-US.md)

[![Build Status](https://github.com/axetroy/chatgpt/workflows/ci/badge.svg)](https://github.com/axetroy/chatgpt/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/chatgpt)](https://goreportcard.com/report/github.com/axetroy/chatgpt)
![Latest Version](https://img.shields.io/github/v/release/axetroy/chatgpt.svg)
[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/chatgpt.svg)

## open.ai SDK

open.ai 的 Golang 非官方 SDK 以及 chatGPT 的命令行工具

### Usage

```bash
$ chatgpt --help
```

### 安装

1. Shell (Mac/Linux)

```bash
curl -fsSL https://github.com/release-lab/install/raw/v1/install.sh | bash -s -- -r=axetroy/chatgpt
```

2. PowerShell (Windows):

```bash
$r="axetroy/chatgpt";iwr https://github.com/release-lab/install/raw/v1/install.ps1 -useb | iex
```

3. [Github release page](https://github.com/axetroy/chatgpt/releases) (全平台支持))

> 下载可执行文件，并且把它加入到`$PATH` 环境变量中

4. 使用 [Golang](https://golang.org) 从源码中构建并安装 (全平台支持)

```bash
go install github.com/axetroy/chatgpt/cmd/chatgpt
```

5. 通过 npm 安装

```sh
npm install @axetroy/chatgpt -g
```

### 测试

```bash
$ make test
```

### 开源许可

The [Anti-996 License](LICENSE_zh-CN)
