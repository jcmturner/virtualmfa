# Virtual MFA

A simple command line tool to generate time based one time passwords from a MFA base32 encoded secret string.

Written to work with AWS' virtual MFA implementation.

## Usage
```
$ ./virtualmfa
Enter MFA Secret: 
```
Paste in your MFA secret and one time passwords will be return with a bar indicating its lifetime.
New one time passwords will be continuously generated until the command is exited with Ctrl-C

For example:
```
$ ./virtualmfa
Enter MFA Secret: 
709801  ╢▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░░░░░╟ 
```
