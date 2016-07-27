package file

/// Contains css and one method to return it
/// Css source modified from:
/// https://github.com/sindresorhus/github-markdown-css/blob/gh-pages/github-markdown.css
object CSS {
  // ...
  final val css = """      @font-face {
        font-family: octicons-link;
        src: url(data:font/woff;charset=utf-8;base64,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) format('woff');
      }
      
      body {
        -ms-text-size-adjust: 100%;
        -webkit-text-size-adjust: 100%;
        color: #333;
        font-family: "Helvetica Neue", Helvetica, "Segoe UI", Arial, freesans, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
        font-size: 16px;
        line-height: 1.6;
        word-wrap: break-word;
        background-color: #eeeeee;
        margin-top: 0;
        margin-bottom: 0;
      }
      
      .enaml {
        width: 75%;
        padding: 15px;
        margin: auto;
        background-color: #fff;
      }
      
      a {
        background-color: transparent;
        -webkit-text-decoration-skip: objects;
      }
      
      a:active,
      a:hover {
        outline-width: 0;
      }
      
      strong {
        font-weight: inherit;
      }
      
      strong {
        font-weight: bolder;
      }
      
      h1 {
        font-size: 2em;
        margin: 0.67em 0;
      }
      
      img {
        border-style: none;
      }
      
      code {
        font-family: monospace, monospace;
        font-size: 1em;
      }
      
      hr {
        box-sizing: content-box;
        height: 0;
        overflow: visible;
      }
      
      * {
        box-sizing: border-box;
      }
      
      a {
        color: #4078c0;
        text-decoration: none;
      }
      
      a:hover,
      a:active {
        text-decoration: underline;
      }
      
      hr {
        height: 0;
        margin: 15px 0;
        overflow: hidden;
        background: transparent;
        border: 0;
        border-bottom: 1px solid #ddd;
      }
      
      hr::before {
        display: table;
        content: "";
      }
      
      hr::after {
        display: table;
        clear: both;
        content: "";
      }
      
      h1,
      h2,
      h3,
      h4,
      h5,
      h6 {
        margin-top: 0;
        margin-bottom: 0;
        line-height: 1.5;
      }
      
      h1 {
        font-size: 30px;
      }
      
      h2 {
        font-size: 21px;
      }
      
      h3 {
        font-size: 16px;
      }
      
      h4 {
        font-size: 14px;
      }
      
      h5 {
        font-size: 12px;
      }
      
      h6 {
        font-size: 11px;
      }
      
      p {
        margin-top: 0;
        margin-bottom: 5px;
      }
      
      blockquote {
        margin: 0;
      }
      
      ul {
        list-style: circle inside;
      }
      
      ul,
      ol {
        padding-left: 2em;
        margin-top: 0;
        margin-bottom: 0;
      }
      
      code {
        font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
        font-size: 12px;
      }
      
      body:before {
        display: table;
        content: "";
      }
      
      body:after {
        display: table;
        clear: both;
        content: "";
      }
      
      body>*:first-child {
        margin-top: 0 !important;
      }
      
      body>*:last-child {
        margin-bottom: 0 !important;
      }
      
      a:not([href]) {
        color: inherit;
        text-decoration: none;
      }
      
      h1,
      h2,
      h3,
      h4,
      h5,
      h6 {
        margin-top: 1em;
        margin-bottom: 16px;
        font-weight: bold;
        line-height: 1.4;
      }
      
      h1 {
        padding-bottom: 0.3em;
        font-size: 2.25em;
        line-height: 1.2;
        border-bottom: 1px solid #eee;
      }
      
      h1 {
        line-height: 1;
      }
      
      h2 {
        /* padding-bottom: 0.3em; */
        font-size: 1.75em;
        line-height: 1.225;
        /* border-bottom: 1px solid #eee; */
      }
      
      h2 {
        line-height: 1;
      }
      
      h3 {
        font-size: 1.5em;
        line-height: 1.43;
      }
      
      h3 {
        line-height: 1.2;
      }
      
      h4 {
        font-size: 1.25em;
      }
      
      h4 {
        line-height: 1.2;
      }
      
      h5 {
        font-size: 1em;
      }
      
      h5 {
        line-height: 1.1;
      }
      
      h6 {
        font-size: 1em;
        color: #777;
      }
      
      h6 {
        line-height: 1.1;
      }
      
      blockquote {
        margin-top: 0;
        margin-bottom: 16px;
      }
      
      hr {
        height: 4px;
        padding: 0;
        margin: 16px 0;
        background-color: #e7e7e7;
        border: 0 none;
      }
      
      blockquote {
        padding: 0 15px;
        color: #777;
        border-left: 4px solid #ddd;
        font-style: italic;
      }
      
      blockquote>:first-child {
        margin-top: 0;
      }
      
      blockquote>:last-child {
        margin-bottom: 0;
      }
      
      code {
        padding: 0;
        padding-top: 0.2em;
        padding-bottom: 0.2em;
        margin: 0;
        font-size: 85%;
        background-color: rgba(0,0,0,0.04);
        border-radius: 3px;
      }
      
      code:before,
      code:after {
        letter-spacing: -0.2em;
        content: "\00a0";
      }
      
      hr {
        border-bottom-color: #eee;
      }"""
  
  def getCss(): String = {
    return css
  }
}