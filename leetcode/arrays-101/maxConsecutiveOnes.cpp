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
  /**
   * The function `findMaxConsecutiveOnes` finds the length of the longest
   * sequence of consecutive ones in a given vector, allowing for at most one
   * zero in the sequence.
   *
   * @param nums The parameter `nums` is a reference to a vector of integers.
   *
   * @return the length of the longest consecutive sequence of 1s in the given
   * vector of integers.
   */
  int bruteForce(vector<int> &nums) {
    // Brute Force Approach
    int longestSequence = 0;
    int n = nums.size();
    for (int left = 0; left < n; left++) {
      int numZeroes = 0;
      for (int right = left; right < n; right++) {
        if (nums[right] == 0) {
          numZeroes += 1;
        }

        if (numZeroes <= 1) {
          longestSequence = max(longestSequence, right - left + 1);
        }
      }
    }
    return longestSequence;
  }

  /**
   * The slidingWindow function finds the length of the longest subarray in a
   * given array that contains at most two zeroes.
   *
   * @param nums The parameter "nums" is a reference to a vector of integers.
   *
   * @return the length of the longest subarray in the given vector `nums` that
   * contains at most two zeroes.
   */
  int slidingWindow(vector<int> &nums) {
    int longestSequence = 0;
    int n = nums.size();
    int right = 0;
    int left = 0;
    int numZeroes = 0;

    while (right < n) {
      if (nums[right] == 0) {
        numZeroes += 1;
      }

      // invalid window, then contract the window
      while (numZeroes == 2) {
        if (nums[left] == 0) {
          numZeroes -= 1;
        }
        left++;
      }
      longestSequence = max(longestSequence, right - left + 1);
      // expand window
      right++;
    }
    return longestSequence;
  }
};

int main() {
  Solution solution;
  vector<int> nums = {1, 0, 1, 1, 0, 1};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tMax consecutive ones in the array using Brute Force approach "
          "are:\n";

  int maxOnes = solution.bruteForce(nums);
  cout << "\t" << maxOnes << endl;

  cout << "\tMax consecutive ones in the array using Sliding Window approach "
          "are:\n";

  maxOnes = solution.slidingWindow(nums);
  cout << "\t" << maxOnes << endl;
}