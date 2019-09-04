##Simplify Path

###题目
Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.
// canonical 最简洁的

In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level. 
For more information, see: Absolute path vs relative path in Linux/Unix

Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. 
The last directory name (if it exists) must not end with a trailing /. 
Also, the canonical path must be the shortest string representing the absolute path.

Example:
```
Input: "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.

Input: "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.

Input: "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.

Input: "/a/./b/../../c/"
Output: "/c"

Input: "/a/../../b/../c//.//"
Output: "/c"

Input: "/a//b////c/d//././/.."
Output: "/a/b/c"
```