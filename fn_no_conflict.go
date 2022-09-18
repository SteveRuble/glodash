package glodash

/**
* Reverts the `_` variable to its previous value and returns a reference to
* the `lodash` function.
*
* @static
* @since 0.1.0
* @memberOf _
* @category Util
* @returns {Function} Returns the `lodash` function.
* @example
*
* var lodash = _.noConflict();
*/
func NoConflict[T any]() func(T) bool {
	var out func(T) bool
	return out
}
	