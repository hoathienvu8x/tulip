---
title: Retrieving Google’s Cache for a Whole Website
excerpt: Some time ago, as some of you noticed, the web server that hosts my blog went down. Unfortunately, some of the sites had no proper backup, so some thing had to be done in case the hard disk couldn&amp;…
date:2009-08-15 05:06:33
author:Guy Rutenberg
tags:backup, Google
categories:Python
---

Some time ago, as some of you noticed, the web server that hosts my blog went down. Unfortunately, some of the sites had no proper backup, so some thing had to be done in case the hard disk couldn’t be recovered. My efforts turned to Google’s cache. Google keeps a copy of the text of the web page in it’s cache, something that is usually useful when the website is temporarily unavailable. The basic idea is to retrieve a copy of all the pages of a certain site that Google has a cache of.

While this is easily done manually when only few pages are cached, the task needs to be automated when a need for retrieving several hundreds of pages rises. This is exactly what the following Python script does.

```python
#!/usr/bin/python
import urllib
import urllib2
import re
import socket
import os
socket.setdefaulttimeout(30)
#adjust the site here
search_term="site:guyrutenberg.com"
def main():
    headers = {'User-Agent': 'Mozilla/5.0 (X11; U; Linux i686 (x86_64); en-US; rv:1.8.1.4) Gecko/20070515 Firefox/2.0.0.4'}
    url = "http://www.google.com/search?q="+search_term
    regex_cache = re.compile(r'<a href="(http://\d*\.\d*\.\d*\.\d*/search\?q\=cache.*?)".*?>Cached</a>')
    regex_next = re.compile('<a href="([^"]*?)"><span id=nn></span>Next</a>')

    #this is the directory we will save files to
    try:
        os.mkdir('files')
    except:
        pass
    counter = 0
    pagenum = 0
    more = True
    while(more):
        pagenum += 1
        print "PAGE "+str(pagenum)+": "+url
        req = urllib2.Request(url, None, headers)
        page = urllib2.urlopen(req).read()
        matches = regex_cache.findall(page)
        for match in matches:
            counter+=1
            tmp_req = urllib2.Request(match.replace('&amp;','&'), None, headers)
            tmp_page = urllib2.urlopen(tmp_req).read()
            print counter,": "+match
            f = open('files/'+str(counter)+'.html','w')
            f.write(tmp_page)
            f.close()
        #now check if there is more pages
        match = regex_next.search(page)
        if match == None:
            more = False
        else:
            url = "http://www.google.com"+match.group(1).replace('&amp;','&')

if __name__=="__main__":
    main()

# vim: ai ts=4 sts=4 et sw=4
```

Before using the script you need to adjust the search_term variable near the beginning of the script. In this variable goes the search term for which all the available cache would be downloaded. E.g. to retrieve the cache of all the pages of http://www.example.org you should set search_term to site:www.example.org
