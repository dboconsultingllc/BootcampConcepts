# Python Notes and Code

**Quick Tips**
* Create functions and classes so code is reusable.

* Use the [`Pylint`](https://www.pylint.org/) to perform linting on scripts to check for warnings, errors and syntax issues.

* Use [`Unit Test Package`](https://docs.python.org/3/library/unittest.html) for unit tests

* Many more tips to come!

## How to use Pylint
Assumimng you have Pylint installed and referencing the correct path, you'll need to execute pylint against your python file.

Windows example below
``` 
c:\users\me\pylint z:\mypathtomypythonfile\myfile.py 

```
After you execute against the file, you'll see output similar to below
``` 
************* Module s3bucketlist
z:\mypathtomypythonfile\myfile.py:1:0: C0114: Missing module docstring (missing-module-docstring)

------------------------------------------------------------------
Your code has been rated at 7.50/10 (previous run: 7.50/10, +0.00) 

```
The higher the score, the better. :joy:

**Note** A nice resource for [pylint codes](http://pylint-messages.wikidot.com/all-codes).