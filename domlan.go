package main

import (
  "flag"
  "./server"
  //"net/http"
  //"io/ioutil"
  "github.com/gin-gonic/gin"
  "github.com/ljgago/glue"
  //"github.com/desertbit/glue"
  //"github.com/elazarl/go-bindata-assetfs"
  //"github.com/gin-gonic/contrib/static"
)

func main() {
  //host := flag.String("host", "0.0.0.0", "IP of host to run webserver on")
  port := flag.String("port", "3000", "Port to run webserver on")
  debug := flag.Bool("debug", true, "Enable the debug log")
  //webcam := flag.Bool("webcam", true, "Enable webcam")
  flag.Parse()

  server.DEBUG_ENABLE = *debug
  
  server.Dev.DeviceWorks()
  //go server.Ruta.Run()
  glue.RunRouter()

  r := gin.New()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())
  // Sirvo los archivos que estan en el sistema
  r.Static("/assets", "./client")
  // Utilizo go-bindata para empotrar los archivos dentro del binario
  //r.ServeFiles("/assets/*filepath", assetFS())
  //r.StaticBinData("/assets", assetFS())
  r.GET("/", func(c *gin.Context) {
    //obj := gin.H{"title": "Domlan"}
    //text, _ := ioutil.ReadFile("./client/views/index.html")
    //c.HTML(http.StatusOK, string(text), gin.H{"title": "Main website",})
    c.File("./client/views/index.html")
  })
  r.GET("/ws", func(c *gin.Context) {
    //server.InitWS(c.Writer, c.Request)
    glue.OnNewSocket(server.OnDataReceive)
  })
  // Escucho por defaul en la dirección 0.0.0.0:3000
  r.Run(":" + *port)
}