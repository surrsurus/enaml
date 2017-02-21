# enaml
**E**naml is **N**ot **A** **M**arkup **L**anguage

<img align="center" src="https://github.com/surrsurus/enaml/blob/master/media/logo.png" alt="enaml" width=250>


## What is enaml then?

Enaml is actually a markup translator, that uses a custom markup language detailed under `examples`. It features a basic gui, and the ability to turn specified enaml/markup files into html code to view in the browser. I did it so that I could have a language that allows the user to make very quick notes for whatever setting, and have them formatted nicely to view in a web browser.

## Note about getting this to work

I spent 6 hours figuring out how to get a jar to compile in case if the user doesn't have scala (which 99% of them won't, 100% if you count the fact no one looks at my repositories). Basically it was beyond a nightmare. Eclipse wasn't helping, Maven wouldn't work, and the only solution was to install a custom plugin for SBT called [SBT Assembly](https://github.com/sbt/sbt-assembly) (Mad shoutout) which magically made this thing work. If you have problems getting an executable jar, please check this out. I'm also hoping the jar runs on other systems. It Works on My Machine(tm) but I can't guarentee that for all systems. All in all, Scala is fun to program in, nightmarish to deploy. If you have any problems please raise an issue.

## Credit

A big thank you to:

  * [SBT Assembly](https://github.com/sbt/sbt-assembly) for literally making this work

  * [Github Markdown CSS](https://github.com/sindresorhus/github-markdown-css) because I like the markdown CSS (and I'm lazy)

  * And Scala Swing, which is packaged in the jar and in the lib folder just in case. I don't own the Swing library, It's not my intellectual property, and it isn't subject to the license I'm using (Unless if specified by the creators).
