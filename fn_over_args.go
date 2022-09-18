package glodash

/**
* Creates a function that invokes `func` with its arguments transformed.
*
* @static
* @since 4.0.0
* @memberOf _
* @category Function
* @param {Function} func The function to wrap.
* @param {...(Function|Function[])} [transforms=[_.identity]]
*  The argument transforms.
* @returns {Function} Returns the new function.
* @example
*
* function doubled(n) {
*   return n * 2;
* }
*
* function square(n) {
*   return n * n;
* }
*
* var func = _.overArgs(function(x, y) {
*   return [x, y];
* }, [square, doubled]);
*
* func(9, 3);
* // => [81, 6]
*
* func(10, 5);
* // => [100, 10]
*/
func OverArgs[T any](fn func(T) bool, transforms any /*...(Function|Function[])*/) func(T) bool {
	var out func(T) bool
	return out
}
	