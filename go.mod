// +heroku goVersion go1.14
module myapp

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.20.12
)
