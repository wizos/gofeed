package main

import (
	"log"
	"regexp"
)

func RegexpFilter(filterReg *regexp.Regexp, data []byte) (outdata []byte) {
	if nil == filterReg {
		return data
	}

	matches := filterReg.FindAllSubmatch(data, -1)
	if nil == matches {
		log.Printf("[ERROR] failed to match filter regex, pattern %s did not match", filterReg.String())
		if *gDebug {
			log.Println("======= debug: target data =======")
			log.Println(string(data))
			log.Println("==============")
		}
		return nil
	}

	for _, match := range matches {
		for patInd, patName := range filterReg.SubexpNames() {
			switch patName {
			case PATTERN_FILTER:
				outdata = append(outdata, match[patInd]...)
			}
		}
	}

	if *gDebug {
		log.Println("======= debug: filter regex ========")
		log.Println(filterReg.String())
		log.Println("======= debug: filtered data =======")
		log.Println(string(outdata))
		log.Println("==============")
	}
	return
}
