'use strict';

angular.module('Domlan').
  directive('dirType', function () {
    return {
      restrict: 'EA', 
      templateUrl: '/assets/views/controls/main.html'
    };
  });


/*
'<div ng-repeat="dbl in db | filterBy: [\'location\']:locate | orderBy:\'name\'">' +
        '<div class="panel">'+
          '<div class="text-center">'+
            '<div class="row">'+
              '<div class="small-6 columns">'+
                '<h2>{{dbl.name}}</h2>'+
              '</div>'+
              '<div class="small-6 columns">'+
                '<div style="margin-bottom: 0px;" class="switch round large">'+
                  '<input id="{{dbl.id}}" type="checkbox" ng-model="data" ng-click=""/>'+
                  '<label for="{{dbl.id}}"></label>'+
                '</div>'+
              '</div>'+
            '</div>'+
          '</div>'+
        '</div>'+
      '</div>',
*/
/*

directive('dirDevice', function () {
    return{
      restrict: 'EA',
      template: '<div><div>Hola {{ nombre }}</div>' +
      'Escribe tu edad: <input ng-model="edad" />' +
      '<a href="#" ng-click="show()">{{ link }}</div>',
      scope: {
      nombre: "@", //variables de alcance($scope) o por valor
      edad: "=", //usado para hacer uso de data-binding(datos entre vista controlador) o por referencia
      link: "@",
      show: "&" //útiles para llamar a funciones
      },
    }
  }).
  directive('dirLocation', function () {
    return{
      restrict: 'EA',
      template: '<dir-device>',
      scope: {
      nombre: "@", //variables de alcance($scope) o por valor
      edad: "=", //usado para hacer uso de data-binding(datos entre vista controlador) o por referencia
      link: "@",
      show: "&" //útiles para llamar a funciones
      },
    }
  }).
  directive('devLight', function () {
    return {
      restrict: 'EA',
      template: '<div class="panel">' +
        '<div class="text-center">' +
          '<div class="row">' +
            '<div class="small-6 columns">' +
              '<h2>{{name}}</h2>'
            '</div>' +
            '<div class="small-6 columns">' +
              '<div style="margin-bottom: 0px;" class="switch round large">' +
                '<input id="switchName" type="checkbox"/>' +
                '<label for="switchName"></label>' +
              '</div>' +
            '</div>' +
          '</div>' +
        '</div>' +
      '</div>',

    };
  }).




Domlan.directive('etiquetas', function () {
  return {
    template : '<p>{{text}}</p>',
    link : function (scope, elem, attrs) {
      scope.text=attrs.etiquetas;
    }
  };
});

Domlan.directive('onoff', function() {
  return {
    template:
      '<div ng-controller="{{onoff}}">
        <div class="btn-group" ng-switch="selected">
          <button class="btn btn-success" ng-switch-when="true" ng-click="button1()"> <b>ON</b> </button>
          <button class="btn btn-danger" ng-switch-when="false" ng-click="button2()"> <b>OFF</b> </button>
        </div>
      </div>',
    restrict: 'A',
    link: function (scope, elem, attrs) {
      //$scope.idcontroller = $scope.onoff;
    }
  };
});

Domlan.directive('contentItem', function ($compile) {
  var imageTemplate = '<div class="entry-photo"><h2>&nbsp;</h2><div class="entry-img"><span><a href="{{rootDirectory}}{{content.data}}"><img ng-src="{{rootDirectory}}{{content.data}}" alt="entry photo"></a></span></div><div class="entry-text"><div class="entry-title">{{content.title}}</div><div class="entry-copy">{{content.description}}</div></div></div>';
  var videoTemplate = '<div class="entry-video"><h2>&nbsp;</h2><div class="entry-vid"><iframe ng-src="{{content.data}}" width="280" height="200" frameborder="0" webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe></div><div class="entry-text"><div class="entry-title">{{content.title}}</div><div class="entry-copy">{{content.description}}</div></div></div>';
  var noteTemplate = '<div class="entry-note"><h2>&nbsp;</h2><div class="entry-text"><div class="entry-title">{{content.title}}</div><div class="entry-copy">{{content.data}}</div></div></div>';

  var getTemplate = function(contentType) {
    var template = '';

    switch(contentType) {
      case 'image':
        template = imageTemplate;
        break;
      case 'video':
        template = videoTemplate;
        break;
      case 'notes':
        template = noteTemplate;
        break;
    }

    return template;
  }

  var linker = function(scope, element, attrs) {
    scope.rootDirectory = 'images/';

    element.html(getTemplate(scope.content.content_type)).show();

    $compile(element.contents())(scope);
  }

  return {
    restrict: "E",
    link: linker,
    scope: {
        content:'='
    }
  };
});

*/
/*
app.directive ('diHref', ['$location', '$route',
  function ($location, $route) {
    return function (scope, element, attrs) {
      scope.$watch ('diHref', function() {
        if(attrs.diHref) {
          element.attr('href', attrs.diHref);
          element.bind('click', function (event) {
            scope.$apply (function () {
              console.log($location.path());
              $route.reload();
            });
          });
        }
      });
    }
}]);

*/