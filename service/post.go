package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
	"myapp/tools"
	"strings"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

//PostCreate Create
func PostCreate(ctx context.Context, input model.NewPost) (*model.Post, error) {
	if strings.Trim(input.Body, " ") == "" {
		return nil, &gqlerror.Error{
			Message: "Content Must Not Be Empty!",
			Extensions: map[string]interface{}{
				"code": "CONTENT_MUST_NOT_BE_EMPTY",
			},
		}
	}

	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	timeNow := tools.TimeNowString()

	post := model.Post{
		Body:      input.Body,
		CreatedAt: timeNow,
		UserID:    ForContext(ctx).ID,
	}

	if err := db.Table("post").Create(&post).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &post, nil
}

//PostGetByID Get By ID
func PostGetByID(ctx context.Context, id int) (*model.Post, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var post model.Post

	if err := db.Table("post").Where("id = ?", id).First(&post).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &post, nil
}

//PostPaginationGetTotalData Pagination Total Data
func PostPaginationGetTotalData(ctx context.Context) (int, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var count int64

	if err := db.Table("post").Count(&count).Error; err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int(count), nil
}

//PostPaginationGetNodes  Pagination Nodes
func PostPaginationGetNodes(ctx context.Context, limit *int, page *int, sortBy *string, ascending *bool) ([]*model.Post, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var posts []*model.Post

	query := db.Table("post")
	tools.QueryMaker(query, limit, page, ascending, sortBy)

	if err := query.Find(&posts).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return posts, nil
}

func PostDelete(ctx context.Context, id int) (string, error) {
	getPost, err := PostGetByID(ctx, id)

	if err != nil {
		fmt.Println(err)
		return "Failed", err
	}

	if getPost.UserID != ForContext(ctx).ID {
		return "Failed", &gqlerror.Error{
			Message: "Not Owner Of The Post",
			Extensions: map[string]interface{}{
				"code": "NOT_OWNER_OF_POST",
			},
		}
	}

	if _, err := PostCommendDeleteByPostID(ctx, id); err != nil {
		fmt.Println(err)
		return "Failed", err
	}

	if _, err := PostLikeDeleteByPostID(ctx, id); err != nil {
		fmt.Println(err)
		return "Failed", err
	}

	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Table("post").Where("id = ?", id).Delete(&model.Post{}).Error; err != nil {
		fmt.Println(err)
		return "Failed", err
	}

	return "Success", nil
}
