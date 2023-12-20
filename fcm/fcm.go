package fcm

import (
	"context"
	"errors"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FCM struct {
	App    *firebase.App
	Client *messaging.Client
}

func NewFCM(firebaseAuthKey string) (*FCM, error) {
	decodedKey, err := DecodeFirebaseAuthKey(firebaseAuthKey)
	if err != nil {
		return &FCM{}, err
	}

	opts := []option.ClientOption{option.WithCredentialsJSON(decodedKey)}
	app, err := firebase.NewApp(context.Background(), nil, opts...)
	if err != nil {
		return &FCM{}, err
	}

	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		return &FCM{}, err
	}

	return &FCM{
		App:    app,
		Client: fcmClient,
	}, nil
}

func (f *FCM) SendTopicCtx(ctx context.Context, topic string, message *messaging.Message) error {
	if topic == "" {
		return errors.New("topic is required")
	}

	message.Topic = topic
	_, err := f.Client.Send(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func (f *FCM) SendTopic(topic string, message *messaging.Message) error {
	if topic == "" {
		return errors.New("topic is required")
	}

	message.Topic = topic
	_, err := f.Client.Send(context.Background(), message)
	if err != nil {
		return err
	}

	return nil
}
