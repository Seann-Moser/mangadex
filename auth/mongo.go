package auth

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ TokenStore = (*MongoTokenStore)(nil)

type MongoTokenStore struct {
	Collection *mongo.Collection
}

func (m *MongoTokenStore) Load(
	ctx context.Context,
	userKey string,
) (*OAuthTokens, *OAuthCredentials, time.Time, error) {

	var doc struct {
		AccessToken  string    `bson:"access_token"`
		RefreshToken string    `bson:"refresh_token"`
		ExpiresAt    time.Time `bson:"expires_at"`

		ClientID     string `bson:"client_id,omitempty"`
		ClientSecret string `bson:"client_secret,omitempty"`
	}

	err := m.Collection.FindOne(ctx, bson.M{"user_key": userKey}).Decode(&doc)
	if err != nil {
		return nil, nil, time.Time{}, err
	}

	tokens := &OAuthTokens{
		AccessToken:  doc.AccessToken,
		RefreshToken: doc.RefreshToken,
	}

	var creds *OAuthCredentials
	if doc.ClientID != "" && doc.ClientSecret != "" {
		creds = &OAuthCredentials{
			ClientID:     doc.ClientID,
			ClientSecret: doc.ClientSecret,
		}
	}

	return tokens, creds, doc.ExpiresAt, nil
}

func (m *MongoTokenStore) Save(
	ctx context.Context,
	userKey string,
	tokens OAuthTokens,
	creds *OAuthCredentials,
	expiresAt time.Time,
) error {

	update := bson.M{
		"user_key":      userKey,
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"expires_at":    expiresAt,
	}

	// only store creds if provided
	if creds != nil {
		update["client_id"] = creds.ClientID
		update["client_secret"] = creds.ClientSecret
	}

	_, err := m.Collection.UpdateOne(
		ctx,
		bson.M{"user_key": userKey},
		bson.M{"$set": update},
		options.Update().SetUpsert(true),
	)

	return err
}

type EncryptedMongoTokenStore struct {
	Collection *mongo.Collection
	Cipher     TokenCipher
}

func (m *EncryptedMongoTokenStore) Save(
	ctx context.Context,
	userKey string,
	tokens OAuthTokens,
	creds *OAuthCredentials,
	expiresAt time.Time,
) error {

	payload := struct {
		Tokens OAuthTokens       `json:"tokens"`
		Creds  *OAuthCredentials `json:"credentials,omitempty"`
	}{
		Tokens: tokens,
		Creds:  creds,
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	enc, err := m.Cipher.Encrypt(raw)
	if err != nil {
		return err
	}

	_, err = m.Collection.UpdateOne(
		ctx,
		bson.M{"user_key": userKey},
		bson.M{
			"$set": bson.M{
				"user_key":   userKey,
				"data":       enc,
				"expires_at": expiresAt,
			},
		},
		options.Update().SetUpsert(true),
	)

	return err
}

func (m *EncryptedMongoTokenStore) Load(
	ctx context.Context,
	userKey string,
) (*OAuthTokens, *OAuthCredentials, time.Time, error) {

	var doc struct {
		Data      []byte    `bson:"data"`
		ExpiresAt time.Time `bson:"expires_at"`
	}

	err := m.Collection.FindOne(ctx, bson.M{"user_key": userKey}).Decode(&doc)
	if err != nil {
		return nil, nil, time.Time{}, err
	}

	raw, err := m.Cipher.Decrypt(doc.Data)
	if err != nil {
		return nil, nil, time.Time{}, err
	}

	var payload struct {
		Tokens OAuthTokens       `json:"tokens"`
		Creds  *OAuthCredentials `json:"credentials,omitempty"`
	}

	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, nil, time.Time{}, err
	}

	return &payload.Tokens, payload.Creds, doc.ExpiresAt, nil
}
