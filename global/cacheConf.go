package global

type cacheConf struct {
	host string
	password string
	db int
	maxOpenConns int
	maxIdleConns int
}

func newCacheConf(host string, password string, db int, maxOpenConns int, maxIdleConns int) *cacheConf {
	return &cacheConf{host: host, password: password, db: db, maxOpenConns: maxOpenConns, maxIdleConns: maxIdleConns}
}

func (c *cacheConf) Host() string {
	return c.host
}

func (c *cacheConf) SetHost(host string) {
	c.host = host
}

func (c *cacheConf) Password() string {
	return c.password
}

func (c *cacheConf) SetPassword(password string) {
	c.password = password
}

func (c *cacheConf) Db() int {
	return c.db
}

func (c *cacheConf) SetDb(db int) {
	c.db = db
}

func (c *cacheConf) MaxOpenConns() int {
	return c.maxOpenConns
}

func (c *cacheConf) SetMaxOpenConns(maxOpenConns int) {
	c.maxOpenConns = maxOpenConns
}

func (c *cacheConf) MaxIdleConns() int {
	return c.maxIdleConns
}

func (c *cacheConf) SetMaxIdleConns(maxIdleConns int) {
	c.maxIdleConns = maxIdleConns
}
