package global

type canalConf struct {
	ip string
	port int
	username string
	password string
	destination string
	soTimeOut int32
	idleTimeOut int32
	subscribe string
}

func newCanalConf(ip string, port int, username string, password string, destination string, soTimeOut int32, idleTimeOut int32, subscribe string) *canalConf {
	return &canalConf{ip: ip, port: port, username: username, password: password, destination: destination, soTimeOut: soTimeOut, idleTimeOut: idleTimeOut, subscribe: subscribe}
}

func (c *canalConf) Ip() string {
	return c.ip
}

func (c *canalConf) SetIp(ip string) {
	c.ip = ip
}

func (c *canalConf) Port() int {
	return c.port
}

func (c *canalConf) SetPort(port int) {
	c.port = port
}

func (c *canalConf) Username() string {
	return c.username
}

func (c *canalConf) SetUsername(username string) {
	c.username = username
}

func (c *canalConf) Password() string {
	return c.password
}

func (c *canalConf) SetPassword(password string) {
	c.password = password
}

func (c *canalConf) Destination() string {
	return c.destination
}

func (c *canalConf) SetDestination(destination string) {
	c.destination = destination
}

func (c *canalConf) SoTimeOut() int32 {
	return c.soTimeOut
}

func (c *canalConf) SetSoTimeOut(soTimeOut int32) {
	c.soTimeOut = soTimeOut
}

func (c *canalConf) IdleTimeOut() int32 {
	return c.idleTimeOut
}

func (c *canalConf) SetIdleTimeOut(idleTimeOut int32) {
	c.idleTimeOut = idleTimeOut
}

func (c *canalConf) Subscribe() string {
	return c.subscribe
}

func (c *canalConf) SetSubscribe(subscribe string) {
	c.subscribe = subscribe
}
