package glodash

/**
* This method is like `_.xor` except that it accepts `comparator` which is
* invoked to compare elements of `arrays`. The order of result values is
* determined by the order they occur in the arrays. The comparator is invoked
* with two arguments: (arrVal, othVal).
*
* @static
* @memberOf _
* @since 4.0.0
* @category Array
* @param {...Array} [arrays] The arrays to inspect.
* @param {Function} [comparator] The comparator invoked per element.
* @returns {Array} Returns the new array of filtered values.
* @example
*
* var objects = [{ 'x': 1, 'y': 2 }, { 'x': 2, 'y': 1 }];
* var others = [{ 'x': 1, 'y': 1 }, { 'x': 1, 'y': 2 }];
*
* _.xorWith(objects, others, _.isEqual);
* // => [{ 'x': 2, 'y': 1 }, { 'x': 1, 'y': 1 }]
*/
func XorWith[T any](arrays any /*...Array*/, comparator func(T) bool) []T {
	var out []T
	return out
}
	