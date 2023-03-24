module github.com/bit-broker/rate-service

go 1.15

require (
	github.com/alicebob/gopher-json v0.0.0-20200520072559-a9ecdc9d1d3a // indirect
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/datawire/ambassador v1.14.4
	github.com/go-redis/redis/v8 v8.8.2
	github.com/gomodule/redigo v1.8.4 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.5.1
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.27.5
	github.com/prometheus/client_golang v1.14.0
	github.com/rs/cors v1.8.3
	github.com/sirupsen/logrus v1.9.0
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da // indirect
	google.golang.org/grpc v1.37.0
)

replace (
	github.com/docker/distribution => github.com/distribution/distribution v2.7.1+incompatible
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
)
