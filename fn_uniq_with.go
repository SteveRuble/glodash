package glodash

/**
* This method is like `_.uniq` except that it accepts `comparator` which
* is invoked to compare elements of `array`. The order of result values is
* determined by the order they occur in the array.The comparator is invoked
* with two arguments: (arrVal, othVal).
*
* @static
* @memberOf _
* @since 4.0.0
* @category Array
* @param {Array} array The array to inspect.
* @param {Function} [comparator] The comparator invoked per element.
* @returns {Array} Returns the new duplicate free array.
* @example
*
* var objects = [{ 'x': 1, 'y': 2 }, { 'x': 2, 'y': 1 }, { 'x': 1, 'y': 2 }];
*
* _.uniqWith(objects, _.isEqual);
* // => [{ 'x': 1, 'y': 2 }, { 'x': 2, 'y': 1 }]
*/
func UniqWith[T any](array []T, comparator func(T) bool) []T {
	var out []T
	return out
}
	