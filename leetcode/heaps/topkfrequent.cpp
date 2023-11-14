#include <algorithm>
#include <iostream>
#include <map>
#include <queue> // for priority_queue
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
  vector<int> topKFrequent(vector<int> &nums, int k) {
    if (k == nums.size()) {
      return nums;
    }

    map<int, int> frequencies;
    for (int i : nums) {
      frequencies[i] += 1;
    }

    /*
     Initialise a heap with most frequent elements at the top
     It is creating a lambda function called `comp` that compares two
    integers `n1` and `n2` based on their frequencies in the `frequencies` map.
    */
    auto comp = [&frequencies](int n1, int n2) {
      return frequencies[n1] > frequencies[n2];
    };

    /* The line `priority_queue<int, vector<int>, decltype(comp)> heap(comp);`
    is creating a priority queue named `heap` of integers. */
    priority_queue<int, vector<int>, decltype(comp)> heap(comp);

    for (pair<int, int> p : frequencies) {
      heap.push(p.first);
      if (heap.size() > k) {
        heap.pop();
      }
    }

    vector<int> top(k);
    for (int i = k - 1; i >= 0; i--) {
      top[i] = heap.top();
      heap.pop();
    }
    return top;
  }
};

int main() {
  cout << "\tTop K Frequent Elements: ";
  vector<int> nums = {1, 1, 1, 2, 2, 3};
  int k = 2;
  Solution solution;
  print(nums);
  vector<int> top = solution.topKFrequent(nums, k);
  cout << "\t";
  for (int i : top) {
    cout << i << "\t";
  }
  cout << endl;
}