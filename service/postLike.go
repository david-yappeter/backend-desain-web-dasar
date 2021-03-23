package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
)

//PostCommendGetByArrayUserID Get By Array ID
func PostLikeGetByArrayPostID(ctx context.Context, ids []int) ([]*model.PostLike, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var postLike []*model.PostLike

	if err := db.Table("post_like").Where("post_id IN (?)", ids).Find(&postLike).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return postLike, nil
}
