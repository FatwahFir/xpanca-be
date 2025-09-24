package jwtx

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	secret []byte
	ttl time.Duration
}

type Claims struct {
	UserID uint `json:"uid"`
	Username string `json:"uname"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func New(secret string, ttlMinutes int) *Manager { return &Manager{secret: []byte(secret), ttl: time.Duration(ttlMinutes)*time.Minute} }

func (m *Manager) Generate(uid uint, uname, role string) (string, error) {
	claims := &Claims{
		UserID: uid,
		Username: uname,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.ttl))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

func (m *Manager) Verify(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) { return m.secret, nil })
	if err != nil { return nil, err }
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { return claims, nil }
	return nil, jwt.ErrTokenInvalidClaims
}
