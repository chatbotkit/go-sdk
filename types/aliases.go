package types

// This file provides hand-maintained, stable aliases for generated type names.
//
// The generated types in this package (types.go) are produced by quicktype,
// which assigns arbitrary, order-dependent names (e.g. "FriskyType",
// "IndecentFeature") to anonymous nested objects and enums in the OpenAPI spec.
// Those names change whenever endpoints are added or reordered, which would
// otherwise break every consumer that references them.
//
// Reference the stable aliases below instead of the raw generated names. When
// types.go is regenerated and a referenced name changes, update ONLY the
// right-hand sides in this file - every consumer keeps using the unchanged
// left-hand names.
//
// This file is hand-maintained and is NOT overwritten by `script:sync-types:go`
// (which only writes types.go).

// --- Conversation complete request message role ---

// CompleteMessageType is the role of a message sent in a conversation complete
// request.
type CompleteMessageType = Type2

const (
	// CompleteMessageRoleUser marks a sent message authored by the end user.
	CompleteMessageRoleUser = MischievousUser
	// CompleteMessageRoleBot marks a sent message authored by the bot.
	CompleteMessageRoleBot = MischievousBot
)

// --- Conversation message list item role ---

// MessageItemType is the role of a message returned from a conversation message
// listing.
type MessageItemType = HilariousType

const (
	// MessageItemRoleUser marks a listed message authored by the end user.
	MessageItemRoleUser = HilariousUser
	// MessageItemRoleBot marks a listed message authored by the bot.
	MessageItemRoleBot = HilariousBot
)

// --- Conversation complete request (stateful) extensions ---

type (
	// CompleteDataset is an inline dataset on a conversation complete request.
	CompleteDataset = AmbitiousDataset
	// CompleteRecord is a record within a CompleteDataset.
	CompleteRecord = IndecentRecord
	// CompleteFeature is an inline feature on a conversation complete request.
	CompleteFeature = IndecentFeature
	// CompleteSkillset is an inline skillset on a conversation complete request.
	CompleteSkillset = AmbitiousSkillset
	// CompleteAbility is an ability within a CompleteSkillset.
	CompleteAbility = IndecentAbility
)

// --- Conversation complete request (stateful) functions ---

type (
	// CompleteFunctionParameters is the JSON-schema parameters of a function on
	// a conversation complete request.
	CompleteFunctionParameters = IndigoParameters
	// CompleteFunctionResult is the result configuration of a function on a
	// conversation complete request.
	CompleteFunctionResult = FriskyResult
)

// CompleteFunctionParametersTypeObject is the "object" value of a function
// parameters schema type on a conversation complete request.
const CompleteFunctionParametersTypeObject = IndigoObject

// --- Conversation message complete request (stateless) extensions ---

type (
	// MessageCompleteDataset is an inline dataset on a conversation message
	// complete request.
	MessageCompleteDataset = TentacledDataset
	// MessageCompleteRecord is a record within a MessageCompleteDataset.
	MessageCompleteRecord = PurpleRecord
	// MessageCompleteFeature is an inline feature on a conversation message
	// complete request.
	MessageCompleteFeature = PurpleFeature
	// MessageCompleteSkillset is an inline skillset on a conversation message
	// complete request.
	MessageCompleteSkillset = TentacledSkillset
	// MessageCompleteAbility is an ability within a MessageCompleteSkillset.
	MessageCompleteAbility = PurpleAbility
)

// --- Conversation message complete request (stateless) functions ---

type (
	// MessageCompleteFunctionParameters is the JSON-schema parameters of a
	// function on a conversation message complete request.
	MessageCompleteFunctionParameters = PurpleParameters
	// MessageCompleteFunctionResult is the result configuration of a function
	// on a conversation message complete request.
	MessageCompleteFunctionResult = HilariousResult
)

// MessageCompleteFunctionParametersTypeObject is the "object" value of a
// function parameters schema type on a conversation message complete request.
const MessageCompleteFunctionParametersTypeObject = PurpleObject
