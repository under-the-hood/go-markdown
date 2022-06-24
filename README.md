# 🔬 Go Markdown research

Research Markdown parsers written on Go.

## Motivation

In developing [✨Sparkle](https://sparkle.wiki/), I've encountered
a need to select an appropriate library for handling Markdown files
because they are at the heart of PKM[^1].

## Requirements

1. **High-level data manipulation:** The library must facilitate
high-level manipulation of Markdown data. This involves providing
capabilities to programmatically interact with and modify the Markdown
content, such as adding, deleting, or altering sections or elements
within the Markdown files. This feature is crucial for dynamic content
generation and customization.

2. **Simplifying validation and transformation of notes:** The library
should streamline the processes of validating and transforming Markdown notes.
This includes checking for syntax correctness, ensuring link integrity,
and converting Markdown into other formats (like HTML). This feature is
essential for maintaining the quality and usability of the Markdown content,
especially when dealing with a large volume of notes or complex documentation.

3. **Concise and user-friendly API:** As with YAML libraries, a brief
and intuitive API is paramount. The library should enable developers
to perform everyday Markdown processing tasks with minimal and straightforward
code. A well-designed API reduces development time, simplifies maintenance,
and makes the library more accessible to developers who might not be deeply
familiar with Markdown's intricacies.

## Resources

- https://github.com/russross/blackfriday
  - github.com/russross/blackfriday/v2
- https://github.com/yuin/goldmark
  - github.com/yuin/goldmark

## The result

Interim results or a detailed follow-up.[^2]

<p align="right">made with ❤️ for everyone by <a href="https://www.octolab.org/">OctoLab</a></p>

[^1]: Personal Knowledge Management.

[^2]: work in progress
