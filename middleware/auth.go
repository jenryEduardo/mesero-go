package middleware

// import (
// 	"mesero-go/part-1/api-main/users/utils" // Importa el paquete utils
// 	"net/http"
// 	"strings"
// 	"github.com/gin-gonic/gin"
// )

// // AuthRequired es el middleware que valida el JWT
// func AuthRequired(c *gin.Context) {
// 	// Obtener el token del encabezado "Authorization"
// 	authHeader := c.GetHeader("Authorization")
// 	if authHeader == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
// 		c.Abort()
// 		return
// 	}

// 	// Eliminar el prefijo "Bearer " del encabezado
// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// 	// Validar el token JWT
// 	token, err := utils.ValidateJWT(tokenString)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
// 		c.Abort()
// 		return
// 	}

// 	// Extraer el user_id del token (usando el token validado)
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
// 		c.Abort()
// 		return
// 	}

// 	userID, ok := claims["user_id"].(float64) // El "user_id" está guardado en el token como float64
// 	if !ok {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID no encontrado en el token"})
// 		c.Abort()
// 		return
// 	}

// 	// Agregar el user_id al contexto para usarlo más tarde en los handlers
// 	c.Set("idUsuario", int(userID)) // Convertir el float64 a int, que es el tipo de tu ID

// 	// Si el token es válido, continúa con la ejecución de la solicitud
// 	c.Next()
// }
