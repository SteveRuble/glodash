package glodash

/**
* This method is like `_.zip` except that it accepts `iteratee` to specify
* how grouped values should be combined. The iteratee is invoked with the
* elements of each group: (...group).
*
* @static
* @memberOf _
* @since 3.8.0
* @category Array
* @param {...Array} [arrays] The arrays to process.
* @param {Function} [iteratee=_.identity] The function to combine
*  grouped values.
* @returns {Array} Returns the new array of grouped elements.
* @example
*
* _.zipWith([1, 2], [10, 20], [100, 200], function(a, b, c) {
*   return a + b + c;
* });
* // => [111, 222]
*/
func ZipWith[T any](arrays any /*...Array*/, iteratee func(T) bool) []T {
	var out []T
	return out
}
	