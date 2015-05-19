(function() {
  'use strict';

  console.log('Congrats! Your JS is running!');

  var app = angular.module('photoMosaic', ['ngFileUpload']);

  app.controller('UploadController', ['$scope', 'Upload', function ($scope, Upload) {
    $scope.$watch('files', function () {
      $scope.upload($scope.files);
    });
    $scope.log = '';

    $scope.upload = function (files) {
      if (files && files.length) {
        for (var i = 0; i < files.length; i++) {
          var file = files[i];
          Upload.upload({
            url: '/process',
            file: file
          }).progress(function (evt) {
            var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
            $scope.log = 'progress: ' + progressPercentage + '% ' +
                         evt.config.file.name + '\n' + $scope.log;
          }).success(function (data, status, headers, config) {
            $scope.log = 'file ' + config.file.name + 'uploaded. Response: ' + JSON.stringify(data) + '\n' + $scope.log;
            $scope.$apply();
          });
        }
      }
    };
  }]);
}());