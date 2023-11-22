/**
 * Given an integer array nums sorted in non-decreasing order, return an array
of the squares of each number sorted in non-decreasing order.

Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]


Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.


Follow up: Squaring each element and sorting the new array is very trivial,
could you find an O(n) solution using a different approach?
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
   * The function takes a vector of integers, squares each element, and returns
   * a new vector with the squared elements in non-decreasing order.
   *
   * @param nums A reference to a vector of integers, named "nums".
   *
   * @return The function `sortedSquares` returns a vector of integers, which
   * contains the squares of the elements in the input vector `nums`, sorted in
   * non-decreasing order.
   */
  vector<int> sortedSquares(vector<int> &nums) {
    int n = nums.size();
    vector<int> answer(n);

    // We will use the approach of 2 pointers to solve this problem in O(n)
    int left = 0;      // First pointer starts from the beginning
    int right = n - 1; // Second pointer starts from the end
    int squaresOf;
    // Iterate over the vector
    for (int i = n - 1; i >= 0; i--) {
      if (abs(nums[left]) < abs(nums[right])) {
        squaresOf = nums[right];
        right--;
      } else {
        squaresOf = nums[left];
        left++;
      }
      answer[i] = squaresOf * squaresOf;
    }
    return answer;
  }
};

/**
 * The main function prints and calculates the squares of numbers in an array
 * and sorts them in non-decreasing order.
 */
int main() {
  vector<int> nums = {-4, -1, 0, 3, 10};
  cout << "\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tAn array of the squares of each number sorted in non-decreasing "
          "order:\n";
  Solution solution;
  vector<int> answer = solution.sortedSquares(nums);
  print(answer);

  nums = {-7, -3, 2, 3, 11};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tAn array of the squares of each number sorted in non-decreasing "
          "order:\n";
  answer = solution.sortedSquares(nums);
  print(answer);
}