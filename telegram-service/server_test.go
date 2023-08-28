package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	grpc_telesvc "grpcsvc/telegram-service"
	"reflect"
	"telesvc/logger"
	"testing"
	"time"
)

func TestServer_SendMessage(t *testing.T) {
	// Init log
	err := logger.NewLogger("./config/logcfg.json")
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		UnimplementedTelegramServiceServer grpc_telesvc.UnimplementedTelegramServiceServer
	}
	type args struct {
		ctx     context.Context
		request *grpc_telesvc.SendMessageRequest
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			"TestSendMessage",
			fields{grpc_telesvc.UnimplementedTelegramServiceServer{}},
			args{
				ctx: context.Background(),
				request: &grpc_telesvc.SendMessageRequest{
					ChatID: 1930587972,
					Text:   time.Now().String() + " Send test message from crypto-telegram-service",
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				UnimplementedTelegramServiceServer: tt.fields.UnimplementedTelegramServiceServer,
			}
			got, err := s.SendMessage(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
