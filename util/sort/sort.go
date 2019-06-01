package sort

import (
	insight "github.com/davidalpert/ginsight/insight"
)

type ByObjectSchemaKey []insight.ObjectSchema

func (s ByObjectSchemaKey) Len() int           { return len(s) }
func (s ByObjectSchemaKey) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByObjectSchemaKey) Less(i, j int) bool { return s[i].Key < s[j].Key }

type ByObjectTypeName []insight.ObjectType

func (s ByObjectTypeName) Len() int           { return len(s) }
func (s ByObjectTypeName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByObjectTypeName) Less(i, j int) bool { return s[i].Name < s[j].Name }

type ByObjectTypeAttributeName []insight.ObjectTypeAttribute

func (s ByObjectTypeAttributeName) Len() int           { return len(s) }
func (s ByObjectTypeAttributeName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByObjectTypeAttributeName) Less(i, j int) bool { return s[i].Name < s[j].Name }
