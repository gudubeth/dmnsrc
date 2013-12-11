Domain Search Tool
===============================

Common line tool for batch checking domain existence. It does this by adding 
suffixes and prefixes to the word list given in words.txt. Only checks .com
domains. 

This is my first real go application and it's also barely working now. I'm
planning to improve the code and add features soon.

Planned features
----------------
* Improve domain check mechanism
* Set multiple suffixes and/or prefixes at once
* Option to use files for suffixes and prefixes
* Option to get full domain names from file. ALso pipelining for this feature
* Use regular expressions for selecting games from word list
* Option to choose max/min domain length
* Flexible TLD selecting (better defaults, file, command line parameters, etc.)
* Output to file
* Sanitize word lists
* Default word libraries
* Option to use more than one word list file
* Concurrent domain check
