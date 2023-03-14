package v2controllers

import (
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type User struct {
	ReadableBaseType
	auth_models.StoredControlledUserFields
	CreatableUser
}

type CreatableUser struct {
	EditableUser
}

type EditableUser struct {
	auth_models.StoredMutableUserFields
}

//nolint:unused
func (u User) toModel(_ *v2models.StoreSet) (v2models.User, error) {
	return v2models.User{
		Model: gorm.Model{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		StoredControlledUserFields: u.StoredControlledUserFields,
		StoredMutableUserFields:    u.StoredMutableUserFields,
	}, nil
}

//nolint:unused
func (u CreatableUser) toModel(storeSet *v2models.StoreSet) (v2models.User, error) {
	return User{CreatableUser: u}.toModel(storeSet)
}

//nolint:unused
func (u EditableUser) toModel(storeSet *v2models.StoreSet) (v2models.User, error) {
	return CreatableUser{EditableUser: u}.toModel(storeSet)
}

type UserController struct {
	ModelController[v2models.User, User, CreatableUser, EditableUser]
}

func newUserController(stores *v2models.StoreSet) *UserController {
	return &UserController{
		ModelController: ModelController[v2models.User, User, CreatableUser, EditableUser]{
			primaryStore:    stores.UserStore,
			allStores:       stores,
			modelToReadable: modelUserToUser,
		},
	}
}

func modelUserToUser(model *v2models.User) *User {
	if model == nil {
		return nil
	}
	return &User{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		StoredControlledUserFields: model.StoredControlledUserFields,
		CreatableUser: CreatableUser{
			EditableUser: EditableUser{
				StoredMutableUserFields: model.StoredMutableUserFields,
			},
		},
	}
}
