package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
)

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
