package glodash

/**
* Creates a function that invokes `func` with `partials` prepended to the
* arguments it receives. This method is like `_.bind` except it does **not**
* alter the `this` binding.
*
* The `_.partial.placeholder` value, which defaults to `_` in monolithic
* builds, may be used as a placeholder for partially applied arguments.
*
* **Note:** This method doesn't set the "length" property of partially
* applied functions.
*
* @static
* @memberOf _
* @since 0.2.0
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
* var sayHelloTo = _.partial(greet, 'hello');
* sayHelloTo('fred');
* // => 'hello fred'
*
* // Partially applied with placeholders.
* var greetFred = _.partial(greet, _, 'fred');
* greetFred('hi');
* // => 'hi fred'
*/
func Partial[T any](fn func(T) bool, partials any /*...**/) func(T) bool {
	var out func(T) bool
	return out
}
	