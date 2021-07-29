
## Description 
Стек — это коллекция, элементы которой получают по принципу «последний вошел, первый вышел». 
LIFO - last in first out

## API

- push(string val) void  { // insert new string into stack }
- pop() string { // remove string from stack, and return removed val}
- isEmpty() bool { // check is the stack empty }
- size() int { // size }
- iterate() ??? { // ??? }

## Perfomance

Для Linked-list

|   #  |   best  | worst |
| :--- |  :----: | :---: | 
| push |    1    |   1   | 
| pop  |    1    |   1   |
| size |    1    |   1   |

Для Resizing-array

|   #  |   best  | worst | amortized |
| :--- |  :----: | :---: | :--: |
| push |    1    |   N   |   1  |
| pop  |    1    |   N   |   1  |
| size |    1    |   1   |   1  |

Amortized = Efficient solution.
- push(): double size of array s[] when array is full.
- pop(): halve size of array s[] when array is 1/4 full.

Linked-list implementation.
- Every operation takes constant time in the worst case.
- Uses extra time and space to deal with the links.

Resizing-array implementation.
- Every operation takes constant amortized time.
- Less wasted space.
