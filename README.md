"# badminton-matching-service"

#Initialize boiler-plate code
go get -u github.com/gin-gonic/gin
go get gopkg.in/yaml.v3
go get -u go.uber.org/zap
go get github.com/redis/go-redis/v9

curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go

### Ref

Gin:
https://gin-gonic.com/en/docs/quickstart/
Redis:
https://github.com/redis/go-redis
Zap:
https://github.com/uber-go/zap

## ADMIN
Macos/Linux
```bash
curl -X POST http://localhost:8080/admin -H 'authorization: Basic Zm9vOmJhcg==' -H 'content-type: application/json' -d '{"value":"bar"}'
```

Windows
```bash
curl -X POST http://localhost:8080/admin -H "authorization: Basic Zm9vOmJhcg==" -H "content-type: application/json" -d "{\"value\":\"bar\"}"
```

## Device
```bash
curl -X GET http://localhost:8080/device/1234 -H "content-type: application/json"

curl -X POST http://localhost:8080/device -H "content-type: application/json" -d "{\"id\":\"TEST-BALL\"}"
```