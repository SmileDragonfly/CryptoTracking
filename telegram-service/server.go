package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	grpc_telesvc "github.com/SmileDragonfly/go-lib/crypto-grpc/telegram-service"
	"github.com/golang/protobuf/ptypes/empty"
	"net/http"
	"telesvc/logger"
)

type Server struct {
	grpc_telesvc.UnimplementedTelegramServiceServer
}

func (s Server) SendMessage(ctx context.Context, request *grpc_telesvc.SendMessageRequest) (*empty.Empty, error) {
	link := "https://api.telegram.org/bot5466150074:AAEIrdGxJYOuGIfP6F2I5LN4DEE7mBql7Rc/sendMessage"
	req := SendMessageRequest{
		ChatID: request.ChatID,
		Text:   request.Text,
	}
	logger.Instance.Info("Request send message:", request.Text)
	reqByte, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(link, "application/json", bytes.NewBuffer(reqByte))
	if err != nil {
		return nil, err
	}
	logger.Instance.Info("Send message successfully")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		strErr := fmt.Sprintf("Send failed. Status: %q", resp.StatusCode)
		return nil, errors.New(strErr)
	}
	return nil, nil
}
