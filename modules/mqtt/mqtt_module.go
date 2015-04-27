package mqttpaho

import (
  "log"
  "time"
  mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

// MqttModule estrutura de datos
type MqttModule struct {
  name     string
  Host     string
  clientID string
  client   *mqtt.Client
}

// NewMqttModule Create a new module MQTT with specified name, host and id
// NewMqttModule crea un nuevo módulo mqtt con el nombre, host e id especificado
func NewMqttModule(name string, host string, clientID string) *MqttModule {
  return &MqttModule{
    name:     name,
    Host:     host,
    clientID: clientID,
  }
}
// Name devuleve el nombre
func (a *MqttModule) Name() string {
  return a.name
}

// Connect return true if the conection is established
// Connect retorna true si la conexión a mqtt es establecida
func (a *MqttModule) Connect() (errs []error) {
  a.client = mqtt.NewClient(createClientOptions(a.clientID, a.Host))
  if token := a.client.Connect(); token.Wait() && token.Error() != nil {
    errs = append(errs, token.Error())
    log.Println(errs)
  }
  return
}

// Disconnect retorna true si la conexión a mqtt es cerrada
func (a *MqttModule) Disconnect() (err error) {
  if a.client != nil {
    a.client.Disconnect(500)
  }
  return
}

// Finalize retorna true si la conexión a mqtt finalizo correctamente
func (a *MqttModule) Finalize() (errs []error) {
  a.Disconnect()
  return
}

// Publish publica un mensaje con el topic específico
func (a *MqttModule) Publish(topic string, message []byte) bool {
  if a.client == nil {
    return false
  }
  a.client.Publish(topic, 0, false, message)
  return true
}

// On Subscribe a un topic, y llama a la función handler cuando un dato es recivido
func (a *MqttModule) On(event string, f func(s []byte)) bool {
  if a.client == nil {
    return false
  }
  // Guardo en una lista los string de eventos
  a.client.Subscribe(event, 0, func(client *mqtt.Client, msg mqtt.Message) {
    f(msg.Payload())
  })
  return true
}

// createClientOptions configura las opciones del cliente
func createClientOptions(clientID, raw string) *mqtt.ClientOptions {
  opts := mqtt.NewClientOptions()
  opts.SetCleanSession(false)
  opts.SetAutoReconnect(true)
  opts.SetKeepAlive(30 * time.Second)
  //opts.SetMaxReconnectInterval(10 * time.Second)
  opts.AddBroker(raw)
  opts.SetClientID(clientID)
  return opts
}