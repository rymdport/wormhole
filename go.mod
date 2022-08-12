module github.com/rymdport/wormhole

go 1.13

require (
	github.com/gorilla/websocket v1.5.0
	github.com/klauspost/compress v1.15.9
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	nhooyr.io/websocket v1.8.7
	salsa.debian.org/vasudev/gospake2 v0.0.0-20210510093858-d91629950ad1
)

replace nhooyr.io/websocket => github.com/rymdport/websocket v1.9.0
