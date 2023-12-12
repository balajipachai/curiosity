/*
Given an integer array nums, move all the even integers at the beginning of the
array followed by all the odd integers.

Return any array that satisfies this condition.



Example 1:

Input: nums = [3,1,2,4]
Output: [2,4,3,1]
Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be
accepted. Example 2:

Input: nums = [0]
Output: [0]


Constraints:

1 <= nums.length <= 5000
0 <= nums[i] <= 5000
*/
#include <iostream>

using namespace std;

/**
 * The function "print" takes a vector of integers as input and prints each
 * element followed by a space.
 *
 * @param vc The parameter "vc" is a reference to a vector of integers.
 */
void print(vector<int> &vc, int n) {
  cout << "\t";
  for (int i = 0; i < n; i++) {
    cout << vc[i] << " ";
  }
  cout << endl;
}

class Solution {
public:
  /**
   * The function sorts an array of integers by moving all odd numbers to the
   * end of the array while keeping the relative order of even numbers
   * unchanged.
   *
   * @param nums The parameter "nums" is a reference to a vector of integers.
   */
  vector<int> sortArrayByParity(vector<int> &nums) {
    int oddIndex = -1;
    int n = nums.size();
    int temp = 0;
    for (int i = 0; i < n; i++) {
      if ((nums[i] % 2) && oddIndex == -1) {
        oddIndex = i;
      } else if (!(nums[i] % 2) && oddIndex >= 0) {
        swap(nums[oddIndex], nums[i]);
        oddIndex++;
      }
    }
    return nums;
  }
};

int main() {
  Solution solution;
  vector<int> nums = {0, 1, 0, 3, 12};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving odd numbers to the end:\n";
  solution.sortArrayByParity(nums);
  print(nums, nums.size());

  nums = {0};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving odd numbers to the end:\n";
  solution.sortArrayByParity(nums);
  print(nums, nums.size());

  nums = {1, 1, 2, 0};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving odd numbers to the end:\n";
  solution.sortArrayByParity(nums);
  print(nums, nums.size());

  nums = {1, 1, 2};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving odd numbers to the end:\n";
  solution.sortArrayByParity(nums);
  print(nums, nums.size());

  nums = {1, 0, 0, 1, 0, 0, 2};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving odd numbers to the end:\n";
  solution.sortArrayByParity(nums);
  print(nums, nums.size());
  return 0;
}
