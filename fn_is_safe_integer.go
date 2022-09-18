package glodash

/**
* Checks if `value` is a safe integer. An integer is safe if it's an IEEE-754
* double precision number which isn't the result of a rounded unsafe integer.
*
* **Note:** This method is based on
* [`Number.isSafeInteger`](https://mdn.io/Number/isSafeInteger).
*
* @static
* @memberOf _
* @since 4.0.0
* @category Lang
* @param {*} value The value to check.
* @returns {boolean} Returns `true` if `value` is a safe integer, else `false`.
* @example
*
* _.isSafeInteger(3);
* // => true
*
* _.isSafeInteger(Number.MIN_VALUE);
* // => false
*
* _.isSafeInteger(Infinity);
* // => false
*
* _.isSafeInteger('3');
* // => false
*/
func IsSafeInteger[T any](value any /***/) any /*boolean*/ {
	var out any /*boolean*/
	return out
}
	