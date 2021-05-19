# youtubenhl-go

My first Golang project via: EQuimper on Youtube.
I used this project to learn about:
1. Goroutines
2. Writing to files in Go 
3. Fetching data from an external API 
4. WaitGroup 
5. Channels 

Have 32 teams and want to fetch the stats from each user of their roster. 
Use 1 API call to recieve all the teams,
 Then 32 API calls to recieve the rest of each team 
 These API calls took 231.1625ms
 Concurrency allowed for all the calls to be made in one shot all together, instead of 32 API calls 1-by-1.  
 
 STEPS
 1. Got static API from statsapi.web.nhl.com --> translated it on JSON-to-Go to get structs 
 2. Opened GoLand --> NewProject --> Renamed project --> updated GOROOT to go1.16.4 --> Create
  | ISSUE: Tried to use, Visual Studio Code but had trouble connecting all the parts needed from Go. Ended up making some Hello World type stuff, but it was getting
  too convoluted.  SOLUTION: Downloaded and set up GoLand environment instead. 
 3. Created a new Go file main.go under the project's directory
 4. Created a new Directory nhlApi under the project's directory and a new file under nhlApi, nhlApi.go
 5. Fetched the API and stored it in, Project name right click --> New Directory nhlApi ---> new File ---> nhlApi.go 
 6. Created a file to manage the teams of the nhlApi, called teams.go and added the struct
    Deleted copyright, changed type to Team, deleted Teams struct from top and botton, and deleted the extra Venue structs 
    | ISSUE: Error after deleting extra vendors. 
     SOLUTION: Searched for syntax errors that may not have come up in Errors. missing ` and extra }
 7. Created a nhlTeamsResponse struct, that is going to be a slice of teams, in teams.go
 8. Create a form that is an exportable function to recive all the teams, in teams.go 
    Created a sting, that is going to be teams that we got from the URL in nhlAPi.go, catch errors, and close the body of the function
    Decoder needs ioRouter interface  to decode the slice of teams gotten from the nhlTeamsResponse struct, and error handle
    Reeturn the slice of teams gotten from the nhlTeamsResponse, and if there was an error
 9. Go to main.go to ask and fetch for api and create main function
 10. Using imported time package to benchmark the request time and format the request time
 11. Open a file using os package, error handle using program termination and close the file
 12. Use io package to write duplicate byte streams and catch errors. Returns standard output and the file we opened, which we set the log output to
 13. Get the slice of Team from teams.go and check for error 
 14. Range over each of the team that we recieved and print the team.Name
     | ISSUE: Rutime error, team.Name undefined (type nhlApi.Team has no field or method name) 
      SOLUTION: Went back to teams.go and re-did my steps from Step 6. Struct was still called Teams []struct {, changed to type Team struct { 
 15. Go to main() and right click --> run go build 
 
 
