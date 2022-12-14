package glodash

/**
* This method is like `_.cloneWith` except that it recursively clones `value`.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Lang
* @param {*} value The value to recursively clone.
* @param {Function} [customizer] The function to customize cloning.
* @returns {*} Returns the deep cloned value.
* @see _.cloneWith
* @example
*
* function customizer(value) {
*   if (_.isElement(value)) {
*     return value.cloneNode(true);
*   }
* }
*
* var el = _.cloneDeepWith(document.body, customizer);
*
* console.log(el === document.body);
* // => false
* console.log(el.nodeName);
* // => 'BODY'
* console.log(el.childNodes.length);
* // => 20
*/
func CloneDeepWith[T any](value any /***/, customizer func(T) bool) any /***/ {
	var out any /***/
	return out
}
	