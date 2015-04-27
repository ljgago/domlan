#!/usr/bin/env bash
echo -e "\nDownloading Go dependencies..."
echo "go get -u github.com/gin-gonic/gin"
go get -u github.com/gin-gonic/gin
echo "go get -u github.com/gorilla/websocket"
go get -u github.com/gorilla/websocket
echo "go get -u github.com/bitly/go-simplejson" 
go get -u github.com/bitly/go-simplejson
echo "go get -u github.com/syndtr/goleveldb/leveldb"
go get -u github.com/syndtr/goleveldb/leveldb
echo "go get -u git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
go get -u git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git
echo "go get -u code.google.com/p/go-uuid/uuid"
go get -u code.google.com/p/go-uuid/uuid
echo "go get -u github.com/franela/goreq"
go get -u github.com/franela/goreq
echo "go get -u github.com/yosssi/gmq/mqtt"
go get -u github.com/yosssi/gmq/mqtt
echo -e "\nDownloading bower dependencies..."
bower install
echo -e "\nCompilado Domlan:"
go build ./domlan.go
echo "Fin!"
