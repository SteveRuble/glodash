package glodash

/**
* Gets `n` random elements at unique keys from `collection` up to the
* size of `collection`.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Collection
* @param {Array|Object} collection The collection to sample.
* @param {number} [n=1] The number of elements to sample.
* @param- {Object} [guard] Enables use as an iteratee for methods like `_.map`.
* @returns {Array} Returns the random elements.
* @example
*
* _.sampleSize([1, 2, 3], 2);
* // => [3, 1]
*
* _.sampleSize([1, 2, 3], 4);
* // => [2, 3, 1]
*/
func SampleSize[T any](collection any /*Array|Object*/, n int) []T {
	var out []T
	return out
}
	