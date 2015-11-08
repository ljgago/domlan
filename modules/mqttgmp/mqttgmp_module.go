package mqttgmp

import (
  "log"
  "github.com/yosssi/gmq/mqtt"
  "github.com/yosssi/gmq/mqtt/client"
)

type MqttModule struct {
  user string
  addr string
  id string
  cli *client.Client
}

// NewMqttModule Create a new module MQTT with specified name, host and id
// NewMqttModule crea un nuevo módulo mqtt con el nombre, host e id especificado
func NewMqttModule(user string, addr string, id string) *MqttModule {
  
  cli := client.New(&client.Options{
    ErrorHandler: func(err error) {
        log.Println(err)
    },
  })
  return &MqttModule{
    user: user,
    addr: addr,
    id: id,
    cli: cli,
  }
}

// Connect return true if the conection is established
// Connect retorna true si la conexión a mqtt es establecida
func (a *MqttModule) Connect() (errs []error) {
  
  err := a.cli.Connect(&client.ConnectOptions{
    CleanSession: false,
    //KeepAlive:    30,
    Network:      "tcp",
    Address:      a.addr,
    ClientID:     []byte(a.id),
  })

  if err != nil {
    log.Println("Error:", err)
  } 
  return
}

// On subscribe to topic specified by event and executes the anonymous function f
// On se subscribe al topic especificado por event y ejecuta la funcion anonima f
func (a *MqttModule) On(event string, f func(s []byte)) bool {
  err := a.cli.Subscribe(&client.SubscribeOptions{
    SubReqs: []*client.SubReq{
      &client.SubReq{
        TopicFilter: []byte(event),
        QoS:         mqtt.QoS0,
        // Define the processing of the message handler.
        Handler: func(topicName, message []byte) {
          f(message)
        },
      },
    },
  })

  if err != nil {
    log.Println("Error:", err)
  }
  return true
}

func (a *MqttModule) Publish(topicName, message string ) {

  err := a.cli.Publish(&client.PublishOptions{
    // QoS is the QoS of the PUBLISH Packet.
    QoS:       mqtt.QoS0,
    // Retain is the Retain of the PUBLISH Packet.
    Retain:    true,
    // TopicName is the Topic Name of the PUBLISH Packet.
    TopicName: []byte(topicName),
    // Message is the Application Message of the PUBLISH Packet.
    Message:   []byte(message),
  })
  if err != nil {
    panic(err)
  }
  return
}
