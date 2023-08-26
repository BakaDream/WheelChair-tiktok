package utils

func GenerateToken(username string, uid int64) string {
	return ""
}
func CheckToken(tokenString string) bool {
	return false
}
func GetUsername(tokenString string) (string, error) {
	var err error
	return "", err
}
func GetUserID(tokenString string) (uint, error) {
	var err error
	return 0, err
}
func GetUser(tokenString string) (uint, string, error) {
	var err error
	return 0, "", err
}
