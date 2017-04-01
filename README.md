# go-server-stats

[![Go Report Card](https://goreportcard.com/badge/github.com/andrewlader/go-server-stats)](https://goreportcard.com/report/github.com/andrewlader/go-server-stats)
[![Build Status](https://travis-ci.org/AndrewLader/go-server-stats.svg?branch=master)](https://travis-ci.org/AndrewLader/go-server-stats)
[![Coverage Status](https://coveralls.io/repos/github/andrewlader/go-server-stats/badge.svg?branch=master)](https://coveralls.io/github/andrewlader/go-server-stats)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/AndrewLader/go-server-stats/blob/master/LICENSE)

The Go-Server-Stats package can be used to count statistics for a web service written in Go. 

```
┌─┐┌─┐   ┌─┐┌─┐┬─┐┬  ┬┌─┐┬─┐   ┌─┐┌┬┐┌─┐┌┬┐┌─┐
│ ┬│ │───└─┐├┤ ├┬┘└┐┌┘├┤ ├┬┘───└─┐ │ ├─┤ │ └─┐
└─┘└─┘   └─┘└─┘┴└─ └┘ └─┘┴└─   └─┘ ┴ ┴ ┴ ┴ └─┘
```

### Installation
To install the package, use the `go` command
```
> go get github.com/andrewlader/go-server-stats
```

### Requirements
Requires Go >= v1.2

### Usage
Add `stats.Stats` to the web service. For example:
```
// the HTTP handler instance
type httpServer struct {
	stats   *stats.Stats
	mux     map[string]apiHandler
}
```

Then when an API is called, use the `Update()` method to increment the counters:
```
server.stats.Update(wasSuccessful, uint64(request.ContentLength), numberOfBytesWritten)
```
