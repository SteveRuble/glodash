package glodash

/**
* Creates a function that accepts up to one argument, ignoring any
* additional arguments.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Function
* @param {Function} func The function to cap arguments for.
* @returns {Function} Returns the new capped function.
* @example
*
* _.map(['6', '8', '10'], _.unary(parseInt));
* // => [6, 8, 10]
*/
func Unary[T any](fn func(T) bool) func(T) bool {
	var out func(T) bool
	return out
}
	