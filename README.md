# jubilant-octo-parakeet

1. Fork this repository, and then clone the fork.
2. Add a `credentials.txt` with your Polygon APIKey and Secret in the following format
```
{
  "apiKey" : "",
  "secret" : ""
}
```
4. Create a new problem on polygon and copy its problem id
5. Run `./createProblem problemId` to create a folder for the problem with skeleton files.
6. Fill in all the details inside the files. Use Markdown for statements and tutorials
7. Once done, open `changes.txt` and add the problem id's of all problems which you wish to update on polygon
8. Run `go build main.go && ./main` to update on Polygon
9. Commit changes on Polygon

Currently unsupported - 
1. Adding a custom checker
2. Interactive problems
3. Adding a validator
4. Adding manual tests
5. Setting a test to use as an example
6. Adding multiple solution files or changing the tags on solution files
7. Running invocations on solutions
8. Creating and downloading packages
