package futuresusdt

import (
	binance "github.com/dirname/Binance"
	"github.com/dirname/Binance/model"
	"reflect"
	"testing"
)

func TestFuturesAllBookTickerWebsocketClient_GetCombined(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestFuturesAllBookTickerWebsocketClient_GetCombined", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			id: 0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FuturesAllBookTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.GetCombined(tt.args.id)
		})
	}
}

func TestFuturesAllBookTickerWebsocketClient_GetSubscribe(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestFuturesAllBookTickerWebsocketClient_GetSubscribe", fields{WebsocketClient: binance.WebsocketClient{}}, args{0}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FuturesAllBookTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.GetSubscribe(tt.args.id)
		})
	}
}

func TestFuturesAllBookTickerWebsocketClient_SetCombined(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		b  bool
		id uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestFuturesAllBookTickerWebsocketClient_SetCombined", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			b:  false,
			id: 0,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FuturesAllBookTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.SetCombined(tt.args.b, tt.args.id)
		})
	}
}

func TestFuturesAllBookTickerWebsocketClient_SetHandler(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		connectHandler  binance.ConnectedHandler
		responseHandler binance.ResponseHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestFuturesAllBookTickerWebsocketClient_SetHandler", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			connectHandler:  nil,
			responseHandler: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FuturesAllBookTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.SetHandler(tt.args.connectHandler, tt.args.responseHandler)
		})
	}
}

func TestFuturesAllBookTickerWebsocketClient_Subscribe(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		id     uint
		params []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestFuturesAllBookTickerWebsocketClient_Subscribe", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FuturesAllBookTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.Subscribe(tt.args.id, tt.args.params...)
		})
	}
}

func TestFuturesAllBookTickerWebsocketClient_Unsubscribe(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		id     uint
		params []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"TestFuturesAllBookTickerWebsocketClient_Unsubscribe", fields{WebsocketClient: binance.WebsocketClient{}}, args{
			id:     0,
			params: nil,
		}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FuturesAllBookTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			u.Unsubscribe(tt.args.id, "")
		})
	}
}

func TestFuturesAllBookTickerWebsocketClient_handleMessage(t *testing.T) {
	type fields struct {
		WebsocketClient binance.WebsocketClient
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"TestFuturesAllBookTickerWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"code\":0,\"msg\":\"Unknown property\",\"id\":0}")}, model.WebsocketErrorResponse{
			Code:    0,
			Message: "Unknown property",
		}, false},
		{"TestFuturesAllBookTickerWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"stream\":\"test\"}")}, AllBookTickerCombinedResponse{
			StreamName: "test",
		}, false},
		{"TestFuturesAllBookTickerWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"e\":\"test\"}")}, AllBookTickerResponse{
			EventType: "test",
		}, false},
		{"TestFuturesAllBookTickerWebsocketClient_handleMessage", fields{WebsocketClient: binance.WebsocketClient{}}, args{msg: []byte("{\"result\":\"test\"}")}, model.WebsocketCommonResponse{
			Result: "test",
		}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &FuturesAllBookTickerWebsocketClient{
				WebsocketClient: tt.fields.WebsocketClient,
			}
			got, err := u.handleMessage(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUSDTFuturesAllBookTickerWebsocketClient(t *testing.T) {
	type args struct {
		streams []string
	}
	tests := []struct {
		name string
		args args
		want *FuturesAllBookTickerWebsocketClient
	}{
		{"TestNewUSDTFuturesAllBookTickerWebsocketClient", args{streams: []string{"BTCUSDT@markPrice"}}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUSDTFuturesAllBookTickerWebsocketClient(tt.args.streams...); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUSDTFuturesAllBookTickerWebsocketClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
