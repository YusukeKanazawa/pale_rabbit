var module = angular.module('pale_rabbit', ['ngRoute', 'ngResource']);

// var itemTypes = [
//   {"_id":"3d5fb08dbdeae2c2a2edf35d88000a84","_rev":"3-dba887d14fd326155f90a015643f838b","doctype":"itemtype","typeCode":"TW","name":"タオル","typeAttr":[{"rank":1,"attrCode":"length","name":"長さ","field":{"type":"number","length":"4"},"default":"","max":"","min":"","unit":"mm","specExpantsion":{"codeLength":"1"}},{"rank":2,"attrCode":"width","name":"幅","field":{"type":"number","length":"4"},"default":"","max":"","min":"","unit":"mm"},{"rank":3,"attrCode":"weight","name":"重さ","field":{"type":"number","length":"4"},"default":"","max":"","min":"","unit":"g"},{"rank":4,"attrCode":"color","name":"色","field":{},"default":"","max":"","min":"","unit":"-","specExpantsion":{"codeLength":"1"}}]},
//   {"_id":"3d5fb08dbdeae2c2a2edf35d8800107c","_rev":"1-0ccc032e90c84b6728e52fd20e33b2df","doctype":"itemtype","typeCode":"GEM","name":"石","typeAttr":[{"rank":1,"attrCode":"phi","name":"径","field":{"type":"number","length":"2"},"default":"","max":"","min":"","unit":"mm","specExpantsion":{"codeLength":"2"}},{"rank":2,"attrCode":"width","name":"色","field":{"type":"number","length":"4"},"default":"","max":"","min":"","unit":"","specExpantsion":{"codeLength":"1"}}]}
// ];
  // var itemTypeResource = $resource('/api/itemType', {}, {getAll: {method: 'GET', isArray: true}});

  // var items = [
  //   {
  //     rank: 1,
  //     code: 'TW001',
  //     name: 'face towel',
  //     type: 'TW',
  //     saleHandle: [
  //       {
  //         rank: 1,
  //         no: 0,
  //         name: '10set',
  //         lotsize: '10',
  //         unit: '枚'
  //       },
  //       {
  //         rank: 2,
  //         no: 1,
  //         name: '20set',
  //         lotsize: '20',
  //         unit: '枚'
  //       }
  //     ],
  //     typeAttrs: [
  //     {
  //       specCode:'3W',
  //       length: '300',
  //       width: '150',
  //       weight: '20',
  //       color: 'white'
  //     },
  //     {
  //       specCode:'3B',
  //       length: '300',
  //       width: '150',
  //       weight: '20',
  //       color: 'blue'
  //     },
  //     {
  //       specCode:'6W',
  //       length: '600',
  //       width: '150',
  //       weight: '20',
  //       color: 'blue'
  //     },
  //     {
  //       specCode:'3K',
  //       length: '300',
  //       width: '150',
  //       weight: '20',
  //       color: 'black'
  //     }
  //     ]
  //   },    
  //   {
  //     rank: 1,
  //     code: 'TW002',
  //     name: 'bath towel',
  //     type: 'TW',
  //     saleHandle: [
  //       {
  //         rank: 1,
  //         no: 0,
  //         name: '5set',
  //         lotsize: '5',
  //         unit: '枚'
  //       },
  //       {
  //         rank: 2,
  //         no: 1,
  //         name: '10set',
  //         lotsize: '10',
  //         unit: '枚'
  //       }
  //     ],
  //     typeAttrs: [
  //     {
  //       length: '300',
  //       width: '150',
  //       weight: '20',
  //       color: 'white'
  //     },
  //     {
  //       length: '300',
  //       width: '150',
  //       weight: '20',
  //       color: 'blue'
  //     }
  //   ]}
  // ];


module.config(function($routeProvider) {
  $routeProvider
  .when('/item', {
    templateUrl: 'view/item.html',
    controller : 'itemController'
  })
  .when('/item/spec', {
    templateUrl: 'view/item_spec.html',
    controller : 'itemSpecController'
  })
  .when('/item/type', {
    templateUrl: 'view/item_type.html',
    controller : 'itemTypesController'
  });

});

// item type
module.controller('itemTypesController', function($scope, $http, $resource) {
  // resource
  var itemTypeResource = $resource('/api/itemType', {}, {getAll: {method: 'GET', isArray: true}});
  // selector
  $scope.selector = {
    select: function(id) {
        $scope.selector.selectedId = id;
      },
    selectedId: 0,
  };
  // 追加
  $scope.append = function() {
    var selected_item_type = $scope.itemTypes[$scope.selector.selectedId];
    console.log($scope.attrCode);
    if(!$scope.attrCode){
      console.log('code is blank');
      return
    }
    if(!$scope.attrName){
      console.log('name is blank');
      return
    }
    if (
      selected_item_type.attrs.find(function(d){
        return d.attr_code ==  $scope.attrCode;
      })
    ){
      console.log('duplicate');
      return
    }

    // append new attr
    var new_attr = {
      "attr_code": $scope.attrCode,
      "attr_name": $scope.attrName,
      "rank": selected_item_type.attrs.length + 1,
      "updated_by": "system",
      "updated_at": "",
      "isInserted": true
    };

    selected_item_type.attrs.push(new_attr)

    // reset form models.;
    $scope.attrCode = '';
    $scope.attrName = '';


  };
  $scope.reset = function(){
    $scope.attrCode = '';
    $scope.attrName = '';
  }
  $scope.commit = function(){

  }
  $scope.itemTypes = itemTypeResource.getAll();
});




module.controller('itemSpecController', function($scope, $http) {
 $scope.attrSelector = {
    select: function(id){
        $scope.attrSelector.selectedId = id;
        $scope.specs=$scope.itemType.typeAttr[id].attrCode;
      },
    selectedId: 0,
  };

  $scope.itemType = itemTypes.find(function(data){
    return data.typeCode == this;
  }, 'TW');

  $scope.items = items;

});

module.controller('itemController', function($scope, $http) {
 $scope.selector = {
    select: function(id){
        $scope.selector.selectedId = id;
        $scope.itemType = itemTypes.find(function(data){
          return data.typeCode == this;
        }, items[id].type);
      },
    selectedId: 0,
  };

  $scope.itemTypes = itemTypes;
  $scope.items = items;
  $scope.detailView = {

  };
});