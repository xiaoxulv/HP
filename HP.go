package main
import(
	"fmt"
	"time"
	"math"
	"math/rand"
	//"code.google.com/p/draw2d/draw2d"
	//"os"
	//"strconv"
	//"errors"
)
func randomFold() []string {
	var r int 
	random := make([]string, 0)
	for i := 0; i < 14; i++{
		rand.Seed(time.Now().UnixNano())
		r = rand.Intn(3)
		if r == 0{
			random = append(random,"l")
		}else if r == 1{
			random = append(random,"r")
		}else{//r = 2
			random = append(random,"f")
		}
	}
	return random
}
func drawFold(random string, hp string){
	pic := CreateNewCanvas(1000, 1000)
	pic.SetLineWidth(1)
	pic.SetStrokeColor(MakeColor(0, 0, 0))
	x := 500.0
	y := 500.0
	pic.ArcTo(x, y, 5, 5, 0, 2*math.Pi)
	t := math.Pi/2
	d := math.Pi/2
	if hp[0] == 'H'{
		pic.Fill()
	}else{//== "P"
		pic.Stroke()
	}
	pic.MoveTo(x,y)
	for i := 0; i < len(random); i++{
		//fmt.Println()
		if random[i] == 'l'{
			t += d
		}else if random[i] == 'r'{
			t -= d
		}
		x += 50*math.Cos(t)
		y += -50*math.Sin(t)//be careful as coordinate-y in go is up to down
		pic.LineTo(x, y)
		pic.Stroke()
		pic.MoveTo(x, y)
		if hp[i+1] == 'H'{
			pic.ArcTo(x, y, 5, 5, 0, 2*math.Pi)
			pic.Fill()
			pic.MoveTo(x,y)
		}else{//== "P"
			pic.ArcTo(x, y, 5, 5, 0, 2*math.Pi)
			pic.Stroke()
			pic.MoveTo(x,y)
		}
	}
    
    
    //pic.FillStroke()
	pic.SaveToPNG("HP.png")

}

func main(){
	//a := randomFold()
	//fmt.Println(a)
	a := "HHPHPHHPHHPPHHH"
	b := "llrfrffrffrrfl"
	//a := ["H","H","P","H","P","H","H","P","H","H","P","P","H","H","H"]
	//b := ["l","l","r","f","r","f","f","r","f","f","r","r","f","l"]
	drawFold(b,a)
}