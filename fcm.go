package apppush

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FCMClient struct {
	client *messaging.Client
}

func (f *FCMClient) Send(
	id string,
	title string,
	body string,
	data map[string]string,
) error {
	_, err := f.client.Send(context.TODO(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: id,
		Data:  data,
	})
	if err != nil {
		return err
	}

	return nil
}

func (f *FCMClient) SendMulticast(
	ids []string,
	title string,
	body string,
	data map[string]string,
) error {
	_, err := f.client.SendMulticast(context.TODO(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Tokens: ids,
		Data:   data,
	})

	if err != nil {
		return err
	}

	return nil
}

func NewFCMClient(credentialJSON []byte) (Client, error) {
	opts := option.WithCredentialsJSON(credentialJSON)
	app, err := firebase.NewApp(context.Background(), nil, opts)
	if err != nil {
		return nil, err
	}
	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}

	return &FCMClient{
		client: client,
	}, nil
}
