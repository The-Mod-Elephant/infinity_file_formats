# Infinity File Formats Go
![](https://img.shields.io/badge/go-65A2BE2?logo=go&style=for-the-badge&logoColor=grey)
[![](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/dark0dave/infinity_file_formats/releases/latest)
[![](https://img.shields.io/badge/Windows-0078D6?&style=for-the-badge&logoColor=white&logo=git-for-windows)](https://github.com/dark0dave/infinity_file_formats/releases/latest)
[![](https://img.shields.io/badge/mac%20os-grey?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/dark0dave/infinity_file_formats/releases/latest)
[![](https://img.shields.io/github/actions/workflow/status/dark0dave/infinity_file_formats/main.yaml?style=for-the-badge)](https://github.com/dark0dave/infinity_file_formats/actions/workflows/main.yaml)
[![](https://img.shields.io/github/license/dark0dave/infinity_file_formats?style=for-the-badge)](./LICENSE)

A Golang library for parsing Infinity file formats into JSON.

## Description

The `infinity_file_formats` package provides functionality to parse files in various Infinity-related formats (such as `.are`, `.bam`, etc.) and convert them into JSON format. This library is designed to handle the unique structure of Infinity files and provide a flexible way to work with their data in modern applications.

This project is not an executable tool but rather a library that can be integrated into other Go projects for file parsing needs.

## Features

- Supports multiple Infinity-related file formats
- Efficient parsing of large files
- Converts parsed data to JSON format

## Usage

Here are some examples of how to use the `infinity_file_formats` library in your Golang projects.

### Basic Parsing Example

```go
package main

import (
	"fmt"
	"github.com/dark0dave/infinity_file_formats/bg"
)

func main() {
	filePath := "path/to/your/file.itm" // Path to your Infinity file
	itm, err := bg.OpenITM(filePath)
	if err != nil {
		fmt.Printf("Reading item failed")
		return
	}
	buf := new(bytes.Buffer)
	err = itm.WriteJson(buf)
	if err != nil {
		fmt.Printf("Writing item to JSON failed")
		return
	}
	fmt.Println(buf.String())
}
```

## Build

```sh
go build ./...
go doc bg
```

## Test

```sh
go test -v ./...
```
