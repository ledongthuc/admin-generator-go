var appServices = angular.module('appServices', []);

appServices.factory('MenusFactory', function ($resource) {
    return $resource('/api/menu', {}, {
        queryList: { method: 'GET', isArray: true},
    })
});

appServices.factory('ColumnsFactory', function ($resource) {
    return $resource('/api/column', {}, {
        queryList: { method: 'GET', isArray: true},
    })
});

appServices.factory('ContentsFactory', function ($resource) {
    return $resource('/api/:table_name', {}, {
        queryList: { method: 'GET', isArray: true},
        create: { method: 'POST', params: {table_name: '@table_name'} }
    })
});

appServices.factory('ContentFactory', function ($resource) {
    return $resource('/api/:table_name/:id', {}, {
        queryOne: { method: 'GET'},
        update: { method: 'PUT', params: { table_name: '@table_name', id: '@id' }},
        delete: { method: 'DELETE', params: { table_name: '@table_name', id: '@id' }},
    })
});
