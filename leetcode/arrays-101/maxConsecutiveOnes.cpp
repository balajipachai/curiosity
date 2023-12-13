/*
Given a binary array nums, return the maximum number of consecutive 1's in the
array if you can flip at most one 0.



Example 1:

Input: nums = [1,0,1,1,0]
Output: 4
Explanation:
- If we flip the first zero, nums becomes [1,1,1,1,0] and we have 4 consecutive
ones.
- If we flip the second zero, nums becomes [1,0,1,1,1] and we have 3 consecutive
ones. The max number of consecutive ones is 4. Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 4
Explanation:
- If we flip the first zero, nums becomes [1,1,1,1,0,1] and we have 4
consecutive ones.
- If we flip the second zero, nums becomes [1,0,1,1,1,1] and we have 4
consecutive ones. The max number of consecutive ones is 4.


Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.


Follow up: What if the input numbers come in one by one as an infinite stream?
In other words, you can't store all numbers coming from the stream as it's too
large to hold in memory. Could you solve it efficiently?
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
  int findMaxConsecutiveOnes(vector<int> &nums) {
    int n = nums.size();
    int maxOnes = 0;
    int count = 0;
    int lastZeroIndex = -1;

    for (int i = 0; i < n; i++) {
      // The count should increase whenever there's a 1 in the array
      // As and when a zero is encountered, flip that too a 1 and continue
      // increasing the one's count Whenever a second 0 is encountered then
      // store the max consecutive ones count in a variable called as maxOnes
      // set the onesCount to the difference of current zero index -
      // previousZeroIndex + 1 (+1 since the array indices are zero based) Check
      // if previousMaxOnes < currentMaxOnes, and set maxOnes count accordingly
      // Return maxOnes count
      // DONE
      if (nums[i]) {
        cout << "if" << endl;
        count++;
      } else if (!nums[i] && lastZeroIndex == -1) {
        cout << "elif" << endl;
        count++;
        lastZeroIndex = i;
      } else {
        cout << "else" << endl;
        // second zero is encountered
        maxOnes = count;
        count = i - (lastZeroIndex + 1);
        lastZeroIndex = i;
      }

      //   return maxOnes;

      //   if (!nums[i] && lastZeroIndex == -1) {
      //     cout << "\tInside first if\t";
      //     lastZeroIndex = i;
      //     count++;
      //     cout << lastZeroIndex << endl;
      //   } else if (!nums[i] && lastZeroIndex >= 0) {
      //     if (count > maxOnes) {
      //       cout << "\tInside second if\t";
      //       maxOnes = count;
      //       cout << maxOnes << endl;
      //     }
      //     count = i - (lastZeroIndex + 1);
      //     lastZeroIndex = i;
      //   } else {
      //     count++;
      //     cout << "\tno ifs executed\t" << count << endl;
      //     if (!maxOnes && count == n) {
      //       maxOnes = count;
      //     }
      //   }
    }
    cout << "\tMax ones " << maxOnes << endl;
    return maxOnes;
  }
};

int main() {
  Solution solution;
  vector<int> nums = {1, 1, 0, 1};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tMax consecutive ones in the array are:\n";

  int maxOnes = solution.findMaxConsecutiveOnes(nums);
  cout << "\t" << maxOnes << endl;
}