// Copyright 2013 Kelsey Hightower
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

// non-working creds from the amazon authentication examples
// http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html#RESTAuthenticationExamples
var (
	AWSAccessKeyId     = "AKIAIOSFODNN7EXAMPLE"
	AWSSecretAccessKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
)

func StringToSign(httpVerb, contentMD5, contentType, date, canonicalizedAmzHeaders,
	canonicalizedResource string) string {

	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s\n%s%s", httpVerb, contentMD5,
		contentType, date, canonicalizedAmzHeaders, canonicalizedResource)
	return stringToSign
}

// Print out an authentication header for use with Amazon authenticated
// REST Requests as defined here:
//     http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html
func main() {
	// Setup a buffer to hold the base64 encoded HMAC-SHA1 of the stringToSign
	signature := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, signature)

	stringToSign := StringToSign("GET", "", "", "Tue, 27 Mar 2007 19:36:42 +0000", "", "/johnsmith/photos/puppy.jpg")

	// Hash the selected elements from the request using the HMAC-SHA1 algorithm
	digest := hmac.New(sha1.New, []byte(AWSSecretAccessKey))
	digest.Write([]byte(stringToSign))

	// base64 encode the digest
	encoder.Write(digest.Sum(nil))
	encoder.Close()

	// construct the authentication header
	AuthenticationHeader := fmt.Sprintf("Authorization: AWS %s:%s", AWSAccessKeyId, signature.String())

	fmt.Println(AuthenticationHeader)
}
