package alltick

import (
	"time"

	"github.com/skip-mev/connect/v2/oracle/config"
)

const (
	// ByBit provides a few different URLs for its Websocket API. The URLs can be found
	// The two production URLs are defined in ProductionURL and TestnetURL.

	// Name is the name of the ByBit provider.
	Name = "alltick_ws"

	// URLProd is the public ByBit Websocket URL.
	URLProd = "wss://quote.alltick.io/quote-stock-b-ws-api?token=328a73ad19c913a0a7e79fb0a0e0f6ac-c-app"

	// DefaultPingInterval is the default ping interval for the ByBit websocket.
	DefaultPingInterval = 10 * time.Second
)

// DefaultWebSocketConfig is the default configuration for the ByBit Websocket.
var DefaultWebSocketConfig = config.WebSocketConfig{
	Name:                          Name,
	Enabled:                       true,
	MaxBufferSize:                 1000,
	ReconnectionTimeout:           config.DefaultReconnectionTimeout,
	PostConnectionTimeout:         config.DefaultPostConnectionTimeout,
	Endpoints:                     []config.Endpoint{{URL: URLProd}},
	ReadBufferSize:                config.DefaultReadBufferSize,
	WriteBufferSize:               config.DefaultWriteBufferSize,
	HandshakeTimeout:              config.DefaultHandshakeTimeout,
	EnableCompression:             config.DefaultEnableCompression,
	ReadTimeout:                   config.DefaultReadTimeout,
	WriteTimeout:                  config.DefaultWriteTimeout,
	PingInterval:                  DefaultPingInterval,
	WriteInterval:                 config.DefaultWriteInterval,
	MaxReadErrorCount:             config.DefaultMaxReadErrorCount,
	MaxSubscriptionsPerConnection: config.DefaultMaxSubscriptionsPerConnection,
	MaxSubscriptionsPerBatch:      config.DefaultMaxSubscriptionsPerBatch,
}
