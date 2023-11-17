/**
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we
get 3, 5, 6 & 9. The sum of these multiples is 23. Find the sum of all the
multiples of 3 or 5 below 1000.
*/

#include <iostream>
using namespace std;

class Solution {
public:
  /**
   * The function calculates the sum of the digits of a given number.
   *
   * @param num The parameter "num" is an integer representing the number for
   * which we want to calculate the sum of its digits.
   *
   * @return The sum of the digits of the given number.
   */
  int sumOfDigits(int num) {
    int sum = 0;
    while (num > 0) {
      sum += num % 10;
      num /= 10;
    }
    return sum;
  }

  /**
   * The function checks if a given number is a multiple of 3 by calculating the
   * sum of its digits and returning true if the sum is divisible by 3,
   * otherwise it returns false.
   *
   * @param num An integer number that we want to check if it is a multiple
   * of 3.
   *
   * @return a boolean value. It returns true if the given number is a multiple
   * of 3, and false otherwise.
   */
  bool isMultipleOf3(int num) {
    int digitSum = sumOfDigits(num);
    if (digitSum % 3 != 0) {
      return false;
    }
    return true;
  }

  /**
   * The function checks if a given number is a multiple of 5 by checking if its
   * unit digit is either 0 or 5.
   *
   * @param num An integer number that we want to check if it is a multiple
   * of 5.
   *
   * @return a boolean value. It returns true if the given number is a multiple
   * of 5 (i.e., the unit digit is either 0 or 5), and false otherwise.
   */
  bool isMultipleOf5(int num) {
    int unitDigit = num % 10;
    if (unitDigit == 0 || unitDigit == 5) {
      return true;
    }
    return false;
  }

  /**
   * The function calculates the sum of all numbers below a given number that
   * are multiples of either 3 or 5.
   *
   * @param num The parameter "num" represents the upper limit or the maximum
   * number up to which we want to find the sum of multiples of 3 or 5. In this
   * case, the function is finding the sum of all multiples of 3 or 5 below
   * 1000.
   *
   * @return the sum of all the multiples of 3 or 5 below the given number.
   */
  int sumOfMultiplesOf3or5Below1000(int num) {
    int sum = 0;
    for (int i = 3; i < num; i++) {
      if ((isMultipleOf3(i)) || (isMultipleOf5(i))) {
        sum += i;
      }
    }
    return sum;
  }
};

int main() {
  cout << "\tSum of multiples of 3 or 5 below 1000:";
  Solution solution;
  cout << "\t" << solution.sumOfMultiplesOf3or5Below1000(1000) << endl;
  return 0;
}