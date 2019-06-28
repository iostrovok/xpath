package way

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var ComputesOperator string = "|"

var Operators []string = []string{
	"+", "-", "*", "div", "mod",
	"!=", ">", "=", "<", ">=", "<=",
	"and", "or",
}

var re, reDigital, reBool *regexp.Regexp

func init() {
	re = regexp.MustCompile(`^([^\[+]+)\[(.+)\]$`)
	reDigital = regexp.MustCompile(`^[0-9]+$`)

	reBool = regexp.MustCompile(`^[-_a-zA-Z0-9]+(\<|\>)[-_a-zA-Z0-9]+$`)

	reBool = regexp.MustCompile(`^[-_a-zA-Z0-9]+(\<|\>)[-_a-zA-Z0-9]+$`)

}

type RuleFunc func(int, interface{}) bool

type Way struct {
	Paths     []*Step `json:"Paths"`
	LastI     int     `json:"LastI"`
	I         int     `json:"I"`
	searchAny bool    `json:"searchAny"`
}

type Step struct {
	Index []int  `json:"Index"`
	Path  string `json:"Path"`
}

func (s *Step) Clone() *Step {
	return &Step{
		Index: s.Index,
		Path:  s.Path,
	}
}

func (w *Way) Dump() string {

	b, err := json.Marshal(w)
	if err != nil {
		log.Fatal("error:", err)
	}

	return string(b)
}

func extractRule(path, key string) (RuleFunc, bool) {
	if !reDigital.MatchString(key) {
		return nil, false
	}

	out := func(i int, m interface{}) bool {
		return false
	}

	return out, true
}

func extractID(path, key string) (*Step, bool) {
	if !reDigital.MatchString(key) {
		return nil, false
	}

	id, err := strconv.Atoi(key)
	if err != nil {
		return nil, false
	}

	return &Step{
		Index: []int{id},
		Path:  path,
	}, true
}

func getStep(path string) (*Step, error) {

	parts := re.FindStringSubmatch(path)
	if len(parts) == 0 {
		return &Step{
			Index: []int{-1},
			Path:  path,
		}, nil
	}

	path = parts[1]
	if reDigital.MatchString(parts[2]) {

		id, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}

		return &Step{
			Index: []int{id},
			Path:  path,
		}, nil
	}

	return &Step{
		Index: []int{-1},
		Path:  path,
	}, nil
}

func New(path string) ([]*Way, error) {

	searchAny := false
	if strings.HasPrefix(path, "//") {
		searchAny = true
	}

	path = strings.TrimPrefix(path, "/")
	paths := strings.Split(path, "/")

	out := &Way{
		searchAny: searchAny,
		Paths:     make([]*Step, len(paths), len(paths)),
		LastI:     len(paths) - 1,
		I:         -1,
	}

	var err error
	for i, p := range paths {
		out.Paths[i], err = getStep(p)
		if err != nil {
			return nil, err
		}
	}

	return []*Way{out}, nil
}

func (w *Way) NextBy(i int) (string, bool) {

	if i > w.LastI {
		return "", false
	}

	return w.Paths[i].Path, true
}

func (w *Way) ArrayRuleBy(i int) (RuleFunc, bool) {
	return nil, false
}

func (w *Way) ArrayIndextBy(i int) ([]int, bool) {
	if i > w.LastI || w.Paths[i].Index[0] == -1 {
		return []int{}, false
	}

	return w.Paths[i].Index, true
}
