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
Swagger:
go get -u github.com/swaggo/gin-swagger && go get -u github.com/swaggo/files && go get -u github.com/swaggo/swag/cmd/swag

# Curl Commands

## Local
http://localhost:8080

## Remote
https://badminton-matching-service.onrender.com

## ADMIN
### Macos/Linux
```bash
curl -X POST https://badminton-matching-service.onrender.com/admin -H 'authorization: Basic Zm9vOmJhcg==' -H 'content-type: application/json' -d '{"value":"bar"}'
```

### Windows
```bash
curl -X POST https://badminton-matching-service.onrender.com/admin -H "authorization: Basic Zm9vOmJhcg==" -H "content-type: application/json" -d "{\"value\":\"bar\"}"
```

## Device
```bash
curl -X GET https://badminton-matching-service.onrender.com/devices/1 -H "content-type: application/json"

curl -X POST https://badminton-matching-service.onrender.com/devices -H "Content-Type: application/json" -d '{"id":"1","last_event_id":"abc123"}'
```

## Event:
### Macos/Linux
```bash
  curl -X GET https://badminton-matching-service.onrender.com/events/event123 -H "Content-Type: application/json"

  curl -X POST https://badminton-matching-service.onrender.com/events -H "Content-Type: application/json" -d '{"event_id":"event123","event_date":"2023-01-01T12:00:00Z"}'
```

### Windows
```bash
  curl -X GET https://badminton-matching-service.onrender.com/events/event123 -H "Content-Type: application/json"

  curl -X POST https://badminton-matching-service.onrender.com/events -H "Content-Type: application/json" -d "{\"event_id\":\"event123\",\"event_date\":\"2023-01-01T12:00:00Z\"}"
```

## Player:
```bash
  curl -X GET https://badminton-matching-service.onrender.com/players/event123/john_doe -H "Content-Type: application/json"

  curl -X POST https://badminton-matching-service.onrender.com/players -H "Content-Type: application/json" -d '{"event_id":"event123","player_name":"john_doe"}'
```

## Match:
```bash
  curl -X GET https://badminton-matching-service.onrender.com/matches/event123/1/2023-01-01T14:00:00Z -H "Content-Type: application/json"

  curl -X POST https://badminton-matching-service.onrender.com/matches -H "Content-Type: application/json" -d '{"event_id":"event123","court_no":1,"date_time":"2023-01-01T14:00:00Z"}'
```

## Partner:
```bash
  curl -X GET https://badminton-matching-service.onrender.com/partners/event123/john_doe -H "Content-Type: application/json"

  curl -X POST https://badminton-matching-service.onrender.com/partners -H "Content-Type: application/json" -d '{"event_id":"event123","player_name":"john_doe"}'
```