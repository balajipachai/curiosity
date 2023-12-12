/*
A school is trying to take an annual photo of all the students. The students are
asked to stand in a single file line in non-decreasing order by height. Let this
ordering be represented by the integer array expected where expected[i] is the
expected height of the ith student in line.

You are given an integer array heights representing the current order that the
students are standing in. Each heights[i] is the height of the ith student in
line (0-indexed).

Return the number of indices where heights[i] != expected[i].

Example 1:

Input: heights = [1,1,4,2,1,3]
Output: 3
Explanation:
heights:  [1,1,4,2,1,3]
expected: [1,1,1,2,3,4]
Indices 2, 4, and 5 do not match.
Example 2:

Input: heights = [5,1,2,3,4]
Output: 5
Explanation:
heights:  [5,1,2,3,4]
expected: [1,2,3,4,5]
All indices do not match.
Example 3:

Input: heights = [1,2,3,4,5]
Output: 0
Explanation:
heights:  [1,2,3,4,5]
expected: [1,2,3,4,5]
All indices match.


Constraints:

1 <= heights.length <= 100
1 <= heights[i] <= 100
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
   * The function heightChecker takes in a vector of heights and returns the
   * number of students who are not in the correct order when standing in
   * non-decreasing order of their heights.
   *
   * @param heights The parameter "heights" is a reference to a vector of
   * integers.
   *
   * @return the count of students whose heights are not in the correct order.
   */
  int heightChecker(vector<int> &heights) {
    int n = heights.size();
    int count = 0;

    vector<int> sortedHeights(n);

    for (int i = 0; i < n; i++) {
      sortedHeights[i] = heights[i];
    }

    // Sort the heights
    sort(sortedHeights.begin(), sortedHeights.end());

    for (int i = 0; i < n; i++) {
      if (heights[i] != sortedHeights[i]) {
        count++;
      }
    }

    return count;
  }
};

int main() {
  Solution solution;
  vector<int> nums = {1, 1, 4, 2, 1, 3};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tNumber of indices where height != expected:\n";

  int notExpectedHeights = solution.heightChecker(nums);
  cout << "\t" << notExpectedHeights << endl;

  nums = {5, 1, 2, 3, 4};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tNumber of indices where height != expected:\n";

  notExpectedHeights = solution.heightChecker(nums);
  cout << "\t" << notExpectedHeights << endl;

  nums = {1, 2, 3, 4, 5};

  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  cout << "\tNumber of indices where height != expected:\n";

  notExpectedHeights = solution.heightChecker(nums);
  cout << "\t" << notExpectedHeights << endl;
  return 0;
}