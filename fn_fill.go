package glodash

/**
* Fills elements of `array` with `value` from `start` up to, but not
* including, `end`.
*
* **Note:** This method mutates `array`.
*
* @static
* @memberOf _
* @since 3.2.0
* @category Array
* @param {Array} array The array to fill.
* @param {*} value The value to fill `array` with.
* @param {number} [start=0] The start position.
* @param {number} [end=array.length] The end position.
* @returns {Array} Returns `array`.
* @example
*
* var array = [1, 2, 3];
*
* _.fill(array, 'a');
* console.log(array);
* // => ['a', 'a', 'a']
*
* _.fill(Array(3), 2);
* // => [2, 2, 2]
*
* _.fill([4, 6, 8, 10], '*', 1, 3);
* // => [4, '*', '*', 10]
*/
func Fill[T any](array []T, value any /***/, start int, end int) []T {
	var out []T
	return out
}
	