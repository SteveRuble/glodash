package glodash

/**
* Creates a slice of `array` with `n` elements dropped from the beginning.
*
* @static
* @memberOf _
* @since 0.5.0
* @category Array
* @param {Array} array The array to query.
* @param {number} [n=1] The number of elements to drop.
* @param- {Object} [guard] Enables use as an iteratee for methods like `_.map`.
* @returns {Array} Returns the slice of `array`.
* @example
*
* _.drop([1, 2, 3]);
* // => [2, 3]
*
* _.drop([1, 2, 3], 2);
* // => [3]
*
* _.drop([1, 2, 3], 5);
* // => []
*
* _.drop([1, 2, 3], 0);
* // => [1, 2, 3]
*/
func Drop[T any](array []T, n int) []T {
	var out []T
	return out
}
	