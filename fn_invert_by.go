package glodash

/**
* This method is like `_.invert` except that the inverted object is generated
* from the results of running each element of `object` thru `iteratee`. The
* corresponding inverted value of each inverted key is an array of keys
* responsible for generating the inverted value. The iteratee is invoked
* with one argument: (value).
*
* @static
* @memberOf _
* @since 4.1.0
* @category Object
* @param {Object} object The object to invert.
* @param {Function} [iteratee=_.identity] The iteratee invoked per element.
* @returns {Object} Returns the new inverted object.
* @example
*
* var object = { 'a': 1, 'b': 2, 'c': 1 };
*
* _.invertBy(object);
* // => { '1': ['a', 'c'], '2': ['b'] }
*
* _.invertBy(object, function(value) {
*   return 'group' + value;
* });
* // => { 'group1': ['a', 'c'], 'group2': ['b'] }
*/
func InvertBy[T any](object T, iteratee func(T) bool) T {
	var out T
	return out
}
	