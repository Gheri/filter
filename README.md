This is filter console app in Go/GoLang.
It takes json file and returns the ids of records with the empty/null/missing
and identical values.
Logic to identify if two records are identical or not, does not consider id value.
If all fields are identical except id, still records are considered identical.

Prerequisite

VSCode >= version 1.55.2 or any other IDE/editor
golang >= 1.16.3

How to run app??

Step 1: Navigate to root directory
Step 2: Install all dependencies
Step 3: run -> "go build . "
Step 4: After step2 binary would be created in root directory
run the binary -> "./filter" (Note you can set the app name while building)

How to run test??

run -> go test

To process another file
Copy file to root directory
Set env variable FILTER_FILE_NAME="<filename>"
run the binary -> "./filter"  (Note you can set the app name while building)

Example:  
// json file  
[  
    {  
      "name": "Viennia Sturm",  
      "address": "",  
      "zip": "17565",  
      "id": "ea0c4"  
    },  
    {  
      "name": null,  
      "address": "Boston Street",  
      "zip": "17565",  
      "id": "ea0c5"  
    },  
    {  
      "name": "Amerah Lang",  
      "address": "5037 Providence Bouled",  
      "zip": "44109",  
      "id": "8d322"  
    },  
    {  
      "name": "Wendolyn Sweat",  
      "address": "1521 Gem Avenue",  
      "zip": "77701",  
      "id": "a6b3b"  
    },  
    {  
      "name": "Wendolyn Sweat",  
      "address": "1521 Gem Avenue",  
      "zip": "77701",  
      "id": "b6b3b"  
    },  
]  


The output of above file would be printed on console  

ea0c4  
ea0c5  
a6b3b  
b6b3b  
