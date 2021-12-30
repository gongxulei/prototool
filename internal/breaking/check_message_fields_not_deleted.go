// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package breaking

import (
	"github.com/gongxulei/prototool/internal/extract"
	"github.com/gongxulei/prototool/internal/text"
)

func checkMessageFieldsNotDeleted(addFailure func(*text.Failure), from *extract.PackageSet, to *extract.PackageSet) error {
	return forEachMessagePair(addFailure, from, to, checkMessageFieldsNotDeletedMessage)
}

func checkMessageFieldsNotDeletedMessage(addFailure func(*text.Failure), from *extract.Message, to *extract.Message) error {
	fromFieldNumberToField := from.FieldNumberToField()
	toFieldNumberToField := to.FieldNumberToField()
	for fieldNumber := range fromFieldNumberToField {
		if _, ok := toFieldNumberToField[fieldNumber]; !ok {
			addFailure(newMessageFieldsNotDeletedFailure(from.FullyQualifiedName(), fieldNumber))
		}
	}
	return nil
}

func newMessageFieldsNotDeletedFailure(messageName string, fieldNumber int32) *text.Failure {
	return newTextFailuref(`Message field "%d" on message %q was deleted.`, fieldNumber, messageName)
}
