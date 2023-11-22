/**
 * Given an array nums of integers, return how many of them contain an even
number of digits.



Example 1:

Input: nums = [12,345,2,6,7896]
Output: 2
Explanation:
12 contains 2 digits (even number of digits).
345 contains 3 digits (odd number of digits).
2 contains 1 digit (odd number of digits).
6 contains 1 digit (odd number of digits).
7896 contains 4 digits (even number of digits).
Therefore only 12 and 7896 contain an even number of digits.
Example 2:

Input: nums = [555,901,482,1771]
Output: 1
Explanation:
Only 1771 contains an even number of digits.


Constraints:

1 <= nums.length <= 500
1 <= nums[i] <= 105
*/

#include <iostream>

using namespace std;

/**
 * The function "print" takes a vector of integers as input and prints each
 * element followed by a space.
 *
 * @param vc The parameter "vc" is a reference to a vector of integers.
 */
void print(vector<int> &vc) {
  cout << "\t";
  for (auto i : vc) {
    cout << i << " ";
  }
  cout << endl;
}

class Solution {
public:
  /**
   * The function "hasEvenDigits" checks if a given number has an even number of
   * digits.
   *
   * @param num The parameter "num" is an integer representing the number for
   * which we want to check if it has even digits.
   *
   * @return a boolean value.
   */
  bool hasEvenDigits(int num) {
    int digitCount = 0;
    while (num) {
      digitCount++;
      num /= 10;
    }
    return (digitCount & 1) == 0;
  }

  /**
   * The function "findNumbers" counts the number of elements in a vector that
   * have an even number of digits.
   *
   * @param nums A reference to a vector of integers.
   *
   * @return the count of numbers in the given vector that have even number of
   * digits.
   */
  int findNumbers(vector<int> &nums) {
    int evenNumberDigitCount = 0;
    for (int i = 0; i < nums.size(); i++) {
      if (hasEvenDigits(nums[i]))
        evenNumberDigitCount++;
    }
    return evenNumberDigitCount;
  }
};

int main() {
  vector<int> nums = {12, 345, 2, 6, 7896};
  cout << "\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tNumbers in the array having even number of digits:\n";
  Solution solution;
  cout << "\t" << solution.findNumbers(nums) << endl;

  nums = {555, 901, 482, 1771};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tNumbers in the array having even number of digits:\n";
  cout << "\t" << solution.findNumbers(nums) << endl;
}