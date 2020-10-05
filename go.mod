module github.com/binarly-io/atlas

go 1.13

require (
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/Shopify/sarama v1.26.1
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-cmd/cmd v1.2.1
	github.com/golang/protobuf v1.4.0
	github.com/google/uuid v1.1.1
	github.com/mailru/go-clickhouse v1.3.0
	github.com/minio/minio-go/v6 v6.0.57
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/viper v1.4.0
	github.com/ulikunitz/xz v0.5.7
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/grpc v1.28.1
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

replace golang.org/x/oauth2 => github.com/sunsingerus/oauth2 v0.0.0-20200410181841-d7afaacd4cbe
