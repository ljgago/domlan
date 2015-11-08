'use strict';

angular.module('Domlan').
  controller('DomlanCtrl', function ($scope, $rootScope, $websocket) {
    //$.material.init();
    $scope.drop = "Hola";

    $scope.protocols = [
      {type: 'mqtt', name: 'MQTT'},
      {type: 'http', name: 'HTTP'}
    ];
    $scope.categories = [
      {type: 'ligth', name: 'Luz'},
      {type: 'sensor', name: 'Sensor'},
      {type: 'trim', name: 'Trim'},
      {type: 'webcam', name: 'Webcam'}
    ];

    $scope.locations = ['Cocina', 'Comedor', 'Baño', 'Cochera', "Webcam"];

    $scope.devices = ['Sensor 1', 'Sensor 2', 'Sensor 3', 'Sensor 4'];
    
    // db va a contener todos los datos recividos del servidor
    $rootScope.db = "";

    // Websockets
    var ws = $websocket.$new({
      url: 'ws://' + window.location.host + '/ws',
      reconnect: true
      //protocols: ['binary', 'base64']
      //reconnectInterval: 500 // it will reconnect after 5 seconds
    });
    ws.$on('$open', function () {
      console.log('Here we are and I\'m pretty sure to get back here for another time at least!');
      ws.$emit('msg-device', {hola:'Hola yo soy un device', pepe: 123, repetido: 'on'});
    })
    .$on('$close', function () {
      console.log('Got close, damn you silly wifi!');
    })
    .$on('start', function (msg) {
      console.log("HolaaaA");
      $rootScope.db = msg;
      $scope.$apply();
      for(var i=0 ; i<msg.length ; i++) {
        console.log(msg[i])
      }
    })
    .$on('update-device', function (msg) {
      $rootScope.db = msg;
      $scope.$apply();
      for(var i=0 ; i<msg.length ; i++) {
        console.log(msg[i])
      }
    });

    /* 
     *  Lista de funciones
     */

    // addDevice: envío los datos del nuevo dispositivo al servidor
    $scope.addDevice = function (dev) {
      dev.device = { "id": dev.id };
      //console.log(dev);
      ws.$emit('add-device', dev);
      //console.log(device);
    };
    
    // Funciones para los diferentes controles
    
    // Control Switch 
    $scope.switchControler = function (dev) {
      ws.$emit('update-device', dev);
      //console.log(dev);
    };

    $scope.removeDevice = function (dev) {
      ws.$emit('remove-device', dev);
    };

    // Control slide
    



    $scope.imparFolders = ['1', '2', '3', '4', '5', '6'];
    $scope.parFolders = ['7', '8', '9', '10'];


    $scope.imparFolderList = function () {
      return imparFolderList($scope.imparFolders);
    };

    $scope.parFolderList = function () {
      return parFolderList($scope.parFolders);
    };

    $scope.sortableConfig = { group: 'folder', animation: 150 };
    'Start End Add Update Remove Sort'.split(' ').forEach(function (name) {
      $scope.sortableConfig['on' + name] = console.log.bind(console, name);
    })
  }).
  controller('gridCtrl', function ($scope, $websocket) {

    $scope.myData = [
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' },
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' },
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' },
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' },
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' },
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' },
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' },
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' },
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' },
      { no: '1', name: 'Hola', location: 'Casa', active: 'OFF' }
    ];
    $scope.gridOptions = {
      data: 'myData',
      enableCellSelection: true,
      enableRowSelection: false,
      enableCellEdit: true,
      columnDefs: [
        {field: 'no', displayName: 'Nº', enableCellEdit: true}, 
        {field: 'name', displayName: 'Nombre', enableCellEdit: true},
        {field: 'location', displayName: 'Ubicación', enableCellEdit: true}, 
        {field: 'active', displayName: 'Activo', enableCellEdit: true}
      ]
    }
  }).
  controller('webcamCtrl', function ($scope, $rootScope, $websocket) {
    var ws = $websocket.$get('ws://' + window.location.host + '/ws');
    var ctx = document.getElementById('canvas').getContext('2d');
    var img = new Image();
    //var img = document.getElementById('miImg');
    //console.log('Out');
    console.log('22');
    ws.$on('webcam', function (msg) {
      console.log('In');
      console.log(msg);
      img.onload = function () {
        ctx.drawImage(img, 0, 0);
      };
      img.src = 'data:image/jpeg;base64,' + msg;
      //$scope.apply();
      //console.log('PPPPPPPPPPPPPPPPPPPPPPPPPPPPPP');
    });
  }).
  controller('homeCtrl', function ($scope) {
    $scope.templateU = '<hr><div class="row"><div class="medium-12 columns"><a class="button radius large expand">Agregar dispositivo</a><a class="button radius large expand">Eliminar dispositivo</a><a class="button radius large expand">Configuración</a></div></div>';
  }).
  controller('locationCtrl', function ($scope, $routeParams) {
    console.log($routeParams);
    $scope.templateUrl = 'partials/' + $routeParams.a;

    $scope.locate = $routeParams.a;
    console.log($scope.locate);
  });


  
/*
    var ws = $websocket('ws://' + window.location.host + '/ws');

    ws.onOpen(function() {
      console.log('connection open');
      ws.send('Hello World');
    });
    ws.onError( function (event) {
      console.log('connection Error', event);
    });
    ws.onClose( function (event) {
      console.log('connection closed', event);
    });
    ws.onMessage( function (msg) {
      var data = JSON.parse(msg.data);
      console.log(data.data[0]);
    });
    */
    
    
// setTimeout(function() {
// ws.close();
// }, 500)


    /*
    console.log("ws://" + window.location.host + "/ws");
    $scope.conn = new WebSocket("ws://" + window.location.host + "/ws");
    $scope.conn.onopen = function() {
      $scope.conn.send("Hola");
      // Write message on receive
    }
    $scope.conn.onmessage = function (msg) {
      console.log(msg.data);
      $scope.drop = msg.data;
      $scope.$apply();
    }
    $scope.conn.onclose = function () {
      console.log("Desconectado del websocket");
    }*/
    


    


    //console.log ("Hola!!!");
    /*var socket = io();
    socket.on("device", function (msg) {
      console.log (msg);
      $scope.drop = msg;
      $scope.$apply();
    });
    socket.emit("device", "Soy en cliente enviado mensaje XD");
    */
    /*
    $scope.startWebsocket = function startWebsocket(websocket_host) {
      console.log(websocket_host);
      $scope.conn = new WebSocket(websocket_host);
      $scope.conn.onopen = function() {
        $scope.conn.send("Hola");
        // Write message on receive
      }
      $scope.conn.onmessage = function (msg) {
        console.log(msg.data);
        $scope.drop = msg.data;
        $scope.$apply();
      }
      $scope.conn.onclose = function () {
        setTimeout(function () { startWebsocket(websocket_host) }, 5000);
      }
    }
    $scope.startWebsocket("ws://" + window.location.host + "/ws");

    $scope.conn.send("Pepe como andas?");
    */


    /*
      var host = "ws://" + window.location.host + "/ws";
      console.log(host);
      var conn = new WebSocket(host);
      conn.onopen = function() {
        conn.send("Connection init");
        // Write message on receive
        
      }
      conn.onmessage = function (msg) {
        console.log(msg.data);
        $scope.drop = msg.data;
        $scope.$apply();
      }
      conn.onclose = function () {
        setTimeout(function(){start(websocketServerLocation)}, 5000);
      }
      */
