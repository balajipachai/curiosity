#include <iomanip>
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
   * The function `naiveMerge` takes two sorted arrays `nums1` and `nums2`,
   * merges them into `nums1`, and sorts the merged array using a naive
   * approach.
   *
   * @param nums1 A vector of integers representing the first array.
   * @param m The parameter `m` represents the size of the `nums1` vector, which
   * is the number of elements currently present in `nums1` before merging with
   * `nums2`.
   * @param nums2 The `nums2` parameter is a vector of integers representing the
   * second array that needs to be merged into `nums1`.
   * @param n The parameter `n` represents the size of the `nums2` vector, which
   * contains `n` elements.
   */
  void naiveMerge(vector<int> &nums1, int m, vector<int> &nums2, int n) {
    time_t start, end;
    time(&start);
    // unsync the I/O of C and C++.
    ios_base::sync_with_stdio(false);

    int k = m - n;
    // Naive approach merge and sort
    for (int i = 0; i < n; i++) {
      nums1[i + k] = nums2[i];
    }
    sort(nums1.begin(), nums1.end());

    // Recording end time.
    time(&end);

    // Calculating total time taken by the program.
    double time_taken = double(end - start);
    cout << "\tTime taken by naive approach : " << fixed << time_taken
         << setprecision(5);
    cout << " sec " << endl;
  }

  /**
   * The function efficientMerge merges two sorted arrays, nums1 and nums2, into
   * nums1 in an efficient manner.
   *
   * @param nums1 A vector of integers representing the first sorted array.
   * @param m The parameter `m` represents the size of the `nums1` vector, which
   * is the number of elements in `nums1` that are currently valid and should be
   * considered in the merge operation.
   * @param nums2 nums2 is a vector of integers that contains the elements to be
   * merged into nums1.
   * @param n The parameter `n` represents the size of the `nums2` vector, which
   * contains `n` elements.
   */
  void efficientMerge(vector<int> &nums1, int m, vector<int> &nums2, int n) {
    time_t start, end;
    time(&start);
    // unsync the I/O of C and C++.
    ios_base::sync_with_stdio(false);

    int p = 0;
    int p1 = 0;
    int p2 = 0;
    vector<int> nums1Copy(m);
    // Copy nums1, as we will be using 3 pointers
    nums1Copy = nums1;

    cout << "\tPrinting nums1Copy" << endl;
    print(nums1Copy);

    while (p < (m + n)) {
      if (p2 >= n || ((p1 < m) && (nums1Copy[p1] < nums2[p2]))) {
        nums1[p++] = nums1Copy[p1++];
      } else {
        nums1[p++] = nums2[p2++];
      }
    }

    // Recording end time.
    time(&end);

    // Calculating total time taken by the program.
    double time_taken = double(end - start);
    cout << "\tTime taken by efficient approach : " << fixed << time_taken
         << setprecision(5);
    cout << " sec " << endl;
  }
};

/**
 * The main function merges two arrays, nums1 and nums2, using both a naive
 * approach and an efficient approach.
 */
int main() {
  Solution solution;
  vector<int> nums1 = {10, 20, 30, 0, 0, 0};
  vector<int> nums2 = {4, 5, 6};
  cout << "\tPrinting nums1" << endl;
  print(nums1);
  cout << "\tPrinting nums2" << endl;
  print(nums2);
  solution.naiveMerge(nums1, nums1.size(), nums2, nums2.size());
  cout << "\tPrinting the array after merge" << endl;
  print(nums1);

  nums1 = {10, 20, 30, 0, 0, 0};
  solution.efficientMerge(nums1, 3, nums2, 3);
  cout << "\tPrinting the array after merge" << endl;
  print(nums1);
}