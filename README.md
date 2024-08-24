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

1. install go
2. install toolset:

   - air-verse/air: `go install github.com/air-verse/air@latest`
   - python3: use any method at disposal, but make sure `python` command is linked to python3, since air_build.py need it

3. setup `.env`, see `.env.example`, self explanatory enough to be copied to `.env`
4. run `go mod download`
5. run `air`
