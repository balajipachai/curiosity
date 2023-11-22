/**
 * Given a binary array nums, return the maximum number of consecutive 1's in
the array.
 *
Example 1:

Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
The maximum number of consecutive 1s is 3.

Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 2

Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
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
   * The function finds the maximum number of consecutive ones in a given vector
   * of integers.
   *
   * @param nums The parameter "nums" is a reference to a vector of integers.
   *
   * @return the maximum number of consecutive ones in the given vector of
   * integers.
   */
  int findMaxConsecutiveOnes(vector<int> &nums) {
    int maxOnes = 0;
    int count = 0;
    for (int i = 0; i < nums.size(); i++) {
      if (nums[i] == 1) {
        count += 1;
      } else {
        maxOnes = max(maxOnes, count);
        count = 0;
      }
    }
    return max(maxOnes, count);
  }
};

int main() {
  vector<int> nums = {1, 1, 0, 1, 1, 1};
  cout << "\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tMax consecutive ones in the given array:\n";
  Solution solution;
  cout << "\t" << solution.findMaxConsecutiveOnes(nums) << endl;

  nums = {1, 0, 1, 1, 0, 1};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tMax consecutive ones in the given array:\n";
  cout << "\t" << solution.findMaxConsecutiveOnes(nums) << endl;
}