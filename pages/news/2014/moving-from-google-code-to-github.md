# Moving From Google Code To GitHub
<p class="date">July 20, 2014</p>

For a while now we have considered which project hosting website makes the most sense for us. Previously we have used Google Code, but this is now all changing. This article outlines the primary reasons why we moved to GitHub and what we think Google Code should work on to support open source projects in the future.

# What We Liked

Google Code is a great project hosting site and has aided us since *Jun 28, 2012*. It is certainly decent and being backed by Google gave us a lot of comfort in mind. We encourage others to try out Google Code because:

* It has great Git support.
* A very easy-to-use wiki system.
* The issue tracker works wonders.

# What We Didn't

Google Code does come with a few faults or lacking features though. Many of these are not the case with GitHub, which encouraged our move there. The things that we didn't like about it where:

* No Pull Requests: Contributing to a GitHub repository is clear and concise through forking and pull requests. Contributing to a Google Code repository involves submitting diffs over the issue tracker or giving, sometimes untrusted, users commit access to the repository. This has caused projects like Go itself to use third-party code review platforms like rietveld.
* Per repository issue trackers: GitHub offers this, but it's almost like multiple repositories don't belong in the same project on Google Code.
* Downloads Removal: Google Code users abused Downloads in some way [according to Google](http://google-opensource.blogspot.com/2013/05/a-change-to-google-code-download-service.html), the fix for this is that the Downloads section on Google Code is no more and it's now recommended to offer downloads through Google Drive. Although certainly not the end of the world, it is a bit strange from a support standpoint.
* Lack of programming API's / repository hooks: Google Code may support these but we couldn't find them anywhere. GitHub has well documented ones for many different features.
* Missing personal association: GitHub users show up with nice profile pictures, etc.

In conclusion, Google Code is a popular and decent project hosting website, but GitHub offers features we find better and it's increasing popularity among developers speaks for itself. Our future is with GitHub.
