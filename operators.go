package main
import("fmt")

func main(){
//wehave the following operators
//Arithemetic
//Assignment 
//Comparism
//Logical
//Bitwise

//Arithemetic Operators
/*
+   Addition
-   Subtraction
/   Division
*   Multiplication
%   Modules
++  Increment
--  Decreement
*/

//MAKING USE OF THE NECCERARY ONE
//modulus
	x:= 10
	y:= 3
	fmt.Println(x % y)




//ASSIGNMENT OPERATORS
//This is used to assign value to a variable

/*
We have the following Assignment Operators

=	Equals
+=	Another way of adding
-=	Another way of subtracting
%=	Modulus
/=	Division
=&	??
|=	??
<<=	??
>>=	??

*/
//COMPARISM OPERATORS

/*
We have th following Comparism Operators
==	Equal to
!=	Not equal to
>	Greater than
<	Less than
<=	Less than or equal to
>=	Greater than or equal to
*/

		A :=10
		B := 5
		fmt.Println(A != B)    //OUTPUT SHOULD BE TRUE BECAUSE THEY ARE NOT EQUAL
//LOGICAL OPERATORS
/*
They ar the following:
&&	Returns true if both statement are true
||	Returns true if one of the statement is true
!`	Reverse the result,returns false if the statement is true
*/


//BITWISE OPERATORS
/*
Bitwise operators are used in binary numbers
&	Setseach bit to 1 if both bit are 0ne
|	Sets each bit to 1 if one or two bit is 1
<<	Moves left By pushing zero in from the right
>	Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off 
*/



	var g = 9
	fmt.Printf("g= %b\n", g)
	fmt.Printf("x >>2 is %b\n", x >>2)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


//GO CONDIITIONS
// This can only be true or false. GO HAS THE FOLLOWILL CONDITION STATEMENT
// IF 
//ELSE
//IF ELSE
//SWITCH

//IF STATEMENT

	if 20>10{
	fmt.Println("hey am greater than you bro")
	}


//IF .. ELSE STATEMENT

	time := 12
	if time < 13 {
	fmt.Println("Good afternoon")
	}else{
	fmt.Println("Good Moening bby")
	}



//ELSE IF STATEMENT
//Use the else if statement to specify a new condition if the first condition is false


	timing := 20

	if timing < 11{
	fmt.Println("Hey good morning")}else if timing >21{
	fmt.Println("Good night")}else{
	fmt.Println("Hmm are you sure") //out put should be this because none met the condition except the else.
	}


/*NOTE: IN A SITUATION WHERE IF CONDITION AND IFELSE CONDITION ARE TRUE THE COMPILER WILL DISPALAY THE FIRST STATEMENT WHIH IS 
THE IF STATEMENT CONDITION*/

// NESTED IF STATEMENT
//WHEN YOU HAVE IF STATEMENT INSIDE AN IF STATEMENT IS WHEN YOU HAVE NESTED IF STATEMENT

age := 40

if age > 12{
fmt.Println("You are a small boy")
if age < 10{
fmt.Println("You are a baby")
}else{
fmt.Println("70 years old man")
}

}

}
