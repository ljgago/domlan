package server

import (
  "log"
  "net/http"
  "github.com/gorilla/websocket"
  "github.com/bitly/go-simplejson"
)

// Client: is an middleman between the websocket connection and the Router.
//
// Client: es un intermediaron entre la conexción websocket y el Router.
type Client struct {
  ws *websocket.Conn  // The websocket connection.
  send chan []byte    // Buffered channel of outbound messages.
}

// Router: structure that hold the set active conections and broadcast messages.
//
// Router: estructura que mantiene el conjunto de conexciones activas
// y los mensajes broadcast.
type Router struct {
  clients     map[*Client]bool      // Registered clients.
  broadcast   chan []byte         // Inbound messages from the clients.
  register    chan *Client         // Register requests from the clients.
  unregister  chan *Client       // Unregister requests from clients.
}

var Ruta = Router{
  clients:    make(map[*Client]bool),
  broadcast:  make(chan []byte),
  register:   make(chan *Client),
  unregister: make(chan *Client),
}

var upgrader = websocket.Upgrader{
  ReadBufferSize:  2048,
  WriteBufferSize: 2048,
}

// Run: gorounite that handle to clients conections.
//
// Run: goroutine que maneja a las conexciones de los clientes.
func (r *Router) Run() {
  for {
    select {
    case c := <-r.register:
      r.clients[c] = true
    case c := <-r.unregister:
      if _, ok := r.clients[c]; ok {
        DEBUG("Websocket: cliente cerrado:", c)
        delete(r.clients, c)
        close(c.send)
      }
    case m := <-r.broadcast:
      for c := range r.clients {
        select {
        case c.send <- m:
        default:
          close(c.send)
          delete(r.clients, c)
        }
      }
    }
  }
}

// SendLoop: loop that management messaging.
//
// sendLoop: loop de Manejo de envío de mensajes.
func (c *Client) sendLoop() {
  for {
    select {
    case msg, ok := <-c.send:
      if !ok {
        c.sendMsg(websocket.CloseMessage, []byte{})
        return
      }
      if err := c.sendMsg(websocket.TextMessage, msg); err != nil {
        return
      }
    }
  }
}

// sendMsg: it send a message with the type and the payload.
//
// sendMsg: envío el mensaje con el tipo y el payload.
func (c *Client) sendMsg(mt int, payload []byte) error {
  return c.ws.WriteMessage(mt, payload)
}

// recvLoop: loop that handle the reception of menssages.
//
// recvLoop: loop que maneja la recepcion de mensajes.
func (c *Client) recvLoop() {
  // Elimino el cliente cuando se desconecta
  defer func() {
    Ruta.unregister <- c
    c.ws.Close()
  }()
  
  for {
    _, msg, err := c.ws.ReadMessage()
    if err != nil {
      break
    }
    //h.broadcast <- msg
    OnEvent(msg) // Proseso el mensaje recivido
  }
}

// Emit: send a message to specified client.
//
// Emit: mando un mensaje al cliente especificado.
func (c *Client) Emit(event string, msg string) {
  str := `{"event":` + event + `,"data":` + msg + `}`
  c.send <- []byte(str)
}

// EmitBroadcast: it send a message to all clients connected.
//
// EmitBroadcast: mando un mensajes a todos los clientes conectados.
func (r *Router) EmitBroadcast(event string, msg string) {
  
  str := `{"event":` + event + `,"data":"` + msg + `"}`
  //log.Println(str)
  r.broadcast <- []byte(str)
}

// OnEvent: it receive a message from the a client.
//
// OnEvent: se recive un mensaje de el cliente.
func OnEvent(msg []byte) {

  js, err := simplejson.NewJson(msg)
  if err != nil {
    log.Println("JSON:", err)
    return
  }
  event := js.Get("event").MustString()
  switch event {
    case "add-device":
      value, _ := js.Get("data").MarshalJSON()
      if err != nil {
        log.Println(err)
      }
      //key := js.Get("data").Get("id").MustString()
      //PutDevice([]byte(key), value) {
      log.Println(string(value))

      //a := dataMerge(msg, []byte(test))
      //log.Println(string(a))
    case "update-device":
    case "delete-device":
    default:
      log.Printf("Websockets: No se encontro el evento: \"%s\"", event)
  }
  /*
  value, _ := js.EncodePretty()
  if err != nil {
    log.Println(err)
    return
  }*/
}


// InitWS: handler websocket que maneja los requests.
func InitWS(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, "Method not allowed", 405)
    return
  }
  ws, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println(err)
    return
  }
  c := &Client{send: make(chan []byte, 256), ws: ws}
  Ruta.register <- c
  
  // Leo la base de datos y se la transmito al cliente
  data := ldb.GetAllDeviceDB()
  log.Println("Datos:", data)
  c.Emit(`"start"`, data)
  
  go c.sendLoop()
  c.recvLoop()
}


/*
go func () {
  goreq.SetConnectTimeout(10 * time.Hour)
  for {
    log.Println("entro")
    res, err := goreq.Request{
      Uri: "http://192.168.1.111/jpg/image.jpg",
      Timeout: 10 * time.Hour,
    }.Do()

    //res, err := climg.Get("http://192.168.1.111/jpg/image.jpg")

    if err != nil {
        log.Printf("http.Get -> %v", err)
        continue
    }

    // We read all the bytes of the image
    // Types: data []byte
    data, err := ioutil.ReadAll(res.Body)

    if err != nil {
        log.Printf("ioutil.ReadAll -> %v", err)
        continue
    }
    //log.Println(data)

    // You have to manually close the body, check docs
    // This is required if you want to use things like
    // Keep-Alive and other HTTP sorcery.
    res.Body.Close()

    // Encode los datos en base64
    uEnc := base64.StdEncoding.EncodeToString(data)
    //log.Println(string(uEnc))
    emitBroadcast(`"webcam"`, string(uEnc), 1)

    time.Sleep(time.Millisecond * 500)
    log.Println("1")
  }
}()
*/
