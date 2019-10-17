# About Litterarum
Litterarum is documentation system for my personal projects.
Litterarum is far superior to [Markdown](https://daringfireball.net/projects/markdown/) for many reasons.

# Language Specificiation

## Sections
Sections are used to break up text into categories.
The semantic meaning of a Section is specific to what [docset](#about/docset) is in use.
Sections are defined as:
`#section#`

## Subsections
Subsections are used to break up text within sections into categories.
Subsections are defined as:
`=subsection=`

## Lists
Guess what? Lists are used to group information into lists.
The characters used to define different types of lists are specific to what [docset](#about/docset) is in use.
By default Lists are defined as:
```
+ item
+ item 2
```

## Notes
Notes are used to remind the reader of possible relevant information.
Notes do not create tags, therefore cannot be referenced specifically.
If a Note is present within a [Section](#sections) or [Subsection](#subsection)
Notes are defined as:
`-note-`

## Blocks
Blocks are used to denote blocks of text seperate from the body/
Blocks are defined as:
`
`` code ``
`

# Comments
Comments are ignored by the parser.
Comments are defined as:
`
// comment
`

and can optionally be used inline:
`// inline comment //`

# Literals
Literals are used to embed literal characters into text.
Literals are defined as:
`\char`

# References
References are used to link to materials internal to the document:
`(about)[#about#]`

or used to link to materials external to the document:
`(golang documentation)[http://golang.org/doc]`

References are defined as:
`(text)[location]`

# Additional Information #
Whitespace is ignored outside of [Blocks](#blocks) and [References](#references)
Sections, subsections, and notes are case insensitive.

## Parser
The parser creates a file in either Markdown format or HTML format.

## Docset
The default `docset` is described within this document but custom `docset`s can be written.

# Credits #
 Andruw Earnhardt
