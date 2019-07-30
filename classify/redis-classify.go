package classify

import "github.com/gomodule/redigo/redis"

func init() {
	setupRedis()
}

func RedisClassify(document string, categories []string) (p map[string]float64) {
	c := redisPool.Get()
	defer c.Close()

	p = make(map[string]float64)

	for _, category := range categories {
		p[category] = wordFrequency(c, category, document) * categoryProb(c, category)
	}
	return
}

// Checks each individual word to see what category they belong to and multiplies them
func wordFrequency(c redis.Conn, category string, document string) (p float64) {
	p = 1.0
	wnc := 0.0
	words := countWords(document)
	for word := range words {
		wnc = (hget(c, category+":"+word, "count") + 1.0) / (hget(c, "word:"+category, "count") + hget(c, "TotalWords", "count"))
		p = p * wnc
	}

	return p
}

func categoryProb(c redis.Conn, category string) float64 {
	return hget(c, "word:"+category, "count") / hget(c, "TotalDocuments", "count")
}
