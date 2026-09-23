module gateway

go 1.27.1

require (
	github.com/KTemka1234/go-micro/contracts v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.19.0
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/joho/godotenv v1.5.1
	github.com/rs/zerolog v1.35.1
	google.golang.org/grpc v1.84.0
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260921155816-b14227669459 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260918162117-cecb64721679 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace github.com/KTemka1234/go-micro/contracts => ../contracts
