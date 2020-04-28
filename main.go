package main

import (
	"encoding/json"
	"github.com/vadomlore/programatic-go-tool/utility"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//Append configuration
type Configuration struct {
	searchpath string
	overwrite  bool                   // set to true will overwrite the original file, or else create a file with -new in the same folder.
	data       map[string]interface{} //values to append
}

func LoadConfig(searchpath, templatepath, overwrite string) Configuration {

	config := Configuration{
		searchpath: searchpath,
		overwrite:  strings.Contains(strings.ToLower(overwrite), "true"),
		data:       make(map[string]interface{}),
	}
	//templatePath := "D:\\vadomlore-go-utility\\template"
	templatePath := templatepath

	files, err := ioutil.ReadDir(templatePath)
	if err != nil {
		log.Fatalf("error load template file %v\n", err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		bytes, _ := ioutil.ReadFile(path.Join(templatePath, f.Name()))
		data := FileTemplate{}
		json.Unmarshal(bytes, &data)
		switch data.File {
		case "deployment.tf":
			{
				v := DeploymentJson{}
				json.Unmarshal(bytes, &v)
				config.data[data.File] = v
			}
		case "environment.yaml":
			{
				v := EnvironmentJson{}
				json.Unmarshal(bytes, &v)
				config.data[data.File] = v
			}
		case "variables.tf":
			{
				v := VariableJson{}
				json.Unmarshal(bytes, &v)
				config.data[data.File] = v
			}
		}
	}
	return config
}

type FileTemplate struct {
	File string `json:"file"`
	Data string `json:"data"`
}

type DeploymentJson struct {
	File string                    `json:"file"`
	Data []opshelper.DeploymentVar `json:"data"`
}

type EnvironmentJson struct {
	File string                         `json:"file"`
	Data []opshelper.EnvironmentOptions `json:"data"`
}

type VariableJson struct {
	File string               `json:"file"`
	Data []opshelper.Variable `json:"data"`
}

func main() {
	if len(os.Args) < 3 {
		log.Println("*****************AppendSkyWalkingLog*************")
		log.Println("Usage: main.go searchPath templatePath overwrite")
		log.Println("Function: append the settings in the json file into ")
		log.Println("the terraform file settings ")
		log.Println("e.g main.go /usr/local/searchpath ./template false")
		log.Println("**************************************************")
		os.Exit(1)
	}
	searchPath, templatePath, overwrite := os.Args[1], os.Args[2], os.Args[3]
	//"D:\\zSiemensWork\\发布配置\\edgeapppublishingmanager", "false"

	config := LoadConfig(searchPath, templatePath, overwrite)
	AppendSkyWalkingLog(config, funcMap)
	log.Println("done")

}

var funcMap = map[string]func(searchpath string, configuration Configuration, overwrite bool){
	"environment.yaml": WriteYamlEnvironment,
	"deployment.tf":    WriteTFDeployment,
	"variables.tf":     WriteTFVariables,
}

func AppendSkyWalkingLog(configuration Configuration, funcs map[string]func(file string, configuration Configuration, overwrite bool)) {
	filepath.Walk(configuration.searchpath, func(fullfilepath string, f os.FileInfo, err error) error {
		for matchFile, fn := range funcs {
			if strings.HasSuffix(f.Name(), matchFile) {
				fn(fullfilepath, configuration, configuration.overwrite)
			}
		}
		return nil
	})
}

func WriteYamlEnvironment(filename string, configuration Configuration, overwrite bool) {
	baseFileName, newPath := newFileToWrite(filename, overwrite)
	log.Printf("update file %v\n", filename)
	opshelper.WriteEnv(configuration.data[baseFileName].(EnvironmentJson).Data, filename, newPath)
}

func WriteTFVariables(filename string, configuration Configuration, overwrite bool) {

	baseFileName, newPath := newFileToWrite(filename, overwrite)
	log.Printf("update file %v\n", filename)
	opshelper.WriteTFVariable(configuration.data[baseFileName].(VariableJson).Data, filename, newPath)
}

func WriteTFDeployment(filename string, configuration Configuration, overwrite bool) {
	baseFileName, newFile := newFileToWrite(filename, overwrite)
	log.Printf("update file %v\n", filename)
	opshelper.InsertStringToFileEndWith(filename, newFile, opshelper.DeploymentVarsString(configuration.data[baseFileName].(DeploymentJson).Data), "}", 2)
}

func newFileToWrite(filename string, overwrite bool) (string, string) {
	baseFileName := filepath.Base(filename)
	basePath := filepath.Dir(filename)
	ext := path.Ext(filename)
	newPath := filename
	if !overwrite {
		newPath = path.Join(basePath, strings.Replace(baseFileName, ext, "-temp"+ext, 1))
	}
	return baseFileName, newPath
}
