package glodash

/**
* Creates a function that invokes `func` with the `this` binding of the
* created function and arguments from `start` and beyond provided as
* an array.
*
* **Note:** This method is based on the
* [rest parameter](https://mdn.io/rest_parameters).
*
* @static
* @memberOf _
* @since 4.0.0
* @category Function
* @param {Function} func The function to apply a rest parameter to.
* @param {number} [start=func.length-1] The start position of the rest parameter.
* @returns {Function} Returns the new function.
* @example
*
* var say = _.rest(function(what, names) {
*   return what + ' ' + _.initial(names).join(', ') +
*     (_.size(names) > 1 ? ', & ' : '') + _.last(names);
* });
*
* say('hello', 'fred', 'barney', 'pebbles');
* // => 'hello fred, barney, & pebbles'
*/
func Rest[T any](fn func(T) bool, start int) func(T) bool {
	var out func(T) bool
	return out
}
	