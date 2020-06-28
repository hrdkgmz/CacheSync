package cache

import (
	"github.com/gomodule/redigo/redis"
)

func (p *RedisPool) DoWitchDb(db int, command string, args ...interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return conn.Do(command, args...)
}

func (p *RedisPool) SetStringWitchDb(db int, key string, value interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return conn.Do("SET", key, value)
}

func (p *RedisPool) GetStringWitchDb(db int, key string) (string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return "", err
	}
	return redis.String(conn.Do("GET", key))
}

func (p *RedisPool) GetBytesWitchDb(db int, key string) ([]byte, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return redis.Bytes(conn.Do("GET", key))
}

func (p *RedisPool) GetIntWitchDb(db int, key string) (int, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return 0, err
	}
	return redis.Int(conn.Do("GET", key))
}

func (p *RedisPool) GetInt64WitchDb(db int, key string) (int64, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return 0, err
	}
	return redis.Int64(conn.Do("GET", key))
}

func (p *RedisPool) DelKeyWitchDb(db int, key string) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return conn.Do("DEL", key)
}

func (p *RedisPool) ExpireKeyWitchDb(db int, key string, seconds int64) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return conn.Do("EXPIRE", key, seconds)
}

func (p *RedisPool) KeysWitchDb(db int, pattern string) ([]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return redis.Strings(conn.Do("KEYS", pattern))
}

func (p *RedisPool) KeysByteSlicesWitchDb(db int, pattern string) ([][]byte, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return redis.ByteSlices(conn.Do("KEYS", pattern))
}

func (p *RedisPool) SetHashMapWitchDb(db int, key string, fieldValue map[string]interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fieldValue)...)
}

func (p *RedisPool) GetHashMapStringWitchDb(db int, key string) (map[string]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return redis.StringMap(conn.Do("HGETALL", key))
}

func (p *RedisPool) MgetHashMapStringWitchDb(db int, keys []string) (map[string]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return redis.StringMap(conn.Do("MGET", redis.Args{}.AddFlat(keys)))
}


func (p *RedisPool) GetHashMapIntWitchDb(db int, key string) (map[string]int, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return redis.IntMap(conn.Do("HGETALL", key))
}

func (p *RedisPool) GetHashStringWitchDb(db int, key string, field string) (string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return "", err
	}
	return redis.String(conn.Do("HGET", key, field))
}

func (p *RedisPool) GetHashIntWitchDb(db int, key string, field string) (int, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return 0, err
	}
	return redis.Int(conn.Do("HGET", key, field))
}

func (p *RedisPool) AddtoSetWitchDb(db int, key string, vals ...string) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	l := []interface{}{key}
	y := make([]interface{}, len(vals))
	for i, v := range vals {
		y[i] = v
	}
	l = append(l, y...)
	return conn.Do("SADD", l...)
}

func (p *RedisPool) RemoveFrmSetWitchDb(db int, key string, vals ...string) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	l := []interface{}{key}
	y := make([]interface{}, len(vals))
	for i, v := range vals {
		y[i] = v
	}
	l = append(l, y...)
	return conn.Do("SREM", l...)
}
func (p *RedisPool) GetSetMembersWitchDb(db int, key string) ([]interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return redis.Values(conn.Do("SMEMBERS", key))
}

func (p *RedisPool) GetSetInterWitchDb(db int, key1 string, key2 string) ([]interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return redis.Values(conn.Do("SINTER", key1, key2))
}

func (p *RedisPool) SetExpireWitchDb(db int, key string, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return conn.Do("EXPIRE", key, ttl)
}

func (p *RedisPool) SetPersistWitchDb(db int, key string) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return conn.Do("PERSIST", key)
}

func (p *RedisPool) GetTTLWitchDb(db int, key string) (int64, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return 0, err
	}
	return redis.Int64(conn.Do("TTL", key))
}

func (p *RedisPool) SetHashAndExpireWitchDb(db int, key string, fieldValue map[string]interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	i, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fieldValue)...)
	if err != nil {
		return i, err
	}
	return conn.Do("EXPIRE", key, ttl)
}

func (p *RedisPool) SetStringAndExpireWitchDb(db int, key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return conn.Do("SETEX", key, ttl, val)
}

func (p *RedisPool) AppendListToHeadWitchDb(db int, key string, val interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	var i interface{}
	if v, ok := val.([]string); ok {
		i, err = conn.Do("LPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("LPUSH", key, val)
	}
	return i, err
}

func (p *RedisPool) SetListToHeadWitchDb(db int, key string, val interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	var i interface{}
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

func (p *RedisPool) AppendListToHeadAndExpireWitchDb(db int, key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	var i interface{}
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

func (p *RedisPool) SetListToHeadAndExpireWitchDb(db int, key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	var i interface{}
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

func (p *RedisPool) AppendListToTailWitchDb(db int, key string, val interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	var i interface{}
	if v, ok := val.([]string); ok {
		i, err = conn.Do("RPUSH", redis.Args{}.Add(key).AddFlat(v)...)
	} else {
		i, err = conn.Do("RPUSH", key, val)
	}
	return i, err
}

func (p *RedisPool) SetListToTailWitchDb(db int, key string, val interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	var i interface{}
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

func (p *RedisPool) AppendListToTailAndExpireWitchDb(db int, key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	var i interface{}
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

func (p *RedisPool) SetListToTailAndExpireWitchDb(db int, key string, val interface{}, ttl int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	var i interface{}
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

func (p *RedisPool) GetListLenWitchDb(db int, key string) (int, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return 0, err
	}
	return redis.Int(conn.Do("LLEN", key))
}

func (p *RedisPool) GetListByRangeWitchDb(db int, key string, start int, stop int) ([]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_,err:=conn.Do("select", db)
	if err != nil {
		return nil, err
	}
	return redis.Strings(conn.Do("LRANGE", key, start, stop))
}
