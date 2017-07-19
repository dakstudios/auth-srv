package handler

import (
	"github.com/micro/go-micro/errors"

	"github.com/dakstudios/auth-srv/db"
	account "github.com/dakstudios/auth-srv/proto/account"
)

type Account struct{}

func validateUser(user *account.User, method string) error {
	if user == nil {
		return errors.BadRequest("org.dakstudios.srv.auth."+method, "invalid_account")
	}

	return nil
}

func (a *Account) ReadUser(context.Context, req *account.ReadUserRequest, res *account.ReadUserResponse) error {
	if req.Id == nil || len(req.Id) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.ReadUser", "invalid_id")
	}

	user, err := db.ReadUser(req.Id)
	if err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.ReadUser", err.Error())
	}

	res.Account = user

	return nil
}

func (a *Account) CreateUser(context.Context, req *account.CreateUserRequest, res *account.CreateUserResponse) error {
	if req.Account == nil {
		return errors.BadRequest("org.dakstudios.srv.auth.CreateUser", "invalid_account")
	}

	if len(req.Account.Email) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.CreateUser", "invalid_email")
	}

	if len(req.Account.Password) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.CreateUser", "invalid_password")
	}

	if err := db.CreateUser(req.Account); err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.CreateUser", err.Error())
	}

	return nil
}

func (a *Account) UpdateUser(context.Context, req *account.UpdateUserRequest, res *account.UpdateUserResponse) error {
	if req.Account == nil {
		return errors.BadRequest("org.dakstudios.srv.auth.UpdateUser", "invalid_account")
	}

	if len(req.Account.Id) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.UpdateUser", "invalid_id")
	}

	if len(req.Account.Email) > 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.UpdateUser", "cant_update_email")
	}

	if err := db.UpdateUser(req.Account); err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.UpdateUser", err.Error())
	}

	return nil
}

func (a *Account) DeleteUser(context.Context, req *account.DeleteUserRequest, res *account.DeleteUserResponse) error {
	if len(req.Id) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.DeleteUser", "invalid_id")
	}

	if err := db.DeleteUser(req.Id); err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.DeleteUser", err.Error())
	}

	return nil
}

func (a *Account) SearchUser(context.Context, *account.SearchUserRequest, *account.SearchUserResponse) error {

}

func (a *Account) ReadRole(context.Context, *ReadRoleRequest, *ReadRoleResponse) error {

}

func (a *Account) ReadAllRoles(context.Context, *ReadAllRolesRequest, *ReadAllRolesResponse) error {

}

func (a *Account) CreateRole(context.Context, *CreateRoleRequest, *CreateRoleResponse) error {

}

func (a *Account) UpdateRole(context.Context, *UpdateRoleRequest, *UpdateRoleResponse) error {

}

func (a *Account) DeleteRole(context.Context, *DeleteRoleRequest, *DeleteRoleResponse) error {

}
