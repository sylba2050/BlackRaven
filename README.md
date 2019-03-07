# BlackRaven

BlackRaven is simple Golang server for generate CHUUNI word.

## demo
### server start
```bash
$ env PASSWORD=YOUR_PASSWORD go run .
```

### get string response
request
```bash
curl -X GET http://153.126.139.150:1417/api/default
```
response
```text
夢幻の騎士
```

### get JSON response
request
```bash
curl -X GET http://153.126.139.150:1417/api/json
```
response
```text
{"data":"夢幻の騎士"}
```

### create new data
sub db
```bash
$ curl -X POST http://153.126.139.150:1417/api/create/top -H 'Content-Type: application/json' -d '{"data":"hoge"}'
$ curl -X POST http://153.126.139.150:1417/api/create/under -H 'Content-Type: application/json' -d '{"data":"hoge"}'
```

main db
```bash
$ curl -X POST http://153.126.139.150:1417/api/create/top/admin -H 'Content-Type: application/json' -d '{"data":"hoge","pass":"YOUR_PASSWORD"}'
$ curl -X POST http://153.126.139.150:1417/api/create/under/admin -H 'Content-Type: application/json' -d '{"data":"hoge","pass":"YOUR_PASSWORD"}'
```

### create

## License
Copyright (c) 2019 Mstn
Released under the [MIT License](https://opensource.org/licenses/mit-license.php)
