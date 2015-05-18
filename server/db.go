package server

import (
  "sync"
  "log"
  "github.com/bitly/go-simplejson"
  "github.com/syndtr/goleveldb/leveldb"
)

/****************************************************************
*  LEVELDB
****************************************************************/

var mutex = &sync.Mutex{}

var ldb DeviceDB

// DeviceDB: main structure of database
//
// DeviceDB: estrutura principal de la base de datos
type DeviceDB struct {
  db *leveldb.DB
}

// OpenDB: open the database
//
// OpenDB: abro la base de datos
func (l *DeviceDB) OpenDB() {
  var err error
  l.db, err = leveldb.OpenFile("./db", nil)
  if err != nil {
    if err.Error() == "ErrCorrupted" {
      log.Println(err)
      l.db, err = leveldb.RecoverFile("./db", nil)
    }
    if err != nil {
      log.Println("Error al abrir la base de datos:", err)
      return
    }
  }
}

// GetAllDeviceDB: get the list of all devices
//
// GetAllDeviceDB: obtengo la lista de todos los dispositivos
func (l *DeviceDB) GetAllDeviceDB() string {
  
  mutex.Lock()
  l.OpenDB()
  defer func() {
    l.db.Close()
    mutex.Unlock()
  }()
  
  iter := l.db.NewIterator(nil, nil)
  data := ""
  for iter.Next() {
    data += string(iter.Value()) + ","
  }

  if data == "" {
    data += "[]"  // Si la base de datos esta vacía
  } else {
    data = "[" + data[:(len(data)-1)] + "]" // Agrego "[]" borrando la ultima coma ","
  }

  iter.Release()
  err := iter.Error()
  if err != nil {
    log.Println(err)
    return ""
  }
  return data
}

// GetDeviceDB: get data of the specified device
//
// GetDeviceDB: obtengo los datos del dispositivo especificado
func (l *DeviceDB) GetDeviceDB(key []byte) string {

  mutex.Lock()
  l.OpenDB()
  defer func() {
    l.db.Close()
    mutex.Unlock()
  }()

  data, err := l.db.Get(key, nil)
  if err != nil {
    log.Println(err)
    return ""
  }
  return string(data)
}

// PutDeviceDB: agrego a la base de datos
//
// PutDeviceDB: add to database
func (l *DeviceDB) PutDeviceDB(key []byte, rawJSON []byte) {

  mutex.Lock()
  l.OpenDB()
  defer func() {
    l.db.Close()
    mutex.Unlock()
  }()

  js, err := simplejson.NewJson(rawJSON)
  if err != nil {
    log.Println(err)
  }
  // Para Producción
  //value, err := js.Encode()
  // Para debug
  value, err := js.EncodePretty()
  
  if err != nil {
    log.Println(err)
  }
  err = l.db.Put(key, value, nil)
  if err != nil {
    log.Println(err)
  }
  DEBUG(string(key), string(value))
}

// DeleteDeviceDB: delete item of the databse.
//
// DeleteDeviceDB: elimino un item de la base de datos.
func (l *DeviceDB) DeleteDeviceDB(key []byte) {

  mutex.Lock()
  l.OpenDB()
  defer func() {
    l.db.Close()
    mutex.Unlock()
  }()

  err := l.db.Delete(key, nil)
  if err != nil {
    log.Println(err)
  }
}

// MergeDeviceDB: merge data of database and the devices
//
// MergeDeviceDB: Mezclo los datos de la base de datos y el de los dispositivos
func (l *DeviceDB) MergeDeviceDB(device []byte) (string, string) {
  
  dv, err := simplejson.NewJson(device)
  if err != nil {
    log.Println("JSON:", err)
    return "", ""
  }
  
  key := dv.Get("device").Get("id").MustString()
  old := l.GetDeviceDB([]byte(key))
  if old == "" {
    log.Println("LevelDB: No hay ningun valor asignado a este id:", key)
    return "", ""
  }

  js, err := simplejson.NewJson([]byte(old))
  if err != nil {
    log.Println("JSON:", err)
    return "", ""
  }

  js.Get("data").Set("device", dv.Get("device").Interface())

  value, _ := js.Get("data").Encode()
  if err != nil {
    log.Println(err)
  }
  // Guardo los datos meclados en la base de datos
  
  DEBUG("JSON:", string(value))
  return key, string(value)

}

