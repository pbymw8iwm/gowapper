package utilswapper

import (
	"time"

	"github.com/kataras/jwt"
)

func JwtEncode(myClaims interface{}, key string, expire time.Duration) string {
	token, err := jwt.Sign(jwt.HS256, key, myClaims, jwt.MaxAge(expire*time.Minute))
	if err != nil {
		return ""
	}
	return string(token)
}
func JwtDecode(token []byte, key string, claims interface{}) (int64, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, key, token)
	if err != nil {
		return 0, err
	}
	err = verifiedToken.Claims(&claims)
	return verifiedToken.StandardClaims.IssuedAt, err
}
