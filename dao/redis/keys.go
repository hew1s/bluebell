package redis

// 冒号分割
// redis 中的key注意使用命名空间的方式 方便查询拆分

const (
	KeyPrefix        = "bluebell:"
	KeyPostTimeZSet  = "post:time"
	KeyPostScoreZSet = "post:score"
	LeyPostVotedPF   = "post:voted:"
)

// 给redis   key加前缀
func getRedisKey(key string)string{
	return KeyPrefix+key
}
