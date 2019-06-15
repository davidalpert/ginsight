package sort

import (
	api "github.com/davidalpert/ginsight/api"
)

type ByObjectSchemaKey []api.ObjectSchema

func (s ByObjectSchemaKey) Len() int           { return len(s) }
func (s ByObjectSchemaKey) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByObjectSchemaKey) Less(i, j int) bool { return s[i].Key < s[j].Key }

type ByObjectTypeName []api.ObjectType

func (s ByObjectTypeName) Len() int           { return len(s) }
func (s ByObjectTypeName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByObjectTypeName) Less(i, j int) bool { return s[i].Name < s[j].Name }

type ByObjectIconID []api.ObjectIcon

func (s ByObjectIconID) Len() int           { return len(s) }
func (s ByObjectIconID) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByObjectIconID) Less(i, j int) bool { return s[i].ID < s[j].ID }

type ByObjectIconName []api.ObjectIcon

func (s ByObjectIconName) Len() int           { return len(s) }
func (s ByObjectIconName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByObjectIconName) Less(i, j int) bool { return s[i].Name < s[j].Name }

type ByObjectTypeAttributeName []api.ObjectTypeAttribute

func (s ByObjectTypeAttributeName) Len() int           { return len(s) }
func (s ByObjectTypeAttributeName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByObjectTypeAttributeName) Less(i, j int) bool { return s[i].Name < s[j].Name }

type ByObjectTypeAttributeID []api.ObjectTypeAttribute

func (s ByObjectTypeAttributeID) Len() int           { return len(s) }
func (s ByObjectTypeAttributeID) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByObjectTypeAttributeID) Less(i, j int) bool { return s[i].ID < s[j].ID }
