package sap

import (
	"log"
	"strings"
	"testing"
)

var source = []string{"enterprise", "ibox-edu", "ibox-ent", "other"}
var prefixNumber = []string{"98", "99", "18"}

func TestConditions(t *testing.T) {
	log.Println(ConditionV1("99"))
	log.Println(ConditionV2("989"))
}

func ConditionV1(parameter string) bool {
	if source[0] == parameter || source[1] == parameter || source[2] == parameter || prefixNumber[0] == parameter || prefixNumber[1] == parameter {
		return true
	}

	return false
}

func ConditionV2(parameter string) bool {
	if IsSetForChannelThirty(parameter) {
		return true
	}

	return false
}

func IsSetForChannelThirty(parameter string) bool {
	switch strings.ToLower(parameter) {
	case "enterprise", "ibox-edu", "ibox-ent", "98", "99", "18": // source
		return true
	default:
		return false
	}
}
