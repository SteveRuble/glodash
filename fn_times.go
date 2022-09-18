package glodash

/**
* Invokes the iteratee `n` times, returning an array of the results of
* each invocation. The iteratee is invoked with one argument; (index).
*
* @static
* @since 0.1.0
* @memberOf _
* @category Util
* @param {number} n The number of times to invoke `iteratee`.
* @param {Function} [iteratee=_.identity] The function invoked per iteration.
* @returns {Array} Returns the array of results.
* @example
*
* _.times(3, String);
* // => ['0', '1', '2']
*
*  _.times(4, _.constant(0));
* // => [0, 0, 0, 0]
*/
func Times[T any](n int, iteratee func(T) bool) []T {
	var out []T
	return out
}
	