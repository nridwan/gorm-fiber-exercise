package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetFiberJwtClaims(c *fiber.Ctx) jwt.MapClaims {
	return c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
}

func GetFiberJwtUserId(c *fiber.Ctx) (string, error) {
	return GetFiberJwtClaims(c).GetSubject()
}
