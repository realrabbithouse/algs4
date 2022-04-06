
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
  * Run-length Coding
  * Huffman Compression
  * LZW Compression

### Huffman Compression
1. Use different number of bits to encode different chars.

2. How do we avoid ambiguity?

  Ex 1. Fixed-length code.
  Ex 2. Append special stop char to each codeword.
  Ex 3. General prefix-free code.
  ==Ensure that no codeword is a *prefix* of another.==

3. How to represent the prefix-free code?

   A binary trie!

4. Compression.

   Method 1: start at leaf; follow path up to the root; print bits in reverse.
   Method 2: create ST of key-value pairs.

5. Expansion.

6. How to write the trie?

   Write preorder traversal of trie; mark leaf (1) and internal nodes (0) with a bit.

7. How to read in the trie?

   Reconstruct from preorder traversal of trie.

8. How to find best prefix-free code?

   *Huffman algorithm demo*
   Count frequency for each character in input.
   Start with one node corresponding to each char i (with weight freq[i]).
   Repeat until single trie formed:
   – select two tries with min weight freq[i] and freq[j]
   – merge into single trie with weight freq[i] + freq[j]



### Tandem repeat

A tandem repeat of a base string b within a string s is a substring of s consisting of
at least one consecutive copy of the base string s. Given b and s, design an algorithm
to find a tandem repeat of b within b of maximum length. Your algorithm should run
in time proportional to M+N, where M is length of b and N is the length s.

### Longest palindromic substring

Given a string s, find the longest substring that is a palindrome in expected linearithmic time.

### Extensions to NFA

Add to NFA.java the ability to handle multiway or, wildcard, and the + closure operator.
