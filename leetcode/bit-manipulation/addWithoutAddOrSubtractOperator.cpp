/**
 * Given two integers a and b, return the sum of the two integers without using
the operators + and -.

Example 1:

Input: a = 1, b = 2
Output: 3
Example 2:

Input: a = 2, b = 3
Output: 5

Constraints:

-1000 <= a, b <= 1000
*/
#include <iostream>
using namespace std;

class Solution {
public:
  /**
   * The function `getSum` adds two numbers without using the `+` or `-`
   * operator by performing bitwise operations.
   *
   * @param a The first number to be added.
   * @param b The parameter "b" is an integer value that represents the second
   * number to be added.
   *
   * @return the sum of the two input integers, `a` and `b`.
   */
  int getSum(int a, int b) {
    /**
     * To add 2 numbers without using + or - operator
     * Doing a ^ b will give us the correct sum iff a,b don't have set bits at
     * same position What if a, b have same bits at same position? In that case
     * do bitwise_and obtain the caryy carry = a & b Shifting carry by 1 bit
     * gives the required sum after adding it to a
     */
    while (b != 0) {
      unsigned carry = a & b;
      a = a ^ b;
      b = carry << 1;
    }
    return a;
  }
};

int main() {
  Solution solution;
  int result;

  int a = 12, b = 34;
  cout << "\tEntered number are a = " << a << " , b = " << b << endl;
  result = solution.getSum(a, b);
  cout << "\t Sum(a,b) = " << result << endl;

  a = -12, b = 34;
  cout << "\tEntered number are a = " << a << " , b = " << b << endl;
  result = solution.getSum(a, b);
  cout << "\t Sum(a,b) = " << result << endl;
}