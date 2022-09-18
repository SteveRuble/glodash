package glodash

/**
* Creates a slice of `array` with `n` elements dropped from the end.
*
* @static
* @memberOf _
* @since 3.0.0
* @category Array
* @param {Array} array The array to query.
* @param {number} [n=1] The number of elements to drop.
* @param- {Object} [guard] Enables use as an iteratee for methods like `_.map`.
* @returns {Array} Returns the slice of `array`.
* @example
*
* _.dropRight([1, 2, 3]);
* // => [1, 2]
*
* _.dropRight([1, 2, 3], 2);
* // => [1]
*
* _.dropRight([1, 2, 3], 5);
* // => []
*
* _.dropRight([1, 2, 3], 0);
* // => [1, 2, 3]
*/
func DropRight[T any](array []T, n int) []T {
	var out []T
	return out
}
	