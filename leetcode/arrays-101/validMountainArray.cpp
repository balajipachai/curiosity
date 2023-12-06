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
   * The function checks if an array is a valid mountain array, meaning it
   * starts with strictly increasing elements, reaches a peak, and then has
   * strictly decreasing elements.
   *
   * @param arr The parameter `arr` is a vector of integers.
   *
   * @return a boolean value. It returns true if the input array is a valid
   * mountain array, and false otherwise.
   */
  bool validMountainArray(vector<int> &arr) {
    int n = arr.size();
    int current = 0;

    /*
    If we walk along the mountain from left to right, we have to move strictly
    up, then strictly down. Let's walk up from left to right until we can't:
    that has to be the peak. We should ensure the peak is not the first or last
    element. Then, we walk down. If we reach the end, the array is valid,
    otherwise its not.
    */
    // Walk up
    while (current < n && arr[current] < arr[current + 1]) {
      current++;
    }
    // When current = n-1 this implies array contains elements which are
    // arranged in ascending order When current = 0 this implies array contains
    // elements which are arranged in descending order
    if (current == n - 1 || current == 0) {
      return false;
    }

    // Walk up
    while (current + 1 < n && arr[current] > arr[current + 1]) {
      current++;
    }

    return current == n - 1;
  }
};

int main() {
  vector<int> nums = {2, 1};
  cout << "\t*****************************Test1*****************************"
       << endl;
  cout << "\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  Solution solution;
  bool result = solution.validMountainArray(nums);

  if (result) {
    cout << "\tValid mountain array" << endl;
  } else {
    cout << "\tNot a mountain array" << endl;
  }

  cout << "\t*****************************Test2*****************************"
       << endl;
  nums = {3, 5, 5};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  result = solution.validMountainArray(nums);
  if (result) {
    cout << "\tValid mountain array" << endl;
  } else {
    cout << "\tNot a mountain array" << endl;
  }

  cout << "\t*****************************Test3*****************************"
       << endl;
  nums = {0, 3, 2, 1};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  result = solution.validMountainArray(nums);
  if (result) {
    cout << "\tValid mountain array" << endl;
  } else {
    cout << "\tNot a mountain array" << endl;
  }

  cout << "\t*****************************Test4*****************************"
       << endl;
  nums = {2};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  result = solution.validMountainArray(nums);
  if (result) {
    cout << "\tValid mountain array" << endl;
  } else {
    cout << "\tNot a mountain array" << endl;
  }

  cout << "\t*****************************Test5*****************************"
       << endl;
  nums = {1, 7, 9, 5, 4, 1, 2};
  cout << "\n\tPrinting the nums array:" << endl;
  print(nums, nums.size());
  result = solution.validMountainArray(nums);
  if (result) {
    cout << "\tValid mountain array" << endl;
  } else {
    cout << "\tNot a mountain array" << endl;
  }
  return 0;
}