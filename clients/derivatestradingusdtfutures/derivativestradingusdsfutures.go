package derivativestradingusdsfutures

import (
	"github.com/Daki-DM/binance-connector-go/clients/derivatestradingusdtfutures common/common"
	BinanceDerivativesTradingUsdsFuturesRestApi "github.com/Daki-DM/binance-connector-go/clients/derivatestradingusdtfutures/src/restapi"
	BinanceDerivativesTradingUsdsFuturesWebsocketApi "github.com/Daki-DM/binance-connector-go/clients/derivatestradingusdtfutures/src/websocketapi"
	BinanceDerivativesTradingUsdsFuturesWebsocketStreams "github.com/Daki-DM/binance-connector-go/clients/derivatestradingusdtfutures/src/websocketstreams"
)

type BinanceDerivativesTradingUsdsFuturesClient struct {
	RestApi          *BinanceDerivativesTradingUsdsFuturesRestApi.RestAPIClient
	WebsocketAPI     *BinanceDerivativesTradingUsdsFuturesWebsocketApi.WebsocketAPIClient
	WebsocketStreams *BinanceDerivativesTradingUsdsFuturesWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceDerivativesTradingUsdsFuturesClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceDerivativesTradingUsdsFuturesClient) {
		c.RestApi = BinanceDerivativesTradingUsdsFuturesRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketAPI(cfg *common.ConfigurationWebsocketApi) Option {
	return func(c *BinanceDerivativesTradingUsdsFuturesClient) {
		c.WebsocketAPI = BinanceDerivativesTradingUsdsFuturesWebsocketApi.NewWebsocketAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceDerivativesTradingUsdsFuturesClient) {
		c.WebsocketStreams = BinanceDerivativesTradingUsdsFuturesWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceDerivativesTradingUsdsFuturesClient(opts ...Option) *BinanceDerivativesTradingUsdsFuturesClient {
	client := &BinanceDerivativesTradingUsdsFuturesClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
