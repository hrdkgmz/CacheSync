package cache

import (
	"github.com/gomodule/redigo/redis"
)

func (p *RedisPool) Close() error {
	err := p.Pool.Close()
	return err
}

func (p *RedisPool) Do(command string, args ...interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do(command, args...)
}

func (p *RedisPool) Select(db int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do("select", db)
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

func (p *RedisPool) RemoveFrmSet(key string, vals ...string) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	l := []interface{}{key}
	y := make([]interface{}, len(vals))
	for i, v := range vals {
		y[i] = v
	}
	l = append(l, y...)
	return conn.Do("SREM", l...)
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

func (p *RedisPool) SetExpire(key string, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do("EXPIRE", key, ttl)
}

func (p *RedisPool) SetPersist(key string) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do("PERSIST", key)
}

func (p *RedisPool) GetTTL(key string) (int64, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("TTL", key))
}

func (p *RedisPool) SetHashAndExpire(key string, fieldValue map[string]interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	i, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fieldValue)...)
	if err != nil {
		return i, err
	}
	return conn.Do("EXPIRE", key, ttl)
}

func (p *RedisPool) SetStringAndExpire(key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return conn.Do("SETEX", key, ttl, val)
}

func (p *RedisPool) AppendListToHead(key string, val interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	var i interface{}
	var err error
	if v, ok := val.([]string); ok {
		i, err = conn.Do("LPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("LPUSH", key, val)
	}
	return i, err
}

func (p *RedisPool) SetListToHead(key string, val interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	var i interface{}
	var err error
	v, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return nil, err
	}
	if v == true {
		i, err = conn.Do("DEL", key)
		if err != nil {
			return i, err
		}
	}
	if v, ok := val.([]string); ok {
		i, err = conn.Do("LPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("LPUSH", key, val)
	}
	return i, err
}

func (p *RedisPool) AppendListToHeadAndExpire(key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	var i interface{}
	var err error
	if v, ok := val.([]string); ok {
		i, err = conn.Do("LPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("LPUSH", key, val)
	}
	if err != nil {
		return i, err
	}
	return conn.Do("EXPIRE", key, ttl)
}

func (p *RedisPool) SetListToHeadAndExpire(key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	var i interface{}
	var err error
	v, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return nil, err
	}
	if v == true {
		i, err = conn.Do("DEL", key)
		if err != nil {
			return i, err
		}
	}
	if v, ok := val.([]string); ok {
		i, err = conn.Do("LPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("LPUSH", key, val)
	}
	if err != nil {
		return i, err
	}
	return conn.Do("EXPIRE", key, ttl)
}

func (p *RedisPool) AppendListToTail(key string, val interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	var i interface{}
	var err error
	if v, ok := val.([]string); ok {
		i, err = conn.Do("RPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("RPUSH", key, val)
	}
	return i, err
}

func (p *RedisPool) SetListToTail(key string, val interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	var i interface{}
	var err error
	v, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return nil, err
	}
	if v == true {
		i, err = conn.Do("DEL", key)
		if err != nil {
			return i, err
		}
	}
	if v, ok := val.([]string); ok {
		i, err = conn.Do("RPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("RPUSH", key, val)
	}
	return i, err
}

func (p *RedisPool) AppendListToTailAndExpire(key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	var i interface{}
	var err error
	if v, ok := val.([]string); ok {
		i, err = conn.Do("RPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("RPUSH", key, val)
	}
	if err != nil {
		return i, err
	}
	return conn.Do("EXPIRE", key, ttl)
}

func (p *RedisPool) SetListToTailAndExpire(key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	var i interface{}
	var err error
	v, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return nil, err
	}
	if v == true {
		i, err = conn.Do("DEL", key)
		if err != nil {
			return i, err
		}
	}
	if v, ok := val.([]string); ok {
		i, err = conn.Do("RPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("RPUSH", key, val)
	}
	if err != nil {
		return i, err
	}
	return conn.Do("EXPIRE", key, ttl)
}

func (p *RedisPool) GetListLen(key string) (int, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("LLEN", key))
}

func (p *RedisPool) GetListByRange(key string, start int, stop int) ([]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("LRANGE", key, start, stop))
}
