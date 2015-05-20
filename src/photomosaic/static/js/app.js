(function() {
  'use strict';

  console.log('Congrats! Your JS is running!');

  var app = angular.module('photoMosaic', ['ngFileUpload']);

  app.controller('UploadController', ['$scope', '$http', 'Upload', function ($scope, $http, Upload) {
    $scope.log = '';
    $scope.targetFiles = [];
    $scope.tileFiles = [];

    $scope.$watch('targetFiles', function () {
      if ($scope.targetFiles) {
        uploadFiles($scope.targetFiles, 'target')
      }
    });

    $scope.$watch('tileFiles', function () {
      if ($scope.tileFiles) {
        uploadFiles($scope.tileFiles, 'tile')
      }
    });


    var uploadFiles = function (files, dataName) {
      if (files && files.length) {
        for (var i = 0; i < files.length; i ++) {
          var file = files[i];
          Upload.upload({
            url: '/process',
            file: file,
            fileFormDataName: dataName
          }).progress(function (evt) {
            var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
            $scope.log = 'progress: ' + progressPercentage + '% ' + evt.config.file.name + '\n' + $scope.log;
          }).success(function (data, status, headers, config) {
            $scope.log = config.file.name + ' uploaded.\n' + $scope.log;
          });
        }
      };
    };

    // $scope.process = function () {
    //   console.log('hi');
    //   $http.get('/image').
    //     success(function (data) {
    //       console.log(data);
    //     }).
    //     error(function (data, status, headers, config) {
    //       console.log(data);
    //     });
    // };

  }]);
}());