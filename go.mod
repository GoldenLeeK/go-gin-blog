module github.com/GoldenLeeK/go-gin-blog

go 1.18

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.8.1
	github.com/go-ini/ini v1.67.0
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
)

require (
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace (
	github.com/GoldenLeeK/go-gin-example/conf => E:/dev/go/src/github.com/GoldenLeeK/go-gin-example/conf
	github.com/GoldenLeeK/go-gin-example/middleware => E:/dev/go/src/github.com/GoldenLeeK/go-gin-example/middleware
	github.com/GoldenLeeK/go-gin-example/models => E:/dev/go/src/github.com/GoldenLeeK/go-gin-example/models
	github.com/GoldenLeeK/go-gin-example/pkg/e => E:/dev/go/src/github.com/GoldenLeeK/go-gin-example/pkg/e
	github.com/GoldenLeeK/go-gin-example/pkg/setting => E:/dev/go/src/github.com/GoldenLeeK/go-gin-example/pkg/setting
	github.com/GoldenLeeK/go-gin-example/pkg/utils => E:/dev/go/src/github.com/GoldenLeeK/go-gin-example/pkg/utils
	github.com/GoldenLeeK/go-gin-example/routers => E:/dev/go/src/github.com/GoldenLeeK/go-gin-example/routers
	github.com/GoldenLeeK/go-gin-example/routers/api => E:/dev/go/src/github.com/GoldenLeeK/go-gin-example/routers/api
)
