## Interview Questions


1. Social network connectivity. Given a social network containing _n_ members and a log file containing _m_
   timestamps at which times pairs of members formed friendships, design an algorithm to determine the earliest
   time at which all members are connect (i.e., every member is a friend of a friend of a friend ... of a friend).
   Assume that the log file is sorted by timestamp and that friendship is an equivalnce relation. The running time
   of your algorithm should be _m_ log _n_ or better and use extra space proportional to _n_.

- A: create a weighted union find that iterates over the list of timestamps unioning all of the friends, if the 
     ending component size is zero than all friends are connected.


2. Union-find with specific canonical element. Add a method `find()` to the union-find data type so that `find(i)`
   returns the largest element in the connected component containing `i`. The operations, `union()` `connected()`,
   and `find()` should all take logartithmic time or better.

   For example, if one of the connected components is `{1, 2, 6, 9}`, then the `find()` method should return `9`
   for each of the four elements in the connected components.

- A: create a compressed weighted union that stores the root as the highest value.


3. Successor with delete. Given a set of _n_ integers `S = {0, 1, ..., n - 1}` and a sequence of requests of
   the following form:
   - Remove _x_ from _S_
   - Find the succesor of _x_: the smallest _y_ in _S_ such that `y >= x`.


- A: union find to mark successor of each integer, when an integer is removed union find connects to the next
     integer in the list.
