package keycloak

import (
	"context"
	"errors"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
	"github.com/psmey/time_tracker/internal/logger"
)

var ErrInvalidToken = errors.New("jwt token is invalid")

type Client struct {
	client *gocloak.GoCloak
	jwks   *gocloak.CertResponse
	realm  string
}

type Config struct {
	BasePath string `yaml:"basePath"`
	Realm    string `yaml:"realm"`
}

func New(config *Config) (*Client, error) {
	client := gocloak.NewClient(config.BasePath)

	jwks, err := client.GetCerts(context.Background(), config.Realm)
	if err != nil {
		logger.LogError("failed to load JSON web key set from keycloak", err)
		return nil, err
	}

	keycloakClient := &Client{
		client: client,
		jwks:   jwks,
		realm:  config.Realm,
	}

	return keycloakClient, nil
}

func (keycloakClient *Client) ValidateToken(context context.Context, token string) (*jwt.MapClaims, error) {
	jwt, claims, err := keycloakClient.client.DecodeAccessToken(
		context,
		token,
		keycloakClient.realm,
	)
	if err != nil {
		logger.LogError("failed to decode access token", err)
		return nil, err
	}

	if !jwt.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
