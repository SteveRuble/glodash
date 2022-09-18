package glodash

/**
* Adds all own enumerable string keyed function properties of a source
* object to the destination object. If `object` is a function, then methods
* are added to its prototype as well.
*
* **Note:** Use `_.runInContext` to create a pristine `lodash` function to
* avoid conflicts caused by modifying the original.
*
* @static
* @since 0.1.0
* @memberOf _
* @category Util
* @param {Function|Object} [object=lodash] The destination object.
* @param {Object} source The object of functions to add.
* @param {Object} [options={}] The options object.
* @param {boolean} [options.chain=true] Specify whether mixins are chainable.
* @returns {Function|Object} Returns `object`.
* @example
*
* function vowels(string) {
*   return _.filter(string, function(v) {
*     return /[aeiou]/i.test(v);
*   });
* }
*
* _.mixin({ 'vowels': vowels });
* _.vowels('fred');
* // => ['e']
*
* _('fred').vowels().value();
* // => ['e']
*
* _.mixin({ 'vowels': vowels }, { 'chain': false });
* _('fred').vowels();
* // => ['e']
*/
func Mixin[T any](object any /*Function|Object*/, source T, options T, options any /*boolean*/) any /*Function|Object*/ {
	var out any /*Function|Object*/
	return out
}
	