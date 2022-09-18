package glodash

/**
* This method is like `_.isMatch` except that it accepts `customizer` which
* is invoked to compare values. If `customizer` returns `undefined`, comparisons
* are handled by the method instead. The `customizer` is invoked with five
* arguments: (objValue, srcValue, index|key, object, source).
*
* @static
* @memberOf _
* @since 4.0.0
* @category Lang
* @param {Object} object The object to inspect.
* @param {Object} source The object of property values to match.
* @param {Function} [customizer] The function to customize comparisons.
* @returns {boolean} Returns `true` if `object` is a match, else `false`.
* @example
*
* function isGreeting(value) {
*   return /^h(?:i|ello)$/.test(value);
* }
*
* function customizer(objValue, srcValue) {
*   if (isGreeting(objValue) && isGreeting(srcValue)) {
*     return true;
*   }
* }
*
* var object = { 'greeting': 'hello' };
* var source = { 'greeting': 'hi' };
*
* _.isMatchWith(object, source, customizer);
* // => true
*/
func IsMatchWith[T any](object T, source T, customizer func(T) bool) any /*boolean*/ {
	var out any /*boolean*/
	return out
}
	