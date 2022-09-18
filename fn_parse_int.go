package glodash

/**
* Converts `string` to an integer of the specified radix. If `radix` is
* `undefined` or `0`, a `radix` of `10` is used unless `value` is a
* hexadecimal, in which case a `radix` of `16` is used.
*
* **Note:** This method aligns with the
* [ES5 implementation](https://es5.github.io/#x15.1.2.2) of `parseInt`.
*
* @static
* @memberOf _
* @since 1.1.0
* @category String
* @param {string} string The string to convert.
* @param {number} [radix=10] The radix to interpret `value` by.
* @param- {Object} [guard] Enables use as an iteratee for methods like `_.map`.
* @returns {number} Returns the converted integer.
* @example
*
* _.parseInt('08');
* // => 8
*
* _.map(['6', '08', '10'], _.parseInt);
* // => [6, 8, 10]
*/
func ParseInt[T any](str string, radix int) int {
	var out int
	return out
}
	