package glodash

/**
* Checks if `n` is between `start` and up to, but not including, `end`. If
* `end` is not specified, it's set to `start` with `start` then set to `0`.
* If `start` is greater than `end` the params are swapped to support
* negative ranges.
*
* @static
* @memberOf _
* @since 3.3.0
* @category Number
* @param {number} number The number to check.
* @param {number} [start=0] The start of the range.
* @param {number} end The end of the range.
* @returns {boolean} Returns `true` if `number` is in the range, else `false`.
* @see _.range, _.rangeRight
* @example
*
* _.inRange(3, 2, 4);
* // => true
*
* _.inRange(4, 8);
* // => true
*
* _.inRange(4, 2);
* // => false
*
* _.inRange(2, 2);
* // => false
*
* _.inRange(1.2, 2);
* // => true
*
* _.inRange(5.2, 4);
* // => false
*
* _.inRange(-3, -2, -6);
* // => true
*/
func InRange[T any](number int, start int, end int) any /*boolean*/ {
	var out any /*boolean*/
	return out
}
	