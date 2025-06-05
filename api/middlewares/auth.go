package middlewares

import (
	"os"
	"strconv"
	"strings"

	"github.com/faridEmilio/api_go_viajate_corporativo/pkg/domains/auth"
	filtros "github.com/faridEmilio/api_go_viajate_corporativo/pkg/filtros/usuarios"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type MiddlewareManager struct {
	authService auth.AuthService
}

func NewMiddlewareManager(auth auth.AuthService) MiddlewareManager {
	return MiddlewareManager{
		authService: auth,
	}
}

func (m *MiddlewareManager) ValidarPermiso(scope string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		bearer := c.Get("Authorization")
		if len(bearer) < 1 {
			return fiber.NewError(fiber.StatusUnauthorized, "acceso no autorizado, debe enviar un token de autenticación")
		}

		parts := strings.Split(bearer, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "formato de token de autenticación inválido")
		}

		tokenString := parts[1]

		// Parsear el token con MapClaims
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "Token de autenticación inválido")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Claims inválidos en el token")
		}

		// Obtener el ID del usuario desde "sub"
		idStr, ok := claims["sub"].(string)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "ID de usuario no válido en el token")
		}
		userID, err := strconv.Atoi(idStr)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "ID de usuario no es numérico")
		}

		// Validar que el permiso esté incluido en el token
		rawPerms, ok := claims["permisos"].([]interface{})
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Permisos no definidos en el token")
		}

		permMap := make(map[string]struct{}, len(rawPerms))
		for _, p := range rawPerms {
			if permStr, ok := p.(string); ok {
				permMap[permStr] = struct{}{}
			}
		}

		if _, hasPerm := permMap[scope]; !hasPerm {
			return fiber.NewError(fiber.StatusForbidden, "No tienes permiso para esta operación")
		}

		user, err := m.authService.GetUserService(filtros.UsuarioFiltro{ID: uint(userID)})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error al obtener detalles del usuario")
		}

		// Verifica si el parámetro comunidad_id está presente
		if comunidadID := c.Params("comunidad_id"); comunidadID != "" {
			comunidad_id, err := strconv.Atoi(comunidadID)
			if err != nil || comunidad_id <= 0 {
				return fiber.NewError(fiber.StatusBadRequest, "ID de comunidad inválido")
			}

			comunidadMap := ExtraerComunidadesMap(claims)
			if _, found := comunidadMap[comunidad_id]; !found {
				return fiber.NewError(fiber.StatusForbidden, "No perteneces a esta comunidad")
			}

			c.Locals("comunidadID", comunidad_id)
		}
		// Establecer los detalles del usuario en el contexto para uso posterior
		c.Locals("user", user)
		return c.Next()
	}
}

func (m *MiddlewareManager) ValidateToken() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		bearer := c.Get("Authorization")
		if len(bearer) < 1 {
			return fiber.NewError(fiber.StatusUnauthorized, "acceso no autorizado, debe enviar un token de autenticación")
		}

		parts := strings.Split(bearer, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "formato de token de autenticación inválido")
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "token de autenticación inválido")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "claims inválidos en el token")
		}

		idStr, ok := claims["sub"].(string)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "ID de usuario no válido en el token")
		}

		userID, err := strconv.Atoi(idStr)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "ID de usuario no es numérico")
		}

		user, err := m.authService.GetUserService(filtros.UsuarioFiltro{ID: uint(userID)})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error al obtener detalles del usuario")
		}

		c.Locals("user", user)
		return c.Next()
	}
}

func ExtraerComunidadesMap(claims jwt.MapClaims) map[int]struct{} {
	result := make(map[int]struct{})
	userRaw, ok := claims["user"].(map[string]interface{})
	if !ok {
		return result
	}

	comunidadesRaw, ok := userRaw["comunidades"].([]interface{})
	if !ok {
		return result
	}

	for _, raw := range comunidadesRaw {
		if comunidadMap, ok := raw.(map[string]interface{}); ok {
			if idFloat, ok := comunidadMap["id"].(float64); ok {
				result[int(idFloat)] = struct{}{}
			}
		}
	}
	return result
}
