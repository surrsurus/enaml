package trans

import file.MarkupFile
import file.CSS.getCss

object Translator {
  /// Method for translating specialized markup into html
  /// Take a markupfile object as an input, which can be initalized
  /// inside of main.scala that links to a user-supplied path
  def translate(file: MarkupFile): Int = {
    // Read from file
    var lines = file.read()
    
    // Look for enaml metadata
    if ( lines(0) != "[enaml]" ) {
      // Return bad exit code
      return 1
    }
    
    // Get title
    var title = lines(1).replace("[", "").replace("]", "")
    
    // Write out to a new html file starting with the essential
    // html data
    htmlHeadBoilerplate(file, title)
    
    file.write("\t\t<h1 style=\"margin: auto;text-align:center;\">" 
      + title + "</h1>\n</br>")
    
    // For each line in the file that was read from,
    // parse each line and write the translated lines to the
    // new file
    for ( i <- 0 to (lines.length - 1) ) {
      var line = lines(i)
      file.write(parse(line))
    }
    
    // End html file
    htmlTailBoilerplate(file)
    
    // Close up the file
    file.close()
    
    // All is good, return positive exit code
    return 0
    
  }
  
  /// Turn text lines into markup based on certain rules
  private def parse(line: String): String = {
    var ignoreLine = false
    var startOfLine = true
    
    // Start of line elements
    var headingLevel = 0
    var isQuote = false
    var isBulleted = false
    
    // Inline elements
    var boldOpen = false
    var italOpen = false
    var undrOpen = false
    var linkOpen = false
    var codeOpen = false
    
    // If the line is blank, insert a spacer
    if ( line == "" ) {
      return "\t\t</br>"
    }
    
    // Division metadata
    if(  line == "[div]") {
      return "\t\t<hr>"
    }
    
    // Split line into characters
    var lineVector = line.split("").toVector
    
    // Link metadata
    if ( line.contains("[link") && lineVector(0) == "[" ) {
      var args = line.split(' ')
      if ( args.length != 3 ) {
        return "\t\t<p><b>Error: Link metadata has improper syntax</b></p>"
      } else {
        if ( args(2).takeRight(1) == "]" ) {
          args(2) = args(2).dropRight(1)
        }
        return "\t\t<a href=" + args(2) + ">" + args(1) + "</a>" 
      }
    }
    
    // Image metadata
    if ( line.contains("[img") && lineVector(0) == "[" ) {
      var args = line.split(' ')
      if( args.length != 2 ) {
        return "\t\t<p><b>Error: Image metadata has improper syntax</b></p>"
      } else {
        if ( args(1).takeRight(1) == "]" ) {
          args(1) = args(1).dropRight(1)
        }
        return "\t\t<img src=" + args(1) + ">" 
      }
    }
    
    // Where the new line (+ html data) will be stored
    var newline = ""
    
    // Iterate through each character
    for( char <- lineVector ) {
      
      if ( !ignoreLine ) {
        
        // Check for start of line syntax first
        if ( "[#*>".contains(char) && startOfLine ) {
          char match {
            case "[" => ignoreLine = true // Comment
            case "#" => headingLevel += 1 // Heading
            case "*" => isBulleted = true // Bulleted list
            case ">" => isQuote = true    // Quote
          }
        // Otherwise check for inline syntax
        } else {
          startOfLine = false
          if ( "%@_`" contains char ) {
            // Each element has a match case to open and close the tag
            if( char == "%") { // Bold
              if ( boldOpen ) newline += "</b>"
              else newline += "<b>"
              boldOpen = !boldOpen
            } else if( char == "@" ) { // Italics
              if ( italOpen ) newline += "</i>"
              else newline += "<i>"
              italOpen = !italOpen
            } else if( char == "_") { // Underlined
              if ( undrOpen ) newline += "</u>"
              else newline += "<u>"
              undrOpen = !undrOpen
            } else if( char == "`") { // Code
              if ( codeOpen ) newline += "</code>"
              else newline += "<code>"
              codeOpen = !codeOpen
            }
          // Otherwise, it must be a regular character
          } else {
            newline += char 
          }
        }
      // Ignored lines don't show up
      } else {
        return ""
      }
      
    }
    
    // Remove trailing and proceeding whitespace
    newline = newline.trim
    
    // Insert tags
    if ( headingLevel > 0 ) {
      // 6 is the smallest heading size
      if( headingLevel > 6 ) {
        headingLevel = 6
      }
      var h = "h" + headingLevel.toString
      return "\t\t<" + h + ">" + newline + "</" + h + ">"
    } else if ( isBulleted ) {
      return "\t\t<ul><li>" + newline + "</li></ul>"
    } else if ( isQuote ) {
      return "\t\t<blockquote>" + newline + "</blockquote>"
    } else {
      return "\t\t<p>" + newline + "</p>"
    }

  }
  
  /// Ugly ugly code up ahead...
  /// It's just html data really, nothing special.
  /// Q: Could this be better?
  /// A: Probably, but this way works as intended, as spagehttirific as it seems
  
  /// Write the top of the html file
  /// (Start html & body tags and also include the head tag metadata)
  private def htmlHeadBoilerplate(file: MarkupFile, title: String): Unit = {
    file.write("""<html lang="en">
    <head>
        <meta charset="utf-8">""")
    file.write("        <title>" + title + "</title>")
    file.write("				<style>" + getCss + "</style>")
    file.write("""        <meta name="description" content="Translated Markup">
        <meta name="author" content="Me">
        <link rel="stylesheet" href="css/styles.css?v=1.0">
    </head>

    <body>
        <div class='enaml'>""")
  }
  
  /// Write the bottom of the html file
  /// (close body and html tags)
  private def htmlTailBoilerplate(file: MarkupFile): Unit = {
    file.write("""    </div>
      </body>
</html>""")
  }
  
}