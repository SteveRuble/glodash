package glodash

/**
* Recursively flatten `array` up to `depth` times.
*
* @static
* @memberOf _
* @since 4.4.0
* @category Array
* @param {Array} array The array to flatten.
* @param {number} [depth=1] The maximum recursion depth.
* @returns {Array} Returns the new flattened array.
* @example
*
* var array = [1, [2, [3, [4]], 5]];
*
* _.flattenDepth(array, 1);
* // => [1, 2, [3, [4]], 5]
*
* _.flattenDepth(array, 2);
* // => [1, 2, 3, [4], 5]
*/
func FlattenDepth[T any](array []T, depth int) []T {
	var out []T
	return out
}
	