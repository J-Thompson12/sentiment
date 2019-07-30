package classify

import (
	"github.com/gomodule/redigo/redis"
)

func init() {
	setupRedis()
}

//RedisTrain trains a document and adds it to a redis serve
func RedisTrain(category string, data string) error {

	c := redisPool.Get()
	defer c.Close()

	wordCategoryTotal := hget(c, "word:"+category, "count")
	totalWords := hget(c, "TotalWords", "count")
	documentCategoryTotal := hget(c, "doc:"+category, "count")
	totalDocuments := hget(c, "TotalDocuments", "count")

	w := countWords(data)
	for word, count := range w {
		reply := hget(c, category+":"+word, "count")
		totalCount := count + reply
		c.Send("HMSET", category+":"+word, "count", totalCount)

		wordCategoryTotal += count
		totalWords += count
	}

	// going into redis
	documentCategoryTotal++
	totalDocuments++

	c.Send("HMSET", "word:"+category, "count", wordCategoryTotal)
	c.Send("HMSET", "TotalWords", "count", totalWords)
	c.Send("HMSET", "doc:"+category, "count", documentCategoryTotal)
	c.Send("HMSET", "TotalDocuments", "count", totalDocuments)

	c.Flush()
	return nil
}

func hget(c redis.Conn, key string, value string) float64 {
	reply, err := redis.Float64(c.Do("HGET", key, value))
	if err != nil {
		//fmt.Println(err)
	}
	return reply
}
