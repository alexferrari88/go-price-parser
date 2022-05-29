
# Go Price Parser

Don't you find it frustrating that whenever you are trying to parse price from scraped data, you can never just use one simple regex and be done with it?

Wouldn't be great, right?

Every website (and country) deciding to now use a comma, now a dot, and why not a space?!

Well, I was fed up with it and made this little library that just does one thing: **you feed it a string with a price, and it gives you back a nicely formatted int**.

Why an int, you might wonder?

Well, that's easy to deal with and not get crazy with all those float shenanigans.
## Features 🚀

- Parses a "dirty" price string into an integer (e.g. "42,069.23 USD" => 4206923)
- Returns values as either directly integer or as Price struct (allows some extra formatting features)

Install it as usual with:

```bash
  go get github.com/alexferrari88/go-price-parser
```
    
## Usage/Examples

You can have the return values as either int or as the Price struct defined in this library

### Return value as int
This is handy if you want to pipe the parsed price in a library such as [Money](https://github.com/Rhymond/go-money).
```go
scrapedPrice := "42.99 USD"
price, err := priceParser.IntFromString(scrapedPrice)
if err != nill {
    panic(err)
}

fmt.Println(price) // Output: 4299
```

### Return value as Price struct
You can use this if you want to keep the currency or for some formatting (see below).
```go
scrapedPrice := "42.99 USD"
price, err := priceParser.PriceFromString(scrapedPrice)
if err != nill {
    panic(err)
}

fmt.Println(price.Amount) // Output: 4299
```

### Other handy functions in the package
If you keep the result of the function as Price struct, you can also use the functions String() and Float().
#### Examples
```go
scrapedPrice := "42.99 USD"
price, err := priceParser.PriceFromString(scrapedPrice)
if err != nill {
    panic(err)
}

fmt.Println(price.Amount.String()) // Output: "42.99 USD"
fmt.Println(price.Amount.Float()) // Output: "42.99"
```



## To Do
- [ ]  Implement currency detection
- [ ]  Add more documentation to the code
- [ ]  Create tests
- [ ]  Optimize code
- [ ]  Improve README.md
## Acknowledgements

 - [hayj/SystemTools](https://github.com/hayj/SystemTools) for the inspiration for this package
## License

[MIT](https://choosealicense.com/licenses/mit/)

