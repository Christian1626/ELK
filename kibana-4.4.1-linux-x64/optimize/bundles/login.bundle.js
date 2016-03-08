webpackJsonp([3],{

/***/ 0:
/***/ function(module, exports, __webpack_require__) {

	
	/**
	 * Test entry file
	 *
	 * This is programatically created and updated, do not modify
	 *
	 * context: {"env":"production","urlBasePath":"","sourceMaps":false,"kbnVersion":"4.4.1","buildNum":9693}
	 * includes code from:
	 *  - elasticsearch@1.0.0
	 *  - kbn_vislib_vis_types@1.0.0
	 *  - kibana@1.0.0
	 *  - markdown_vis@1.0.0
	 *  - metric_vis@1.0.0
	 *  - shield@2.2.0
	 *  - spyModes@1.0.0
	 *  - statusPage@1.0.0
	 *  - table_vis@1.0.0
	 *
	 */

	'use strict';

	__webpack_require__(1070);
	__webpack_require__(1761);
	__webpack_require__(1597);
	__webpack_require__(1598);
	__webpack_require__(1599);
	__webpack_require__(1600);
	__webpack_require__(1601);
	__webpack_require__(1602);
	__webpack_require__(1603);
	__webpack_require__(1604);
	__webpack_require__(1605);
	__webpack_require__(1606);
	__webpack_require__(1607);
	__webpack_require__(1608);
	__webpack_require__(1609);
	__webpack_require__(1610);
	__webpack_require__(1611);
	__webpack_require__(1612);
	__webpack_require__(1597);
	__webpack_require__(1613);
	__webpack_require__(1614);
	__webpack_require__(1759);
	__webpack_require__(1070).bootstrap();
	/* xoxo */

/***/ },

/***/ 1761:
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	__webpack_require__(1762);
	var kibanaLogoUrl = __webpack_require__(1763);

	__webpack_require__(1070).setVisible(false).setRootTemplate(__webpack_require__(1764)).setRootController('login', function ($http) {
	  return {
	    kibanaLogoUrl: kibanaLogoUrl,
	    submit: function submit(username, password) {
	      var _this = this;

	      $http.post('./api/shield/v1/login', { username: username, password: password }).then(function (response) {
	        return window.location.href = './';
	      }, function (error) {
	        return _this.error = true;
	      });
	    }
	  };
	});

/***/ },

/***/ 1762:
/***/ function(module, exports) {

	// removed by extract-text-webpack-plugin

/***/ },

/***/ 1763:
/***/ function(module, exports, __webpack_require__) {

	module.exports = __webpack_require__.p + "installedPlugins/shield/public/images/kibana.svg"

/***/ },

/***/ 1764:
/***/ function(module, exports) {

	module.exports = "<div class=\"container\">\n  <h1><img ng-src=\"{{login.kibanaLogoUrl}}\" /></h1>\n\n  <form id=\"login-form\" ng-submit=\"login.submit(username, password)\">\n    <div ng-show=\"login.error\" class=\"form-group has-error\">\n      <label class=\"control-label\">Oops! Invalid username/password.</label>\n    </div>\n    <div class=\"form-group inner-addon left-addon\">\n      <i class=\"fa fa-user fa-lg fa-fw\"></i>\n      <input type=\"text\" ng-model=\"username\" class=\"form-control\" id=\"username\" name=\"username\" placeholder=\"Username\" />\n    </div>\n    <div class=\"form-group  inner-addon left-addon\">\n      <i class=\"fa fa-lock fa-lg fa-fw\"></i>\n      <input type=\"password\" ng-model=\"password\" class=\"form-control\" id=\"password\" name=\"password\" placeholder=\"Password\" />\n    </div>\n    <button type=\"submit\" ng-disabled=\"!username || !password\" class=\"btn btn-default login\">LOG IN</button>\n  </form>\n</div>"

/***/ }

});