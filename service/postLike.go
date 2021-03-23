package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
	"myapp/tools"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

//PostLikeCreate Create
func PostLikeCreate(ctx context.Context, input model.NewPostLike) (*model.PostLike, error) {
	user := ForContext(ctx)
	getPostLike, err := PostLikeGetByUserIDAndPostID(ctx, user.ID, input.PostID)

	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		return nil, err
	}

	if getPostLike != nil {
		return nil, &gqlerror.Error{
			Message: "Already Liked",
			Extensions: map[string]interface{}{
				"code": "ALREADY_LIKED",
			},
		}
	}

	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	timeNow := tools.TimeNowString()

	postLike := model.PostLike{
		Body:      input.Body,
		CreatedAt: timeNow,
		UserID:    user.ID,
		PostID:    input.PostID,
	}

	if err := db.Table("post_like").Create(&postLike).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &postLike, nil
}

//PostLikeGetByUserIDAndPostID Get By UserID and PostID
func PostLikeGetByUserIDAndPostID(ctx context.Context, userID int, postID int) (*model.PostLike, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var postLike model.PostLike

	if err := db.Table("post_like").Where("user_id = ? AND post_id = ?", userID, postID).First(&postLike).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &postLike, nil
}

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

//PostLikeDelete Delete
func PostLikeDelete(ctx context.Context, id int) (string, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Table("post_like").Where("id = ?", id).Delete(&model.PostLike{}).Error; err != nil {
		fmt.Println(err)
		return "Failed", err
	}

	return "Success", nil
}
