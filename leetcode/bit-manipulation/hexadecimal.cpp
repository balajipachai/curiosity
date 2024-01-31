/**
 Given an integer num, return a string representing its hexadecimal
representation. For negative integers, two’s complement method is used.

All the letters in the answer string should be lowercase characters, and there
should not be any leading zeros in the answer except for the zero itself.

Note: You are not allowed to use any built-in library method to directly solve
this problem.



Example 1:

Input: num = 26
Output: "1a"
Example 2:

Input: num = -1
Output: "ffffffff"


Constraints:

-231 <= num <= 231 - 1
*/
#include <iostream>
using namespace std;

class Solution {
public:
  /**
   * The function "toHex" converts a given integer into its hexadecimal
   * representation.
   *
   * @param num The parameter `num` is an integer that represents the number
   * to be converted to hexadecimal.
   *
   * @return a string representation of the input integer in hexadecimal
   * format.
   */
  string toHex(int num) {
    string result = "";
    int quotient;
    int remainder;
    bool isNegative = false;

    if (num == 0) {
      return "0";
    }

    if (num < 0) {
      u_int n = num;
      while (n) {
        quotient = n / 16;
        remainder = n % 16;
        if (remainder < 10) {
          result += tolower(remainder + 48);
        } else {
          result += tolower(remainder + 55);
        }
        n = quotient;
      }
    } else {
      while (num > 0) {
        quotient = num / 16;
        remainder = num % 16;
        if (remainder < 10) {
          result += tolower(remainder + 48);
        } else {
          result += tolower(remainder + 55);
        }
        num = quotient;
      }
    }
    // Reverse the string, as the conversion is the inverse of the remainder
    reverse(result.begin(), result.end());
    return result;
  }
};

int main() {
  Solution solution;
  string result;

  int num = 100;
  cout << "\tEntered number is: " << num;
  printf("\tHexadecimal representation of %d is = ", num);
  result = solution.toHex(num);
  cout << result << endl;

  num = -1;
  cout << "\tEntered number is: " << num;
  printf("\tHexadecimal representation of %d is = ", num);
  result = solution.toHex(num);
  cout << result << endl;

  num = 2545;
  cout << "\tEntered number is: " << num;
  printf("\tHexadecimal representation of %d is = ", num);
  result = solution.toHex(num);
  cout << result << endl;

  num = 0;
  cout << "\tEntered number is: " << num;
  printf("\tHexadecimal representation of %d is = ", num);
  result = solution.toHex(num);
  cout << result << endl;

  return 0;
}