/*
Given an integer array nums, return the third distinct maximum number in this
array. If the third maximum does not exist, return the maximum number.



Example 1:

Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.
Example 2:

Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned
instead. Example 3:

Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have
the same value). The third distinct maximum is 1.


Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1


Follow up: Can you find an O(n) solution?
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
  int bruteForce(vector<int> &nums) {
    int n = nums.size();
    if (n == 1) {
      return nums[0];
    }
    // Sort the vector in descending order
    sort(nums.begin(), nums.end(), greater<int>());
    // Remove Duplicates and have only distinct elements
    auto ip = unique(nums.begin(), nums.end());
    nums.resize(distance(nums.begin(), ip));
    n = nums.size();
    if (n >= 3) {
      return nums[2];
    }
    return nums[1];
  }
};

int main() {
  Solution solution;
  vector<int> nums = {2, 2, 3, 1};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tThird distinct max is:\n";
  int max = solution.bruteForce(nums);
  cout << "\tPrinting the nums array after sort, unique & resize:" << endl;
  print(nums, nums.size());
  cout << "\t" << max << endl;

  nums = {3, 2, 1};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tThird distinct max is:\n";
  max = solution.bruteForce(nums);
  cout << "\tPrinting the nums array after sort, unique & resize:" << endl;
  print(nums, nums.size());
  cout << "\t" << max << endl;

  nums = {1, 2};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tThird distinct max is:\n";
  max = solution.bruteForce(nums);
  cout << "\tPrinting the nums array after sort, unique & resize:" << endl;
  print(nums, nums.size());
  cout << "\t" << max << endl;

  nums = {1, 2, 3, 3, 3, 10, 1, 2, 3, 7, 7, 8};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tThird distinct max is:\n";
  max = solution.bruteForce(nums);
  cout << "\tPrinting the nums array after sort, unique & resize:" << endl;
  print(nums, nums.size());
  cout << "\t" << max << endl;
}