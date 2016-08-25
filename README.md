# bbcrowler
Bug Bounty Crawler

# Functionality of crawler

* Fetching hackerone.com new programs using JSON parsing
* Fetching new programs from bugcrowd.com using HTML parsing
* Fetching new hacktivity events from hackerone.com using JSON parsing

# History

* This crawler created special for using inside any type of information bot and adapted for usage at Heroku

# Architecture

* In-memory data saving using standard Golang maps
* For synchronisation purposes used Firebase storage

# Example of usage 

* See ```example/main.go```

# MIT License

Copyright (c) 2016 Maddevs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.