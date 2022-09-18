package glodash

/**
* Creates a function that invokes `func` with the `this` binding of `thisArg`
* and `partials` prepended to the arguments it receives.
*
* The `_.bind.placeholder` value, which defaults to `_` in monolithic builds,
* may be used as a placeholder for partially applied arguments.
*
* **Note:** Unlike native `Function#bind`, this method doesn't set the "length"
* property of bound functions.
*
* @static
* @memberOf _
* @since 0.1.0
* @category Function
* @param {Function} func The function to bind.
* @param {*} thisArg The `this` binding of `func`.
* @param {...*} [partials] The arguments to be partially applied.
* @returns {Function} Returns the new bound function.
* @example
*
* function greet(greeting, punctuation) {
*   return greeting + ' ' + this.user + punctuation;
* }
*
* var object = { 'user': 'fred' };
*
* var bound = _.bind(greet, object, 'hi');
* bound('!');
* // => 'hi fred!'
*
* // Bound with placeholders.
* var bound = _.bind(greet, object, _, '!');
* bound('hi');
* // => 'hi fred!'
*/
func Bind[T any](fn func(T) bool, thisArg any /***/, partials any /*...**/) func(T) bool {
	var out func(T) bool
	return out
}
	