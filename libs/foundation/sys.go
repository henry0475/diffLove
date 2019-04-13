package foundation

import (
	"bytes"
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	configure "code.aliyun.com/liwenbin0320/alumnusAssociation_api/config"
	"github.com/go-redis/redis"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

var httpClient *http.Client
var redisClient *redis.Client
var mysqlClient *sql.DB

var conf configure.Config

const (
	// MaxIdleConns means the maximum keep-alive's links.
	MaxIdleConns int = 10
	// MaxIdleConnsPerHost means (keep-alive) connections to keep per-host.
	MaxIdleConnsPerHost int = 10
	// IdleConnTimeout means (keep-alive) connection will remain idle before closing itself.
	IdleConnTimeout int = 30
)

func init() {
	log.Println("init foundation")

	config.Load(file.NewSource(
		file.WithPath(os.Getenv("GOPATH") + "/src/code.aliyun.com/liwenbin0320/alumnusAssociation_api/config/config.json"),
	))
	config.Scan(&conf)

	httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
		Timeout: 20 * time.Second,
	}

	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     conf.Hosts.Cache.Address + ":" + strconv.Itoa(conf.Hosts.Cache.Port),
			Password: conf.Hosts.Cache.Auth,
			DB:       0, // use default DB
			PoolSize: 3,
		},
	)
	pong, err := redisClient.Ping().Result()
	if pong != "PONG" {
		log.Print("Linking redis failed, because: ", err.Error())
		return
	}
}

// GetMysqlClient will return a client for mysql
func GetMysqlClient() *sql.DB {
	if mysqlClient != nil {
		return mysqlClient
	}
	log.Println("recreate a mysql client")
	mClient, err := sql.Open("mysql", conf.Hosts.Database.User+":"+conf.Hosts.Database.Password+"@tcp("+conf.Hosts.Database.Address+":"+strconv.Itoa(conf.Hosts.Database.Port)+")/"+conf.Hosts.Database.Db+"?charset="+conf.Hosts.Database.Charset)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	mClient.SetMaxOpenConns(5)
	mClient.SetMaxIdleConns(3)
	//10 mins
	mClient.SetConnMaxLifetime(10 * 60 * 1000)

	mysqlClient = mClient
	return mysqlClient
}

// GetHTTPRequestHandle will return a handle for HTTP request
func GetHTTPRequestHandle(url string, method string, content []byte, headerMap map[string]string) (req *http.Request) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(content))
	if err != nil {
		log.Println(err.Error())
		return
	}
	for head, value := range headerMap {
		req.Header.Set(head, value)
	}
	return
}

// GetRedisClient will return a client for redis
func GetRedisClient() *redis.Client {
	return redisClient
}

// GetHTTPClient will return a client for HTTP
func GetHTTPClient() *http.Client {
	return httpClient
}

// GetConf will return conf handle
func GetConf() configure.Config {
	return conf
}

// FilterSensitiveCharacters will filter some sensitive characters in case of a special case happened
// The result could be different depends on the parameter of level
// level 0 means there is no changes (as default)
// level 1 means only ' will be omitted
// level 2 means some characters could be omitted
func FilterSensitiveCharacters(str string, level int8) string {
	switch level {
	case 0:
		return str
	case 1:
		return strings.Replace(str, "'", "", -1)
	case 2:
		str = strings.Replace(str, "'", "", -1)
		str = strings.Replace(str, "&", "", -1)
		str = strings.Replace(str, "\"", "", -1)
		return str
	default:
		return str
	}
}
