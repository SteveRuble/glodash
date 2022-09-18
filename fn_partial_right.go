package glodash

/**
* This method is like `_.partial` except that partially applied arguments
* are appended to the arguments it receives.
*
* The `_.partialRight.placeholder` value, which defaults to `_` in monolithic
* builds, may be used as a placeholder for partially applied arguments.
*
* **Note:** This method doesn't set the "length" property of partially
* applied functions.
*
* @static
* @memberOf _
* @since 1.0.0
* @category Function
* @param {Function} func The function to partially apply arguments to.
* @param {...*} [partials] The arguments to be partially applied.
* @returns {Function} Returns the new partially applied function.
* @example
*
* function greet(greeting, name) {
*   return greeting + ' ' + name;
* }
*
* var greetFred = _.partialRight(greet, 'fred');
* greetFred('hi');
* // => 'hi fred'
*
* // Partially applied with placeholders.
* var sayHelloTo = _.partialRight(greet, 'hello', _);
* sayHelloTo('fred');
* // => 'hello fred'
*/
func PartialRight[T any](fn func(T) bool, partials any /*...**/) func(T) bool {
	var out func(T) bool
	return out
}
	