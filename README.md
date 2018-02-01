# MS Office Online preview link generator

[![Build Status](https://travis-ci.org/ildarusmanov/msofficepreview.svg?branch=master)](https://travis-ci.org/ildarusmanov/msofficepreview)

# Setup

## Create config from config.example.com
```
cp config.example.yml config.yml
vim config.yml
```
## Setup dependencies
```
dep ensure
```

## Run tests


```
go test ./...
```

## Try some requests like this
```
POST http://0.0.0.0:8080/api/v1/previews

{"file_path": "/path/to/file/in/your/storage/example.xls"}
```
