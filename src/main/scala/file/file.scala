package file

import scala.io.Source
import java.io._

class MarkupFile(val path: String) {
  var inPath = path;
  var outPath = generate_outpath();
  val writer = new PrintWriter(new File(outPath));
  
  def close(): Unit = {
    writer.close();
  }
  
  def read(): Vector[String] = {
    return Source.fromFile(inPath).getLines.toVector;
  }
  
  def write(text: String): Unit = {
    if( text != "" ) {
      writer.write(text + "\n"); 
    }
  }
  
  private def generate_outpath(): String = {
    return inPath.split('.').toVector.apply(0) + ".html";
  }
  
}