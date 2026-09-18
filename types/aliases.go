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
type CompleteMessageType = Type12

const (
	// CompleteMessageRoleUser marks a sent message authored by the end user.
	CompleteMessageRoleUser = AmbitiousUser
	// CompleteMessageRoleBot marks a sent message authored by the bot.
	CompleteMessageRoleBot = AmbitiousBot
)

// --- Conversation message list item role ---

// MessageItemType is the role of a message returned from a conversation message
// listing.
type MessageItemType = Type22

const (
	// MessageItemRoleUser marks a listed message authored by the end user.
	MessageItemRoleUser = User1
	// MessageItemRoleBot marks a listed message authored by the bot.
	MessageItemRoleBot = Bot1
)

// --- Conversation complete request (stateful) extensions ---

type (
	// CompleteDataset is an inline dataset on a conversation complete request.
	CompleteDataset = StickyDataset
	// CompleteRecord is a record within a CompleteDataset.
	CompleteRecord = FluffyRecord
	// CompleteFeature is an inline feature on a conversation complete request.
	CompleteFeature = FluffyFeature
	// CompleteSkillset is an inline skillset on a conversation complete request.
	CompleteSkillset = StickySkillset
	// CompleteAbility is an ability within a CompleteSkillset.
	CompleteAbility = FluffyAbility
)

// --- Conversation complete request (stateful) functions ---

type (
	// CompleteFunctionParameters is the JSON-schema parameters of a function on
	// a conversation complete request.
	CompleteFunctionParameters = FluffyParameters
	// CompleteFunctionResult is the result configuration of a function on a
	// conversation complete request.
	CompleteFunctionResult = AmbitiousResult
)

// CompleteFunctionParametersTypeObject is the "object" value of a function
// parameters schema type on a conversation complete request.
const CompleteFunctionParametersTypeObject = IndigoObject

// --- Conversation message complete request (stateless) extensions ---

type (
	// MessageCompleteDataset is an inline dataset on a conversation message
	// complete request.
	MessageCompleteDataset = CunningDataset
	// MessageCompleteRecord is a record within a MessageCompleteDataset.
	MessageCompleteRecord = HilariousRecord
	// MessageCompleteFeature is an inline feature on a conversation message
	// complete request.
	MessageCompleteFeature = HilariousFeature
	// MessageCompleteSkillset is an inline skillset on a conversation message
	// complete request.
	MessageCompleteSkillset = CunningSkillset
	// MessageCompleteAbility is an ability within a MessageCompleteSkillset.
	MessageCompleteAbility = HilariousAbility
)

// --- Conversation message complete request (stateless) functions ---

type (
	// MessageCompleteFunctionParameters is the JSON-schema parameters of a
	// function on a conversation message complete request.
	MessageCompleteFunctionParameters = IndecentParameters
	// MessageCompleteFunctionResult is the result configuration of a function
	// on a conversation message complete request.
	MessageCompleteFunctionResult = MischievousResult
)

// MessageCompleteFunctionParametersTypeObject is the "object" value of a
// function parameters schema type on a conversation message complete request.
const MessageCompleteFunctionParametersTypeObject = CunningObject

// --- Decision question and answer kinds ---

// DecisionQuestionKind is the type of a decision question and of its answer.
type DecisionQuestionKind = QuestionType

const (
	// DecisionQuestionKindBoolean is a yes or no question.
	DecisionQuestionKindBoolean = PurpleBoolean
	// DecisionQuestionKindChoice is a choice between several named options.
	DecisionQuestionKindChoice = PurpleChoice
	// DecisionQuestionKindScore is a level on an ordered scale.
	DecisionQuestionKindScore = PurpleScore
)

// DecisionCreateAnswer is an answer on a decision create response.
type DecisionCreateAnswer = Answer
