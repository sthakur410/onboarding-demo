module onboarding-demo

go 1.15

// replace github.com/infobloxopen/atlas-app-toolkit => github.com/infobloxopen/atlas-app-toolkit v0.18.2

replace github.com/lyft/protoc-gen-validate => github.com/envoyproxy/protoc-gen-validate v0.0.7

// require (
// 	github.com/envoyproxy/protoc-gen-validate v0.4.1
// 	github.com/gogo/protobuf v1.2.1 // indirect
// 	github.com/golang/protobuf v1.4.2
// 	github.com/grpc-ecosystem/grpc-gateway v1.15.0
// 	github.com/infobloxopen/protoc-gen-gorm v0.20.0
// 	github.com/jinzhu/gorm v1.9.16
// 	golang.org/x/net v0.0.0-20200923182212-328152dc79b1 // indirect
// 	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
// 	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
// 	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
// 	google.golang.org/grpc v1.32.0
// )

// module onboarding-demo

// go 1.15

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.15.0
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.16
	github.com/prometheus/client_golang v1.7.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	golang.org/x/net v0.0.0-20200923182212-328152dc79b1 // indirect
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.32.0
)
