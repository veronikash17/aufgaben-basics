package rectangles
import "fmt"
// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	// TODO
	for i:=0;i<height;i++{
		for a:=0; a<width; a++{
			fmt.Print("#")
		}
		fmt.Println()
	}

}
