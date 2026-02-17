package main
import "fmt"

//len is said to be the exact array a slice contains. it is the visible value available
//cap  is the amount of value your slice can hold
// go duplicate values in other for other incoming values to enter. 
// look at the code below carefully before to find out whats going on 
func backingAr(){

myslice:= []int{}

myslice= append(myslice,1)
fmt.Printf("first capacity is %d\n",cap(myslice))
fmt.Printf("The 1st  length is %d\n", len(myslice))
fmt.Printf("The 1st value is %v\n", myslice)
//this is the first apend to the empty slice. here first value is appended 
//so the exact len is 1 and the cap is 1.



myslice= append(myslice,2)
fmt.Printf("my second capacity is %d\n",
 cap(myslice))
fmt.Printf("The 2nd length is %d\n", len(myslice))
fmt.Printf("The 2nd value is %v\n", myslice)
// Here the second append comes in. go has already create 1 space which has len 1 and cap len.
// since new value is coming in it will create another space for the value. so it has to duplicate the first value. 
//since it duplicated the first value. it will have a cap of 2 len of 2. 
//Note: The capacity is full here so for the next append it need a space where it will be inserted to.


myslice= append(myslice,3)
fmt.Printf("my capacity third is %d\n",
 cap(myslice))

fmt.Printf("The 3rd length is %d\n", len(myslice))
fmt.Printf("The 3rd  value is %v\n", myslice)
//HERE STARTS THE PROBLEM YOU MAY FACE  BECAUSE MAY QUESTION YOURSELF WHY cap is 4 and len is 3.
//Recall this . len is the exact amount of value present while cap is the amount of value it can hold .
//since we created another value and the slice is already full it needs to create space. 
//so it duplicates the value availabe which is "1,2", so now we have total of 4. 
// Here the thing will  include the third append ,remember it has already created 4 space
// so the  len which is the exact amount is 3 while the cap which is the amount of value it can hold is 4



myslice= append(myslice,4)
fmt.Printf("my fouth capacity is %d\n",
 cap(myslice))
fmt.Printf("The 4th length is %d\n", len(myslice))
fmt.Printf("The 4th value is %v\n", myslice)
//Here the fourth apppend comes in .
//remember it has already created a space for 4. So the fourth append goes in directly filling the space completely again.
//so the len will be 4 and the cap also 4
// NOTE: for another append it need to create another space because it has been completely filled

myslice= append(myslice,5)
fmt.Printf("my second fifth is %d\n",
 cap(myslice))
fmt.Printf("The 5th length is %d\n", len(myslice))
fmt.Printf("The 5th value is %v\n", myslice)
//Since the space is completely filled it has to create another one
// So it duplicates the present value thereby having the total space of 8.
//So the append will insert  there by having the len of  5 and the cap of 8


myslice= append(myslice,6)
fmt.Printf("value %v\n",myslice)
fmt.Printf(" My capacity is %d\n:",cap(myslice))
fmt.Printf(" my length %d\n:",len(myslice))
// Another insert comes in since there is still space it will append directly
// so the len will be 6 and the cap is 8


}
func main(){

 backingAr()
}
