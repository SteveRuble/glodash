package glodash

/**
* Creates a function that invokes `func` with arguments arranged according
* to the specified `indexes` where the argument value at the first index is
* provided as the first argument, the argument value at the second index is
* provided as the second argument, and so on.
*
* @static
* @memberOf _
* @since 3.0.0
* @category Function
* @param {Function} func The function to rearrange arguments for.
* @param {...(number|number[])} indexes The arranged argument indexes.
* @returns {Function} Returns the new function.
* @example
*
* var rearged = _.rearg(function(a, b, c) {
*   return [a, b, c];
* }, [2, 0, 1]);
*
* rearged('b', 'c', 'a')
* // => ['a', 'b', 'c']
*/
func Rearg[T any](fn func(T) bool, indexes any /*...(number|number[])*/) func(T) bool {
	var out func(T) bool
	return out
}
	