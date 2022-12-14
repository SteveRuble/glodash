package glodash

/**
* Attempts to invoke `func`, returning either the result or the caught error
* object. Any additional arguments are provided to `func` when it's invoked.
*
* @static
* @memberOf _
* @since 3.0.0
* @category Util
* @param {Function} func The function to attempt.
* @param {...*} [args] The arguments to invoke `func` with.
* @returns {*} Returns the `func` result or error object.
* @example
*
* // Avoid throwing errors for invalid selectors.
* var elements = _.attempt(function(selector) {
*   return document.querySelectorAll(selector);
* }, '>_>');
*
* if (_.isError(elements)) {
*   elements = [];
* }
*/
func Attempt[T any](fn func(T) bool, args any /*...**/) any /***/ {
	var out any /***/
	return out
}
	