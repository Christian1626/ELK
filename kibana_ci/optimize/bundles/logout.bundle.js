webpackJsonp([4],{

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
	__webpack_require__(1765);
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

/***/ 1765:
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	__webpack_require__(1070).setVisible(false).setRootController('logout', function ($http) {
	  $http.post('./api/shield/v1/logout', {}).then(function (response) {
	    return window.location.href = './';
	  });
	});

/***/ }

});