package glodash

/**
* Creates a function that invokes the method at `path` of a given object.
* Any additional arguments are provided to the invoked method.
*
* @static
* @memberOf _
* @since 3.7.0
* @category Util
* @param {Array|string} path The path of the method to invoke.
* @param {...*} [args] The arguments to invoke the method with.
* @returns {Function} Returns the new invoker function.
* @example
*
* var objects = [
*   { 'a': { 'b': _.constant(2) } },
*   { 'a': { 'b': _.constant(1) } }
* ];
*
* _.map(objects, _.method('a.b'));
* // => [2, 1]
*
* _.map(objects, _.method(['a', 'b']));
* // => [2, 1]
*/
func Method[T any](path any /*Array|string*/, args any /*...**/) func(T) bool {
	var out func(T) bool
	return out
}
	