# IPLocationFinder

I created a simple service that uses a 3rd party api to read and parse the db file. It then checks the input ip against the database to determine the country and then checks it against a csv that represents the whitelisted countries. If your ip is in the list you will get a response

The IP for {countryname} is whitelisted.

Otherwise you get a response

The IP for {country} is NOT on the whitelist.

I tested by sending a get request to localhost:8080/{ip}
223.255.128.0 not whitelisted
81.2.69.142 whitelisted
