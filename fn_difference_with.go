package glodash

/**
* This method is like `_.difference` except that it accepts `comparator`
* which is invoked to compare elements of `array` to `values`. The order and
* references of result values are determined by the first array. The comparator
* is invoked with two arguments: (arrVal, othVal).
*
* **Note:** Unlike `_.pullAllWith`, this method returns a new array.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Array
* @param {Array} array The array to inspect.
* @param {...Array} [values] The values to exclude.
* @param {Function} [comparator] The comparator invoked per element.
* @returns {Array} Returns the new array of filtered values.
* @example
*
* var objects = [{ 'x': 1, 'y': 2 }, { 'x': 2, 'y': 1 }];
*
* _.differenceWith(objects, [{ 'x': 1, 'y': 2 }], _.isEqual);
* // => [{ 'x': 2, 'y': 1 }]
*/
func DifferenceWith[T any](array []T, values any /*...Array*/, comparator func(T) bool) []T {
	var out []T
	return out
}
	