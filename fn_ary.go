package glodash

/**
* Creates a function that invokes `func`, with up to `n` arguments,
* ignoring any additional arguments.
*
* @static
* @memberOf _
* @since 3.0.0
* @category Function
* @param {Function} func The function to cap arguments for.
* @param {number} [n=func.length] The arity cap.
* @param- {Object} [guard] Enables use as an iteratee for methods like `_.map`.
* @returns {Function} Returns the new capped function.
* @example
*
* _.map(['6', '8', '10'], _.ary(parseInt, 1));
* // => [6, 8, 10]
*/
func Ary[T any](fn func(T) bool, n int) func(T) bool {
	var out func(T) bool
	return out
}
	