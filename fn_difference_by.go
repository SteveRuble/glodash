package glodash

/**
* This method is like `_.difference` except that it accepts `iteratee` which
* is invoked for each element of `array` and `values` to generate the criterion
* by which they're compared. The order and references of result values are
* determined by the first array. The iteratee is invoked with one argument:
* (value).
*
* **Note:** Unlike `_.pullAllBy`, this method returns a new array.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Array
* @param {Array} array The array to inspect.
* @param {...Array} [values] The values to exclude.
* @param {Function} [iteratee=_.identity] The iteratee invoked per element.
* @returns {Array} Returns the new array of filtered values.
* @example
*
* _.differenceBy([2.1, 1.2], [2.3, 3.4], Math.floor);
* // => [1.2]
*
* // The `_.property` iteratee shorthand.
* _.differenceBy([{ 'x': 2 }, { 'x': 1 }], [{ 'x': 1 }], 'x');
* // => [{ 'x': 2 }]
*/
func DifferenceBy[T any](array []T, values any /*...Array*/, iteratee func(T) bool) []T {
	var out []T
	return out
}
	