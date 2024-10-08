# Gorm Fiber Exercise

Exercise to implement real world API implementation using go fiber.
Some feature will be preserved in my future go boilerplate.

Features:

- Graceful shutdown
- Dependency Injection by FX
- .env based config
- Hot Reload
- Gorm database model
- JWT auth

## Requirement

- Golang
- Postgresql

## How to run

1. install go
2. install toolset:

   - air-verse/air: `go install github.com/air-verse/air@latest`
   - python3: use any method at disposal, but make sure `python` command is linked to python3, since air_build.py need it

3. setup `.env`, see `.env.example`, self explanatory enough to be copied to `.env`
4. run `go mod download`
5. run `air`

### Extras
see [feat/extras](https://github.com/nridwan/gorm-fiber-exercise/tree/feat/extras) for more feature like:
- manual migration with atlas
- opentelemetry integration

Update: these extra features are now refined in https://github.com/nridwan/gorm-fiber-boilerplate . since it's quite redundant this repo won't get maintained
