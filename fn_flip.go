package glodash

/**
* Creates a function that invokes `func` with arguments reversed.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Function
* @param {Function} func The function to flip arguments for.
* @returns {Function} Returns the new flipped function.
* @example
*
* var flipped = _.flip(function() {
*   return _.toArray(arguments);
* });
*
* flipped('a', 'b', 'c', 'd');
* // => ['d', 'c', 'b', 'a']
*/
func Flip[T any](fn func(T) bool) func(T) bool {
	var out func(T) bool
	return out
}
	