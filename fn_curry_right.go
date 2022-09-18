package glodash

/**
* This method is like `_.curry` except that arguments are applied to `func`
* in the manner of `_.partialRight` instead of `_.partial`.
*
* The `_.curryRight.placeholder` value, which defaults to `_` in monolithic
* builds, may be used as a placeholder for provided arguments.
*
* **Note:** This method doesn't set the "length" property of curried functions.
*
* @static
* @memberOf _
* @since 3.0.0
* @category Function
* @param {Function} func The function to curry.
* @param {number} [arity=func.length] The arity of `func`.
* @param- {Object} [guard] Enables use as an iteratee for methods like `_.map`.
* @returns {Function} Returns the new curried function.
* @example
*
* var abc = function(a, b, c) {
*   return [a, b, c];
* };
*
* var curried = _.curryRight(abc);
*
* curried(3)(2)(1);
* // => [1, 2, 3]
*
* curried(2, 3)(1);
* // => [1, 2, 3]
*
* curried(1, 2, 3);
* // => [1, 2, 3]
*
* // Curried with placeholders.
* curried(3)(1, _)(2);
* // => [1, 2, 3]
*/
func CurryRight[T any](fn func(T) bool, arity int) func(T) bool {
	var out func(T) bool
	return out
}
	