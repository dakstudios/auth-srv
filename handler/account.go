package handler

import (
	"github.com/micro/go-micro/errors"

	"github.com/dakstudios/auth-srv/db"
	account "github.com/dakstudios/auth-srv/proto/account"
)

type Account struct{}

func (a *Account) ReadUser(ctx context.Context, req *account.ReadUserRequest, res *account.ReadUserResponse) error {
	if len(req.Id) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.ReadUser", "invalid_id")
	}

	user, err := db.ReadUser(req.Id)
	if err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.ReadUser", err.Error())
	}

	res.Account = user

	return nil
}

func (a *Account) CreateUser(ctx context.Context, req *account.CreateUserRequest, res *account.CreateUserResponse) error {
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

func (a *Account) UpdateUser(ctx context.Context, req *account.UpdateUserRequest, res *account.UpdateUserResponse) error {
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

func (a *Account) DeleteUser(ctx context.Context, req *account.DeleteUserRequest, res *account.DeleteUserResponse) error {
	if len(req.Id) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.DeleteUser", "invalid_id")
	}

	if err := db.DeleteUser(req.Id); err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.DeleteUser", err.Error())
	}

	return nil
}

func (a *Account) SearchUser(ctx context.Context, req *account.SearchUserRequest, res *account.SearchUserResponse) error {
	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Offset < 0 {
		req.Offset = 0
	}

	accounts, err := db.SearchUsers(req.Limit, req.Offset)
	if err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.SearchUser", err.Error())
	}

	res.Accounts = accounts

	return nil
}

func (a *Account) ReadRole(ctx context.Context, req *account.ReadRoleRequest, res *account.ReadRoleResponse) error {
	if len(req.Id) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.ReadRole", "invalid_id")
	}

	role, err := db.ReadRole(req.Id)
	if err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.ReadRole", err.Error())
	}

	res.Role = role

	return nil
}

func (a *Account) ReadAllRoles(ctx context.Context, req *account.ReadAllRolesRequest, res *account.ReadAllRolesResponse) error {
	roles, err := db.ReadAllRoles()
	if err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.ReadAllRoles", err.Error())
	}

	res.Roles = roles

	return nil
}

func (a *Account) CreateRole(ctx context.Context, req *account.CreateRoleRequest, res *account.CreateRoleResponse) error {
	if req.Role == nil {
		return errors.BadRequest("org.dakstudios.srv.auth.CreateRole", "invalid_role")
	}

	if len(req.Role.Name) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.CreateRole", "invalid_name")
	}

	if err := db.CreateRole(req.Role); err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.CreateRole", err.Error())
	}

	return nil
}

func (a *Account) UpdateRole(ctx context.Context, req *account.UpdateRoleRequest, res *account.UpdateRoleResponse) error {
	if req.Role == nil {
		return errors.BadRequest("org.dakstudios.srv.auth.UpdateRole", "invalid_role")
	}

	if len(req.Role.Id) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.UpdateRole", "invalid_id")
	}

	if err := db.UpdateRole(req.Role); err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.UpdateRole", err.Error())
	}

	return nil
}

func (a *Account) DeleteRole(ctx context.Context, req *account.DeleteRoleRequest, res *account.DeleteRoleResponse) error {
	if len(req.Id) == 0 {
		return errors.BadRequest("org.dakstudios.srv.auth.DeleteRole", "invalid_id")
	}

	if err := db.DeleteRole(req.Id); err != nil {
		return errors.InternalServerError("org.dakstudios.srv.auth.DeleteRole", err.Error())
	}

	return nil
}
