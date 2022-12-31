package jwt

import (
	"testing"
)

func TestJWTParse_JWTTokenParse(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyZGF0YSI6eyJpZCI6IjgyNzE2MTc2Njk2NjU1ODcyMSIsInVzZXJuYW1lIjoiaHZ0dXJpbmdnYSIsIm1haWwiOiJ4QGRpc2lzbS5jb20ifSwiand0Xy5fc3RhbmRhcmRfY2xhaW1zIjp7ImlzcyI6Im9hdXRoMi5kaXNpc20uY29tIiwiYXVkIjpbIiJdLCJleHAiOjE2NzM4ODc4NDgsIm5iZiI6MTY3MjUwNTQ0OCwiaWF0IjoxNjcyNTA1NDQ4LCJqdGkiOiJlOGU1ZTZjNS03YThkLTQxMDctOGRiYi0xNjNmOWIyMDE0ZmIifX0.46RkH7qSVOQHFcNqPjdrZVYt0hr01I8HQm05iYg1Aqc"
	parse, err := NewParseJWTToken(token, "373d27763b55593ca1964ee4b23ac04a1376fc67e5f74027").JWTTokenParse()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(parse)
}
