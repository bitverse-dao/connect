package alltick

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/skip-mev/connect/v2/providers/base/websocket/handlers"
)

type (
	// Operation is the operation to perform. This is used to construct subscription messages
	// when initially connecting to the websocket. This can later be extended to support
	// other operations.
	Operation uint64

	// Channel is the channel to subscribe to. The channel is used to determine the type of
	// price data that we want. This can later be extended to support other channels.
)

const (
	// OperationSubscribe is the operation to subscribe to a channel.
	OperationSubscribe Operation = 22005
	OperationTicker              = 22998
	OperationPing      Operation = 22001
)

type BaseRequest struct {
	CmdId uint64 `json:"cmd_id"`
	SeqId uint64 `json:"seq_id"`
	Trace string `json:"trace"`
}

type SubscriptionRequest struct {
	BaseRequest
	Data Data `json:"data"`
}

type Data struct {
	SymbolList []Code `json:"symbol_list"`
}

type Code struct {
	Code string `json:"code"`
}

// NewSubscriptionRequestMessage creates subscription messages corresponding to the provided tickers.
// If the number of tickers is greater than 10, the requests will be broken into 10-ticker messages.
func NewSubscriptionRequestMessage(tickers []string) ([]handlers.WebsocketEncodedMessage, error) {
	numTickers := len(tickers)
	if numTickers == 0 {
		return nil, fmt.Errorf("tickers cannot be empty")
	}

	messages := make([]handlers.WebsocketEncodedMessage, len(tickers))

	for i, ticker := range tickers {
		code := strings.ReplaceAll(ticker, "USDT", ".US")
		bz, err := json.Marshal(
			SubscriptionRequest{
				BaseRequest: BaseRequest{
					CmdId: 22004,
					SeqId: 123,
					Trace: fmt.Sprintf("3baaa938-f92c-4a74-a228-fd49d5e2f8bc-%d", time.Now().UnixMilli()),
				},
				Data: Data{
					SymbolList: []Code{
						{
							Code: code,
						},
					},
				},
			},
		)
		if err != nil {
			return messages, fmt.Errorf("unable to marshal message: %w", err)
		}

		messages[i] = bz

	}

	return messages, nil
}

type HeartbeatPing struct {
	CmdId uint64   `json:"cmd_id"`
	SeqId uint64   `json:"seq_id"`
	Trace string   `json:"trace"`
	Data  struct{} `json:"data"`
}

// NewHeartbeatPingMessage returns the encoded message for sending a heartbeat message to a peer.
func NewHeartbeatPingMessage() ([]handlers.WebsocketEncodedMessage, error) {
	bz, err := json.Marshal(
		HeartbeatPing{
			CmdId: 22000,
			SeqId: 123,
			Trace: fmt.Sprintf("3baaa938-f92c-4a74-a228-fd49d5e2f8bc-%d", time.Now().UnixMilli()),
			Data:  struct{}{},
		},
	)

	return []handlers.WebsocketEncodedMessage{bz}, err
}

type BaseMessage struct {
	CmdId uint64 `json:"cmd_id"`
}

// HeartbeatPong is the pong sent back from the server after a ping.
//
// Example:
//
//	{
//	   "success": true,
//	   "ret_msg": "pong",
//	   "conn_id": "0970e817-426e-429a-a679-ff7f55e0b16a",
//	   "op": "ping"
//	}
type HeartbeatPong struct {
	Ret   uint64   `json:"ret"`
	Msg   string   `json:"msg"`
	CmdId uint64   `json:"cmd_id"`
	SeqId uint64   `json:"seq_id"`
	Trace string   `json:"trace"`
	Data  struct{} `json:"data"`
}

// SubscriptionResponse is the response for a subscribe event.
type SubscriptionResponse struct {
	Ret   uint64   `json:"ret"`
	Msg   string   `json:"msg"`
	CmdId uint64   `json:"cmd_id"`
	SeqId uint64   `json:"seq_id"`
	Trace string   `json:"trace"`
	Data  struct{} `json:"data"`
}

// TickerUpdateMessage is the update sent for a subscribed ticker on the ByBit websocket API.
type TickerUpdateMessage struct {
	CmdId uint64           `json:"cmd_id"`
	Data  TickerUpdateData `json:"data"`
}

// TickerUpdateData is the data stored inside a ticker update message.
type TickerUpdateData struct {
	Code     string `json:"code"`
	TickTime string `json:"tick_time"`
	Price    string `json:"price"`
}
