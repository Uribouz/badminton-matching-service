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


# Curl Commands

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
curl -X GET http://localhost:8080/devices/1 -H "content-type: application/json"

curl -X POST http://localhost:8080/devices -H "Content-Type: application/json" -d '{"id":"1","last_event_id":"abc123"}'
```

## Event:
```bash
  curl -X GET http://localhost:8080/events/event123 -H "Content-Type: application/json"

  curl -X POST http://localhost:8080/events -H "Content-Type: application/json" -d '{"event_id":"event123","event_date":"2023-01-01T12:00:00Z"}'
```

## Player:
```bash
  curl -X GET http://localhost:8080/players/event123/john_doe -H "Content-Type: application/json"

  curl -X POST http://localhost:8080/players -H "Content-Type: application/json" -d '{"event_id":"event123","player_name":"john_doe"}'
```

## Match:
```bash
  curl -X GET http://localhost:8080/matches/event123/1/2023-01-01T14:00:00Z -H "Content-Type: application/json"

  curl -X POST http://localhost:8080/matches -H "Content-Type: application/json" -d '{"event_id":"event123","court_no":1,"date_time":"2023-01-01T14:00:00Z"}'
```

## Partner:
```bash
  curl -X GET http://localhost:8080/partners/event123/john_doe -H "Content-Type: application/json"

  curl -X POST http://localhost:8080/partners -H "Content-Type: application/json" -d '{"event_id":"event123","player_name":"john_doe"}'
```