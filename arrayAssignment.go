package main
import "fmt"


	func main(){
	var myint int =200
	var myfloat64 float64 = 300
	var mystring string = "Luckify"
	var mybool bool = true

	fmt.Printf("MY datatype: %T\n and my value: %v\n",myint,myint)

	fmt.Printf("My datatype:%T\n and Myvalue:%v\n ", myfloat64,myfloat64)

	fmt.Printf("My datatype:%T\n and myvalue is = %v\n", mystring,mystring)

	fmt.Printf("my datatype:%T\n and my value:%v\n", mybool,mybool)
// question 2



//No3

	ages := [5]int{1,2,3,4,5}
	var totalSum int = 0

	for i:=0; i < len(ages); i++{
	totalSum += ages[i]
	}

	count := len(ages)
	mean :=  totalSum / int(count)

	fmt.Println("the mean is=",mean)
	fmt.Println("total sum is =", totalSum)


//4 maximum number

	ClassGrades := [5]int{30,40,50,20,10}

	maxNum:= ClassGrades[0]

	for i:= 1; i < len(ClassGrades); i++ {

	if ClassGrades[i] > maxNum{

	maxNum = ClassGrades[i]

	}
	}

	fmt.Printf("the maxnum is : %d\n", maxNum)



//parrallel array for data assiation

	names := [4]string{"Ada", "Obi","Emeka","Amaka"}
	scores := [4]int{10,20,30,40}

	for i:= 0;  i < len(names); i++ {
	studentName:= names[i]
	scoreValue:= scores[i]
	fmt.Printf("Name: %-8s | Score: %d\n", studentName, scoreValue )
	}

//no7

	var myInt int
	var myFloat float64
	var myString string
	var myBool bool

	var myNumbers [4]int

	fmt.Println("Zero Values:")
	fmt.Printf("Integer: [%d]\n", myInt)
	fmt.Printf("Float:   [%f]\n", myFloat)
	fmt.Printf("String:  [%s] (Empty quotes)\n", myString)
	fmt.Printf("Boolean: [%t]\n", myBool)
	fmt.Printf("Array:   %v\n", myNumbers)
//n06

box := [3]int{3,4,5}
_=box
secondbox := [3]int{1,2,3}

fmt.Printf("This is the value of box:%v\n",box)
box = secondbox
fmt.Printf("This is the value of box:%v\n",box)

}

