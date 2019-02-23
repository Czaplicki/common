package oogl

type Uniform int32

//Unform is just a "shell" for all the openGl Uniform (set) functions
//To give them a more Object-oriented feel

//makeing the functions "members" of a Uniform varible

//// TODO: seperate the std set funcs form set array funs (for readability)
//// TODO: desisde if singel set matrix is a thing, and if , paramenter format

//format chanages :

	//old :
	// 1	: glUniform<n><t>( location, ... )
	// 2	: glUniform<n><t>v( location, count, *values )
	// 3	: glUniformMatrix<n(xn2)><t>( location, count, *values )


	//new :
	//1		: uniformObj.Set<Type><n>( ... )
	//2		: uniformObj.Set<Type><n>Array( count, *values)

	//extra	: uniformObj.SetMatrix<n(xn2)><t>( []values )
	//3		: uniformObj.SetMatrix<n(xn2)><t>Array( count, *values )
