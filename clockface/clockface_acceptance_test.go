package clockface_test

import (
	"bytes"
	"encoding/xml"
	"lgwt/clockface"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(2000, time.January, 1, 0, 0, 30, 0, time.UTC)
	want := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSvgWriterSecondHand(t *testing.T) {

	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

// func TestSVGWriterMinuteHand(t *testing.T) {
// 	cases := []struct {
// 		time time.Time
// 		line Line
// 	}{
// 		{
// 			simpleTime(0, 0, 0),
// 			Line{150, 150, 150, 70},
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(testName(c.time), func(t *testing.T) {
// 			b := bytes.Buffer{}
// 			clockface.SVGWriter(&b, c.time)

// 			svg := SVG{}
// 			xml.Unmarshal(b.Bytes(), &svg)

// 			if !containsLine(c.line, svg.Line) {
// 				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
// 			}
// 		})
// 	}
// }

func containsLine(ln Line, lns []Line) bool {
	for _, line := range lns {
		if line == ln {
			return true
		}
	}

	return false
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(2000, 9, 20, h, m, s, 0, time.UTC)
}

func testName(time time.Time) string {
	return time.Format("15:34:11")
}
