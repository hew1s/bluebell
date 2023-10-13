package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExipire = errors.New("投票时间已过")
)

func CreatePost(pid int64)error{
	_ ,err := rdb.ZAdd(getRedisKey(KeyPostTimeZSet),redis.Z{
		Score:float64(time.Now().Unix()),
		Member: pid,
	}).Result()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 查询是否超过七天
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExipire
	}
	// 先查询之前投票记录
	ov := rdb.ZScore(getRedisKey(LeyPostVotedPF+postID), userID).Val()
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value)
	_, err := rdb.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID).Result()
	if err != nil {
		return err
	}
	if value == 0 {
		_, err = rdb.ZRem(getRedisKey(LeyPostVotedPF+postID), userID).Result()
	} else {
		_, err = rdb.ZAdd(getRedisKey(LeyPostVotedPF+postID), redis.Z{
			Score:  value, //赞成还是反对
			Member: userID,
		}).Result()
	}
	return err
}
