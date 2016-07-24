package ui

import scala.swing._
import scala.swing.event._
import file.MarkupFile
import trans.Translator.translate
import sun.swing.FilePane.FileChooserUIAccessor

object gui extends SimpleSwingApplication {
  
  val searchField = new TextField { columns = 32; }
  searchField.text = "No file selected";
  val translateButton = new Button("Translate");
  translateButton.enabled = false;
  val browseButton = new Button("Browse");
  val results = new TextArea("Translate with enaml!");
  
  def top = new MainFrame {
    
    title = "Translate Markup";
    val searchLine = new BoxPanel(Orientation.Horizontal) {
      contents += searchField;
      contents += Swing.HStrut(20);
      contents += translateButton;
      contents += Swing.HStrut(20);
      contents += browseButton;
    }
      
    // make sure that resizing only changes the resultField:
    restrictHeight(searchLine);
    
    contents = new BoxPanel(Orientation.Vertical) {
      contents += searchLine;
      contents += Swing.VStrut(10);
      border = Swing.EmptyBorder(10, 10, 10, 10);
      contents += results;
    }
  
    listenTo(translateButton);
    listenTo(browseButton);
    reactions += {
      case EditDone(`searchField`) => enableTranslate()
      case ButtonClicked(`translateButton`) => translateNew()
      case ButtonClicked(`browseButton`) => browseFiles()
    }
    
  }
  
  private def enableTranslate(): Unit = {
    if( searchField.text != "" ) {
      translateButton.enabled = true;
    }
  }
    
  private def translateNew(): Unit = {
    results.text = "Translating...";
    var file = new MarkupFile(searchField.text);
    var err = translate(file);
    if( err == 1 ) {
      results.text = "Error! File is not enaml data!";
    } else if( err == 0 ) {
      results.text = "Translated!";
    }
    
  }
  
  private def browseFiles(): Unit = {
    val fc = new FileChooser;
    fc.title = "Choose a file to be translated";
    val result = fc.showOpenDialog(null)
    
    if (result == FileChooser.Result.Approve) {
      translateButton.enabled = true;
      searchField.text = fc.selectedFile.getAbsolutePath;
    } else {
      translateButton.enabled = false;
      results.text = "Open command cancelled by user.";
    }
  }
  
  private def restrictHeight(s: Component) {
    s.maximumSize = new Dimension(Short.MaxValue, s.preferredSize.height);
  }

}