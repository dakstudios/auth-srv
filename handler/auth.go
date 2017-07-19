package handler

import (
	"github.com/micro/go-micro/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	"github.com/dgrijalva/jwt-go"

	"github.com/dakstudios/auth-srv/db"
	auth "github.com/dakstudios/auth-srv/proto/auth"
)

const (
	jwtSecret = "some_secure_secret"
)

type userClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type Auth struct{}

func (a *Auth) Authenticate(ctx context.Context, req *auth.AuthenticateRequest, res *auth.AuthenticateResponse) error {
	if len(req.Email) == 0 || len(req.Password) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth", "invalid_request")
	}

	user, err := db.FindUser(req.Email)
	if err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth", "server_error")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.BadRequest("org.dakstudios.srv.auth", "access_denied")
	}

	claims := userClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "org.dakstudios.srv.auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth", "server_error")
	}

	res.Token = &auth.Token{Token: signed}

	return nil
}

func (a *Auth) Authorize(ctx context.Context, req *auth.AuthorizeRequest, res *auth.AuthorizeResponse) error {
	if len(req.Token.Token) == 0 || len(req.Permission) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth", "invalid_request")
	}

	token, err := jwt.ParseWithClaims(req.Token.Token, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if !token.Valid {
		return errors.BadRequest("org.dakstudios.srv.auth", "access_denied")
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return errors.InternalServerError("org.dakstudios.srv.auth", "server_error")
	}

	authorized, err := db.Authorize(claims.ID, req.Permission)
	if err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth", "server_error")
	}

	res.Authorized = authorized

	return nil
}
