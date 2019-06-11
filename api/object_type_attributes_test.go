package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"

	api "github.com/davidalpert/ginsight/api"
)

var _ = Describe("Client", func() {
	Describe("ObjectTypeAttributes", func() {
		Describe("Get All For ObjectType", func() {
			endpoint := "/rest/insight/1.0/objecttype/57/attributes"

			Context("7 attributes of various types", func() {
				fixture := `[
  {
      "id": 448,
      "objectType": {
          "id": 58,
          "name": "Vlan",
          "type": 0,
          "icon": {
              "id": 85,
              "name": "Software Installer",
              "url16": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=16&inherited=false&abstract=false&time=1558951873000",
              "url48": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=48&inherited=false&abstract=false&time=1558951873000"
          },
          "position": 12,
          "created": "27/May/19 3:05 AM",
          "updated": "27/May/19 3:11 AM",
          "objectCount": 0,
          "objectSchemaId": 2,
          "inherited": false,
          "abstractObjectType": false,
          "parentObjectTypeInherited": false
      },
      "name": "Key",
      "label": false,
      "type": 0,
      "defaultType": {
          "id": 0,
          "name": "Text"
      },
      "editable": false,
      "system": true,
      "sortable": true,
      "summable": false,
      "minimumCardinality": 1,
      "maximumCardinality": 1,
      "removable": false,
      "hidden": false,
      "includeChildObjectTypes": false,
      "uniqueAttribute": false,
      "options": "",
      "position": 0
  },
  {
      "id": 449,
      "objectType": {
          "id": 58,
          "name": "Vlan",
          "type": 0,
          "icon": {
              "id": 85,
              "name": "Software Installer",
              "url16": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=16&inherited=false&abstract=false&time=1558951873000",
              "url48": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=48&inherited=false&abstract=false&time=1558951873000"
          },
          "position": 12,
          "created": "27/May/19 3:05 AM",
          "updated": "27/May/19 3:11 AM",
          "objectCount": 0,
          "objectSchemaId": 2,
          "inherited": false,
          "abstractObjectType": false,
          "parentObjectTypeInherited": false
      },
      "name": "Name",
      "label": true,
      "type": 0,
      "description": "The name of the object",
      "defaultType": {
          "id": 0,
          "name": "Text"
      },
      "editable": true,
      "system": false,
      "sortable": true,
      "summable": false,
      "minimumCardinality": 1,
      "maximumCardinality": 1,
      "removable": false,
      "hidden": false,
      "includeChildObjectTypes": false,
      "uniqueAttribute": false,
      "options": "",
      "position": 1
  },
  {
      "id": 450,
      "objectType": {
          "id": 58,
          "name": "Vlan",
          "type": 0,
          "icon": {
              "id": 85,
              "name": "Software Installer",
              "url16": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=16&inherited=false&abstract=false&time=1558951873000",
              "url48": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=48&inherited=false&abstract=false&time=1558951873000"
          },
          "position": 12,
          "created": "27/May/19 3:05 AM",
          "updated": "27/May/19 3:11 AM",
          "objectCount": 0,
          "objectSchemaId": 2,
          "inherited": false,
          "abstractObjectType": false,
          "parentObjectTypeInherited": false
      },
      "name": "Created",
      "label": false,
      "type": 0,
      "defaultType": {
          "id": 6,
          "name": "DateTime"
      },
      "editable": false,
      "system": true,
      "sortable": true,
      "summable": false,
      "minimumCardinality": 1,
      "maximumCardinality": 1,
      "removable": false,
      "hidden": false,
      "includeChildObjectTypes": false,
      "uniqueAttribute": false,
      "options": "",
      "position": 2
  },
  {
      "id": 451,
      "objectType": {
          "id": 58,
          "name": "Vlan",
          "type": 0,
          "icon": {
              "id": 85,
              "name": "Software Installer",
              "url16": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=16&inherited=false&abstract=false&time=1558951873000",
              "url48": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=48&inherited=false&abstract=false&time=1558951873000"
          },
          "position": 12,
          "created": "27/May/19 3:05 AM",
          "updated": "27/May/19 3:11 AM",
          "objectCount": 0,
          "objectSchemaId": 2,
          "inherited": false,
          "abstractObjectType": false,
          "parentObjectTypeInherited": false
      },
      "name": "Updated",
      "label": false,
      "type": 0,
      "defaultType": {
          "id": 6,
          "name": "DateTime"
      },
      "editable": false,
      "system": true,
      "sortable": true,
      "summable": false,
      "minimumCardinality": 1,
      "maximumCardinality": 1,
      "removable": false,
      "hidden": false,
      "includeChildObjectTypes": false,
      "uniqueAttribute": false,
      "options": "",
      "position": 3
  },
  {
      "id": 452,
      "objectType": {
          "id": 58,
          "name": "Vlan",
          "type": 0,
          "icon": {
              "id": 85,
              "name": "Software Installer",
              "url16": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=16&inherited=false&abstract=false&time=1558951873000",
              "url48": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=48&inherited=false&abstract=false&time=1558951873000"
          },
          "position": 12,
          "created": "27/May/19 3:05 AM",
          "updated": "27/May/19 3:11 AM",
          "objectCount": 0,
          "objectSchemaId": 2,
          "inherited": false,
          "abstractObjectType": false,
          "parentObjectTypeInherited": false
      },
      "name": "Number",
      "label": false,
      "type": 0,
      "defaultType": {
          "id": 1,
          "name": "Integer"
      },
      "editable": true,
      "system": false,
      "sortable": true,
      "summable": false,
      "minimumCardinality": 0,
      "maximumCardinality": 1,
      "removable": true,
      "hidden": false,
      "includeChildObjectTypes": false,
      "uniqueAttribute": false,
      "options": "",
      "position": 4
  },
  {
      "id": 453,
      "objectType": {
          "id": 58,
          "name": "Vlan",
          "type": 0,
          "icon": {
              "id": 85,
              "name": "Software Installer",
              "url16": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=16&inherited=false&abstract=false&time=1558951873000",
              "url48": "https://jira.mydomain.com/rest/insight/1.0/objecttype/58/icon.png?size=48&inherited=false&abstract=false&time=1558951873000"
          },
          "position": 12,
          "created": "27/May/19 3:05 AM",
          "updated": "27/May/19 3:11 AM",
          "objectCount": 0,
          "objectSchemaId": 2,
          "inherited": false,
          "abstractObjectType": false,
          "parentObjectTypeInherited": false
      },
      "name": "Description",
      "label": false,
      "type": 0,
      "defaultType": {
          "id": 0,
          "name": "Text"
      },
      "editable": true,
      "system": false,
      "sortable": true,
      "summable": false,
      "minimumCardinality": 0,
      "maximumCardinality": 1,
      "removable": true,
      "hidden": false,
      "includeChildObjectTypes": false,
      "uniqueAttribute": false,
      "options": "",
      "position": 5
  },
  {
    "id": 193,
    "objectType": {
        "id": 18,
        "name": "Codebase",
        "type": 0,
        "icon": {
            "id": 97,
            "name": "Git",
            "url16": "https://jira.mydomain.com/rest/insight/1.0/objecttype/18/icon.png?size=16&inherited=false&abstract=false&time=1558951783000",
            "url48": "https://jira.mydomain.com/rest/insight/1.0/objecttype/18/icon.png?size=48&inherited=false&abstract=false&time=1558951783000"
        },
        "position": 1,
        "created": "25/Apr/19 7:12 PM",
        "updated": "27/May/19 3:09 AM",
        "objectCount": 0,
        "parentObjectTypeId": 16,
        "objectSchemaId": 2,
        "inherited": false,
        "abstractObjectType": false,
        "parentObjectTypeInherited": false
    },
    "name": "URI",
    "label": false,
    "type": 0,
    "description": "The URI of the Source control Repo",
    "defaultType": {
        "id": 7,
        "name": "URL"
    },
    "additionalValue": "DISABLED",
    "editable": true,
    "system": false,
    "sortable": true,
    "summable": false,
    "minimumCardinality": 0,
    "maximumCardinality": 1,
    "removable": true,
    "hidden": false,
    "includeChildObjectTypes": false,
    "uniqueAttribute": false,
    "options": "",
    "position": 4
}
]`
				It("returns a slice of InsightObjectAttribute", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(200, fixture))

					attributeList, err := testClient.GetObjectTypeAttributesForObjectTypeID("57")

					Expect(err).To(BeNil())

					var sliceOfAttributes *[]api.ObjectTypeAttribute
					Expect(attributeList).To(BeAssignableToTypeOf(sliceOfAttributes))
				})

				It("should have 7 attributes", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(200, fixture))

					attributeList, err := testClient.GetObjectTypeAttributesForObjectTypeID("57")

					Expect(err).To(BeNil())
					Expect(*attributeList).To(HaveLen(7))
				})
			})
		})
	})
})
