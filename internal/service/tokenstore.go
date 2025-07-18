package service

import (
    "golang.org/x/oauth2"
    "sync"
)

var (
    tokenStore = make(map[string]*oauth2.Token)
    mu         sync.RWMutex
)

func SaveToken(userID string, token *oauth2.Token) {
    mu.Lock()
    defer mu.Unlock()
    tokenStore[userID] = token
}

func GetToken(userID string) (*oauth2.Token, bool) {
    mu.RLock()
    defer mu.RUnlock()
    token, exists := tokenStore[userID]
    return token, exists
}
