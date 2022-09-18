package glodash

/**
* This method is like `_.zip` except that it accepts an array of grouped
* elements and creates an array regrouping the elements to their pre-zip
* configuration.
*
* @static
* @memberOf _
* @since 1.2.0
* @category Array
* @param {Array} array The array of grouped elements to process.
* @returns {Array} Returns the new array of regrouped elements.
* @example
*
* var zipped = _.zip(['a', 'b'], [1, 2], [true, false]);
* // => [['a', 1, true], ['b', 2, false]]
*
* _.unzip(zipped);
* // => [['a', 'b'], [1, 2], [true, false]]
*/
func Unzip[T any](array []T) []T {
	var out []T
	return out
}
	