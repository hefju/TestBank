package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/hefju/TestBank/models"
    "strconv"
)



func main() {

    router := gin.Default()
    router.LoadHTMLGlob("views/*")
    router.GET("/", func(c *gin.Context) { //测试，获取数据表信息
        // log.Println("visit homepage\r\n")
       // logger.Debug("visit homepage")
        c.String(http.StatusOK, "Welcome to TestBank...")
    })
    router.GET("/index",func(c *gin.Context){
       //
        x:=models.GetTopics(3)
        obj:=gin.H{"Topics": x}
        c.HTML(http.StatusOK,"index.html",obj)
        // c.HTMLString(http.StatusOK,"index.html",obj)

    })

    router.GET("/one/:index",func(c *gin.Context){
        str := c.Params.ByName("index")
      //  fmt.Println("str:",str)
        index,_:=strconv.Atoi(str)
       // fmt.Println("x:",index)
        x:=models.GetOne( index   )
       // fmt.Println("x:",x)
        c.JSON(http.StatusOK, gin.H{"topic": x})

    })


    router.Run(":8086")
    fmt.Println("end")
}

