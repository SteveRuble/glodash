package glodash

/**
* Creates a new array concatenating `array` with any additional arrays
* and/or values.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Array
* @param {Array} array The array to concatenate.
* @param {...*} [values] The values to concatenate.
* @returns {Array} Returns the new concatenated array.
* @example
*
* var array = [1];
* var other = _.concat(array, 2, [3], [[4]]);
*
* console.log(other);
* // => [1, 2, 3, [4]]
*
* console.log(array);
* // => [1]
*/
func Concat[T any](array []T, values any /*...**/) []T {
	var out []T
	return out
}
	