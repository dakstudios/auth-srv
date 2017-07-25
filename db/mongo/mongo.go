package mongo

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dakstudios/auth-srv/db"
	account "github.com/dakstudios/auth-srv/proto/account"
)

type user struct {
	ID      bson.ObjectId   `bson:"_id,omitempty"`
	RoleIDs []bson.ObjectId `bson:"role_ids,omitempty"`

	Email        string `bson:"email,omitempty"`
	PasswordHash string `bson:"password,omitempty"`

	FirstName string `bson:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty"`

	Created time.Time `bson:"created,omitempty"`
	Updated time.Time `bson:"updated,omitempty"`
}

type role struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Permissions []string      `bson:"permissions,omitempty"`

	Name string `bson:"name,omitempty"`
}

type permission struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"name,omitempty"`
}

type mongo struct {
	db *mgo.Database
}

var (
	Url = "127.0.0.1"
)

const (
	hashCost int = 16
)

func init() {
	db.Register(new(mongo))
}

func (m *mongo) Init() error {
	var session *mgo.Session
	var err error

	if session, err = mgo.Dial(Url); err != nil {
		return err
	}

	session.SetMode(mgo.Monotonic, true)
	m.db = session.DB("account")

	// create basic roles
	id := bson.NewObjectId()
	name := "superadmin"
	m.db.C("roles").Upsert(
		bson.M{"name": name},
		bson.M{"$set": &role{
			ID:          id,
			Permissions: []string{"org.dakstudios.*"},
			Name:        name,
		}},
	)

	// create superadmin user
	ids := []bson.ObjectId{id}
	id = bson.NewObjectId()
	email := "admin@admin.com"

	var pass []byte
	if pass, err = bcrypt.GenerateFromPassword([]byte("admin"), hashCost); err != nil {
		return err
	}

	m.db.C("accounts").Upsert(
		bson.M{"email": email},
		bson.M{"$set": &user{
			ID:           id,
			RoleIDs:      ids,
			Email:        email,
			PasswordHash: string(pass),
			FirstName:    "Admin",
			LastName:     "Admin",
			Created:      time.Now(),
			Updated:      time.Unix(0, 0),
		}},
	)

	return nil
}

// account
func (m *mongo) FindUser(email string) (*account.User, error) {
	acc := user{}
	if err := m.db.C("accounts").Find(bson.M{"email": email}).One(&acc); err != nil {
		return nil, err
	}

	user := &account.User{
		Id:        acc.ID.String(),
		Roles:     []*account.Role{},
		FirstName: acc.FirstName,
		LastName:  acc.LastName,
		Email:     acc.Email,
		Password:  acc.PasswordHash,
		Created:   acc.Created.Unix(),
		Updated:   acc.Updated.Unix(),
	}

	roles := []role{}
	if err := m.db.C("roles").Find(bson.M{"_id": bson.M{"$in": acc.RoleIDs}}).All(&roles); err != nil {
		return nil, err
	}

	for _, role := range roles {
		r := &account.Role{
			Id:          role.ID.String(),
			Name:        role.Name,
			Permissions: role.Permissions,
		}

		user.Roles = append(user.Roles, r)
	}

	return user, nil
}

func (m *mongo) ReadUser(id string) (*account.User, error) {
	acc := user{}
	if err := m.db.C("accounts").FindId(id).One(&acc); err != nil {
		return nil, err
	}

	user := &account.User{
		Id:        acc.ID.String(),
		Roles:     []*account.Role{},
		FirstName: acc.FirstName,
		LastName:  acc.LastName,
		Email:     acc.Email,
		Password:  "",
		Created:   acc.Created.Unix(),
		Updated:   acc.Updated.Unix(),
	}

	var roles []role
	if err := m.db.C("roles").Find(bson.M{"_id": bson.M{"$in": acc.RoleIDs}}).All(&roles); err != nil {
		return nil, err
	}

	for _, role := range roles {
		r := &account.Role{
			Id:          role.ID.String(),
			Name:        role.Name,
			Permissions: role.Permissions,
		}

		user.Roles = append(user.Roles, r)
	}

	return user, nil
}

func (m *mongo) CreateUser(account *account.User) error {
	roleIds := []bson.ObjectId{}
	for _, role := range account.Roles {
		roleIds = append(roleIds, bson.ObjectId(role.Id))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), hashCost)
	if err != nil {
		return err
	}

	acc := user{
		ID:           bson.NewObjectId(),
		RoleIDs:      roleIds,
		Email:        account.Email,
		PasswordHash: string(hash),
		FirstName:    account.FirstName,
		LastName:     account.LastName,
		Created:      time.Now(),
		Updated:      time.Unix(0, 0),
	}

	return m.db.C("accounts").Insert(&acc)
}

func (m *mongo) UpdateUser(account *account.User) error {
	roleIds := []bson.ObjectId{}
	for _, role := range account.Roles {
		roleIds = append(roleIds, bson.ObjectId(role.Id))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), hashCost)
	if err != nil {
		return err
	}

	acc := user{
		RoleIDs:      roleIds,
		Email:        account.Email,
		PasswordHash: string(hash),
		FirstName:    account.FirstName,
		LastName:     account.LastName,
		Updated:      time.Now(),
	}

	return m.db.C("accounts").UpdateId(account.Id, &acc)
}

func (m *mongo) DeleteUser(id string) error {
	return m.db.C("accounts").RemoveId(id)
}

func (m *mongo) SearchUsers(limit, offset int64) ([]*account.User, error) {
	users := []user{}
	if err := m.db.C("accounts").Find(nil).Skip(int(offset)).Limit(int(limit)).All(&users); err != nil {
		return nil, err
	}

	accs := []*account.User{}
	for _, user := range users {
		roles := []*account.Role{}
		for _, roleId := range user.RoleIDs {
			r := role{}
			if err := m.db.C("roles").FindId(roleId).One(&r); err != nil {
				return nil, err
			}

			roles = append(roles, &account.Role{
				Id:          r.ID.String(),
				Name:        r.Name,
				Permissions: r.Permissions,
			})
		}

		accs = append(accs, &account.User{
			Id:        user.ID.String(),
			Roles:     roles,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Created:   user.Created.Unix(),
			Updated:   user.Updated.Unix(),
		})
	}

	return accs, nil
}

// role
func (m *mongo) ReadRole(id string) (*account.Role, error) {
	r := role{}
	if err := m.db.C("roles").FindId(id).One(&r); err != nil {
		return nil, err
	}

	role := account.Role{
		Id:          r.ID.String(),
		Name:        r.Name,
		Permissions: r.Permissions,
	}

	return &role, nil
}

func (m *mongo) ReadAllRoles() ([]*account.Role, error) {
	r := []role{}
	if err := m.db.C("roles").Find(nil).All(&r); err != nil {
		return nil, err
	}

	roles := []*account.Role{}
	for _, role := range r {
		roles = append(roles, &account.Role{
			Id:          role.ID.String(),
			Name:        role.Name,
			Permissions: role.Permissions,
		})
	}

	return roles, nil
}

func (m *mongo) CreateRole(_role *account.Role) error {
	r := role{
		ID:          bson.NewObjectId(),
		Name:        _role.Name,
		Permissions: _role.Permissions,
	}

	return m.db.C("roles").Insert(&r)
}

func (m *mongo) UpdateRole(_role *account.Role) error {
	r := role{
		Name:        _role.Name,
		Permissions: _role.Permissions,
	}

	return m.db.C("roles").UpdateId(_role.Id, &r)
}

func (m *mongo) DeleteRole(id string) error {
	return m.db.C("roles").RemoveId(id)
}

// auth
func (m *mongo) Authorize(id, permission string) (bool, error) {
	acc := user{}
	if err := m.db.C("accounts").FindId(id).One(&acc); err != nil {
		return false, err
	}

	// collect all user permissions
	iter := m.db.C("roles").Find(bson.M{"_id": bson.M{"$in": acc.RoleIDs}}).Iter()
	r := role{}
	perms := []string{}
	for iter.Next(&r) {
		perms = append(perms, r.Permissions...)
	}

	if strings.Contains(permission, ".") { // permission supports masks
		for _, perm := range perms {
			if string(perm[len(perm)-1]) == "*" { // supports * only as last character
				withoutDots := strings.Replace(perm, ".", "", -1)
				withoutMask := strings.Replace(withoutDots, "*", "", -1)
				permissionWithoutDots := strings.Replace(permission, ".", "", -1)

				if strings.HasPrefix(permissionWithoutDots, withoutMask) {
					return true, nil
				}
			} else if perm == permission {
				return true, nil
			}
		}

		return false, nil
	} else {
		return sliceContains(perms, permission), nil
	}
}

type filter func(s string) bool

func sliceFilter(fn filter, arr []string) []string {
	res := []string{}
	for _, s := range arr {
		if fn(s) {
			res = append(res, s)
		}
	}

	return res
}

func sliceContains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}

	return false
}
