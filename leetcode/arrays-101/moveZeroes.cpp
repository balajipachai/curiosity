/*
Given an integer array nums, move all 0's to the end of it while maintaining the
relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]


Example 2:
Input: nums = [0]
Output: [0]


Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1


Follow up: Could you minimize the total number of operations done?
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
private:
  /* data */
public:
  /**
   * The function moves all the zeroes in a given vector to the end while
   * maintaining the relative order of the non-zero elements.
   *
   * @param nums The parameter "nums" is a reference to a vector of integers.
   */
  void moveZeroes(vector<int> &nums) {
    int zeroIndex = -1;
    int n = nums.size();
    for (int i = 0; i < n; i++) {
      if (nums[i] == 0 && zeroIndex == -1) {
        zeroIndex = i;
      } else if (nums[i] != 0 && zeroIndex >= 0) {
        nums[zeroIndex] = nums[i];
        nums[i] = 0;
        zeroIndex++;
      }
    }
  }
};

int main() {
  Solution solution;
  vector<int> nums = {0, 1, 0, 3, 12};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving zeroes to the end:\n";
  solution.moveZeroes(nums);
  print(nums, nums.size());

  nums = {0};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving zeroes to the end:\n";
  solution.moveZeroes(nums);
  print(nums, nums.size());

  nums = {1, 1, 2, 0};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving zeroes to the end:\n";
  solution.moveZeroes(nums);
  print(nums, nums.size());

  nums = {1, 1, 2};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving zeroes to the end:\n";
  solution.moveZeroes(nums);
  print(nums, nums.size());

  nums = {1, 0, 0, 1, 0, 0, 2};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after moving zeroes to the end:\n";
  solution.moveZeroes(nums);
  print(nums, nums.size());
  return 0;
}
