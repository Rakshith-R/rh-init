package student

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//Student ..
type Student struct {
	Name  string `yaml:"Name"`
	Class string `yaml:"Class"`
	Age   int    `yaml:"Age"`
}

// Init initializes student
func (student *Student) Init() {
	student.Write()
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

//GetHTML returns html
func (student *Student) GetHTML() string {
	HTML := `<body>
	<h2>Name : %s</h2>
	<h2>Class : %s</h2>
	<h2>Age : %d</h2>
	<br/>
	<form method="post" action ="/">
	<label for="Name">Name:</label><br>
	<input type="text" id="Name" name="Name" value="%s"><br>
	<label for="Class">Class:</label><br>
	<input type="text" id="Class" name="Class" value="%s" ><br>
	<label for="Age">Age:</label><br>
	<input type="number" id="Age" name="Age" value=%d><br>
	<input type="submit" value="Submit">
	</form>
	</body>`
	return fmt.Sprintf(HTML, student.Name, student.Class, student.Age, student.Name, student.Class, student.Age)
}
