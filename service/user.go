package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
	"myapp/tools"
	"strings"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

//UserRegister register user
func UserRegister(ctx context.Context, input model.NewUser) (*model.AuthentificationToken, error) {
	input.Email = strings.Trim(input.Email, " ")

	if strings.EqualFold(input.Email, "") {
		return nil, &gqlerror.Error{
			Message: "Email Empty!",
			Extensions: map[string]interface{}{
				"code": "Email Empty",
			},
		}
	}

	if err := tools.EmailValidate(strings.ToLower(input.Email)); err != nil {
		fmt.Println(err)
		return nil, &gqlerror.Error{
			Message: "Email Not Valid!",
			Extensions: map[string]interface{}{
				"code": "INVALID_EMAIL",
			},
		}
	}

	if input.Password != input.ConfirmPassword {
		return nil, &gqlerror.Error{
			Message: "Password & Confirm Password Different!",
			Extensions: map[string]interface{}{
				"code": "INVALID_CONFIRM_PASSWORD",
			}}
	}

	if len(input.Password) == 0 {
		return nil, &gqlerror.Error{
			Message: "Password Empty!",
			Extensions: map[string]interface{}{
				"code": "PASSWORD_EMPTY",
			}}
	}

	_, err := UserGetByEmail(ctx, strings.ToLower(input.Email))

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		return nil, err
	}

	if err == nil {
		fmt.Println(err)
		return nil, &gqlerror.Error{
			Message: "User Already Exists",
			Extensions: map[string]interface{}{
				"code": "USER_ALREADY_EXISTS",
			},
		}
	}

	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	timeNow := tools.TimeNowString()

	user := model.User{
		Name:      input.Name,
		Email:     strings.ToLower(input.Email),
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

	if err := db.Table("user").Where("id = ?", user.ID).Updates(map[string]interface{}{
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
	name = strings.Trim(name, " ")
	if name == "" {
		return "Failed", &gqlerror.Error{
			Message: "Name Must Not Be Empty!",
			Extensions: map[string]interface{}{
				"code": "EMPTY_NAME",
			},
		}
	}
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

	if user.Avatar != nil {
		tempAvatar := tools.GdriveViewLink(*user.Avatar)
		user.Avatar = &tempAvatar
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

	for index, val := range users {
		if val.Avatar != nil {
			tempAvatar := tools.GdriveViewLink(*val.Avatar)
			users[index].Avatar = &tempAvatar
		}
	}

	return users, nil
}

//UserGetByEmail Get By Email
func UserGetByEmail(ctx context.Context, email string) (*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user model.User

	if err := db.Table("user").Where("lower(email) = ?", email).First(&user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	if user.Avatar != nil {
		tempAvatar := tools.GdriveViewLink(*user.Avatar)
		user.Avatar = &tempAvatar
	}

	return &user, nil
}

//UserLogin Login
func UserLogin(ctx context.Context, email string, password string) (*model.AuthentificationToken, error) {
	if strings.EqualFold(strings.Trim(email, " "), "") {
		return nil, &gqlerror.Error{
			Message: "Email Must Not Be Empty!",
			Extensions: map[string]interface{}{
				"code": "EMAIL_EMPTY",
			}}
	}

	if err := tools.EmailValidate(strings.ToLower(email)); err != nil {
		fmt.Println(err)
		return nil, &gqlerror.Error{
			Message: "Email Not Valid!",
			Extensions: map[string]interface{}{
				"code": "INVALID_EMAIL",
			},
		}
	}

	getUser, err := UserGetByEmail(ctx, email)

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		return nil, err
	}

	if getUser == nil {
		fmt.Println(err)
		return nil, &gqlerror.Error{
			Message: "User Not Found",
			Extensions: map[string]interface{}{
				"code": "USER_NOT_FOUND",
			},
		}
	}

	if !tools.PasswordValidate(password, getUser.Password) {
		fmt.Println(err)
		return nil, &gqlerror.Error{
			Message: "Wrong Password",
			Extensions: map[string]interface{}{
				"code": "INVALID_PASSWORD",
			},
		}
	}

	return TokenGenerate(ctx, *getUser)
}

//UserGetByArrayID Get By Array ID
func UserGetByArrayID(ctx context.Context, ids []int) ([]*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var users []*model.User

	if err := db.Table("user").Where("id IN (?)", ids).Find(&users).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	for index, val := range users {
		if val.Avatar != nil {
			tempAvatar := tools.GdriveViewLink(*val.Avatar)
			users[index].Avatar = &tempAvatar
		}
	}

	return users, nil
}

//UserEditAvatar Edit Avatar
func UserEditAvatar(ctx context.Context, input model.EditAvatar) (string, error) {
	if input.Avatar != nil {
		if input.Avatar.ContentType == "image/jpeg" || input.Avatar.ContentType == "image/png" {
			if input.Avatar.Size < 26214400 {
				resp, err := UploadFile(ctx, *input.Avatar)

				if err != nil {
					fmt.Println(err)
					return "Failed", err
				}

				if _, err := UserUpdateSingleColumn(ctx, "avatar", resp); err != nil {
					fmt.Println(err)
					return "Failed", err
				}

				return tools.GdriveViewLink(resp), nil
			} else {
				return "Failed", &gqlerror.Error{
					Message: "File Exceeded 25MB",
					Extensions: map[string]interface{}{
						"code": "INVALID_FILE_SIZE",
					},
				}
			}
		} else {
			return "Failed", &gqlerror.Error{
				Message: "File Extensions Must Be .png .jpg .jpeg",
				Extensions: map[string]interface{}{
					"code": "INVALID_FILE_EXTENSIONS",
				},
			}
		}
	} else {
		if _, err := UserUpdateSingleColumn(ctx, "avatar", nil); err != nil {
			fmt.Println(err)
			return "Failed", err
		}
	}

	return "Success", nil
}
