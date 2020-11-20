pget is a script that calls the go program to grab a webpage 
(or any file, really) and pipes it to a formatter. I used
prettier because it is available on my system. It is a dependency
for many js frameworks since they use it in their CLI.
after the formatting, the output is piped to pymentize. This creates
colored output (ansi screen codes by default) that enhances
the readability of the code. 
