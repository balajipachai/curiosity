/*
Given a fixed-length integer array arr, duplicate each occurrence of zero,
shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written. Do
the above modifications to the input array in place and do not return anything.

Example 1:

Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to:
[1,0,0,2,3,0,0,4]

Example 2:

Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to:
[1,2,3]

Constraints:

1 <= arr.length <= 104
0 <= arr[i] <= 9
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
   * The function `duplicateZeros` takes a vector of integers as input and
   * duplicates any zeros in the vector, shifting the rest of the elements to
   * the right.
   *
   * @param arr The parameter `arr` is a reference to a vector of integers.
   */
  void duplicateZeros(vector<int> &arr) {
    int n = arr.size() - 1;
    int possibleDuplicates = 0;

    for (int left = 0; left <= n - possibleDuplicates; left++) {
      if (arr[left] == 0) {
        if (left == n - possibleDuplicates) {
          arr[n] = 0;
          n -= 1;
          break;
        }
        possibleDuplicates++;
      }
    }

    int last = n - possibleDuplicates;
    for (int i = last; i >= 0; i--) {
      if (arr[i] == 0) {
        arr[i + possibleDuplicates] = 0;
        possibleDuplicates--;
        arr[i + possibleDuplicates] = 0;
      } else {
        arr[i + possibleDuplicates] = arr[i];
      }
    }
  }
};

int main() {
  vector<int> nums = {1, 0, 2, 3, 0, 4, 5, 0};
  cout << "\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tPrinting the nums array after duplication zeros:\n";
  Solution solution;
  solution.duplicateZeros(nums);
  print(nums);

  nums = {1, 0, 1, 1, 0, 1};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums);
  cout << "\tPrinting the nums array after duplication zeros:\n";
  solution.duplicateZeros(nums);
  print(nums);
}
