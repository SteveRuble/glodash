package glodash

/**
* This method invokes `interceptor` and returns `value`. The interceptor
* is invoked with one argument; (value). The purpose of this method is to
* "tap into" a method chain sequence in order to modify intermediate results.
*
* @static
* @memberOf _
* @since 0.1.0
* @category Seq
* @param {*} value The value to provide to `interceptor`.
* @param {Function} interceptor The function to invoke.
* @returns {*} Returns `value`.
* @example
*
* _([1, 2, 3])
*  .tap(function(array) {
*    // Mutate input array.
*    array.pop();
*  })
*  .reverse()
*  .value();
* // => [2, 1]
*/
func Tap[T any](value any /***/, interceptor func(T) bool) any /***/ {
	var out any /***/
	return out
}
	