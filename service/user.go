package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
	"myapp/tools"
)

//UserRegister register user
func UserRegister(ctx context.Context, input model.NewUser) (*model.AuthentificationToken, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	timeNow := tools.TimeNowString()

	user := model.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  tools.PasswordHash(input.Password),
		CreatedAt: timeNow,
		UpdatedAt: nil,
	}

	if err := db.Table("user").Create(&user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return TokenGenerate(ctx, user)
}

//UserUpdateSingleColumn Update Single Column
func UserUpdateSingleColumn(ctx context.Context, columnName string, value interface{}) (interface{}, error) {
	user := ForContext(ctx)

	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	timeNow := tools.TimeNowString()

	if err := db.Table("user").Where("user_id = ?", user.ID).Updates(map[string]interface{}{
		columnName:   value,
		"updated_at": timeNow,
	}).Error; err != nil {
		fmt.Println(err)
		return "Failed", err
	}

	return "Success", nil
}

//UserUpdateName Update Name
func UserUpdateName(ctx context.Context, name string) (string, error) {
	resp, err := UserUpdateSingleColumn(ctx, "name", name)
	return resp.(string), err
}

//UserUpdatePassword Update Password
func UserUpdatePassword(ctx context.Context, password string) (string, error) {
	resp, err := UserUpdateSingleColumn(ctx, "name", tools.PasswordHash(password))
	return resp.(string), err
}

//UserGetByID Get By ID
func UserGetByID(ctx context.Context, id int) (*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user model.User

	if err := db.Table("user").Where("id = ?", id).First(&user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

//UserGetByToken Get By Token
func UserGetByToken(ctx context.Context) (*model.User, error) {
	return UserGetByID(ctx, ForContext(ctx).ID)
}

//UserPaginationGetTotalData Pagination Total Data
func UserPaginationGetTotalData(ctx context.Context) (int, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Table("user").Count(&count).Error; err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int(count), nil
}

//UserPaginationGetNodes Pagination Nodes
func UserPaginationGetNodes(ctx context.Context, limit *int, page *int, sortBy *string, ascending *bool) ([]*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var users []*model.User

	query := db.Table("user")
	tools.QueryMaker(query, limit, page, ascending, sortBy)

	if err := query.Find(&users).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}
