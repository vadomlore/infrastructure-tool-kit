package opshelper

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type EnvironmentOptions struct {
	Name      string `yaml:"name" json:"name"`
	ValueFrom struct {
		SecretKeyRef struct {
			Name string `yaml:"name" json:"name"`
			Key  string `yaml:"key" json:"key"`
		} `yaml:"secretKeyRef,omitempty" json:"secretKeyRef,omitempty"`
		FieldRef struct {
			FieldPath string `yaml:"fieldPath" json:"fieldPath"`
		} `yaml:"fieldRef,omitempty" json:"fieldRef,omitempty"`
	} `yaml:"valueFrom,omitempty" json:"valueFrom,omitempty"`
	Value string `yaml:"value,omitempty" json:"value,omitempty"`
}
type Environment struct {
	Env []EnvironmentOptions `yaml:"env" json:"env"`
}

func WriteEnv(values []EnvironmentOptions, src string, dest string) {
	b, error := ioutil.ReadFile(src)
	if error != nil {
		log.Fatalf("fail to read file %v\n", error)
	}
	env := Environment{}
	yaml.Unmarshal(b, &env)

	env.Env = append(env.Env, values...)

	d, err := yaml.Marshal(&env)

	if err != nil {
		log.Fatalf("error:%v", err)
	}
	ioutil.WriteFile(dest, d, 0644)
}

func WriteTFVariable(variables []Variable, src string, dest string) {
	b, error := ioutil.ReadFile(src)
	if error != nil {
		log.Fatalf("fail to read file %v\n", error)
	}
	for _, v := range variables {
		b = append(b, []byte("\n"+v.String()+"\n")...)
	}
	ioutil.WriteFile(dest, b, 0644)
}

type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	B     bool   `json:"b"`
}

type DeploymentVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (v Variable) String() string {
	if v.B {
		return fmt.Sprintf("variable %s { default = %v }", v.Name, v.Value)
	}
	return fmt.Sprintf("variable %s {}", v.Name)
}

func DeploymentVarsString(vars []DeploymentVar) string {
	s := ""
	for _, v := range vars {
		s += fmt.Sprintf("\t\t%v = %v", v.Name, v.Value)
		s += "\n"
	}
	return s
}
