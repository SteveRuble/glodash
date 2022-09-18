package glodash

/**
* This method is like `_.max` except that it accepts `iteratee` which is
* invoked for each element in `array` to generate the criterion by which
* the value is ranked. The iteratee is invoked with one argument: (value).
*
* @static
* @memberOf _
* @since 4.0.0
* @category Math
* @param {Array} array The array to iterate over.
* @param {Function} [iteratee=_.identity] The iteratee invoked per element.
* @returns {*} Returns the maximum value.
* @example
*
* var objects = [{ 'n': 1 }, { 'n': 2 }];
*
* _.maxBy(objects, function(o) { return o.n; });
* // => { 'n': 2 }
*
* // The `_.property` iteratee shorthand.
* _.maxBy(objects, 'n');
* // => { 'n': 2 }
*/
func MaxBy[T any](array []T, iteratee func(T) bool) any /***/ {
	var out any /***/
	return out
}
	