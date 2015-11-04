var appServices = angular.module('appServices', []);

appServices.factory('MenusFactory', function ($resource) {
    return $resource('/api/menu', {}, {
        query: { method: 'GET', isArray: true},
    })
});

appServices.factory('ColumnsFactory', function ($resource) {
    return $resource('/api/column', {}, {
        query: { method: 'GET', isArray: true},
    })
});

appServices.factory('ContentsFactory', function ($resource) {
    return $resource('/api/:table_name', {}, {
        query: { method: 'GET', isArray: true},
        create: { method: 'POST', params: {table_name: '@table_name'} }
    })
});

appServices.factory('ContentFactory', function ($resource) {
    return $resource('/api/:table_name/:id', {}, {
        show: { method: 'GET'},
    })
});
