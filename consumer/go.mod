module myproject/consumer

go 1.21.0

toolchain go1.24.3

require (
	github.com/go-sql-driver/mysql v1.9.3
	github.com/segmentio/kafka-go v0.4.36
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
)

replace myproject/modules/model => ./modules/model
