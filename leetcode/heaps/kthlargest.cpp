#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

/**
 * The function "print" takes a vector of integers as input and prints each
 * element followed by a space.
 *
 * @param vc The parameter "vc" is a reference to a vector of integers.
 */
void print(vector<int> &vc) {
  for (auto i : vc) {
    cout << i << " ";
  }
  cout << endl;
}

class Solution {
public:
  int findKthLargest(vector<int> &nums, int k) {
    // Create heap using the vector by using the C++ STL's make_heap method
    // By default this creates a max-heap
    make_heap(nums.begin(), nums.end());
    // For max-heap, the element at index 0 is the largest
    // Thus, to obtain the Kth Largest Element
    // Perform the below operations K times
    // delete the topmost element
    // heapify the remaining element
    // Once the loop exits for K times then print the heap[0] this will be the
    // kth largest element
    int i = 0;
    int secondLargest;
    while (i < k) {
      // Moves the largest element to the end of the array
      // pop_heap ensures that heapify property is maintained while removing the
      // elements
      pop_heap(nums.begin(), nums.end());
      secondLargest = nums.back();
      nums.pop_back();
      i++;
    }
    return secondLargest;
  }
};

int main() {
  vector<int> nums = {3, 2, 1, 5, 6, 4};
  int k = 2;
  cout << "\tFind the 2nd largest element from: ";
  print(nums);
  Solution solution;
  int secondLargest = solution.findKthLargest(nums, k);
  cout << "\t2nd largest element is: " << secondLargest << endl;
}