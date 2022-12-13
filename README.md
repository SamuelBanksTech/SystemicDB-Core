<div align="center" style="background: black">
<img height="300" src="./res/logo-black.png" alt="SystemicDB Logo" />

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/SamuelBanksTech/SystemicDB-Core)](https://goreportcard.com/report/github.com/SamuelBanksTech/SystemicDB-Core)
[![Twitter Handle](https://img.shields.io/twitter/follow/samuelbankstech)](https://twitter.com/samuelbankstech)

</div>


# SystemicDB - Core

### Introduction

This is the core SystemicDB package that is used in the full SystemicDB Server application. I can be imported and used in any Go application for standalone usage or one could wrap their own server around package for custom usage.

### Install

```bash 
go get github.com/SamuelBanksTech/SystemicDB-Core
```

### Usage

First of all you will need to instantiate the core SystemicDB struct.

```go
sdb := NewSystemicDB()
```

Insert, take a string key, a byte slice as its data, and a expiry time

```go
sdb.Insert("my-key", []byte("This is my byte data"), 15 * time.Minute)
```


