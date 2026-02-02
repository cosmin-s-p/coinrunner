package coinrunner

// slice to simulate data we would store in a database
// used to store Tokens, keys represent idempotency keys for ease of use
var Memory []Token = []Token{}

func CheckInMemoryData(token Token) bool {
	for _, t := range Memory {
		if t.IdempotencyKey == token.IdempotencyKey {
			// found it
			return true
		}
	}

	return false
}
