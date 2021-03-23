// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthOps struct {
	Register *AuthentificationToken `json:"register"`
	Login    *AuthentificationToken `json:"login"`
}

type AuthentificationToken struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type EditUserName struct {
	Name string `json:"name"`
}

type EditUserPassword struct {
	Password string `json:"password"`
}

type NewPost struct {
	Body string `json:"body"`
}

type NewPostCommend struct {
	Body   string `json:"body"`
	PostID int    `json:"post_id"`
}

type NewPostLike struct {
	Body   string `json:"body"`
	PostID int    `json:"post_id"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PostCommendOps struct {
	Create *PostCommend `json:"create"`
	Delete string       `json:"delete"`
}

type PostLikeOps struct {
	Create *PostLike `json:"create"`
	Delete string    `json:"delete"`
}

type PostOps struct {
	Create *Post  `json:"create"`
	Delete string `json:"delete"`
}

type PostPagination struct {
	Limit     *int    `json:"limit"`
	Page      *int    `json:"page"`
	SortBy    *string `json:"sort_by"`
	Ascending *bool   `json:"ascending"`
	TotalData int     `json:"total_data"`
	Nodes     []*Post `json:"nodes"`
}

type UserOps struct {
	EditName     string `json:"edit_name"`
	EditPassword string `json:"edit_password"`
}

type UserPagination struct {
	Limit     *int    `json:"limit"`
	Page      *int    `json:"page"`
	SortBy    *string `json:"sort_by"`
	Ascending *bool   `json:"ascending"`
	TotalData int     `json:"total_data"`
	Nodes     []*User `json:"nodes"`
}