This application simulates a shopping cart structure
where you can create carts with different articles. It supports
the CRUD methods. 
The data from the articles list are provided from the API http://challenge.getsandbox.com/articles.
They are stored in a slice of structs, and are manipulated as an index of the slice, so here is the 
convention map to know the number for each article:

    0-Banana
    1-Apple
    3-Cookies
    4-Noodles
    5-Olive Oil
    6-Water
    7-Beer
    8-Vodka
    9-Bread
    10-Grapes
    11-Rice
    12-Pizza

Data is input manually in JSON format directly in Postman, although there are 2 carts created in
the main.go file, to start the application. Each one with the same set of articles
