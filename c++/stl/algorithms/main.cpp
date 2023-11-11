#include <algorithm>
#include <iostream>
#include <numeric> // For accumulate, iota operation
#include <vector>

using namespace std;

void printArray(int a[], int n, bool printAsChar) {
  for (int i = 0; i < n; ++i) {
    cout << "\t";
    if (printAsChar) {
      printf("%c\t", a[i]);
    } else {
      printf("%d\t", a[i]);
    }
  }
  cout << endl;
}

void sorting(int numbers[], int characters[]) {}

/**
 * The main function sorts an array of numbers and an array of characters and
 * then prints them out.
 * @return The main function is returning an integer value of 0.
 */
int main() {
  int numbers[] = {4, 8, 2, 1, 47, 80, 0, 2, 6, 9};
  int characters[] = {97, 101, 105, 111, 113, 121, 122};
  int n = sizeof(numbers) / sizeof(numbers[0]);
  int c = sizeof(characters) / sizeof(characters[0]);

  cout << "\t**************************************SORTING*********************"
          "**"
          "***************"
       << endl;
  cout << "\tPrinting the unsorted arrays::" << endl;
  printArray(numbers, n, false);
  printArray(characters, c, true);

  cout << "\tSorting the arrays::" << endl;
  sort(begin(numbers), end(numbers));
  sort(begin(characters), end(characters));

  cout << "\tPrinting the sorted arrays::" << endl;
  printArray(numbers, n, false);
  printArray(characters, c, true);

  cout << "\tSorting the arrays in descending order::" << endl;
  sort(numbers, numbers + n, greater<int>()); // without using begin() & end()
  sort(begin(characters), end(characters), greater<int>());

  cout << "\tPrinting the sorted arrays in descending order::" << endl;
  printArray(numbers, n, false);
  printArray(characters, c, true);

  cout << "\t**************************************SEARCHING*******************"
          "**"
          "**"
          "***************"
       << endl;

  // Since Binary Search works on an Ascending Sorted Array
  sort(numbers, numbers + n);
  if (binary_search(numbers, numbers + n, 80)) {
    cout << "\t80 is found in the numbers array" << endl;
  } else {
    cout << "\t80 is NOT found in the numbers array" << endl;
  }

  cout << "\t**************************************VECTOR*******************"
          "**"
          "**"
          "***************"
       << endl;
  // Initializing vector from array values
  vector<int> vect(numbers, numbers + n);
  cout << "\tVector is: " << endl;
  for (int i = 0; i < n; i++) {
    cout << "\t" << vect[i];
  }
  cout << endl;

  vector<int> unsortedNumbers(numbers, numbers + n);
  for (int i = 0; i < n; i++) {
    unsortedNumbers[i] = (i + 1) * 10;
  }

  reverse(unsortedNumbers.begin(), unsortedNumbers.end());

  cout << "\tUnsorted Vector is: " << endl;
  for (int i = 0; i < n; i++) {
    cout << "\t" << unsortedNumbers[i];
  }
  cout << endl;

  // sort the vectors in descending order
  sort(unsortedNumbers.begin(), unsortedNumbers.end());
  cout << "\tSorted Vector is: " << endl;
  for (int i = 0; i < n; i++) {
    cout << "\t" << unsortedNumbers[i];
  }
  cout << endl;

  cout << "\tReversing the vector `vect`";
  reverse(vect.begin(), vect.end());
  cout << "\tVector after reversing is: " << endl;
  for (int i = 0; i < n; i++) {
    cout << "\t" << vect[i];
  }
  cout << endl;

  cout << "\tMax element of `vect` is "
       << *max_element(vect.begin(), vect.end()) << endl;
  cout << "\tMin element of `unsortedNumbers` is "
       << *min_element(unsortedNumbers.begin(), unsortedNumbers.end()) << endl;

  cout << "\tAccumulation of `vect` element is "
       << accumulate(vect.begin(), vect.end(), 0) << endl;

  cout << "\tCount the occurence of 2 in the vector `vect`\t"
       << count(vect.begin(), vect.end(), 2) << endl;

  cout << "\tFinds 40 in unsortedNubmers vector" << endl;
  find(unsortedNumbers.begin(), unsortedNumbers.end(), 40) !=
          unsortedNumbers.end()
      ? cout << "\t\tElement found" << endl
      : cout << "\t\tElement not found" << endl;

  cout << "\t**************************************ARRAY "
          "ALGORITHMS*******************"
          "**"
          "**"
          "***************"
       << endl;
  int num[6] = {1, 2, 3, 4, 5, -6};
  printArray(num, 6, false);
  cout << "\tall_of()\n";
  all_of(num, num + 6, [](int x) { return x > 0; })
      ? cout << "\tAll elements are positive\n"
      : cout << "\tAll elements are not positive\n";

  cout << "\tany_of(): Checking if any number is negative\n";
  any_of(num, num + 6, [](int x) { return x < 0; })
      ? cout << "\tThere is a negative number\n"
      : cout << "\tNo negative numbers found\n";

  cout << "\tany_of(): Checking if any number is negative in `unsortedNumbers "
          "vector`\n";
  any_of(unsortedNumbers.begin(), unsortedNumbers.end(),
         [](int x) { return x < 0; })
      ? cout << "\tThere is a negative number\n"
      : cout << "\tNo negative numbers found\n";

  int copiedArray[6];
  cout << "\tPrinting the copied array\n";
  printArray(copiedArray, 6, false);

  cout << "\tCopying `num` to `copiedArray`\n";
  copy_n(num, 6, copiedArray);

  cout << "\tPrinting the copied array after copy_n() operations\n";
  printArray(copiedArray, 6, false);

  cout << "\tUsing iota() to assign continuous values to the array\n";
  int continuousValues[10];
  iota(continuousValues, continuousValues + 10, 10);
  printArray(continuousValues, 10, false);
  return 0;
}
