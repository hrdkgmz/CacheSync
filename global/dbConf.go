package global

type dbConf struct {
	host string
	database string
	username string
	password string
	charset string
	maxOpenConns int
	maxIdleConns int
}

func newDbConf(host string, database string, username string, password string, charset string, maxOpenConns int, maxIdleConns int) *dbConf {
	return &dbConf{host: host, database: database, username: username, password: password, charset: charset, maxOpenConns: maxOpenConns, maxIdleConns: maxIdleConns}
}

func (d *dbConf) Host() string {
	return d.host
}

func (d *dbConf) SetHost(host string) {
	d.host = host
}

func (d *dbConf) Database() string {
	return d.database
}

func (d *dbConf) SetDatabase(database string) {
	d.database = database
}

func (d *dbConf) Username() string {
	return d.username
}

func (d *dbConf) SetUsername(username string) {
	d.username = username
}

func (d *dbConf) Password() string {
	return d.password
}

func (d *dbConf) SetPassword(password string) {
	d.password = password
}

func (d *dbConf) Charset() string {
	return d.charset
}

func (d *dbConf) SetCharset(charset string) {
	d.charset = charset
}

func (d *dbConf) MaxOpenConns() int {
	return d.maxOpenConns
}

func (d *dbConf) SetMaxOpenConns(maxOpenConns int) {
	d.maxOpenConns = maxOpenConns
}

func (d *dbConf) MaxIdleConns() int {
	return d.maxIdleConns
}

func (d *dbConf) SetMaxIdleConns(maxIdleConns int) {
	d.maxIdleConns = maxIdleConns
}