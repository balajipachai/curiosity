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

struct greaters {
  bool operator()(const int &a, const int &b) { return a > b; }
};

class KthLargest {
public:
  int _k;
  vector<int> _nums;
  KthLargest(int k, vector<int> &nums) {}

  int add(int val) {
    // Uses Sorting
    _nums.push_back(val);
    sort(_nums.begin(), _nums.end());
    return _nums[_nums.size() - _k];
  }

  void KthLargestUsingHeaps(int k, vector<int> &nums) {
    _k = k;
    for (int i = 0; i < nums.size(); i++)
      _nums.push_back(nums[i]);
    // Create a min-heap using the elements from nums
    make_heap(_nums.begin(), _nums.end(), greaters());
    cout << "isHeap 1\t" << is_heap(_nums.begin(), _nums.end()) << endl;
    cout << "Printing _nums\t";
    print(_nums);
    // cout << "1: \n";
    // print(nums);
    // pop from the heap until heap.length == k.
    while (_nums.size() != _k) {
      //   cout << "2: \n";
      //   print(nums);
      pop_heap(_nums.begin(), _nums.end());
      //   cout << "3: \n";
      //   print(nums);
      _nums.pop_back(); // removes element from the vector
      //   cout << "4: \n";
      //   print(nums);
      cout << "isHeap 2\t" << is_heap(_nums.begin(), _nums.end()) << endl;
    }
  }

  int addUsingHeaps(int val) {
    _nums.push_back(val);
    push_heap(_nums.begin(), _nums.end());
    cout << "\nnums size: ***********\t" << _nums.size()
         << "\t*****************\n";
    cout << "isHeap 3\t" << is_heap(_nums.begin(), _nums.end()) << endl;
    while (_nums.size() > _k) {
      //   cout << "5: \n";
      //   print(_nums);
      pop_heap(_nums.begin(), _nums.end());
      //   cout << "6: \n";
      //   print(_nums);
      _nums.pop_back();
      //   cout << "7: \n";
      //   print(_nums);
      cout << "isHeap 4\t" << is_heap(_nums.begin(), _nums.end()) << endl;
    }
    cout << "Printing _nums\t";
    print(_nums);
    return _nums.front();
  }
};

int main() {
  vector<int> nums = {4, 5, 8, 2};
  int k = 3;
  KthLargest *obj = new KthLargest(k, nums);
  //   cout << "\tKthLargest after adding 3 in the stream\t" << obj->add(3) <<
  //   endl;
  //   ;
  //   cout << "\tKthLargest after adding 5 in the stream\t" << obj->add(5) <<
  //   endl;
  //   ;
  //   cout << "\tKthLargest after adding 10 in the stream\t" << obj->add(10)
  //        << endl;
  //   ;
  //   cout << "\tKthLargest after adding 9 in the stream\t" << obj->add(9) <<
  //   endl;
  //   ;
  //   cout << "\tKthLargest after adding 4 in the stream\t" << obj->add(4) <<
  //   endl;
  //   ;

  cout << "\tSolution using heaps\n";
  nums = {4, 5, 8, 2};
  obj->KthLargestUsingHeaps(k, nums);
  int answer = obj->addUsingHeaps(3);
  cout << "\tKthLargest after adding 3 in the stream***********\t" << answer
       << " ******************\n";
  ;
  answer = obj->addUsingHeaps(5);
  cout << "\tKthLargest after adding 5 in the stream***********\t" << answer
       << " ******************\n";
  ;

  answer = obj->addUsingHeaps(10);
  cout << "\tKthLargest after adding 10 in the stream***********\t" << answer
       << " ******************\n";
  ;

  answer = obj->addUsingHeaps(9);
  cout << "\tKthLargest after adding 9 in the stream***********\t" << answer
       << " ******************\n";
  ;

  answer = obj->addUsingHeaps(4);
  cout << "\tKthLargest after adding 4 in the stream***********\t" << answer
       << " ******************\n";
  ;
}