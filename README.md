# MS Office Online preview link generator

[![Build Status](https://travis-ci.org/ildarusmanov/msofficepreview.svg?branch=master)](https://travis-ci.org/ildarusmanov/msofficepreview)
[![Maintainability](https://api.codeclimate.com/v1/badges/406e0ad48e6c05fd6a5e/maintainability)](https://codeclimate.com/github/ildarusmanov/msofficepreview/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/406e0ad48e6c05fd6a5e/test_coverage)](https://codeclimate.com/github/ildarusmanov/msofficepreview/test_coverage)

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

Get preview link
```
POST http://0.0.0.0:8001/api/v1/previews

{"file_path": "/path/to/file/in/your/storage/example.xls"}
```

Status check
```
GET http://0.0.0.0:8001/api/v1/status/check
```

## Run with docker
```
cd [project path]

sudo docker build -t msofficepreview .

// prod
sudo docker run --restart=always -d -p 10.90.137.73:8001:8001 --network host --mount type=bind,source=/storage/dir/path,target=/storage,readonly msofficepreview

// or dev
sudo docker run -d -p 8001:8001 --network host --mount type=bind,source=/home/storage,target=/storage,readonly msofficepreview 
// list containers
sudo docker ps
```


