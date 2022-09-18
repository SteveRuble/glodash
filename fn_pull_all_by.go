package glodash

/**
* This method is like `_.pullAll` except that it accepts `iteratee` which is
* invoked for each element of `array` and `values` to generate the criterion
* by which they're compared. The iteratee is invoked with one argument: (value).
*
* **Note:** Unlike `_.differenceBy`, this method mutates `array`.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Array
* @param {Array} array The array to modify.
* @param {Array} values The values to remove.
* @param {Function} [iteratee=_.identity] The iteratee invoked per element.
* @returns {Array} Returns `array`.
* @example
*
* var array = [{ 'x': 1 }, { 'x': 2 }, { 'x': 3 }, { 'x': 1 }];
*
* _.pullAllBy(array, [{ 'x': 1 }, { 'x': 3 }], 'x');
* console.log(array);
* // => [{ 'x': 2 }]
*/
func PullAllBy[T any](array []T, values []T, iteratee func(T) bool) []T {
	var out []T
	return out
}
	