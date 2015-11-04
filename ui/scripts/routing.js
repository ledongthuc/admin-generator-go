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
