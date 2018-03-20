
'use strict';

var app = angular.module('application', []);

app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.queryAllSupp = function(){

		appFactory.queryAllSupp(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_supp = array;
		});
	}

	// $scope.querySupp = function(){

	// 	var id = $scope.supp_id;

	// 	appFactory.querySupp(id, function(data){
	// 		$scope.query_supp = data;

	// 		if ($scope.query_supp == "Could not locate"){
	// 			console.log()
	// 			$("#error_query").show();
	// 		} else{
	// 			$("#error_query").hide();
	// 		}
	// 	});
	// }

	$scope.recordSupp = function(){

		appFactory.recordSupp($scope.supp, function(data){
			$scope.create_supp = data;
			$("#success_create").show();
		});
	}

	// $scope.changeHolder = function(){

	// 	appFactory.changeHolder($scope.holder, function(data){
	// 		$scope.change_holder = data;
	// 		if ($scope.change_holder == "Error: no tuna catch found"){
	// 			$("#error_holder").show();
	// 			$("#success_holder").hide();
	// 		} else{
	// 			$("#success_holder").show();
	// 			$("#error_holder").hide();
	// 		}
	// 	});
	// }

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllSupp = function(callback){

    	$http.get('/get_all_supp/').success(function(output){
			callback(output)
		});
	}

	// factory.queryTuna = function(id, callback){
 //    	$http.get('/get_tuna/'+id).success(function(output){
	// 		callback(output)
	// 	});
	// }

	factory.recordSupp = function(data, callback){

		data.location = data.longitude + ", "+ data.latitude;

		var supp = data.id + "-" + data.emailaddress + "-" + data.name + "-" + data.role + "-" + data.holder + "-" + data.balance;

    	$http.get('/add_supp/'+supp).success(function(output){
			callback(output)
		});
	}

	// factory.changeHolder = function(data, callback){

	// 	var holder = data.id + "-" + data.name;

 //    	$http.get('/change_holder/'+holder).success(function(output){
	// 		callback(output)
	// 	});
	// }

	return factory;
});


