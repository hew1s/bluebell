package redis

import "time"

func Login(userID int64, token string) error {
	userStrID := string(userID)
	err := rdb.Set(userStrID, token, time.Hour*2).Err()
	return err
}

func UserIDAndToken(userID int64,token string)bool{
	newtoken ,_:= rdb.Get(string(userID)).Result()
	return newtoken == token
}