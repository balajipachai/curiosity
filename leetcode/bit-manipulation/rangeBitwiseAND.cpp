/**
 * Given two integers left and right that represent the range [left, right],
return the bitwise AND of all numbers in this range, inclusive.



Example 1:

Input: left = 5, right = 7
Output: 4
Example 2:

Input: left = 0, right = 0
Output: 0
Example 3:

Input: left = 1, right = 2147483647
Output: 0


Constraints:

0 <= left <= right <= 231 - 1
*/

#include <iostream>
using namespace std;

class Solution {
public:
  /**
   * The function calculates the common prefix of two integers by shifting them
   * right until they are equal, and then shifting the result left by the number
   * of shifts performed.
   *
   * @param m The parameter `m` represents the first integer value.
   * @param n The parameter `n` represents an integer value.
   *
   * @return the common prefix of the two input integers, m and n.
   */
  int usingCommonPrefix(int m, int n) {
    int shift = 0;
    while (m < n) {
      m = m >> 1;
      n = n >> 1;
      shift += 1;
    }
    return m << shift;
  }

  /**
   * The function uses the Brian Kernighan algorithm to find the bitwise AND of
   * two numbers m and n.
   *
   * @param m The parameter "m" represents the first integer value.
   * @param n The parameter "n" represents an integer value.
   *
   * @return the bitwise AND of the variables m and n.
   */
  int usingBrianKernighanAlgo(int m, int n) {
    while (m < n) {
      n = n & (n - 1);
    }
    return m & n;
  }

  /**
   * The function rangeBitwiseAnd calculates the bitwise AND of all numbers in
   * the range [left, right] using either the common prefix approach or the
   * Brian Kernighan algorithm, depending on the value of the whichApproach
   * parameter.
   *
   * @param left The starting number of the range.
   * @param right The "right" parameter represents the upper bound of the range
   * for which we want to perform the bitwise AND operation.
   * @param whichApproach A boolean variable that determines which approach to
   * use for calculating the bitwise AND of the range. If `whichApproach` is
   * `false` (0), the function will use the `usingCommonPrefix` approach. If
   * `whichApproach` is `true` (1), the function will use
   *
   * @return the result of either the "usingCommonPrefix" or
   * "usingBrianKernighanAlgo" function, depending on the value of the
   * "whichApproach" parameter.
   */
  int rangeBitwiseAnd(int left, int right, bool whichApproach) {
    // whichApproach = 0 -> usingCommonPrefix
    // whichApproach = 1 -> usingBrianKernighan algorithm
    if (!whichApproach) {
      return usingCommonPrefix(left, right);
    }
    return usingBrianKernighanAlgo(left, right);
  }
};

int main() {
  Solution solution;
  int left = 5, right = 7;
  cout << "\tCommon Prefix: range bitwise AND of " << left << " & " << right
       << " = " << solution.rangeBitwiseAnd(left, right, 0) << endl;
  cout << "\tBrian Kernighan: range bitwise AND of " << left << " & " << right
       << " = " << solution.rangeBitwiseAnd(left, right, 1) << endl;

  left = 0, right = 0;
  cout << "\tCommon Prefix: range bitwise AND of " << left << " & " << right
       << " = " << solution.rangeBitwiseAnd(left, right, 0) << endl;
  cout << "\tBrian Kernighan: range bitwise AND of " << left << " & " << right
       << " = " << solution.rangeBitwiseAnd(left, right, 1) << endl;

  left = 1, right = 2147483647;
  cout << "\tCommon Prefix: range bitwise AND of " << left << " & " << right
       << " = " << solution.rangeBitwiseAnd(left, right, 0) << endl;
  cout << "\tBrian Kernighan: range bitwise AND of " << left << " & " << right
       << " = " << solution.rangeBitwiseAnd(left, right, 1) << endl;
  return 0;
}