var menuControllers = angular.module('menuControllers', []);
menuControllers.controller('MenuListCtrl', function ($scope, $rootScope, MenusFactory) {
    $rootScope.menus = MenusFactory.queryList()
});

var contentControllers = angular.module('contentControllers', []);
contentControllers.controller('ContentListCtrl', function ($scope, $routeParams, $rootScope, ColumnsFactory, ContentsFactory, ContentFactory) {
    var pageName = $routeParams.name
    $rootScope.pageName = pageName
    $rootScope.action = "list"

    $scope.columns = ColumnsFactory.queryList({table_name: pageName}, function(columns) {
        angular.forEach(columns, function(value, key) {
            if(value.primary) {
                $scope.primaryName = value.name
                return
            }
        });
    })

    $scope.cells = ContentsFactory.queryList({table_name: pageName}, function(cells) {
        if(cells == undefined || cells.length == 0) {
            $scope.emptyErrorStyle = {'display':'block'}
            return
        }
    })
    $scope.delete = function(id) {
      var isDelete = confirm("Do you want to delete?")
      if(isDelete) {
          console.log(id)
          ContentFactory.delete({ table_name: pageName, id: id }, function() {
              $scope.cells = ContentsFactory.queryList({table_name: pageName}, function(cells) {
                  if(cells == undefined || cells.length == 0) {
                      $scope.emptyErrorStyle = {'display':'block'}
                      return
                  }
              })
          })

      }
    }
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
            $scope.contentItem = {table_name: pageName} // Reset
        },
        function(data, status) {
            $scope.alertSuccessStyle={'display':'none'}
            $scope.alertErrorStyle={'display':'none'}
            $scope.alertErrorStyle={'display':'block'}
            $scope.contentItem = {table_name: pageName} // Reset
            $scope.errorMessage = data.data
        })
    }

    $scope.columns = ColumnsFactory.queryList({table_name: pageName})
});

contentControllers.controller('ContentEditCtrl', function ($scope, $routeParams, $rootScope, ColumnsFactory, ContentFactory) {
    var pageName = $routeParams.name
    var id = $routeParams.id
    $rootScope.pageName = pageName
    $rootScope.action = "edit"
    $scope.columns = ColumnsFactory.queryList({table_name: pageName})
    $scope.contentItem = ContentFactory.queryOne({table_name: pageName, id: id})

    $scope.updateContent = function() {
        $scope.contentItem.table_name = pageName
        $scope.contentItem.id = id
        ContentFactory.update($scope.contentItem,
        function(data) {
            $scope.alertSuccessStyle={'display':'none'}
            $scope.alertErrorStyle={'display':'none'}
            $scope.alertSuccessStyle={'display':'block'}
        },
        function(data, status) {
            $scope.alertSuccessStyle={'display':'none'}
            $scope.alertErrorStyle={'display':'none'}
            $scope.alertErrorStyle={'display':'block'}
            $scope.errorMessage = data.data
        })
    }
});
