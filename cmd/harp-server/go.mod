module github.com/elastic/harp/cmd/harp-server

go 1.15

replace github.com/elastic/harp => ../../

require (
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/dchest/uniuri v0.0.0-20200228104902-7aecb25e1fe5
	github.com/elastic/harp v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.10.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/google/wire v0.4.0
	github.com/gosimple/slug v1.9.0
	github.com/json-iterator/go v1.1.10
	github.com/magefile/mage v1.10.0
	github.com/oklog/run v1.1.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b	 // indirect
	github.com/spf13/cobra v1.1.1
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.33.1
)
