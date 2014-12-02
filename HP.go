package main
import(
	"fmt"
	"time"
	"math"
	"math/rand"
	//"strings"
	//"code.google.com/p/draw2d/draw2d"
	"os"
	//"strconv"
	"errors"
)
// return a random sequence string of l,r,f
func randomFold(length int) string {
	var r int 
	var str string
	for i := 0; i < length; i++{
		rand.Seed(time.Now().UnixNano())
		r = rand.Intn(3)
		if r == 0{
			str += "l"
		}else if r == 1{
			str += "r"
		}else{//r = 2
			str += "f"
		}
	}
	return str
}
// return the fold onto a suffcient large 2D matrix
// start form (25,25), init "O" means not used
// odd--points: H/P ; even--edges: E:edges used
func DrawFold(random string, hp string) [][]string {
	var m [][] string = make([][]string, 70)
	for i := 0; i < 70; i++{
		m[i] = make([]string, 70)
	}
	for i := 0; i < 70; i++{
		for j := 0; j < 70; j++{
			m[i][j] = "O"
		}
	}
	x := 35
	y := 35
	t := math.Pi/2
	d := math.Pi/2
	if hp[0] == 'H'{
		m[x][y] = "H"
	}else{//'P'
		m[x][y] = "P"
	}
	for i := 0; i < len(random); i++{
		if random[i] == 'l'{
			t +=d
		}else if random[i] == 'r'{
			t -=d
		}
		y += int(math.Cos(t))
		x += -int(math.Sin(t))
		m[x][y] = "E" //line
		y += int(math.Cos(t))
		x += -int(math.Sin(t)) 
		if hp[i+1] == 'H'{
			m[x][y] = "H"
		}else{//'P'
			m[x][y] = "P"
		}
	}
 	return m 
}
// draw the fold on a canvas and save to png
func PaintFold(random string, hp string){
	pic := CreateNewCanvas(1000, 1000)
	pic.SetLineWidth(1)
	pic.SetStrokeColor(MakeColor(0, 0, 0))
	x := 500.0
	y := 500.0
	pic.gc.ArcTo(x, y, 5, 5, 0, 2*math.Pi)
	t := math.Pi/2
	d := math.Pi/2
	if hp[0] == 'H'{
		pic.Fill()
	}else{//== "P"
		pic.Stroke()
	}
	pic.MoveTo(x,y)
	for i := 0; i < len(random); i++{
		if random[i] == 'l'{
			t += d
		}else if random[i] == 'r'{
			t -= d
		}
		x += 50*math.Cos(t)
		y += -50*math.Sin(t)//be careful as coordinate-y in go is up to down
		pic.LineTo(x, y)
		pic.Stroke()
		if hp[i+1] == 'H'{
			pic.gc.ArcTo(x, y, 5, 5, 0, 2*math.Pi)
			pic.Fill()
			pic.MoveTo(x,y)
		}else{//== "P"
			pic.gc.ArcTo(x, y, 5, 5, 0, 2*math.Pi)
			pic.Stroke()
			pic.MoveTo(x,y)
		}
	}
	pic.SaveToPNG("fold.png")
}
// return energy(S,p) = 10x - \sum p*s,
// x: time crossed, p = 1 if P, p = 0 if H
// s: neighbour points in the walk and not adjacent
// if crossed structure, return true
func energy(random string, hp string) (bool, int){
	m := DrawFold(random, hp)
	cross := false
	count := 0
	s := 0
	for i := 2; i < len(m)-2; i++{
		for j:= 2; j < len(m[0])-2; j++{
			if m[i][j] == "H" || m[i][j] == "P" {
				count++// count total num of points to calculate cross times
				if m[i][j] == "H"{// 'P' is 0, no influence, only 'H' matters
					// diagonal existence must result in s++
					if m[i-2][j-2] == "H" || m[i-2][j-2]== "P"{
						s++
					}
					if m[i-2][j+2] == "H" || m[i-2][j+2]== "P"{
						s++
					}
					if m[i+2][j-2] == "H" || m[i+2][j-2]== "P"{
						s++
					}
					if m[i+2][j+2] == "H" || m[i+2][j+2]== "P"{
						s++
					}
					// not adjacent in the structure
					// && precedent ||
					if (m[i-2][j] == "H" || m[i-2][j]== "P") && m[i-1][j] == "O"{
						s++
					}
					if (m[i][j-2] == "H" || m[i][j-2]== "P") && m[i][j-1] == "O"{
						s++
					}	
					if (m[i][j+2] == "H" || m[i][j+2]== "P") && m[i][j+1] == "O"{
						s++
					}	
					if (m[i+2][j] == "H" || m[i+2][j]== "P") && m[i+1][j] == "O"{
						s++
					}	
					//fmt.Println("find a H")
					//fmt.Println(s)															
				}
			}
		}
	}
	if count != 0{
		cross = true
	}
	x := len(hp) - count
	energy := 10 * x - s
	//fmt.Println(x)
	//fmt.Println(energy)
	return cross, energy
}
//  Takes a fold and randomly changes one of its com- mands
func RandomFoldChange(str string) string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(str))//index to change
	s := rand.Intn(2)// two choices over change
	// fmt.Println(r)
	// fmt.Println(s)
	// fmt.Println(str)
	
	if str[r] == 'l'{
		if s == 0{
			str = str[0:r] + "r" + str[r+1:]
		}else {
			str = str[0:r] + "f" + str[r+1:]
		}
	}
	if str[r] == 'r'{
		if s == 0{
			str = str[0:r] + "l" + str[r+1:]
		}else {
			str = str[0:r] + "f" + str[r+1:]
		}
	}
	if str[r] == 'f'{
		if s == 0{
			str = str[0:r] + "l" + str[r+1:]
		}else {
			str = str[0:r] + "r" + str[r+1:]
		}
	}
	//fmt.Println(str)
	return str
}
// use simulate anneal to find local optimal enery fold structure
// lower energy results in fold change a command and higher energy 
// results in fold change with a probabilty (math.e(-delta/kT))
func OptimizeFold(hp string) (string, int){
	m := 100000//iteration times
	T := 1.0//temperature
	k := 0.998//parameter
	q := 0.0
	p := 0.0
	random := randomFold(len(hp)-1)
	_, energyOri := energy(random, hp)
	for i := 0; i < m; i++{
		randChange := RandomFoldChange(random)
		cross, energyChange := energy(randChange, hp)
		if cross {// reject the crossed ones
			continue
		} else {
			delta := energyChange - energyOri
			if delta < 0{
				random = randChange
				energyOri = energyChange
			}else{
				q = math.Exp(-float64(delta)/(k*T))
				//fmt.Println(100*q)
				rand.Seed(time.Now().UnixNano())
				p = float64(rand.Intn(100))
				if(p < 100*q){
					random = randChange
					energyOri = energyChange				
				}
			}
			if i%100 == 0{
				T = 0.999*T// run faster
			}
		}
		
	} 
	//fmt.Println(energyOri)
	return random, energyOri
}
func main(){
	//x := randomFold()
	//fmt.Println(x)
	//a := "HHPHPHHPHHPPHHH"
	//b := "llrfrffrffrrfl"
	
	//PaintFold(x,a)
	// m := DrawFold(b, a)
	// for i := 0; i < len(m); i++{
	// 	for j := 0; j < len(m[0]); j++{
	// 		fmt.Print(m[i][j])
	// 	}
	// 	fmt.Println()
	// }
	//PaintFold(b,a)

	
	//fmt.Println(energy(b,a))
	//RandomFoldChange(x)
	//fmt.Println("##########################")
	//s := OptimizeFold(a)
	//fmt.Println(s)
	//PaintFold(s,a)
	if(len(os.Args) != 2){
        err := errors.New("Error: sorry, number of your input is not OK")
        fmt.Println(err)
        return
    }
    hp := os.Args[1]
    opt, ener := OptimizeFold(hp)
    PaintFold(opt, hp)
    fmt.Println("Energy: " ,ener)
    fmt.Println("Structure: " ,opt)

}