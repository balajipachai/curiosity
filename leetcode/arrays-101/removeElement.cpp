/**
 * Given an integer array nums and an integer val, remove all occurrences of val
in nums in-place. The order of the elements may be changed. Then return the
number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get
accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the
elements which are not equal to val. The remaining elements of nums are not
important as well as the size of nums. Return k.

Example 1:

Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of
nums being 2. It does not matter what you leave beyond the returned k (hence
they are underscores). Example 2:

Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of
nums containing 0, 0, 1, 3, and 4. Note that the five elements can be returned
in any order. It does not matter what you leave beyond the returned k (hence
they are underscores).


Constraints:

0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100
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
   * The function removes all occurrences of a given value from a vector and
   * returns the new size of the vector.
   *
   * @param nums A reference to a vector of integers, which represents the array
   * from which we want to remove elements.
   * @param val The value that needs to be removed from the vector.
   *
   * @return the new size of the vector after removing all occurrences of the
   * given value.
   */
  int removeElement(vector<int> &nums, int val) {
    int i = 0;
    int n = nums.size();
    // Efficient Approach
    while (i < n) {
      if (nums[i] == val) {
        // Copy the last element at i^th position and remove the last element
        // reducing the size of the array this avoids unnecessary copy
        // operations
        nums[i] = nums[n - 1];
        nums.pop_back();
        n--;
      } else {
        i++;
      }
    }
    cout << "\tNew size is \t" << n << endl;
    return n;
  }
};

int main() {
  vector<int> nums = {1, 0, 2, 3, 0, 4, 5, 0};
  int val = 0;
  cout << "\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tPrinting the nums array after removing the val = 0:\n";
  Solution solution;
  solution.removeElement(nums, val);
  print(nums);

  nums = {3, 2, 2, 3};
  val = 3;
  cout << "\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tPrinting the nums array after removing the val = 3:\n";
  solution.removeElement(nums, val);
  print(nums);

  nums = {0, 1, 2, 2, 3, 0, 4, 2};
  val = 2;
  cout << "\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tPrinting the nums array after removing the val = 2:\n";
  solution.removeElement(nums, val);
  print(nums);
}