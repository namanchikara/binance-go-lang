package spotclient

import (
	"encoding/json"
	"github.com/dirname/Binance"
	"github.com/dirname/Binance/config"
	logger "github.com/dirname/Binance/logging"
	"github.com/dirname/Binance/model"
)

// SpotPartialBookDepthWebsocketClient responsible to handle market price data from websocket
type SpotPartialBookDepthWebsocketClient struct {
	binance.WebsocketClient
}

// NewSpotPartialBookDepthWebsocketClient Factory function
func NewSpotPartialBookDepthWebsocketClient(streams ...string) *SpotPartialBookDepthWebsocketClient {
	c := new(SpotPartialBookDepthWebsocketClient)
	c.WebsocketClient.Init(config.SpotWssHost, streams...)
	return c
}

// Subscribe subscribe market price data
func (u *SpotPartialBookDepthWebsocketClient) Subscribe(id uint, params ...string) {

	req := model.WebsocketCommon{
		Method: "SUBSCRIBE",
		Params: params,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket subscribed, params: %s id: %d", params, id)
}

// Unsubscribe unsubscribe market price data
func (u *SpotPartialBookDepthWebsocketClient) Unsubscribe(id uint, params ...string) {
	req := model.WebsocketCommon{
		Method: "UNSUBSCRIBE",
		Params: params,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket unsubscribed, params: %s id: %d", params, id)
}

// SetCombined set combined, it is true when stream's length is greater than one, and false if stream's length is equal to one
func (u *SpotPartialBookDepthWebsocketClient) SetCombined(b bool, id uint) {
	req := model.WebsocketCommon{
		Method: "SET_PROPERTY",
		Params: []interface{}{"combined", b},
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket set combined, params: %v", req)
}

// GetCombined get combined
func (u *SpotPartialBookDepthWebsocketClient) GetCombined(id uint) {
	req := model.WebsocketCommon{
		Method: "GET_PROPERTY",
		Params: []interface{}{"combined"},
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket set combined, params: %v", req)
}

// GetSubscribe get subscribed list
func (u *SpotPartialBookDepthWebsocketClient) GetSubscribe(id uint) {
	req := model.WebsocketCommon{
		Method: "LIST_SUBSCRIPTIONS",
		Params: nil,
		ID:     id,
	}
	u.WebsocketClient.SendJSON(req)

	logger.Info("Websocket get subscribed, params: %v", req)
}

// SetHandler set callback handler
func (u *SpotPartialBookDepthWebsocketClient) SetHandler(connectHandler binance.ConnectedHandler, responseHandler binance.ResponseHandler) {
	u.WebsocketClient.SetConnectedHandler(connectHandler)
	u.WebsocketClient.SetResponseHandler(responseHandler)
	u.WebsocketClient.SetMessageHandler(u.handleMessage)
}

func (u *SpotPartialBookDepthWebsocketClient) handleMessage(msg []byte) (interface{}, error) {
	var parser map[string]interface{}
	var err error
	err = json.Unmarshal(msg, &parser)
	if _, ok := parser["stream"]; ok {
		result := PartialBookDepthCombinedResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	if _, ok := parser["result"]; ok {
		result := model.WebsocketCommonResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	if _, ok := parser["lastUpdateId"]; ok {
		result := PartialBookDepthResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	if _, ok := parser["code"]; ok {
		result := model.WebsocketErrorResponse{}
		err = json.Unmarshal(msg, &result)
		return result, err
	}
	return parser, err
}
