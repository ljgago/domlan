package server

import (
  //"../modules/mqtt"
  "../modules/mqttgmp"
  //"github.com/ljgago/glue"
  //"time"
)

var Dev Devices

type Devices struct {
  // Add the the device modules
  //mqttDev *mqttpaho.MqttModule
  mqttDev *mqttgmp.MqttModule
}

func (d *Devices) DeviceWorks() {
  
  /****************************************************************
  *
  *  Configure the differents device modules
  *  Configuro los diferentes módulos de dispositivos
  *
  ****************************************************************/
  /****************************************************************
  *  MQTT
  ****************************************************************/
  //d.mqttDev = mqttpaho.NewMqttModule("server", "tcp://0.0.0.0:1883", "DomlanMQTT")
  d.mqttDev = mqttgmp.NewMqttModule("server", "0.0.0.0:1883", "DomlanMQTT")
  d.mqttDev.Connect()
  d.mqttDev.On("server/#", func(device []byte) {
    key, value := ldb.MergeDeviceDB(device)
    ldb.PutDeviceDB([]byte(key), []byte(value))
    data := ldb.GetAllDeviceDB()
    Ruta.EmitBroadcast(`"update-devices"`, string(data))
    DEBUG("Entro al handler mqtt:", string(device))
  })



  /****************************************************************
  *  WEBCAMS
  ****************************************************************/
  
}

/****************************************************************
*  WEBSOCKETS
****************************************************************/
