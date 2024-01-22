/**

Given an integer num, return a string of its base 7 representation.
Example 1:

Input: num = 100
Output: "202"
Example 2:

Input: num = -7
Output: "-10"

Constraints:
-10^7 <= num <= 10^7
*/

#include <iostream>
using namespace std;

class Solution {
public:
  /**
   * The function "convertToBase7" takes an integer as input and returns its
   * base 7 representation as a string.
   *
   * @param num The parameter "num" is an integer that represents the number to
   * be converted to base 7.
   *
   * @return a string representation of the given number in base 7.
   */
  string convertToBase7(int num) {
    string result = "";
    int quotient;
    int remainder;
    bool isNegative = false;

    if (num == 0) {
      return "0";
    }

    if (num < 0) {
      isNegative = true;
      num = abs(num);
    }

    while (num > 0) {
      quotient = num / 7;
      remainder = num % 7;
      result += to_string(remainder);
      num = quotient;
    }
    // Reverse the string, as the conversion is the inverse of the remainder
    reverse(result.begin(), result.end());
    if (isNegative) {
      result = "-" + result;
    }
    return result;
  }
};

int main() {
  Solution solution;
  int num = -7;
  cout << "\tEntered number is: " << num;
  printf("\tBase 7 representation of %d is = ", num);
  string result = solution.convertToBase7(num);
  cout << result << endl;

  num = 100;
  cout << "\tEntered number is: " << num;
  printf("\tBase 7 representation of %d is = ", num);
  result = solution.convertToBase7(num);
  cout << result << endl;

  num = 0;
  cout << "\tEntered number is: " << num;
  printf("\tBase 7 representation of %d is = ", num);
  result = solution.convertToBase7(num);
  cout << result << endl;

  return 0;
}