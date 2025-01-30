# Don't use directories just to organize files

This is a subtle but important point, especially if you're new to Go.

You shouldn't create new directories just to organize your .go files. In Go, creating a directory creates a new package, and placing a file in that directory makes it part of that package.

Create a directory only when you have a specific reason to create a new package – not because you want a neater/cleaner/clearer directory structure for your files.
