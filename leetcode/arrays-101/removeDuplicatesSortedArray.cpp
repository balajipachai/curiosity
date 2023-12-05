/*
Given an integer array nums sorted in non-decreasing order, remove the
duplicates in-place such that each unique element appears only once. The
relative order of the elements should be kept the same. Then return the number
of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you
need to do the following things:

Change the array nums such that the first k elements of nums contain the unique
elements in the order they were present in nums initially. The remaining
elements of nums are not important as well as the size of nums. Return k. Custom
Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.



Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of
nums being 1 and 2 respectively. It does not matter what you leave beyond the
returned k (hence they are underscores). Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of
nums being 0, 1, 2, 3, and 4 respectively. It does not matter what you leave
beyond the returned k (hence they are underscores).


Constraints:

1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.
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
   * The function removes duplicates from a given vector of integers and returns
   * the number of unique elements.
   *
   * @param nums The parameter "nums" is a reference to a vector of integers.
   *
   * @return the number of unique elements in the vector after removing
   * duplicates.
   */
  int removeDuplicates(vector<int> &nums) {
    int unique = 0;
    int n = nums.size();

    if (n == 0 || n == 1) {
      return n;
    }

    for (int current = 0; current < n - 1; current++) {
      if (nums[current] != nums[current + 1]) {
        nums[unique++] = nums[current];
      }
    }
    // Copy the last element
    nums[unique++] = nums[n - 1];
    return unique;
  }
};

int main() {
  vector<int> nums = {1, 1, 2};
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after removing duplicates:\n";
  Solution solution;
  int uniqueElements = solution.removeDuplicates(nums);
  print(nums, uniqueElements);

  nums = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tPrinting the nums array after removing duplicates:\n";
  uniqueElements = solution.removeDuplicates(nums);
  print(nums, uniqueElements);
  return 0;
}