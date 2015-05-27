package models
import (
    "log"
  //  "strings"
    "fmt"
)

type Topics struct {
    Id        int64
    Index     int
    Title     string //标题(分类: 健康,统计的,)
    Content   string //详细内容
    Answer    string
//    CreatedAt int64 `xorm:"created"`
//    UpdatedAt int64 `xorm:"updated"`
}

func GetTopics(index int)[]Topics{
    topics:=make([]Topics,0)
    err:=engine.Where("`Index`=?",index).Find(&topics)
    if err!=nil{
        log.Fatal(err)
    }
//    for k,v:=range topics{
//        topics[k].Content=strings.Replace( v.Content,"\n","<br>",-1)
//    }
    return  topics
}

func GetOne(index int)*Topics{
    topic:=new(Topics)
    _,err:=engine.Where("`Index`=?",index).Get(topic)
    if err!=nil{
        log.Fatal(err)
    }
    fmt.Println("topic:",topic)
    return  topic
}