module auth

go 1.27.1

replace github.com/KTemka1234/go-micro/contracts => ../contracts

require (
	github.com/KTemka1234/go-micro/contracts v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.19.0
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.12.3
	github.com/pressly/goose/v3 v3.28.0
	github.com/rs/zerolog v1.35.1
	golang.org/x/crypto v0.55.0
	google.golang.org/grpc v1.84.0
	google.golang.org/protobuf v1.36.12
	gorm.io/driver/postgres v1.6.3
	gorm.io/gorm v1.31.2
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.10.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/rogpeppe/go-internal v1.16.0 // indirect
	github.com/sethvargo/go-retry v0.4.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260918162117-cecb64721679 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
