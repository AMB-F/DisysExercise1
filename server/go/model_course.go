/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Course struct {
	Id int32 `json:"id"`

	Name string `json:"name,omitempty"`

	// Expected amount of hours per week to work
	ExpectedWorkload int32 `json:"expectedWorkload,omitempty"`

	// Decimal number from 1-5
	Rating float32 `json:"rating,omitempty"`
}
