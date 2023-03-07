module github.com/LvffY/ansible

go 1.17

replace (
	// These are replace directives copied from porter
	// When you use a newer version of Porter, if you run into trouble with go mod tidy
	// Copy any additional replace directives from Porter's go.mod file
	// They must match the replaces used by porter everything to compile
	github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
	github.com/spf13/viper => github.com/getporter/viper v1.7.1-porter.2.0.20210514172839-3ea827168363
)

require (
	get.porter.sh/magefiles v0.3.1
	get.porter.sh/porter v1.0.0
	github.com/ghodss/yaml v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.1
	github.com/xeipuuv/gojsonschema v1.2.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/PaesslerAG/gval v1.0.0 // indirect
	github.com/PaesslerAG/jsonpath v0.1.1 // indirect
	github.com/PuerkitoBio/goquery v1.5.0 // indirect
	github.com/andybalholm/brotli v1.0.0 // indirect
	github.com/andybalholm/cascadia v1.0.0 // indirect
	github.com/carolynvs/aferox v0.3.0 // indirect
	github.com/carolynvs/magex v0.8.0 // indirect
	github.com/cbroglie/mustache v1.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/gobuffalo/logger v1.0.3 // indirect
	github.com/gobuffalo/packd v1.0.0 // indirect
	github.com/gobuffalo/packr/v2 v2.8.0 // indirect
	github.com/golang/snappy v0.0.4-0.20210608040537-544b4180ac70 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/karrick/godirwalk v1.15.3 // indirect
	github.com/klauspost/compress v1.15.1 // indirect
	github.com/klauspost/pgzip v1.2.4 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/magefile/mage v1.11.0 // indirect
	github.com/markbates/errx v1.1.0 // indirect
	github.com/markbates/oncer v1.0.0 // indirect
	github.com/markbates/safe v1.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mholt/archiver/v3 v3.5.0 // indirect
	github.com/mmcdole/gofeed v1.0.0-beta2 // indirect
	github.com/mmcdole/goxpp v0.0.0-20181012175147-0068e33feabf // indirect
	github.com/nwaples/rardecode v1.1.0 // indirect
	github.com/olekukonko/tablewriter v0.0.4 // indirect
	github.com/pierrec/lz4/v4 v4.0.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ulikunitz/xz v0.5.7 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/term v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
