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

/**
 * The struct greaters defines a functor that compares two integers and returns
 * true if the first integer is greater than the second.
 *
 */
struct greaters {
  bool operator()(const int &a, const int &b) { return a > b; }
};

int main() {
  vector<int> vect = {20, 30, 45, 15, 10};
  // Converts the vector into a heap
  make_heap(vect.begin(),
            vect.end()); // make-heap, creates a max-heap by default

  cout << "\tPrints the vector:\t";
  print(vect);

  cout << "\tInitial max heap element:\t";
  cout << vect.front() << endl;

  cout << "\tPop heap" << endl;
  pop_heap(vect.begin(),
           vect.end()); // Moves the largest element to the end of the array
  vect.pop_back();

  cout << "\tPrints the vector after pop:\t";
  print(vect);

  cout << "\tMax heap element after pop:\t";
  cout << vect.front() << endl;

  cout << "\tPush 100 into the heap" << endl;
  vect.push_back(100);

  cout << "\tHeap just after push_back(): ";
  print(vect);
  push_heap(vect.begin(), vect.end());
  cout << "\tHeap after push_heap(): ";
  print(vect);

  cout << "\tBefore sort is_heap(vect):\t" << is_heap(vect.begin(), vect.end())
       << endl;

  cout << "\tSort hep or Heap sort\n";
  sort_heap(vect.begin(), vect.end());
  cout << "\tPrinting the heap after sorting\t";
  print(vect);

  cout << "\tAfter sort is_heap(vect):\t" << is_heap(vect.begin(), vect.end())
       << endl;

  cout << "\tCreating a min-heap using make-heap\n";
  vector<int> minVect = {25, 54, 78, 95, 41, 10, 5, 7};
  cout << "\tPrints the minVect:\t";
  print(minVect);

  cout << "\tCreates min-heap from minVect\n";
  // TODO: What is the role of Compare comp, the last argument to make_heap
  // TODO below?
  make_heap(minVect.begin(), minVect.end(), greaters());
  cout << "\tPrints the minVect after creating a min-heap:\t";
  print(minVect);

  cout << "\tMin heap element is:\t";
  cout << minVect.front() << endl;
}