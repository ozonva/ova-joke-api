module github.com/ozonva/ova-joke-api

go 1.16

require (
	github.com/benbjohnson/clock v1.1.0
	github.com/fsnotify/fsnotify v1.5.0 // indirect
	github.com/golang/mock v1.6.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.16.0
	github.com/ozonva/ova-joke-api/pkg/ova-joke-api v0.0.0-20210904092453-a0d5e5d51b5e
	github.com/rs/zerolog v1.24.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/goleak v1.1.10
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
)

replace github.com/ozonva/ova-joke-api/pkg/ova-joke-api => ./pkg/ova-joke-api
