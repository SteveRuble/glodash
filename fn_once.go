package glodash

/**
* Creates a function that is restricted to invoking `func` once. Repeat calls
* to the function return the value of the first invocation. The `func` is
* invoked with the `this` binding and arguments of the created function.
*
* @static
* @memberOf _
* @since 0.1.0
* @category Function
* @param {Function} func The function to restrict.
* @returns {Function} Returns the new restricted function.
* @example
*
* var initialize = _.once(createApplication);
* initialize();
* initialize();
* // => `createApplication` is invoked once
*/
func Once[T any](fn func(T) bool) func(T) bool {
	var out func(T) bool
	return out
}
	