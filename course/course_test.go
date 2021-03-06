package course

import (
	"io/ioutil"
	"testing"
)

func TestFull(t *testing.T) {
	c := Course{
		Capacity:   10,
		Enrollment: 10,
	}

	if !c.IsFull() {
		t.Fail()
	}
}

func TestBelowCapactiy(t *testing.T) {
	c := Course{
		Capacity:   50,
		Enrollment: 25,
	}

	if c.IsFull() {
		t.Fail()
	}
}

func TestAboveCapacity(t *testing.T) {
	c := Course{
		Capacity:   100,
		Enrollment: 101,
	}

	if !c.IsFull() {
		t.Fail()
	}
}

func TestPullFromNextCell(t *testing.T) {
	inpt := `
	<html><head>
	<tr valign=top>
	 <td colspan=2 bgcolor=#99CCFF><b><br>
		<font size=+1>Spring 2022 English CC1010 section 001</font><br>
		<font size=+2>UNIVERSITY WRITING</font><br>
		<i>UW: CONTEMPORARY ESSAYS</i><br>
		<br></b></td></tr>
	<tr valign=top><td bgcolor=#99CCFF>Call Number</td>
	 <td bgcolor=#DADADA>15243</td></tr>
	<tr valign=top><td bgcolor=#99CCFF>Day &amp; Time<br>Location</td>
	 <td bgcolor=#DADADA>MW 8:40am-9:55am<br>201D Philosophy Hall</td></tr>
	<tr valign=top><td bgcolor=#99CCFF>Points</td>
	 <td bgcolor=#DADADA>3</td></tr>
	`

	outpt := pullFromNextCell(inpt, "Call Number")

	if (outpt != "15243") {
		t.Fail()
	}
}

func TestParseIntAtSimple(t *testing.T) {
	out := parseIntAt("12345", 0)

	if out != 12345 {
		t.Fail()
	}
}

func TestParseIntNoNumber(t *testing.T) {
	out := parseIntAt("Hello World", 3)

	if out != 0 {
		t.Fail()
	}
}

func TestParseIntDecimal(t *testing.T) {
	out := parseIntAt("Hello 351341.432 World", 6)

	if out != 351341 {
		t.Fail()
	}
}

func TestParseIntSingleDigit(t *testing.T) {
	out := parseIntAt("No. 1 !!!", 4)

	if out != 1 {
		t.Fail()
	}
}

func TestParseEnrollmentDetails(t *testing.T) {
	enrollment, capacity := parseEnrollmentDetails(" 	14 students (18 max)")

	if enrollment != 14 || capacity != 18 {
		t.Fail()
	}
}


func TestStripFrom(t *testing.T) {
	bytes, err := ioutil.ReadFile("_testdata/example.txt")

	if err != nil {
		t.Fail()
	}

	input := string(bytes)

	actualCourse, err := StripFrom(input)

	if err != nil {
		t.Fail()
	}

	expectedCourse := Course{
		Professor: "Ali Yalgin",
		CallNum: 15243,
		Capacity: 14,
		Enrollment: 10,
	}

	if actualCourse != expectedCourse {
		t.Fail()
	}
}

func TestStringer(t *testing.T) {
	myCourse := Course{
		Professor: "John Cena",
		CallNum: 69420,
		Capacity: 50,
		Enrollment: 45,
	}

	expected := "Call #: 69420\nProfessor: John Cena\nCapacity: 50\nEnrollment: 45\nFull?: false"

	if myCourse.String() != expected {
		t.Fail()
	}
}
