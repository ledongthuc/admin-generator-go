var generatedAdminApp = angular.module('generatedAdmin', [
    "ngRoute",
    "menuControllers",
    "contentControllers",
]);

generatedAdminApp.filter('capitalize', function() {
    return function(input) {
      return (!!input) ? input.charAt(0).toUpperCase() + input.substr(1).toLowerCase() : '';
    }
});

generatedAdminApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/list/:name', {
        templateUrl: 'partials/list.html',
        controller: 'ContentListCtrl'
      }).
      when('/add/:name', {
        templateUrl: 'partials/add.html',
        controller: 'ContentAddCtrl'
      }).
      when('/edit/:name/:id', {
        templateUrl: 'partials/edit.html',
        controller: 'ContentEditCtrl'
      }).
      otherwise({
        redirectTo: '/home'
      });
  }]);

var menuControllers = angular.module('menuControllers', []);
menuControllers.controller('MenuListCtrl', function ($scope, $http, $rootScope) {
    $http.get('/api/menu').success(function(data) {
        $rootScope.menus = data;
    });

});

var contentControllers = angular.module('contentControllers', []);
contentControllers.controller('ContentListCtrl', function ($scope, $http, $routeParams, $rootScope) {
    var pageName = $routeParams.name
    $rootScope.pageName = pageName
    $rootScope.action = "list"
    $http.get('/api/column/'+pageName).success(function(data) {
        $scope.columns = data;
        angular.forEach(data, function(value, key) {
            if(value.primary) {
                $scope.primaryName = value.name
            }
        })
    });

    $http.get('/api/'+pageName).success(function(data) {
        $scope.cells = data;
        $scope.delete = function(id) {
          confirm("Do you want to delete?")
        };
    });
});
contentControllers.controller('ContentAddCtrl', function ($scope, $http, $routeParams, $rootScope) {
    var pageName = $routeParams.name
    $rootScope.pageName = pageName
    $rootScope.action = "add"
    $http.get('/api/column/'+pageName).success(function(data) {
        $scope.columns = data;
    });
});
contentControllers.controller('ContentEditCtrl', function ($scope, $http, $routeParams, $rootScope) {
    var pageName = $routeParams.name
    var id = $routeParams.id
    $rootScope.pageName = pageName
    $rootScope.action = "edit"
    $http.get('/api/column/'+pageName).success(function(data) {
        $scope.columns = data;
    });

    $http.get('/api/'+pageName+"/"+id).success(function(data) {
        $scope.data = data;
    });
});
