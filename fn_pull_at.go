package glodash

/**
* Removes elements from `array` corresponding to `indexes` and returns an
* array of removed elements.
*
* **Note:** Unlike `_.at`, this method mutates `array`.
*
* @static
* @memberOf _
* @since 3.0.0
* @category Array
* @param {Array} array The array to modify.
* @param {...(number|number[])} [indexes] The indexes of elements to remove.
* @returns {Array} Returns the new array of removed elements.
* @example
*
* var array = ['a', 'b', 'c', 'd'];
* var pulled = _.pullAt(array, [1, 3]);
*
* console.log(array);
* // => ['a', 'c']
*
* console.log(pulled);
* // => ['b', 'd']
*/
func PullAt[T any](array []T, indexes any /*...(number|number[])*/) []T {
	var out []T
	return out
}
	