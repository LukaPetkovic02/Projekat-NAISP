package tokenBucket

import (
	"fmt"
	"time"

	"github.com/LukaPetkovicSV16/Projekat-NAISP/config"
)

func Init() *TokenBucket {

	return &TokenBucket{
		rate:           uint(config.Values.TokenBucket.Rate),
		currentTokens:  config.Values.TokenBucket.Size,
		lastRefillTime: uint(time.Now().Unix()),
	}

}

type TokenBucket struct {
	rate           uint //vremenski period nakon kog reffilujemo bucket
	currentTokens  uint //trenutni broj zahtjeva u bucketu
	lastRefillTime uint //vrijeme poslednjeg refreshovanja bucketa

}

func (tb *TokenBucket) RequestApproval() bool {

	now := uint(time.Now().Unix())

	if now-tb.lastRefillTime > tb.rate { //ako je proslo vise od dozvoljenog vremena refillujemo bucket i dodamo zahtjev koji je upravo poslat
		tb.currentTokens = config.Values.TokenBucket.Size
		tb.lastRefillTime = now
		return true
	} else {
		if tb.currentTokens == 0 { //ako je bucket pun odbijamo zahtjev uz poruku
			fmt.Println("Request denied !")
			return false
		} else { //ukoliko ima mjesta propustamo zahtjev
			tb.currentTokens -= 1
			return true
		}
	}

}
