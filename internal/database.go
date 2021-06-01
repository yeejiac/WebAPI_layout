package internal

func ConnRedis_pool()
{
	pool := &redis.Pool{
		// 連線的 callback 定義
		Dial: func() (redis.Conn, error) {

			//建構一條連線
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			//在這邊可以做連線池初始化 選擇 redis db的動作
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},

		//定期對 redis server 做 ping/pong 測試
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}