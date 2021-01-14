module github.com/SpiralMindJP/at-sdk/examples/at-web-server

go 1.16

require (
	cloud.google.com/go/storage v1.12.0 // indirect
	firebase.google.com/go/v4 v4.1.0
	github.com/BurntSushi/toml v0.3.1
	github.com/SpiralMindJP/at-sdk/go v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v1.5.1
	github.com/gorilla/websocket v1.4.2
	golang.org/x/oauth2 v0.0.0-20201208152858-08078c50e5b5 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/api v0.36.0
	google.golang.org/grpc v1.35.0
)

replace github.com/SpiralMindJP/at-sdk/go => ../../go
