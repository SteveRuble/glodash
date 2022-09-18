package glodash

/**
* Creates a function that iterates over `pairs` and invokes the corresponding
* function of the first predicate to return truthy. The predicate-function
* pairs are invoked with the `this` binding and arguments of the created
* function.
*
* @static
* @memberOf _
* @since 4.0.0
* @category Util
* @param {Array} pairs The predicate-function pairs.
* @returns {Function} Returns the new composite function.
* @example
*
* var func = _.cond([
*   [_.matches({ 'a': 1 }),           _.constant('matches A')],
*   [_.conforms({ 'b': _.isNumber }), _.constant('matches B')],
*   [_.stubTrue,                      _.constant('no match')]
* ]);
*
* func({ 'a': 1, 'b': 2 });
* // => 'matches A'
*
* func({ 'a': 0, 'b': 1 });
* // => 'matches B'
*
* func({ 'a': '1', 'b': '2' });
* // => 'no match'
*/
func Cond[T any](pairs []T) func(T) bool {
	var out func(T) bool
	return out
}
	