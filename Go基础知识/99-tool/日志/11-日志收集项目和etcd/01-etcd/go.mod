module 01-etcd

go 1.20

require (
	github.com/coreos/etcd v3.3.27+incompatible // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	go.etcd.io/bbolt v1.3.10 // indirect
	go.etcd.io/etcd v3.3.27+incompatible // indirect
	golang.org/x/sys v0.4.0 // indirect
)

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
