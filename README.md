# Gorm Fiber Exercise

Exercise to implement real world API implementation using go fiber. 
Some feature will be preserved in my future go boilerplate.

Features:

- Graceful shutdown
- Dependency Injection (manual & Fx)
- .env based config
- Hot Reload
- Gorm database model
- JWT auth

## How to run
```
⚠️ This guide refer to Linux Users. after rechecking in windows, .air.toml command must be:
go build -o ./tmp/main.exe .

and bin should be:
./tmp/main.exe

Update it manually first if using windows, working on a way for air command to be cross platform
```
1. install go
2. install air-verse/air: `go install github.com/air-verse/air@latest`
3. setup `.env`, see `.env.example`, self explanatory enough to be copied to `.env`
4. run `go mod download`
5. run `air`
