package websocket

import (
  "net/http"
  "github.com/bitly/go-simplejson"
  "github.com/gorilla/websocket"
)


type WebsocketModule struct {
  hub   *Hub
  conn  *Conn
}

type Hub struct {
  connections map[*connection]bool  // Registered connections.
  broadcast chan []byte             // Inbound messages from the connections.
  register chan *connection         // Register requests from the connections.
  unregister chan *connection       // Unregister requests from connections.
}

type Conn struct {
  ws *websocket.Conn  // The websocket connection.
  send chan []byte    // Buffered channel of outbound messages.
}

type MessageHandler func(*Conn, Message)

type Message struct {
  Event string
  Payload []byte
}

type route struct {
  topic    string
  callback MessageHandler
}

type router struct {
  sync.RWMutex
  routes         *list.List
  defaultHandler MessageHandler
  messages       chan *packets.PublishPacket
  stop           chan bool
}

func NewWebsocketModule() *WebsocketModule {
  // Revisar!!!
  go Run()
  return &MqttModule
}


func (a *WebsocketModule) On(event string, f func(s []byte)) bool {
  if a.conn.ws == nil {
    return false
  }
  a.Event(event, func(conn *Conn, msg Message) {
    f(msg.Payload())
  })
  return true
}


func (h *Hub) Run() {
  for {
    select {
    case c := <-h.register:
      h.connections[c] = true
    case c := <-h.unregister:
      if _, ok := h.connections[c]; ok {
        log.Println("Websocket: cliente cerrado:", c)
        delete(h.connections, c)
        close(c.send)
      }
    case m := <-h.broadcast:
      for c := range h.connections {
        select {
        case c.send <- m:
        default:
          close(c.send)
          delete(h.connections, c)
        }
      }
    }
  }
}






func onEvent(msg []byte) {

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
}

// recvLoop: Loop de manejo de recepcion de mensajes.
func (a *WebsocketModule) recvLoop(callback MessageHandler) {
  // Elimino el cliente cuando se desconecta
  defer func() {
    a.hub.unregister <- c
    a.conn.ws.Close()
  }()
  
  for {
    _, msg, err := a.conn.ws.ReadMessage()
    if err != nil {
      break
    }
    //h.broadcast <- msg
    onEvent(msg) // Proseso el mensaje recivido
  }
}



  

var H = Hub {
  connections: make(map[*connection]bool),
  broadcast:   make(chan []byte),
  register:    make(chan *connection),
  unregister:  make(chan *connection),
}

func (h *Hub) Run() {
  for {
    select {
    case c := <-h.register:
      h.connections[c] = true
    case c := <-h.unregister:
      if _, ok := h.connections[c]; ok {
        log.Println("Websocket: cliente cerrado:", c)
        delete(h.connections, c)
        close(c.send)
      }
    case m := <-h.broadcast:
      for c := range h.connections {
        select {
        case c.send <- m:
        default:
          close(c.send)
          delete(h.connections, c)
        }
      }
    }
  }
}

var upgrader = websocket.Upgrader{
  ReadBufferSize:  2048,
  WriteBufferSize: 2048,
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
  ws *websocket.Conn  // The websocket connection.
  send chan []byte    // Buffered channel of outbound messages.
}

// write: Envio el mensaje con el tipo y el payload.
func (c *connection) write(mt int, payload []byte) error {
  return c.ws.WriteMessage(mt, payload)
}

// sendLoop: Loop de Manejo de envío de mensajes.
func (c *connection) sendLoop() {
  for {
    select {
    case msg, ok := <-c.send:
      if !ok {
        c.write(websocket.CloseMessage, []byte{})
        return
      }
      if err := c.write(websocket.TextMessage, msg); err != nil {
        return
      }
    }
  }
}

// recvLoop: Loop de manejo de recepcion de mensajes.
func (c *connection) recvLoop() {
  // Elimino el cliente cuando se desconecta
  defer func() {
    H.unregister <- c
    c.ws.Close()
  }()
  
  for {
    _, msg, err := c.ws.ReadMessage()
    if err != nil {
      break
    }
    //h.broadcast <- msg
    onEvent(msg) // Proseso el mensaje recivido
  }
}



