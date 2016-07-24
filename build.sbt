lazy val root = (project in file("."))
libraryDependencies += "org.scala-lang" % "scala-swing" % "2.10+"
mainClass in (Compile, run) := Some("ui.gui")