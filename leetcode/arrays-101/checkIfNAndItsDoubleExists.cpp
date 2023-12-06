/*
Check If N and Its Double Exist
Given an array arr of integers, check if there exist two indices i and j such
that :

i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]


Example 1:

Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]
Example 2:

Input: arr = [3,1,7,11]
Output: false
Explanation: There is no i and j that satisfy the conditions.


Constraints:

2 <= arr.length <= 500
-103 <= arr[i] <= 103
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
   * The function checks if there exists an element in the given vector that is
   * equal to twice of another element in the vector.
   *
   * @param arr arr is a vector of integers.
   *
   * @return a boolean value. It returns true if there exists an element in the
   * input vector `arr` such that it is equal to twice of another element in the
   * vector (excluding itself). Otherwise, it returns false.
   */
  bool checkIfExist(vector<int> &arr) {
    int n = arr.size();

    if (n == 0 || n == 1)
      return false; // Empty Array
    int current = 0;
    while (current < n) {
      for (int i = 0; i < n; i++) {
        if (arr[current] == (2 * arr[i]) && (i != current)) {
          cout << "\tN & its Double Exists at i, j = " << current << ", " << i
               << endl;
          return true;
        }
      }
      current++;
    }
    cout << "\tN & its Double Does Not Exist" << endl;
    return false;
  }
};

int main() {
  vector<int> nums = {10, 2, 5, 3};
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  Solution solution;
  solution.checkIfExist(nums);

  nums = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  solution.checkIfExist(nums);

  nums = {3, 1, 7, 11};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  solution.checkIfExist(nums);
  return 0;
}