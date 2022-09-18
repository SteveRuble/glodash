package glodash

/**
* Invokes `func` after `wait` milliseconds. Any additional arguments are
* provided to `func` when it's invoked.
*
* @static
* @memberOf _
* @since 0.1.0
* @category Function
* @param {Function} func The function to delay.
* @param {number} wait The number of milliseconds to delay invocation.
* @param {...*} [args] The arguments to invoke `func` with.
* @returns {number} Returns the timer id.
* @example
*
* _.delay(function(text) {
*   console.log(text);
* }, 1000, 'later');
* // => Logs 'later' after one second.
*/
func Delay[T any](fn func(T) bool, wait int, args any /*...**/) int {
	var out int
	return out
}
	