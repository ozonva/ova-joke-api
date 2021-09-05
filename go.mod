module github.com/ozonva/ova-joke-api

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Masterminds/squirrel v1.5.0
	github.com/Shopify/sarama v1.29.1 // indirect
	github.com/benbjohnson/clock v1.1.0
	github.com/fsnotify/fsnotify v1.5.0 // indirect
	github.com/golang/mock v1.6.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.2
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.16.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozonva/ova-joke-api/pkg/ova-joke-api v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.24.0
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/zhashkevych/go-sqlxmock v1.5.1
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/goleak v1.1.10
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
)

replace github.com/ozonva/ova-joke-api/pkg/ova-joke-api => ./pkg/ova-joke-api
