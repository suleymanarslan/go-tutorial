package redis

import (
	"github.com/garyburd/redigo/redis"
	"hoditgo/settings"
)

type RedisCli struct {
	conn redis.Conn
}

var instanceRedisCli *RedisCli = nil

func Connect() (conn *RedisCli) {
	if instanceRedisCli == nil {
		instanceRedisCli = new(RedisCli)
		var err error

		instanceRedisCli.conn, err = redis.Dial("tcp",  settings.Get().RedisPort)

		if err != nil {
			panic(err)
		}

		if _, err := instanceRedisCli.conn.Do("AUTH", settings.Get().RedisPassword); err != nil {
			instanceRedisCli.conn.Close()
			panic(err)
		}
	}

	return instanceRedisCli
}

func (redisCli *RedisCli) SetValue(key string, value string, expiration ...interface{}) error {
	_, err := redisCli.conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}

	return err
}

func (redisCli *RedisCli) GetValue(key string) (interface{}, error) {
	return redisCli.conn.Do("GET", key)
}

func (redisCli *RedisCli) AddUserToRoom(userName string, roomName string)  {
	 redisCli.conn.Do("LPUSH", roomName, userName)
}

func (redisCli *RedisCli) RoomExists(userName string, roomName string ) (interface{}, error){
	return redisCli.conn.Do("EXISTS", roomName)
}


