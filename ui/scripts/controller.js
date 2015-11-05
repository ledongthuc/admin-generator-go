var menuControllers = angular.module('menuControllers', []);
menuControllers.controller('MenuListCtrl', function ($scope, $rootScope, MenusFactory) {
    $rootScope.menus = MenusFactory.query()
});

var contentControllers = angular.module('contentControllers', []);
contentControllers.controller('ContentListCtrl', function ($scope, $routeParams, $rootScope, ColumnsFactory, ContentsFactory) {
    var pageName = $routeParams.name
    $rootScope.pageName = pageName
    $rootScope.action = "list"

    $scope.columns = ColumnsFactory.query({table_name: pageName}, function(columns) {
        angular.forEach(columns, function(value, key) {
            if(value.primary) {
                $scope.primaryName = value.name
                return
            }
        });
    })

    $scope.cells = ContentsFactory.query({table_name: pageName})
    $scope.delete = function(id) {
      confirm("Do you want to delete?")
    };
});

contentControllers.controller('ContentAddCtrl', function ($scope, $routeParams, $rootScope, ColumnsFactory, ContentsFactory) {
    var pageName = $routeParams.name
    $rootScope.pageName = pageName
    $rootScope.action = "add"
    $scope.contentItem = {table_name: pageName}
    $scope.addContent = function() {
        ContentsFactory.create($scope.contentItem,
        function(data) {
            $scope.alertSuccessStyle={'display':'none'}
            $scope.alertErrorStyle={'display':'none'}
            $scope.alertSuccessStyle={'display':'block'}
            $scope.contentItem = {table_name: pageName}
        },
        function(data, status) {
            $scope.alertSuccessStyle={'display':'none'}
            $scope.alertErrorStyle={'display':'none'}
            $scope.alertErrorStyle={'display':'block'}
            $scope.contentItem = {table_name: pageName}
            $scope.errorMessage = data.data
        })
    }

    $scope.columns = ColumnsFactory.query({table_name: pageName})
});

contentControllers.controller('ContentEditCtrl', function ($scope, $routeParams, $rootScope, ColumnsFactory, ContentFactory) {
    var pageName = $routeParams.name
    var id = $routeParams.id
    $rootScope.pageName = pageName
    $rootScope.action = "edit"
    $scope.columns = ColumnsFactory.query({table_name: pageName})
    $scope.contentItem = ContentFactory.show({table_name: pageName, id: id})
});
