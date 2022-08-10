package main

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"sync"
	"time"
)

// 缓存内容信息的结构
type cacheContent struct {
	content     []byte
	expiresTime int64
}

//localCache 缓存结构
type localCache struct {
	stop  chan struct{}
	wg    sync.WaitGroup
	mu    sync.RWMutex            //读写锁
	cache map[string]cacheContent //缓存信息
}

//NewLocalCache 创建缓存对象
func newLocalCache(cleanInterval time.Duration) *localCache {
	lc := &localCache{
		cache: make(map[string]cacheContent),
		stop:  make(chan struct{}),
	}
	lc.wg.Add(1)
	go func(interval time.Duration) {
		defer lc.wg.Done()
		lc.cleanUp(interval)
	}(cleanInterval)
	return lc
}

//定期清除过期对象
func (lc *localCache) cleanUp(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for key, c := range lc.cache {
				if c.expiresTime <= time.Now().Unix() {
					//删除缓存
					delete(lc.cache, key)
				}
			}
			lc.mu.Unlock()
		}
	}
}

func (lc *localCache) stopCleanUp() {
	close(lc.stop)
	lc.wg.Wait()
}

//Update 更新缓存
func (lc *localCache) update(key string, content []byte, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	//内容压缩
	compress := compress(content)
	lc.cache[key] = cacheContent{compress, expireAtTimestamp}
}

//Read 读取缓存
func (lc *localCache) read(key string) ([]byte, error) {
	lc.mu.RLock()
	content, ok := lc.cache[key]
	lc.mu.RUnlock()
	if !ok {
		return nil, errors.New("cache content not find")
	}
	if content.expiresTime <= time.Now().Unix() {
		lc.delete(key)
		return nil, errors.New("cache content not find")
	}
	return unCompress(content.content), nil
}

//删除内容
func (lc *localCache) delete(key string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	delete(lc.cache, key)
}

//压缩[]byte
func compress(data []byte) []byte {
	var out bytes.Buffer
	g := gzip.NewWriter(&out)
	_, _ = g.Write(data)
	_ = g.Close()
	return out.Bytes()
}

//解压[]byte
func unCompress(data []byte) []byte {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(data)
	r, _ := gzip.NewReader(&in)
	_ = r.Close()
	_, _ = io.Copy(&out, r)
	return out.Bytes()
}
