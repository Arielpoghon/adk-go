// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openaimodel

import "google.golang.org/adk/v2/model/openaimodel/internal/openaicommon"

// The sentinels are defined in the internal package the endpoint paths return
// them from, and aliased here. Assignment preserves identity, so errors.Is is
// unaffected.
var (
	// ErrModelNameRequired is returned when a model name is not provided.
	ErrModelNameRequired = openaicommon.ErrModelNameRequired
	// ErrUnsupportedAPI is returned when ClientConfig.API names an API this package does not implement.
	ErrUnsupportedAPI = openaicommon.ErrUnsupportedAPI
	// ErrNoChoices is returned when a Chat Completions response carries no choices.
	ErrNoChoices = openaicommon.ErrNoChoices
	// ErrRequestNil is returned when the provided request is nil.
	ErrRequestNil = openaicommon.ErrRequestNil
	// ErrNoContents is returned when the LLM request has no contents.
	ErrNoContents = openaicommon.ErrNoContents
	// ErrFunctionCallMissingName is returned when a function call is missing a name.
	ErrFunctionCallMissingName = openaicommon.ErrFunctionCallMissingName
	// ErrTopKNotSupported is returned when TopK is used, which is not supported.
	ErrTopKNotSupported = openaicommon.ErrTopKNotSupported
	// ErrStopSequencesNotSupported is returned when stop sequences are used with the Responses API, which does not support them.
	ErrStopSequencesNotSupported = openaicommon.ErrStopSequencesNotSupported
	// ErrMultipleCandidatesNotSupported is returned when multiple candidates are requested, which is not supported.
	ErrMultipleCandidatesNotSupported = openaicommon.ErrMultipleCandidatesNotSupported
	// ErrPenaltiesNotSupported is returned when frequency/presence penalties are used with the Responses API, which does not support them.
	ErrPenaltiesNotSupported = openaicommon.ErrPenaltiesNotSupported
	// ErrLabelsNotSupported is returned when request labels are used, which is not supported.
	ErrLabelsNotSupported = openaicommon.ErrLabelsNotSupported
	// ErrSafetySettingsNotSupported is returned when Gemini safety settings are used, which is not supported.
	ErrSafetySettingsNotSupported = openaicommon.ErrSafetySettingsNotSupported
	// ErrUnsupportedMIMEType is returned when an unsupported MIME type is used.
	ErrUnsupportedMIMEType = openaicommon.ErrUnsupportedMIMEType
	// ErrUnsupportedConfigField is returned when a generation config field has no equivalent on the selected API; the message names it.
	ErrUnsupportedConfigField = openaicommon.ErrUnsupportedConfigField
	// ErrEmptyJSONSchema is returned when an empty JSON schema is provided.
	ErrEmptyJSONSchema = openaicommon.ErrEmptyJSONSchema
	// ErrEmptyResponse is returned when the OpenAI API returns an empty response.
	ErrEmptyResponse = openaicommon.ErrEmptyResponse
	// ErrNoOutputItems is returned when the response contains no output items.
	ErrNoOutputItems = openaicommon.ErrNoOutputItems
	// ErrUnsupportedMessageContentType is returned when an unsupported message content type is used.
	ErrUnsupportedMessageContentType = openaicommon.ErrUnsupportedMessageContentType
	// ErrUnsupportedOutputItemType is returned when an unsupported output item type is used.
	ErrUnsupportedOutputItemType = openaicommon.ErrUnsupportedOutputItemType
	// ErrFunctionCallArgs is returned when a function call's arguments are not a decodable JSON object.
	ErrFunctionCallArgs = openaicommon.ErrFunctionCallArgs
	// ErrNoTextOrToolContent is returned when the response output does not contain text or tool content.
	ErrNoTextOrToolContent = openaicommon.ErrNoTextOrToolContent
	// ErrResponseFailed is returned when the server reports the response itself
	// as failed, whatever output it came with. Such a failure arrives as HTTP
	// 200, and both a blocking call and a stream report it in place of the
	// output that would otherwise have read as a turn.
	ErrResponseFailed = openaicommon.ErrResponseFailed
)
