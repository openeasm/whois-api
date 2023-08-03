# A API for whois query
## Install
build from source
```bash
git clone https://github.com/openeasm/whois_api
go run app.go 8000
```
or use docker
```bash
docker build -t my-golang-app .
docker run -p 8000:8000 my-golang-app
```

```bash
## Usage
### 1. 查域名
```bash
curl http://localhost:8000/whois?item=qianxin.com
```
### 2. 查 IP
暂不支持

This is a part of [OpenEASM]() project.