# goweb
## API
### GET
- hostname/
- hostname/users
- hostname/users/{id}
- hostname/groups
- hostname/groups/{id}
- hostname/messages

### PUT
- hostname/users
```
curl -H "Content-Type: application/json" -d '{"name":"keisuke"}' http://localhost:8080/users
```

- hostname/groups
```
curl -H "Content-Type: application/json" -d '{"name":"tennis_club"}' http://localhost:8080/groups
```

- hostname/messages
```
curl -H "Content-Type: application/json" -d '{"UserId":2, "GroupId":1, "text":"hhhhhhhhhh", "mode":"dialogue"}' http://localhost:8080/messages
```

