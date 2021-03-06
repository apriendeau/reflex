package main

import "regexp"

var defaultExcludes = []string{
	// VCS dirs
	`(^|/)\.git/`,
	`(^|/)\.hg/`,
	// Vim
	`~$`,
	`\.swp$`,
	// Emacs
	`\.#`,
	`(^|/)#.*#$`,
	// OS X
	`(^|/)\.DS_Store$`,
	// Go vendor
	`(^|/)\.vendor`,
	`(^|/)\vendor`,
	// node modules
	`(^|/)\node_modules`,
}

var defaultExcludeMatcher multiMatcher

func init() {
	for _, pattern := range defaultExcludes {
		m := newRegexMatcher(regexp.MustCompile(pattern), true)
		defaultExcludeMatcher = append(defaultExcludeMatcher, m)
	}
}
