package student

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//Student ..
type Student struct {
	Name  string `yaml:"Name"`
	Class string `yaml:"Class"`
	Age   int    `yaml:"Age"`
	HTML  string `yaml:"html"`
}

// Init initializes student
func (student *Student) Init() {
	file, err := os.Open("student.yaml")
	if err != nil {
		log.Fatalf("Error in Init opening file : %v", err)
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(student); err != nil {
		log.Fatalf("Error in Init decoding : %v", err)
	}
}
func (student *Student) Write() {
	err := os.Remove("student.yaml")
	file, err := os.OpenFile("student.yaml", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("Error in Init opening file : %v", err)
	}
	defer file.Close()
	d := yaml.NewEncoder(file)
	if err := d.Encode(student); err != nil {
		log.Fatalf("Error in Init decoding : %v", err)
	}
}
