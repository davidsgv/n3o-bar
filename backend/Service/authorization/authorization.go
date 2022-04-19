package authorization

import (
	"fmt"
	"os"
	"time"

	"github.com/davidsgv/n3o-bar/dto"
	"github.com/davidsgv/n3o-bar/entity"
	"github.com/davidsgv/n3o-bar/model/authorization"
	"github.com/dgrijalva/jwt-go"
)

//interfaz para uso externo
type AuthorizationService interface {
	GenerateToken(dto.UsuarioInicioSesion) string
	ValidateToken(tokenString string) (*jwt.Token, error)
	Login(dto.Credentials) (*dto.UsuarioInicioSesion, bool)
}

//Claims para firma jwt
type jwtCustomClaims struct {
	Usuario string `json:"usuario"`
	Id      int64  `json:"usuarioId"`
	Rol     string `json:"rol"`
	jwt.StandardClaims
}

//implementa la interfaz para devolver este objeto
type jwtService struct {
	secretKey string
	issuer    string
}

//contructor
func NewAuthorizationService() AuthorizationService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "neobar.com",
	}
}

//Implementación interfaz (publicos)
func (jwtSrv jwtService) Login(dtoEntrada dto.Credentials) (*dto.UsuarioInicioSesion, bool) {
	//se crea el modelo y la entidad con el dto entrante
	modelAutho := authorization.NewAuthorizationModel()
	entity := entity.Usuario{
		Correo:   dtoEntrada.Username,
		Password: dtoEntrada.Password,
	}

	//se trata de iniciar sesion
	usuario := modelAutho.Login(entity)

	//se valida si se logro iniciar sesión
	if usuario.Correo != "" && usuario.Id > 0 && usuario.Rol.Id > 0 && usuario.Rol.Nombre != "" {
		dtoRespuesta := dto.UsuarioInicioSesion{
			Correo:    usuario.Correo,
			Rol:       usuario.Rol.Nombre,
			UsuarioId: usuario.Id,
		}
		return &dtoRespuesta, true
	}

	return nil, false
}
func (jwtSrv jwtService) GenerateToken(dto dto.UsuarioInicioSesion) string {

	// Set custom and standard claims
	claims := &jwtCustomClaims{
		Usuario: dto.Correo,
		Id:      dto.UsuarioId,
		Rol:     dto.Rol,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
func (jwtSrv jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
}

//funciones privadas
func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}
