// Ref: https://www.sohamkamani.com/golang/json/
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Test struct {
	Name string
	// Ref: https://bitfieldconsulting.com/golang/map-string-interface
	Outputs map[string]interface{}
}

func main() {

	var str Test

	// var result map[string]any

	jsonString := `{
  "outputs": {
    "iam_policy_arn": {
      "value": "arn:aws:iam::801658281203:policy/k8s-partner-portal-financial-groc-web-bff-groc-conciliation-bff",
      "type": "string"
    },
    "iam_policy_document": {
      "value": {
        "Statement": [
          {
            "Action": [
              "s3:ListBucket",
              "s3:GetObject"
            ],
            "Effect": "Allow",
            "Resource": [
              "arn:aws:s3:::development-merchant-static-full/*",
              "arn:aws:s3:::development-merchant-static-full"
            ],
            "Sid": "TFStmt01"
          }
        ],
        "Version": "2012-10-17"
      },
      "type": [
        "object",
        {
          "Statement": [
            "tuple",
            [
              [
                "object",
                {
                  "Action": [
                    "tuple",
                    [
                      "string",
                      "string"
                    ]
                  ],
                  "Effect": "string",
                  "Resource": [
                    "tuple",
                    [
                      "string",
                      "string"
                    ]
                  ],
                  "Sid": "string"
                }
              ]
            ]
          ],
          "Version": "string"
        }
      ]
    },
    "iam_policy_name": {
      "value": "k8s-partner-portal-financial-groc-web-bff-groc-conciliation-bff",
      "type": "string"
    },
    "iam_role_arn": {
      "value": "arn:aws:iam::801658281203:role/k8s-partner-portal-financial-groc-web-bff-groc-conciliation-bff",
      "type": "string"
    },
    "iam_role_name": {
      "value": "k8s-partner-portal-financial-groc-web-bff-groc-conciliation-bff",
      "type": "string"
    }
  }
}`

	fmt.Println(reflect.TypeOf(jsonString))

	json.Unmarshal([]byte(jsonString), &str.Outputs)

	fmt.Println(str.Outputs)

}
