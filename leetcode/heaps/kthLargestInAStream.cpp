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
  cout << "\t";
  for (auto i : vc) {
    cout << i << " ";
  }
  cout << endl;
}

/* The KthLargest class maintains a min heap of the k largest elements from a
given set of numbers and provides a method to add a new number and return the
kth largest element. */
class KthLargest {
public:
  int k;
  vector<int> nums;
  /**
   * The function creates a min heap using the elements from the given vector
   * and pops elements from the heap until its length is equal to the given
   * value of k.
   *
   * @param k The parameter "k" represents the value of kth largest element that
   * we want to find.
   * @param nums A vector of integers containing the elements from which we want
   * to find the kth largest element.
   */
  KthLargest(int k, vector<int> &nums) {
    // In the constructor, create a min heap using the elements from nums. Then,
    // pop from the heap until heap.length == k.
    this->k = k;
    this->nums = nums;
    // Creation of min-heap
    make_heap(this->nums.begin(), this->nums.end(), greater<int>());
    while (this->nums.size() > this->k) {
      // Pop from the heap, and arrange the heap as a min-heap
      pop_heap(this->nums.begin(), this->nums.end(), greater<int>());
      this->nums.pop_back();
    }
  }

  /**
   * The add function adds a value to a vector, maintains the vector as a
   * min-heap, and removes the smallest elements if the vector size exceeds a
   * given threshold.
   *
   * @param val The parameter "val" is an integer value that you want to add to
   * the "nums" vector.
   *
   * @return the smallest element in the `nums` vector, which is the front
   * element of the vector after the heap has been arranged as a min-heap.
   */
  int add(int val) {
    this->nums.push_back(val);
    // Push in the heap, and arrange the heap as a min-heap
    push_heap(this->nums.begin(), this->nums.end(), greater<int>());
    while (this->nums.size() > this->k) {
      // Pop from the heap, and arrange the heap as a min-heap
      pop_heap(this->nums.begin(), this->nums.end(), greater<int>());
      this->nums.pop_back(); // removes the last element
    }
    return this->nums.front();
  }
};

/**
 * The code creates an object of class KthLargest and adds elements to it, then
 * prints the result of each addition.
 */
int main() {
  vector<int> nums = {4, 5, 8, 2};
  int k = 3;
  KthLargest *obj = new KthLargest(k, nums);
  cout << "\tInput stream:" << endl;
  print(nums);
  cout << "\tKthLargest in a stream after adding 3, 5, 10, 9 & 4 one by "
          "one:\n\t";
  cout << obj->add(3) << "\t";
  cout << obj->add(5) << "\t";
  cout << obj->add(10) << "\t";
  cout << obj->add(9) << "\t";
  cout << obj->add(4) << endl;
}