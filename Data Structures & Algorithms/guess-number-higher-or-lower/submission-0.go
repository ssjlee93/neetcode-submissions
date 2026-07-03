/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    // easy binay search 
    low, high := 0, n
    for low <= high {
        mid := low + (high-low)/2
        
        if guess(mid) > 0 {
            low = mid +1
        } else if guess(mid) < 0 {
            high = mid - 1
        } else {
            return mid
        }
    }
    return -1
}
