package glodash

/**
* Creates a function that checks if **all** of the `predicates` return
* truthy when invoked with the arguments it receives.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Util
* @param {...(Function|Function[])} [predicates=[_.identity]]
*  The predicates to check.
* @returns {Function} Returns the new function.
* @example
*
* var func = _.overEvery([Boolean, isFinite]);
*
* func('1');
* // => true
*
* func(null);
* // => false
*
* func(NaN);
* // => false
*/
func OverEvery[T any](predicates any /*...(Function|Function[])*/) func(T) bool {
	var out func(T) bool
	return out
}
	