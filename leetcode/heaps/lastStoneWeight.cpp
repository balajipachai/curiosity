/**
 * Problem:
 * You are given an array of integers stones where stones[i] is the weight of
 * the ith stone. We are playing a game with the stones. On each turn, we choose
 * the heaviest two stones and smash them together. Suppose the heaviest two
 * stones have weights x and y with x <= y. The result of this smash is: If x ==
 * y, both stones are destroyed, and If x != y, the stone of weight x is
 * destroyed, and the stone of weight y has new weight y - x. At the end of the
 * game, there is at most one stone left.
 *
 * Return the weight of the last remaining stone. If there are no stones left,
 * return 0.
 *
 * Example 1
 * Input: stones = [2,7,4,1,8,1]
 * Output: 1
 * Explanation:
 *
 * We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
 * we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
 * we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
 * we combine 1 and 1 to get 0 so the array converts to [1] then that's the
 * value of the last stone.
 *
 * Example 2:
 * Input: stones = [1]
 * Output: 1
 */

#include <algorithm>
#include <iostream>
#include <queue>
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

class Solution {
public:
  /**
   * The function "lastStoneWeight" takes a vector of integers representing the
   * weights of stones and returns the weight of the last stone remaining after
   * repeatedly smashing the two heaviest stones together until there is only
   * one stone left.
   *
   * @param stones The parameter "stones" is a reference to a vector of
   * integers.
   *
   * @return the last stone weight, which is the weight of the stone that
   * remains after all other stones have been destroyed.
   */
  int lastStoneWeight(vector<int> &stones) {
    int x, y;
    priority_queue<int> heap;
    // Add stones to the heap array
    for (int i = 0; i < stones.size(); i++) {
      heap.push(stones[i]);
    }

    while (heap.size() != 1) {
      y = heap.top();
      heap.pop();
      x = heap.top();
      heap.pop();
      int diff = y - x;
      heap.push(diff);
    }
    return heap.top();
  }
};

int main() {
  cout << "\tLast stone weight: ";
  vector<int> stones = {2, 2};
  Solution solution;
  print(stones);
  int answer = solution.lastStoneWeight(stones);
  cout << "\tLast stone weight = \t" << answer << endl;

  cout << "***********************************" << endl;
  stones = {2, 7, 4, 1, 8, 1};
  print(stones);
  answer = solution.lastStoneWeight(stones);
  cout << "\tLast stone weight = \t" << answer << endl;
}