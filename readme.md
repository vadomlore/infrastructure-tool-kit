# IE Terraform file ddd variable tool pre-process 

Introduction
------------

the program enable you to add variables in the template folder map to the
target files, it is able to extend the files with your logic of adding
variables. this is just a pre batch add variable, you need to check the
result after the processor. 




Example
-------
main.exe /usr/local/ie/terrform ./template true

* the first parameter is the file search path
* the second parameter is the adding information metadata 
* the third parameter will overwrite the original file if set to true, otherwise 
will generate a new file with -temp suffix in the same folder as the original file.
