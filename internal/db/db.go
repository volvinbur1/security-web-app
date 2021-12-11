package db

import (
	"context"
	"errors"
	"github.com/volvinbur1/security-web-app/internal/cmn"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Manager struct {
	client *mongo.Client
}

func New() *Manager {
	m := &Manager{}
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	var err error
	m.client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	return m
}

// Disconnect should be called before Manager release
func (m *Manager) Disconnect() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	err := m.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Manager) AddUser(user cmn.User) error {
	authTable := m.client.Database("secapp").Collection("users")

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()
	_, err := authTable.InsertOne(ctx, user)

	//	bson.D{
	//	{Key: "login", Value: user.Login}, {Key: "name", Value: user.Name}, {Key: "surname", Value: user.Surname},
	//	{Key: "pwdhash", Value: user.Password}, {Key: "pwdsalt", Value: user.PwdSalt},
	//})
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetUsers() ([]cmn.User, error) {
	authTable := m.client.Database("secapp").Collection("users")

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	cursor, err := authTable.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var usersBson []bson.M
	if err = cursor.All(ctx, &usersBson); err != nil {
		return nil, err
	}

	var users []cmn.User
	for i := 0; i < len(usersBson); i++ {
		user, err := parseBson(usersBson[i])
		if err != nil {
			continue
		}
		users = append(users, user)
	}

	return users, nil
}

func parseBson(data bson.M) (cmn.User, error) {
	u := cmn.User{}
	isOkay := true
	u.Guid, isOkay = data["guid"].(string)
	if !isOkay {
		return cmn.User{}, errors.New("error when parsing bson value `guid`")
	}

	u.Name, isOkay = data["name"].(string)
	if !isOkay {
		return cmn.User{}, errors.New("error when parsing bson value `name`")
	}

	u.Surname, isOkay = data["surname"].(string)
	if !isOkay {
		return cmn.User{}, errors.New("error when parsing bson value `surname`")
	}

	u.Phone, isOkay = data["phone"].(string)
	if !isOkay {
		return cmn.User{}, errors.New("error when parsing bson value `phone`")
	}

	u.Email, isOkay = data["email"].(string)
	if !isOkay {
		return cmn.User{}, errors.New("error when parsing bson value `email`")
	}

	u.Login, isOkay = data["login"].(string)
	if !isOkay {
		return cmn.User{}, errors.New("error when parsing bson value `login`")
	}

	u.Password, isOkay = data["pwdhash"].(string)
	if !isOkay {
		return cmn.User{}, errors.New("error when parsing bson value `pwdhash`")
	}

	u.PwdSalt, isOkay = data["pwdsalt"].(string)
	if !isOkay {
		return cmn.User{}, errors.New("error when parsing bson value `pwdsalt`")
	}

	return u, nil
}
