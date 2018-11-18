package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Embed", func() {
	BeforeEach(func() {
		templatePath = ""
	})

	AfterEach(func() {
		os.RemoveAll("./testdata/cfg.go")
	})

	It("sanity", func() {
		os.Args = []string{"app", "-source=./testdata/cfg.yaml", "-target=./testdata/cfg.go", "-package=testpackage", "-name=Config"}
		main()
		actualContent, _ := ioutil.ReadFile("./testdata/cfg.go")
		expectedContent, _ := ioutil.ReadFile("./testdata/goldenCfg.go")
		Ω(removeSpecialSymbols(actualContent)).Should(Equal(removeSpecialSymbols(expectedContent)))
	})

	It("negative", func() {
		os.Args = []string{"app", "-source=./testdata/cfgNotExisting.yaml", "-target=./testdata/cfg.go", "-package=testpackage", "-name=Config"}
		defer func() {
			Ω(recover()).ShouldNot(BeNil())
		}()
		main()
	})
})

func removeSpecialSymbols(b []byte) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s := string(b)
	s = strings.Replace(s, "0xd, ", "", -1)
	s = reg.ReplaceAllString(s, "")
	return s
}
