package config

type ServerProductListener struct {
	GrpcListenerHost string `env:"PRODUCT_GRPC_LISTENER_HOST" env-default:"0.0.0.0"`
	RestPort         string `env:"PRODUCT_REST_PORT" env-default:"9080"`
	GrpcPort         string `env:"PRODUCT_GRPC_PORT" env-default:"9082"`
}
