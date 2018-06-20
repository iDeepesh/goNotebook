package txt

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

type popularSport struct {
	Country string
	Name    string
}

//GetNumberOfPlayersInTeam - Returns number of players in cricket team
func (s popularSport) GetNumberOfPlayersInTeam() int {
	if s.Name == "Cricket" {
		return 11
	}
	return 10
}

//GetTotalPlayers - Returns total number of players
func (s popularSport) GetTotalPlayers() int {
	return s.GetNumberOfPlayersInTeam() * 2
}

//RenderOneTemplate - Reads hello.gohtml and renders the same
func RenderOneTemplate() {
	fmt.Println("Inside txt.RenderOneTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderOneTemplate method for template examples")

	temp, err := template.ParseFiles("txt/tmpl/helloWorld.gohtml")
	if err != nil {
		panic(err)
	}

	temp.Execute(os.Stdout, nil)
}

//RenderMultiTemplate - Reads multiple templates and renders the same
func RenderMultiTemplate() {
	fmt.Println("Inside txt.RenderMultiTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderMultiTemplate method for template examples")

	temp, err := template.ParseFiles("txt/tmpl/helloWorld.gohtml", "txt/tmpl/helloWorldAgain.gohtml")
	if err != nil {
		panic(err)
	}

	temp.ExecuteTemplate(os.Stdout, "helloWorld.gohtml", nil)
	temp.ExecuteTemplate(os.Stdout, "helloWorldAgain.gohtml", nil)
}

//RenderMultiTemplateRegex - Reads all templates in a folder and renders the same
func RenderMultiTemplateRegex() {
	fmt.Println("Inside txt.RenderMultiTemplateRegex method for template examples")
	defer fmt.Println("Finished executing txt.RenderMultiTemplateRegex method for template examples")

	temp, err := template.ParseGlob("txt/tmpl/helloWorl*.gohtml")
	if err != nil {
		panic(err)
	}

	temp.ExecuteTemplate(os.Stdout, "helloWorld.gohtml", nil)
	temp.ExecuteTemplate(os.Stdout, "helloWorldAgain.gohtml", nil)
}

//RenderTimeTemplate - Reads all templates in a folder and renders the same
func RenderTimeTemplate() {
	fmt.Println("Inside txt.RenderTimeTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderTimeTemplate method for template examples")

	temp, err := template.ParseFiles("txt/tmpl/helloTime.gohtml")
	if err != nil {
		panic(err)
	}

	temp.ExecuteTemplate(os.Stdout, "helloTime.gohtml", time.Now())
}

//RenderVariableTemplate - Reads all templates in a folder and renders the same
func RenderVariableTemplate() {
	fmt.Println("Inside txt.RenderVariableTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderVariableTemplate method for template examples")

	temp, err := template.ParseFiles("txt/tmpl/helloVariable.gohtml")
	if err != nil {
		panic(err)
	}

	temp.ExecuteTemplate(os.Stdout, "helloVariable.gohtml", time.Now())
}

//RenderSliceTemplate - Reads all templates in a folder and renders the same
func RenderSliceTemplate() {
	fmt.Println("Inside txt.RenderSliceTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderSliceTemplate method for template examples")

	temp, err := template.ParseFiles("txt/tmpl/helloSlice.gohtml")
	if err != nil {
		panic(err)
	}

	sports := []string{"Cricket", "Soccer", "Ping Pong", "Baseball"}

	temp.ExecuteTemplate(os.Stdout, "helloSlice.gohtml", sports)
}

//RenderMapTemplate - Reads all templates in a folder and renders the same
func RenderMapTemplate() {
	fmt.Println("Inside txt.RenderMapTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderMapTemplate method for template examples")

	temp, err := template.ParseFiles("txt/tmpl/helloMap.gohtml")
	if err != nil {
		panic(err)
	}

	sports := map[string]string{"India": "Cricket", "Brazil": "Soccer", "China": "Ping Pong", "USA": "Baseball"}

	temp.ExecuteTemplate(os.Stdout, "helloMap.gohtml", sports)
}

//RenderStructTemplate - Reads all templates in a folder and renders the same
func RenderStructTemplate() {
	fmt.Println("Inside txt.RenderStructTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderStructTemplate method for template examples")

	temp, err := template.ParseFiles("txt/tmpl/helloStruct.gohtml")
	if err != nil {
		panic(err)
	}

	sport := popularSport{"India", "Cricket"}

	temp.ExecuteTemplate(os.Stdout, "helloStruct.gohtml", sport)
}

//RenderStructMapTemplate - Reads all templates in a folder and renders the same
func RenderStructMapTemplate() {
	fmt.Println("Inside txt.RenderStructMapTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderStructMapTemplate method for template examples")

	temp, err := template.ParseFiles("txt/tmpl/helloStructMap.gohtml")
	if err != nil {
		panic(err)
	}

	sports := map[string]popularSport{"Sachin": {"India", "Cricket"}, "Pele": {"Brazil", "Soccer"}, "Xing": {"China", "Ping Pong"}, "Joe": {"USA", "Baseball"}}

	temp.ExecuteTemplate(os.Stdout, "helloStructMap.gohtml", sports)
}

//RenderFunctionalTemplate - Reads all templates in a folder and renders the same
func RenderFunctionalTemplate() {
	fmt.Println("Inside txt.RenderFunctionalTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderFunctionalTemplate method for template examples")

	fm := template.FuncMap{
		"uc": strings.ToUpper,
		"firstThree": func(s string) string {
			ss := strings.TrimSpace(s)
			return ss[:3]
		},
	}

	temp, err := template.New("").Funcs(fm).ParseFiles("txt/tmpl/helloFunctions.gohtml")
	if err != nil {
		panic(err)
	}

	sports := map[string]popularSport{"Sachin": {"India", "Cricket"}, "Pele": {"Brazil", "Soccer"}, "Xing": {"China", "Ping Pong"}, "Joe": {"USA", "Baseball"}}

	temp.ExecuteTemplate(os.Stdout, "helloFunctions.gohtml", sports)
}

//RenderPipelineTemplate - Reads all templates in a folder and renders the same
func RenderPipelineTemplate() {
	fmt.Println("Inside txt.RenderPipelineTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderPipelineTemplate method for template examples")

	fm := template.FuncMap{
		"fMMDDYY": func(t time.Time) string {
			return t.Format("01/02/2006")
		},
	}

	temp := template.Must(template.New("").Funcs(fm).ParseFiles("txt/tmpl/helloPipeline.gohtml"))

	temp.ExecuteTemplate(os.Stdout, "helloPipeline.gohtml", time.Now())
}

//RenderNestedTemplate - Reads all templates in a folder and renders the same
func RenderNestedTemplate() {
	fmt.Println("Inside txt.RenderNestedTemplate method for template examples")
	defer fmt.Println("Finished executing txt.RenderNestedTemplate method for template examples")

	temp, err := template.ParseFiles("txt/tmpl/helloNesting.gohtml", "txt/tmpl/helloNested.gohtml")
	if err != nil {
		panic(err)
	}

	temp.ExecuteTemplate(os.Stdout, "helloNesting.gohtml", "Arya")
}
