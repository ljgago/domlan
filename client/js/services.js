'use strict';
/*
angular.module('Domlan')
  
  .factory('rimoPouchDB', function ($rootScope) {
    var urldb = 'http://' + location.hostname + ':' + 5984 + '/devices';
    return {
      localDB: PouchDB('devices'),
      remoteDB: PouchDB(urldb), // no usado por el momento
      replicate: function () {
        var sync1 = PouchDB.replicate('devices', urldb, {live: true});
        var sync2 = PouchDB.replicate(urldb, 'devices', {live: true})
          .on('change', function (info) {
            // handle change
            console.log('up-to-date: la base de datos!');
            $rootScope.$broadcast('@sync:all');
            
          })
          .on('uptodate', function (info) {
            // handle up-to-date
            // se dispara el evento cuando los datos estan actualizados

            
          })
          .on('error', function (err) {
            // handle error
          });
      }
    };
  });
*/

/*
  .service('rimoPouchDB', function ($rootScope) {
    //this.urldb = 'http://' + location.hostname + ':' + 5984 + '/devices';
    
    //ṭhis.localDB = new PouchDB('devices');
    //this.remoteDB = new PouchDB(urldb);

    this.sync1 = PouchDB.replicate('devices', 'http://localhost:5984/devices', {live: true});
    this.sync2 = PouchDB.replicate('http://localhost:5984/devices', 'devices', {live: true})
      .on('uptodate', function (info) {
        // handle up-to-date
        // se dispara el evento cuando los datos estan actualizados
        console.log('up-to-date la base de datos!');
        $rootScope.$broadcast('algo');
      })
      .on('error', function (err) {
        // handle error
      });
  });*/

/*
 $provide.decorator('$browser', ['$delegate', function($delegate) {
        var superUrl = $delegate.url;
        $delegate.url = function(url, replace) {
            if(url !== undefined) {
                return superUrl(url.replace(/\%20/g,"+"), replace);
            } else {
                return superUrl().replace(/\+/g,"%20");
            }
        }
        return $delegate;
    }]);
*/






/*
Rimocon.factory('pouchWrapper', ['$q', '$rootScope', 'myPouch', function($q, $rootScope, myPouch) {
  return {
    add: function(text) {
      var deferred = $q.defer();
      var doc = {
        type: 'todo',
        text: text
      };
      myPouch.post(doc, function(err, res) {
        $rootScope.$apply(function() {
          if (err) {
            deferred.reject(err)
          } else {
            deferred.resolve(res)
          }
        });
      });
      return deferred.promise;
    },
    remove: function(id) {
      var deferred = $q.defer();
      myPouch.get(id, function(err, doc) {
        $rootScope.$apply(function() {
          if (err) {
            deferred.reject(err);
          } else {
            myPouch.remove(doc, function(err, res) {
              $rootScope.$apply(function() {
                if (err) {
                  deferred.reject(err)
                } else {
                  deferred.resolve(res)
                }
              });
            });
          }
        });
      });
      return deferred.promise;
    }
  }
}]);

Rimocon.factory ('listener', ['$rootScope', 'myPouch', function($rootScope, myPouch) {
  myPouch.changes({
    live: true
  })

  .on('change', function (change) {
    $rootScope.$apply(function () {
      myPouch.get(change.id, function (err, doc) {
        $rootScope.$apply(function () {
          if (err) console.log(err);
          $rootScope.$broadcast('change', doc);
        })
      });
    })
  })
  .on('uptodate', function (uptodate) {
    colsole.log("Updated!");
  })
  .on('error', function (err) {
    console.log(err);
  })

  
    else {
      $rootScope.$apply(function() {
        $rootScope.$broadcast('delTodo', change.id);
      });
    }
  }).

      
    }
  })
}]);

*/

