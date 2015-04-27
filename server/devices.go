package server

import (
  //"../modules/mqtt"
  "../modules/mqttgmp"
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
  d.mqttDev.On("device/id/#", func(data []byte) {
    //key, value := ldb.MergeDeviceDB(data)
    //ldb.PutDeviceDB([]byte(key), []byte(value))
    //H.EmitBroadcast(`"update-devices"`, value)
    DEBUG("Entro al handler mqtt:", string(data))
  })

  /****************************************************************
  *  WEBCAMS
  ****************************************************************/



  
}