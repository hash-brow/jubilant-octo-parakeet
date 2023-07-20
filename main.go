package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/variety-jones/polygon"
	"os"
)

func CreateApiObjectFromLocal() (api polygon.PolygonApi) {
	api = polygon.PolygonApi{}

	data, err := os.ReadFile("credentials.txt")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &api)
	if err != nil {
		panic(err)
	}

	return api
}

func GetContestProblems(api polygon.PolygonApi, contestId string) (problems []polygon.ProblemObject) {
	parameters := make(map[string]string)
	parameters["contestId"] = contestId

	problems, err := api.ContestProblems(parameters)
	fmt.Println(problems)

	if err != nil {
		fmt.Println(err)
	}

	return problems
}

func setGenInfo(api polygon.PolygonApi, fileName string) (err error) {
	// info.txt
	/*
		{
			"timeLimit" : "",
			"memoryLimit" : "",
			"interactive" : ""
		}
	*/

	wrapper := struct {
		TimeLimit   string
		MemoryLimit string
		Interactive string
	}{}

	data, err := os.ReadFile(fileName)
	err = json.Unmarshal(data, &wrapper)

	parameters := make(map[string]string)
	parameters["timeLimit"] = wrapper.TimeLimit
	parameters["memoryLimit"] = wrapper.MemoryLimit
	parameters["interactive"] = wrapper.Interactive

	err = api.ProblemUpdateInfo(parameters)
	return err
}

func setStatement(api polygon.PolygonApi, fileName string) (err error) {
	// statement.txt
	/*
		{
			"lang" : ,
			"name" : ,
			"legend" : ,
			"input" : ,
			"output" : ,
			"notes" : ,
		}
	*/

	wrapper := struct {
		Lang   string
		Name   string
		Legend string
		Input  string
		Output string
		Notes  string
	}{}

	data, err := os.ReadFile(fileName)
	err = json.Unmarshal(data, &wrapper)

	parameters := make(map[string]string)

	parameters["lang"] = wrapper.Lang
	parameters["name"] = wrapper.Name
	parameters["legend"] = wrapper.Legend
	parameters["input"] = wrapper.Input
	parameters["output"] = wrapper.Output
	parameters["notes"] = wrapper.Notes

	err = api.ProblemSaveStatement(parameters)

	return err
}

func setTutorial(api polygon.PolygonApi, fileName string) (err error) {
	// tutorial.txt
	/*
		{
			"tutorial" : ""
		}
	*/

	wrapper := struct {
		Tutorial string
	}{}
	data, err := os.ReadFile(fileName)
	err = json.Unmarshal(data, &wrapper)

	parameters := make(map[string]string)
	parameters["tutorial"] = wrapper.Tutorial

	err = api.ProblemSaveGeneralTutorial(parameters)
	return err
}

func setGenerator(api polygon.PolygonApi, fileName string) (err error) {
	// generator gen.cpp
	file, err := os.ReadFile(fileName)
	body := string(file)

	parameters := make(map[string]string)
	parameters["checkExisting"] = "false"
	parameters["name"] = "gen.cpp"
	parameters["type"] = "source"
	parameters["file"] = body

	err = api.ProblemSaveFile(parameters)

	return err
}

func setChecker(api polygon.PolygonApi, fileName string) (err error) {
	// checker.txt
	/*
		{
			"checker" : ""
		}
	*/
	/*
		Options
		std::fcmp.cpp - Lines, doesn't ignore whitespaces
		std::hcmp.cpp - Single huge integer
		std::lcmp.cpp - Lines, ignores whitespaces
		std::ncmp.cpp - Single or more int64, ignores whitespaces
		std::nyesno.cpp - Zero or more yes/no, case insensitive
		std::rcmp4.cpp - Single or more double, max error 1E-4
		std::rcmp6.cpp - Single or more double, max error 1E-6
		std::rcmp9.cpp - Single or more double, max error 1E-9
		std::wcmp.cpp - Sequence of tokens
		std::yesno.cpp - Single yes/no, case insensitive
	*/

	wrapper := struct {
		Checker string
	}{}

	data, err := os.ReadFile(fileName)
	err = json.Unmarshal(data, &wrapper)

	parameters := make(map[string]string)
	parameters["checker"] = wrapper.Checker

	err = api.ProblemSetChecker(parameters)

	return err
}

func setTestScript(api polygon.PolygonApi, fileName string) (err error) {
	// tests.txt
	/*
		{
			"testScript" : ""
		}
	*/

	wrapper := struct {
		TestScript string
	}{}

	data, err := os.ReadFile(fileName)
	err = json.Unmarshal(data, &wrapper)

	parameters := make(map[string]string)
	parameters["testset"] = "tests"
	parameters["source"] = wrapper.TestScript

	err = api.ProblemSaveScript(parameters)

	return err
}

func setSolution(api polygon.PolygonApi, fileName string) (err error) {
	file, err := os.ReadFile(fileName)
	body := string(file)

	parameters := make(map[string]string)
	parameters["name"] = "sol.cpp"
	parameters["file"] = body

	err = api.ProblemSaveSolution(parameters)

	return err
}

func makeProblem(api polygon.PolygonApi, problemId string) (err error) {
	err = setGenInfo(api, string("./"+problemId+"/info.txt"))
	if err != nil {
		fmt.Printf("Error while setting general info {%s}\n", err)
	}

	err = setStatement(api, string(problemId+"/statement.txt"))
	if err != nil {
		fmt.Printf("Error while setting statement {%s}\n", err)
	}

	err = setTutorial(api, string(problemId+"/tutorial.txt"))
	if err != nil {
		fmt.Printf("Error while setting tutorial {%s}\n", err)

	}
	err = setGenerator(api, string(problemId+"/gen.cpp"))
	if err != nil {
		fmt.Printf("Error while setting generator {%s}\n", err)
	}

	err = setChecker(api, string(problemId+"/checker.txt"))
	if err != nil {
		fmt.Printf("Error while setting checker {%s}\n", err)
	}

	err = setTestScript(api, string(problemId+"/tests.txt"))
	if err != nil {
		fmt.Printf("Error while setting test script {%s}\n", err)
	}

	// manual tests

	err = setSolution(api, string(problemId+"/sol.cpp"))
	if err != nil {
		fmt.Printf("Error while setting solution {%s}\n", err)
	}

	return nil
}

func main() {
	api := CreateApiObjectFromLocal()

	file, err := os.Open("changes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		problemId := scanner.Text()
		api.ProblemId = problemId
		err = makeProblem(api, problemId)

		if err != nil {
			fmt.Printf("Error handling problemId {%s} with error {%s}", problemId, err)
		}
	}
}
