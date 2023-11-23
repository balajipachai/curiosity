/**
 * A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 * 99. Find
the largest palindrome made from the product of two 3-digit numbers.
*/

#include <iostream>

using namespace std;

class Solution {
public:
  /**
   * The function checks if a given number is a palindrome or not.
   *
   * @param number The parameter "number" is an integer that represents the
   * number we want to check if it is a palindrome.
   *
   * @return a boolean value indicating whether the given number is a palindrome
   * or not.
   */
  bool isPallindrome(int number) {
    int reverse = 0;
    int x = number;
    while (x > 0) {
      reverse = 10 * reverse + x % 10;
      x /= 10;
    }
    return reverse == number;
  }

  int largestPallindrome() {
    int max = -1;
    int p = 0;
    int n1, n2 = 0;
    for (int i = 999; i >= 100; i--) {
      if (max >= i * 999) {
        break;
      }
      for (int j = 999; j >= 100; j--) {
        p = i * j;
        if (max < p && isPallindrome(p)) {
          max = p;
          n2 = j;
        }
      }
      n1 = i;
    }
    cout << "\tThe two 3-digit numbers are\t" << n1 << " , " << n2 << endl;
    return max;
  }
};

int main() {
  Solution solution;
  int answer = solution.largestPallindrome();
  cout << "\tLargest pallindrome product of a 3-digit number is: " << answer
       << endl;
}