# Arrays

* Arrays are the most common datastructure that one come across in various programming languages.
* Arrays store homogeneous data types.
* In Go, len() returns the length of the array.
* Arrays can be traversed using a traditional for loop OR
* Arrays can be traversed by using the `range` keyword.
* While using `range`, both `index` & `value` can be returned.
* In Go, arrays have a fixed size.

# Slices

* Slices in Go can be thought of as dynamic arrays.
* Using len() & cap() functions we get the length and capacity of slices.
* `_` is the blank identifier, when you don't want to use some returned values make use of underscore.
* Difference between len() & cap()
*	len() tells you how many elements are currently stored in the slice.
*	cap() tells you the maximum number of elements the slice can hold without reallocation.

# Maps

* Maps in GO can be thought of as JSON objects or Hashtables that have a key and an associated value.

# Database Operations

* For any database operation, what you need? what are the basics?
    - For instance, if you want to withdraw money from the bank account (NO INTERNET BANKING FACILITY IS AVAILABLE), you must know the BANK_ADDRESS, ACC_NO, ACC_SIGNATURE.
    - Similary to go ahead and query the database you need a few things, those are:
    
        -  BANK_ADDRESS => DATABASE_NAME
        - ACC_NO        => DATABASE_USER
        - ACC_SIGNATURE => DATABASE_PASSWORD

    - Using the above we get a database connection object, bingo! once you have a connection object you can go ahead and start querying the databe.

    `THIS IS THE APPLICABLE TO ALL PROGRAMMING LANGUAGES WHEN IT COMES TO QUERYING DATABASE`

# Web forms

- The different concepts that were learnt in database operations are now being called from behing an http route.
- `net/http` package in Go is responsible for defining routes which becomes the API endpoints and those can be exposed.


# What I have learned from Chapter 2?

- An introduction to Array, Slice and Maps, IMO these are the most basic datastructure that one must know in order to proceed with learning Go at a deeper level.
- The difference between array and slice is

    - Array is fixe size
    - Slice is dynamic
- Maps can be imagined as JSON object having key value pairs or in terms of PHP an associative array's equivalent is Maps.
- Database operations introduces to defining database connections, the parameters that are required for establishing the db connection.
- Once the connection is made the ways in which the db can be queried and the ways in which insert, update, delete happens.
- For queuering we use `db.Query`, and for inserting, updating & deleting we use `db.Prepare`, where Prepare symbolises prepared statement.

As a learned having an understanding of arrays, slices, maps & the learning objective of database operations will help in the coming chapter to tackle difficult problems.

With this let us venture ahead....
And be a programmer of Value....
Be a programmer of value not just a name bearing programmer....

@iambatman....