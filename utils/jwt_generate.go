package utils

// generate token example
//func JWTTokenGenerate(data entity.LoginGetData) (entity.Login, error) {
//	currentTime := time.Now()
//	claims := entity.MyCustomClaims{
//		UserID: data.UserID.String,
//		Nama:   data.Name.String,
//		Email:  data.Email.String,
//		RoleID: data.RoleID,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: currentTime.Add(time.Hour * 24).Unix(),
//			IssuedAt:  currentTime.Unix(),
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	idToken := entity.Login{}
//	tokenKey := os.Getenv(constant.JWTKey)
//
//	var err error
//	idToken.IDToken, err = token.SignedString([]byte(tokenKey))
//
//	return idToken, err
//}
