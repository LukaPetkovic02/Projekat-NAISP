package tokenBucket

import (
	"fmt"
	"time"
)

func Init(maxTokens uint64, rate uint64) *TokenBucket {

	return &TokenBucket{
		rate:           time.Duration(rate * uint64(time.Millisecond)),
		maxTokens:      maxTokens,
		currentTokens:  0,
		lastRefillTime: time.Now(),
	}

}

type TokenBucket struct {
	rate           time.Duration //vremenski period nakon kog reffilujemo bucket
	maxTokens      uint64        //maksimalan broj zahtjeva u odredjenom vremenskom periodu
	currentTokens  uint64        //trenutni broj zahtjeva u bucketu
	lastRefillTime time.Time     //vrijeme poslednjeg refreshovanja bucketa

}

func (tb *TokenBucket) RequestApproval() bool {

	now := time.Now()

	if now.Sub(tb.lastRefillTime) > tb.rate { //ako je proslo vise od 1 minute refillujemo bucket i dodamo zahtjev koji je upravo poslat
		tb.currentTokens = 1
		tb.lastRefillTime = time.Now()
		return true
	} else {
		if tb.currentTokens == tb.maxTokens { //ako je bucket pun odbijamo zahtjev uz poruku
			fmt.Println("Request denied !")
			return false
		} else { //ukoliko ima mjesta propustamo zahtjev
			tb.currentTokens += 1
			return true
		}
	}

}
