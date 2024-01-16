/*
Given an array nums of n integers where nums[i] is in the range [1, n], return
an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Example 2:

Input: nums = [1,1]
Output: [2]


Constraints:

n == nums.length
1 <= n <= 105
1 <= nums[i] <= n


Follow up: Could you do it without extra space and in O(n) runtime? You may
assume the returned list does not count as extra space.
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
   * The bruteForce function takes in a vector of integers and returns a vector
   * of integers that are missing from the input vector.
   *
   * @param nums The parameter "nums" is a reference to a vector of integers.
   *
   * @return The function `bruteForce` returns a vector of integers.
   */
  vector<int> bruteForce(vector<int> &nums) {
    int start = 1;
    int end = nums.size();
    vector<int> result;
    bool found;
    while (start <= end) {
      for (int i = 0; i < end; i++) {
        if (nums[i] == start) {
          found = true;
        }
      }
      if (!found) {
        result.push_back(start);
      }
      start += 1;
      found = false;
    }
    return result;
  }
  /**
   * The function `findDisappearedNumbers` takes in a vector of integers and
   * returns a vector of integers containing the numbers that are missing from
   * the input vector.
   *
   * @param nums The parameter `nums` is a reference to a vector of integers.
   *
   * @return a vector of integers, which contains the numbers that are missing
   * from the input vector.
   */
  vector<int> findDisappearedNumbers(vector<int> &nums) {
    int n = nums.size();
    /* The code snippet is used to mark the numbers that are present in the
       input vector `nums`. */
    for (int i = 0; i < n; i++) {
      int newIndex = abs(nums[i]) - 1;
      if (nums[newIndex] > 0) {
        nums[newIndex] *= -1;
      }
    }

    vector<int> result;

    /* The code snippet is used to find the numbers that are missing from the
       input vector `nums`. */
    for (int i = 1; i <= n; i++) {
      if (nums[i - 1] > 0) {
        result.push_back(i);
      }
    }

    return result;
  }
};

int main() {
  Solution solution;
  vector<int> nums = {2, 2, 3, 1};
  cout << "\tBRUTE FORCE" << endl;
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  vector<int> result = solution.bruteForce(nums);
  printf("\tPrinting all the integers in the range [1, %lu] that do not appear "
         "in nums:\n",
         nums.size());
  print(result, result.size());
  cout << "\t************************************************\n";

  nums = {4, 3, 2, 7, 8, 2, 3, 1};
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  result = solution.bruteForce(nums);
  printf("\tPrinting all the integers in the range [1, %lu] that do not appear "
         "in nums:\n",
         nums.size());
  print(result, result.size());
  cout << "\t************************************************\n";

  nums = {1, 1};
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  result = solution.bruteForce(nums);
  printf("\tPrinting all the integers in the range [1, %lu] that do not appear "
         "in nums:\n",
         nums.size());
  print(result, result.size());
  cout << "\t************************************************\n";

  cout << "\tEFFICIENT APPROACH" << endl;
  nums = {2, 2, 3, 1};
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  result = solution.findDisappearedNumbers(nums);
  printf("\tPrinting all the integers in the range [1, %lu] that do not appear "
         "in nums:\n",
         nums.size());
  print(result, result.size());
  cout << "\t************************************************\n";

  nums = {4, 3, 2, 7, 8, 2, 3, 1};
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  result = solution.findDisappearedNumbers(nums);
  printf("\tPrinting all the integers in the range [1, %lu] that do not appear "
         "in nums:\n",
         nums.size());
  print(result, result.size());
  cout << "\t************************************************\n";

  nums = {1, 1};
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  result = solution.findDisappearedNumbers(nums);
  printf("\tPrinting all the integers in the range [1, %lu] that do not appear "
         "in nums:\n",
         nums.size());
  print(result, result.size());
  cout << "\t************************************************\n";

  return 0;
}