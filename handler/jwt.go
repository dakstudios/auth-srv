package handler

import (
	"github.com/micro/go-micro/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	"github.com/dakstudios/auth-srv/db"
	jwtAuth "github.com/dakstudios/auth-srv/proto/jwt"
)

type JWT struct{}

func (j *JWT) Authenticate(ctx context.Context, req *jwtAuth.AuthenticateRequest, res *jwtAuth.AuthenticateResponse) error {
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

	res.Token = &jwtAuth.Token{Token: "some_jwt_token"}

	return nil
}

func (j *JWT) Authorize(ctx context.Context, req *jwtAuth.AuthorizeRequest, res *jwtAuth.AuthorizeResponse) error {
	if len(req.Token.Token) == 0 || len(req.Permission) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth", "invalid_request")
	}

	// check token and get user id from it

	authorized, err := db.Authorize("some_id", req.Permission)
	if err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth", "server_error")
	}

	res.Authorized = authorized

	return nil
}
