package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type cliOptions struct {
	IncludeQsValues bool
	FilterQs        bool
}

var options cliOptions

func main() {
	flag.BoolVar(&options.IncludeQsValues, "qv", false, "If enabled, include query string values (e.g. if enabled /?q=123 - 123 would be included in results")
	flag.BoolVar(&options.FilterQs, "fq", false, "If enabled, filter out query strings (e.g. if enabled /?q=123 - q would NOT be included in results")
	flag.Parse()

	allComponents := make(map[string]bool)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		host := strings.ToLower(sc.Text())
		components := getUrlComponents(host)
		for _, component := range components {
			if allComponents[component] {
				continue
			}
			allComponents[component] = true
			fmt.Println(component)
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read input: %s\n", err)
	}
}

func getUrlComponents(host string) []string {
	var components []string
	if !strings.HasPrefix(host, "/") && !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}

	u, err := url.Parse(host)

	// If URL can't be parsed, ignore and move on
	if err != nil {
		return components
	}

	path := u.Path
	pathFragments := strings.Split(path, "/")

	// Remove first item from the slice as it will be blank
	if len(pathFragments) > 0 {
		pathFragments = pathFragments[1:]
	}

	// If query strings can't be parsed, set query strings as empty
	queryStrings, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		queryStrings = nil
	}

	domain := u.Host
	domainFragments := strings.Split(domain, ".")

	// Remove last item from the slice as it will be the extension (.com, .net, .etc)
	if len(domainFragments) > 0 {
		domainFragments = domainFragments[:len(domainFragments)-1]
	}

	for _, fragment := range domainFragments {
		components = append(components, fragment)
	}

	for _, pathFragment := range pathFragments {
		components = append(components, pathFragment)
	}

	if queryStrings != nil {
		for qs, values := range queryStrings {
			if !options.FilterQs {
				components = append(components, qs)
			}
			if options.IncludeQsValues {
				for _, val := range values {
					components = append(components, val)
				}
			}
		}
	}

	return components
}
