package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
	"myapp/tools"
)

//PostCommendCreate Create
func PostCommendCreate(ctx context.Context, input model.NewPostCommend) (*model.PostCommend, error) {
	user := ForContext(ctx)

	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	timeNow := tools.TimeNowString()

	postCommend := model.PostCommend{
		Body:      input.Body,
		CreatedAt: timeNow,
		UserID:    user.ID,
		PostID:    input.PostID,
	}

	if err := db.Table("post_commend").Create(&postCommend).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &postCommend, nil
}

//PostCommendGetByArrayUserID Get By Array ID
func PostCommendGetByArrayPostID(ctx context.Context, ids []int) ([]*model.PostCommend, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var postCommends []*model.PostCommend

	if err := db.Table("post_commend").Where("post_id IN (?)", ids).Find(&postCommends).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return postCommends, nil
}

//PostCommendDelete Delete
func PostCommendDelete(ctx context.Context, id int) (string, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Table("post_commend").Where("id = ?", id).Delete(&model.PostCommend{}).Error; err != nil {
		fmt.Println(err)
		return "Failed", err
	}

	return "Success", nil
}
