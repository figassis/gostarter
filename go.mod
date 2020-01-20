module geeks-accelerator/oss/saas-starter-kit

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/aws/aws-sdk-go v1.27.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dimfeld/httptreemux v5.0.1+incompatible
	github.com/dustin/go-humanize v1.0.0
	github.com/geeks-accelerator/files v0.0.0-20190704085106-630677cd5c14
	github.com/geeks-accelerator/sqlxmigrate v0.0.0-20190823021348-d047c980bb66
	github.com/geeks-accelerator/swag v1.6.3
	github.com/go-openapi/spec v0.19.2 // indirect
	github.com/go-openapi/swag v0.19.4 // indirect
	github.com/go-playground/locales v0.12.1
	github.com/go-playground/pkg v0.0.0-20190522230805-792a755e6910
	github.com/go-playground/universal-translator v0.16.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/google/go-cmp v0.4.0
	github.com/gorilla/schema v1.1.0
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.0
	github.com/huandu/go-sqlbuilder v1.4.1
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/ikeikeikeike/go-sitemap-generator/v2 v2.0.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.3.0
	github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/onsi/ginkgo v1.11.0 // indirect
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/pborman/uuid v1.2.0
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pkg/errors v0.9.0
	github.com/sethgrid/pester v0.0.0-20190127155807-68a33a018ad0
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/stretchr/testify v1.4.0
	github.com/sudo-suhas/symcrypto v1.0.0
	github.com/tdewolff/minify v2.3.6+incompatible
	github.com/tdewolff/parse v2.3.4+incompatible // indirect
	github.com/tinylib/msgp v1.1.0 // indirect
	github.com/urfave/cli v1.22.2
	github.com/xwb1989/sqlparser v0.0.0-20180606152119-120387863bf2
	gitlab.com/geeks-accelerator/oss/devops v1.0.57
	golang.org/x/crypto v0.0.0-20200109152110-61a87790db17
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	golang.org/x/sys v0.0.0-20200113162924-86b910548bc1 // indirect
	golang.org/x/tools v0.0.0-20200113223816-544dc8ea2d5f // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.16.1
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gotest.tools v2.2.0+incompatible // indirect
)

// replace gitlab.com/geeks-accelerator/oss/devops => ../devops

go 1.13
