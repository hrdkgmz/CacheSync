package cache

import "github.com/gomodule/redigo/redis"

func (p *RedisPool) Close() error {
	err := p.Pool.Close()
	return err
}

func (p *RedisPool) Do(command string, args ...interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do(command, args...)
}

func (p *RedisPool) SetString(key string, value interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do("SET", key, value)
}

func (p *RedisPool) GetString(key string) (string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

func (p *RedisPool) GetBytes(key string) ([]byte, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

func (p *RedisPool) GetInt(key string) (int, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))
}

func (p *RedisPool) GetInt64(key string) (int64, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("GET", key))
}

func (p *RedisPool) DelKey(key string) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do("DEL", key)
}

func (p *RedisPool) ExpireKey(key string, seconds int64) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do("EXPIRE", key, seconds)
}

func (p *RedisPool) Keys(pattern string) ([]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", pattern))
}

func (p *RedisPool) KeysByteSlices(pattern string) ([][]byte, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.ByteSlices(conn.Do("KEYS", pattern))
}

func (p *RedisPool) SetHashMap(key string, fieldValue map[string]interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fieldValue)...)
}

func (p *RedisPool) GetHashMapString(key string) (map[string]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", key))
}

func (p *RedisPool) GetHashMapInt(key string) (map[string]int, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.IntMap(conn.Do("HGETALL", key))
}

func (p *RedisPool) GetHashString(key string, field string) (string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("HGET", key, field))
}

func (p *RedisPool) GetHashInt(key string, field string) (int, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("HGET", key, field))
}

func (p *RedisPool) AddtoSet(key string, vals ...string) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	l := []interface{}{key}
	y := make([]interface{}, len(vals))
	for i, v := range vals {
		y[i] = v
	}
	l = append(l, y...)
	return conn.Do("SADD", l...)
}

func (p *RedisPool) GetSetMembers(key string) ([]interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Values(conn.Do("SMEMBERS", key))
}

func (p *RedisPool) GetSetInters(key1 string, key2 string) ([]interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Values(conn.Do("SINTER", key1, key2))
}
