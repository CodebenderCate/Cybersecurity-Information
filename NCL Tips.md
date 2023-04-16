# National Cyber League Resources
A list of resources that are useful for the National Cyber League Competition.

## Open Source Intelligence
- [Google](google.com): Google should almost always be your first stop for OSINT questions. 
  - Get familiar with Google's more advanced search operators. [Here](https://static.semrush.com/blog/uploads/files/39/12/39121580a18160d3587274faed6323e2.pdf) is a handy cheat sheet.
  - [Google reverse image search](https://www.google.com/imghp?hl=EN) is another great tool to help track down photos.
  - [Google Maps](https://www.google.com/maps) is great for tracking down buildings
  - [Google Earth](https://earth.google.com/web/) is cool, too
- [Wigle](wigle.net): Find and identify wifi networks. 
- [EXIF](https://exif.tools/): Exif tools can help reveal file meta data that can give answers out for easy OSINT questions.
- [OSINT Framework](https://osintframework.com/): Lots of OSINT tools sorted by category.

## Cryptography
- [CyberChef](https://gchq.github.io/CyberChef/): Convert between number bases and encoding schemes
- [Cryptii](https://cryptii.com/): Similar to CyberChef but it has some other options that may be useful. 
- [Rumkin](https://rumkin.com/tools/cipher/rail-fence/): Encode/Decode railfence ciphers
- [dcode.fr](https://www.dcode.fr/en) Encode/Decode many ciphers. This website is usually one of the first sites on google when you google for a specific cipher.
- [Cipher Identifier](https://www.boxentriq.com/code-breaking/cipher-identifier): Some websites can help you identify waht cipher is being used on a given cipher text. There are many good options. But, I have attached one here. 
- [Strings](https://www.howtogeek.com/427805/how-to-use-the-strings-command-on-linux/): The linux strings command will search through any file type and return strings of characters it found. Useful for finding flags in arbitrary file types. 
- [Digital Invisible Ink Toolkit](https://diit.sourceforge.net/): Find messages that are hidden in pictures. Many steganography questions use this tool.
- [Steghide](https://www.kali.org/tools/steghide/): Another way to find messages hidden in a picture.
  - Command to extract message: `steghide -sF <filename>`
- [Binwalk](https://www.kali.org/tools/binwalk/): Binwalk is used to extract files that are hidden in another file. Run `binwalk --dd='.*' <filename>` to extract all hidden files.

## Password Cracking
- [Hash Generator](https://www.browserling.com/tools/all-hashes): You will probably be asked to compute the hashes of some passwords. This site generates many hashes all at once. Select what you need and paste it into the answer box. 
- [Hashes.com](hashes.com): Lookup known hash values here. It is a good first step on any password cracking questions. Often, even a couple of the hard passwords can be found here.
- [CrackStation](https://crackstation.net/): Similar to hashes.com. Lookup hashes and see if there is a known password.
- [hashkill.io](hashkill.io): Similar to the last two. Used to lookup passsword hashes. 
- [Hashcat](https://www.kali.org/tools/hashcat/): Crack password hashes with a wordlist, brute force, or a combination.
  - **Common Flags**
  - -m: Hash Type - Select what type of hash to crack. 0 is an MD5 hash and is most common in NCL. ( -m 0)
  - -a: Attack Mode - Select what type of attack. 0 is a dictionary attack and is most common in NCL. (-a 0)
  - -o: Output - Select what file to output results to. (-O cracked.txt)
  - Example command: `hashcat -m 0 -a 0 -o cracked.txt target_hashes.txt wordlist.txt`
  - **Word Lists**: Many wordlists can be found in every kali install in `/usr/share/wordlists/`
  - **Rule Lists**: Google for the following rule lists
    - dive.rule
    - HoboRules
    - NSARules
    - OneRuleToRuleThemAll
- [Ophcrack](https://ophcrack.sourceforge.io/): Used to crack windows NTLM hashes. 
  - Download ophcrack [here](https://ophcrack.sourceforge.io/download.php?type=ophcrack)
  - Download table [here](https://ophcrack.sourceforge.io/tables.php): Most password for NCL are found in XP Free Small or XP Free Special.
- [CeWL](https://www.kali.org/tools/cewl/): CeWL (Custom Word List generator) is used to crawl a website an make a list of words that can be used for cracking a password. 
- [John the Ripper](https://www.kali.org/tools/john/): Used similarly to hashcat. John the Ripper is used to crack password hashes. I am not familiar with John the Ripper amd prefer hashcat so, I will not offer much of a guide on it
- [Mentalist](https://en.kali.tools/?p=1310): Generate large wordlists based on a dictionary, or create rules for hashcat or John the Ripper.
- [Hydra](https://www.kali.org/tools/hydra/): Brute force server passwords and try to connect.
- [SecLists](https://github.com/danielmiessler/SecLists): GitHub repo with great wordlists and lots of other hacking goodies.

## Log Analysis
- [cat](https://www.geeksforgeeks.org/cat-command-in-linux-with-examples/): Used to read the contets of a file `cat example.txt`
- [grep](https://www.geeksforgeeks.org/grep-command-in-unixlinux/): Search for strings or regular expressions to find information in a log. Regular expressions can be a great way to find specific information. But, they can be tricky.
  - **Common Flags**
  - -E: Use regular expressions [Cheat Sheet](https://ryanstutorials.net/linuxtutorial/cheatsheetgrep.php)
  - -o: Only out put the part of the line that matched the regular expression or string
  - -v: Invert the matches. Only output lines that do not match the regular expression or string.
  - -i: Ignore case
  - **Helpful Tips**
  - Pipe the output from one grep statement into another grep statement. This way you can filter down data in stages understand what you are doing.
  - Mix and match flags on consecutive grep statements to fine tune what you are looking for
  - **Example**: `cat log.log | grep -i "user" | grep -iv "invalid" | grep -E -o "([0-9]{1,3}[\.]){3}[0-9]{1,3}"`
  - **Explanation**: cat the log file and look for any line that contains the word "user" regardless of case to find any line that contains a usernam. Then, grep for anyline that says "invalid" and select only lines that do not match to sort out invalid user attempts. Finally, that last regular expression searches for IP addresses and only outputs the IP address.
- [cut](https://www.geeksforgeeks.org/cut-command-linux-examples/): Cut is used to split a line in bash. 
  - **Common Flags**;
  - -d: Specify a delimter
  - -f: Select which part of a cut you want to select
- [sort](https://www.geeksforgeeks.org/sort-command-linuxunix-examples/): Sort lines in alphabetical order.
- [uniq](https://www.geeksforgeeks.org/uniq-command-in-linux-with-examples/#:~:text=The%20uniq%20command%20in%20Linux,also%20deletes%20the%20duplicate%20lines.): Remove adjacent duplicate lines. It is important to sort the lines before doing this so that all duplicates are removed 
  - `cat log.txt | sort | uniq`
  - **Common Flags**
  - -c: Return the count of how many duplicate lines there were of each type
- [wc](https://www.geeksforgeeks.org/wc-command-linux-examples/#:~:text=wc%20stands%20for%20word%20count,specified%20in%20the%20file%20arguments.): Return the word count of a file. 
  - **Commmon Flags**
  - -l: Count the number of lines
  - -w: Count the number of bytes in the file
  - -m: Count the number of characters in the file
- [awk](https://www.geeksforgeeks.org/awk-command-unixlinux-examples/): Awk can be used to scan over files and pull information out by selecting certain fields. Awk is a little more complex than grep but is extremely maleable to do complex analysis. 
- [Excel](https://www.sans.org/white-papers/37745/): Excel can be a great resource to analyze log files quickly. Import a log as a csv file, and change the delimiter to something that makes sense for the log you are looking at.
- [VS Code](https://code.visualstudio.com/download): Any text editor can be good to view the data and start to understand what you are looking at. 
- [Python](https://www.pythontutorial.net/python-basics/python-read-text-file/#:~:text=To%20read%20a%20text%20file%20in%20Python%2C%20you%20follow%20these,the%20file%20close()%20method.): Python has amazing utilities to read and write text files. Python can be used for more complex analysis of many log files. 

## Network Traffic Analysis
- [Wireshark](https://www.kali.org/tools/wireshark/): Open, filter, and analyze pcap files.
  - [Plugins](https://www.wireshark.org/docs/wsug_html_chunked/ChPluginFolders.html): You can install additional plugins to extend the capabilities of Wireshark. For example, there is an h264 extractor to extract MP4 data from RTP packets and an ffmpeg plugin to reassemble media. 
- [AircapNG](https://www.kali.org/tools/aircrack-ng/#:~:text=aircrack%2Dng%20is%20an%20802.11,or%20simply%20by%20brute%20force.): Analyze wireless traffic and extract the password from unencryptes 802.11 traffic.
  - Run `aircrack-ng wifi.pcap` to show the wifi password
- [Zeek](https://docs.zeek.org/en/master/): Zeek is a very powerful  open-source network traffic analyzer. 
- [Capture Passwords with Wireshark](https://www.infosecmatter.com/capture-passwords-using-wireshark/): A neat walkthrough on how to find passwords for many different protocols using Wireshark.

## Forensics
- [EXIF](https://exif.tools/): Exif tools can help reveal file meta data that helps identify what a file is and where it came from.
- [Autopsy](https://www.kali.org/tools/autopsy/): Great for analyzing disk images and extracting files from a disc image.
  - [This](https://lepouvoirclapratique.com/blog/wp-content/uploads/2018/02/Tutorial_6_-_Kali_Linux_-_Sleuthkit.pdf) is a great walkthrough to get started with autopsy.
- [FTK Imager](https://www.exterro.com/ftk-imager#:~:text=FTK%C2%AE%20Imager%20is%20a,(FTK%C2%AE)%20is%20warranted): Used to disect forensic images. 

## Scanning & Reconnaissance
- [Nmap](https://nmap.org/book/man.html): Nmap is a great resource for enumerating open ports and what services are running on them. Nmap is an important tool and you should be very familiar with it for any CTF event.
  - **Commmon Flags**
  - -sV: Enable version detection
  - -sC: Run with standard scripts
  - -Pn: Do not ping - run scan even if the host looks like it is down
  - -p: Specify a port or port range (-p 8090, -p 100-1123, etc)
  - -O: Enable OS detection
  - Nmap Cheat Sheet [here](https://www.stationx.net/nmap-cheat-sheet/)
- [Dirbuster](https://www.kali.org/tools/dirbuster/): Dirbuster is used to brute force directories on a web server. It is super simple to use and quick to get going.
- [Gobuster](https://www.kali.org/tools/gobuster/): Similar to dirbuster, used to scan for directories on a web server
- [Wappalyzer](https://www.wappalyzer.com/): Wappalyzer is a browser extention that allows you to 
- [Netcat](https://allabouttesting.org/netcat-tutorial-for-beginners/): Netcat can be used to connect to a remote machine and enumerate what is running on that machine.
  - To connect to a machine run `nc <hostname orIP> <port num>`
- [Telnet](https://www.process.com/docs/tcpware6_1/users/chapter_12.html](https://www.javatpoint.com/linux-telnet-command ): Like netcat, telnet is used to connect to a remote computer. Once you are connected, you can send the computer commands and try to figure out what it is running. 
  - To connect to a remote machine run `telnet <hostname or IP> <port num>`
- [Metasploit](https://docs.rapid7.com/metasploit/getting-started/): Metasploit is not just for hacking. There are a lot of good tools for scanning and recon as well. For instance, the ftp_version module allows you to find what version of an ftp server a server is running. No hacking required!

## Web Application Exploitation
- [Web Browser](https://www.webdevelopersnotes.com/browsers-list): There are many web browsers out there. Pick one and start exploring the website. Knowing how the website functions normally can help you understand where vulnerabilities may exist. 
- [Developer Tools](https://nira.com/chrome-developer-tools/): Hit F12 or right click and inspect a web page to view the developer tools that are available in any browser. This lets you see the 
  - **Dev Tool Tips**
  - Read the HTML carefully. Hints are usually hidden in the HTML, especially in earlier challenges.
  - Pay attention to the network tab and understand what requests are being made. 
  - Look at the cookies! Cookies can be found under teh application tab. Often times, a challenge can be completed by just changing a cookie value. 
- [Burp Suite](https://portswigger.net/burp/documentation/desktop/penetration-testing): Burp Suite is a great tool for figuring out how a website works and hacking into it. Burp Suite has many capabilities. So, the most important features are briefly discussed below.
  - [Proxy](https://portswigger.net/burp/documentation/desktop/tools/proxy/getting-started#:~:text=Burp%20Proxy%20lies%20at%20the,traffic%20passing%20in%20both%20directions.): Proxy your web traffic through Burp and analyze what is happening
  - [Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater/using): After intercepting traffic in the proxy, got to "Actions" > "Send to Repeater" to use the repeater with a web request. You can view, modify, and send web requests to the server, and catch the response so that you can analyze how the web server is working. 
  - [HTTP History](https://portswigger.net/burp/documentation/desktop/tools/proxy/history/getting-started-http-history#:~:text=The%20HTTP%20history%20tab%20displays,useful%20parameters%20and%20other%20inputs.): This is a list of all the URLs that have been sent through the proxy. This can let you see what resources are being requested and look for any interesting or anomalous requests. 
- [Postman](https://www.postman.com/downloads/): Postman helps create and repeate API calls. This can help get information out of a server anv easily view what the server is sending back. It is easy to save API calls for later to rerun or modify after you discover more information. 
- [sqlmap](SQLMap): sqlmap is used to automatically detect and exploit sql injection vulnerabilities.
- [Exploit DB](https://www.exploit-db.com/): Once you know what services or what OS is running on a server, look that up in Exploit DB and see what vulnerabilities may exist that you can take advantage of. 
- [Metasploit](https://docs.rapid7.com/metasploit/getting-started/): Metasploit is very versitile and can help with many problems. Once you understand the server, what is running on it, and what vulnerabilities it may have, jump into Metasploit and see if there are any modules that may be helpful. 

## Enumeration & Exploitation
- [Ascii Table](https://www.asciitable.com/): Many enumeration and exploitation questions use ascii in some way. So, knowing how to decode ascii is an important skill. 
- [Python](https://www.python.org/): Lots of the easy Enumeration & Exploitation problems use python. Understanding how to read and write python code is a good skill to have. 
- [Bash](https://www.gnu.org/software/bash/): Same as above, except its Bash this time. Since Bash is used for the linux command line, it is used for many of the utilities mentioned in this document. So, you need to be familiar with this, and it is often seen in Enumeration & Exploitation problems. 
- [Decompilers](https://en.wikipedia.org/wiki/Decompiler): Sometimes you will be given code that has already been compiled. A decompiler can take this compiled code and help us understand what it is doing.
  - [Python Decompiler Online](https://www.toolnb.com/tools-lang-en/pyc.html)
  - [Java Decompiler Online](http://www.javadecompilers.com/)
- [Ghidra](https://www.kalilinux.in/2021/06/ghidra-reverse-engineering-kali-linux.html): Ghidra is a tool that was developed by the NSA to reverse engineer software. Now it is open source and comes installed on every instance of kali linux. It procides a lot of good information but it can be difficult to understand. 
- [GNU Debugger](www.cs.umd.edu/~srhuang/teaching/cmsc212/gdb-tutorial-handout.pdf): GDB is used to step through code and is used primarily for debugging. But, it can be used to understand code that has already been compiled.

## Other Tips
- Pay close attention to the title of the challenge. It's usually a hint at how to solve the problem.
- When I'm stuck on a challenge, I often go look at the gymnasium again to see if there are any tips or tools they list in the gym that might help me.
