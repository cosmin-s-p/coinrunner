package coinrunner

import (
	"math/rand/v2"
	"time"
)

func GetIPS() []string {
	return []string{
		"ipTest1",
		"ipTest2",
		"ipTest3",
		"ipTest4",
		"ipTest5",
	}
}

func GetIdempotencyKeys() []string {
	return []string{
		"keyTest1",
		"keyTest2",
		"keyTest3",
		"keyTest4",
		"keyTest5",
		"keyTest6",
		"keyTest7",
		"keyTest8",
		"keyTest9",
		"keyTest10",
	}
}

func GetRandomIP() string {
	ips := GetIPS()

	randomIndex := rand.IntN(len(ips))

	return ips[randomIndex]
}

func GetRandomIdempotencyKey() string {
	keys := GetIdempotencyKeys()

	randomIndex := rand.IntN(len(keys))

	return keys[randomIndex]
}

func InitializeRandomToken() Token {
	return Token{
		IdempotencyKey: GetRandomIdempotencyKey(),
		RiskScore:      rand.IntN(100),
		PaidAmount:     rand.Float32() * 500,
		SenderIp:       GetRandomIP(),
	}
}

func SimulateParallelRequests() {

	// create requests with random tokens for duplicate requests and velocity checks
	for len(Memory) < 100 {
		time.Sleep(time.Second * 2)

		token := InitializeRandomToken()
		Memory = append(Memory, token)
	}
}
