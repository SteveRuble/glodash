package glodash

/**
* Creates a function that performs a partial deep comparison between the
* value at `path` of a given object to `srcValue`, returning `true` if the
* object value is equivalent, else `false`.
*
* **Note:** Partial comparisons will match empty array and empty object
* `srcValue` values against any array or object value, respectively. See
* `_.isEqual` for a list of supported value comparisons.
*
* @static
* @memberOf _
* @since 3.2.0
* @category Util
* @param {Array|string} path The path of the property to get.
* @param {*} srcValue The value to match.
* @returns {Function} Returns the new spec function.
* @example
*
* var objects = [
*   { 'a': 1, 'b': 2, 'c': 3 },
*   { 'a': 4, 'b': 5, 'c': 6 }
* ];
*
* _.find(objects, _.matchesProperty('a', 4));
* // => { 'a': 4, 'b': 5, 'c': 6 }
*/
func MatchesProperty[T any](path any /*Array|string*/, srcValue any /***/) func(T) bool {
	var out func(T) bool
	return out
}
	