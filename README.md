# cli
Command Line Interface (CLI) package

When testing ideas and concepts, it is useful to have an easy to use Command
Line Interface (CLI).  This is my attempt to create a simple one.

To use it, you need to build a table of Items (see the data structure Items).
The fields are:

OrderList:  This is a slice of strings which is the ordered list of the internal
names of the items to be displayed to the user.  The names are for internal
use and are not displayed.  They are indices into the ItemList table.  
Make them meaningful for you.

ItemList: Is a map of the Items to be displayed on the screen.  The fields for
the structure "Item" are listed below.  The index into the map is the names
that you entered into the OrderList slice.

ActionLines:  After the list of items to choose from are displayed, the contents
of the ActionLines slice are displayed one line per slice element.  This is where
you put the instructions to the user such as enter the number of the item You
want to change and the like.

Item: is the data structure that is repeated as many times as you have items
to modify.  The fields for the structure Item are:

Name:  This is the name of the item and it must be the same as the index into
the OrderList.

Prompt:  This is the line that is displayed next to a number on the screen.
It needs to tell the user the value that is to be changed.

Response:  This is the line that is displayed once the user enters a valid
response.

Value: is the default value of the field that is displayed to the user.

Validator:  This is a function that takes in a string and returns a bool.
You can supply any validation function you choose and if the function
does not return true, the user is asked to re-enter the field.

To use the package, you call the command function and pass the the pointer
to the Items table that you have built.  It returns a channel that will emit
valid user entered data.  The channel returns the structure ItemResponse which
has a name and a value field.  The name is the name of the data item and
the value is the value of the validated value of the data item.

There are two special names provided by the cli package.  One of them is continue
and the other is Quit.  They are sent when the user enters c or q.  How they are
used is a function of the using application.  The intent here is to signify
the users instruction is to run the function as defined by the cli values
entered or quit the program, but of course, they can be used in other ways.
