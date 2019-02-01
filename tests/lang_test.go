package test

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/AplaProject/apla-lang"
)

type contract struct {
	Source string
	Line   int    // the start line of the contract
	Gas    int64  // expecting gas
	Result string // expecting result
}

func (cnt *contract) check(gas int64, get string) error {
	if get != cnt.Result {
		return fmt.Errorf("Line %d: get != want;\n%s !=\n%s", cnt.Line, get, cnt.Result)
	}
	if cnt.Gas > 0 && gas != cnt.Gas {
		return fmt.Errorf("Line %d: got gas %d != %d", cnt.Line, gas, cnt.Gas)
	}
	return nil
}

func (cnt *contract) checkError(err error) error {
	if err.Error() != cnt.Result {
		return fmt.Errorf("Line %d: get != want;\n%s !=\n%s", cnt.Line, err.Error(), cnt.Result)
	}
	return nil
}

func loadTest(filename string) (ret []*contract, err error) {
	var (
		input []byte
		start int
	)
	input, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	ret = make([]*contract, 0, 1000)

	list := strings.Split(string(input), "\n")
	source := make([]string, 0, 32)
	re := regexp.MustCompile(`====\s*(\d+)\s*:(.*)`)

	for i, line := range list {
		if strings.HasPrefix(line, `====`) {
			var (
				gas    int64
				result string
			)
			if re.MatchString(line) {
				match := re.FindSubmatch([]byte(line))
				if gas, err = strconv.ParseInt(strings.TrimSpace(string(match[1])), 10, 64); err != nil {
					return
				}
				result = string(match[2])
			} else {
				result = line[4:]
			}
			ret = append(ret, &contract{
				Source: strings.Join(source, "\r\n"),
				Line:   start + 1,
				Gas:    gas,
				Result: strings.Replace(strings.TrimSpace(result), `\n`, "\n", -1),
			})
			source = source[:0]
			start = i + 1
		} else {
			source = append(source, line)
		}
	}

	return
}

func TestLang(t *testing.T) {
	vm := simvolio.NewVM()
	contracts, err := loadTest(`default_test`)
	if err != nil {
		t.Error(err)
		return
	}
	for i := int64(len(contracts)) - 1; i >= 0; i-- {
		cnt := contracts[i]
		if err = vm.LoadContract(cnt.Source, i); err != nil {
			if err = cnt.checkError(err); err != nil {
				t.Error(err)
				return
			}
		}
		result, gas, err := vm.Run(vm.Contracts[len(vm.Contracts)-1])
		if err = cnt.check(gas, result); err != nil {
			t.Error(err)
			return
		}
	}
	t.Error(`OK`, vm)
}
