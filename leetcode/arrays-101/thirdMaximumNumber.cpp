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
  /**
   * The function bruteForce takes a vector of integers, sorts it in descending
   * order, removes duplicates, and returns the third largest distinct element
   * if it exists, otherwise it returns the second largest distinct element.
   *
   * @param nums The parameter "nums" is a reference to a vector of integers.
   *
   * @return The function `bruteForce` returns the third largest distinct
   * element from the input vector `nums`. If there are less than 3 distinct
   * elements, it returns the second largest distinct element.
   */
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

  /**
   * The function "thirdMax" returns the third largest element in a given vector
   * of integers, ignoring duplicates.
   *
   * @param nums A reference to a vector of integers, named "nums".
   *
   * @return the third maximum element from the given vector of integers.
   */
  int thirdMax(vector<int> &nums) {
    long long firstMax = numeric_limits<long long>::min();
    long long secondMax = numeric_limits<long long>::min();
    long long thirdMax = numeric_limits<long long>::min();

    for (int i = 0; i < nums.size(); i++) {
      // If the num is already stored in either firstMax, secondMax or thirdMax
      // then skip it
      if (firstMax == nums[i] || secondMax == nums[i] || thirdMax == nums[i]) {
        continue;
      }

      // if nums[i] is greather than firstMax then update both secondMax and
      // thirdMax
      if (firstMax <= nums[i]) {
        thirdMax = secondMax;
        secondMax = firstMax;
        firstMax = nums[i];
      } else if (secondMax <= nums[i]) {
        // when nums[i] > secondMax then update secondMax and thirdmax
        thirdMax = secondMax;
        secondMax = nums[i];
      } else if (thirdMax <= nums[i]) {
        // when nums[i] > thirdMax then update thirdMax
        thirdMax = nums[i];
      }
    }

    // When thirdMax is not at all update in that case return firstMax
    if (thirdMax == numeric_limits<long long>::min()) {
      return firstMax;
    }

    return thirdMax;
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

  cout << "\t*************Efficient Approach**********" << endl;
  nums = {1, 2, 3, 3, 3, 10, 1, 2, 3, 7, 7, 8};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tThird distinct max is:\n";
  max = solution.thirdMax(nums);
  cout << "\t" << max << endl;

  nums = {1, 2};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tThird distinct max is:\n";
  max = solution.thirdMax(nums);
  cout << "\t" << max << endl;

  nums = {3, 2, 1};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tThird distinct max is:\n";
  max = solution.thirdMax(nums);
  cout << "\t" << max << endl;

  nums = {1, 2, -2147483648};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tThird distinct max is:\n";
  max = solution.thirdMax(nums);
  cout << "\t" << max << endl;
}