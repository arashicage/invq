package main

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
	"gopkg.in/macaron.v1"
)

// redis 连接池
// setting
// cron 系统状态 连接数 系统压力等

func main() {

	m := macaron.Classic()

	//	默认为静态 dir 为 public 目录
	m.Use(macaron.Static("public"))
	//	设置模板目录,模板后缀
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
	}))

	m.Get("/", func(ctx *macaron.Context) {

		ctx.Data["servers"] = string(main1())

		ctx.HTML(200, "index")

	})

	m.Run(8080)
}

func main1() []byte {
	defer func() {
		if e := recover(); e != nil {
			log.Println("error occers")
		}
	}()

	//  机动车
	//     sj not null
	//     zj nullable
	//
	// kprq cyr_ybsbh swjg_dm

	connectTimeout := time.Duration(3) * time.Second
	readTimeout := time.Duration(3) * time.Second
	writeTimeout := time.Duration(3) * time.Second

	url := "127.0.0.1:6379"
	rs, err := redis.DialTimeout("tcp", url, connectTimeout, readTimeout, writeTimeout)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer rs.Close()

	v, err := redis.Values(rs.Do("HGETALL", "04:key000000027"))
	if err != nil {
		panic(err)
	}

	m, err := redis.StringMap(v, err)
	var j []byte
	j, err = ffjson.Marshal(&m)
	return j

}
