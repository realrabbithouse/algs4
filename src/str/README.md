
### Outlines

* **String Sorts**
  * LSD radix sort
  * MSD radix sort
  * 3-way radix quicksort
* **Tries**
  * R-way tries
  * ternary search tries
* **Substring Search**
  * Knuth-Morris-Pratt
  * Boyer-Moore
  * Rabin-Karp
* **Regular expressions**
  * Nondeterministic finite state automata (NFA)
* **Data Compression**

### Tandem repeat

A tandem repeat of a base string b within a string s is a substring of s consisting of
at least one consecutive copy of the base string s. Given b and s, design an algorithm
to find a tandem repeat of b within b of maximum length. Your algorithm should run
in time proportional to M+N, where M is length of b and N is the length s.

### Longest palindromic substring

Given a string s, find the longest substring that is a palindrome in expected linearithmic time.

### Extensions to NFA

Add to NFA.java the ability to handle multiway or, wildcard, and the + closure operator.
