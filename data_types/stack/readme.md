LIFO - last in first out

## API

- push(string val) void  { // insert new string into stack }
- pop() string { // remove string from stack, and return removed val}
- isEmpty() bool { // check is the stack empty }
- size() int { // size }
- iterate() ??? { // ??? }

## Perfomance

|   #  |   best  | worst |
| :--- |  :----: | :---: | 
| push |    1    |   N   | 
| pop  |    1    |   N   |
| size |    1    |   1   |