### Basic JWT Implementations
A basic api that accepts username and password authenticates it with static values and returns a JWT token

Server

```bash

chmod +x generate_key.sh
# generate the jwt signing key
./generate_key.sh 
 
# start go server
go run ./src
```

Client
```bash
curl -X POST http://localhost:8000/login -d '{"email":"buddha@gautam.com","password":"Bhandinaa"}' -H "Content-Type: application/json"
```


