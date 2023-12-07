/*
Given an array arr, replace every element in that array with the greatest
element among the elements to its right, and replace the last element with -1.

After doing so, return the array.



Example 1:

Input: arr = [17,18,5,4,6,1]
Output: [18,6,6,6,1,-1]
Explanation:
- index 0 --> the greatest element to the right of index 0 is index 1 (18).
- index 1 --> the greatest element to the right of index 1 is index 4 (6).
- index 2 --> the greatest element to the right of index 2 is index 4 (6).
- index 3 --> the greatest element to the right of index 3 is index 4 (6).
- index 4 --> the greatest element to the right of index 4 is index 5 (1).
- index 5 --> there are no elements to the right of index 5, so we put -1.
Example 2:

Input: arr = [400]
Output: [-1]
Explanation: There are no elements to the right of index 0.


Constraints:

1 <= arr.length <= 104
1 <= arr[i] <= 105
*/

#include <iomanip>
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
   * The function replaces each element in a vector with the maximum element to
   * its right, and sets the last element to -1.
   *
   * @param arr A reference to a vector of integers.
   *
   * @return a vector of integers.
   */
  vector<int> replaceElementsNaiveApproach(vector<int> &arr) {
    time_t start, end;
    time(&start);
    // unsync the I/O of C and C++.
    ios_base::sync_with_stdio(false);
    int current = 0;
    int n = arr.size();
    int max = 0;

    if (n == 1) {
      arr[current] = -1;
      return arr;
    }

    while (current < n) {
      for (int i = current + 1; i < n; i++) {
        if (arr[i] > max) {
          max = arr[i];
        }
      }
      arr[current] = max;
      max = 0;
      current++;
    }
    // last element
    arr[n - 1] = -1;
    // Recording end time.
    time(&end);

    // Calculating total time taken by the program.
    double time_taken = double(end - start);
    cout << "\tTime taken by naive approach : " << fixed << time_taken
         << setprecision(5);
    cout << " sec " << endl;
    return arr;
  }

  /**
   * The function replaces each element in a vector with the maximum element to
   * its right, and returns the modified vector.
   *
   * @param arr A reference to a vector of integers.
   *
   * @return a vector of integers.
   */
  vector<int> replaceElementsEfficientApproach(vector<int> &arr) {
    time_t start, end;
    time(&start);
    // unsync the I/O of C and C++.
    ios_base::sync_with_stdio(false);
    int n = arr.size();
    int temp = 0;
    int maxFromRight = -1;

    for (int i = n - 1; i >= 0; i--) {
      temp = maxFromRight;
      if (arr[i] > maxFromRight) {
        maxFromRight = arr[i];
      }
      arr[i] = temp;
    }
    // Recording end time.
    time(&end);

    // Calculating total time taken by the program.
    double time_taken = double(end - start);
    cout << "\tTime taken by efficient approach : " << fixed << time_taken
         << setprecision(5);
    cout << " sec " << endl;
    return arr;
  }
};

int main() {
  vector<int> nums = {
      7, 2, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3};
  cout << "\t*****************************Test1*****************************"
       << endl;
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  Solution solution;
  solution.replaceElementsNaiveApproach(nums);
  cout << "\tPrinting the nums array AFTER replace:" << endl;
  print(nums, nums.size());

  nums = {2};
  cout << "\t*****************************Test2*****************************"
       << endl;
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  solution.replaceElementsNaiveApproach(nums);
  cout << "\tPrinting the nums array AFTER replace:" << endl;
  print(nums, nums.size());

  nums = {17, 18, 5, 4, 6, 1};
  cout << "\t*****************************Test3*****************************"
       << endl;
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  solution.replaceElementsNaiveApproach(nums);
  cout << "\tPrinting the nums array AFTER replace:" << endl;
  print(nums, nums.size());

  nums = {
      7, 2, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0,
      6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5,
      6, 7, 2, 9, 3, 4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3,
      4, 1, 0, 8, 5, 2, 7, 4, 9, 0, 6, 3, 1, 8, 5, 6, 7, 2, 9, 3, 4, 1, 0, 8, 5,
      2, 7, 4, 9, 0, 6, 3};
  cout << "\t*****************************Test4*****************************"
       << endl;
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  solution.replaceElementsEfficientApproach(nums);
  cout << "\tPrinting the nums array AFTER replace:" << endl;
  print(nums, nums.size());

  nums = {2};
  cout << "\t*****************************Test5*****************************"
       << endl;
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  solution.replaceElementsEfficientApproach(nums);
  cout << "\tPrinting the nums array AFTER replace:" << endl;
  print(nums, nums.size());

  nums = {17, 18, 5, 4, 6, 1};
  cout << "\t*****************************Test6*****************************"
       << endl;
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  solution.replaceElementsEfficientApproach(nums);
  cout << "\tPrinting the nums array AFTER replace:" << endl;
  print(nums, nums.size());
}