package glodash

/**
* The opposite of `_.before`; this method creates a function that invokes
* `func` once it's called `n` or more times.
*
* @static
* @memberOf _
* @since 0.1.0
* @category Function
* @param {number} n The number of calls before `func` is invoked.
* @param {Function} func The function to restrict.
* @returns {Function} Returns the new restricted function.
* @example
*
* var saves = ['profile', 'settings'];
*
* var done = _.after(saves.length, function() {
*   console.log('done saving!');
* });
*
* _.forEach(saves, function(type) {
*   asyncSave({ 'type': type, 'complete': done });
* });
* // => Logs 'done saving!' after the two async saves have completed.
*/
func After[T any](n int, fn func(T) bool) func(T) bool {
	var out func(T) bool
	return out
}
	