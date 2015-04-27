'use strict';

// Declare app level module which depends on filters, and services
// 'geoJobs.filters', 'geoJobs.services', 'geoJobs.directives'
//var app = angular;

var Domlan = angular.module ('Domlan', ['ngRoute','ngWebsocket']).
  config (['$locationProvider', '$routeProvider', 
    function ($locationProvider, $routeProvider) {
    
    $routeProvider.
      when('/', {
        template: '<hr><div class="row"><div class="medium-12 columns"><a href="#/add" class="button radius large expand">Agregar dispositivo</a><a href="#/remove" class="button radius large expand">Eliminar dispositivo</a><a class="button radius large expand">Configuración</a></div></div>',
      }).
      when('/add', {
        templateUrl: '/assets/views/add.html'
      }).
      when('/remove', {
        templateUrl: '/assets/views/remove.html'
      }).
      when("/Webcam", {
        templateUrl: '/assets/views/webcam.html',
        controller: 'webcamCtrl'
      }).
      when('/:a', {
        template: '<div>Loading...</div>',
        controller: 'locationCtrl'
      })
}]);

/*controller('DynamicController', function ($scope, $routeParams) {
console.log($routeParams);
$scope.templateUrl = 'partials/' + $routeParams.a;
}).
*/
/*
function deviceCompare(a, b) {
  if (typeof a.Name !== 'undefined' && typeof b.Name !== 'undefined') {
    if (a.Name < b.Name)
      return -1;
    return a.Name > b.Name;
  }
  if (a.DeviceID < b.DeviceID) {
    return -1;
  }
  return a.DeviceID > b.DeviceID;
}

function folderCompare(a, b) {
  if (a.ID < b.ID) {
    return -1;
  }
  return a.ID > b.ID;
}

function folderMap(l) {
  var m = {};
  l.forEach(function (r) {
    m[r.ID] = r;
  });
  return m;
}

function folderList(m) {
  var l = [];
  for (var id in m) {
    l.push(m[id]);
  }
  l.sort(folderCompare);
  return l;
}

function decimals(val, num) {
  var digits, decs;

  if (val === 0) {
    return 0;
  }

  digits = Math.floor(Math.log(Math.abs(val)) / Math.log(10));
  decs = Math.max(0, num - digits);
  return decs;
}

function randomString(len) {
  var i, result = '', chars = '01234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-';
  for (i = 0; i < len; i++) {
    result += chars[Math.round(Math.random() * (chars.length - 1))];
  }
  return result;
}

function isEmptyObject(obj) {
  var name;
  for (name in obj) {
    return false;
  }
  return true;
}

function debounce(func, wait) {
  var timeout, args, context, timestamp, result, again;

  var later = function () {
    var last = Date.now() - timestamp;
    if (last < wait) {
        timeout = setTimeout(later, wait - last);
    } else {
      timeout = null;
      if (again) {
        again = false;
        result = func.apply(context, args);
        context = args = null;
      }
    }
  };

  return function () {
    context = this;
    args = arguments;
    timestamp = Date.now();
    var callNow = !timeout;
    if (!timeout) {
      timeout = setTimeout(later, wait);
      result = func.apply(context, args);
      context = args = null;
    } else {
        again = true;
    }
    return result;
  };
}

*/

