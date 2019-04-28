module retail

require (
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	//github.com/gin-gonic/gin v1.3.0
	github.com/gin-gonic/gin v1.3.0
	github.com/go-openapi/jsonreference v0.19.0 // indirect
	github.com/go-openapi/spec v0.19.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/lexkong/log v0.0.0-20180607165131-972f9cd951fc
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.2
	github.com/swaggo/gin-swagger v1.0.0
	github.com/swaggo/swag v1.4.0
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	golang.org/x/tools v0.0.0-20190208222737-3744606dbb67
	gopkg.in/gin-gonic/gin.v1 v1.3.0
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	github.com/graphql-go/graphql v0.7.8
)

//replace gopkg.in/gin-gonic/gin.v1 => github.com/gin-gonic/gin v1.3.0

replace golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20181205085412-a5c9d58dba9a

replace golang.org/x/net v0.0.0-20181005035420-146acd28ed58 => github.com/golang/net v0.0.0-20181005035420-146acd28ed58

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190208222737-3744606dbb67
