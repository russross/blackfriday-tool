Blackfriday Tool
=================

Blackfriday is a [Markdown][1] processor implemented in [Go][2]. It
lives on github, where you can find more information about it:

* <http://github.com/russross/blackfriday>

Blackfriday-tool is an example command-line tool that uses
blackfriday to process markdown input. It provides a complete
example of how to use blackfriday in a Go project.

Here is the help message for the tool:

    Blackfriday Markdown Processor v1.1
    Available at http://github.com/russross/blackfriday

    Copyright © 2011 Russ Ross <russ@russross.com>
    Distributed under the Simplified BSD License
    See website for details

    Usage:
      ./blackfriday-tool [options] [inputfile [outputfile]]

    Options:
      -cpuprofile="": Write cpu profile to a file
      -css="": Link to a CSS stylesheet (implies -page)
      -format=html: output format: html, latex, or deck
      -fractions=true: Use improved fraction rules for smartypants
      -latexdashes=true: Use LaTeX-style dash rules for smartypants
      -page=false: Generate a standalone HTML page (implies -latex=false)
      -repeat=1: Process the input multiple times (for benchmarking)
      -smartypants=true: Apply smartypants-style substitutions
      -toc=false: Generate a table of contents (implies -latex=false)
      -toconly=false: Generate a table of contents only (implies -toc)
      -xhtml=true: Use XHTML-style tags in HTML output


License
-------

Blackfriday is distributed under the Simplified BSD License:

> Copyright © 2011 Russ Ross. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
> 
> The views and conclusions contained in the software and documentation are those of the
> authors and should not be interpreted as representing official policies, either expressed
> or implied, of the copyright holder.


   [1]: http://daringfireball.net/projects/markdown/ "Markdown"
   [2]: http://golang.org/ "Go Language"
   [3]: http://github.com/tanoku/upskirt "Upskirt"
