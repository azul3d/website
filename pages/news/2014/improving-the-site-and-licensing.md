# Improving The Site And Licensing
<p class="date">September 15, 2014</p>

Today we're happy to announce several improvements to the website and licensing information of Azul3D. We truly want this website to be a comprehensive database about all things Azul3D.

# The Site

This is by no means a slam on Go HTML templates -- they incredibly useful and *in fact we're still using them to render the final HTML pages*. But what started out as just [a few minor changes](https://groups.google.com/forum/#!topic/azul3d/iDeeEHDVbok) to the website eventually turned into moving all of the news articles, documentation (including installation instructions and FAQ), etc over to Markdown instead of Go HTML templates.

We have found that Markdown is more suitable for writing content like news articles and documentation, and looking back it feels like we were abusing templates before. Markdown documents easier to read and write by hand, in addition to the fact that they can be styled independantly and their rendered HTML can be fed into a Go HTML template.

To render the Markdown we are using *Russ Ross*'s Markdown parser, [Blackfriday](https://github.com/russross/blackfriday), which we have found to be very fast both in sheer speed and with it's incredibly clear API. We're really happy with it!

You can see how much more clear and concise this method is by comparing an old Go [HTML template](https://raw.githubusercontent.com/azul3d/cmd-webgen/9a76decd02028b4d7c2ca3b8734eb7c031a3b6b7/pages/news/2014/mac-osx-support.tmpl) with it's new [Markdown version](https://raw.githubusercontent.com/azul3d/cmd-webgen/master/pages/news/2014/mac-osx-support.md).

Additionally, there were several pages that did not have any dedicated navigation links to them, except through the [FAQ](/doc/faq.html). This started to become a problem because the [FAQ](/doc/faq.html) *started to look like our navigation area* (how strange!).

We've made a [Doc](/doc/) page in the navigation bar that serves solely to give you an overview of all the Azul3D documentation (excluding [package documentation](/packages.html)) available on the website.

# Licensing

Previously, we kept a copy of the `LICENSE` file in each package directory. This works well because no matter which package you clone the source code also has a copy of the license, *we're keeping this practice*. But for some people understanding legal documents like these is just plain time consuming. *We're not an alternative to a real lawyer, and we can't provide legal advice to you* but to help with understanding the license there is now a dedicated [license page](/doc/license.html).

Also, each package had an `AUTHORS` file distributed with it. This was kind of a mishap on our part because none of them were in synchronization. The question we really want to answer is *"Who all has actually worked on this code?"* -- because Azul3D truly is an open source project and it needs the help of people like you.

We've removed all of the `AUTHOR` files and instead made links to the new [authors page](/doc/authors.html) which is maintained as an always up-to-date list of who has contributed to Azul3D. For those wanting to see who has actually worked on what -- `git` history is much better than these old `AUTHOR` files anyway.

*Now contributers get the proper credit that they deserve*!

One slight concern in moving to an online [authors page](/doc/authors.html) was spam (we want to give credit to you, not spam you), and for this reason we've made the file (optionally) have ROT13 encoded email addresses in it. The email addresses are decoded client-side in Javascript, so that *most* simple web crawlers won't find the proper email addresses.

# Conclusion

We hope to increase the amount and quality of content being brought to you in a legible form through the Azul3D website. It's one of the things that we are constantly trying to improve.

In the future we hope to provide detailed documentation relating to Azul3D packages, for instance we want to document how the graphics pipeline works, how digital audio works, etc all with real-world Azul3D code examples.

