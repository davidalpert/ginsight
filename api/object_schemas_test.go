package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"

	api "github.com/davidalpert/ginsight/api"
)

var _ = Describe("Client", func() {
	Describe("ObjectSchemas", func() {
		Describe("List All", func() {
			endpoint := "/rest/insight/1.0/objectschema/list"

			Context("With no schemas", func() {
				fixture := `{
          "objectschemas": []
        }`

				It("returns a slice of InsightObjectSchema", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(200, fixture))

					schemaList, err := testClient.GetObjectSchemas()

					Expect(err).To(BeNil())

					var sliceOfObjectSchema *api.ObjectSchemaList
					Expect(schemaList).To(BeAssignableToTypeOf(sliceOfObjectSchema))
				})

				It("should be empty", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(200, fixture))

					schemaList, err := testClient.GetObjectSchemas()

					Expect(err).To(BeNil())
					Expect(schemaList.Schemas).To(HaveLen(0))
				})
			})

			Context("With 2 schemas", func() {
				fixture := `
          {
            "objectschemas": [
               {
                  "id": 3,
                  "name": "Hardware",
                  "objectSchemaKey": "HARD",
                  "status": "Ok",
                  "created": "06/May/19 6:33 AM",
                  "updated": "06/May/19 6:33 AM",
                  "objectCount": 0,
                  "objectTypeCount": 0
               },
               {
                 "id": 5,
                 "name": "Computers",
                 "objectSchemaKey": "IT",
                 "status": "Ok",
                 "description": "The IT department schema",
                 "created": "25/May/19 3:42 AM",
                 "updated": "25/May/19 3:42 AM",
                 "objectCount": 0,
                 "objectTypeCount": 0
              }
            ]
          }`

				It("should deserialize to two schemas", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(200, fixture))

					schemaList, err := testClient.GetObjectSchemas()

					Expect(err).To(BeNil())
					Expect(schemaList.Schemas).To(HaveLen(2))
				})
			})
		})

		Describe("Create one ObjectSchema", func() {
			endpoint := "/rest/insight/1.0/objectschema/create"
			schemaCreate := api.ObjectSchemaCreateUpdateRequest{
				Name:        "Computers",
				Key:         "COMP",
				Description: "The IT department schema",
			}

			Context("Successful", func() {
				fixture := `{
          "id": 5,
          "name": "Computers",
          "objectSchemaKey": "COMP",
          "status": "Ok",
          "description": "The IT department schema",
          "created": "25/May/19 3:42 AM",
          "updated": "25/May/19 3:42 AM",
          "objectCount": 0,
          "objectTypeCount": 0
        }`

				It("completes successfully", func() {
					httpmock.RegisterResponder("POST", endpoint, InsightApiResponder(201, fixture))

					schemaCreate := api.ObjectSchemaCreateUpdateRequest{
						Name:        "Computers",
						Key:         "COMP",
						Description: "The IT department schema",
					}

					newSchema, err := testClient.CreateSchema(&schemaCreate)

					Expect(err).To(BeNil())
					Expect(newSchema).To(Not(BeNil()))
					Expect(newSchema.Name).To(Equal(schemaCreate.Name))
					Expect(newSchema.Key).To(Equal(schemaCreate.Key))
					Expect(newSchema.Description).To(Equal(schemaCreate.Description))
				})
			})

			Context("Name already exists", func() {
				fixture := `{
          "errorMessages": [],
          "errors": {
            "name": "Name has to be unique!"
          }
        }`

				It("returns an error", func() {
					httpmock.RegisterResponder("POST", endpoint, InsightApiResponder(400, fixture))

					_, err := testClient.CreateSchema(&schemaCreate)

					Expect(err).To(Not(BeNil()))
					clientError := api.ClientError{}
					Expect(err).To(BeAssignableToTypeOf(&clientError))
					Expect(err.Error()).To(Equal("400 " + fixture + "\n"))
				})
			})
		})

		Describe("Get one ObjectSchema", func() {
			endpoint := "/rest/insight/1.0/objectschema/5"

			Context("Successful", func() {
				fixture := `{
          "id": 5,
          "name": "Computers",
          "objectSchemaKey": "IT",
          "status": "Ok",
          "description": "The IT department schema",
          "created": "25/May/19 3:42 AM",
          "updated": "25/May/19 3:42 AM",
          "objectCount": 0,
          "objectTypeCount": 0
        }`

				It("schema exists", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(200, fixture))

					schema, err := testClient.GetObjectSchemaById("5")

					Expect(err).To(BeNil())
					Expect(schema).To(Not(BeNil()))
					Expect(schema.Name).To(Equal("Computers"))
					Expect(schema.Key).To(Equal("IT"))
					Expect(schema.Description).To(Equal("The IT department schema"))
				})
			})

			Context("id does not exist", func() {
				fixture := `{
          "errorMessages": [
              "NotFoundInsightException: Could not find SchemaObject with id: 5"
          ],
          "errors": {}
        }`

				It("returns an error", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(404, fixture))

					_, err := testClient.GetObjectSchemaById("5")

					Expect(err).To(Not(BeNil()))
					clientError := api.ClientError{}
					Expect(err).To(BeAssignableToTypeOf(&clientError))
					Expect(err.Error()).To(Equal("404 " + fixture + "\n"))
				})
			})
		})

		Describe("Delete one ObjectSchema by key", func() {
			listSchemasResponse := `
      {
        "objectschemas": [
           {
              "id": 3,
              "name": "Hardware",
              "objectSchemaKey": "HARD",
              "status": "Ok",
              "created": "06/May/19 6:33 AM",
              "updated": "06/May/19 6:33 AM",
              "objectCount": 0,
              "objectTypeCount": 0
           },
           {
             "id": 5,
             "name": "Computers",
             "objectSchemaKey": "IT",
             "status": "Ok",
             "description": "The IT department schema",
             "created": "25/May/19 3:42 AM",
             "updated": "25/May/19 3:42 AM",
             "objectCount": 0,
             "objectTypeCount": 0
          }
        ]
      }`

			Context("Schema exists", func() {
				It("returns no error", func() {
					httpmock.RegisterResponder("GET", "/rest/insight/1.0/objectschema/list", InsightApiResponder(200, listSchemasResponse))
					httpmock.RegisterResponder("DELETE", "/rest/insight/1.0/objectschema/5", InsightApiResponder(200, ""))

					err := testClient.DeleteSchemaByKey("IT")

					Expect(err).To(BeNil())
				})
			})

			Context("id does not exist", func() {
				notFoundResponse := `{
          "errorMessages": [
              "NotFoundInsightException: Could not find SchemaObject with key: ITX"
          ],
          "errors": {}
        }`

				It("returns an error", func() {
					httpmock.RegisterResponder("GET", "/rest/insight/1.0/objectschema/list", InsightApiResponder(200, listSchemasResponse))
					httpmock.RegisterResponder("DELETE", "/rest/insight/1.0/objectschema/5", InsightApiResponder(404, notFoundResponse))

					err := testClient.DeleteSchemaByKey("ITX")

					Expect(err).To(Not(BeNil()))
					clientError := api.ObjectSchemaNotFoundError{}
					Expect(err).To(BeAssignableToTypeOf(&clientError))
					Expect(err.Error()).To(Equal("Did not find schema 'ITX'\n\nAre you looking for one of these schemas? [HARD IT]\n"))
				})
			})
		})

		Describe("Delete one ObjectSchema", func() {
			endpoint := MockURLFor("/rest/insight/1.0/objectschema/5")

			Context("Schema exists", func() {
				It("returns no error", func() {
					httpmock.RegisterResponder("DELETE", endpoint, InsightApiResponder(200, ""))

					err := testClient.DeleteSchema("5")

					Expect(err).To(BeNil())
				})
			})

			Context("id does not exist", func() {
				fixture := `{
          "errorMessages": [
              "NotFoundInsightException: Could not find SchemaObject with id: 5"
          ],
          "errors": {}
        }`

				It("returns an error", func() {
					httpmock.RegisterResponder("DELETE", endpoint, InsightApiResponder(404, fixture))

					err := testClient.DeleteSchema("5")

					Expect(err).To(Not(BeNil()))
					clientError := api.ClientError{}
					Expect(err).To(BeAssignableToTypeOf(&clientError))
					Expect(err.Error()).To(Equal("404 " + fixture + "\n"))
				})
			})
		})

		Describe("Update one ObjectSchema", func() {
			endpoint := "/rest/insight/1.0/objectschema/5"
			schemaUpdate := api.ObjectSchemaCreateUpdateRequest{
				Name:        "Computers",
				Key:         "IT",
				Description: "The IT department schema",
			}
			existingSchemaResponse := `{
    "id": 5,
    "name": "Computer",
    "objectSchemaKey": "IT",
    "status": "Ok",
    "description": "The IT department schema",
    "created": "25/May/19 3:42 AM",
    "updated": "25/May/19 3:42 AM",
    "objectCount": 0,
    "objectTypeCount": 0
   }`

			Context("Successful", func() {
				fixture := `{
          "id": 5,
          "name": "Computers",
          "objectSchemaKey": "IT",
          "status": "Ok",
          "description": "The IT department schema",
          "created": "25/May/19 3:42 AM",
          "updated": "25/May/19 3:42 AM",
          "objectCount": 0,
          "objectTypeCount": 0
        }`

				It("completes successfully", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(200, existingSchemaResponse))
					httpmock.RegisterResponder("PUT", endpoint, InsightApiResponder(201, fixture))

					updatedSchema, err := testClient.UpdateSchema("5", &schemaUpdate)

					Expect(err).To(BeNil())
					Expect(updatedSchema).To(Not(BeNil()))
					Expect(updatedSchema.Name).To(Equal(schemaUpdate.Name))
					Expect(updatedSchema.Key).To(Equal(schemaUpdate.Key))
					Expect(updatedSchema.Description).To(Equal(schemaUpdate.Description))
				})
			})

			Context("Name does not exist", func() {
				fixture := `{
          "errorMessages": [],
          "errors": {
            "name": "Name has to be unique!"
          }
        }`

				It("returns an error", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(200, existingSchemaResponse))
					httpmock.RegisterResponder("PUT", endpoint, InsightApiResponder(404, fixture))

					_, err := testClient.UpdateSchema("5", &schemaUpdate)

					Expect(err).To(Not(BeNil()))
					clientError := api.ClientError{}
					Expect(err).To(BeAssignableToTypeOf(&clientError))
					Expect(err.Error()).To(Equal("404 " + fixture + "\n"))
				})
			})

			Context("Changed key raises an error", func() {
				// the insight api does not let you change the Jira KEY once a schem is created
				schemaUpdate2 := api.ObjectSchemaCreateUpdateRequest{
					Name:        "Computers",
					Key:         "IT2", // <-- assume was already set as 'IT'; see existingSchemaResponse.objectSchemaKey
					Description: "The IT department schema",
				}

				It("returns an error", func() {
					httpmock.RegisterResponder("GET", endpoint, InsightApiResponder(200, existingSchemaResponse))

					_, err := testClient.UpdateSchema("5", &schemaUpdate2)

					Expect(err).To(Not(BeNil()))
					clientError := api.ObjectSchemaKeyMismatchError{}
					Expect(err).To(BeAssignableToTypeOf(&clientError))
				})
			})
		})
	})
})
