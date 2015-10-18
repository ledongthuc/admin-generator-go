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
      otherwise({
        redirectTo: '/home'
      });
  }]);

var menuControllers = angular.module('menuControllers', []);
menuControllers.controller('MenuListCtrl', function ($scope, $http) {
    $http.get('/api/menu').success(function(data) {
        $scope.menus = data;
    });

});

var contentControllers = angular.module('contentControllers', []);
contentControllers.controller('ContentListCtrl', function ($scope, $http, $routeParams, $rootScope) {
    var pageName = $routeParams.name
    $rootScope.pageName = pageName
    $rootScope.action = "list"
    $http.get('/api/column/'+pageName).success(function(data) {
        $scope.columns = data;
    });

    $http.get('/api/'+pageName).success(function(data) {
        $scope.cells = data;
    });
});
