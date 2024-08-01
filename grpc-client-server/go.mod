module rqa

go 1.21.0

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.65.0

require (
	github.com/sanjay7178/go-basics v0.0.0-20240801092216-5642d722ed60
	google.golang.org/grpc v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)
