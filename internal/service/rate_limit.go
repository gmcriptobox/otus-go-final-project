package service

import (
	"time"

	"golang.org/x/time/rate"
)

type BucketWithLastEventTime struct {
	limiter       *rate.Limiter
	lastEventTime time.Time // Время последнего обращения к бакету
}

type RateLimit struct {
	limit     int
	bucketTTL int // Время жизни бакета с момента последнего обращения к нему (в секундах)
	bucketMap map[string]*BucketWithLastEventTime
}

func NewRateLimit(limit int, bucketTTL int) *RateLimit {
	rateLimit := &RateLimit{
		limit:     limit,
		bucketTTL: bucketTTL,
		bucketMap: make(map[string]*BucketWithLastEventTime),
	}
	go deleteUnusedBuckets(rateLimit)
	return rateLimit
}

func (r *RateLimit) Allow(key string) bool {
	bucket, ok := r.bucketMap[key]
	if !ok {
		bucket = newBucket(r.limit)
		r.bucketMap[key] = bucket
	}

	// Записываем время последнего обращения к бакету для последующего удаления неиспользуемых бакетов
	bucket.lastEventTime = time.Now()
	return bucket.limiter.Allow()
}

func (r *RateLimit) ResetBucket(key string) {
	delete(r.bucketMap, key)
}

func deleteUnusedBuckets(r *RateLimit) {
	ticker := time.NewTicker(60 * time.Second)
	for {
		<-ticker.C
		for key, bucket := range r.bucketMap {
			if time.Since(bucket.lastEventTime) > time.Duration(r.bucketTTL)*time.Second {
				delete(r.bucketMap, key)
			}
		}
	}
}

func newBucket(limit int) *BucketWithLastEventTime {
	return &BucketWithLastEventTime{
		limiter: rate.NewLimiter(rate.Limit(float64(limit)/time.Duration.Seconds(60*time.Second)), limit),
	}
}
