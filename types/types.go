// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    graphqlRequest, err := UnmarshalGraphqlRequest(bytes)
//    bytes, err = graphqlRequest.Marshal()
//
//    graphqlResponse, err := UnmarshalGraphqlResponse(bytes)
//    bytes, err = graphqlResponse.Marshal()
//
//    platformReportsListParams, err := UnmarshalPlatformReportsListParams(bytes)
//    bytes, err = platformReportsListParams.Marshal()
//
//    platformReportsListResponse, err := UnmarshalPlatformReportsListResponse(bytes)
//    bytes, err = platformReportsListResponse.Marshal()
//
//    reportGenerateParams, err := UnmarshalReportGenerateParams(bytes)
//    bytes, err = reportGenerateParams.Marshal()
//
//    reportGenerateRequest, err := UnmarshalReportGenerateRequest(bytes)
//    bytes, err = reportGenerateRequest.Marshal()
//
//    reportGenerateResponse, err := UnmarshalReportGenerateResponse(bytes)
//    bytes, err = reportGenerateResponse.Marshal()
//
//    reportsGenerateRequest, err := UnmarshalReportsGenerateRequest(bytes)
//    bytes, err = reportsGenerateRequest.Marshal()
//
//    reportsGenerateResponse, err := UnmarshalReportsGenerateResponse(bytes)
//    bytes, err = reportsGenerateResponse.Marshal()
//
//    blueprintCloneParams, err := UnmarshalBlueprintCloneParams(bytes)
//    bytes, err = blueprintCloneParams.Marshal()
//
//    blueprintCloneRequest, err := UnmarshalBlueprintCloneRequest(bytes)
//    bytes, err = blueprintCloneRequest.Marshal()
//
//    blueprintCloneResponse, err := UnmarshalBlueprintCloneResponse(bytes)
//    bytes, err = blueprintCloneResponse.Marshal()
//
//    blueprintDeleteParams, err := UnmarshalBlueprintDeleteParams(bytes)
//    bytes, err = blueprintDeleteParams.Marshal()
//
//    blueprintDeleteRequest, err := UnmarshalBlueprintDeleteRequest(bytes)
//    bytes, err = blueprintDeleteRequest.Marshal()
//
//    blueprintDeleteResponse, err := UnmarshalBlueprintDeleteResponse(bytes)
//    bytes, err = blueprintDeleteResponse.Marshal()
//
//    blueprintFetchParams, err := UnmarshalBlueprintFetchParams(bytes)
//    bytes, err = blueprintFetchParams.Marshal()
//
//    blueprintFetchResponse, err := UnmarshalBlueprintFetchResponse(bytes)
//    bytes, err = blueprintFetchResponse.Marshal()
//
//    blueprintResourcesListParams, err := UnmarshalBlueprintResourcesListParams(bytes)
//    bytes, err = blueprintResourcesListParams.Marshal()
//
//    blueprintResourcesListResponse, err := UnmarshalBlueprintResourcesListResponse(bytes)
//    bytes, err = blueprintResourcesListResponse.Marshal()
//
//    blueprintUpdateParams, err := UnmarshalBlueprintUpdateParams(bytes)
//    bytes, err = blueprintUpdateParams.Marshal()
//
//    blueprintUpdateRequest, err := UnmarshalBlueprintUpdateRequest(bytes)
//    bytes, err = blueprintUpdateRequest.Marshal()
//
//    blueprintUpdateResponse, err := UnmarshalBlueprintUpdateResponse(bytes)
//    bytes, err = blueprintUpdateResponse.Marshal()
//
//    blueprintCreateRequest, err := UnmarshalBlueprintCreateRequest(bytes)
//    bytes, err = blueprintCreateRequest.Marshal()
//
//    blueprintCreateResponse, err := UnmarshalBlueprintCreateResponse(bytes)
//    bytes, err = blueprintCreateResponse.Marshal()
//
//    blueprintsListParams, err := UnmarshalBlueprintsListParams(bytes)
//    bytes, err = blueprintsListParams.Marshal()
//
//    blueprintsListResponse, err := UnmarshalBlueprintsListResponse(bytes)
//    bytes, err = blueprintsListResponse.Marshal()
//
//    botCloneParams, err := UnmarshalBotCloneParams(bytes)
//    bytes, err = botCloneParams.Marshal()
//
//    botCloneRequest, err := UnmarshalBotCloneRequest(bytes)
//    bytes, err = botCloneRequest.Marshal()
//
//    botCloneResponse, err := UnmarshalBotCloneResponse(bytes)
//    bytes, err = botCloneResponse.Marshal()
//
//    botDeleteParams, err := UnmarshalBotDeleteParams(bytes)
//    bytes, err = botDeleteParams.Marshal()
//
//    botDeleteRequest, err := UnmarshalBotDeleteRequest(bytes)
//    bytes, err = botDeleteRequest.Marshal()
//
//    botDeleteResponse, err := UnmarshalBotDeleteResponse(bytes)
//    bytes, err = botDeleteResponse.Marshal()
//
//    botDownvoteParams, err := UnmarshalBotDownvoteParams(bytes)
//    bytes, err = botDownvoteParams.Marshal()
//
//    botDownvoteRequest, err := UnmarshalBotDownvoteRequest(bytes)
//    bytes, err = botDownvoteRequest.Marshal()
//
//    botDownvoteResponse, err := UnmarshalBotDownvoteResponse(bytes)
//    bytes, err = botDownvoteResponse.Marshal()
//
//    botFetchParams, err := UnmarshalBotFetchParams(bytes)
//    bytes, err = botFetchParams.Marshal()
//
//    botFetchResponse, err := UnmarshalBotFetchResponse(bytes)
//    bytes, err = botFetchResponse.Marshal()
//
//    botMemorySearchParams, err := UnmarshalBotMemorySearchParams(bytes)
//    bytes, err = botMemorySearchParams.Marshal()
//
//    botMemorySearchRequest, err := UnmarshalBotMemorySearchRequest(bytes)
//    bytes, err = botMemorySearchRequest.Marshal()
//
//    botMemorySearchResponse, err := UnmarshalBotMemorySearchResponse(bytes)
//    bytes, err = botMemorySearchResponse.Marshal()
//
//    botSessionCreateParams, err := UnmarshalBotSessionCreateParams(bytes)
//    bytes, err = botSessionCreateParams.Marshal()
//
//    botSessionCreateRequest, err := UnmarshalBotSessionCreateRequest(bytes)
//    bytes, err = botSessionCreateRequest.Marshal()
//
//    botSessionCreateResponse, err := UnmarshalBotSessionCreateResponse(bytes)
//    bytes, err = botSessionCreateResponse.Marshal()
//
//    botUpdateParams, err := UnmarshalBotUpdateParams(bytes)
//    bytes, err = botUpdateParams.Marshal()
//
//    botUpdateRequest, err := UnmarshalBotUpdateRequest(bytes)
//    bytes, err = botUpdateRequest.Marshal()
//
//    botUpdateResponse, err := UnmarshalBotUpdateResponse(bytes)
//    bytes, err = botUpdateResponse.Marshal()
//
//    botUpvoteParams, err := UnmarshalBotUpvoteParams(bytes)
//    bytes, err = botUpvoteParams.Marshal()
//
//    botUpvoteRequest, err := UnmarshalBotUpvoteRequest(bytes)
//    bytes, err = botUpvoteRequest.Marshal()
//
//    botUpvoteResponse, err := UnmarshalBotUpvoteResponse(bytes)
//    bytes, err = botUpvoteResponse.Marshal()
//
//    botUsageFetchParams, err := UnmarshalBotUsageFetchParams(bytes)
//    bytes, err = botUsageFetchParams.Marshal()
//
//    botUsageFetchResponse, err := UnmarshalBotUsageFetchResponse(bytes)
//    bytes, err = botUsageFetchResponse.Marshal()
//
//    botCreateRequest, err := UnmarshalBotCreateRequest(bytes)
//    bytes, err = botCreateRequest.Marshal()
//
//    botCreateResponse, err := UnmarshalBotCreateResponse(bytes)
//    bytes, err = botCreateResponse.Marshal()
//
//    botsListParams, err := UnmarshalBotsListParams(bytes)
//    bytes, err = botsListParams.Marshal()
//
//    botsListResponse, err := UnmarshalBotsListResponse(bytes)
//    bytes, err = botsListResponse.Marshal()
//
//    channelMessagePublishParams, err := UnmarshalChannelMessagePublishParams(bytes)
//    bytes, err = channelMessagePublishParams.Marshal()
//
//    channelMessagePublishRequest, err := UnmarshalChannelMessagePublishRequest(bytes)
//    bytes, err = channelMessagePublishRequest.Marshal()
//
//    channelMessagePublishResponse, err := UnmarshalChannelMessagePublishResponse(bytes)
//    bytes, err = channelMessagePublishResponse.Marshal()
//
//    channelMessagesSubscribeParams, err := UnmarshalChannelMessagesSubscribeParams(bytes)
//    bytes, err = channelMessagesSubscribeParams.Marshal()
//
//    channelMessagesSubscribeRequest, err := UnmarshalChannelMessagesSubscribeRequest(bytes)
//    bytes, err = channelMessagesSubscribeRequest.Marshal()
//
//    contactConversationsListParams, err := UnmarshalContactConversationsListParams(bytes)
//    bytes, err = contactConversationsListParams.Marshal()
//
//    contactConversationsListResponse, err := UnmarshalContactConversationsListResponse(bytes)
//    bytes, err = contactConversationsListResponse.Marshal()
//
//    contactDeleteParams, err := UnmarshalContactDeleteParams(bytes)
//    bytes, err = contactDeleteParams.Marshal()
//
//    contactDeleteRequest, err := UnmarshalContactDeleteRequest(bytes)
//    bytes, err = contactDeleteRequest.Marshal()
//
//    contactDeleteResponse, err := UnmarshalContactDeleteResponse(bytes)
//    bytes, err = contactDeleteResponse.Marshal()
//
//    contactFetchParams, err := UnmarshalContactFetchParams(bytes)
//    bytes, err = contactFetchParams.Marshal()
//
//    contactFetchResponse, err := UnmarshalContactFetchResponse(bytes)
//    bytes, err = contactFetchResponse.Marshal()
//
//    contactMemoriesListParams, err := UnmarshalContactMemoriesListParams(bytes)
//    bytes, err = contactMemoriesListParams.Marshal()
//
//    contactMemoriesListResponse, err := UnmarshalContactMemoriesListResponse(bytes)
//    bytes, err = contactMemoriesListResponse.Marshal()
//
//    contactMemorySearchParams, err := UnmarshalContactMemorySearchParams(bytes)
//    bytes, err = contactMemorySearchParams.Marshal()
//
//    contactMemorySearchRequest, err := UnmarshalContactMemorySearchRequest(bytes)
//    bytes, err = contactMemorySearchRequest.Marshal()
//
//    contactMemorySearchResponse, err := UnmarshalContactMemorySearchResponse(bytes)
//    bytes, err = contactMemorySearchResponse.Marshal()
//
//    contactSecretAuthenticateParams, err := UnmarshalContactSecretAuthenticateParams(bytes)
//    bytes, err = contactSecretAuthenticateParams.Marshal()
//
//    contactSecretAuthenticateRequest, err := UnmarshalContactSecretAuthenticateRequest(bytes)
//    bytes, err = contactSecretAuthenticateRequest.Marshal()
//
//    contactSecretAuthenticateResponse, err := UnmarshalContactSecretAuthenticateResponse(bytes)
//    bytes, err = contactSecretAuthenticateResponse.Marshal()
//
//    contactSecretRevokeParams, err := UnmarshalContactSecretRevokeParams(bytes)
//    bytes, err = contactSecretRevokeParams.Marshal()
//
//    contactSecretRevokeRequest, err := UnmarshalContactSecretRevokeRequest(bytes)
//    bytes, err = contactSecretRevokeRequest.Marshal()
//
//    contactSecretRevokeResponse, err := UnmarshalContactSecretRevokeResponse(bytes)
//    bytes, err = contactSecretRevokeResponse.Marshal()
//
//    contactSecretVerifyParams, err := UnmarshalContactSecretVerifyParams(bytes)
//    bytes, err = contactSecretVerifyParams.Marshal()
//
//    contactSecretVerifyRequest, err := UnmarshalContactSecretVerifyRequest(bytes)
//    bytes, err = contactSecretVerifyRequest.Marshal()
//
//    contactSecretVerifyResponse, err := UnmarshalContactSecretVerifyResponse(bytes)
//    bytes, err = contactSecretVerifyResponse.Marshal()
//
//    contactSecretsListParams, err := UnmarshalContactSecretsListParams(bytes)
//    bytes, err = contactSecretsListParams.Marshal()
//
//    contactSecretsListResponse, err := UnmarshalContactSecretsListResponse(bytes)
//    bytes, err = contactSecretsListResponse.Marshal()
//
//    contactSpacesListParams, err := UnmarshalContactSpacesListParams(bytes)
//    bytes, err = contactSpacesListParams.Marshal()
//
//    contactSpacesListResponse, err := UnmarshalContactSpacesListResponse(bytes)
//    bytes, err = contactSpacesListResponse.Marshal()
//
//    contactTasksListParams, err := UnmarshalContactTasksListParams(bytes)
//    bytes, err = contactTasksListParams.Marshal()
//
//    contactTasksListResponse, err := UnmarshalContactTasksListResponse(bytes)
//    bytes, err = contactTasksListResponse.Marshal()
//
//    contactUpdateParams, err := UnmarshalContactUpdateParams(bytes)
//    bytes, err = contactUpdateParams.Marshal()
//
//    contactUpdateRequest, err := UnmarshalContactUpdateRequest(bytes)
//    bytes, err = contactUpdateRequest.Marshal()
//
//    contactUpdateResponse, err := UnmarshalContactUpdateResponse(bytes)
//    bytes, err = contactUpdateResponse.Marshal()
//
//    contactCreateRequest, err := UnmarshalContactCreateRequest(bytes)
//    bytes, err = contactCreateRequest.Marshal()
//
//    contactCreateResponse, err := UnmarshalContactCreateResponse(bytes)
//    bytes, err = contactCreateResponse.Marshal()
//
//    contactEnsureRequest, err := UnmarshalContactEnsureRequest(bytes)
//    bytes, err = contactEnsureRequest.Marshal()
//
//    contactEnsureResponse, err := UnmarshalContactEnsureResponse(bytes)
//    bytes, err = contactEnsureResponse.Marshal()
//
//    contactsExportParams, err := UnmarshalContactsExportParams(bytes)
//    bytes, err = contactsExportParams.Marshal()
//
//    contactsExportResponse, err := UnmarshalContactsExportResponse(bytes)
//    bytes, err = contactsExportResponse.Marshal()
//
//    contactsListParams, err := UnmarshalContactsListParams(bytes)
//    bytes, err = contactsListParams.Marshal()
//
//    contactsListResponse, err := UnmarshalContactsListResponse(bytes)
//    bytes, err = contactsListResponse.Marshal()
//
//    conversationAttachmentUploadParams, err := UnmarshalConversationAttachmentUploadParams(bytes)
//    bytes, err = conversationAttachmentUploadParams.Marshal()
//
//    conversationAttachmentUploadRequest, err := UnmarshalConversationAttachmentUploadRequest(bytes)
//    bytes, err = conversationAttachmentUploadRequest.Marshal()
//
//    conversationAttachmentUploadResponse, err := UnmarshalConversationAttachmentUploadResponse(bytes)
//    bytes, err = conversationAttachmentUploadResponse.Marshal()
//
//    conversationMessageCompleteParams, err := UnmarshalConversationMessageCompleteParams(bytes)
//    bytes, err = conversationMessageCompleteParams.Marshal()
//
//    conversationMessageCompleteRequest, err := UnmarshalConversationMessageCompleteRequest(bytes)
//    bytes, err = conversationMessageCompleteRequest.Marshal()
//
//    conversationMessageCompleteResponse, err := UnmarshalConversationMessageCompleteResponse(bytes)
//    bytes, err = conversationMessageCompleteResponse.Marshal()
//
//    conversationContactUpsertParams, err := UnmarshalConversationContactUpsertParams(bytes)
//    bytes, err = conversationContactUpsertParams.Marshal()
//
//    conversationContactUpsertRequest, err := UnmarshalConversationContactUpsertRequest(bytes)
//    bytes, err = conversationContactUpsertRequest.Marshal()
//
//    conversationContactUpsertResponse, err := UnmarshalConversationContactUpsertResponse(bytes)
//    bytes, err = conversationContactUpsertResponse.Marshal()
//
//    conversationDeleteParams, err := UnmarshalConversationDeleteParams(bytes)
//    bytes, err = conversationDeleteParams.Marshal()
//
//    conversationDeleteRequest, err := UnmarshalConversationDeleteRequest(bytes)
//    bytes, err = conversationDeleteRequest.Marshal()
//
//    conversationDeleteResponse, err := UnmarshalConversationDeleteResponse(bytes)
//    bytes, err = conversationDeleteResponse.Marshal()
//
//    statefulConversationDispatchRequest, err := UnmarshalStatefulConversationDispatchRequest(bytes)
//    bytes, err = statefulConversationDispatchRequest.Marshal()
//
//    statefulConversationDispatchResponse, err := UnmarshalStatefulConversationDispatchResponse(bytes)
//    bytes, err = statefulConversationDispatchResponse.Marshal()
//
//    conversationDownvoteParams, err := UnmarshalConversationDownvoteParams(bytes)
//    bytes, err = conversationDownvoteParams.Marshal()
//
//    conversationDownvoteRequest, err := UnmarshalConversationDownvoteRequest(bytes)
//    bytes, err = conversationDownvoteRequest.Marshal()
//
//    conversationDownvoteResponse, err := UnmarshalConversationDownvoteResponse(bytes)
//    bytes, err = conversationDownvoteResponse.Marshal()
//
//    conversationFetchParams, err := UnmarshalConversationFetchParams(bytes)
//    bytes, err = conversationFetchParams.Marshal()
//
//    conversationFetchResponse, err := UnmarshalConversationFetchResponse(bytes)
//    bytes, err = conversationFetchResponse.Marshal()
//
//    conversationMessageDeleteParams, err := UnmarshalConversationMessageDeleteParams(bytes)
//    bytes, err = conversationMessageDeleteParams.Marshal()
//
//    conversationMessageDeleteRequest, err := UnmarshalConversationMessageDeleteRequest(bytes)
//    bytes, err = conversationMessageDeleteRequest.Marshal()
//
//    conversationMessageDeleteResponse, err := UnmarshalConversationMessageDeleteResponse(bytes)
//    bytes, err = conversationMessageDeleteResponse.Marshal()
//
//    conversationMessageDownvoteParams, err := UnmarshalConversationMessageDownvoteParams(bytes)
//    bytes, err = conversationMessageDownvoteParams.Marshal()
//
//    conversationMessageDownvoteRequest, err := UnmarshalConversationMessageDownvoteRequest(bytes)
//    bytes, err = conversationMessageDownvoteRequest.Marshal()
//
//    conversationMessageDownvoteResponse, err := UnmarshalConversationMessageDownvoteResponse(bytes)
//    bytes, err = conversationMessageDownvoteResponse.Marshal()
//
//    conversationMessageFetchParams, err := UnmarshalConversationMessageFetchParams(bytes)
//    bytes, err = conversationMessageFetchParams.Marshal()
//
//    conversationMessageFetchResponse, err := UnmarshalConversationMessageFetchResponse(bytes)
//    bytes, err = conversationMessageFetchResponse.Marshal()
//
//    conversationMessageSynthesizeParams, err := UnmarshalConversationMessageSynthesizeParams(bytes)
//    bytes, err = conversationMessageSynthesizeParams.Marshal()
//
//    conversationMessageSynthesizeRequest, err := UnmarshalConversationMessageSynthesizeRequest(bytes)
//    bytes, err = conversationMessageSynthesizeRequest.Marshal()
//
//    conversationMessageSynthesizeResponse, err := UnmarshalConversationMessageSynthesizeResponse(bytes)
//    bytes, err = conversationMessageSynthesizeResponse.Marshal()
//
//    conversationMessageUpdateParams, err := UnmarshalConversationMessageUpdateParams(bytes)
//    bytes, err = conversationMessageUpdateParams.Marshal()
//
//    conversationMessageUpdateRequest, err := UnmarshalConversationMessageUpdateRequest(bytes)
//    bytes, err = conversationMessageUpdateRequest.Marshal()
//
//    conversationMessageUpdateResponse, err := UnmarshalConversationMessageUpdateResponse(bytes)
//    bytes, err = conversationMessageUpdateResponse.Marshal()
//
//    conversationMessageUpvoteParams, err := UnmarshalConversationMessageUpvoteParams(bytes)
//    bytes, err = conversationMessageUpvoteParams.Marshal()
//
//    conversationMessageUpvoteRequest, err := UnmarshalConversationMessageUpvoteRequest(bytes)
//    bytes, err = conversationMessageUpvoteRequest.Marshal()
//
//    conversationMessageUpvoteResponse, err := UnmarshalConversationMessageUpvoteResponse(bytes)
//    bytes, err = conversationMessageUpvoteResponse.Marshal()
//
//    conversationMessageCreateParams, err := UnmarshalConversationMessageCreateParams(bytes)
//    bytes, err = conversationMessageCreateParams.Marshal()
//
//    conversationMessageCreateRequest, err := UnmarshalConversationMessageCreateRequest(bytes)
//    bytes, err = conversationMessageCreateRequest.Marshal()
//
//    conversationMessageCreateResponse, err := UnmarshalConversationMessageCreateResponse(bytes)
//    bytes, err = conversationMessageCreateResponse.Marshal()
//
//    conversationMessagesListParams, err := UnmarshalConversationMessagesListParams(bytes)
//    bytes, err = conversationMessagesListParams.Marshal()
//
//    conversationMessagesListResponse, err := UnmarshalConversationMessagesListResponse(bytes)
//    bytes, err = conversationMessagesListResponse.Marshal()
//
//    conversationMessageReceiveParams, err := UnmarshalConversationMessageReceiveParams(bytes)
//    bytes, err = conversationMessageReceiveParams.Marshal()
//
//    conversationMessageReceiveRequest, err := UnmarshalConversationMessageReceiveRequest(bytes)
//    bytes, err = conversationMessageReceiveRequest.Marshal()
//
//    conversationMessageReceiveResponse, err := UnmarshalConversationMessageReceiveResponse(bytes)
//    bytes, err = conversationMessageReceiveResponse.Marshal()
//
//    conversationMessageSendParams, err := UnmarshalConversationMessageSendParams(bytes)
//    bytes, err = conversationMessageSendParams.Marshal()
//
//    conversationMessageSendRequest, err := UnmarshalConversationMessageSendRequest(bytes)
//    bytes, err = conversationMessageSendRequest.Marshal()
//
//    conversationMessageSendResponse, err := UnmarshalConversationMessageSendResponse(bytes)
//    bytes, err = conversationMessageSendResponse.Marshal()
//
//    conversationSessionCreateParams, err := UnmarshalConversationSessionCreateParams(bytes)
//    bytes, err = conversationSessionCreateParams.Marshal()
//
//    conversationSessionCreateRequest, err := UnmarshalConversationSessionCreateRequest(bytes)
//    bytes, err = conversationSessionCreateRequest.Marshal()
//
//    conversationSessionCreateResponse, err := UnmarshalConversationSessionCreateResponse(bytes)
//    bytes, err = conversationSessionCreateResponse.Marshal()
//
//    conversationUpdateParams, err := UnmarshalConversationUpdateParams(bytes)
//    bytes, err = conversationUpdateParams.Marshal()
//
//    conversationUpdateRequest, err := UnmarshalConversationUpdateRequest(bytes)
//    bytes, err = conversationUpdateRequest.Marshal()
//
//    conversationUpdateResponse, err := UnmarshalConversationUpdateResponse(bytes)
//    bytes, err = conversationUpdateResponse.Marshal()
//
//    conversationUpvoteParams, err := UnmarshalConversationUpvoteParams(bytes)
//    bytes, err = conversationUpvoteParams.Marshal()
//
//    conversationUpvoteRequest, err := UnmarshalConversationUpvoteRequest(bytes)
//    bytes, err = conversationUpvoteRequest.Marshal()
//
//    conversationUpvoteResponse, err := UnmarshalConversationUpvoteResponse(bytes)
//    bytes, err = conversationUpvoteResponse.Marshal()
//
//    conversationUsageFetchParams, err := UnmarshalConversationUsageFetchParams(bytes)
//    bytes, err = conversationUsageFetchParams.Marshal()
//
//    conversationUsageFetchResponse, err := UnmarshalConversationUsageFetchResponse(bytes)
//    bytes, err = conversationUsageFetchResponse.Marshal()
//
//    conversationCompleteRequest, err := UnmarshalConversationCompleteRequest(bytes)
//    bytes, err = conversationCompleteRequest.Marshal()
//
//    conversationCompleteResponse, err := UnmarshalConversationCompleteResponse(bytes)
//    bytes, err = conversationCompleteResponse.Marshal()
//
//    conversationCreateRequest, err := UnmarshalConversationCreateRequest(bytes)
//    bytes, err = conversationCreateRequest.Marshal()
//
//    conversationCreateResponse, err := UnmarshalConversationCreateResponse(bytes)
//    bytes, err = conversationCreateResponse.Marshal()
//
//    conversationDispatchRequest, err := UnmarshalConversationDispatchRequest(bytes)
//    bytes, err = conversationDispatchRequest.Marshal()
//
//    conversationDispatchResponse, err := UnmarshalConversationDispatchResponse(bytes)
//    bytes, err = conversationDispatchResponse.Marshal()
//
//    conversationsExportParams, err := UnmarshalConversationsExportParams(bytes)
//    bytes, err = conversationsExportParams.Marshal()
//
//    conversationsExportResponse, err := UnmarshalConversationsExportResponse(bytes)
//    bytes, err = conversationsExportResponse.Marshal()
//
//    conversationsListParams, err := UnmarshalConversationsListParams(bytes)
//    bytes, err = conversationsListParams.Marshal()
//
//    conversationsListResponse, err := UnmarshalConversationsListResponse(bytes)
//    bytes, err = conversationsListResponse.Marshal()
//
//    datasetDeleteParams, err := UnmarshalDatasetDeleteParams(bytes)
//    bytes, err = datasetDeleteParams.Marshal()
//
//    datasetDeleteRequest, err := UnmarshalDatasetDeleteRequest(bytes)
//    bytes, err = datasetDeleteRequest.Marshal()
//
//    datasetDeleteResponse, err := UnmarshalDatasetDeleteResponse(bytes)
//    bytes, err = datasetDeleteResponse.Marshal()
//
//    datasetFetchParams, err := UnmarshalDatasetFetchParams(bytes)
//    bytes, err = datasetFetchParams.Marshal()
//
//    datasetFetchResponse, err := UnmarshalDatasetFetchResponse(bytes)
//    bytes, err = datasetFetchResponse.Marshal()
//
//    datasetFileAttachParams, err := UnmarshalDatasetFileAttachParams(bytes)
//    bytes, err = datasetFileAttachParams.Marshal()
//
//    datasetFileAttachRequest, err := UnmarshalDatasetFileAttachRequest(bytes)
//    bytes, err = datasetFileAttachRequest.Marshal()
//
//    datasetFileAttachResponse, err := UnmarshalDatasetFileAttachResponse(bytes)
//    bytes, err = datasetFileAttachResponse.Marshal()
//
//    datasetFileDetachParams, err := UnmarshalDatasetFileDetachParams(bytes)
//    bytes, err = datasetFileDetachParams.Marshal()
//
//    datasetFileDetachRequest, err := UnmarshalDatasetFileDetachRequest(bytes)
//    bytes, err = datasetFileDetachRequest.Marshal()
//
//    datasetFileDetachResponse, err := UnmarshalDatasetFileDetachResponse(bytes)
//    bytes, err = datasetFileDetachResponse.Marshal()
//
//    datasetFileSyncParams, err := UnmarshalDatasetFileSyncParams(bytes)
//    bytes, err = datasetFileSyncParams.Marshal()
//
//    datasetFileSyncRequest, err := UnmarshalDatasetFileSyncRequest(bytes)
//    bytes, err = datasetFileSyncRequest.Marshal()
//
//    datasetFileSyncResponse, err := UnmarshalDatasetFileSyncResponse(bytes)
//    bytes, err = datasetFileSyncResponse.Marshal()
//
//    datasetFilesListParams, err := UnmarshalDatasetFilesListParams(bytes)
//    bytes, err = datasetFilesListParams.Marshal()
//
//    datasetFilesListResponse, err := UnmarshalDatasetFilesListResponse(bytes)
//    bytes, err = datasetFilesListResponse.Marshal()
//
//    datasetRecordDeleteParams, err := UnmarshalDatasetRecordDeleteParams(bytes)
//    bytes, err = datasetRecordDeleteParams.Marshal()
//
//    datasetRecordDeleteRequest, err := UnmarshalDatasetRecordDeleteRequest(bytes)
//    bytes, err = datasetRecordDeleteRequest.Marshal()
//
//    datasetRecordDeleteResponse, err := UnmarshalDatasetRecordDeleteResponse(bytes)
//    bytes, err = datasetRecordDeleteResponse.Marshal()
//
//    datasetRecordFetchParams, err := UnmarshalDatasetRecordFetchParams(bytes)
//    bytes, err = datasetRecordFetchParams.Marshal()
//
//    datasetRecordFetchResponse, err := UnmarshalDatasetRecordFetchResponse(bytes)
//    bytes, err = datasetRecordFetchResponse.Marshal()
//
//    datasetRecordUpdateParams, err := UnmarshalDatasetRecordUpdateParams(bytes)
//    bytes, err = datasetRecordUpdateParams.Marshal()
//
//    datasetRecordUpdateRequest, err := UnmarshalDatasetRecordUpdateRequest(bytes)
//    bytes, err = datasetRecordUpdateRequest.Marshal()
//
//    datasetRecordUpdateResponse, err := UnmarshalDatasetRecordUpdateResponse(bytes)
//    bytes, err = datasetRecordUpdateResponse.Marshal()
//
//    datasetRecordCreateParams, err := UnmarshalDatasetRecordCreateParams(bytes)
//    bytes, err = datasetRecordCreateParams.Marshal()
//
//    datasetRecordCreateRequest, err := UnmarshalDatasetRecordCreateRequest(bytes)
//    bytes, err = datasetRecordCreateRequest.Marshal()
//
//    datasetRecordCreateResponse, err := UnmarshalDatasetRecordCreateResponse(bytes)
//    bytes, err = datasetRecordCreateResponse.Marshal()
//
//    datasetRecordsExportParams, err := UnmarshalDatasetRecordsExportParams(bytes)
//    bytes, err = datasetRecordsExportParams.Marshal()
//
//    datasetRecordsExportResponse, err := UnmarshalDatasetRecordsExportResponse(bytes)
//    bytes, err = datasetRecordsExportResponse.Marshal()
//
//    datasetRecordsListParams, err := UnmarshalDatasetRecordsListParams(bytes)
//    bytes, err = datasetRecordsListParams.Marshal()
//
//    datasetRecordsListResponse, err := UnmarshalDatasetRecordsListResponse(bytes)
//    bytes, err = datasetRecordsListResponse.Marshal()
//
//    datasetSearchParams, err := UnmarshalDatasetSearchParams(bytes)
//    bytes, err = datasetSearchParams.Marshal()
//
//    datasetSearchRequest, err := UnmarshalDatasetSearchRequest(bytes)
//    bytes, err = datasetSearchRequest.Marshal()
//
//    datasetSearchResponse, err := UnmarshalDatasetSearchResponse(bytes)
//    bytes, err = datasetSearchResponse.Marshal()
//
//    datasetUpdateParams, err := UnmarshalDatasetUpdateParams(bytes)
//    bytes, err = datasetUpdateParams.Marshal()
//
//    datasetUpdateRequest, err := UnmarshalDatasetUpdateRequest(bytes)
//    bytes, err = datasetUpdateRequest.Marshal()
//
//    datasetUpdateResponse, err := UnmarshalDatasetUpdateResponse(bytes)
//    bytes, err = datasetUpdateResponse.Marshal()
//
//    datasetCreateRequest, err := UnmarshalDatasetCreateRequest(bytes)
//    bytes, err = datasetCreateRequest.Marshal()
//
//    datasetCreateResponse, err := UnmarshalDatasetCreateResponse(bytes)
//    bytes, err = datasetCreateResponse.Marshal()
//
//    datasetsListParams, err := UnmarshalDatasetsListParams(bytes)
//    bytes, err = datasetsListParams.Marshal()
//
//    datasetsListResponse, err := UnmarshalDatasetsListResponse(bytes)
//    bytes, err = datasetsListResponse.Marshal()
//
//    eventLogsExportParams, err := UnmarshalEventLogsExportParams(bytes)
//    bytes, err = eventLogsExportParams.Marshal()
//
//    eventLogsExportResponse, err := UnmarshalEventLogsExportResponse(bytes)
//    bytes, err = eventLogsExportResponse.Marshal()
//
//    eventLogsListParams, err := UnmarshalEventLogsListParams(bytes)
//    bytes, err = eventLogsListParams.Marshal()
//
//    eventLogsListResponse, err := UnmarshalEventLogsListResponse(bytes)
//    bytes, err = eventLogsListResponse.Marshal()
//
//    eventLogsSubscribeRequest, err := UnmarshalEventLogsSubscribeRequest(bytes)
//    bytes, err = eventLogsSubscribeRequest.Marshal()
//
//    fileDeleteParams, err := UnmarshalFileDeleteParams(bytes)
//    bytes, err = fileDeleteParams.Marshal()
//
//    fileDeleteRequest, err := UnmarshalFileDeleteRequest(bytes)
//    bytes, err = fileDeleteRequest.Marshal()
//
//    fileDeleteResponse, err := UnmarshalFileDeleteResponse(bytes)
//    bytes, err = fileDeleteResponse.Marshal()
//
//    fileDownloadParams, err := UnmarshalFileDownloadParams(bytes)
//    bytes, err = fileDownloadParams.Marshal()
//
//    fileDownloadResponse, err := UnmarshalFileDownloadResponse(bytes)
//    bytes, err = fileDownloadResponse.Marshal()
//
//    fileFetchParams, err := UnmarshalFileFetchParams(bytes)
//    bytes, err = fileFetchParams.Marshal()
//
//    fileFetchResponse, err := UnmarshalFileFetchResponse(bytes)
//    bytes, err = fileFetchResponse.Marshal()
//
//    fileSyncParams, err := UnmarshalFileSyncParams(bytes)
//    bytes, err = fileSyncParams.Marshal()
//
//    fileSyncRequest, err := UnmarshalFileSyncRequest(bytes)
//    bytes, err = fileSyncRequest.Marshal()
//
//    fileSyncResponse, err := UnmarshalFileSyncResponse(bytes)
//    bytes, err = fileSyncResponse.Marshal()
//
//    fileUpdateParams, err := UnmarshalFileUpdateParams(bytes)
//    bytes, err = fileUpdateParams.Marshal()
//
//    fileUpdateRequest, err := UnmarshalFileUpdateRequest(bytes)
//    bytes, err = fileUpdateRequest.Marshal()
//
//    fileUpdateResponse, err := UnmarshalFileUpdateResponse(bytes)
//    bytes, err = fileUpdateResponse.Marshal()
//
//    fileUploadParams, err := UnmarshalFileUploadParams(bytes)
//    bytes, err = fileUploadParams.Marshal()
//
//    fileUploadRequest, err := UnmarshalFileUploadRequest(bytes)
//    bytes, err = fileUploadRequest.Marshal()
//
//    fileUploadResponse, err := UnmarshalFileUploadResponse(bytes)
//    bytes, err = fileUploadResponse.Marshal()
//
//    fileCreateRequest, err := UnmarshalFileCreateRequest(bytes)
//    bytes, err = fileCreateRequest.Marshal()
//
//    fileCreateResponse, err := UnmarshalFileCreateResponse(bytes)
//    bytes, err = fileCreateResponse.Marshal()
//
//    filesListParams, err := UnmarshalFilesListParams(bytes)
//    bytes, err = filesListParams.Marshal()
//
//    filesListResponse, err := UnmarshalFilesListResponse(bytes)
//    bytes, err = filesListResponse.Marshal()
//
//    discordIntegrationDeleteParams, err := UnmarshalDiscordIntegrationDeleteParams(bytes)
//    bytes, err = discordIntegrationDeleteParams.Marshal()
//
//    discordIntegrationDeleteRequest, err := UnmarshalDiscordIntegrationDeleteRequest(bytes)
//    bytes, err = discordIntegrationDeleteRequest.Marshal()
//
//    discordIntegrationDeleteResponse, err := UnmarshalDiscordIntegrationDeleteResponse(bytes)
//    bytes, err = discordIntegrationDeleteResponse.Marshal()
//
//    discordIntegrationFetchParams, err := UnmarshalDiscordIntegrationFetchParams(bytes)
//    bytes, err = discordIntegrationFetchParams.Marshal()
//
//    discordIntegrationFetchResponse, err := UnmarshalDiscordIntegrationFetchResponse(bytes)
//    bytes, err = discordIntegrationFetchResponse.Marshal()
//
//    discordIntegrationSetupParams, err := UnmarshalDiscordIntegrationSetupParams(bytes)
//    bytes, err = discordIntegrationSetupParams.Marshal()
//
//    discordIntegrationSetupRequest, err := UnmarshalDiscordIntegrationSetupRequest(bytes)
//    bytes, err = discordIntegrationSetupRequest.Marshal()
//
//    discordIntegrationSetupResponse, err := UnmarshalDiscordIntegrationSetupResponse(bytes)
//    bytes, err = discordIntegrationSetupResponse.Marshal()
//
//    discordIntegrationUpdateParams, err := UnmarshalDiscordIntegrationUpdateParams(bytes)
//    bytes, err = discordIntegrationUpdateParams.Marshal()
//
//    discordIntegrationUpdateRequest, err := UnmarshalDiscordIntegrationUpdateRequest(bytes)
//    bytes, err = discordIntegrationUpdateRequest.Marshal()
//
//    discordIntegrationUpdateResponse, err := UnmarshalDiscordIntegrationUpdateResponse(bytes)
//    bytes, err = discordIntegrationUpdateResponse.Marshal()
//
//    discordIntegrationCreateRequest, err := UnmarshalDiscordIntegrationCreateRequest(bytes)
//    bytes, err = discordIntegrationCreateRequest.Marshal()
//
//    discordIntegrationCreateResponse, err := UnmarshalDiscordIntegrationCreateResponse(bytes)
//    bytes, err = discordIntegrationCreateResponse.Marshal()
//
//    discordIntegrationsListParams, err := UnmarshalDiscordIntegrationsListParams(bytes)
//    bytes, err = discordIntegrationsListParams.Marshal()
//
//    discordIntegrationsListResponse, err := UnmarshalDiscordIntegrationsListResponse(bytes)
//    bytes, err = discordIntegrationsListResponse.Marshal()
//
//    emailIntegrationDeleteParams, err := UnmarshalEmailIntegrationDeleteParams(bytes)
//    bytes, err = emailIntegrationDeleteParams.Marshal()
//
//    emailIntegrationDeleteRequest, err := UnmarshalEmailIntegrationDeleteRequest(bytes)
//    bytes, err = emailIntegrationDeleteRequest.Marshal()
//
//    emailIntegrationDeleteResponse, err := UnmarshalEmailIntegrationDeleteResponse(bytes)
//    bytes, err = emailIntegrationDeleteResponse.Marshal()
//
//    emailIntegrationFetchParams, err := UnmarshalEmailIntegrationFetchParams(bytes)
//    bytes, err = emailIntegrationFetchParams.Marshal()
//
//    emailIntegrationFetchResponse, err := UnmarshalEmailIntegrationFetchResponse(bytes)
//    bytes, err = emailIntegrationFetchResponse.Marshal()
//
//    emailIntegrationSetupParams, err := UnmarshalEmailIntegrationSetupParams(bytes)
//    bytes, err = emailIntegrationSetupParams.Marshal()
//
//    emailIntegrationSetupRequest, err := UnmarshalEmailIntegrationSetupRequest(bytes)
//    bytes, err = emailIntegrationSetupRequest.Marshal()
//
//    emailIntegrationSetupResponse, err := UnmarshalEmailIntegrationSetupResponse(bytes)
//    bytes, err = emailIntegrationSetupResponse.Marshal()
//
//    emailIntegrationUpdateParams, err := UnmarshalEmailIntegrationUpdateParams(bytes)
//    bytes, err = emailIntegrationUpdateParams.Marshal()
//
//    emailIntegrationUpdateRequest, err := UnmarshalEmailIntegrationUpdateRequest(bytes)
//    bytes, err = emailIntegrationUpdateRequest.Marshal()
//
//    emailIntegrationUpdateResponse, err := UnmarshalEmailIntegrationUpdateResponse(bytes)
//    bytes, err = emailIntegrationUpdateResponse.Marshal()
//
//    emailIntegrationCreateRequest, err := UnmarshalEmailIntegrationCreateRequest(bytes)
//    bytes, err = emailIntegrationCreateRequest.Marshal()
//
//    emailIntegrationCreateResponse, err := UnmarshalEmailIntegrationCreateResponse(bytes)
//    bytes, err = emailIntegrationCreateResponse.Marshal()
//
//    emailIntegrationsListParams, err := UnmarshalEmailIntegrationsListParams(bytes)
//    bytes, err = emailIntegrationsListParams.Marshal()
//
//    emailIntegrationsListResponse, err := UnmarshalEmailIntegrationsListResponse(bytes)
//    bytes, err = emailIntegrationsListResponse.Marshal()
//
//    extractIntegrationDeleteParams, err := UnmarshalExtractIntegrationDeleteParams(bytes)
//    bytes, err = extractIntegrationDeleteParams.Marshal()
//
//    extractIntegrationDeleteRequest, err := UnmarshalExtractIntegrationDeleteRequest(bytes)
//    bytes, err = extractIntegrationDeleteRequest.Marshal()
//
//    extractIntegrationDeleteResponse, err := UnmarshalExtractIntegrationDeleteResponse(bytes)
//    bytes, err = extractIntegrationDeleteResponse.Marshal()
//
//    extractIntegrationFetchParams, err := UnmarshalExtractIntegrationFetchParams(bytes)
//    bytes, err = extractIntegrationFetchParams.Marshal()
//
//    extractIntegrationFetchResponse, err := UnmarshalExtractIntegrationFetchResponse(bytes)
//    bytes, err = extractIntegrationFetchResponse.Marshal()
//
//    extractIntegrationUpdateParams, err := UnmarshalExtractIntegrationUpdateParams(bytes)
//    bytes, err = extractIntegrationUpdateParams.Marshal()
//
//    extractIntegrationUpdateRequest, err := UnmarshalExtractIntegrationUpdateRequest(bytes)
//    bytes, err = extractIntegrationUpdateRequest.Marshal()
//
//    extractIntegrationUpdateResponse, err := UnmarshalExtractIntegrationUpdateResponse(bytes)
//    bytes, err = extractIntegrationUpdateResponse.Marshal()
//
//    extractIntegrationCreateRequest, err := UnmarshalExtractIntegrationCreateRequest(bytes)
//    bytes, err = extractIntegrationCreateRequest.Marshal()
//
//    extractIntegrationCreateResponse, err := UnmarshalExtractIntegrationCreateResponse(bytes)
//    bytes, err = extractIntegrationCreateResponse.Marshal()
//
//    extractIntegrationsListParams, err := UnmarshalExtractIntegrationsListParams(bytes)
//    bytes, err = extractIntegrationsListParams.Marshal()
//
//    extractIntegrationsListResponse, err := UnmarshalExtractIntegrationsListResponse(bytes)
//    bytes, err = extractIntegrationsListResponse.Marshal()
//
//    instagramIntegrationDeleteParams, err := UnmarshalInstagramIntegrationDeleteParams(bytes)
//    bytes, err = instagramIntegrationDeleteParams.Marshal()
//
//    instagramIntegrationDeleteRequest, err := UnmarshalInstagramIntegrationDeleteRequest(bytes)
//    bytes, err = instagramIntegrationDeleteRequest.Marshal()
//
//    instagramIntegrationDeleteResponse, err := UnmarshalInstagramIntegrationDeleteResponse(bytes)
//    bytes, err = instagramIntegrationDeleteResponse.Marshal()
//
//    instagramIntegrationFetchParams, err := UnmarshalInstagramIntegrationFetchParams(bytes)
//    bytes, err = instagramIntegrationFetchParams.Marshal()
//
//    instagramIntegrationFetchResponse, err := UnmarshalInstagramIntegrationFetchResponse(bytes)
//    bytes, err = instagramIntegrationFetchResponse.Marshal()
//
//    instagramIntegrationSetupParams, err := UnmarshalInstagramIntegrationSetupParams(bytes)
//    bytes, err = instagramIntegrationSetupParams.Marshal()
//
//    instagramIntegrationSetupRequest, err := UnmarshalInstagramIntegrationSetupRequest(bytes)
//    bytes, err = instagramIntegrationSetupRequest.Marshal()
//
//    instagramIntegrationSetupResponse, err := UnmarshalInstagramIntegrationSetupResponse(bytes)
//    bytes, err = instagramIntegrationSetupResponse.Marshal()
//
//    instagramIntegrationUpdateParams, err := UnmarshalInstagramIntegrationUpdateParams(bytes)
//    bytes, err = instagramIntegrationUpdateParams.Marshal()
//
//    instagramIntegrationUpdateRequest, err := UnmarshalInstagramIntegrationUpdateRequest(bytes)
//    bytes, err = instagramIntegrationUpdateRequest.Marshal()
//
//    instagramIntegrationUpdateResponse, err := UnmarshalInstagramIntegrationUpdateResponse(bytes)
//    bytes, err = instagramIntegrationUpdateResponse.Marshal()
//
//    instagramIntegrationCreateRequest, err := UnmarshalInstagramIntegrationCreateRequest(bytes)
//    bytes, err = instagramIntegrationCreateRequest.Marshal()
//
//    instagramIntegrationCreateResponse, err := UnmarshalInstagramIntegrationCreateResponse(bytes)
//    bytes, err = instagramIntegrationCreateResponse.Marshal()
//
//    instagramIntegrationsListParams, err := UnmarshalInstagramIntegrationsListParams(bytes)
//    bytes, err = instagramIntegrationsListParams.Marshal()
//
//    instagramIntegrationsListResponse, err := UnmarshalInstagramIntegrationsListResponse(bytes)
//    bytes, err = instagramIntegrationsListResponse.Marshal()
//
//    mCPServerIntegrationDeleteParams, err := UnmarshalMCPServerIntegrationDeleteParams(bytes)
//    bytes, err = mCPServerIntegrationDeleteParams.Marshal()
//
//    mCPServerIntegrationDeleteRequest, err := UnmarshalMCPServerIntegrationDeleteRequest(bytes)
//    bytes, err = mCPServerIntegrationDeleteRequest.Marshal()
//
//    mCPServerIntegrationDeleteResponse, err := UnmarshalMCPServerIntegrationDeleteResponse(bytes)
//    bytes, err = mCPServerIntegrationDeleteResponse.Marshal()
//
//    mCPServerIntegrationFetchParams, err := UnmarshalMCPServerIntegrationFetchParams(bytes)
//    bytes, err = mCPServerIntegrationFetchParams.Marshal()
//
//    mCPServerIntegrationFetchResponse, err := UnmarshalMCPServerIntegrationFetchResponse(bytes)
//    bytes, err = mCPServerIntegrationFetchResponse.Marshal()
//
//    mCPServerIntegrationUpdateParams, err := UnmarshalMCPServerIntegrationUpdateParams(bytes)
//    bytes, err = mCPServerIntegrationUpdateParams.Marshal()
//
//    mCPServerIntegrationUpdateRequest, err := UnmarshalMCPServerIntegrationUpdateRequest(bytes)
//    bytes, err = mCPServerIntegrationUpdateRequest.Marshal()
//
//    mCPServerIntegrationUpdateResponse, err := UnmarshalMCPServerIntegrationUpdateResponse(bytes)
//    bytes, err = mCPServerIntegrationUpdateResponse.Marshal()
//
//    mCPServerIntegrationCreateRequest, err := UnmarshalMCPServerIntegrationCreateRequest(bytes)
//    bytes, err = mCPServerIntegrationCreateRequest.Marshal()
//
//    mCPServerIntegrationCreateResponse, err := UnmarshalMCPServerIntegrationCreateResponse(bytes)
//    bytes, err = mCPServerIntegrationCreateResponse.Marshal()
//
//    mCPServerIntegrationsListParams, err := UnmarshalMCPServerIntegrationsListParams(bytes)
//    bytes, err = mCPServerIntegrationsListParams.Marshal()
//
//    mCPServerIntegrationsListResponse, err := UnmarshalMCPServerIntegrationsListResponse(bytes)
//    bytes, err = mCPServerIntegrationsListResponse.Marshal()
//
//    messengerIntegrationDeleteParams, err := UnmarshalMessengerIntegrationDeleteParams(bytes)
//    bytes, err = messengerIntegrationDeleteParams.Marshal()
//
//    messengerIntegrationDeleteRequest, err := UnmarshalMessengerIntegrationDeleteRequest(bytes)
//    bytes, err = messengerIntegrationDeleteRequest.Marshal()
//
//    messengerIntegrationDeleteResponse, err := UnmarshalMessengerIntegrationDeleteResponse(bytes)
//    bytes, err = messengerIntegrationDeleteResponse.Marshal()
//
//    messengerIntegrationFetchParams, err := UnmarshalMessengerIntegrationFetchParams(bytes)
//    bytes, err = messengerIntegrationFetchParams.Marshal()
//
//    messengerIntegrationFetchResponse, err := UnmarshalMessengerIntegrationFetchResponse(bytes)
//    bytes, err = messengerIntegrationFetchResponse.Marshal()
//
//    messengerIntegrationSetupParams, err := UnmarshalMessengerIntegrationSetupParams(bytes)
//    bytes, err = messengerIntegrationSetupParams.Marshal()
//
//    messengerIntegrationSetupRequest, err := UnmarshalMessengerIntegrationSetupRequest(bytes)
//    bytes, err = messengerIntegrationSetupRequest.Marshal()
//
//    messengerIntegrationSetupResponse, err := UnmarshalMessengerIntegrationSetupResponse(bytes)
//    bytes, err = messengerIntegrationSetupResponse.Marshal()
//
//    messengerIntegrationUpdateParams, err := UnmarshalMessengerIntegrationUpdateParams(bytes)
//    bytes, err = messengerIntegrationUpdateParams.Marshal()
//
//    messengerIntegrationUpdateRequest, err := UnmarshalMessengerIntegrationUpdateRequest(bytes)
//    bytes, err = messengerIntegrationUpdateRequest.Marshal()
//
//    messengerIntegrationUpdateResponse, err := UnmarshalMessengerIntegrationUpdateResponse(bytes)
//    bytes, err = messengerIntegrationUpdateResponse.Marshal()
//
//    messengerIntegrationCreateRequest, err := UnmarshalMessengerIntegrationCreateRequest(bytes)
//    bytes, err = messengerIntegrationCreateRequest.Marshal()
//
//    messengerIntegrationCreateResponse, err := UnmarshalMessengerIntegrationCreateResponse(bytes)
//    bytes, err = messengerIntegrationCreateResponse.Marshal()
//
//    messengerIntegrationsListParams, err := UnmarshalMessengerIntegrationsListParams(bytes)
//    bytes, err = messengerIntegrationsListParams.Marshal()
//
//    messengerIntegrationsListResponse, err := UnmarshalMessengerIntegrationsListResponse(bytes)
//    bytes, err = messengerIntegrationsListResponse.Marshal()
//
//    notionIntegrationDeleteParams, err := UnmarshalNotionIntegrationDeleteParams(bytes)
//    bytes, err = notionIntegrationDeleteParams.Marshal()
//
//    notionIntegrationDeleteRequest, err := UnmarshalNotionIntegrationDeleteRequest(bytes)
//    bytes, err = notionIntegrationDeleteRequest.Marshal()
//
//    notionIntegrationDeleteResponse, err := UnmarshalNotionIntegrationDeleteResponse(bytes)
//    bytes, err = notionIntegrationDeleteResponse.Marshal()
//
//    notionIntegrationFetchParams, err := UnmarshalNotionIntegrationFetchParams(bytes)
//    bytes, err = notionIntegrationFetchParams.Marshal()
//
//    notionIntegrationFetchResponse, err := UnmarshalNotionIntegrationFetchResponse(bytes)
//    bytes, err = notionIntegrationFetchResponse.Marshal()
//
//    notionIntegrationSyncParams, err := UnmarshalNotionIntegrationSyncParams(bytes)
//    bytes, err = notionIntegrationSyncParams.Marshal()
//
//    notionIntegrationSyncRequest, err := UnmarshalNotionIntegrationSyncRequest(bytes)
//    bytes, err = notionIntegrationSyncRequest.Marshal()
//
//    notionIntegrationSyncResponse, err := UnmarshalNotionIntegrationSyncResponse(bytes)
//    bytes, err = notionIntegrationSyncResponse.Marshal()
//
//    notionIntegrationUpdateParams, err := UnmarshalNotionIntegrationUpdateParams(bytes)
//    bytes, err = notionIntegrationUpdateParams.Marshal()
//
//    notionIntegrationUpdateRequest, err := UnmarshalNotionIntegrationUpdateRequest(bytes)
//    bytes, err = notionIntegrationUpdateRequest.Marshal()
//
//    notionIntegrationUpdateResponse, err := UnmarshalNotionIntegrationUpdateResponse(bytes)
//    bytes, err = notionIntegrationUpdateResponse.Marshal()
//
//    notionIntegrationCreateRequest, err := UnmarshalNotionIntegrationCreateRequest(bytes)
//    bytes, err = notionIntegrationCreateRequest.Marshal()
//
//    notionIntegrationCreateResponse, err := UnmarshalNotionIntegrationCreateResponse(bytes)
//    bytes, err = notionIntegrationCreateResponse.Marshal()
//
//    notionIntegrationsListParams, err := UnmarshalNotionIntegrationsListParams(bytes)
//    bytes, err = notionIntegrationsListParams.Marshal()
//
//    notionIntegrationsListResponse, err := UnmarshalNotionIntegrationsListResponse(bytes)
//    bytes, err = notionIntegrationsListResponse.Marshal()
//
//    sitemapIntegrationDeleteParams, err := UnmarshalSitemapIntegrationDeleteParams(bytes)
//    bytes, err = sitemapIntegrationDeleteParams.Marshal()
//
//    sitemapIntegrationDeleteRequest, err := UnmarshalSitemapIntegrationDeleteRequest(bytes)
//    bytes, err = sitemapIntegrationDeleteRequest.Marshal()
//
//    sitemapIntegrationDeleteResponse, err := UnmarshalSitemapIntegrationDeleteResponse(bytes)
//    bytes, err = sitemapIntegrationDeleteResponse.Marshal()
//
//    sitemapIntegrationFetchParams, err := UnmarshalSitemapIntegrationFetchParams(bytes)
//    bytes, err = sitemapIntegrationFetchParams.Marshal()
//
//    sitemapIntegrationFetchResponse, err := UnmarshalSitemapIntegrationFetchResponse(bytes)
//    bytes, err = sitemapIntegrationFetchResponse.Marshal()
//
//    sitemapIntegrationSyncParams, err := UnmarshalSitemapIntegrationSyncParams(bytes)
//    bytes, err = sitemapIntegrationSyncParams.Marshal()
//
//    sitemapIntegrationSyncRequest, err := UnmarshalSitemapIntegrationSyncRequest(bytes)
//    bytes, err = sitemapIntegrationSyncRequest.Marshal()
//
//    sitemapIntegrationSyncResponse, err := UnmarshalSitemapIntegrationSyncResponse(bytes)
//    bytes, err = sitemapIntegrationSyncResponse.Marshal()
//
//    sitemapIntegrationUpdateParams, err := UnmarshalSitemapIntegrationUpdateParams(bytes)
//    bytes, err = sitemapIntegrationUpdateParams.Marshal()
//
//    sitemapIntegrationUpdateRequest, err := UnmarshalSitemapIntegrationUpdateRequest(bytes)
//    bytes, err = sitemapIntegrationUpdateRequest.Marshal()
//
//    sitemapIntegrationUpdateResponse, err := UnmarshalSitemapIntegrationUpdateResponse(bytes)
//    bytes, err = sitemapIntegrationUpdateResponse.Marshal()
//
//    sitemapIntegrationCreateRequest, err := UnmarshalSitemapIntegrationCreateRequest(bytes)
//    bytes, err = sitemapIntegrationCreateRequest.Marshal()
//
//    sitemapIntegrationCreateResponse, err := UnmarshalSitemapIntegrationCreateResponse(bytes)
//    bytes, err = sitemapIntegrationCreateResponse.Marshal()
//
//    sitemapIntegrationsListParams, err := UnmarshalSitemapIntegrationsListParams(bytes)
//    bytes, err = sitemapIntegrationsListParams.Marshal()
//
//    sitemapIntegrationsListResponse, err := UnmarshalSitemapIntegrationsListResponse(bytes)
//    bytes, err = sitemapIntegrationsListResponse.Marshal()
//
//    slackIntegrationDeleteParams, err := UnmarshalSlackIntegrationDeleteParams(bytes)
//    bytes, err = slackIntegrationDeleteParams.Marshal()
//
//    slackIntegrationDeleteRequest, err := UnmarshalSlackIntegrationDeleteRequest(bytes)
//    bytes, err = slackIntegrationDeleteRequest.Marshal()
//
//    slackIntegrationDeleteResponse, err := UnmarshalSlackIntegrationDeleteResponse(bytes)
//    bytes, err = slackIntegrationDeleteResponse.Marshal()
//
//    slackIntegrationFetchParams, err := UnmarshalSlackIntegrationFetchParams(bytes)
//    bytes, err = slackIntegrationFetchParams.Marshal()
//
//    slackIntegrationFetchResponse, err := UnmarshalSlackIntegrationFetchResponse(bytes)
//    bytes, err = slackIntegrationFetchResponse.Marshal()
//
//    slackIntegrationSetupParams, err := UnmarshalSlackIntegrationSetupParams(bytes)
//    bytes, err = slackIntegrationSetupParams.Marshal()
//
//    slackIntegrationSetupRequest, err := UnmarshalSlackIntegrationSetupRequest(bytes)
//    bytes, err = slackIntegrationSetupRequest.Marshal()
//
//    slackIntegrationSetupResponse, err := UnmarshalSlackIntegrationSetupResponse(bytes)
//    bytes, err = slackIntegrationSetupResponse.Marshal()
//
//    slackIntegrationUpdateParams, err := UnmarshalSlackIntegrationUpdateParams(bytes)
//    bytes, err = slackIntegrationUpdateParams.Marshal()
//
//    slackIntegrationUpdateRequest, err := UnmarshalSlackIntegrationUpdateRequest(bytes)
//    bytes, err = slackIntegrationUpdateRequest.Marshal()
//
//    slackIntegrationUpdateResponse, err := UnmarshalSlackIntegrationUpdateResponse(bytes)
//    bytes, err = slackIntegrationUpdateResponse.Marshal()
//
//    slackIntegrationCreateRequest, err := UnmarshalSlackIntegrationCreateRequest(bytes)
//    bytes, err = slackIntegrationCreateRequest.Marshal()
//
//    slackIntegrationCreateResponse, err := UnmarshalSlackIntegrationCreateResponse(bytes)
//    bytes, err = slackIntegrationCreateResponse.Marshal()
//
//    slackIntegrationsListParams, err := UnmarshalSlackIntegrationsListParams(bytes)
//    bytes, err = slackIntegrationsListParams.Marshal()
//
//    slackIntegrationsListResponse, err := UnmarshalSlackIntegrationsListResponse(bytes)
//    bytes, err = slackIntegrationsListResponse.Marshal()
//
//    supportIntegrationDeleteParams, err := UnmarshalSupportIntegrationDeleteParams(bytes)
//    bytes, err = supportIntegrationDeleteParams.Marshal()
//
//    supportIntegrationDeleteRequest, err := UnmarshalSupportIntegrationDeleteRequest(bytes)
//    bytes, err = supportIntegrationDeleteRequest.Marshal()
//
//    supportIntegrationDeleteResponse, err := UnmarshalSupportIntegrationDeleteResponse(bytes)
//    bytes, err = supportIntegrationDeleteResponse.Marshal()
//
//    supportIntegrationFetchParams, err := UnmarshalSupportIntegrationFetchParams(bytes)
//    bytes, err = supportIntegrationFetchParams.Marshal()
//
//    supportIntegrationFetchResponse, err := UnmarshalSupportIntegrationFetchResponse(bytes)
//    bytes, err = supportIntegrationFetchResponse.Marshal()
//
//    supportIntegrationUpdateParams, err := UnmarshalSupportIntegrationUpdateParams(bytes)
//    bytes, err = supportIntegrationUpdateParams.Marshal()
//
//    supportIntegrationUpdateRequest, err := UnmarshalSupportIntegrationUpdateRequest(bytes)
//    bytes, err = supportIntegrationUpdateRequest.Marshal()
//
//    supportIntegrationUpdateResponse, err := UnmarshalSupportIntegrationUpdateResponse(bytes)
//    bytes, err = supportIntegrationUpdateResponse.Marshal()
//
//    supportIntegrationCreateRequest, err := UnmarshalSupportIntegrationCreateRequest(bytes)
//    bytes, err = supportIntegrationCreateRequest.Marshal()
//
//    supportIntegrationCreateResponse, err := UnmarshalSupportIntegrationCreateResponse(bytes)
//    bytes, err = supportIntegrationCreateResponse.Marshal()
//
//    supportIntegrationsListParams, err := UnmarshalSupportIntegrationsListParams(bytes)
//    bytes, err = supportIntegrationsListParams.Marshal()
//
//    supportIntegrationsListResponse, err := UnmarshalSupportIntegrationsListResponse(bytes)
//    bytes, err = supportIntegrationsListResponse.Marshal()
//
//    telegramIntegrationDeleteParams, err := UnmarshalTelegramIntegrationDeleteParams(bytes)
//    bytes, err = telegramIntegrationDeleteParams.Marshal()
//
//    telegramIntegrationDeleteRequest, err := UnmarshalTelegramIntegrationDeleteRequest(bytes)
//    bytes, err = telegramIntegrationDeleteRequest.Marshal()
//
//    telegramIntegrationDeleteResponse, err := UnmarshalTelegramIntegrationDeleteResponse(bytes)
//    bytes, err = telegramIntegrationDeleteResponse.Marshal()
//
//    telegramIntegrationFetchParams, err := UnmarshalTelegramIntegrationFetchParams(bytes)
//    bytes, err = telegramIntegrationFetchParams.Marshal()
//
//    telegramIntegrationFetchResponse, err := UnmarshalTelegramIntegrationFetchResponse(bytes)
//    bytes, err = telegramIntegrationFetchResponse.Marshal()
//
//    telegramIntegrationSetupParams, err := UnmarshalTelegramIntegrationSetupParams(bytes)
//    bytes, err = telegramIntegrationSetupParams.Marshal()
//
//    telegramIntegrationSetupRequest, err := UnmarshalTelegramIntegrationSetupRequest(bytes)
//    bytes, err = telegramIntegrationSetupRequest.Marshal()
//
//    telegramIntegrationSetupResponse, err := UnmarshalTelegramIntegrationSetupResponse(bytes)
//    bytes, err = telegramIntegrationSetupResponse.Marshal()
//
//    telegramIntegrationUpdateParams, err := UnmarshalTelegramIntegrationUpdateParams(bytes)
//    bytes, err = telegramIntegrationUpdateParams.Marshal()
//
//    telegramIntegrationUpdateRequest, err := UnmarshalTelegramIntegrationUpdateRequest(bytes)
//    bytes, err = telegramIntegrationUpdateRequest.Marshal()
//
//    telegramIntegrationUpdateResponse, err := UnmarshalTelegramIntegrationUpdateResponse(bytes)
//    bytes, err = telegramIntegrationUpdateResponse.Marshal()
//
//    telegramIntegrationCreateRequest, err := UnmarshalTelegramIntegrationCreateRequest(bytes)
//    bytes, err = telegramIntegrationCreateRequest.Marshal()
//
//    telegramIntegrationCreateResponse, err := UnmarshalTelegramIntegrationCreateResponse(bytes)
//    bytes, err = telegramIntegrationCreateResponse.Marshal()
//
//    telegramIntegrationsListParams, err := UnmarshalTelegramIntegrationsListParams(bytes)
//    bytes, err = telegramIntegrationsListParams.Marshal()
//
//    telegramIntegrationsListResponse, err := UnmarshalTelegramIntegrationsListResponse(bytes)
//    bytes, err = telegramIntegrationsListResponse.Marshal()
//
//    triggerIntegrationDeleteParams, err := UnmarshalTriggerIntegrationDeleteParams(bytes)
//    bytes, err = triggerIntegrationDeleteParams.Marshal()
//
//    triggerIntegrationDeleteRequest, err := UnmarshalTriggerIntegrationDeleteRequest(bytes)
//    bytes, err = triggerIntegrationDeleteRequest.Marshal()
//
//    triggerIntegrationDeleteResponse, err := UnmarshalTriggerIntegrationDeleteResponse(bytes)
//    bytes, err = triggerIntegrationDeleteResponse.Marshal()
//
//    triggerIntegrationFetchParams, err := UnmarshalTriggerIntegrationFetchParams(bytes)
//    bytes, err = triggerIntegrationFetchParams.Marshal()
//
//    triggerIntegrationFetchResponse, err := UnmarshalTriggerIntegrationFetchResponse(bytes)
//    bytes, err = triggerIntegrationFetchResponse.Marshal()
//
//    triggerIntegrationInvokeParams, err := UnmarshalTriggerIntegrationInvokeParams(bytes)
//    bytes, err = triggerIntegrationInvokeParams.Marshal()
//
//    triggerIntegrationInvokeRequest, err := UnmarshalTriggerIntegrationInvokeRequest(bytes)
//    bytes, err = triggerIntegrationInvokeRequest.Marshal()
//
//    triggerIntegrationInvokeResponse, err := UnmarshalTriggerIntegrationInvokeResponse(bytes)
//    bytes, err = triggerIntegrationInvokeResponse.Marshal()
//
//    triggerIntegrationSetupParams, err := UnmarshalTriggerIntegrationSetupParams(bytes)
//    bytes, err = triggerIntegrationSetupParams.Marshal()
//
//    triggerIntegrationSetupRequest, err := UnmarshalTriggerIntegrationSetupRequest(bytes)
//    bytes, err = triggerIntegrationSetupRequest.Marshal()
//
//    triggerIntegrationSetupResponse, err := UnmarshalTriggerIntegrationSetupResponse(bytes)
//    bytes, err = triggerIntegrationSetupResponse.Marshal()
//
//    triggerIntegrationUpdateParams, err := UnmarshalTriggerIntegrationUpdateParams(bytes)
//    bytes, err = triggerIntegrationUpdateParams.Marshal()
//
//    triggerIntegrationUpdateRequest, err := UnmarshalTriggerIntegrationUpdateRequest(bytes)
//    bytes, err = triggerIntegrationUpdateRequest.Marshal()
//
//    triggerIntegrationUpdateResponse, err := UnmarshalTriggerIntegrationUpdateResponse(bytes)
//    bytes, err = triggerIntegrationUpdateResponse.Marshal()
//
//    triggerIntegrationCreateRequest, err := UnmarshalTriggerIntegrationCreateRequest(bytes)
//    bytes, err = triggerIntegrationCreateRequest.Marshal()
//
//    triggerIntegrationCreateResponse, err := UnmarshalTriggerIntegrationCreateResponse(bytes)
//    bytes, err = triggerIntegrationCreateResponse.Marshal()
//
//    triggerIntegrationsListParams, err := UnmarshalTriggerIntegrationsListParams(bytes)
//    bytes, err = triggerIntegrationsListParams.Marshal()
//
//    triggerIntegrationsListResponse, err := UnmarshalTriggerIntegrationsListResponse(bytes)
//    bytes, err = triggerIntegrationsListResponse.Marshal()
//
//    twilioIntegrationDeleteParams, err := UnmarshalTwilioIntegrationDeleteParams(bytes)
//    bytes, err = twilioIntegrationDeleteParams.Marshal()
//
//    twilioIntegrationDeleteRequest, err := UnmarshalTwilioIntegrationDeleteRequest(bytes)
//    bytes, err = twilioIntegrationDeleteRequest.Marshal()
//
//    twilioIntegrationDeleteResponse, err := UnmarshalTwilioIntegrationDeleteResponse(bytes)
//    bytes, err = twilioIntegrationDeleteResponse.Marshal()
//
//    twilioIntegrationFetchParams, err := UnmarshalTwilioIntegrationFetchParams(bytes)
//    bytes, err = twilioIntegrationFetchParams.Marshal()
//
//    twilioIntegrationFetchResponse, err := UnmarshalTwilioIntegrationFetchResponse(bytes)
//    bytes, err = twilioIntegrationFetchResponse.Marshal()
//
//    twilioIntegrationSetupParams, err := UnmarshalTwilioIntegrationSetupParams(bytes)
//    bytes, err = twilioIntegrationSetupParams.Marshal()
//
//    twilioIntegrationSetupRequest, err := UnmarshalTwilioIntegrationSetupRequest(bytes)
//    bytes, err = twilioIntegrationSetupRequest.Marshal()
//
//    twilioIntegrationSetupResponse, err := UnmarshalTwilioIntegrationSetupResponse(bytes)
//    bytes, err = twilioIntegrationSetupResponse.Marshal()
//
//    twilioIntegrationUpdateParams, err := UnmarshalTwilioIntegrationUpdateParams(bytes)
//    bytes, err = twilioIntegrationUpdateParams.Marshal()
//
//    twilioIntegrationUpdateRequest, err := UnmarshalTwilioIntegrationUpdateRequest(bytes)
//    bytes, err = twilioIntegrationUpdateRequest.Marshal()
//
//    twilioIntegrationUpdateResponse, err := UnmarshalTwilioIntegrationUpdateResponse(bytes)
//    bytes, err = twilioIntegrationUpdateResponse.Marshal()
//
//    twilioIntegrationCreateRequest, err := UnmarshalTwilioIntegrationCreateRequest(bytes)
//    bytes, err = twilioIntegrationCreateRequest.Marshal()
//
//    twilioIntegrationCreateResponse, err := UnmarshalTwilioIntegrationCreateResponse(bytes)
//    bytes, err = twilioIntegrationCreateResponse.Marshal()
//
//    twilioIntegrationsListParams, err := UnmarshalTwilioIntegrationsListParams(bytes)
//    bytes, err = twilioIntegrationsListParams.Marshal()
//
//    twilioIntegrationsListResponse, err := UnmarshalTwilioIntegrationsListResponse(bytes)
//    bytes, err = twilioIntegrationsListResponse.Marshal()
//
//    whatsAppIntegrationDeleteParams, err := UnmarshalWhatsAppIntegrationDeleteParams(bytes)
//    bytes, err = whatsAppIntegrationDeleteParams.Marshal()
//
//    whatsAppIntegrationDeleteRequest, err := UnmarshalWhatsAppIntegrationDeleteRequest(bytes)
//    bytes, err = whatsAppIntegrationDeleteRequest.Marshal()
//
//    whatsAppIntegrationDeleteResponse, err := UnmarshalWhatsAppIntegrationDeleteResponse(bytes)
//    bytes, err = whatsAppIntegrationDeleteResponse.Marshal()
//
//    whatsAppIntegrationFetchParams, err := UnmarshalWhatsAppIntegrationFetchParams(bytes)
//    bytes, err = whatsAppIntegrationFetchParams.Marshal()
//
//    whatsAppIntegrationFetchResponse, err := UnmarshalWhatsAppIntegrationFetchResponse(bytes)
//    bytes, err = whatsAppIntegrationFetchResponse.Marshal()
//
//    whatsAppIntegrationSetupParams, err := UnmarshalWhatsAppIntegrationSetupParams(bytes)
//    bytes, err = whatsAppIntegrationSetupParams.Marshal()
//
//    whatsAppIntegrationSetupRequest, err := UnmarshalWhatsAppIntegrationSetupRequest(bytes)
//    bytes, err = whatsAppIntegrationSetupRequest.Marshal()
//
//    whatsAppIntegrationSetupResponse, err := UnmarshalWhatsAppIntegrationSetupResponse(bytes)
//    bytes, err = whatsAppIntegrationSetupResponse.Marshal()
//
//    whatsAppIntegrationUpdateParams, err := UnmarshalWhatsAppIntegrationUpdateParams(bytes)
//    bytes, err = whatsAppIntegrationUpdateParams.Marshal()
//
//    whatsAppIntegrationUpdateRequest, err := UnmarshalWhatsAppIntegrationUpdateRequest(bytes)
//    bytes, err = whatsAppIntegrationUpdateRequest.Marshal()
//
//    whatsAppIntegrationUpdateResponse, err := UnmarshalWhatsAppIntegrationUpdateResponse(bytes)
//    bytes, err = whatsAppIntegrationUpdateResponse.Marshal()
//
//    whatsAppIntegrationCreateRequest, err := UnmarshalWhatsAppIntegrationCreateRequest(bytes)
//    bytes, err = whatsAppIntegrationCreateRequest.Marshal()
//
//    whatsAppIntegrationCreateResponse, err := UnmarshalWhatsAppIntegrationCreateResponse(bytes)
//    bytes, err = whatsAppIntegrationCreateResponse.Marshal()
//
//    whatsAppIntegrationsListParams, err := UnmarshalWhatsAppIntegrationsListParams(bytes)
//    bytes, err = whatsAppIntegrationsListParams.Marshal()
//
//    whatsAppIntegrationsListResponse, err := UnmarshalWhatsAppIntegrationsListResponse(bytes)
//    bytes, err = whatsAppIntegrationsListResponse.Marshal()
//
//    widgetIntegrationDeleteParams, err := UnmarshalWidgetIntegrationDeleteParams(bytes)
//    bytes, err = widgetIntegrationDeleteParams.Marshal()
//
//    widgetIntegrationDeleteRequest, err := UnmarshalWidgetIntegrationDeleteRequest(bytes)
//    bytes, err = widgetIntegrationDeleteRequest.Marshal()
//
//    widgetIntegrationDeleteResponse, err := UnmarshalWidgetIntegrationDeleteResponse(bytes)
//    bytes, err = widgetIntegrationDeleteResponse.Marshal()
//
//    widgetIntegrationFetchParams, err := UnmarshalWidgetIntegrationFetchParams(bytes)
//    bytes, err = widgetIntegrationFetchParams.Marshal()
//
//    widgetIntegrationFetchResponse, err := UnmarshalWidgetIntegrationFetchResponse(bytes)
//    bytes, err = widgetIntegrationFetchResponse.Marshal()
//
//    widgetIntegrationSetupParams, err := UnmarshalWidgetIntegrationSetupParams(bytes)
//    bytes, err = widgetIntegrationSetupParams.Marshal()
//
//    widgetIntegrationSetupRequest, err := UnmarshalWidgetIntegrationSetupRequest(bytes)
//    bytes, err = widgetIntegrationSetupRequest.Marshal()
//
//    widgetIntegrationSetupResponse, err := UnmarshalWidgetIntegrationSetupResponse(bytes)
//    bytes, err = widgetIntegrationSetupResponse.Marshal()
//
//    widgetIntegrationUpdateParams, err := UnmarshalWidgetIntegrationUpdateParams(bytes)
//    bytes, err = widgetIntegrationUpdateParams.Marshal()
//
//    widgetIntegrationUpdateRequest, err := UnmarshalWidgetIntegrationUpdateRequest(bytes)
//    bytes, err = widgetIntegrationUpdateRequest.Marshal()
//
//    widgetIntegrationUpdateResponse, err := UnmarshalWidgetIntegrationUpdateResponse(bytes)
//    bytes, err = widgetIntegrationUpdateResponse.Marshal()
//
//    widgetIntegrationCreateRequest, err := UnmarshalWidgetIntegrationCreateRequest(bytes)
//    bytes, err = widgetIntegrationCreateRequest.Marshal()
//
//    widgetIntegrationCreateResponse, err := UnmarshalWidgetIntegrationCreateResponse(bytes)
//    bytes, err = widgetIntegrationCreateResponse.Marshal()
//
//    widgetIntegrationsListParams, err := UnmarshalWidgetIntegrationsListParams(bytes)
//    bytes, err = widgetIntegrationsListParams.Marshal()
//
//    widgetIntegrationsListResponse, err := UnmarshalWidgetIntegrationsListResponse(bytes)
//    bytes, err = widgetIntegrationsListResponse.Marshal()
//
//    magicFromPromptGenerateParams, err := UnmarshalMagicFromPromptGenerateParams(bytes)
//    bytes, err = magicFromPromptGenerateParams.Marshal()
//
//    magicFromPromptGenerateRequest, err := UnmarshalMagicFromPromptGenerateRequest(bytes)
//    bytes, err = magicFromPromptGenerateRequest.Marshal()
//
//    magicFromPromptGenerateResponse, err := UnmarshalMagicFromPromptGenerateResponse(bytes)
//    bytes, err = magicFromPromptGenerateResponse.Marshal()
//
//    magicPromptsListParams, err := UnmarshalMagicPromptsListParams(bytes)
//    bytes, err = magicPromptsListParams.Marshal()
//
//    magicPromptsListResponse, err := UnmarshalMagicPromptsListResponse(bytes)
//    bytes, err = magicPromptsListResponse.Marshal()
//
//    memoryDeleteParams, err := UnmarshalMemoryDeleteParams(bytes)
//    bytes, err = memoryDeleteParams.Marshal()
//
//    memoryDeleteRequest, err := UnmarshalMemoryDeleteRequest(bytes)
//    bytes, err = memoryDeleteRequest.Marshal()
//
//    memoryDeleteResponse, err := UnmarshalMemoryDeleteResponse(bytes)
//    bytes, err = memoryDeleteResponse.Marshal()
//
//    memoryFetchParams, err := UnmarshalMemoryFetchParams(bytes)
//    bytes, err = memoryFetchParams.Marshal()
//
//    memoryFetchResponse, err := UnmarshalMemoryFetchResponse(bytes)
//    bytes, err = memoryFetchResponse.Marshal()
//
//    memoryUpdateParams, err := UnmarshalMemoryUpdateParams(bytes)
//    bytes, err = memoryUpdateParams.Marshal()
//
//    memoryUpdateRequest, err := UnmarshalMemoryUpdateRequest(bytes)
//    bytes, err = memoryUpdateRequest.Marshal()
//
//    memoryUpdateResponse, err := UnmarshalMemoryUpdateResponse(bytes)
//    bytes, err = memoryUpdateResponse.Marshal()
//
//    memoryCreateRequest, err := UnmarshalMemoryCreateRequest(bytes)
//    bytes, err = memoryCreateRequest.Marshal()
//
//    memoryCreateResponse, err := UnmarshalMemoryCreateResponse(bytes)
//    bytes, err = memoryCreateResponse.Marshal()
//
//    memoriesExportParams, err := UnmarshalMemoriesExportParams(bytes)
//    bytes, err = memoriesExportParams.Marshal()
//
//    memoriesExportResponse, err := UnmarshalMemoriesExportResponse(bytes)
//    bytes, err = memoriesExportResponse.Marshal()
//
//    memoriesListParams, err := UnmarshalMemoriesListParams(bytes)
//    bytes, err = memoriesListParams.Marshal()
//
//    memoriesListResponse, err := UnmarshalMemoriesListResponse(bytes)
//    bytes, err = memoriesListResponse.Marshal()
//
//    memorySearchRequest, err := UnmarshalMemorySearchRequest(bytes)
//    bytes, err = memorySearchRequest.Marshal()
//
//    memorySearchResponse, err := UnmarshalMemorySearchResponse(bytes)
//    bytes, err = memorySearchResponse.Marshal()
//
//    partnerUserDeleteParams, err := UnmarshalPartnerUserDeleteParams(bytes)
//    bytes, err = partnerUserDeleteParams.Marshal()
//
//    partnerUserDeleteRequest, err := UnmarshalPartnerUserDeleteRequest(bytes)
//    bytes, err = partnerUserDeleteRequest.Marshal()
//
//    partnerUserDeleteResponse, err := UnmarshalPartnerUserDeleteResponse(bytes)
//    bytes, err = partnerUserDeleteResponse.Marshal()
//
//    partnerUserFetchParams, err := UnmarshalPartnerUserFetchParams(bytes)
//    bytes, err = partnerUserFetchParams.Marshal()
//
//    partnerUserFetchResponse, err := UnmarshalPartnerUserFetchResponse(bytes)
//    bytes, err = partnerUserFetchResponse.Marshal()
//
//    partnerUserTokenDeleteParams, err := UnmarshalPartnerUserTokenDeleteParams(bytes)
//    bytes, err = partnerUserTokenDeleteParams.Marshal()
//
//    partnerUserTokenDeleteRequest, err := UnmarshalPartnerUserTokenDeleteRequest(bytes)
//    bytes, err = partnerUserTokenDeleteRequest.Marshal()
//
//    partnerUserTokenDeleteResponse, err := UnmarshalPartnerUserTokenDeleteResponse(bytes)
//    bytes, err = partnerUserTokenDeleteResponse.Marshal()
//
//    partnerUserTokenCreateParams, err := UnmarshalPartnerUserTokenCreateParams(bytes)
//    bytes, err = partnerUserTokenCreateParams.Marshal()
//
//    partnerUserTokenCreateRequest, err := UnmarshalPartnerUserTokenCreateRequest(bytes)
//    bytes, err = partnerUserTokenCreateRequest.Marshal()
//
//    partnerUserTokenCreateResponse, err := UnmarshalPartnerUserTokenCreateResponse(bytes)
//    bytes, err = partnerUserTokenCreateResponse.Marshal()
//
//    partnerUserTokensListParams, err := UnmarshalPartnerUserTokensListParams(bytes)
//    bytes, err = partnerUserTokensListParams.Marshal()
//
//    partnerUserTokensListResponse, err := UnmarshalPartnerUserTokensListResponse(bytes)
//    bytes, err = partnerUserTokensListResponse.Marshal()
//
//    partnerUserUpdateParams, err := UnmarshalPartnerUserUpdateParams(bytes)
//    bytes, err = partnerUserUpdateParams.Marshal()
//
//    partnerUserUpdateRequest, err := UnmarshalPartnerUserUpdateRequest(bytes)
//    bytes, err = partnerUserUpdateRequest.Marshal()
//
//    partnerUserUpdateResponse, err := UnmarshalPartnerUserUpdateResponse(bytes)
//    bytes, err = partnerUserUpdateResponse.Marshal()
//
//    partnerUserCreateRequest, err := UnmarshalPartnerUserCreateRequest(bytes)
//    bytes, err = partnerUserCreateRequest.Marshal()
//
//    partnerUserCreateResponse, err := UnmarshalPartnerUserCreateResponse(bytes)
//    bytes, err = partnerUserCreateResponse.Marshal()
//
//    partnerUsersListParams, err := UnmarshalPartnerUsersListParams(bytes)
//    bytes, err = partnerUsersListParams.Marshal()
//
//    partnerUsersListResponse, err := UnmarshalPartnerUsersListResponse(bytes)
//    bytes, err = partnerUsersListResponse.Marshal()
//
//    platformAbilitiesListParams, err := UnmarshalPlatformAbilitiesListParams(bytes)
//    bytes, err = platformAbilitiesListParams.Marshal()
//
//    platformAbilitiesListResponse, err := UnmarshalPlatformAbilitiesListResponse(bytes)
//    bytes, err = platformAbilitiesListResponse.Marshal()
//
//    platformActionsListParams, err := UnmarshalPlatformActionsListParams(bytes)
//    bytes, err = platformActionsListParams.Marshal()
//
//    platformActionsListResponse, err := UnmarshalPlatformActionsListResponse(bytes)
//    bytes, err = platformActionsListResponse.Marshal()
//
//    platformDocFetchParams, err := UnmarshalPlatformDocFetchParams(bytes)
//    bytes, err = platformDocFetchParams.Marshal()
//
//    platformDocFetchResponse, err := UnmarshalPlatformDocFetchResponse(bytes)
//    bytes, err = platformDocFetchResponse.Marshal()
//
//    platformDocsListParams, err := UnmarshalPlatformDocsListParams(bytes)
//    bytes, err = platformDocsListParams.Marshal()
//
//    platformDocsListResponse, err := UnmarshalPlatformDocsListResponse(bytes)
//    bytes, err = platformDocsListResponse.Marshal()
//
//    platformDocsSearchRequest, err := UnmarshalPlatformDocsSearchRequest(bytes)
//    bytes, err = platformDocsSearchRequest.Marshal()
//
//    platformDocsSearchResponse, err := UnmarshalPlatformDocsSearchResponse(bytes)
//    bytes, err = platformDocsSearchResponse.Marshal()
//
//    platformExampleCloneParams, err := UnmarshalPlatformExampleCloneParams(bytes)
//    bytes, err = platformExampleCloneParams.Marshal()
//
//    platformExampleCloneRequest, err := UnmarshalPlatformExampleCloneRequest(bytes)
//    bytes, err = platformExampleCloneRequest.Marshal()
//
//    platformExampleCloneResponse, err := UnmarshalPlatformExampleCloneResponse(bytes)
//    bytes, err = platformExampleCloneResponse.Marshal()
//
//    platformExampleFetchParams, err := UnmarshalPlatformExampleFetchParams(bytes)
//    bytes, err = platformExampleFetchParams.Marshal()
//
//    platformExampleFetchResponse, err := UnmarshalPlatformExampleFetchResponse(bytes)
//    bytes, err = platformExampleFetchResponse.Marshal()
//
//    platformExamplesListParams, err := UnmarshalPlatformExamplesListParams(bytes)
//    bytes, err = platformExamplesListParams.Marshal()
//
//    platformExamplesListResponse, err := UnmarshalPlatformExamplesListResponse(bytes)
//    bytes, err = platformExamplesListResponse.Marshal()
//
//    platformExamplesSearchRequest, err := UnmarshalPlatformExamplesSearchRequest(bytes)
//    bytes, err = platformExamplesSearchRequest.Marshal()
//
//    platformExamplesSearchResponse, err := UnmarshalPlatformExamplesSearchResponse(bytes)
//    bytes, err = platformExamplesSearchResponse.Marshal()
//
//    platformGuideFetchParams, err := UnmarshalPlatformGuideFetchParams(bytes)
//    bytes, err = platformGuideFetchParams.Marshal()
//
//    platformGuideFetchResponse, err := UnmarshalPlatformGuideFetchResponse(bytes)
//    bytes, err = platformGuideFetchResponse.Marshal()
//
//    platformGuidesListParams, err := UnmarshalPlatformGuidesListParams(bytes)
//    bytes, err = platformGuidesListParams.Marshal()
//
//    platformGuidesListResponse, err := UnmarshalPlatformGuidesListResponse(bytes)
//    bytes, err = platformGuidesListResponse.Marshal()
//
//    platformGuidesSearchRequest, err := UnmarshalPlatformGuidesSearchRequest(bytes)
//    bytes, err = platformGuidesSearchRequest.Marshal()
//
//    platformGuidesSearchResponse, err := UnmarshalPlatformGuidesSearchResponse(bytes)
//    bytes, err = platformGuidesSearchResponse.Marshal()
//
//    platformManualFetchParams, err := UnmarshalPlatformManualFetchParams(bytes)
//    bytes, err = platformManualFetchParams.Marshal()
//
//    platformManualFetchResponse, err := UnmarshalPlatformManualFetchResponse(bytes)
//    bytes, err = platformManualFetchResponse.Marshal()
//
//    platformManualsListParams, err := UnmarshalPlatformManualsListParams(bytes)
//    bytes, err = platformManualsListParams.Marshal()
//
//    platformManualsListResponse, err := UnmarshalPlatformManualsListResponse(bytes)
//    bytes, err = platformManualsListResponse.Marshal()
//
//    platformManualsSearchRequest, err := UnmarshalPlatformManualsSearchRequest(bytes)
//    bytes, err = platformManualsSearchRequest.Marshal()
//
//    platformManualsSearchResponse, err := UnmarshalPlatformManualsSearchResponse(bytes)
//    bytes, err = platformManualsSearchResponse.Marshal()
//
//    platformModelsListParams, err := UnmarshalPlatformModelsListParams(bytes)
//    bytes, err = platformModelsListParams.Marshal()
//
//    platformModelsListResponse, err := UnmarshalPlatformModelsListResponse(bytes)
//    bytes, err = platformModelsListResponse.Marshal()
//
//    platformSecretsListParams, err := UnmarshalPlatformSecretsListParams(bytes)
//    bytes, err = platformSecretsListParams.Marshal()
//
//    platformSecretsListResponse, err := UnmarshalPlatformSecretsListResponse(bytes)
//    bytes, err = platformSecretsListResponse.Marshal()
//
//    platformTutorialFetchParams, err := UnmarshalPlatformTutorialFetchParams(bytes)
//    bytes, err = platformTutorialFetchParams.Marshal()
//
//    platformTutorialFetchResponse, err := UnmarshalPlatformTutorialFetchResponse(bytes)
//    bytes, err = platformTutorialFetchResponse.Marshal()
//
//    platformTutorialsListParams, err := UnmarshalPlatformTutorialsListParams(bytes)
//    bytes, err = platformTutorialsListParams.Marshal()
//
//    platformTutorialsListResponse, err := UnmarshalPlatformTutorialsListResponse(bytes)
//    bytes, err = platformTutorialsListResponse.Marshal()
//
//    platformTutorialsSearchRequest, err := UnmarshalPlatformTutorialsSearchRequest(bytes)
//    bytes, err = platformTutorialsSearchRequest.Marshal()
//
//    platformTutorialsSearchResponse, err := UnmarshalPlatformTutorialsSearchResponse(bytes)
//    bytes, err = platformTutorialsSearchResponse.Marshal()
//
//    policyDeleteParams, err := UnmarshalPolicyDeleteParams(bytes)
//    bytes, err = policyDeleteParams.Marshal()
//
//    policyDeleteRequest, err := UnmarshalPolicyDeleteRequest(bytes)
//    bytes, err = policyDeleteRequest.Marshal()
//
//    policyDeleteResponse, err := UnmarshalPolicyDeleteResponse(bytes)
//    bytes, err = policyDeleteResponse.Marshal()
//
//    policyFetchParams, err := UnmarshalPolicyFetchParams(bytes)
//    bytes, err = policyFetchParams.Marshal()
//
//    policyFetchResponse, err := UnmarshalPolicyFetchResponse(bytes)
//    bytes, err = policyFetchResponse.Marshal()
//
//    policyUpdateParams, err := UnmarshalPolicyUpdateParams(bytes)
//    bytes, err = policyUpdateParams.Marshal()
//
//    policyUpdateRequest, err := UnmarshalPolicyUpdateRequest(bytes)
//    bytes, err = policyUpdateRequest.Marshal()
//
//    policyUpdateResponse, err := UnmarshalPolicyUpdateResponse(bytes)
//    bytes, err = policyUpdateResponse.Marshal()
//
//    policyCreateRequest, err := UnmarshalPolicyCreateRequest(bytes)
//    bytes, err = policyCreateRequest.Marshal()
//
//    policyCreateResponse, err := UnmarshalPolicyCreateResponse(bytes)
//    bytes, err = policyCreateResponse.Marshal()
//
//    policiesListParams, err := UnmarshalPoliciesListParams(bytes)
//    bytes, err = policiesListParams.Marshal()
//
//    policiesListResponse, err := UnmarshalPoliciesListResponse(bytes)
//    bytes, err = policiesListResponse.Marshal()
//
//    portalDeleteParams, err := UnmarshalPortalDeleteParams(bytes)
//    bytes, err = portalDeleteParams.Marshal()
//
//    portalDeleteRequest, err := UnmarshalPortalDeleteRequest(bytes)
//    bytes, err = portalDeleteRequest.Marshal()
//
//    portalDeleteResponse, err := UnmarshalPortalDeleteResponse(bytes)
//    bytes, err = portalDeleteResponse.Marshal()
//
//    portalFetchParams, err := UnmarshalPortalFetchParams(bytes)
//    bytes, err = portalFetchParams.Marshal()
//
//    portalFetchResponse, err := UnmarshalPortalFetchResponse(bytes)
//    bytes, err = portalFetchResponse.Marshal()
//
//    portalUpdateParams, err := UnmarshalPortalUpdateParams(bytes)
//    bytes, err = portalUpdateParams.Marshal()
//
//    portalUpdateRequest, err := UnmarshalPortalUpdateRequest(bytes)
//    bytes, err = portalUpdateRequest.Marshal()
//
//    portalUpdateResponse, err := UnmarshalPortalUpdateResponse(bytes)
//    bytes, err = portalUpdateResponse.Marshal()
//
//    portalCreateRequest, err := UnmarshalPortalCreateRequest(bytes)
//    bytes, err = portalCreateRequest.Marshal()
//
//    portalCreateResponse, err := UnmarshalPortalCreateResponse(bytes)
//    bytes, err = portalCreateResponse.Marshal()
//
//    portalsListParams, err := UnmarshalPortalsListParams(bytes)
//    bytes, err = portalsListParams.Marshal()
//
//    portalsListResponse, err := UnmarshalPortalsListResponse(bytes)
//    bytes, err = portalsListResponse.Marshal()
//
//    secretAuthenticateParams, err := UnmarshalSecretAuthenticateParams(bytes)
//    bytes, err = secretAuthenticateParams.Marshal()
//
//    secretAuthenticateRequest, err := UnmarshalSecretAuthenticateRequest(bytes)
//    bytes, err = secretAuthenticateRequest.Marshal()
//
//    secretAuthenticateResponse, err := UnmarshalSecretAuthenticateResponse(bytes)
//    bytes, err = secretAuthenticateResponse.Marshal()
//
//    secretDeleteParams, err := UnmarshalSecretDeleteParams(bytes)
//    bytes, err = secretDeleteParams.Marshal()
//
//    secretDeleteRequest, err := UnmarshalSecretDeleteRequest(bytes)
//    bytes, err = secretDeleteRequest.Marshal()
//
//    secretDeleteResponse, err := UnmarshalSecretDeleteResponse(bytes)
//    bytes, err = secretDeleteResponse.Marshal()
//
//    secretFetchParams, err := UnmarshalSecretFetchParams(bytes)
//    bytes, err = secretFetchParams.Marshal()
//
//    secretFetchResponse, err := UnmarshalSecretFetchResponse(bytes)
//    bytes, err = secretFetchResponse.Marshal()
//
//    secretRevokeParams, err := UnmarshalSecretRevokeParams(bytes)
//    bytes, err = secretRevokeParams.Marshal()
//
//    secretRevokeRequest, err := UnmarshalSecretRevokeRequest(bytes)
//    bytes, err = secretRevokeRequest.Marshal()
//
//    secretRevokeResponse, err := UnmarshalSecretRevokeResponse(bytes)
//    bytes, err = secretRevokeResponse.Marshal()
//
//    secretUpdateParams, err := UnmarshalSecretUpdateParams(bytes)
//    bytes, err = secretUpdateParams.Marshal()
//
//    secretUpdateRequest, err := UnmarshalSecretUpdateRequest(bytes)
//    bytes, err = secretUpdateRequest.Marshal()
//
//    secretUpdateResponse, err := UnmarshalSecretUpdateResponse(bytes)
//    bytes, err = secretUpdateResponse.Marshal()
//
//    secretVerifyParams, err := UnmarshalSecretVerifyParams(bytes)
//    bytes, err = secretVerifyParams.Marshal()
//
//    secretVerifyRequest, err := UnmarshalSecretVerifyRequest(bytes)
//    bytes, err = secretVerifyRequest.Marshal()
//
//    secretVerifyResponse, err := UnmarshalSecretVerifyResponse(bytes)
//    bytes, err = secretVerifyResponse.Marshal()
//
//    secretCreateRequest, err := UnmarshalSecretCreateRequest(bytes)
//    bytes, err = secretCreateRequest.Marshal()
//
//    secretCreateResponse, err := UnmarshalSecretCreateResponse(bytes)
//    bytes, err = secretCreateResponse.Marshal()
//
//    secretsListParams, err := UnmarshalSecretsListParams(bytes)
//    bytes, err = secretsListParams.Marshal()
//
//    secretsListResponse, err := UnmarshalSecretsListResponse(bytes)
//    bytes, err = secretsListResponse.Marshal()
//
//    skillsetAbilityDeleteParams, err := UnmarshalSkillsetAbilityDeleteParams(bytes)
//    bytes, err = skillsetAbilityDeleteParams.Marshal()
//
//    skillsetAbilityDeleteRequest, err := UnmarshalSkillsetAbilityDeleteRequest(bytes)
//    bytes, err = skillsetAbilityDeleteRequest.Marshal()
//
//    skillsetAbilityDeleteResponse, err := UnmarshalSkillsetAbilityDeleteResponse(bytes)
//    bytes, err = skillsetAbilityDeleteResponse.Marshal()
//
//    skillsetAbilityExecuteParams, err := UnmarshalSkillsetAbilityExecuteParams(bytes)
//    bytes, err = skillsetAbilityExecuteParams.Marshal()
//
//    skillsetAbilityExecuteRequest, err := UnmarshalSkillsetAbilityExecuteRequest(bytes)
//    bytes, err = skillsetAbilityExecuteRequest.Marshal()
//
//    skillsetAbilityExecuteResponse, err := UnmarshalSkillsetAbilityExecuteResponse(bytes)
//    bytes, err = skillsetAbilityExecuteResponse.Marshal()
//
//    skillsetAbilityFetchParams, err := UnmarshalSkillsetAbilityFetchParams(bytes)
//    bytes, err = skillsetAbilityFetchParams.Marshal()
//
//    skillsetAbilityFetchResponse, err := UnmarshalSkillsetAbilityFetchResponse(bytes)
//    bytes, err = skillsetAbilityFetchResponse.Marshal()
//
//    skillsetAbilityUpdateParams, err := UnmarshalSkillsetAbilityUpdateParams(bytes)
//    bytes, err = skillsetAbilityUpdateParams.Marshal()
//
//    skillsetAbilityUpdateRequest, err := UnmarshalSkillsetAbilityUpdateRequest(bytes)
//    bytes, err = skillsetAbilityUpdateRequest.Marshal()
//
//    skillsetAbilityUpdateResponse, err := UnmarshalSkillsetAbilityUpdateResponse(bytes)
//    bytes, err = skillsetAbilityUpdateResponse.Marshal()
//
//    skillsetAbilityCreateParams, err := UnmarshalSkillsetAbilityCreateParams(bytes)
//    bytes, err = skillsetAbilityCreateParams.Marshal()
//
//    skillsetAbilityCreateRequest, err := UnmarshalSkillsetAbilityCreateRequest(bytes)
//    bytes, err = skillsetAbilityCreateRequest.Marshal()
//
//    skillsetAbilityCreateResponse, err := UnmarshalSkillsetAbilityCreateResponse(bytes)
//    bytes, err = skillsetAbilityCreateResponse.Marshal()
//
//    skillsetAbilitiesExportParams, err := UnmarshalSkillsetAbilitiesExportParams(bytes)
//    bytes, err = skillsetAbilitiesExportParams.Marshal()
//
//    skillsetAbilitiesExportResponse, err := UnmarshalSkillsetAbilitiesExportResponse(bytes)
//    bytes, err = skillsetAbilitiesExportResponse.Marshal()
//
//    skillsetAbilitiesListParams, err := UnmarshalSkillsetAbilitiesListParams(bytes)
//    bytes, err = skillsetAbilitiesListParams.Marshal()
//
//    skillsetAbilitiesListResponse, err := UnmarshalSkillsetAbilitiesListResponse(bytes)
//    bytes, err = skillsetAbilitiesListResponse.Marshal()
//
//    skillsetDeleteParams, err := UnmarshalSkillsetDeleteParams(bytes)
//    bytes, err = skillsetDeleteParams.Marshal()
//
//    skillsetDeleteRequest, err := UnmarshalSkillsetDeleteRequest(bytes)
//    bytes, err = skillsetDeleteRequest.Marshal()
//
//    skillsetDeleteResponse, err := UnmarshalSkillsetDeleteResponse(bytes)
//    bytes, err = skillsetDeleteResponse.Marshal()
//
//    skillsetFetchParams, err := UnmarshalSkillsetFetchParams(bytes)
//    bytes, err = skillsetFetchParams.Marshal()
//
//    skillsetFetchResponse, err := UnmarshalSkillsetFetchResponse(bytes)
//    bytes, err = skillsetFetchResponse.Marshal()
//
//    skillsetUpdateParams, err := UnmarshalSkillsetUpdateParams(bytes)
//    bytes, err = skillsetUpdateParams.Marshal()
//
//    skillsetUpdateRequest, err := UnmarshalSkillsetUpdateRequest(bytes)
//    bytes, err = skillsetUpdateRequest.Marshal()
//
//    skillsetUpdateResponse, err := UnmarshalSkillsetUpdateResponse(bytes)
//    bytes, err = skillsetUpdateResponse.Marshal()
//
//    skillsetCreateRequest, err := UnmarshalSkillsetCreateRequest(bytes)
//    bytes, err = skillsetCreateRequest.Marshal()
//
//    skillsetCreateResponse, err := UnmarshalSkillsetCreateResponse(bytes)
//    bytes, err = skillsetCreateResponse.Marshal()
//
//    skillsetsListParams, err := UnmarshalSkillsetsListParams(bytes)
//    bytes, err = skillsetsListParams.Marshal()
//
//    skillsetsListResponse, err := UnmarshalSkillsetsListResponse(bytes)
//    bytes, err = skillsetsListResponse.Marshal()
//
//    spaceFetchParams, err := UnmarshalSpaceFetchParams(bytes)
//    bytes, err = spaceFetchParams.Marshal()
//
//    spaceFetchResponse, err := UnmarshalSpaceFetchResponse(bytes)
//    bytes, err = spaceFetchResponse.Marshal()
//
//    spaceUpdateParams, err := UnmarshalSpaceUpdateParams(bytes)
//    bytes, err = spaceUpdateParams.Marshal()
//
//    spaceUpdateRequest, err := UnmarshalSpaceUpdateRequest(bytes)
//    bytes, err = spaceUpdateRequest.Marshal()
//
//    spaceUpdateResponse, err := UnmarshalSpaceUpdateResponse(bytes)
//    bytes, err = spaceUpdateResponse.Marshal()
//
//    spaceCreateRequest, err := UnmarshalSpaceCreateRequest(bytes)
//    bytes, err = spaceCreateRequest.Marshal()
//
//    spaceCreateResponse, err := UnmarshalSpaceCreateResponse(bytes)
//    bytes, err = spaceCreateResponse.Marshal()
//
//    spacesExportParams, err := UnmarshalSpacesExportParams(bytes)
//    bytes, err = spacesExportParams.Marshal()
//
//    spacesExportResponse, err := UnmarshalSpacesExportResponse(bytes)
//    bytes, err = spacesExportResponse.Marshal()
//
//    spacesListParams, err := UnmarshalSpacesListParams(bytes)
//    bytes, err = spacesListParams.Marshal()
//
//    spacesListResponse, err := UnmarshalSpacesListResponse(bytes)
//    bytes, err = spacesListResponse.Marshal()
//
//    taskDeleteParams, err := UnmarshalTaskDeleteParams(bytes)
//    bytes, err = taskDeleteParams.Marshal()
//
//    taskDeleteRequest, err := UnmarshalTaskDeleteRequest(bytes)
//    bytes, err = taskDeleteRequest.Marshal()
//
//    taskDeleteResponse, err := UnmarshalTaskDeleteResponse(bytes)
//    bytes, err = taskDeleteResponse.Marshal()
//
//    taskFetchParams, err := UnmarshalTaskFetchParams(bytes)
//    bytes, err = taskFetchParams.Marshal()
//
//    taskFetchResponse, err := UnmarshalTaskFetchResponse(bytes)
//    bytes, err = taskFetchResponse.Marshal()
//
//    taskTriggerParams, err := UnmarshalTaskTriggerParams(bytes)
//    bytes, err = taskTriggerParams.Marshal()
//
//    taskTriggerRequest, err := UnmarshalTaskTriggerRequest(bytes)
//    bytes, err = taskTriggerRequest.Marshal()
//
//    taskTriggerResponse, err := UnmarshalTaskTriggerResponse(bytes)
//    bytes, err = taskTriggerResponse.Marshal()
//
//    taskUpdateParams, err := UnmarshalTaskUpdateParams(bytes)
//    bytes, err = taskUpdateParams.Marshal()
//
//    taskUpdateRequest, err := UnmarshalTaskUpdateRequest(bytes)
//    bytes, err = taskUpdateRequest.Marshal()
//
//    taskUpdateResponse, err := UnmarshalTaskUpdateResponse(bytes)
//    bytes, err = taskUpdateResponse.Marshal()
//
//    taskCreateRequest, err := UnmarshalTaskCreateRequest(bytes)
//    bytes, err = taskCreateRequest.Marshal()
//
//    taskCreateResponse, err := UnmarshalTaskCreateResponse(bytes)
//    bytes, err = taskCreateResponse.Marshal()
//
//    tasksExportParams, err := UnmarshalTasksExportParams(bytes)
//    bytes, err = tasksExportParams.Marshal()
//
//    tasksExportResponse, err := UnmarshalTasksExportResponse(bytes)
//    bytes, err = tasksExportResponse.Marshal()
//
//    tasksListParams, err := UnmarshalTasksListParams(bytes)
//    bytes, err = tasksListParams.Marshal()
//
//    tasksListResponse, err := UnmarshalTasksListResponse(bytes)
//    bytes, err = tasksListResponse.Marshal()
//
//    teamsListParams, err := UnmarshalTeamsListParams(bytes)
//    bytes, err = teamsListParams.Marshal()
//
//    teamsListResponse, err := UnmarshalTeamsListResponse(bytes)
//    bytes, err = teamsListResponse.Marshal()
//
//    usageFetchResponse, err := UnmarshalUsageFetchResponse(bytes)
//    bytes, err = usageFetchResponse.Marshal()
//
//    usageSeriesFetchResponse, err := UnmarshalUsageSeriesFetchResponse(bytes)
//    bytes, err = usageSeriesFetchResponse.Marshal()
//
//    message, err := UnmarshalMessage(bytes)
//    bytes, err = message.Marshal()
//
//    entity, err := UnmarshalEntity(bytes)
//    bytes, err = entity.Marshal()
//
//    messageType, err := UnmarshalMessageType(bytes)
//    bytes, err = messageType.Marshal()
//
//    trigger, err := UnmarshalTrigger(bytes)
//    bytes, err = trigger.Marshal()
//
//    schedule, err := UnmarshalSchedule(bytes)
//    bytes, err = schedule.Marshal()
//
//    syncStatus, err := UnmarshalSyncStatus(bytes)
//    bytes, err = syncStatus.Marshal()
//
//    taskStatus, err := UnmarshalTaskStatus(bytes)
//    bytes, err = taskStatus.Marshal()
//
//    taskOutcome, err := UnmarshalTaskOutcome(bytes)
//    bytes, err = taskOutcome.Marshal()
//
//    blueprintVisibility, err := UnmarshalBlueprintVisibility(bytes)
//    bytes, err = blueprintVisibility.Marshal()
//
//    botVisibility, err := UnmarshalBotVisibility(bytes)
//    bytes, err = botVisibility.Marshal()
//
//    datasetVisibility, err := UnmarshalDatasetVisibility(bytes)
//    bytes, err = datasetVisibility.Marshal()
//
//    datasetFileAttachmentType, err := UnmarshalDatasetFileAttachmentType(bytes)
//    bytes, err = datasetFileAttachmentType.Marshal()
//
//    datasetFilter, err := UnmarshalDatasetFilter(bytes)
//    bytes, err = datasetFilter.Marshal()
//
//    skillsetVisibility, err := UnmarshalSkillsetVisibility(bytes)
//    bytes, err = skillsetVisibility.Marshal()
//
//    fileVisibility, err := UnmarshalFileVisibility(bytes)
//    bytes, err = fileVisibility.Marshal()
//
//    secretType, err := UnmarshalSecretType(bytes)
//    bytes, err = secretType.Marshal()
//
//    secretKind, err := UnmarshalSecretKind(bytes)
//    bytes, err = secretKind.Marshal()
//
//    secretVisibility, err := UnmarshalSecretVisibility(bytes)
//    bytes, err = secretVisibility.Marshal()
//
//    usage, err := UnmarshalUsage(bytes)
//    bytes, err = usage.Marshal()
//
//    completeReason, err := UnmarshalCompleteReason(bytes)
//    bytes, err = completeReason.Marshal()
//
//    completeEnd, err := UnmarshalCompleteEnd(bytes)
//    bytes, err = completeEnd.Marshal()
//
//    executionLimits, err := UnmarshalExecutionLimits(bytes)
//    bytes, err = executionLimits.Marshal()
//
//    policyType, err := UnmarshalPolicyType(bytes)
//    bytes, err = policyType.Marshal()
//
//    limits, err := UnmarshalLimits(bytes)
//    bytes, err = limits.Marshal()
//
//    meta, err := UnmarshalMeta(bytes)
//    bytes, err = meta.Marshal()
//
//    model, err := UnmarshalModel(bytes)
//    bytes, err = model.Marshal()
//
//    botRef, err := UnmarshalBotRef(bytes)
//    bytes, err = botRef.Marshal()
//
//    botConfig, err := UnmarshalBotConfig(bytes)
//    bytes, err = botConfig.Marshal()
//
//    botRefOrConfig, err := UnmarshalBotRefOrConfig(bytes)
//    bytes, err = botRefOrConfig.Marshal()
//
//    blueprintProps, err := UnmarshalBlueprintProps(bytes)
//    bytes, err = blueprintProps.Marshal()
//
//    instanceRefProperties, err := UnmarshalInstanceRefProperties(bytes)
//    bytes, err = instanceRefProperties.Marshal()
//
//    instanceMetaProps, err := UnmarshalInstanceMetaProps(bytes)
//    bytes, err = instanceMetaProps.Marshal()
//
//    instanceCRUDProps, err := UnmarshalInstanceCRUDProps(bytes)
//    bytes, err = instanceCRUDProps.Marshal()
//
//    instanceListProps, err := UnmarshalInstanceListProps(bytes)
//    bytes, err = instanceListProps.Marshal()
//
//    jSONSchemaObject, err := UnmarshalJSONSchemaObject(bytes)
//    bytes, err = jSONSchemaObject.Marshal()
//
//    functionsDefinition, err := UnmarshalFunctionsDefinition(bytes)
//    bytes, err = functionsDefinition.Marshal()
//
//    extensionsDefinition, err := UnmarshalExtensionsDefinition(bytes)
//    bytes, err = extensionsDefinition.Marshal()
//
//    completeStreamingResponseItem, err := UnmarshalCompleteStreamingResponseItem(bytes)
//    bytes, err = completeStreamingResponseItem.Marshal()

package types

import "bytes"
import "errors"
import "time"

import "encoding/json"

func UnmarshalGraphqlRequest(data []byte) (GraphqlRequest, error) {
	var r GraphqlRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GraphqlRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalGraphqlResponse(data []byte) (GraphqlResponse, error) {
	var r GraphqlResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GraphqlResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformReportsListParams(data []byte) (PlatformReportsListParams, error) {
	var r PlatformReportsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformReportsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformReportsListResponse(data []byte) (PlatformReportsListResponse, error) {
	var r PlatformReportsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformReportsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalReportGenerateParams(data []byte) (ReportGenerateParams, error) {
	var r ReportGenerateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReportGenerateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReportGenerateRequest map[string]interface{}

func UnmarshalReportGenerateRequest(data []byte) (ReportGenerateRequest, error) {
	var r ReportGenerateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReportGenerateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReportGenerateResponse map[string]interface{}

func UnmarshalReportGenerateResponse(data []byte) (ReportGenerateResponse, error) {
	var r ReportGenerateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReportGenerateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReportsGenerateRequest map[string]map[string]interface{}

func UnmarshalReportsGenerateRequest(data []byte) (ReportsGenerateRequest, error) {
	var r ReportsGenerateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReportsGenerateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ReportsGenerateResponse map[string]ReportsGenerateResponseValue

func UnmarshalReportsGenerateResponse(data []byte) (ReportsGenerateResponse, error) {
	var r ReportsGenerateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReportsGenerateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintCloneParams(data []byte) (BlueprintCloneParams, error) {
	var r BlueprintCloneParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintCloneParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BlueprintCloneRequest map[string]interface{}

func UnmarshalBlueprintCloneRequest(data []byte) (BlueprintCloneRequest, error) {
	var r BlueprintCloneRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintCloneRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintCloneResponse(data []byte) (BlueprintCloneResponse, error) {
	var r BlueprintCloneResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintCloneResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintDeleteParams(data []byte) (BlueprintDeleteParams, error) {
	var r BlueprintDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintDeleteRequest(data []byte) (BlueprintDeleteRequest, error) {
	var r BlueprintDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintDeleteResponse(data []byte) (BlueprintDeleteResponse, error) {
	var r BlueprintDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintFetchParams(data []byte) (BlueprintFetchParams, error) {
	var r BlueprintFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintFetchResponse(data []byte) (BlueprintFetchResponse, error) {
	var r BlueprintFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintResourcesListParams(data []byte) (BlueprintResourcesListParams, error) {
	var r BlueprintResourcesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintResourcesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintResourcesListResponse(data []byte) (BlueprintResourcesListResponse, error) {
	var r BlueprintResourcesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintResourcesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintUpdateParams(data []byte) (BlueprintUpdateParams, error) {
	var r BlueprintUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintUpdateRequest(data []byte) (BlueprintUpdateRequest, error) {
	var r BlueprintUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintUpdateResponse(data []byte) (BlueprintUpdateResponse, error) {
	var r BlueprintUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintCreateRequest(data []byte) (BlueprintCreateRequest, error) {
	var r BlueprintCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintCreateResponse(data []byte) (BlueprintCreateResponse, error) {
	var r BlueprintCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintsListParams(data []byte) (BlueprintsListParams, error) {
	var r BlueprintsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintsListResponse(data []byte) (BlueprintsListResponse, error) {
	var r BlueprintsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotCloneParams(data []byte) (BotCloneParams, error) {
	var r BotCloneParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotCloneParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BotCloneRequest interface{}

func UnmarshalBotCloneRequest(data []byte) (BotCloneRequest, error) {
	var r BotCloneRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotCloneRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotCloneResponse(data []byte) (BotCloneResponse, error) {
	var r BotCloneResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotCloneResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotDeleteParams(data []byte) (BotDeleteParams, error) {
	var r BotDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BotDeleteRequest map[string]interface{}

func UnmarshalBotDeleteRequest(data []byte) (BotDeleteRequest, error) {
	var r BotDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotDeleteResponse(data []byte) (BotDeleteResponse, error) {
	var r BotDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotDownvoteParams(data []byte) (BotDownvoteParams, error) {
	var r BotDownvoteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotDownvoteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotDownvoteRequest(data []byte) (BotDownvoteRequest, error) {
	var r BotDownvoteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotDownvoteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotDownvoteResponse(data []byte) (BotDownvoteResponse, error) {
	var r BotDownvoteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotDownvoteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotFetchParams(data []byte) (BotFetchParams, error) {
	var r BotFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotFetchResponse(data []byte) (BotFetchResponse, error) {
	var r BotFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotMemorySearchParams(data []byte) (BotMemorySearchParams, error) {
	var r BotMemorySearchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotMemorySearchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotMemorySearchRequest(data []byte) (BotMemorySearchRequest, error) {
	var r BotMemorySearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotMemorySearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotMemorySearchResponse(data []byte) (BotMemorySearchResponse, error) {
	var r BotMemorySearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotMemorySearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotSessionCreateParams(data []byte) (BotSessionCreateParams, error) {
	var r BotSessionCreateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotSessionCreateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotSessionCreateRequest(data []byte) (BotSessionCreateRequest, error) {
	var r BotSessionCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotSessionCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotSessionCreateResponse(data []byte) (BotSessionCreateResponse, error) {
	var r BotSessionCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotSessionCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotUpdateParams(data []byte) (BotUpdateParams, error) {
	var r BotUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotUpdateRequest(data []byte) (BotUpdateRequest, error) {
	var r BotUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotUpdateResponse(data []byte) (BotUpdateResponse, error) {
	var r BotUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotUpvoteParams(data []byte) (BotUpvoteParams, error) {
	var r BotUpvoteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotUpvoteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotUpvoteRequest(data []byte) (BotUpvoteRequest, error) {
	var r BotUpvoteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotUpvoteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotUpvoteResponse(data []byte) (BotUpvoteResponse, error) {
	var r BotUpvoteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotUpvoteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotUsageFetchParams(data []byte) (BotUsageFetchParams, error) {
	var r BotUsageFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotUsageFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotUsageFetchResponse(data []byte) (BotUsageFetchResponse, error) {
	var r BotUsageFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotUsageFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotCreateRequest(data []byte) (BotCreateRequest, error) {
	var r BotCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotCreateResponse(data []byte) (BotCreateResponse, error) {
	var r BotCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotsListParams(data []byte) (BotsListParams, error) {
	var r BotsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotsListResponse(data []byte) (BotsListResponse, error) {
	var r BotsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalChannelMessagePublishParams(data []byte) (ChannelMessagePublishParams, error) {
	var r ChannelMessagePublishParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelMessagePublishParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalChannelMessagePublishRequest(data []byte) (ChannelMessagePublishRequest, error) {
	var r ChannelMessagePublishRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelMessagePublishRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalChannelMessagePublishResponse(data []byte) (ChannelMessagePublishResponse, error) {
	var r ChannelMessagePublishResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelMessagePublishResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalChannelMessagesSubscribeParams(data []byte) (ChannelMessagesSubscribeParams, error) {
	var r ChannelMessagesSubscribeParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelMessagesSubscribeParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalChannelMessagesSubscribeRequest(data []byte) (ChannelMessagesSubscribeRequest, error) {
	var r ChannelMessagesSubscribeRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChannelMessagesSubscribeRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactConversationsListParams(data []byte) (ContactConversationsListParams, error) {
	var r ContactConversationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactConversationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactConversationsListResponse(data []byte) (ContactConversationsListResponse, error) {
	var r ContactConversationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactConversationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactDeleteParams(data []byte) (ContactDeleteParams, error) {
	var r ContactDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ContactDeleteRequest map[string]interface{}

func UnmarshalContactDeleteRequest(data []byte) (ContactDeleteRequest, error) {
	var r ContactDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactDeleteResponse(data []byte) (ContactDeleteResponse, error) {
	var r ContactDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactFetchParams(data []byte) (ContactFetchParams, error) {
	var r ContactFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactFetchResponse(data []byte) (ContactFetchResponse, error) {
	var r ContactFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactMemoriesListParams(data []byte) (ContactMemoriesListParams, error) {
	var r ContactMemoriesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactMemoriesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactMemoriesListResponse(data []byte) (ContactMemoriesListResponse, error) {
	var r ContactMemoriesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactMemoriesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactMemorySearchParams(data []byte) (ContactMemorySearchParams, error) {
	var r ContactMemorySearchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactMemorySearchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactMemorySearchRequest(data []byte) (ContactMemorySearchRequest, error) {
	var r ContactMemorySearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactMemorySearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactMemorySearchResponse(data []byte) (ContactMemorySearchResponse, error) {
	var r ContactMemorySearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactMemorySearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSecretAuthenticateParams(data []byte) (ContactSecretAuthenticateParams, error) {
	var r ContactSecretAuthenticateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretAuthenticateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ContactSecretAuthenticateRequest map[string]interface{}

func UnmarshalContactSecretAuthenticateRequest(data []byte) (ContactSecretAuthenticateRequest, error) {
	var r ContactSecretAuthenticateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretAuthenticateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSecretAuthenticateResponse(data []byte) (ContactSecretAuthenticateResponse, error) {
	var r ContactSecretAuthenticateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretAuthenticateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSecretRevokeParams(data []byte) (ContactSecretRevokeParams, error) {
	var r ContactSecretRevokeParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretRevokeParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ContactSecretRevokeRequest map[string]interface{}

func UnmarshalContactSecretRevokeRequest(data []byte) (ContactSecretRevokeRequest, error) {
	var r ContactSecretRevokeRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretRevokeRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSecretRevokeResponse(data []byte) (ContactSecretRevokeResponse, error) {
	var r ContactSecretRevokeResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretRevokeResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSecretVerifyParams(data []byte) (ContactSecretVerifyParams, error) {
	var r ContactSecretVerifyParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretVerifyParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ContactSecretVerifyRequest map[string]interface{}

func UnmarshalContactSecretVerifyRequest(data []byte) (ContactSecretVerifyRequest, error) {
	var r ContactSecretVerifyRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretVerifyRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSecretVerifyResponse(data []byte) (ContactSecretVerifyResponse, error) {
	var r ContactSecretVerifyResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretVerifyResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSecretsListParams(data []byte) (ContactSecretsListParams, error) {
	var r ContactSecretsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSecretsListResponse(data []byte) (ContactSecretsListResponse, error) {
	var r ContactSecretsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSecretsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSpacesListParams(data []byte) (ContactSpacesListParams, error) {
	var r ContactSpacesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSpacesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactSpacesListResponse(data []byte) (ContactSpacesListResponse, error) {
	var r ContactSpacesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactSpacesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactTasksListParams(data []byte) (ContactTasksListParams, error) {
	var r ContactTasksListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactTasksListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactTasksListResponse(data []byte) (ContactTasksListResponse, error) {
	var r ContactTasksListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactTasksListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactUpdateParams(data []byte) (ContactUpdateParams, error) {
	var r ContactUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactUpdateRequest(data []byte) (ContactUpdateRequest, error) {
	var r ContactUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactUpdateResponse(data []byte) (ContactUpdateResponse, error) {
	var r ContactUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactCreateRequest(data []byte) (ContactCreateRequest, error) {
	var r ContactCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactCreateResponse(data []byte) (ContactCreateResponse, error) {
	var r ContactCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactEnsureRequest(data []byte) (ContactEnsureRequest, error) {
	var r ContactEnsureRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactEnsureRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactEnsureResponse(data []byte) (ContactEnsureResponse, error) {
	var r ContactEnsureResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactEnsureResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactsExportParams(data []byte) (ContactsExportParams, error) {
	var r ContactsExportParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactsExportParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactsExportResponse(data []byte) (ContactsExportResponse, error) {
	var r ContactsExportResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactsExportResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactsListParams(data []byte) (ContactsListParams, error) {
	var r ContactsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalContactsListResponse(data []byte) (ContactsListResponse, error) {
	var r ContactsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationAttachmentUploadParams(data []byte) (ConversationAttachmentUploadParams, error) {
	var r ConversationAttachmentUploadParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationAttachmentUploadParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationAttachmentUploadRequest(data []byte) (ConversationAttachmentUploadRequest, error) {
	var r ConversationAttachmentUploadRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationAttachmentUploadRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationAttachmentUploadResponse(data []byte) (ConversationAttachmentUploadResponse, error) {
	var r ConversationAttachmentUploadResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationAttachmentUploadResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageCompleteParams(data []byte) (ConversationMessageCompleteParams, error) {
	var r ConversationMessageCompleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageCompleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageCompleteRequest(data []byte) (ConversationMessageCompleteRequest, error) {
	var r ConversationMessageCompleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageCompleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageCompleteResponse(data []byte) (ConversationMessageCompleteResponse, error) {
	var r ConversationMessageCompleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageCompleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationContactUpsertParams(data []byte) (ConversationContactUpsertParams, error) {
	var r ConversationContactUpsertParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationContactUpsertParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationContactUpsertRequest(data []byte) (ConversationContactUpsertRequest, error) {
	var r ConversationContactUpsertRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationContactUpsertRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationContactUpsertResponse(data []byte) (ConversationContactUpsertResponse, error) {
	var r ConversationContactUpsertResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationContactUpsertResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationDeleteParams(data []byte) (ConversationDeleteParams, error) {
	var r ConversationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationDeleteRequest map[string]interface{}

func UnmarshalConversationDeleteRequest(data []byte) (ConversationDeleteRequest, error) {
	var r ConversationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationDeleteResponse(data []byte) (ConversationDeleteResponse, error) {
	var r ConversationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStatefulConversationDispatchRequest(data []byte) (StatefulConversationDispatchRequest, error) {
	var r StatefulConversationDispatchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StatefulConversationDispatchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStatefulConversationDispatchResponse(data []byte) (StatefulConversationDispatchResponse, error) {
	var r StatefulConversationDispatchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StatefulConversationDispatchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationDownvoteParams(data []byte) (ConversationDownvoteParams, error) {
	var r ConversationDownvoteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationDownvoteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationDownvoteRequest(data []byte) (ConversationDownvoteRequest, error) {
	var r ConversationDownvoteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationDownvoteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationDownvoteResponse(data []byte) (ConversationDownvoteResponse, error) {
	var r ConversationDownvoteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationDownvoteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationFetchParams(data []byte) (ConversationFetchParams, error) {
	var r ConversationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationFetchResponse(data []byte) (ConversationFetchResponse, error) {
	var r ConversationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageDeleteParams(data []byte) (ConversationMessageDeleteParams, error) {
	var r ConversationMessageDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationMessageDeleteRequest map[string]interface{}

func UnmarshalConversationMessageDeleteRequest(data []byte) (ConversationMessageDeleteRequest, error) {
	var r ConversationMessageDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageDeleteResponse(data []byte) (ConversationMessageDeleteResponse, error) {
	var r ConversationMessageDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageDownvoteParams(data []byte) (ConversationMessageDownvoteParams, error) {
	var r ConversationMessageDownvoteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageDownvoteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageDownvoteRequest(data []byte) (ConversationMessageDownvoteRequest, error) {
	var r ConversationMessageDownvoteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageDownvoteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageDownvoteResponse(data []byte) (ConversationMessageDownvoteResponse, error) {
	var r ConversationMessageDownvoteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageDownvoteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageFetchParams(data []byte) (ConversationMessageFetchParams, error) {
	var r ConversationMessageFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageFetchResponse(data []byte) (ConversationMessageFetchResponse, error) {
	var r ConversationMessageFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageSynthesizeParams(data []byte) (ConversationMessageSynthesizeParams, error) {
	var r ConversationMessageSynthesizeParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageSynthesizeParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ConversationMessageSynthesizeRequest map[string]interface{}

func UnmarshalConversationMessageSynthesizeRequest(data []byte) (ConversationMessageSynthesizeRequest, error) {
	var r ConversationMessageSynthesizeRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageSynthesizeRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageSynthesizeResponse(data []byte) (ConversationMessageSynthesizeResponse, error) {
	var r ConversationMessageSynthesizeResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageSynthesizeResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageUpdateParams(data []byte) (ConversationMessageUpdateParams, error) {
	var r ConversationMessageUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageUpdateRequest(data []byte) (ConversationMessageUpdateRequest, error) {
	var r ConversationMessageUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageUpdateResponse(data []byte) (ConversationMessageUpdateResponse, error) {
	var r ConversationMessageUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageUpvoteParams(data []byte) (ConversationMessageUpvoteParams, error) {
	var r ConversationMessageUpvoteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageUpvoteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageUpvoteRequest(data []byte) (ConversationMessageUpvoteRequest, error) {
	var r ConversationMessageUpvoteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageUpvoteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageUpvoteResponse(data []byte) (ConversationMessageUpvoteResponse, error) {
	var r ConversationMessageUpvoteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageUpvoteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageCreateParams(data []byte) (ConversationMessageCreateParams, error) {
	var r ConversationMessageCreateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageCreateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageCreateRequest(data []byte) (ConversationMessageCreateRequest, error) {
	var r ConversationMessageCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageCreateResponse(data []byte) (ConversationMessageCreateResponse, error) {
	var r ConversationMessageCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessagesListParams(data []byte) (ConversationMessagesListParams, error) {
	var r ConversationMessagesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessagesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessagesListResponse(data []byte) (ConversationMessagesListResponse, error) {
	var r ConversationMessagesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessagesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageReceiveParams(data []byte) (ConversationMessageReceiveParams, error) {
	var r ConversationMessageReceiveParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageReceiveParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageReceiveRequest(data []byte) (ConversationMessageReceiveRequest, error) {
	var r ConversationMessageReceiveRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageReceiveRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageReceiveResponse(data []byte) (ConversationMessageReceiveResponse, error) {
	var r ConversationMessageReceiveResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageReceiveResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageSendParams(data []byte) (ConversationMessageSendParams, error) {
	var r ConversationMessageSendParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageSendParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageSendRequest(data []byte) (ConversationMessageSendRequest, error) {
	var r ConversationMessageSendRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageSendRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationMessageSendResponse(data []byte) (ConversationMessageSendResponse, error) {
	var r ConversationMessageSendResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationMessageSendResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationSessionCreateParams(data []byte) (ConversationSessionCreateParams, error) {
	var r ConversationSessionCreateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationSessionCreateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationSessionCreateRequest(data []byte) (ConversationSessionCreateRequest, error) {
	var r ConversationSessionCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationSessionCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationSessionCreateResponse(data []byte) (ConversationSessionCreateResponse, error) {
	var r ConversationSessionCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationSessionCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationUpdateParams(data []byte) (ConversationUpdateParams, error) {
	var r ConversationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationUpdateRequest(data []byte) (ConversationUpdateRequest, error) {
	var r ConversationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationUpdateResponse(data []byte) (ConversationUpdateResponse, error) {
	var r ConversationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationUpvoteParams(data []byte) (ConversationUpvoteParams, error) {
	var r ConversationUpvoteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationUpvoteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationUpvoteRequest(data []byte) (ConversationUpvoteRequest, error) {
	var r ConversationUpvoteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationUpvoteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationUpvoteResponse(data []byte) (ConversationUpvoteResponse, error) {
	var r ConversationUpvoteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationUpvoteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationUsageFetchParams(data []byte) (ConversationUsageFetchParams, error) {
	var r ConversationUsageFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationUsageFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationUsageFetchResponse(data []byte) (ConversationUsageFetchResponse, error) {
	var r ConversationUsageFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationUsageFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationCompleteRequest(data []byte) (ConversationCompleteRequest, error) {
	var r ConversationCompleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationCompleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationCompleteResponse(data []byte) (ConversationCompleteResponse, error) {
	var r ConversationCompleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationCompleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationCreateRequest(data []byte) (ConversationCreateRequest, error) {
	var r ConversationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationCreateResponse(data []byte) (ConversationCreateResponse, error) {
	var r ConversationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationDispatchRequest(data []byte) (ConversationDispatchRequest, error) {
	var r ConversationDispatchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationDispatchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationDispatchResponse(data []byte) (ConversationDispatchResponse, error) {
	var r ConversationDispatchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationDispatchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationsExportParams(data []byte) (ConversationsExportParams, error) {
	var r ConversationsExportParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsExportParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationsExportResponse(data []byte) (ConversationsExportResponse, error) {
	var r ConversationsExportResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsExportResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationsListParams(data []byte) (ConversationsListParams, error) {
	var r ConversationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConversationsListResponse(data []byte) (ConversationsListResponse, error) {
	var r ConversationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConversationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetDeleteParams(data []byte) (DatasetDeleteParams, error) {
	var r DatasetDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DatasetDeleteRequest map[string]interface{}

func UnmarshalDatasetDeleteRequest(data []byte) (DatasetDeleteRequest, error) {
	var r DatasetDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetDeleteResponse(data []byte) (DatasetDeleteResponse, error) {
	var r DatasetDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFetchParams(data []byte) (DatasetFetchParams, error) {
	var r DatasetFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFetchResponse(data []byte) (DatasetFetchResponse, error) {
	var r DatasetFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFileAttachParams(data []byte) (DatasetFileAttachParams, error) {
	var r DatasetFileAttachParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileAttachParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFileAttachRequest(data []byte) (DatasetFileAttachRequest, error) {
	var r DatasetFileAttachRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileAttachRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFileAttachResponse(data []byte) (DatasetFileAttachResponse, error) {
	var r DatasetFileAttachResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileAttachResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFileDetachParams(data []byte) (DatasetFileDetachParams, error) {
	var r DatasetFileDetachParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileDetachParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFileDetachRequest(data []byte) (DatasetFileDetachRequest, error) {
	var r DatasetFileDetachRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileDetachRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFileDetachResponse(data []byte) (DatasetFileDetachResponse, error) {
	var r DatasetFileDetachResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileDetachResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFileSyncParams(data []byte) (DatasetFileSyncParams, error) {
	var r DatasetFileSyncParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileSyncParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DatasetFileSyncRequest map[string]interface{}

func UnmarshalDatasetFileSyncRequest(data []byte) (DatasetFileSyncRequest, error) {
	var r DatasetFileSyncRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileSyncRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFileSyncResponse(data []byte) (DatasetFileSyncResponse, error) {
	var r DatasetFileSyncResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileSyncResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFilesListParams(data []byte) (DatasetFilesListParams, error) {
	var r DatasetFilesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFilesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFilesListResponse(data []byte) (DatasetFilesListResponse, error) {
	var r DatasetFilesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFilesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordDeleteParams(data []byte) (DatasetRecordDeleteParams, error) {
	var r DatasetRecordDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DatasetRecordDeleteRequest map[string]interface{}

func UnmarshalDatasetRecordDeleteRequest(data []byte) (DatasetRecordDeleteRequest, error) {
	var r DatasetRecordDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordDeleteResponse(data []byte) (DatasetRecordDeleteResponse, error) {
	var r DatasetRecordDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordFetchParams(data []byte) (DatasetRecordFetchParams, error) {
	var r DatasetRecordFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordFetchResponse(data []byte) (DatasetRecordFetchResponse, error) {
	var r DatasetRecordFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordUpdateParams(data []byte) (DatasetRecordUpdateParams, error) {
	var r DatasetRecordUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordUpdateRequest(data []byte) (DatasetRecordUpdateRequest, error) {
	var r DatasetRecordUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordUpdateResponse(data []byte) (DatasetRecordUpdateResponse, error) {
	var r DatasetRecordUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordCreateParams(data []byte) (DatasetRecordCreateParams, error) {
	var r DatasetRecordCreateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordCreateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordCreateRequest(data []byte) (DatasetRecordCreateRequest, error) {
	var r DatasetRecordCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordCreateResponse(data []byte) (DatasetRecordCreateResponse, error) {
	var r DatasetRecordCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordsExportParams(data []byte) (DatasetRecordsExportParams, error) {
	var r DatasetRecordsExportParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordsExportParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordsExportResponse(data []byte) (DatasetRecordsExportResponse, error) {
	var r DatasetRecordsExportResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordsExportResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordsListParams(data []byte) (DatasetRecordsListParams, error) {
	var r DatasetRecordsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetRecordsListResponse(data []byte) (DatasetRecordsListResponse, error) {
	var r DatasetRecordsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetRecordsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetSearchParams(data []byte) (DatasetSearchParams, error) {
	var r DatasetSearchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetSearchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetSearchRequest(data []byte) (DatasetSearchRequest, error) {
	var r DatasetSearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetSearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetSearchResponse(data []byte) (DatasetSearchResponse, error) {
	var r DatasetSearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetSearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetUpdateParams(data []byte) (DatasetUpdateParams, error) {
	var r DatasetUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetUpdateRequest(data []byte) (DatasetUpdateRequest, error) {
	var r DatasetUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetUpdateResponse(data []byte) (DatasetUpdateResponse, error) {
	var r DatasetUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetCreateRequest(data []byte) (DatasetCreateRequest, error) {
	var r DatasetCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetCreateResponse(data []byte) (DatasetCreateResponse, error) {
	var r DatasetCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetsListParams(data []byte) (DatasetsListParams, error) {
	var r DatasetsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetsListResponse(data []byte) (DatasetsListResponse, error) {
	var r DatasetsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEventLogsExportParams(data []byte) (EventLogsExportParams, error) {
	var r EventLogsExportParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventLogsExportParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEventLogsExportResponse(data []byte) (EventLogsExportResponse, error) {
	var r EventLogsExportResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventLogsExportResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEventLogsListParams(data []byte) (EventLogsListParams, error) {
	var r EventLogsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventLogsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEventLogsListResponse(data []byte) (EventLogsListResponse, error) {
	var r EventLogsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventLogsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEventLogsSubscribeRequest(data []byte) (EventLogsSubscribeRequest, error) {
	var r EventLogsSubscribeRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventLogsSubscribeRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileDeleteParams(data []byte) (FileDeleteParams, error) {
	var r FileDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FileDeleteRequest map[string]interface{}

func UnmarshalFileDeleteRequest(data []byte) (FileDeleteRequest, error) {
	var r FileDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileDeleteResponse(data []byte) (FileDeleteResponse, error) {
	var r FileDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileDownloadParams(data []byte) (FileDownloadParams, error) {
	var r FileDownloadParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileDownloadParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileDownloadResponse(data []byte) (FileDownloadResponse, error) {
	var r FileDownloadResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileDownloadResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileFetchParams(data []byte) (FileFetchParams, error) {
	var r FileFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileFetchResponse(data []byte) (FileFetchResponse, error) {
	var r FileFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileSyncParams(data []byte) (FileSyncParams, error) {
	var r FileSyncParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileSyncParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FileSyncRequest map[string]interface{}

func UnmarshalFileSyncRequest(data []byte) (FileSyncRequest, error) {
	var r FileSyncRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileSyncRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileSyncResponse(data []byte) (FileSyncResponse, error) {
	var r FileSyncResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileSyncResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileUpdateParams(data []byte) (FileUpdateParams, error) {
	var r FileUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileUpdateRequest(data []byte) (FileUpdateRequest, error) {
	var r FileUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileUpdateResponse(data []byte) (FileUpdateResponse, error) {
	var r FileUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileUploadParams(data []byte) (FileUploadParams, error) {
	var r FileUploadParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileUploadParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileUploadRequest(data []byte) (FileUploadRequest, error) {
	var r FileUploadRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileUploadRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileUploadResponse(data []byte) (FileUploadResponse, error) {
	var r FileUploadResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileUploadResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileCreateRequest(data []byte) (FileCreateRequest, error) {
	var r FileCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileCreateResponse(data []byte) (FileCreateResponse, error) {
	var r FileCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFilesListParams(data []byte) (FilesListParams, error) {
	var r FilesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FilesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFilesListResponse(data []byte) (FilesListResponse, error) {
	var r FilesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FilesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationDeleteParams(data []byte) (DiscordIntegrationDeleteParams, error) {
	var r DiscordIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DiscordIntegrationDeleteRequest map[string]interface{}

func UnmarshalDiscordIntegrationDeleteRequest(data []byte) (DiscordIntegrationDeleteRequest, error) {
	var r DiscordIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationDeleteResponse(data []byte) (DiscordIntegrationDeleteResponse, error) {
	var r DiscordIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationFetchParams(data []byte) (DiscordIntegrationFetchParams, error) {
	var r DiscordIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationFetchResponse(data []byte) (DiscordIntegrationFetchResponse, error) {
	var r DiscordIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationSetupParams(data []byte) (DiscordIntegrationSetupParams, error) {
	var r DiscordIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DiscordIntegrationSetupRequest map[string]interface{}

func UnmarshalDiscordIntegrationSetupRequest(data []byte) (DiscordIntegrationSetupRequest, error) {
	var r DiscordIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationSetupResponse(data []byte) (DiscordIntegrationSetupResponse, error) {
	var r DiscordIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationUpdateParams(data []byte) (DiscordIntegrationUpdateParams, error) {
	var r DiscordIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationUpdateRequest(data []byte) (DiscordIntegrationUpdateRequest, error) {
	var r DiscordIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationUpdateResponse(data []byte) (DiscordIntegrationUpdateResponse, error) {
	var r DiscordIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationCreateRequest(data []byte) (DiscordIntegrationCreateRequest, error) {
	var r DiscordIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationCreateResponse(data []byte) (DiscordIntegrationCreateResponse, error) {
	var r DiscordIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationsListParams(data []byte) (DiscordIntegrationsListParams, error) {
	var r DiscordIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDiscordIntegrationsListResponse(data []byte) (DiscordIntegrationsListResponse, error) {
	var r DiscordIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DiscordIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationDeleteParams(data []byte) (EmailIntegrationDeleteParams, error) {
	var r EmailIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type EmailIntegrationDeleteRequest map[string]interface{}

func UnmarshalEmailIntegrationDeleteRequest(data []byte) (EmailIntegrationDeleteRequest, error) {
	var r EmailIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationDeleteResponse(data []byte) (EmailIntegrationDeleteResponse, error) {
	var r EmailIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationFetchParams(data []byte) (EmailIntegrationFetchParams, error) {
	var r EmailIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationFetchResponse(data []byte) (EmailIntegrationFetchResponse, error) {
	var r EmailIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationSetupParams(data []byte) (EmailIntegrationSetupParams, error) {
	var r EmailIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type EmailIntegrationSetupRequest map[string]interface{}

func UnmarshalEmailIntegrationSetupRequest(data []byte) (EmailIntegrationSetupRequest, error) {
	var r EmailIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationSetupResponse(data []byte) (EmailIntegrationSetupResponse, error) {
	var r EmailIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationUpdateParams(data []byte) (EmailIntegrationUpdateParams, error) {
	var r EmailIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationUpdateRequest(data []byte) (EmailIntegrationUpdateRequest, error) {
	var r EmailIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationUpdateResponse(data []byte) (EmailIntegrationUpdateResponse, error) {
	var r EmailIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationCreateRequest(data []byte) (EmailIntegrationCreateRequest, error) {
	var r EmailIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationCreateResponse(data []byte) (EmailIntegrationCreateResponse, error) {
	var r EmailIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationsListParams(data []byte) (EmailIntegrationsListParams, error) {
	var r EmailIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEmailIntegrationsListResponse(data []byte) (EmailIntegrationsListResponse, error) {
	var r EmailIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EmailIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationDeleteParams(data []byte) (ExtractIntegrationDeleteParams, error) {
	var r ExtractIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ExtractIntegrationDeleteRequest map[string]interface{}

func UnmarshalExtractIntegrationDeleteRequest(data []byte) (ExtractIntegrationDeleteRequest, error) {
	var r ExtractIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationDeleteResponse(data []byte) (ExtractIntegrationDeleteResponse, error) {
	var r ExtractIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationFetchParams(data []byte) (ExtractIntegrationFetchParams, error) {
	var r ExtractIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationFetchResponse(data []byte) (ExtractIntegrationFetchResponse, error) {
	var r ExtractIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationUpdateParams(data []byte) (ExtractIntegrationUpdateParams, error) {
	var r ExtractIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationUpdateRequest(data []byte) (ExtractIntegrationUpdateRequest, error) {
	var r ExtractIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationUpdateResponse(data []byte) (ExtractIntegrationUpdateResponse, error) {
	var r ExtractIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationCreateRequest(data []byte) (ExtractIntegrationCreateRequest, error) {
	var r ExtractIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationCreateResponse(data []byte) (ExtractIntegrationCreateResponse, error) {
	var r ExtractIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationsListParams(data []byte) (ExtractIntegrationsListParams, error) {
	var r ExtractIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtractIntegrationsListResponse(data []byte) (ExtractIntegrationsListResponse, error) {
	var r ExtractIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtractIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationDeleteParams(data []byte) (InstagramIntegrationDeleteParams, error) {
	var r InstagramIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InstagramIntegrationDeleteRequest map[string]interface{}

func UnmarshalInstagramIntegrationDeleteRequest(data []byte) (InstagramIntegrationDeleteRequest, error) {
	var r InstagramIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationDeleteResponse(data []byte) (InstagramIntegrationDeleteResponse, error) {
	var r InstagramIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationFetchParams(data []byte) (InstagramIntegrationFetchParams, error) {
	var r InstagramIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationFetchResponse(data []byte) (InstagramIntegrationFetchResponse, error) {
	var r InstagramIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationSetupParams(data []byte) (InstagramIntegrationSetupParams, error) {
	var r InstagramIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InstagramIntegrationSetupRequest map[string]interface{}

func UnmarshalInstagramIntegrationSetupRequest(data []byte) (InstagramIntegrationSetupRequest, error) {
	var r InstagramIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationSetupResponse(data []byte) (InstagramIntegrationSetupResponse, error) {
	var r InstagramIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationUpdateParams(data []byte) (InstagramIntegrationUpdateParams, error) {
	var r InstagramIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationUpdateRequest(data []byte) (InstagramIntegrationUpdateRequest, error) {
	var r InstagramIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationUpdateResponse(data []byte) (InstagramIntegrationUpdateResponse, error) {
	var r InstagramIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationCreateRequest(data []byte) (InstagramIntegrationCreateRequest, error) {
	var r InstagramIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationCreateResponse(data []byte) (InstagramIntegrationCreateResponse, error) {
	var r InstagramIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationsListParams(data []byte) (InstagramIntegrationsListParams, error) {
	var r InstagramIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstagramIntegrationsListResponse(data []byte) (InstagramIntegrationsListResponse, error) {
	var r InstagramIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstagramIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationDeleteParams(data []byte) (MCPServerIntegrationDeleteParams, error) {
	var r MCPServerIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MCPServerIntegrationDeleteRequest map[string]interface{}

func UnmarshalMCPServerIntegrationDeleteRequest(data []byte) (MCPServerIntegrationDeleteRequest, error) {
	var r MCPServerIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationDeleteResponse(data []byte) (MCPServerIntegrationDeleteResponse, error) {
	var r MCPServerIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationFetchParams(data []byte) (MCPServerIntegrationFetchParams, error) {
	var r MCPServerIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationFetchResponse(data []byte) (MCPServerIntegrationFetchResponse, error) {
	var r MCPServerIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationUpdateParams(data []byte) (MCPServerIntegrationUpdateParams, error) {
	var r MCPServerIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationUpdateRequest(data []byte) (MCPServerIntegrationUpdateRequest, error) {
	var r MCPServerIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationUpdateResponse(data []byte) (MCPServerIntegrationUpdateResponse, error) {
	var r MCPServerIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationCreateRequest(data []byte) (MCPServerIntegrationCreateRequest, error) {
	var r MCPServerIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationCreateResponse(data []byte) (MCPServerIntegrationCreateResponse, error) {
	var r MCPServerIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationsListParams(data []byte) (MCPServerIntegrationsListParams, error) {
	var r MCPServerIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMCPServerIntegrationsListResponse(data []byte) (MCPServerIntegrationsListResponse, error) {
	var r MCPServerIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MCPServerIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationDeleteParams(data []byte) (MessengerIntegrationDeleteParams, error) {
	var r MessengerIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessengerIntegrationDeleteRequest map[string]interface{}

func UnmarshalMessengerIntegrationDeleteRequest(data []byte) (MessengerIntegrationDeleteRequest, error) {
	var r MessengerIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationDeleteResponse(data []byte) (MessengerIntegrationDeleteResponse, error) {
	var r MessengerIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationFetchParams(data []byte) (MessengerIntegrationFetchParams, error) {
	var r MessengerIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationFetchResponse(data []byte) (MessengerIntegrationFetchResponse, error) {
	var r MessengerIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationSetupParams(data []byte) (MessengerIntegrationSetupParams, error) {
	var r MessengerIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessengerIntegrationSetupRequest map[string]interface{}

func UnmarshalMessengerIntegrationSetupRequest(data []byte) (MessengerIntegrationSetupRequest, error) {
	var r MessengerIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationSetupResponse(data []byte) (MessengerIntegrationSetupResponse, error) {
	var r MessengerIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationUpdateParams(data []byte) (MessengerIntegrationUpdateParams, error) {
	var r MessengerIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationUpdateRequest(data []byte) (MessengerIntegrationUpdateRequest, error) {
	var r MessengerIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationUpdateResponse(data []byte) (MessengerIntegrationUpdateResponse, error) {
	var r MessengerIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationCreateRequest(data []byte) (MessengerIntegrationCreateRequest, error) {
	var r MessengerIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationCreateResponse(data []byte) (MessengerIntegrationCreateResponse, error) {
	var r MessengerIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationsListParams(data []byte) (MessengerIntegrationsListParams, error) {
	var r MessengerIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessengerIntegrationsListResponse(data []byte) (MessengerIntegrationsListResponse, error) {
	var r MessengerIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessengerIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationDeleteParams(data []byte) (NotionIntegrationDeleteParams, error) {
	var r NotionIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NotionIntegrationDeleteRequest map[string]interface{}

func UnmarshalNotionIntegrationDeleteRequest(data []byte) (NotionIntegrationDeleteRequest, error) {
	var r NotionIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationDeleteResponse(data []byte) (NotionIntegrationDeleteResponse, error) {
	var r NotionIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationFetchParams(data []byte) (NotionIntegrationFetchParams, error) {
	var r NotionIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationFetchResponse(data []byte) (NotionIntegrationFetchResponse, error) {
	var r NotionIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationSyncParams(data []byte) (NotionIntegrationSyncParams, error) {
	var r NotionIntegrationSyncParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationSyncParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NotionIntegrationSyncRequest map[string]interface{}

func UnmarshalNotionIntegrationSyncRequest(data []byte) (NotionIntegrationSyncRequest, error) {
	var r NotionIntegrationSyncRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationSyncRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationSyncResponse(data []byte) (NotionIntegrationSyncResponse, error) {
	var r NotionIntegrationSyncResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationSyncResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationUpdateParams(data []byte) (NotionIntegrationUpdateParams, error) {
	var r NotionIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationUpdateRequest(data []byte) (NotionIntegrationUpdateRequest, error) {
	var r NotionIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationUpdateResponse(data []byte) (NotionIntegrationUpdateResponse, error) {
	var r NotionIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationCreateRequest(data []byte) (NotionIntegrationCreateRequest, error) {
	var r NotionIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationCreateResponse(data []byte) (NotionIntegrationCreateResponse, error) {
	var r NotionIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationsListParams(data []byte) (NotionIntegrationsListParams, error) {
	var r NotionIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNotionIntegrationsListResponse(data []byte) (NotionIntegrationsListResponse, error) {
	var r NotionIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NotionIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationDeleteParams(data []byte) (SitemapIntegrationDeleteParams, error) {
	var r SitemapIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SitemapIntegrationDeleteRequest map[string]interface{}

func UnmarshalSitemapIntegrationDeleteRequest(data []byte) (SitemapIntegrationDeleteRequest, error) {
	var r SitemapIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationDeleteResponse(data []byte) (SitemapIntegrationDeleteResponse, error) {
	var r SitemapIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationFetchParams(data []byte) (SitemapIntegrationFetchParams, error) {
	var r SitemapIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationFetchResponse(data []byte) (SitemapIntegrationFetchResponse, error) {
	var r SitemapIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationSyncParams(data []byte) (SitemapIntegrationSyncParams, error) {
	var r SitemapIntegrationSyncParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationSyncParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SitemapIntegrationSyncRequest map[string]interface{}

func UnmarshalSitemapIntegrationSyncRequest(data []byte) (SitemapIntegrationSyncRequest, error) {
	var r SitemapIntegrationSyncRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationSyncRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationSyncResponse(data []byte) (SitemapIntegrationSyncResponse, error) {
	var r SitemapIntegrationSyncResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationSyncResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationUpdateParams(data []byte) (SitemapIntegrationUpdateParams, error) {
	var r SitemapIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationUpdateRequest(data []byte) (SitemapIntegrationUpdateRequest, error) {
	var r SitemapIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationUpdateResponse(data []byte) (SitemapIntegrationUpdateResponse, error) {
	var r SitemapIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationCreateRequest(data []byte) (SitemapIntegrationCreateRequest, error) {
	var r SitemapIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationCreateResponse(data []byte) (SitemapIntegrationCreateResponse, error) {
	var r SitemapIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationsListParams(data []byte) (SitemapIntegrationsListParams, error) {
	var r SitemapIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSitemapIntegrationsListResponse(data []byte) (SitemapIntegrationsListResponse, error) {
	var r SitemapIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SitemapIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationDeleteParams(data []byte) (SlackIntegrationDeleteParams, error) {
	var r SlackIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SlackIntegrationDeleteRequest map[string]interface{}

func UnmarshalSlackIntegrationDeleteRequest(data []byte) (SlackIntegrationDeleteRequest, error) {
	var r SlackIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationDeleteResponse(data []byte) (SlackIntegrationDeleteResponse, error) {
	var r SlackIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationFetchParams(data []byte) (SlackIntegrationFetchParams, error) {
	var r SlackIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationFetchResponse(data []byte) (SlackIntegrationFetchResponse, error) {
	var r SlackIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationSetupParams(data []byte) (SlackIntegrationSetupParams, error) {
	var r SlackIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SlackIntegrationSetupRequest map[string]interface{}

func UnmarshalSlackIntegrationSetupRequest(data []byte) (SlackIntegrationSetupRequest, error) {
	var r SlackIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationSetupResponse(data []byte) (SlackIntegrationSetupResponse, error) {
	var r SlackIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationUpdateParams(data []byte) (SlackIntegrationUpdateParams, error) {
	var r SlackIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationUpdateRequest(data []byte) (SlackIntegrationUpdateRequest, error) {
	var r SlackIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationUpdateResponse(data []byte) (SlackIntegrationUpdateResponse, error) {
	var r SlackIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationCreateRequest(data []byte) (SlackIntegrationCreateRequest, error) {
	var r SlackIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationCreateResponse(data []byte) (SlackIntegrationCreateResponse, error) {
	var r SlackIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationsListParams(data []byte) (SlackIntegrationsListParams, error) {
	var r SlackIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSlackIntegrationsListResponse(data []byte) (SlackIntegrationsListResponse, error) {
	var r SlackIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SlackIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationDeleteParams(data []byte) (SupportIntegrationDeleteParams, error) {
	var r SupportIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SupportIntegrationDeleteRequest map[string]interface{}

func UnmarshalSupportIntegrationDeleteRequest(data []byte) (SupportIntegrationDeleteRequest, error) {
	var r SupportIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationDeleteResponse(data []byte) (SupportIntegrationDeleteResponse, error) {
	var r SupportIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationFetchParams(data []byte) (SupportIntegrationFetchParams, error) {
	var r SupportIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationFetchResponse(data []byte) (SupportIntegrationFetchResponse, error) {
	var r SupportIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationUpdateParams(data []byte) (SupportIntegrationUpdateParams, error) {
	var r SupportIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationUpdateRequest(data []byte) (SupportIntegrationUpdateRequest, error) {
	var r SupportIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationUpdateResponse(data []byte) (SupportIntegrationUpdateResponse, error) {
	var r SupportIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationCreateRequest(data []byte) (SupportIntegrationCreateRequest, error) {
	var r SupportIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationCreateResponse(data []byte) (SupportIntegrationCreateResponse, error) {
	var r SupportIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationsListParams(data []byte) (SupportIntegrationsListParams, error) {
	var r SupportIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSupportIntegrationsListResponse(data []byte) (SupportIntegrationsListResponse, error) {
	var r SupportIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SupportIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationDeleteParams(data []byte) (TelegramIntegrationDeleteParams, error) {
	var r TelegramIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TelegramIntegrationDeleteRequest map[string]interface{}

func UnmarshalTelegramIntegrationDeleteRequest(data []byte) (TelegramIntegrationDeleteRequest, error) {
	var r TelegramIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationDeleteResponse(data []byte) (TelegramIntegrationDeleteResponse, error) {
	var r TelegramIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationFetchParams(data []byte) (TelegramIntegrationFetchParams, error) {
	var r TelegramIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationFetchResponse(data []byte) (TelegramIntegrationFetchResponse, error) {
	var r TelegramIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationSetupParams(data []byte) (TelegramIntegrationSetupParams, error) {
	var r TelegramIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TelegramIntegrationSetupRequest map[string]interface{}

func UnmarshalTelegramIntegrationSetupRequest(data []byte) (TelegramIntegrationSetupRequest, error) {
	var r TelegramIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationSetupResponse(data []byte) (TelegramIntegrationSetupResponse, error) {
	var r TelegramIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationUpdateParams(data []byte) (TelegramIntegrationUpdateParams, error) {
	var r TelegramIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationUpdateRequest(data []byte) (TelegramIntegrationUpdateRequest, error) {
	var r TelegramIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationUpdateResponse(data []byte) (TelegramIntegrationUpdateResponse, error) {
	var r TelegramIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationCreateRequest(data []byte) (TelegramIntegrationCreateRequest, error) {
	var r TelegramIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationCreateResponse(data []byte) (TelegramIntegrationCreateResponse, error) {
	var r TelegramIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationsListParams(data []byte) (TelegramIntegrationsListParams, error) {
	var r TelegramIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTelegramIntegrationsListResponse(data []byte) (TelegramIntegrationsListResponse, error) {
	var r TelegramIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationDeleteParams(data []byte) (TriggerIntegrationDeleteParams, error) {
	var r TriggerIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TriggerIntegrationDeleteRequest map[string]interface{}

func UnmarshalTriggerIntegrationDeleteRequest(data []byte) (TriggerIntegrationDeleteRequest, error) {
	var r TriggerIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationDeleteResponse(data []byte) (TriggerIntegrationDeleteResponse, error) {
	var r TriggerIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationFetchParams(data []byte) (TriggerIntegrationFetchParams, error) {
	var r TriggerIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationFetchResponse(data []byte) (TriggerIntegrationFetchResponse, error) {
	var r TriggerIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationInvokeParams(data []byte) (TriggerIntegrationInvokeParams, error) {
	var r TriggerIntegrationInvokeParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationInvokeParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TriggerIntegrationInvokeRequest map[string]interface{}

func UnmarshalTriggerIntegrationInvokeRequest(data []byte) (TriggerIntegrationInvokeRequest, error) {
	var r TriggerIntegrationInvokeRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationInvokeRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationInvokeResponse(data []byte) (TriggerIntegrationInvokeResponse, error) {
	var r TriggerIntegrationInvokeResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationInvokeResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationSetupParams(data []byte) (TriggerIntegrationSetupParams, error) {
	var r TriggerIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TriggerIntegrationSetupRequest map[string]interface{}

func UnmarshalTriggerIntegrationSetupRequest(data []byte) (TriggerIntegrationSetupRequest, error) {
	var r TriggerIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationSetupResponse(data []byte) (TriggerIntegrationSetupResponse, error) {
	var r TriggerIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationUpdateParams(data []byte) (TriggerIntegrationUpdateParams, error) {
	var r TriggerIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationUpdateRequest(data []byte) (TriggerIntegrationUpdateRequest, error) {
	var r TriggerIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationUpdateResponse(data []byte) (TriggerIntegrationUpdateResponse, error) {
	var r TriggerIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationCreateRequest(data []byte) (TriggerIntegrationCreateRequest, error) {
	var r TriggerIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationCreateResponse(data []byte) (TriggerIntegrationCreateResponse, error) {
	var r TriggerIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationsListParams(data []byte) (TriggerIntegrationsListParams, error) {
	var r TriggerIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerIntegrationsListResponse(data []byte) (TriggerIntegrationsListResponse, error) {
	var r TriggerIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationDeleteParams(data []byte) (TwilioIntegrationDeleteParams, error) {
	var r TwilioIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TwilioIntegrationDeleteRequest map[string]interface{}

func UnmarshalTwilioIntegrationDeleteRequest(data []byte) (TwilioIntegrationDeleteRequest, error) {
	var r TwilioIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationDeleteResponse(data []byte) (TwilioIntegrationDeleteResponse, error) {
	var r TwilioIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationFetchParams(data []byte) (TwilioIntegrationFetchParams, error) {
	var r TwilioIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationFetchResponse(data []byte) (TwilioIntegrationFetchResponse, error) {
	var r TwilioIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationSetupParams(data []byte) (TwilioIntegrationSetupParams, error) {
	var r TwilioIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TwilioIntegrationSetupRequest map[string]interface{}

func UnmarshalTwilioIntegrationSetupRequest(data []byte) (TwilioIntegrationSetupRequest, error) {
	var r TwilioIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationSetupResponse(data []byte) (TwilioIntegrationSetupResponse, error) {
	var r TwilioIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationUpdateParams(data []byte) (TwilioIntegrationUpdateParams, error) {
	var r TwilioIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationUpdateRequest(data []byte) (TwilioIntegrationUpdateRequest, error) {
	var r TwilioIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationUpdateResponse(data []byte) (TwilioIntegrationUpdateResponse, error) {
	var r TwilioIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationCreateRequest(data []byte) (TwilioIntegrationCreateRequest, error) {
	var r TwilioIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationCreateResponse(data []byte) (TwilioIntegrationCreateResponse, error) {
	var r TwilioIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationsListParams(data []byte) (TwilioIntegrationsListParams, error) {
	var r TwilioIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTwilioIntegrationsListResponse(data []byte) (TwilioIntegrationsListResponse, error) {
	var r TwilioIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TwilioIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationDeleteParams(data []byte) (WhatsAppIntegrationDeleteParams, error) {
	var r WhatsAppIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WhatsAppIntegrationDeleteRequest map[string]interface{}

func UnmarshalWhatsAppIntegrationDeleteRequest(data []byte) (WhatsAppIntegrationDeleteRequest, error) {
	var r WhatsAppIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationDeleteResponse(data []byte) (WhatsAppIntegrationDeleteResponse, error) {
	var r WhatsAppIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationFetchParams(data []byte) (WhatsAppIntegrationFetchParams, error) {
	var r WhatsAppIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationFetchResponse(data []byte) (WhatsAppIntegrationFetchResponse, error) {
	var r WhatsAppIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationSetupParams(data []byte) (WhatsAppIntegrationSetupParams, error) {
	var r WhatsAppIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WhatsAppIntegrationSetupRequest map[string]interface{}

func UnmarshalWhatsAppIntegrationSetupRequest(data []byte) (WhatsAppIntegrationSetupRequest, error) {
	var r WhatsAppIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationSetupResponse(data []byte) (WhatsAppIntegrationSetupResponse, error) {
	var r WhatsAppIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationUpdateParams(data []byte) (WhatsAppIntegrationUpdateParams, error) {
	var r WhatsAppIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationUpdateRequest(data []byte) (WhatsAppIntegrationUpdateRequest, error) {
	var r WhatsAppIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationUpdateResponse(data []byte) (WhatsAppIntegrationUpdateResponse, error) {
	var r WhatsAppIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationCreateRequest(data []byte) (WhatsAppIntegrationCreateRequest, error) {
	var r WhatsAppIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationCreateResponse(data []byte) (WhatsAppIntegrationCreateResponse, error) {
	var r WhatsAppIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationsListParams(data []byte) (WhatsAppIntegrationsListParams, error) {
	var r WhatsAppIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWhatsAppIntegrationsListResponse(data []byte) (WhatsAppIntegrationsListResponse, error) {
	var r WhatsAppIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WhatsAppIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationDeleteParams(data []byte) (WidgetIntegrationDeleteParams, error) {
	var r WidgetIntegrationDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WidgetIntegrationDeleteRequest map[string]interface{}

func UnmarshalWidgetIntegrationDeleteRequest(data []byte) (WidgetIntegrationDeleteRequest, error) {
	var r WidgetIntegrationDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationDeleteResponse(data []byte) (WidgetIntegrationDeleteResponse, error) {
	var r WidgetIntegrationDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationFetchParams(data []byte) (WidgetIntegrationFetchParams, error) {
	var r WidgetIntegrationFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationFetchResponse(data []byte) (WidgetIntegrationFetchResponse, error) {
	var r WidgetIntegrationFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationSetupParams(data []byte) (WidgetIntegrationSetupParams, error) {
	var r WidgetIntegrationSetupParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationSetupParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WidgetIntegrationSetupRequest map[string]interface{}

func UnmarshalWidgetIntegrationSetupRequest(data []byte) (WidgetIntegrationSetupRequest, error) {
	var r WidgetIntegrationSetupRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationSetupRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationSetupResponse(data []byte) (WidgetIntegrationSetupResponse, error) {
	var r WidgetIntegrationSetupResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationSetupResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationUpdateParams(data []byte) (WidgetIntegrationUpdateParams, error) {
	var r WidgetIntegrationUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationUpdateRequest(data []byte) (WidgetIntegrationUpdateRequest, error) {
	var r WidgetIntegrationUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationUpdateResponse(data []byte) (WidgetIntegrationUpdateResponse, error) {
	var r WidgetIntegrationUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationCreateRequest(data []byte) (WidgetIntegrationCreateRequest, error) {
	var r WidgetIntegrationCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationCreateResponse(data []byte) (WidgetIntegrationCreateResponse, error) {
	var r WidgetIntegrationCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationsListParams(data []byte) (WidgetIntegrationsListParams, error) {
	var r WidgetIntegrationsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalWidgetIntegrationsListResponse(data []byte) (WidgetIntegrationsListResponse, error) {
	var r WidgetIntegrationsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WidgetIntegrationsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMagicFromPromptGenerateParams(data []byte) (MagicFromPromptGenerateParams, error) {
	var r MagicFromPromptGenerateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MagicFromPromptGenerateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMagicFromPromptGenerateRequest(data []byte) (MagicFromPromptGenerateRequest, error) {
	var r MagicFromPromptGenerateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MagicFromPromptGenerateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMagicFromPromptGenerateResponse(data []byte) (MagicFromPromptGenerateResponse, error) {
	var r MagicFromPromptGenerateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MagicFromPromptGenerateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMagicPromptsListParams(data []byte) (MagicPromptsListParams, error) {
	var r MagicPromptsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MagicPromptsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMagicPromptsListResponse(data []byte) (MagicPromptsListResponse, error) {
	var r MagicPromptsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MagicPromptsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoryDeleteParams(data []byte) (MemoryDeleteParams, error) {
	var r MemoryDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MemoryDeleteRequest map[string]interface{}

func UnmarshalMemoryDeleteRequest(data []byte) (MemoryDeleteRequest, error) {
	var r MemoryDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoryDeleteResponse(data []byte) (MemoryDeleteResponse, error) {
	var r MemoryDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoryFetchParams(data []byte) (MemoryFetchParams, error) {
	var r MemoryFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoryFetchResponse(data []byte) (MemoryFetchResponse, error) {
	var r MemoryFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoryUpdateParams(data []byte) (MemoryUpdateParams, error) {
	var r MemoryUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoryUpdateRequest(data []byte) (MemoryUpdateRequest, error) {
	var r MemoryUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoryUpdateResponse(data []byte) (MemoryUpdateResponse, error) {
	var r MemoryUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoryCreateRequest(data []byte) (MemoryCreateRequest, error) {
	var r MemoryCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoryCreateResponse(data []byte) (MemoryCreateResponse, error) {
	var r MemoryCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoryCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoriesExportParams(data []byte) (MemoriesExportParams, error) {
	var r MemoriesExportParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoriesExportParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoriesExportResponse(data []byte) (MemoriesExportResponse, error) {
	var r MemoriesExportResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoriesExportResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoriesListParams(data []byte) (MemoriesListParams, error) {
	var r MemoriesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoriesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemoriesListResponse(data []byte) (MemoriesListResponse, error) {
	var r MemoriesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemoriesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemorySearchRequest(data []byte) (MemorySearchRequest, error) {
	var r MemorySearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemorySearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMemorySearchResponse(data []byte) (MemorySearchResponse, error) {
	var r MemorySearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MemorySearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserDeleteParams(data []byte) (PartnerUserDeleteParams, error) {
	var r PartnerUserDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PartnerUserDeleteRequest map[string]interface{}

func UnmarshalPartnerUserDeleteRequest(data []byte) (PartnerUserDeleteRequest, error) {
	var r PartnerUserDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserDeleteResponse(data []byte) (PartnerUserDeleteResponse, error) {
	var r PartnerUserDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserFetchParams(data []byte) (PartnerUserFetchParams, error) {
	var r PartnerUserFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserFetchResponse(data []byte) (PartnerUserFetchResponse, error) {
	var r PartnerUserFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserTokenDeleteParams(data []byte) (PartnerUserTokenDeleteParams, error) {
	var r PartnerUserTokenDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserTokenDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PartnerUserTokenDeleteRequest map[string]interface{}

func UnmarshalPartnerUserTokenDeleteRequest(data []byte) (PartnerUserTokenDeleteRequest, error) {
	var r PartnerUserTokenDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserTokenDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserTokenDeleteResponse(data []byte) (PartnerUserTokenDeleteResponse, error) {
	var r PartnerUserTokenDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserTokenDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserTokenCreateParams(data []byte) (PartnerUserTokenCreateParams, error) {
	var r PartnerUserTokenCreateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserTokenCreateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PartnerUserTokenCreateRequest map[string]interface{}

func UnmarshalPartnerUserTokenCreateRequest(data []byte) (PartnerUserTokenCreateRequest, error) {
	var r PartnerUserTokenCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserTokenCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserTokenCreateResponse(data []byte) (PartnerUserTokenCreateResponse, error) {
	var r PartnerUserTokenCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserTokenCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserTokensListParams(data []byte) (PartnerUserTokensListParams, error) {
	var r PartnerUserTokensListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserTokensListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserTokensListResponse(data []byte) (PartnerUserTokensListResponse, error) {
	var r PartnerUserTokensListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserTokensListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserUpdateParams(data []byte) (PartnerUserUpdateParams, error) {
	var r PartnerUserUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserUpdateRequest(data []byte) (PartnerUserUpdateRequest, error) {
	var r PartnerUserUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserUpdateResponse(data []byte) (PartnerUserUpdateResponse, error) {
	var r PartnerUserUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserCreateRequest(data []byte) (PartnerUserCreateRequest, error) {
	var r PartnerUserCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUserCreateResponse(data []byte) (PartnerUserCreateResponse, error) {
	var r PartnerUserCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUserCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUsersListParams(data []byte) (PartnerUsersListParams, error) {
	var r PartnerUsersListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUsersListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPartnerUsersListResponse(data []byte) (PartnerUsersListResponse, error) {
	var r PartnerUsersListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PartnerUsersListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformAbilitiesListParams(data []byte) (PlatformAbilitiesListParams, error) {
	var r PlatformAbilitiesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformAbilitiesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformAbilitiesListResponse(data []byte) (PlatformAbilitiesListResponse, error) {
	var r PlatformAbilitiesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformAbilitiesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformActionsListParams(data []byte) (PlatformActionsListParams, error) {
	var r PlatformActionsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformActionsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformActionsListResponse(data []byte) (PlatformActionsListResponse, error) {
	var r PlatformActionsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformActionsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformDocFetchParams(data []byte) (PlatformDocFetchParams, error) {
	var r PlatformDocFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformDocFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformDocFetchResponse(data []byte) (PlatformDocFetchResponse, error) {
	var r PlatformDocFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformDocFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformDocsListParams(data []byte) (PlatformDocsListParams, error) {
	var r PlatformDocsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformDocsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformDocsListResponse(data []byte) (PlatformDocsListResponse, error) {
	var r PlatformDocsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformDocsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformDocsSearchRequest(data []byte) (PlatformDocsSearchRequest, error) {
	var r PlatformDocsSearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformDocsSearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformDocsSearchResponse(data []byte) (PlatformDocsSearchResponse, error) {
	var r PlatformDocsSearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformDocsSearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformExampleCloneParams(data []byte) (PlatformExampleCloneParams, error) {
	var r PlatformExampleCloneParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformExampleCloneParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PlatformExampleCloneRequest map[string]interface{}

func UnmarshalPlatformExampleCloneRequest(data []byte) (PlatformExampleCloneRequest, error) {
	var r PlatformExampleCloneRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformExampleCloneRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformExampleCloneResponse(data []byte) (PlatformExampleCloneResponse, error) {
	var r PlatformExampleCloneResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformExampleCloneResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformExampleFetchParams(data []byte) (PlatformExampleFetchParams, error) {
	var r PlatformExampleFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformExampleFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformExampleFetchResponse(data []byte) (PlatformExampleFetchResponse, error) {
	var r PlatformExampleFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformExampleFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformExamplesListParams(data []byte) (PlatformExamplesListParams, error) {
	var r PlatformExamplesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformExamplesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformExamplesListResponse(data []byte) (PlatformExamplesListResponse, error) {
	var r PlatformExamplesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformExamplesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformExamplesSearchRequest(data []byte) (PlatformExamplesSearchRequest, error) {
	var r PlatformExamplesSearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformExamplesSearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformExamplesSearchResponse(data []byte) (PlatformExamplesSearchResponse, error) {
	var r PlatformExamplesSearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformExamplesSearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformGuideFetchParams(data []byte) (PlatformGuideFetchParams, error) {
	var r PlatformGuideFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformGuideFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformGuideFetchResponse(data []byte) (PlatformGuideFetchResponse, error) {
	var r PlatformGuideFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformGuideFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformGuidesListParams(data []byte) (PlatformGuidesListParams, error) {
	var r PlatformGuidesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformGuidesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformGuidesListResponse(data []byte) (PlatformGuidesListResponse, error) {
	var r PlatformGuidesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformGuidesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformGuidesSearchRequest(data []byte) (PlatformGuidesSearchRequest, error) {
	var r PlatformGuidesSearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformGuidesSearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformGuidesSearchResponse(data []byte) (PlatformGuidesSearchResponse, error) {
	var r PlatformGuidesSearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformGuidesSearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformManualFetchParams(data []byte) (PlatformManualFetchParams, error) {
	var r PlatformManualFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformManualFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformManualFetchResponse(data []byte) (PlatformManualFetchResponse, error) {
	var r PlatformManualFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformManualFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformManualsListParams(data []byte) (PlatformManualsListParams, error) {
	var r PlatformManualsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformManualsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformManualsListResponse(data []byte) (PlatformManualsListResponse, error) {
	var r PlatformManualsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformManualsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformManualsSearchRequest(data []byte) (PlatformManualsSearchRequest, error) {
	var r PlatformManualsSearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformManualsSearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformManualsSearchResponse(data []byte) (PlatformManualsSearchResponse, error) {
	var r PlatformManualsSearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformManualsSearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformModelsListParams(data []byte) (PlatformModelsListParams, error) {
	var r PlatformModelsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformModelsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformModelsListResponse(data []byte) (PlatformModelsListResponse, error) {
	var r PlatformModelsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformModelsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformSecretsListParams(data []byte) (PlatformSecretsListParams, error) {
	var r PlatformSecretsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformSecretsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformSecretsListResponse(data []byte) (PlatformSecretsListResponse, error) {
	var r PlatformSecretsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformSecretsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformTutorialFetchParams(data []byte) (PlatformTutorialFetchParams, error) {
	var r PlatformTutorialFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformTutorialFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformTutorialFetchResponse(data []byte) (PlatformTutorialFetchResponse, error) {
	var r PlatformTutorialFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformTutorialFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformTutorialsListParams(data []byte) (PlatformTutorialsListParams, error) {
	var r PlatformTutorialsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformTutorialsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformTutorialsListResponse(data []byte) (PlatformTutorialsListResponse, error) {
	var r PlatformTutorialsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformTutorialsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformTutorialsSearchRequest(data []byte) (PlatformTutorialsSearchRequest, error) {
	var r PlatformTutorialsSearchRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformTutorialsSearchRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPlatformTutorialsSearchResponse(data []byte) (PlatformTutorialsSearchResponse, error) {
	var r PlatformTutorialsSearchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlatformTutorialsSearchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyDeleteParams(data []byte) (PolicyDeleteParams, error) {
	var r PolicyDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PolicyDeleteRequest map[string]interface{}

func UnmarshalPolicyDeleteRequest(data []byte) (PolicyDeleteRequest, error) {
	var r PolicyDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyDeleteResponse(data []byte) (PolicyDeleteResponse, error) {
	var r PolicyDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyFetchParams(data []byte) (PolicyFetchParams, error) {
	var r PolicyFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyFetchResponse(data []byte) (PolicyFetchResponse, error) {
	var r PolicyFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyUpdateParams(data []byte) (PolicyUpdateParams, error) {
	var r PolicyUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyUpdateRequest(data []byte) (PolicyUpdateRequest, error) {
	var r PolicyUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyUpdateResponse(data []byte) (PolicyUpdateResponse, error) {
	var r PolicyUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyCreateRequest(data []byte) (PolicyCreateRequest, error) {
	var r PolicyCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyCreateResponse(data []byte) (PolicyCreateResponse, error) {
	var r PolicyCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPoliciesListParams(data []byte) (PoliciesListParams, error) {
	var r PoliciesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PoliciesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPoliciesListResponse(data []byte) (PoliciesListResponse, error) {
	var r PoliciesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PoliciesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalDeleteParams(data []byte) (PortalDeleteParams, error) {
	var r PortalDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PortalDeleteRequest map[string]interface{}

func UnmarshalPortalDeleteRequest(data []byte) (PortalDeleteRequest, error) {
	var r PortalDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalDeleteResponse(data []byte) (PortalDeleteResponse, error) {
	var r PortalDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalFetchParams(data []byte) (PortalFetchParams, error) {
	var r PortalFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalFetchResponse(data []byte) (PortalFetchResponse, error) {
	var r PortalFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalUpdateParams(data []byte) (PortalUpdateParams, error) {
	var r PortalUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalUpdateRequest(data []byte) (PortalUpdateRequest, error) {
	var r PortalUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalUpdateResponse(data []byte) (PortalUpdateResponse, error) {
	var r PortalUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalCreateRequest(data []byte) (PortalCreateRequest, error) {
	var r PortalCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalCreateResponse(data []byte) (PortalCreateResponse, error) {
	var r PortalCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalsListParams(data []byte) (PortalsListParams, error) {
	var r PortalsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPortalsListResponse(data []byte) (PortalsListResponse, error) {
	var r PortalsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PortalsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretAuthenticateParams(data []byte) (SecretAuthenticateParams, error) {
	var r SecretAuthenticateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretAuthenticateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SecretAuthenticateRequest map[string]interface{}

func UnmarshalSecretAuthenticateRequest(data []byte) (SecretAuthenticateRequest, error) {
	var r SecretAuthenticateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretAuthenticateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretAuthenticateResponse(data []byte) (SecretAuthenticateResponse, error) {
	var r SecretAuthenticateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretAuthenticateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretDeleteParams(data []byte) (SecretDeleteParams, error) {
	var r SecretDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SecretDeleteRequest map[string]interface{}

func UnmarshalSecretDeleteRequest(data []byte) (SecretDeleteRequest, error) {
	var r SecretDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretDeleteResponse(data []byte) (SecretDeleteResponse, error) {
	var r SecretDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretFetchParams(data []byte) (SecretFetchParams, error) {
	var r SecretFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretFetchResponse(data []byte) (SecretFetchResponse, error) {
	var r SecretFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretRevokeParams(data []byte) (SecretRevokeParams, error) {
	var r SecretRevokeParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretRevokeParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SecretRevokeRequest map[string]interface{}

func UnmarshalSecretRevokeRequest(data []byte) (SecretRevokeRequest, error) {
	var r SecretRevokeRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretRevokeRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretRevokeResponse(data []byte) (SecretRevokeResponse, error) {
	var r SecretRevokeResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretRevokeResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretUpdateParams(data []byte) (SecretUpdateParams, error) {
	var r SecretUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretUpdateRequest(data []byte) (SecretUpdateRequest, error) {
	var r SecretUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretUpdateResponse(data []byte) (SecretUpdateResponse, error) {
	var r SecretUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretVerifyParams(data []byte) (SecretVerifyParams, error) {
	var r SecretVerifyParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretVerifyParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SecretVerifyRequest map[string]interface{}

func UnmarshalSecretVerifyRequest(data []byte) (SecretVerifyRequest, error) {
	var r SecretVerifyRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretVerifyRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretVerifyResponse(data []byte) (SecretVerifyResponse, error) {
	var r SecretVerifyResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretVerifyResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretCreateRequest(data []byte) (SecretCreateRequest, error) {
	var r SecretCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretCreateResponse(data []byte) (SecretCreateResponse, error) {
	var r SecretCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretsListParams(data []byte) (SecretsListParams, error) {
	var r SecretsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretsListResponse(data []byte) (SecretsListResponse, error) {
	var r SecretsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityDeleteParams(data []byte) (SkillsetAbilityDeleteParams, error) {
	var r SkillsetAbilityDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SkillsetAbilityDeleteRequest map[string]interface{}

func UnmarshalSkillsetAbilityDeleteRequest(data []byte) (SkillsetAbilityDeleteRequest, error) {
	var r SkillsetAbilityDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityDeleteResponse(data []byte) (SkillsetAbilityDeleteResponse, error) {
	var r SkillsetAbilityDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityExecuteParams(data []byte) (SkillsetAbilityExecuteParams, error) {
	var r SkillsetAbilityExecuteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityExecuteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityExecuteRequest(data []byte) (SkillsetAbilityExecuteRequest, error) {
	var r SkillsetAbilityExecuteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityExecuteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityExecuteResponse(data []byte) (SkillsetAbilityExecuteResponse, error) {
	var r SkillsetAbilityExecuteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityExecuteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityFetchParams(data []byte) (SkillsetAbilityFetchParams, error) {
	var r SkillsetAbilityFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityFetchResponse(data []byte) (SkillsetAbilityFetchResponse, error) {
	var r SkillsetAbilityFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityUpdateParams(data []byte) (SkillsetAbilityUpdateParams, error) {
	var r SkillsetAbilityUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityUpdateRequest(data []byte) (SkillsetAbilityUpdateRequest, error) {
	var r SkillsetAbilityUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityUpdateResponse(data []byte) (SkillsetAbilityUpdateResponse, error) {
	var r SkillsetAbilityUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityCreateParams(data []byte) (SkillsetAbilityCreateParams, error) {
	var r SkillsetAbilityCreateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityCreateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityCreateRequest(data []byte) (SkillsetAbilityCreateRequest, error) {
	var r SkillsetAbilityCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilityCreateResponse(data []byte) (SkillsetAbilityCreateResponse, error) {
	var r SkillsetAbilityCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilityCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilitiesExportParams(data []byte) (SkillsetAbilitiesExportParams, error) {
	var r SkillsetAbilitiesExportParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilitiesExportParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilitiesExportResponse(data []byte) (SkillsetAbilitiesExportResponse, error) {
	var r SkillsetAbilitiesExportResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilitiesExportResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilitiesListParams(data []byte) (SkillsetAbilitiesListParams, error) {
	var r SkillsetAbilitiesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilitiesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetAbilitiesListResponse(data []byte) (SkillsetAbilitiesListResponse, error) {
	var r SkillsetAbilitiesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetAbilitiesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetDeleteParams(data []byte) (SkillsetDeleteParams, error) {
	var r SkillsetDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SkillsetDeleteRequest map[string]interface{}

func UnmarshalSkillsetDeleteRequest(data []byte) (SkillsetDeleteRequest, error) {
	var r SkillsetDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetDeleteResponse(data []byte) (SkillsetDeleteResponse, error) {
	var r SkillsetDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetFetchParams(data []byte) (SkillsetFetchParams, error) {
	var r SkillsetFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetFetchResponse(data []byte) (SkillsetFetchResponse, error) {
	var r SkillsetFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetUpdateParams(data []byte) (SkillsetUpdateParams, error) {
	var r SkillsetUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetUpdateRequest(data []byte) (SkillsetUpdateRequest, error) {
	var r SkillsetUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetUpdateResponse(data []byte) (SkillsetUpdateResponse, error) {
	var r SkillsetUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetCreateRequest(data []byte) (SkillsetCreateRequest, error) {
	var r SkillsetCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetCreateResponse(data []byte) (SkillsetCreateResponse, error) {
	var r SkillsetCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetsListParams(data []byte) (SkillsetsListParams, error) {
	var r SkillsetsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetsListResponse(data []byte) (SkillsetsListResponse, error) {
	var r SkillsetsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpaceFetchParams(data []byte) (SpaceFetchParams, error) {
	var r SpaceFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpaceFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpaceFetchResponse(data []byte) (SpaceFetchResponse, error) {
	var r SpaceFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpaceFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpaceUpdateParams(data []byte) (SpaceUpdateParams, error) {
	var r SpaceUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpaceUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpaceUpdateRequest(data []byte) (SpaceUpdateRequest, error) {
	var r SpaceUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpaceUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpaceUpdateResponse(data []byte) (SpaceUpdateResponse, error) {
	var r SpaceUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpaceUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpaceCreateRequest(data []byte) (SpaceCreateRequest, error) {
	var r SpaceCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpaceCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpaceCreateResponse(data []byte) (SpaceCreateResponse, error) {
	var r SpaceCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpaceCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpacesExportParams(data []byte) (SpacesExportParams, error) {
	var r SpacesExportParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpacesExportParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpacesExportResponse(data []byte) (SpacesExportResponse, error) {
	var r SpacesExportResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpacesExportResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpacesListParams(data []byte) (SpacesListParams, error) {
	var r SpacesListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpacesListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSpacesListResponse(data []byte) (SpacesListResponse, error) {
	var r SpacesListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SpacesListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskDeleteParams(data []byte) (TaskDeleteParams, error) {
	var r TaskDeleteParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskDeleteParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TaskDeleteRequest map[string]interface{}

func UnmarshalTaskDeleteRequest(data []byte) (TaskDeleteRequest, error) {
	var r TaskDeleteRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskDeleteRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskDeleteResponse(data []byte) (TaskDeleteResponse, error) {
	var r TaskDeleteResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskDeleteResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskFetchParams(data []byte) (TaskFetchParams, error) {
	var r TaskFetchParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskFetchParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskFetchResponse(data []byte) (TaskFetchResponse, error) {
	var r TaskFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskTriggerParams(data []byte) (TaskTriggerParams, error) {
	var r TaskTriggerParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskTriggerParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TaskTriggerRequest map[string]interface{}

func UnmarshalTaskTriggerRequest(data []byte) (TaskTriggerRequest, error) {
	var r TaskTriggerRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskTriggerRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskTriggerResponse(data []byte) (TaskTriggerResponse, error) {
	var r TaskTriggerResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskTriggerResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskUpdateParams(data []byte) (TaskUpdateParams, error) {
	var r TaskUpdateParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskUpdateParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskUpdateRequest(data []byte) (TaskUpdateRequest, error) {
	var r TaskUpdateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskUpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskUpdateResponse(data []byte) (TaskUpdateResponse, error) {
	var r TaskUpdateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskUpdateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskCreateRequest(data []byte) (TaskCreateRequest, error) {
	var r TaskCreateRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskCreateRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskCreateResponse(data []byte) (TaskCreateResponse, error) {
	var r TaskCreateResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskCreateResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTasksExportParams(data []byte) (TasksExportParams, error) {
	var r TasksExportParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TasksExportParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTasksExportResponse(data []byte) (TasksExportResponse, error) {
	var r TasksExportResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TasksExportResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTasksListParams(data []byte) (TasksListParams, error) {
	var r TasksListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TasksListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTasksListResponse(data []byte) (TasksListResponse, error) {
	var r TasksListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TasksListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTeamsListParams(data []byte) (TeamsListParams, error) {
	var r TeamsListParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TeamsListParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTeamsListResponse(data []byte) (TeamsListResponse, error) {
	var r TeamsListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TeamsListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUsageFetchResponse(data []byte) (UsageFetchResponse, error) {
	var r UsageFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsageFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUsageSeriesFetchResponse(data []byte) (UsageSeriesFetchResponse, error) {
	var r UsageSeriesFetchResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UsageSeriesFetchResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessage(data []byte) (Message, error) {
	var r Message
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Message) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEntity(data []byte) (Entity, error) {
	var r Entity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Entity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMessageType(data []byte) (MessageType, error) {
	var r MessageType
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MessageType) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTrigger(data []byte) (Trigger, error) {
	var r Trigger
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Trigger) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSchedule(data []byte) (Schedule, error) {
	var r Schedule
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Schedule) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSyncStatus(data []byte) (SyncStatus, error) {
	var r SyncStatus
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskStatus(data []byte) (TaskStatus, error) {
	var r TaskStatus
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTaskOutcome(data []byte) (TaskOutcome, error) {
	var r TaskOutcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TaskOutcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintVisibility(data []byte) (BlueprintVisibility, error) {
	var r BlueprintVisibility
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintVisibility) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotVisibility(data []byte) (BotVisibility, error) {
	var r BotVisibility
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotVisibility) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetVisibility(data []byte) (DatasetVisibility, error) {
	var r DatasetVisibility
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetVisibility) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDatasetFileAttachmentType(data []byte) (DatasetFileAttachmentType, error) {
	var r DatasetFileAttachmentType
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFileAttachmentType) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DatasetFilter map[string]*DatasetFilterValue

func UnmarshalDatasetFilter(data []byte) (DatasetFilter, error) {
	var r DatasetFilter
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DatasetFilter) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSkillsetVisibility(data []byte) (SkillsetVisibility, error) {
	var r SkillsetVisibility
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkillsetVisibility) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileVisibility(data []byte) (FileVisibility, error) {
	var r FileVisibility
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FileVisibility) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretType(data []byte) (SecretType, error) {
	var r SecretType
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretType) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretKind(data []byte) (SecretKind, error) {
	var r SecretKind
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretKind) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSecretVisibility(data []byte) (SecretVisibility, error) {
	var r SecretVisibility
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SecretVisibility) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUsage(data []byte) (Usage, error) {
	var r Usage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Usage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCompleteReason(data []byte) (CompleteReason, error) {
	var r CompleteReason
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompleteReason) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCompleteEnd(data []byte) (CompleteEnd, error) {
	var r CompleteEnd
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompleteEnd) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExecutionLimits(data []byte) (ExecutionLimits, error) {
	var r ExecutionLimits
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExecutionLimits) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPolicyType(data []byte) (PolicyType, error) {
	var r PolicyType
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PolicyType) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalLimits(data []byte) (Limits, error) {
	var r Limits
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Limits) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Meta map[string]interface{}

func UnmarshalMeta(data []byte) (Meta, error) {
	var r Meta
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Meta) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Model string

func UnmarshalModel(data []byte) (Model, error) {
	var r Model
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Model) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotRef(data []byte) (BotRef, error) {
	var r BotRef
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotRef) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotConfig(data []byte) (BotConfig, error) {
	var r BotConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBotRefOrConfig(data []byte) (BotRefOrConfig, error) {
	var r BotRefOrConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BotRefOrConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalBlueprintProps(data []byte) (BlueprintProps, error) {
	var r BlueprintProps
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BlueprintProps) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstanceRefProperties(data []byte) (InstanceRefProperties, error) {
	var r InstanceRefProperties
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstanceRefProperties) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstanceMetaProps(data []byte) (InstanceMetaProps, error) {
	var r InstanceMetaProps
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstanceMetaProps) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstanceCRUDProps(data []byte) (InstanceCRUDProps, error) {
	var r InstanceCRUDProps
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstanceCRUDProps) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstanceListProps(data []byte) (InstanceListProps, error) {
	var r InstanceListProps
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstanceListProps) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalJSONSchemaObject(data []byte) (JSONSchemaObject, error) {
	var r JSONSchemaObject
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JSONSchemaObject) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FunctionsDefinition []FunctionsDefinitionElement

func UnmarshalFunctionsDefinition(data []byte) (FunctionsDefinition, error) {
	var r FunctionsDefinition
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FunctionsDefinition) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExtensionsDefinition(data []byte) (ExtensionsDefinition, error) {
	var r ExtensionsDefinition
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExtensionsDefinition) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCompleteStreamingResponseItem(data []byte) (CompleteStreamingResponseItem, error) {
	var r CompleteStreamingResponseItem
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompleteStreamingResponseItem) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GraphqlRequest struct {
	// The name of the operation to execute                          
	OperationName                             *string                `json:"operationName,omitempty"`
	// The GraphQL query or mutation string                          
	Query                                     string                 `json:"query"`
	// The variables for the GraphQL operation                       
	Variables                                 map[string]interface{} `json:"variables,omitempty"`
}

type GraphqlResponse struct {
	// The data returned by the GraphQL operation                         
	Data                                           map[string]interface{} `json:"data,omitempty"`
	// Any errors returned by the GraphQL operation                       
	Errors                                         []Error                `json:"errors,omitempty"`
}

type Error struct {
	Message *string `json:"message,omitempty"`
}

type PlatformReportsListParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type PlatformReportsListResponse struct {
	Items []PlatformReportsListResponseItem `json:"items"`
}

// Instance list properties
type PlatformReportsListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type ReportGenerateParams struct {
	// The ID of the report to generate       
	ReportID                           string `json:"reportId"`
}

// Successful report output data
type ReportsGenerateResponseValue struct {
	// Error message if report generation failed        
	Error                                       *string `json:"error,omitempty"`
}

type BlueprintCloneParams struct {
	// The ID of the blueprint to clone       
	BlueprintID                        string `json:"blueprintId"`
}

type BlueprintCloneResponse struct {
	// The ID of the cloned blueprint                                
	ID                                        string                 `json:"id"`
	// A map of the resources that were cloned                       
	Resources                                 map[string]interface{} `json:"resources"`
}

type BlueprintDeleteParams struct {
	// The ID of the blueprint to delete       
	BlueprintID                         string `json:"blueprintId"`
}

type BlueprintDeleteRequest struct {
	// If true, deletes all resources associated with the blueprint. If false or omitted, only      
	// the blueprint is deleted.                                                                    
	DeleteResources                                                                           *bool `json:"deleteResources,omitempty"`
}

type BlueprintDeleteResponse struct {
	// The ID of the deleted blueprint       
	ID                                string `json:"id"`
}

type BlueprintFetchParams struct {
	// The ID of the blueprint to retrieve       
	BlueprintID                           string `json:"blueprintId"`
}

// Instance list properties
type BlueprintFetchResponse struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The blueprint visibility                                               
	Visibility                                         *SecretVisibility      `json:"visibility,omitempty"`
}

type BlueprintResourcesListParams struct {
	// The ID of the blueprint to clone       
	BlueprintID                        string `json:"blueprintId"`
}

type BlueprintResourcesListResponse struct {
	// The ID of the blueprint                       
	ID                        string                 `json:"id"`
	// A map of the resources                        
	Resources                 map[string]interface{} `json:"resources"`
}

type BlueprintUpdateParams struct {
	BlueprintID string `json:"blueprintId"`
}

// Instance crud properties
type BlueprintUpdateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The blueprint visibility                                
	Visibility                          *SecretVisibility      `json:"visibility,omitempty"`
}

type BlueprintUpdateResponse struct {
	// The ID of the updated blueprint       
	ID                                string `json:"id"`
}

// Instance crud properties
type BlueprintCreateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The blueprint visibility                                
	Visibility                          *SecretVisibility      `json:"visibility,omitempty"`
}

type BlueprintCreateResponse struct {
	// The ID of the created blueprint       
	ID                                string `json:"id"`
}

type BlueprintsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type BlueprintsListResponse struct {
	Items []BlueprintsListResponseItem `json:"items"`
}

// Instance list properties
type BlueprintsListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The blueprint visibility                                               
	Visibility                                         *SecretVisibility      `json:"visibility,omitempty"`
}

type BotCloneParams struct {
	BotID string `json:"botId"`
}

type BotCloneResponse struct {
	// The ID of the cloned bot       
	ID                         string `json:"id"`
}

type BotDeleteParams struct {
	// The ID of the bot to delete       
	BotID                         string `json:"botId"`
}

type BotDeleteResponse struct {
	// The ID of the deleted bot       
	ID                          string `json:"id"`
}

type BotDownvoteParams struct {
	// The ID of the bot       
	BotID               string `json:"botId"`
}

type BotDownvoteRequest struct {
	// The reason for the downvote        
	Reason                        *string `json:"reason,omitempty"`
	// The value of the downvote          
	Value                         *int64  `json:"value,omitempty"`
}

type BotDownvoteResponse struct {
	// The bot ID of the downvoted bot       
	ID                                string `json:"id"`
}

type BotFetchParams struct {
	// The ID of the bot to retrieve       
	BotID                           string `json:"botId"`
}

// Blueprint properties
type BotFetchResponse struct {
	// The backstory this configuration is using                                
	Backstory                                            *string                `json:"backstory,omitempty"`
	// The ID of the blueprint                                                  
	BlueprintID                                          *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                         
	CreatedAt                                            float64                `json:"createdAt"`
	// The id of the dataset this configuration is using                        
	DatasetID                                            *string                `json:"datasetId,omitempty"`
	// The associated description                                               
	Description                                          *string                `json:"description,omitempty"`
	// The instance ID                                                          
	ID                                                   string                 `json:"id"`
	// Meta data information                                                    
	Meta                                                 map[string]interface{} `json:"meta,omitempty"`
	// A model definition                                                       
	Model                                                *string                `json:"model,omitempty"`
	// The moderation flag for this configuration                               
	Moderation                                           *bool                  `json:"moderation,omitempty"`
	// The associated name                                                      
	Name                                                 *string                `json:"name,omitempty"`
	// The privacy flag for this configuration                                  
	Privacy                                              *bool                  `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                       
	SkillsetID                                           *string                `json:"skillsetId,omitempty"`
	// The timestamp (ms) when the instance was updated                         
	UpdatedAt                                            float64                `json:"updatedAt"`
	// The bot visibility                                                       
	Visibility                                           *SecretVisibility      `json:"visibility,omitempty"`
}

type BotMemorySearchParams struct {
	// The ID of the bot to search memories for       
	BotID                                      string `json:"botId"`
}

type BotMemorySearchRequest struct {
	// The keyword/phrase to search for       
	Search                             string `json:"search"`
}

type BotMemorySearchResponse struct {
	// An array of memories matching the search query                              
	Items                                            []BotMemorySearchResponseItem `json:"items"`
}

type BotMemorySearchResponseItem struct {
	ID   string                 `json:"id"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Text string                 `json:"text"`
}

type BotSessionCreateParams struct {
	// The ID of the bot for this session       
	BotID                                string `json:"botId"`
}

type BotSessionCreateRequest struct {
	// The maximum amount of time this session will stay open                                  
	DurationInSeconds                                         *float64                         `json:"durationInSeconds,omitempty"`
	// An array of messages to be included in the conversation                                 
	Messages                                                  []BotSessionCreateRequestMessage `json:"messages,omitempty"`
	// Meta data information                                                                   
	Meta                                                      map[string]interface{}           `json:"meta,omitempty"`
}

type BotSessionCreateRequestMessage struct {
	// The text of the message            
	Text                      string      `json:"text"`
	// The type of the message            
	Type                      MessageType `json:"type"`
}

type BotSessionCreateResponse struct {
	// The ID of the conversation                                                         
	ConversationID                                      string                            `json:"conversationId"`
	// The time the token will expire in milliseconds                                     
	ExpiresAt                                           float64                           `json:"expiresAt"`
	// The ID of the bot                                                                  
	ID                                                  string                            `json:"id"`
	// An array of messages included in the conversation                                  
	Messages                                            []BotSessionCreateResponseMessage `json:"messages,omitempty"`
	// The token for this conversation                                                    
	Token                                               string                            `json:"token"`
}

type BotSessionCreateResponseMessage struct {
	// The text of the message            
	Text                      string      `json:"text"`
	// The type of the message            
	Type                      MessageType `json:"type"`
}

type BotUpdateParams struct {
	BotID string `json:"botId"`
}

// Blueprint properties
type BotUpdateRequest struct {
	// The unique alias for the instance                                        
	Alias                                                *string                `json:"alias,omitempty"`
	// The backstory this configuration is using                                
	Backstory                                            *string                `json:"backstory,omitempty"`
	// The ID of the blueprint                                                  
	BlueprintID                                          *string                `json:"blueprintId,omitempty"`
	// The id of the dataset this configuration is using                        
	DatasetID                                            *string                `json:"datasetId,omitempty"`
	// The associated description                                               
	Description                                          *string                `json:"description,omitempty"`
	// Meta data information                                                    
	Meta                                                 map[string]interface{} `json:"meta,omitempty"`
	// A model definition                                                       
	Model                                                *string                `json:"model,omitempty"`
	// The moderation flag for this configuration                               
	Moderation                                           *bool                  `json:"moderation,omitempty"`
	// The associated name                                                      
	Name                                                 *string                `json:"name,omitempty"`
	// The privacy flag for this configuration                                  
	Privacy                                              *bool                  `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                       
	SkillsetID                                           *string                `json:"skillsetId,omitempty"`
	// The bot visibility                                                       
	Visibility                                           *SecretVisibility      `json:"visibility,omitempty"`
}

type BotUpdateResponse struct {
	// The ID of the updated bot       
	ID                          string `json:"id"`
}

type BotUpvoteParams struct {
	// The ID of the bot       
	BotID               string `json:"botId"`
}

type BotUpvoteRequest struct {
	// The reason for the upvote        
	Reason                      *string `json:"reason,omitempty"`
	// The value of the upvote          
	Value                       *int64  `json:"value,omitempty"`
}

type BotUpvoteResponse struct {
	// The ID of the upvoted bot       
	ID                          string `json:"id"`
}

type BotUsageFetchParams struct {
	// The ID of the bot                                     
	BotID                                         string     `json:"botId"`
	// Start date for the period (ISO 8601 format)           
	From                                          *time.Time `json:"from,omitempty"`
	// End date for the period (ISO 8601 format)             
	To                                            *time.Time `json:"to,omitempty"`
}

type BotUsageFetchResponse struct {
	// Total number of conversations          
	Conversations                      *int64 `json:"conversations,omitempty"`
	// Total number of messages               
	Messages                           *int64 `json:"messages,omitempty"`
	// Total number of BASE tokens used       
	Tokens                             *int64 `json:"tokens,omitempty"`
}

// Blueprint properties
type BotCreateRequest struct {
	// The unique alias for the instance                                        
	Alias                                                *string                `json:"alias,omitempty"`
	// The backstory this configuration is using                                
	Backstory                                            *string                `json:"backstory,omitempty"`
	// The ID of the blueprint                                                  
	BlueprintID                                          *string                `json:"blueprintId,omitempty"`
	// The id of the dataset this configuration is using                        
	DatasetID                                            *string                `json:"datasetId,omitempty"`
	// The associated description                                               
	Description                                          *string                `json:"description,omitempty"`
	// Meta data information                                                    
	Meta                                                 map[string]interface{} `json:"meta,omitempty"`
	// A model definition                                                       
	Model                                                *string                `json:"model,omitempty"`
	// The moderation flag for this configuration                               
	Moderation                                           *bool                  `json:"moderation,omitempty"`
	// The associated name                                                      
	Name                                                 *string                `json:"name,omitempty"`
	// The privacy flag for this configuration                                  
	Privacy                                              *bool                  `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                       
	SkillsetID                                           *string                `json:"skillsetId,omitempty"`
	// The bot visibility                                                       
	Visibility                                           *SecretVisibility      `json:"visibility,omitempty"`
}

type BotCreateResponse struct {
	// The ID of the created bot       
	ID                          string `json:"id"`
}

type BotsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type BotsListResponse struct {
	Items []BotsListResponseItem `json:"items"`
}

// Blueprint properties
type BotsListResponseItem struct {
	// The backstory this configuration is using                                
	Backstory                                            *string                `json:"backstory,omitempty"`
	// The ID of the blueprint                                                  
	BlueprintID                                          *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                         
	CreatedAt                                            float64                `json:"createdAt"`
	// The id of the dataset this configuration is using                        
	DatasetID                                            *string                `json:"datasetId,omitempty"`
	// The associated description                                               
	Description                                          *string                `json:"description,omitempty"`
	// The instance ID                                                          
	ID                                                   string                 `json:"id"`
	// Meta data information                                                    
	Meta                                                 map[string]interface{} `json:"meta,omitempty"`
	// A model definition                                                       
	Model                                                *string                `json:"model,omitempty"`
	// The moderation flag for this configuration                               
	Moderation                                           *bool                  `json:"moderation,omitempty"`
	// The associated name                                                      
	Name                                                 *string                `json:"name,omitempty"`
	// The privacy flag for this configuration                                  
	Privacy                                              *bool                  `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                       
	SkillsetID                                           *string                `json:"skillsetId,omitempty"`
	// The timestamp (ms) when the instance was updated                         
	UpdatedAt                                            float64                `json:"updatedAt"`
	// The bot visibility                                                       
	Visibility                                           *SecretVisibility      `json:"visibility,omitempty"`
}

type ChannelMessagePublishParams struct {
	// The ID of the channel to publish to (minimum 16 characters)       
	ChannelID                                                     string `json:"channelId"`
}

type ChannelMessagePublishRequest struct {
	// The message to publish to the channel                       
	Message                                 map[string]interface{} `json:"message"`
}

type ChannelMessagePublishResponse struct {
	// The ID of the channel the message was published to       
	ID                                                   string `json:"id"`
}

type ChannelMessagesSubscribeParams struct {
	// The ID of the channel to subscribe to (minimum 16 characters)       
	ChannelID                                                       string `json:"channelId"`
}

type ChannelMessagesSubscribeRequest struct {
	// Number of historical messages to replay from the channel       
	// before subscribing to live updates. When provided, the         
	// subscriber will first receive up to this many recent           
	// messages that were published before the subscription           
	// started. This is useful for catching up on messages that       
	// may have been published during connection setup.               
	HistoryLength                                              *int64 `json:"historyLength,omitempty"`
}

type ContactConversationsListParams struct {
	// The ID of the contact to list conversations for        
	ContactID                                         string  `json:"contactId"`
	// The cursor to use for pagination                       
	Cursor                                            *string `json:"cursor,omitempty"`
	// The order of the paginated items                       
	Order                                             *Order  `json:"order,omitempty"`
	// The number of items to retrieve                        
	Take                                              *int64  `json:"take,omitempty"`
}

type ContactConversationsListResponse struct {
	Items []ContactConversationsListResponseItem `json:"items"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ContactConversationsListResponseItem struct {
	// The contact id assigned to this conversation                             
	ContactID                                            *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                         
	CreatedAt                                            float64                `json:"createdAt"`
	// The associated description                                               
	Description                                          *string                `json:"description,omitempty"`
	// The instance ID                                                          
	ID                                                   string                 `json:"id"`
	// Meta data information                                                    
	Meta                                                 map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                      
	Name                                                 *string                `json:"name,omitempty"`
	// The task id assigned to this conversation                                
	TaskID                                               *string                `json:"taskId,omitempty"`
	// The timestamp (ms) when the instance was updated                         
	UpdatedAt                                            float64                `json:"updatedAt"`
	// The ID of the bot this configuration is using                            
	BotID                                                *string                `json:"botId,omitempty"`
	// The backstory this configuration is using                                
	Backstory                                            *string                `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using                        
	DatasetID                                            *string                `json:"datasetId,omitempty"`
	// A model definition                                                       
	Model                                                *string                `json:"model,omitempty"`
	// The moderation flag for this configuration                               
	Moderation                                           *bool                  `json:"moderation,omitempty"`
	// The privacy flag for this configuration                                  
	Privacy                                              *bool                  `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                       
	SkillsetID                                           *string                `json:"skillsetId,omitempty"`
}

type ContactDeleteParams struct {
	// The ID of the contact to delete       
	ContactID                         string `json:"contactId"`
}

type ContactDeleteResponse struct {
	// The ID of the deleted contact       
	ID                              string `json:"id"`
}

type ContactFetchParams struct {
	// The ID of the contact to retrieve       
	ContactID                           string `json:"contactId"`
}

// Instance list properties
type ContactFetchResponse struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The email address of the contact                                       
	Email                                              *string                `json:"email,omitempty"`
	// The fingerprint of the contact                                         
	Fingerprint                                        string                 `json:"fingerprint"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The nickname of the contact                                            
	Nick                                               *string                `json:"nick,omitempty"`
	// The phone number of the contact                                        
	Phone                                              *string                `json:"phone,omitempty"`
	// The preferences of the contact                                         
	Preferences                                        *string                `json:"preferences,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The timestamp (ms) when the contact was verified                       
	VerifiedAt                                         *float64               `json:"verifiedAt,omitempty"`
}

type ContactMemoriesListParams struct {
	// The ID of the contact to list memories for        
	ContactID                                    string  `json:"contactId"`
	// The cursor to use for pagination                  
	Cursor                                       *string `json:"cursor,omitempty"`
	// The order of the paginated items                  
	Order                                        *Order  `json:"order,omitempty"`
	// The number of items to retrieve                   
	Take                                         *int64  `json:"take,omitempty"`
}

type ContactMemoriesListResponse struct {
	Items []ContactMemoriesListResponseItem `json:"items"`
}

// Instance list properties
type ContactMemoriesListResponseItem struct {
	// The ID of the bot the memory belongs to                                
	BotID                                              *string                `json:"botId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The text of the memory                                                 
	Text                                               string                 `json:"text"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type ContactMemorySearchParams struct {
	// The ID of the contact to search memories for       
	ContactID                                      string `json:"contactId"`
}

type ContactMemorySearchRequest struct {
	// The keyword/phrase to search for       
	Search                             string `json:"search"`
}

type ContactMemorySearchResponse struct {
	// An array of memories matching the search query                                  
	Items                                            []ContactMemorySearchResponseItem `json:"items"`
}

type ContactMemorySearchResponseItem struct {
	ID   string                 `json:"id"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Text string                 `json:"text"`
}

type ContactSecretAuthenticateParams struct {
	// The ID of the contact the secret belongs to       
	ContactID                                     string `json:"contactId"`
	// The ID of the secret to authenticate              
	SecretID                                      string `json:"secretId"`
}

type ContactSecretAuthenticateResponse struct {
	// The ID of the secret to authenticate       
	ID                                     string `json:"id"`
	// The URL to authenticate the secret         
	URL                                    string `json:"url"`
}

type ContactSecretRevokeParams struct {
	// The ID of the contact the secret belongs to       
	ContactID                                     string `json:"contactId"`
	// The ID of the secret to be revoked                
	SecretID                                      string `json:"secretId"`
}

type ContactSecretRevokeResponse struct {
	// The ID of the revoked secret       
	ID                             string `json:"id"`
}

type ContactSecretVerifyParams struct {
	// The ID of the contact the secret belongs to       
	ContactID                                     string `json:"contactId"`
	// The ID of the secret to be verified               
	SecretID                                      string `json:"secretId"`
}

type ContactSecretVerifyResponse struct {
	Action                          *ContactSecretVerifyResponseAction `json:"action,omitempty"`
	// The ID of the verified secret                                   
	ID                              string                             `json:"id"`
	// The status of the secret                                        
	Status                          Status                             `json:"status"`
}

// The action to take next
type ContactSecretVerifyResponseAction struct {
	// The type of action to take                   
	Type                                 ActionType `json:"type"`
	// The URL to authenticate the secret           
	URL                                  string     `json:"url"`
}

type ContactSecretsListParams struct {
	// The ID of the contact to list secrets for        
	ContactID                                   string  `json:"contactId"`
	// The cursor to use for pagination                 
	Cursor                                      *string `json:"cursor,omitempty"`
	// The order of the paginated items                 
	Order                                       *Order  `json:"order,omitempty"`
	// The number of items to retrieve                  
	Take                                        *int64  `json:"take,omitempty"`
}

type ContactSecretsListResponse struct {
	Items []ContactSecretsListResponseItem `json:"items"`
}

// Instance list properties
type ContactSecretsListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The type of the secret                                                 
	Type                                               string                 `json:"type"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type ContactSpacesListParams struct {
	// The ID of the contact to list spaces for        
	ContactID                                  string  `json:"contactId"`
	// The cursor to use for pagination                
	Cursor                                     *string `json:"cursor,omitempty"`
	// The order of the paginated items                
	Order                                      *Order  `json:"order,omitempty"`
	// The number of items to retrieve                 
	Take                                       *int64  `json:"take,omitempty"`
}

type ContactSpacesListResponse struct {
	Items []ContactSpacesListResponseItem `json:"items"`
}

// Instance list properties
type ContactSpacesListResponseItem struct {
	// The contact id assigned to this rating                                 
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type ContactTasksListParams struct {
	// The ID of the contact to list tasks for        
	ContactID                                 string  `json:"contactId"`
	// The cursor to use for pagination               
	Cursor                                    *string `json:"cursor,omitempty"`
	// The order of the paginated items               
	Order                                     *Order  `json:"order,omitempty"`
	// The number of items to retrieve                
	Take                                      *int64  `json:"take,omitempty"`
}

type ContactTasksListResponse struct {
	Items []ContactTasksListResponseItem `json:"items"`
}

// Instance list properties
type ContactTasksListResponseItem struct {
	// The bot associated with the task                                       
	BotID                                              *string                `json:"botId,omitempty"`
	// The contact id assigned to this task                                   
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The task execution outcome                                             
	Outcome                                            *TaskOutcome           `json:"outcome,omitempty"`
	// The schedule of the task                                               
	Schedule                                           *string                `json:"schedule,omitempty"`
	// The task execution status                                              
	Status                                             *TaskStatus            `json:"status,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type ContactUpdateParams struct {
	ContactID string `json:"contactId"`
}

// Instance crud properties
type ContactUpdateRequest struct {
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The email address of the contact                                       
	Email                                              *string                `json:"email,omitempty"`
	// The fingerprint of the contact                                         
	Fingerprint                                        *string                `json:"fingerprint,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The nickname of the contact                                            
	Nick                                               *string                `json:"nick,omitempty"`
	// The phone number of the contact                                        
	Phone                                              *string                `json:"phone,omitempty"`
	// The preferences of the contact                                         
	Preferences                                        *string                `json:"preferences,omitempty"`
	// The timestamp (ms) when the contact was verified                       
	VerifiedAt                                         *float64               `json:"verifiedAt,omitempty"`
}

type ContactUpdateResponse struct {
	// The ID of the updated contact       
	ID                              string `json:"id"`
}

// Instance crud properties
type ContactCreateRequest struct {
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The email address of the contact                                       
	Email                                              *string                `json:"email,omitempty"`
	// The fingerprint of the contact                                         
	Fingerprint                                        *string                `json:"fingerprint,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The nickname of the contact                                            
	Nick                                               *string                `json:"nick,omitempty"`
	// The phone number of the contact                                        
	Phone                                              *string                `json:"phone,omitempty"`
	// The preferences of the contact                                         
	Preferences                                        *string                `json:"preferences,omitempty"`
	// The timestamp (ms) when the contact was verified                       
	VerifiedAt                                         *float64               `json:"verifiedAt,omitempty"`
}

type ContactCreateResponse struct {
	// The ID of the created contact       
	ID                              string `json:"id"`
}

// Instance crud properties
type ContactEnsureRequest struct {
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The email address of the contact                                       
	Email                                              *string                `json:"email,omitempty"`
	// The fingerprint of the contact                                         
	Fingerprint                                        string                 `json:"fingerprint"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The nickname of the contact                                            
	Nick                                               *string                `json:"nick,omitempty"`
	// The phone number of the contact                                        
	Phone                                              *string                `json:"phone,omitempty"`
	// The preferences of the contact                                         
	Preferences                                        *string                `json:"preferences,omitempty"`
	// The timestamp (ms) when the contact was verified                       
	VerifiedAt                                         *float64               `json:"verifiedAt,omitempty"`
}

type ContactEnsureResponse struct {
	// The ID of the ensured contact       
	ID                              string `json:"id"`
}

type ContactsExportParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ContactsExportResponse struct {
	Items []ContactsExportResponseItem `json:"items"`
}

// Instance list properties
type ContactsExportResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The email address of the contact                                       
	Email                                              *string                `json:"email,omitempty"`
	// The fingerprint of the contact                                         
	Fingerprint                                        string                 `json:"fingerprint"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The nickname of the contact                                            
	Nick                                               *string                `json:"nick,omitempty"`
	// The phone number of the contact                                        
	Phone                                              *string                `json:"phone,omitempty"`
	// The preferences of the contact                                         
	Preferences                                        *string                `json:"preferences,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The timestamp (ms) when the contact was verified                       
	VerifiedAt                                         *float64               `json:"verifiedAt,omitempty"`
}

type ContactsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ContactsListResponse struct {
	Items []ContactsListResponseItem `json:"items"`
}

// Instance list properties
type ContactsListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The email address of the contact                                       
	Email                                              *string                `json:"email,omitempty"`
	// The fingerprint of the contact                                         
	Fingerprint                                        string                 `json:"fingerprint"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The nickname of the contact                                            
	Nick                                               *string                `json:"nick,omitempty"`
	// The phone number of the contact                                        
	Phone                                              *string                `json:"phone,omitempty"`
	// The preferences of the contact                                         
	Preferences                                        *string                `json:"preferences,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The timestamp (ms) when the contact was verified                       
	VerifiedAt                                         *float64               `json:"verifiedAt,omitempty"`
}

type ConversationAttachmentUploadParams struct {
	ConversationID string `json:"conversationId"`
}

type ConversationAttachmentUploadRequest struct {
	// The file to upload either as http: or data: URL                                         
	//                                                                                         
	// The file definition to upload                                                           
	File                                              *ConversationAttachmentUploadRequestFile `json:"file"`
}

// The file definition to upload
type PurpleFile struct {
	// The file name        
	Name            *string `json:"name,omitempty"`
	// The file size        
	Size            float64 `json:"size"`
	// The file type        
	Type            string  `json:"type"`
}

type ConversationAttachmentUploadResponse struct {
	// The ID of the upload file                                                                 
	ID                                        string                                             `json:"id"`
	// The name of the uploaded file                                                             
	Name                                      *string                                            `json:"name,omitempty"`
	// The request required to upload the file                                                   
	UploadRequest                             *ConversationAttachmentUploadResponseUploadRequest `json:"uploadRequest,omitempty"`
}

// The request required to upload the file
type ConversationAttachmentUploadResponseUploadRequest struct {
	// The HTTP headers to use                       
	Headers                   map[string]interface{} `json:"headers"`
	// The HTTP method to use                        
	Method                    string                 `json:"method"`
	// The HTTP url to use                           
	URL                       string                 `json:"url"`
}

type ConversationMessageCompleteParams struct {
	// The ID of the conversation to receive message from       
	ConversationID                                       string `json:"conversationId"`
}

type ConversationMessageCompleteRequest struct {
	// Known entities                                                                                          
	Entities                                                     []ConversationMessageCompleteRequestEntity    `json:"entities,omitempty"`
	// Extensions to enhance the bot's capabilities                                                            
	Extensions                                                   *ConversationMessageCompleteRequestExtensions `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                                   
	Functions                                                    []ConversationMessageCompleteRequestFunction  `json:"functions,omitempty"`
	// Execution limits to control conversation processing bounds                                              
	Limits                                                       *ConversationMessageCompleteRequestLimits     `json:"limits,omitempty"`
	// The text of the message to send                                                                         
	Text                                                         string                                        `json:"text"`
}

// Extracted entity from the message
type ConversationMessageCompleteRequestEntity struct {
	// Start offset                                   
	Begin                          float64            `json:"begin"`
	// End offset                                     
	End                            float64            `json:"end"`
	Replacement                    *PurpleReplacement `json:"replacement,omitempty"`
	// The text value of the entity                   
	Text                           string             `json:"text"`
	// The entity type                                
	Type                           string             `json:"type"`
}

type PurpleReplacement struct {
	// Start offset                             
	Begin                               float64 `json:"begin"`
	// End offset                               
	End                                 float64 `json:"end"`
	// The text value of the replacement        
	Text                                string  `json:"text"`
}

// Extensions to enhance the bot's capabilities
type ConversationMessageCompleteRequestExtensions struct {
	// Additional backstory for the bot                                  
	Backstory                                           *string          `json:"backstory,omitempty"`
	// Inline datasets to provide additional context                     
	Datasets                                            []PurpleDataset  `json:"datasets,omitempty"`
	// Feature flags to enable specific bot capabilities                 
	Features                                            []PurpleFeature  `json:"features,omitempty"`
	// Inline skillsets to provide additional abilities                  
	Skillsets                                           []PurpleSkillset `json:"skillsets,omitempty"`
}

type PurpleDataset struct {
	// The description of the dataset               
	Description                      *string        `json:"description,omitempty"`
	// The name of the dataset                      
	Name                             *string        `json:"name,omitempty"`
	// The records in the dataset                   
	Records                          []PurpleRecord `json:"records"`
}

type PurpleRecord struct {
	// Additional metadata for the record                       
	Meta                                 map[string]interface{} `json:"meta,omitempty"`
	// The text content of the record                           
	Text                                 string                 `json:"text"`
}

type PurpleFeature struct {
	// The name of the feature to enable                                    
	Name                                             string                 `json:"name"`
	// Optional configuration options for the feature                       
	Options                                          map[string]interface{} `json:"options,omitempty"`
}

type PurpleSkillset struct {
	// The abilities in the skillset                  
	Abilities                         []PurpleAbility `json:"abilities"`
	// The description of the skillset                
	Description                       *string         `json:"description,omitempty"`
	// The name of the skillset                       
	Name                              *string         `json:"name,omitempty"`
}

type PurpleAbility struct {
	// The description of the ability                            
	Description                           string                 `json:"description"`
	// The instruction for the ability                           
	Instruction                           string                 `json:"instruction"`
	// Additional metadata for the ability                       
	Meta                                  map[string]interface{} `json:"meta,omitempty"`
	// The name of the ability                                   
	Name                                  string                 `json:"name"`
	// Optional secret ID for the ability                        
	SecretID                              *string                `json:"secretId,omitempty"`
}

type ConversationMessageCompleteRequestFunction struct {
	// Configuration for when this function should be automatically called                    
	Call                                                                     *PurpleCall      `json:"call,omitempty"`
	// The description of the function                                                        
	Description                                                              string           `json:"description"`
	// The name of the function (must be a valid JS identifier, max 64 chars)                 
	Name                                                                     string           `json:"name"`
	// JSON Schema definition for the function parameters                                     
	Parameters                                                               PurpleParameters `json:"parameters"`
	// The result of the function execution                                                   
	Result                                                                   *PurpleResult    `json:"result,omitempty"`
}

// Configuration for when this function should be automatically called
type PurpleCall struct {
	// If true, this function will be force-called at the end of the conversation        
	End                                                                            *bool `json:"end,omitempty"`
	// If true, this function will be force-called at the start of the conversation      
	Start                                                                          *bool `json:"start,omitempty"`
}

// JSON Schema definition for the function parameters
type PurpleParameters struct {
	// Object property definitions                             
	Properties                          map[string]interface{} `json:"properties"`
	// Required property names                                 
	Required                            []string               `json:"required,omitempty"`
	// The schema type, must be "object"                       
	Type                                ParametersType         `json:"type"`
}

// The result of the function execution
type PurpleResult struct {
	// The data returned by the function (can be any type)            
	Data                                                  interface{} `json:"data"`
	// The channel for streaming function results                     
	Channel                                               *string     `json:"channel,omitempty"`
}

// Execution limits to control conversation processing bounds
type ConversationMessageCompleteRequestLimits struct {
	// Maximum number of function/tool calls. Controls how many total function calls can be made       
	// during the conversation.                                                                        
	Calls                                                                                       *int64 `json:"calls,omitempty"`
	// Maximum number of model continuations. Controls how many times the model can continue           
	// generating after reaching a stop condition.                                                     
	Continuations                                                                               *int64 `json:"continuations,omitempty"`
	// Maximum number of agentic iterations. Controls how many times the model can iterate             
	// through tool calls and responses.                                                               
	Iterations                                                                                  *int64 `json:"iterations,omitempty"`
}

type ConversationMessageCompleteResponse struct {
	// Information about why the completion ended                                         
	End                                          ConversationMessageCompleteResponseEnd   `json:"end"`
	// The ID of the created message                                                      
	ID                                           string                                   `json:"id"`
	// The text of the message received                                                   
	Text                                         string                                   `json:"text"`
	// Usage information                                                                  
	Usage                                        ConversationMessageCompleteResponseUsage `json:"usage"`
}

// Information about why the completion ended
type ConversationMessageCompleteResponseEnd struct {
	// The reason why the completion ended               
	Reason                                CompleteReason `json:"reason"`
}

// Usage information
type ConversationMessageCompleteResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

type ConversationContactUpsertParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

// Instance crud properties
type ConversationContactUpsertRequest struct {
	// The associated description                             
	Description                        *string                `json:"description,omitempty"`
	// The email address of the contact                       
	Email                              *string                `json:"email,omitempty"`
	// The fingerprint of the contact                         
	Fingerprint                        *string                `json:"fingerprint,omitempty"`
	// Meta data information                                  
	Meta                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                    
	Name                               *string                `json:"name,omitempty"`
	// The nickname of the contact                            
	Nick                               *string                `json:"nick,omitempty"`
	// The phone number of the contact                        
	Phone                              *string                `json:"phone,omitempty"`
}

type ConversationContactUpsertResponse struct {
	// The ID of the created contact       
	ID                              string `json:"id"`
}

type ConversationDeleteParams struct {
	// The ID of the conversation to delete       
	ConversationID                         string `json:"conversationId"`
}

type ConversationDeleteResponse struct {
	// The ID of the deleted conversation       
	ID                                   string `json:"id"`
}

type StatefulConversationDispatchRequest struct {
	// A unique ID to deduplicate dispatch requests                                                             
	ChannelID                                                    *string                                        `json:"channelId,omitempty"`
	// Known entities                                                                                           
	Entities                                                     []StatefulConversationDispatchRequestEntity    `json:"entities,omitempty"`
	// Extensions to enhance the bot's capabilities                                                             
	Extensions                                                   *StatefulConversationDispatchRequestExtensions `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                                    
	Functions                                                    []StatefulConversationDispatchRequestFunction  `json:"functions,omitempty"`
	// Execution limits to control conversation processing bounds                                               
	Limits                                                       *StatefulConversationDispatchRequestLimits     `json:"limits,omitempty"`
	// The text of the message to send                                                                          
	Text                                                         string                                         `json:"text"`
}

// Extracted entity from the message
type StatefulConversationDispatchRequestEntity struct {
	// Start offset                                   
	Begin                          float64            `json:"begin"`
	// End offset                                     
	End                            float64            `json:"end"`
	Replacement                    *FluffyReplacement `json:"replacement,omitempty"`
	// The text value of the entity                   
	Text                           string             `json:"text"`
	// The entity type                                
	Type                           string             `json:"type"`
}

type FluffyReplacement struct {
	// Start offset                             
	Begin                               float64 `json:"begin"`
	// End offset                               
	End                                 float64 `json:"end"`
	// The text value of the replacement        
	Text                                string  `json:"text"`
}

// Extensions to enhance the bot's capabilities
type StatefulConversationDispatchRequestExtensions struct {
	// Additional backstory for the bot                                  
	Backstory                                           *string          `json:"backstory,omitempty"`
	// Inline datasets to provide additional context                     
	Datasets                                            []FluffyDataset  `json:"datasets,omitempty"`
	// Feature flags to enable specific bot capabilities                 
	Features                                            []FluffyFeature  `json:"features,omitempty"`
	// Inline skillsets to provide additional abilities                  
	Skillsets                                           []FluffySkillset `json:"skillsets,omitempty"`
}

type FluffyDataset struct {
	// The description of the dataset               
	Description                      *string        `json:"description,omitempty"`
	// The name of the dataset                      
	Name                             *string        `json:"name,omitempty"`
	// The records in the dataset                   
	Records                          []FluffyRecord `json:"records"`
}

type FluffyRecord struct {
	// Additional metadata for the record                       
	Meta                                 map[string]interface{} `json:"meta,omitempty"`
	// The text content of the record                           
	Text                                 string                 `json:"text"`
}

type FluffyFeature struct {
	// The name of the feature to enable                                    
	Name                                             string                 `json:"name"`
	// Optional configuration options for the feature                       
	Options                                          map[string]interface{} `json:"options,omitempty"`
}

type FluffySkillset struct {
	// The abilities in the skillset                  
	Abilities                         []FluffyAbility `json:"abilities"`
	// The description of the skillset                
	Description                       *string         `json:"description,omitempty"`
	// The name of the skillset                       
	Name                              *string         `json:"name,omitempty"`
}

type FluffyAbility struct {
	// The description of the ability                            
	Description                           string                 `json:"description"`
	// The instruction for the ability                           
	Instruction                           string                 `json:"instruction"`
	// Additional metadata for the ability                       
	Meta                                  map[string]interface{} `json:"meta,omitempty"`
	// The name of the ability                                   
	Name                                  string                 `json:"name"`
	// Optional secret ID for the ability                        
	SecretID                              *string                `json:"secretId,omitempty"`
}

type StatefulConversationDispatchRequestFunction struct {
	// Configuration for when this function should be automatically called                    
	Call                                                                     *FluffyCall      `json:"call,omitempty"`
	// The description of the function                                                        
	Description                                                              string           `json:"description"`
	// The name of the function (must be a valid JS identifier, max 64 chars)                 
	Name                                                                     string           `json:"name"`
	// JSON Schema definition for the function parameters                                     
	Parameters                                                               FluffyParameters `json:"parameters"`
	// The result of the function execution                                                   
	Result                                                                   *FluffyResult    `json:"result,omitempty"`
}

// Configuration for when this function should be automatically called
type FluffyCall struct {
	// If true, this function will be force-called at the end of the conversation        
	End                                                                            *bool `json:"end,omitempty"`
	// If true, this function will be force-called at the start of the conversation      
	Start                                                                          *bool `json:"start,omitempty"`
}

// JSON Schema definition for the function parameters
type FluffyParameters struct {
	// Object property definitions                             
	Properties                          map[string]interface{} `json:"properties"`
	// Required property names                                 
	Required                            []string               `json:"required,omitempty"`
	// The schema type, must be "object"                       
	Type                                ParametersType         `json:"type"`
}

// The result of the function execution
type FluffyResult struct {
	// The data returned by the function (can be any type)            
	Data                                                  interface{} `json:"data"`
	// The channel for streaming function results                     
	Channel                                               *string     `json:"channel,omitempty"`
}

// Execution limits to control conversation processing bounds
type StatefulConversationDispatchRequestLimits struct {
	// Maximum number of function/tool calls. Controls how many total function calls can be made       
	// during the conversation.                                                                        
	Calls                                                                                       *int64 `json:"calls,omitempty"`
	// Maximum number of model continuations. Controls how many times the model can continue           
	// generating after reaching a stop condition.                                                     
	Continuations                                                                               *int64 `json:"continuations,omitempty"`
	// Maximum number of agentic iterations. Controls how many times the model can iterate             
	// through tool calls and responses.                                                               
	Iterations                                                                                  *int64 `json:"iterations,omitempty"`
}

type StatefulConversationDispatchResponse struct {
	// The channel ID to subscribe to for completion events       
	ChannelID                                              string `json:"channelId"`
}

type ConversationDownvoteParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

type ConversationDownvoteRequest struct {
	// The reason for the downvote        
	Reason                        *string `json:"reason,omitempty"`
	// The value of the downvote          
	Value                         *int64  `json:"value,omitempty"`
}

type ConversationDownvoteResponse struct {
	// The conversation ID of the downvoted conversation       
	ID                                                  string `json:"id"`
}

type ConversationFetchParams struct {
	// The ID of the conversation to retrieve       
	ConversationID                           string `json:"conversationId"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ConversationFetchResponse struct {
	// The contact id assigned to this conversation                             
	ContactID                                            *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                         
	CreatedAt                                            float64                `json:"createdAt"`
	// The associated description                                               
	Description                                          *string                `json:"description,omitempty"`
	// The instance ID                                                          
	ID                                                   string                 `json:"id"`
	// Meta data information                                                    
	Meta                                                 map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                      
	Name                                                 *string                `json:"name,omitempty"`
	// The task id assigned to this conversation                                
	TaskID                                               *string                `json:"taskId,omitempty"`
	// The timestamp (ms) when the instance was updated                         
	UpdatedAt                                            float64                `json:"updatedAt"`
	// The ID of the bot this configuration is using                            
	BotID                                                *string                `json:"botId,omitempty"`
	// The backstory this configuration is using                                
	Backstory                                            *string                `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using                        
	DatasetID                                            *string                `json:"datasetId,omitempty"`
	// A model definition                                                       
	Model                                                *string                `json:"model,omitempty"`
	// The moderation flag for this configuration                               
	Moderation                                           *bool                  `json:"moderation,omitempty"`
	// The privacy flag for this configuration                                  
	Privacy                                              *bool                  `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                       
	SkillsetID                                           *string                `json:"skillsetId,omitempty"`
}

type ConversationMessageDeleteParams struct {
	// The ID of the conversation containing the message       
	ConversationID                                      string `json:"conversationId"`
	// The ID of the message to delete                         
	MessageID                                           string `json:"messageId"`
}

type ConversationMessageDeleteResponse struct {
	// The ID of the deleted message       
	ID                              string `json:"id"`
}

type ConversationMessageDownvoteParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
	// The ID of the message            
	MessageID                    string `json:"messageId"`
}

type ConversationMessageDownvoteRequest struct {
	// The reason for the downvote        
	Reason                        *string `json:"reason,omitempty"`
	// The value of the downvote          
	Value                         *int64  `json:"value,omitempty"`
}

type ConversationMessageDownvoteResponse struct {
	// The ID of the downvoted message       
	ID                                string `json:"id"`
}

type ConversationMessageFetchParams struct {
	// The ID of the conversation containing the message       
	ConversationID                                      string `json:"conversationId"`
	// The ID of the message to retrieve                       
	MessageID                                           string `json:"messageId"`
}

// Instance list properties
type ConversationMessageFetchResponse struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The text of the fetched message                                        
	Text                                               string                 `json:"text"`
	// The type of the message                                                
	Type                                               MessageType            `json:"type"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type ConversationMessageSynthesizeParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
	// The ID of the message            
	MessageID                    string `json:"messageId"`
}

type ConversationMessageSynthesizeResponse struct {
	// The ID of the synthesized message       
	ID                                  string `json:"id"`
}

type ConversationMessageUpdateParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
	// The ID of the message            
	MessageID                    string `json:"messageId"`
}

// Instance crud properties
type ConversationMessageUpdateRequest struct {
	// The associated description                                              
	Description                       *string                                  `json:"description,omitempty"`
	// Known entities                                                          
	Entities                          []ConversationMessageUpdateRequestEntity `json:"entities,omitempty"`
	// Meta data information                                                   
	Meta                              map[string]interface{}                   `json:"meta,omitempty"`
	// The associated name                                                     
	Name                              *string                                  `json:"name,omitempty"`
	// The updated text of the message                                         
	Text                              *string                                  `json:"text,omitempty"`
	// The type of the message                                                 
	Type                              *MessageType                             `json:"type,omitempty"`
}

// Extracted entity from the message
type ConversationMessageUpdateRequestEntity struct {
	// Start offset                                      
	Begin                          float64               `json:"begin"`
	// End offset                                        
	End                            float64               `json:"end"`
	Replacement                    *TentacledReplacement `json:"replacement,omitempty"`
	// The text value of the entity                      
	Text                           string                `json:"text"`
	// The entity type                                   
	Type                           string                `json:"type"`
}

type TentacledReplacement struct {
	// Start offset                             
	Begin                               float64 `json:"begin"`
	// End offset                               
	End                                 float64 `json:"end"`
	// The text value of the replacement        
	Text                                string  `json:"text"`
}

type ConversationMessageUpdateResponse struct {
	// The ID of the updated message       
	ID                              string `json:"id"`
}

type ConversationMessageUpvoteParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
	// The ID of the message            
	MessageID                    string `json:"messageId"`
}

type ConversationMessageUpvoteRequest struct {
	// The reason for the upvote        
	Reason                      *string `json:"reason,omitempty"`
	// The value of the upvote          
	Value                       *int64  `json:"value,omitempty"`
}

type ConversationMessageUpvoteResponse struct {
	// The ID of the upvoted message       
	ID                              string `json:"id"`
}

type ConversationMessageCreateParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

// Instance crud properties
type ConversationMessageCreateRequest struct {
	// The associated description                                         
	Description                  *string                                  `json:"description,omitempty"`
	// Known entities                                                     
	Entities                     []ConversationMessageCreateRequestEntity `json:"entities,omitempty"`
	// Meta data information                                              
	Meta                         map[string]interface{}                   `json:"meta,omitempty"`
	// The associated name                                                
	Name                         *string                                  `json:"name,omitempty"`
	// The text of the message                                            
	Text                         string                                   `json:"text"`
	// The type of the message                                            
	Type                         MessageType                              `json:"type"`
}

// Extracted entity from the message
type ConversationMessageCreateRequestEntity struct {
	// Start offset                                   
	Begin                          float64            `json:"begin"`
	// End offset                                     
	End                            float64            `json:"end"`
	Replacement                    *StickyReplacement `json:"replacement,omitempty"`
	// The text value of the entity                   
	Text                           string             `json:"text"`
	// The entity type                                
	Type                           string             `json:"type"`
}

type StickyReplacement struct {
	// Start offset                             
	Begin                               float64 `json:"begin"`
	// End offset                               
	End                                 float64 `json:"end"`
	// The text value of the replacement        
	Text                                string  `json:"text"`
}

type ConversationMessageCreateResponse struct {
	// Extracted entities from the message                                          
	Entities                              []ConversationMessageCreateResponseEntity `json:"entities"`
	// The ID of the created message                                                
	ID                                    string                                    `json:"id"`
}

// Extracted entity from the message
type ConversationMessageCreateResponseEntity struct {
	// Start offset                                   
	Begin                          float64            `json:"begin"`
	// End offset                                     
	End                            float64            `json:"end"`
	Replacement                    *IndigoReplacement `json:"replacement,omitempty"`
	// The text value of the entity                   
	Text                           string             `json:"text"`
	// The entity type                                
	Type                           string             `json:"type"`
}

type IndigoReplacement struct {
	// Start offset                             
	Begin                               float64 `json:"begin"`
	// End offset                               
	End                                 float64 `json:"end"`
	// The text value of the replacement        
	Text                                string  `json:"text"`
}

type ConversationMessagesListParams struct {
	// The ID of the conversation to list messages for        
	ConversationID                                    string  `json:"conversationId"`
	// The cursor to use for pagination                       
	Cursor                                            *string `json:"cursor,omitempty"`
	// The order of the paginated items                       
	Order                                             *Order  `json:"order,omitempty"`
	// The number of items to retrieve                        
	Take                                              *int64  `json:"take,omitempty"`
}

type ConversationMessagesListResponse struct {
	Items []ConversationMessagesListResponseItem `json:"items"`
}

// Instance list properties
type ConversationMessagesListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The text of the message                                                
	Text                                               string                 `json:"text"`
	// The type of the message                                                
	Type                                               MessageType            `json:"type"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type ConversationMessageReceiveParams struct {
	// The ID of the conversation to receive message from       
	ConversationID                                       string `json:"conversationId"`
}

type ConversationMessageReceiveRequest struct {
	// Extensions to enhance the bot's capabilities                                                      
	Extensions                                              *ConversationMessageReceiveRequestExtensions `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                             
	Functions                                               []ConversationMessageReceiveRequestFunction  `json:"functions,omitempty"`
}

// Extensions to enhance the bot's capabilities
type ConversationMessageReceiveRequestExtensions struct {
	// Additional backstory for the bot                                     
	Backstory                                           *string             `json:"backstory,omitempty"`
	// Inline datasets to provide additional context                        
	Datasets                                            []TentacledDataset  `json:"datasets,omitempty"`
	// Feature flags to enable specific bot capabilities                    
	Features                                            []TentacledFeature  `json:"features,omitempty"`
	// Inline skillsets to provide additional abilities                     
	Skillsets                                           []TentacledSkillset `json:"skillsets,omitempty"`
}

type TentacledDataset struct {
	// The description of the dataset                  
	Description                      *string           `json:"description,omitempty"`
	// The name of the dataset                         
	Name                             *string           `json:"name,omitempty"`
	// The records in the dataset                      
	Records                          []TentacledRecord `json:"records"`
}

type TentacledRecord struct {
	// Additional metadata for the record                       
	Meta                                 map[string]interface{} `json:"meta,omitempty"`
	// The text content of the record                           
	Text                                 string                 `json:"text"`
}

type TentacledFeature struct {
	// The name of the feature to enable                                    
	Name                                             string                 `json:"name"`
	// Optional configuration options for the feature                       
	Options                                          map[string]interface{} `json:"options,omitempty"`
}

type TentacledSkillset struct {
	// The abilities in the skillset                     
	Abilities                         []TentacledAbility `json:"abilities"`
	// The description of the skillset                   
	Description                       *string            `json:"description,omitempty"`
	// The name of the skillset                          
	Name                              *string            `json:"name,omitempty"`
}

type TentacledAbility struct {
	// The description of the ability                            
	Description                           string                 `json:"description"`
	// The instruction for the ability                           
	Instruction                           string                 `json:"instruction"`
	// Additional metadata for the ability                       
	Meta                                  map[string]interface{} `json:"meta,omitempty"`
	// The name of the ability                                   
	Name                                  string                 `json:"name"`
	// Optional secret ID for the ability                        
	SecretID                              *string                `json:"secretId,omitempty"`
}

type ConversationMessageReceiveRequestFunction struct {
	// Configuration for when this function should be automatically called                       
	Call                                                                     *TentacledCall      `json:"call,omitempty"`
	// The description of the function                                                           
	Description                                                              string              `json:"description"`
	// The name of the function (must be a valid JS identifier, max 64 chars)                    
	Name                                                                     string              `json:"name"`
	// JSON Schema definition for the function parameters                                        
	Parameters                                                               TentacledParameters `json:"parameters"`
	// The result of the function execution                                                      
	Result                                                                   *TentacledResult    `json:"result,omitempty"`
}

// Configuration for when this function should be automatically called
type TentacledCall struct {
	// If true, this function will be force-called at the end of the conversation        
	End                                                                            *bool `json:"end,omitempty"`
	// If true, this function will be force-called at the start of the conversation      
	Start                                                                          *bool `json:"start,omitempty"`
}

// JSON Schema definition for the function parameters
type TentacledParameters struct {
	// Object property definitions                             
	Properties                          map[string]interface{} `json:"properties"`
	// Required property names                                 
	Required                            []string               `json:"required,omitempty"`
	// The schema type, must be "object"                       
	Type                                ParametersType         `json:"type"`
}

// The result of the function execution
type TentacledResult struct {
	// The data returned by the function (can be any type)            
	Data                                                  interface{} `json:"data"`
	// The channel for streaming function results                     
	Channel                                               *string     `json:"channel,omitempty"`
}

type ConversationMessageReceiveResponse struct {
	// The ID of the created message                                           
	ID                                 string                                  `json:"id"`
	// The text of the message received                                        
	Text                               string                                  `json:"text"`
	// Usage information                                                       
	Usage                              ConversationMessageReceiveResponseUsage `json:"usage"`
}

// Usage information
type ConversationMessageReceiveResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

type ConversationMessageSendParams struct {
	// The ID of the conversation to send the message to       
	ConversationID                                      string `json:"conversationId"`
}

type ConversationMessageSendRequest struct {
	// Known entities                                                                                 
	Entities                                                []ConversationMessageSendRequestEntity    `json:"entities,omitempty"`
	// Extensions to enhance the bot's capabilities                                                   
	Extensions                                              *ConversationMessageSendRequestExtensions `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                          
	Functions                                               []ConversationMessageSendRequestFunction  `json:"functions,omitempty"`
	// The text of the message to send                                                                
	Text                                                    string                                    `json:"text"`
}

// Extracted entity from the message
type ConversationMessageSendRequestEntity struct {
	// Start offset                                     
	Begin                          float64              `json:"begin"`
	// End offset                                       
	End                            float64              `json:"end"`
	Replacement                    *IndecentReplacement `json:"replacement,omitempty"`
	// The text value of the entity                     
	Text                           string               `json:"text"`
	// The entity type                                  
	Type                           string               `json:"type"`
}

type IndecentReplacement struct {
	// Start offset                             
	Begin                               float64 `json:"begin"`
	// End offset                               
	End                                 float64 `json:"end"`
	// The text value of the replacement        
	Text                                string  `json:"text"`
}

// Extensions to enhance the bot's capabilities
type ConversationMessageSendRequestExtensions struct {
	// Additional backstory for the bot                                  
	Backstory                                           *string          `json:"backstory,omitempty"`
	// Inline datasets to provide additional context                     
	Datasets                                            []StickyDataset  `json:"datasets,omitempty"`
	// Feature flags to enable specific bot capabilities                 
	Features                                            []StickyFeature  `json:"features,omitempty"`
	// Inline skillsets to provide additional abilities                  
	Skillsets                                           []StickySkillset `json:"skillsets,omitempty"`
}

type StickyDataset struct {
	// The description of the dataset               
	Description                      *string        `json:"description,omitempty"`
	// The name of the dataset                      
	Name                             *string        `json:"name,omitempty"`
	// The records in the dataset                   
	Records                          []StickyRecord `json:"records"`
}

type StickyRecord struct {
	// Additional metadata for the record                       
	Meta                                 map[string]interface{} `json:"meta,omitempty"`
	// The text content of the record                           
	Text                                 string                 `json:"text"`
}

type StickyFeature struct {
	// The name of the feature to enable                                    
	Name                                             string                 `json:"name"`
	// Optional configuration options for the feature                       
	Options                                          map[string]interface{} `json:"options,omitempty"`
}

type StickySkillset struct {
	// The abilities in the skillset                  
	Abilities                         []StickyAbility `json:"abilities"`
	// The description of the skillset                
	Description                       *string         `json:"description,omitempty"`
	// The name of the skillset                       
	Name                              *string         `json:"name,omitempty"`
}

type StickyAbility struct {
	// The description of the ability                            
	Description                           string                 `json:"description"`
	// The instruction for the ability                           
	Instruction                           string                 `json:"instruction"`
	// Additional metadata for the ability                       
	Meta                                  map[string]interface{} `json:"meta,omitempty"`
	// The name of the ability                                   
	Name                                  string                 `json:"name"`
	// Optional secret ID for the ability                        
	SecretID                              *string                `json:"secretId,omitempty"`
}

type ConversationMessageSendRequestFunction struct {
	// Configuration for when this function should be automatically called                    
	Call                                                                     *StickyCall      `json:"call,omitempty"`
	// The description of the function                                                        
	Description                                                              string           `json:"description"`
	// The name of the function (must be a valid JS identifier, max 64 chars)                 
	Name                                                                     string           `json:"name"`
	// JSON Schema definition for the function parameters                                     
	Parameters                                                               StickyParameters `json:"parameters"`
	// The result of the function execution                                                   
	Result                                                                   *StickyResult    `json:"result,omitempty"`
}

// Configuration for when this function should be automatically called
type StickyCall struct {
	// If true, this function will be force-called at the end of the conversation        
	End                                                                            *bool `json:"end,omitempty"`
	// If true, this function will be force-called at the start of the conversation      
	Start                                                                          *bool `json:"start,omitempty"`
}

// JSON Schema definition for the function parameters
type StickyParameters struct {
	// Object property definitions                             
	Properties                          map[string]interface{} `json:"properties"`
	// Required property names                                 
	Required                            []string               `json:"required,omitempty"`
	// The schema type, must be "object"                       
	Type                                ParametersType         `json:"type"`
}

// The result of the function execution
type StickyResult struct {
	// The data returned by the function (can be any type)            
	Data                                                  interface{} `json:"data"`
	// The channel for streaming function results                     
	Channel                                               *string     `json:"channel,omitempty"`
}

type ConversationMessageSendResponse struct {
	// Extracted entities from the message                                        
	Entities                              []ConversationMessageSendResponseEntity `json:"entities"`
	// The ID of the sent message                                                 
	ID                                    string                                  `json:"id"`
}

// Extracted entity from the message
type ConversationMessageSendResponseEntity struct {
	// Start offset                                      
	Begin                          float64               `json:"begin"`
	// End offset                                        
	End                            float64               `json:"end"`
	Replacement                    *HilariousReplacement `json:"replacement,omitempty"`
	// The text value of the entity                      
	Text                           string                `json:"text"`
	// The entity type                                   
	Type                           string                `json:"type"`
}

type HilariousReplacement struct {
	// Start offset                             
	Begin                               float64 `json:"begin"`
	// End offset                               
	End                                 float64 `json:"end"`
	// The text value of the replacement        
	Text                                string  `json:"text"`
}

type ConversationSessionCreateParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

type ConversationSessionCreateRequest struct {
	// The maximum amount of time this session will stay open         
	DurationInSeconds                                        *float64 `json:"durationInSeconds,omitempty"`
}

type ConversationSessionCreateResponse struct {
	// The time the token will expire in milliseconds        
	ExpiresAt                                        float64 `json:"expiresAt"`
	// The ID of the conversation                            
	ID                                               string  `json:"id"`
	// The token for this conversation                       
	Token                                            string  `json:"token"`
}

type ConversationUpdateParams struct {
	ConversationID string `json:"conversationId"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ConversationUpdateRequest struct {
	// The contact id assigned to this conversation                             
	ContactID                                            *string                `json:"contactId,omitempty"`
	// The associated description                                               
	Description                                          *string                `json:"description,omitempty"`
	// Meta data information                                                    
	Meta                                                 map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                      
	Name                                                 *string                `json:"name,omitempty"`
	// The space id assigned to this conversation                               
	SpaceID                                              *string                `json:"spaceId,omitempty"`
	// The task id assigned to this conversation                                
	TaskID                                               *string                `json:"taskId,omitempty"`
	// The ID of the bot this configuration is using                            
	BotID                                                *string                `json:"botId,omitempty"`
	// The backstory this configuration is using                                
	Backstory                                            *string                `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using                        
	DatasetID                                            *string                `json:"datasetId,omitempty"`
	// A model definition                                                       
	Model                                                *string                `json:"model,omitempty"`
	// The moderation flag for this configuration                               
	Moderation                                           *bool                  `json:"moderation,omitempty"`
	// The privacy flag for this configuration                                  
	Privacy                                              *bool                  `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                       
	SkillsetID                                           *string                `json:"skillsetId,omitempty"`
}

type ConversationUpdateResponse struct {
	// The ID of the updated conversation       
	ID                                   string `json:"id"`
}

type ConversationUpvoteParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

type ConversationUpvoteRequest struct {
	// The reason for the upvote        
	Reason                      *string `json:"reason,omitempty"`
	// The value of the upvote          
	Value                       *int64  `json:"value,omitempty"`
}

type ConversationUpvoteResponse struct {
	// The ID of the upvoted conversation       
	ID                                   string `json:"id"`
}

type ConversationUsageFetchParams struct {
	// The ID of the conversation                            
	ConversationID                                string     `json:"conversationId"`
	// Start date for the period (ISO 8601 format)           
	From                                          *time.Time `json:"from,omitempty"`
	// End date for the period (ISO 8601 format)             
	To                                            *time.Time `json:"to,omitempty"`
}

type ConversationUsageFetchResponse struct {
	// Total number of messages               
	Messages                           *int64 `json:"messages,omitempty"`
	// Total number of BASE tokens used       
	Tokens                             *int64 `json:"tokens,omitempty"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ConversationCompleteRequest struct {
	// An array of attachments to be added to the conversation                                           
	Attachments                                                  []ConversationCompleteRequestAttachment `json:"attachments,omitempty"`
	// The contact ID to associate with this conversation                                                
	ContactID                                                    *ConversationCompleteRequestContactID   `json:"contactId"`
	// Extensions to enhance the bot's capabilities                                                      
	Extensions                                                   *ConversationCompleteRequestExtensions  `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                             
	Functions                                                    []ConversationCompleteRequestFunction   `json:"functions,omitempty"`
	// Execution limits to control conversation processing bounds                                        
	Limits                                                       *ConversationCompleteRequestLimits      `json:"limits,omitempty"`
	// An array of messages to be added to the conversation                                              
	Messages                                                     []ConversationCompleteRequestMessage    `json:"messages"`
	// The ID of the bot this configuration is using                                                     
	BotID                                                        *string                                 `json:"botId,omitempty"`
	// The backstory this configuration is using                                                         
	Backstory                                                    *string                                 `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using                                                 
	DatasetID                                                    *string                                 `json:"datasetId,omitempty"`
	// A model definition                                                                                
	Model                                                        *string                                 `json:"model,omitempty"`
	// The moderation flag for this configuration                                                        
	Moderation                                                   *bool                                   `json:"moderation,omitempty"`
	// The privacy flag for this configuration                                                           
	Privacy                                                      *bool                                   `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                                                
	SkillsetID                                                   *string                                 `json:"skillsetId,omitempty"`
}

type ConversationCompleteRequestAttachment struct {
	// The URL of the attachment        
	URL                         *string `json:"url,omitempty"`
}

// A contact object to create or retrieve a trusted contact
type PurpleContactID struct {
	// A description of the contact                                       
	Description                                    *string                `json:"description,omitempty"`
	// The email address of the contact                                   
	Email                                          *string                `json:"email,omitempty"`
	// A unique fingerprint to identify the contact                       
	Fingerprint                                    string                 `json:"fingerprint"`
	// Additional metadata for the contact                                
	Meta                                           map[string]interface{} `json:"meta,omitempty"`
	// The name of the contact                                            
	Name                                           *string                `json:"name,omitempty"`
	// A nickname for the contact                                         
	Nick                                           *string                `json:"nick,omitempty"`
	// The phone number of the contact                                    
	Phone                                          *string                `json:"phone,omitempty"`
}

// Extensions to enhance the bot's capabilities
type ConversationCompleteRequestExtensions struct {
	// Additional backstory for the bot                                  
	Backstory                                           *string          `json:"backstory,omitempty"`
	// Inline datasets to provide additional context                     
	Datasets                                            []IndigoDataset  `json:"datasets,omitempty"`
	// Feature flags to enable specific bot capabilities                 
	Features                                            []IndigoFeature  `json:"features,omitempty"`
	// Inline skillsets to provide additional abilities                  
	Skillsets                                           []IndigoSkillset `json:"skillsets,omitempty"`
}

type IndigoDataset struct {
	// The description of the dataset               
	Description                      *string        `json:"description,omitempty"`
	// The name of the dataset                      
	Name                             *string        `json:"name,omitempty"`
	// The records in the dataset                   
	Records                          []IndigoRecord `json:"records"`
}

type IndigoRecord struct {
	// Additional metadata for the record                       
	Meta                                 map[string]interface{} `json:"meta,omitempty"`
	// The text content of the record                           
	Text                                 string                 `json:"text"`
}

type IndigoFeature struct {
	// The name of the feature to enable                                    
	Name                                             string                 `json:"name"`
	// Optional configuration options for the feature                       
	Options                                          map[string]interface{} `json:"options,omitempty"`
}

type IndigoSkillset struct {
	// The abilities in the skillset                  
	Abilities                         []IndigoAbility `json:"abilities"`
	// The description of the skillset                
	Description                       *string         `json:"description,omitempty"`
	// The name of the skillset                       
	Name                              *string         `json:"name,omitempty"`
}

type IndigoAbility struct {
	// The description of the ability                            
	Description                           string                 `json:"description"`
	// The instruction for the ability                           
	Instruction                           string                 `json:"instruction"`
	// Additional metadata for the ability                       
	Meta                                  map[string]interface{} `json:"meta,omitempty"`
	// The name of the ability                                   
	Name                                  string                 `json:"name"`
	// Optional secret ID for the ability                        
	SecretID                              *string                `json:"secretId,omitempty"`
}

type ConversationCompleteRequestFunction struct {
	// Configuration for when this function should be automatically called                    
	Call                                                                     *IndigoCall      `json:"call,omitempty"`
	// The description of the function                                                        
	Description                                                              string           `json:"description"`
	// The name of the function (must be a valid JS identifier, max 64 chars)                 
	Name                                                                     string           `json:"name"`
	// JSON Schema definition for the function parameters                                     
	Parameters                                                               IndigoParameters `json:"parameters"`
	// The result of the function execution                                                   
	Result                                                                   *IndigoResult    `json:"result,omitempty"`
}

// Configuration for when this function should be automatically called
type IndigoCall struct {
	// If true, this function will be force-called at the end of the conversation        
	End                                                                            *bool `json:"end,omitempty"`
	// If true, this function will be force-called at the start of the conversation      
	Start                                                                          *bool `json:"start,omitempty"`
}

// JSON Schema definition for the function parameters
type IndigoParameters struct {
	// Object property definitions                             
	Properties                          map[string]interface{} `json:"properties"`
	// Required property names                                 
	Required                            []string               `json:"required,omitempty"`
	// The schema type, must be "object"                       
	Type                                ParametersType         `json:"type"`
}

// The result of the function execution
type IndigoResult struct {
	// The data returned by the function (can be any type)            
	Data                                                  interface{} `json:"data"`
	// The channel for streaming function results                     
	Channel                                               *string     `json:"channel,omitempty"`
}

// Execution limits to control conversation processing bounds
type ConversationCompleteRequestLimits struct {
	// Maximum number of function/tool calls. Controls how many total function calls can be made       
	// during the conversation.                                                                        
	Calls                                                                                       *int64 `json:"calls,omitempty"`
	// Maximum number of model continuations. Controls how many times the model can continue           
	// generating after reaching a stop condition.                                                     
	Continuations                                                                               *int64 `json:"continuations,omitempty"`
	// Maximum number of agentic iterations. Controls how many times the model can iterate             
	// through tool calls and responses.                                                               
	Iterations                                                                                  *int64 `json:"iterations,omitempty"`
}

// A message in the conversation
type ConversationCompleteRequestMessage struct {
	// Meta data information                         
	Meta                      map[string]interface{} `json:"meta,omitempty"`
	// The text of the message                       
	Text                      string                 `json:"text"`
	// The type of the message                       
	Type                      MessageType            `json:"type"`
}

type ConversationCompleteResponse struct {
	// Information about why the completion ended                                  
	End                                          ConversationCompleteResponseEnd   `json:"end"`
	// The text of the message received                                            
	Text                                         string                            `json:"text"`
	// Usage information                                                           
	Usage                                        ConversationCompleteResponseUsage `json:"usage"`
}

// Information about why the completion ended
type ConversationCompleteResponseEnd struct {
	// The reason why the completion ended               
	Reason                                CompleteReason `json:"reason"`
}

// Usage information
type ConversationCompleteResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ConversationCreateRequest struct {
	// The contact id assigned to this conversation                                           
	ContactID                                              *string                            `json:"contactId,omitempty"`
	// The associated description                                                             
	Description                                            *string                            `json:"description,omitempty"`
	// An array of messages to be added to the conversation                                   
	Messages                                               []ConversationCreateRequestMessage `json:"messages,omitempty"`
	// Meta data information                                                                  
	Meta                                                   map[string]interface{}             `json:"meta,omitempty"`
	// The associated name                                                                    
	Name                                                   *string                            `json:"name,omitempty"`
	// The space id assigned to this conversation                                             
	SpaceID                                                *string                            `json:"spaceId,omitempty"`
	// The task id assigned to this conversation                                              
	TaskID                                                 *string                            `json:"taskId,omitempty"`
	// The ID of the bot this configuration is using                                          
	BotID                                                  *string                            `json:"botId,omitempty"`
	// The backstory this configuration is using                                              
	Backstory                                              *string                            `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using                                      
	DatasetID                                              *string                            `json:"datasetId,omitempty"`
	// A model definition                                                                     
	Model                                                  *string                            `json:"model,omitempty"`
	// The moderation flag for this configuration                                             
	Moderation                                             *bool                              `json:"moderation,omitempty"`
	// The privacy flag for this configuration                                                
	Privacy                                                *bool                              `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                                     
	SkillsetID                                             *string                            `json:"skillsetId,omitempty"`
}

type ConversationCreateRequestMessage struct {
	// The text of the message            
	Text                      string      `json:"text"`
	// The type of the message            
	Type                      MessageType `json:"type"`
}

type ConversationCreateResponse struct {
	// The ID of the created conversation                                                   
	ID                                                  string                              `json:"id"`
	// An array of messages included in the conversation                                    
	Messages                                            []ConversationCreateResponseMessage `json:"messages,omitempty"`
}

type ConversationCreateResponseMessage struct {
	// The text of the message            
	Text                      string      `json:"text"`
	// The type of the message            
	Type                      MessageType `json:"type"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ConversationDispatchRequest struct {
	// An array of attachments to be added to the conversation                                           
	Attachments                                                  []ConversationDispatchRequestAttachment `json:"attachments,omitempty"`
	// A unique channel ID to subscribe to for completion events                                         
	ChannelID                                                    *string                                 `json:"channelId,omitempty"`
	// The contact ID to associate with this conversation                                                
	ContactID                                                    *ConversationDispatchRequestContactID   `json:"contactId"`
	// Extensions to enhance the bot's capabilities                                                      
	Extensions                                                   *ConversationDispatchRequestExtensions  `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                             
	Functions                                                    []ConversationDispatchRequestFunction   `json:"functions,omitempty"`
	// Execution limits to control conversation processing bounds                                        
	Limits                                                       *ConversationDispatchRequestLimits      `json:"limits,omitempty"`
	// An array of messages to be added to the conversation                                              
	Messages                                                     []ConversationDispatchRequestMessage    `json:"messages"`
	// The ID of the bot this configuration is using                                                     
	BotID                                                        *string                                 `json:"botId,omitempty"`
	// The backstory this configuration is using                                                         
	Backstory                                                    *string                                 `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using                                                 
	DatasetID                                                    *string                                 `json:"datasetId,omitempty"`
	// A model definition                                                                                
	Model                                                        *string                                 `json:"model,omitempty"`
	// The moderation flag for this configuration                                                        
	Moderation                                                   *bool                                   `json:"moderation,omitempty"`
	// The privacy flag for this configuration                                                           
	Privacy                                                      *bool                                   `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                                                
	SkillsetID                                                   *string                                 `json:"skillsetId,omitempty"`
}

type ConversationDispatchRequestAttachment struct {
	// The URL of the attachment        
	URL                         *string `json:"url,omitempty"`
}

// A contact object to create or retrieve a trusted contact
type FluffyContactID struct {
	// A description of the contact                                       
	Description                                    *string                `json:"description,omitempty"`
	// The email address of the contact                                   
	Email                                          *string                `json:"email,omitempty"`
	// A unique fingerprint to identify the contact                       
	Fingerprint                                    string                 `json:"fingerprint"`
	// Additional metadata for the contact                                
	Meta                                           map[string]interface{} `json:"meta,omitempty"`
	// The name of the contact                                            
	Name                                           *string                `json:"name,omitempty"`
	// A nickname for the contact                                         
	Nick                                           *string                `json:"nick,omitempty"`
	// The phone number of the contact                                    
	Phone                                          *string                `json:"phone,omitempty"`
}

// Extensions to enhance the bot's capabilities
type ConversationDispatchRequestExtensions struct {
	// Additional backstory for the bot                                    
	Backstory                                           *string            `json:"backstory,omitempty"`
	// Inline datasets to provide additional context                       
	Datasets                                            []IndecentDataset  `json:"datasets,omitempty"`
	// Feature flags to enable specific bot capabilities                   
	Features                                            []IndecentFeature  `json:"features,omitempty"`
	// Inline skillsets to provide additional abilities                    
	Skillsets                                           []IndecentSkillset `json:"skillsets,omitempty"`
}

type IndecentDataset struct {
	// The description of the dataset                 
	Description                      *string          `json:"description,omitempty"`
	// The name of the dataset                        
	Name                             *string          `json:"name,omitempty"`
	// The records in the dataset                     
	Records                          []IndecentRecord `json:"records"`
}

type IndecentRecord struct {
	// Additional metadata for the record                       
	Meta                                 map[string]interface{} `json:"meta,omitempty"`
	// The text content of the record                           
	Text                                 string                 `json:"text"`
}

type IndecentFeature struct {
	// The name of the feature to enable                                    
	Name                                             string                 `json:"name"`
	// Optional configuration options for the feature                       
	Options                                          map[string]interface{} `json:"options,omitempty"`
}

type IndecentSkillset struct {
	// The abilities in the skillset                    
	Abilities                         []IndecentAbility `json:"abilities"`
	// The description of the skillset                  
	Description                       *string           `json:"description,omitempty"`
	// The name of the skillset                         
	Name                              *string           `json:"name,omitempty"`
}

type IndecentAbility struct {
	// The description of the ability                            
	Description                           string                 `json:"description"`
	// The instruction for the ability                           
	Instruction                           string                 `json:"instruction"`
	// Additional metadata for the ability                       
	Meta                                  map[string]interface{} `json:"meta,omitempty"`
	// The name of the ability                                   
	Name                                  string                 `json:"name"`
	// Optional secret ID for the ability                        
	SecretID                              *string                `json:"secretId,omitempty"`
}

type ConversationDispatchRequestFunction struct {
	// Configuration for when this function should be automatically called                      
	Call                                                                     *IndecentCall      `json:"call,omitempty"`
	// The description of the function                                                          
	Description                                                              string             `json:"description"`
	// The name of the function (must be a valid JS identifier, max 64 chars)                   
	Name                                                                     string             `json:"name"`
	// JSON Schema definition for the function parameters                                       
	Parameters                                                               IndecentParameters `json:"parameters"`
	// The result of the function execution                                                     
	Result                                                                   *IndecentResult    `json:"result,omitempty"`
}

// Configuration for when this function should be automatically called
type IndecentCall struct {
	// If true, this function will be force-called at the end of the conversation        
	End                                                                            *bool `json:"end,omitempty"`
	// If true, this function will be force-called at the start of the conversation      
	Start                                                                          *bool `json:"start,omitempty"`
}

// JSON Schema definition for the function parameters
type IndecentParameters struct {
	// Object property definitions                             
	Properties                          map[string]interface{} `json:"properties"`
	// Required property names                                 
	Required                            []string               `json:"required,omitempty"`
	// The schema type, must be "object"                       
	Type                                ParametersType         `json:"type"`
}

// The result of the function execution
type IndecentResult struct {
	// The data returned by the function (can be any type)            
	Data                                                  interface{} `json:"data"`
	// The channel for streaming function results                     
	Channel                                               *string     `json:"channel,omitempty"`
}

// Execution limits to control conversation processing bounds
type ConversationDispatchRequestLimits struct {
	// Maximum number of function/tool calls. Controls how many total function calls can be made       
	// during the conversation.                                                                        
	Calls                                                                                       *int64 `json:"calls,omitempty"`
	// Maximum number of model continuations. Controls how many times the model can continue           
	// generating after reaching a stop condition.                                                     
	Continuations                                                                               *int64 `json:"continuations,omitempty"`
	// Maximum number of agentic iterations. Controls how many times the model can iterate             
	// through tool calls and responses.                                                               
	Iterations                                                                                  *int64 `json:"iterations,omitempty"`
}

// A message in the conversation
type ConversationDispatchRequestMessage struct {
	// Meta data information                         
	Meta                      map[string]interface{} `json:"meta,omitempty"`
	// The text of the message                       
	Text                      string                 `json:"text"`
	// The type of the message                       
	Type                      MessageType            `json:"type"`
}

type ConversationDispatchResponse struct {
	// The channel ID to subscribe to for completion events       
	ChannelID                                              string `json:"channelId"`
}

type ConversationsExportParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ConversationsExportResponse struct {
	Items []ConversationsExportResponseItem `json:"items"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ConversationsExportResponseItem struct {
	// The contact id assigned to this conversation                             
	ContactID                                            *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                         
	CreatedAt                                            float64                `json:"createdAt"`
	// The associated description                                               
	Description                                          *string                `json:"description,omitempty"`
	// The instance ID                                                          
	ID                                                   string                 `json:"id"`
	// Meta data information                                                    
	Meta                                                 map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                      
	Name                                                 *string                `json:"name,omitempty"`
	// The space id assigned to this conversation                               
	SpaceID                                              *string                `json:"spaceId,omitempty"`
	// The task id assigned to this conversation                                
	TaskID                                               *string                `json:"taskId,omitempty"`
	// The timestamp (ms) when the instance was updated                         
	UpdatedAt                                            float64                `json:"updatedAt"`
	// The ID of the bot this configuration is using                            
	BotID                                                *string                `json:"botId,omitempty"`
	// The backstory this configuration is using                                
	Backstory                                            *string                `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using                        
	DatasetID                                            *string                `json:"datasetId,omitempty"`
	// A model definition                                                       
	Model                                                *string                `json:"model,omitempty"`
	// The moderation flag for this configuration                               
	Moderation                                           *bool                  `json:"moderation,omitempty"`
	// The privacy flag for this configuration                                  
	Privacy                                              *bool                  `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                       
	SkillsetID                                           *string                `json:"skillsetId,omitempty"`
}

type ConversationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ConversationsListResponse struct {
	Items []ConversationsListResponseItem `json:"items"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ConversationsListResponseItem struct {
	// The contact id assigned to this conversation                             
	ContactID                                            *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                         
	CreatedAt                                            float64                `json:"createdAt"`
	// The associated description                                               
	Description                                          *string                `json:"description,omitempty"`
	// The instance ID                                                          
	ID                                                   string                 `json:"id"`
	// Meta data information                                                    
	Meta                                                 map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                      
	Name                                                 *string                `json:"name,omitempty"`
	// The space id assigned to this conversation                               
	SpaceID                                              *string                `json:"spaceId,omitempty"`
	// The task id assigned to this conversation                                
	TaskID                                               *string                `json:"taskId,omitempty"`
	// The timestamp (ms) when the instance was updated                         
	UpdatedAt                                            float64                `json:"updatedAt"`
	// The ID of the bot this configuration is using                            
	BotID                                                *string                `json:"botId,omitempty"`
	// The backstory this configuration is using                                
	Backstory                                            *string                `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using                        
	DatasetID                                            *string                `json:"datasetId,omitempty"`
	// A model definition                                                       
	Model                                                *string                `json:"model,omitempty"`
	// The moderation flag for this configuration                               
	Moderation                                           *bool                  `json:"moderation,omitempty"`
	// The privacy flag for this configuration                                  
	Privacy                                              *bool                  `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using                       
	SkillsetID                                           *string                `json:"skillsetId,omitempty"`
}

type DatasetDeleteParams struct {
	// The ID of the dataset to delete       
	DatasetID                         string `json:"datasetId"`
}

type DatasetDeleteResponse struct {
	// The ID of the deleted dataset       
	ID                              string `json:"id"`
}

type DatasetFetchParams struct {
	// The ID of the dataset to retrieve       
	DatasetID                           string `json:"datasetId"`
}

// Blueprint properties
type DatasetFetchResponse struct {
	// The ID of the blueprint                                                   
	BlueprintID                                           *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                          
	CreatedAt                                             float64                `json:"createdAt"`
	// The associated description                                                
	Description                                           *string                `json:"description,omitempty"`
	// The instance ID                                                           
	ID                                                    string                 `json:"id"`
	// An instruction to include before found records                            
	MatchInstruction                                      *string                `json:"matchInstruction,omitempty"`
	// Meta data information                                                     
	Meta                                                  map[string]interface{} `json:"meta,omitempty"`
	// An instruction to include if no records where found                       
	MismatchInstruction                                   *string                `json:"mismatchInstruction,omitempty"`
	// The associated name                                                       
	Name                                                  *string                `json:"name,omitempty"`
	// The total number of tokens for each record                                
	RecordMaxTokens                                       *float64               `json:"recordMaxTokens,omitempty"`
	// The reranker class for the dataset                                        
	Reranker                                              *string                `json:"reranker,omitempty"`
	// The total number of records to return during search                       
	SearchMaxRecords                                      *float64               `json:"searchMaxRecords,omitempty"`
	// The total number of tokens to use during search                           
	SearchMaxTokens                                       *float64               `json:"searchMaxTokens,omitempty"`
	// The minimum score to filter search results by                             
	SearchMinScore                                        *float64               `json:"searchMinScore,omitempty"`
	// A list of separators to use when tokenizing text                          
	Separators                                            *string                `json:"separators,omitempty"`
	// The storage class for the dataset                                         
	Store                                                 string                 `json:"store"`
	// The timestamp (ms) when the instance was updated                          
	UpdatedAt                                             float64                `json:"updatedAt"`
	// The dataset visibility                                                    
	Visibility                                            *SecretVisibility      `json:"visibility,omitempty"`
}

type DatasetFileAttachParams struct {
	// The ID of the dataset       
	DatasetID               string `json:"datasetId"`
	// The ID of the file          
	FileID                  string `json:"fileId"`
}

type DatasetFileAttachRequest struct {
	// The dataset file attachment type                           
	Type                               *DatasetFileAttachmentType `json:"type,omitempty"`
}

type DatasetFileAttachResponse struct {
	// The ID of the dataset file       
	ID                           string `json:"id"`
}

type DatasetFileDetachParams struct {
	// The ID of the dataset       
	DatasetID               string `json:"datasetId"`
	// The ID of the file          
	FileID                  string `json:"fileId"`
}

type DatasetFileDetachRequest struct {
	// Delete records associated with the file      
	DeleteRecords                             *bool `json:"deleteRecords,omitempty"`
}

type DatasetFileDetachResponse struct {
	// The ID of the dataset file       
	ID                           string `json:"id"`
}

type DatasetFileSyncParams struct {
	// The ID of the dataset       
	DatasetID               string `json:"datasetId"`
	// The ID of the file          
	FileID                  string `json:"fileId"`
}

type DatasetFileSyncResponse struct {
	// The ID of the dataset file       
	ID                           string `json:"id"`
}

type DatasetFilesListParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The ID of the dataset                   
	DatasetID                          string  `json:"datasetId"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type DatasetFilesListResponse struct {
	Items []DatasetFilesListResponseItem `json:"items"`
}

// Instance list properties
type DatasetFilesListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The file visibility                                                    
	Visibility                                         *SecretVisibility      `json:"visibility,omitempty"`
}

type DatasetRecordDeleteParams struct {
	// The ID of the dataset                
	DatasetID                        string `json:"datasetId"`
	// The ID of the record to delete       
	RecordID                         string `json:"recordId"`
}

type DatasetRecordDeleteResponse struct {
	// The ID of the deleted record       
	ID                             string `json:"id"`
}

type DatasetRecordFetchParams struct {
	// The ID of the dataset                  
	DatasetID                          string `json:"datasetId"`
	// The ID of the record to retrieve       
	RecordID                           string `json:"recordId"`
}

// Instance list properties
type DatasetRecordFetchResponse struct {
	// The timestamp (ms) when the instance was created        
	CreatedAt                                          float64 `json:"createdAt"`
	// The instance ID                                         
	ID                                                 string  `json:"id"`
	// The source of the dataset record                        
	Source                                             *string `json:"source,omitempty"`
	// The text of the dataset record                          
	Text                                               string  `json:"text"`
	// The timestamp (ms) when the instance was updated        
	UpdatedAt                                          float64 `json:"updatedAt"`
}

type DatasetRecordUpdateParams struct {
	DatasetID string `json:"datasetId"`
	RecordID  string `json:"recordId"`
}

type DatasetRecordUpdateRequest struct {
	// Meta data information                                      
	Meta                                   map[string]interface{} `json:"meta,omitempty"`
	// The source to update the record with                       
	Source                                 *string                `json:"source,omitempty"`
	// The text to update the record with                         
	Text                                   *string                `json:"text,omitempty"`
}

type DatasetRecordUpdateResponse struct {
	// The ID of the updated record       
	ID                             string `json:"id"`
}

type DatasetRecordCreateParams struct {
	DatasetID string `json:"datasetId"`
}

type DatasetRecordCreateRequest struct {
	// Meta data information                          
	Meta                       map[string]interface{} `json:"meta,omitempty"`
	// The source of the record                       
	Source                     *string                `json:"source,omitempty"`
	// The text of the record                         
	Text                       string                 `json:"text"`
}

type DatasetRecordCreateResponse struct {
	// The ID of the created record       
	ID                             string `json:"id"`
}

type DatasetRecordsExportParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The ID of the dataset to export         
	DatasetID                          string  `json:"datasetId"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type DatasetRecordsExportResponse struct {
	Items []DatasetRecordsExportResponseItem `json:"items"`
}

// Instance list properties
type DatasetRecordsExportResponseItem struct {
	// The timestamp (ms) when the instance was created        
	CreatedAt                                          float64 `json:"createdAt"`
	// The instance ID                                         
	ID                                                 string  `json:"id"`
	Source                                             *string `json:"source,omitempty"`
	Text                                               string  `json:"text"`
	// The timestamp (ms) when the instance was updated        
	UpdatedAt                                          float64 `json:"updatedAt"`
}

type DatasetRecordsListParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The ID of the dataset                   
	DatasetID                          string  `json:"datasetId"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type DatasetRecordsListResponse struct {
	Items []DatasetRecordsListResponseItem `json:"items"`
}

// Instance list properties
type DatasetRecordsListResponseItem struct {
	// The timestamp (ms) when the instance was created        
	CreatedAt                                          float64 `json:"createdAt"`
	// The instance ID                                         
	ID                                                 string  `json:"id"`
	Source                                             *string `json:"source,omitempty"`
	Text                                               string  `json:"text"`
	// The timestamp (ms) when the instance was updated        
	UpdatedAt                                          float64 `json:"updatedAt"`
}

type DatasetSearchParams struct {
	// The ID of the dataset to search       
	DatasetID                         string `json:"datasetId"`
}

type DatasetSearchRequest struct {
	Filter                             map[string]*FilterValue `json:"filter,omitempty"`
	// The keyword/phrase to search for                        
	Search                             string                  `json:"search"`
}

type FilterClass struct {
	Eq  *Eq      `json:"$eq"`
	Ne  *Eq      `json:"$ne"`
	Gt  *float64 `json:"$gt,omitempty"`
	Gte *float64 `json:"$gte,omitempty"`
	Lt  *float64 `json:"$lt,omitempty"`
	LTE *float64 `json:"$lte,omitempty"`
}

type DatasetSearchResponse struct {
	// The ID of the dataset that was searched                                    
	ID                                              string                        `json:"id"`
	// An array of records matching the search query                              
	Records                                         []DatasetSearchResponseRecord `json:"records"`
}

type DatasetSearchResponseRecord struct {
	ID     string                 `json:"id"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
	Score  float64                `json:"score"`
	Source *string                `json:"source,omitempty"`
	Text   string                 `json:"text"`
}

type DatasetUpdateParams struct {
	DatasetID string `json:"datasetId"`
}

// Blueprint properties
type DatasetUpdateRequest struct {
	// The unique alias for the instance                                         
	Alias                                                 *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                                   
	BlueprintID                                           *string                `json:"blueprintId,omitempty"`
	// The associated description                                                
	Description                                           *string                `json:"description,omitempty"`
	// An instruction to include before found records                            
	MatchInstruction                                      *string                `json:"matchInstruction,omitempty"`
	// Meta data information                                                     
	Meta                                                  map[string]interface{} `json:"meta,omitempty"`
	// An instruction to include if no records where found                       
	MismatchInstruction                                   *string                `json:"mismatchInstruction,omitempty"`
	// The associated name                                                       
	Name                                                  *string                `json:"name,omitempty"`
	// The total number of tokens to for each record                             
	RecordMaxTokens                                       *float64               `json:"recordMaxTokens,omitempty"`
	// The reranker class for the dataset                                        
	Reranker                                              *string                `json:"reranker,omitempty"`
	// The total number of records to return during search                       
	SearchMaxRecords                                      *float64               `json:"searchMaxRecords,omitempty"`
	// The total number of tokens to use during search                           
	SearchMaxTokens                                       *float64               `json:"searchMaxTokens,omitempty"`
	// The minimum score to filter search results by                             
	SearchMinScore                                        *float64               `json:"searchMinScore,omitempty"`
	// A list of separators to use when tokenizing text                          
	Separators                                            *string                `json:"separators,omitempty"`
	// The dataset visibility                                                    
	Visibility                                            *SecretVisibility      `json:"visibility,omitempty"`
}

type DatasetUpdateResponse struct {
	// The ID of the updated dataset       
	ID                              string `json:"id"`
}

// Blueprint properties
type DatasetCreateRequest struct {
	// The unique alias for the instance                                         
	Alias                                                 *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                                   
	BlueprintID                                           *string                `json:"blueprintId,omitempty"`
	// The associated description                                                
	Description                                           *string                `json:"description,omitempty"`
	// An instruction to include before found records                            
	MatchInstruction                                      *string                `json:"matchInstruction,omitempty"`
	// Meta data information                                                     
	Meta                                                  map[string]interface{} `json:"meta,omitempty"`
	// An instruction to include if no records where found                       
	MismatchInstruction                                   *string                `json:"mismatchInstruction,omitempty"`
	// The associated name                                                       
	Name                                                  *string                `json:"name,omitempty"`
	// The total number of tokens for each record                                
	RecordMaxTokens                                       *float64               `json:"recordMaxTokens,omitempty"`
	// The reranker class for the dataset                                        
	Reranker                                              *string                `json:"reranker,omitempty"`
	// The total number of records to return during search                       
	SearchMaxRecords                                      *float64               `json:"searchMaxRecords,omitempty"`
	// The total number of tokens to use during search                           
	SearchMaxTokens                                       *float64               `json:"searchMaxTokens,omitempty"`
	// The minimum score to filter search results by                             
	SearchMinScore                                        *float64               `json:"searchMinScore,omitempty"`
	// A list of separators to use when tokenizing text                          
	Separators                                            *string                `json:"separators,omitempty"`
	// The storage class for the dataset                                         
	Store                                                 *string                `json:"store,omitempty"`
	// The dataset visibility                                                    
	Visibility                                            *SecretVisibility      `json:"visibility,omitempty"`
}

type DatasetCreateResponse struct {
	// The ID of the created dataset       
	ID                              string `json:"id"`
}

type DatasetsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type DatasetsListResponse struct {
	Items []DatasetsListResponseItem `json:"items"`
}

// Blueprint properties
type DatasetsListResponseItem struct {
	// The ID of the blueprint                                                   
	BlueprintID                                           *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                          
	CreatedAt                                             float64                `json:"createdAt"`
	// The associated description                                                
	Description                                           *string                `json:"description,omitempty"`
	// The instance ID                                                           
	ID                                                    string                 `json:"id"`
	// An instruction to include before found records                            
	MatchInstruction                                      *string                `json:"matchInstruction,omitempty"`
	// Meta data information                                                     
	Meta                                                  map[string]interface{} `json:"meta,omitempty"`
	// An instruction to include if no records where found                       
	MismatchInstruction                                   *string                `json:"mismatchInstruction,omitempty"`
	// The associated name                                                       
	Name                                                  *string                `json:"name,omitempty"`
	// The total number of tokens for each record                                
	RecordMaxTokens                                       *float64               `json:"recordMaxTokens,omitempty"`
	// The reranker class for the dataset                                        
	Reranker                                              *string                `json:"reranker,omitempty"`
	// The total number of records to return during search                       
	SearchMaxRecords                                      *float64               `json:"searchMaxRecords,omitempty"`
	// The total number of tokens to use during search                           
	SearchMaxTokens                                       *float64               `json:"searchMaxTokens,omitempty"`
	// The minimum score to filter search results by                             
	SearchMinScore                                        *float64               `json:"searchMinScore,omitempty"`
	// A list of separators to use when tokenizing text                          
	Separators                                            *string                `json:"separators,omitempty"`
	// The storage class for the dataset                                         
	Store                                                 string                 `json:"store"`
	// The timestamp (ms) when the instance was updated                          
	UpdatedAt                                             float64                `json:"updatedAt"`
	// The dataset visibility                                                    
	Visibility                                            *SecretVisibility      `json:"visibility,omitempty"`
}

type EventLogsExportParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type EventLogsExportResponse struct {
	Items []EventLogsExportResponseItem `json:"items"`
}

// Instance list properties
type EventLogsExportResponseItem struct {
	// Related ability ID if applicable                                       
	AbilityID                                          *string                `json:"abilityId,omitempty"`
	// Related blueprint ID if applicable                                     
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// Related bot ID if applicable                                           
	BotID                                              *string                `json:"botId,omitempty"`
	// Related contact ID if applicable                                       
	ContactID                                          *string                `json:"contactId,omitempty"`
	// Related conversation ID if applicable                                  
	ConversationID                                     *string                `json:"conversationId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// Related dataset ID if applicable                                       
	DatasetID                                          *string                `json:"datasetId,omitempty"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// Related Discord integration ID if applicable                           
	DiscordIntegrationID                               *string                `json:"discordIntegrationId,omitempty"`
	// Related email integration ID if applicable                             
	EmailIntegrationID                                 *string                `json:"emailIntegrationId,omitempty"`
	// Related extract integration ID if applicable                           
	ExtractIntegrationID                               *string                `json:"extractIntegrationId,omitempty"`
	// Related file ID if applicable                                          
	FileID                                             *string                `json:"fileId,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Related MCP server integration ID if applicable                        
	McpserverIntegrationID                             *string                `json:"mcpserverIntegrationId,omitempty"`
	// Related Messenger integration ID if applicable                         
	MessengerIntegrationID                             *string                `json:"messengerIntegrationId,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// Related Notion integration ID if applicable                            
	NotionIntegrationID                                *string                `json:"notionIntegrationId,omitempty"`
	// Related portal ID if applicable                                        
	PortalID                                           *string                `json:"portalId,omitempty"`
	// Related record ID if applicable                                        
	RecordID                                           *string                `json:"recordId,omitempty"`
	// Related secret ID if applicable                                        
	SecretID                                           *string                `json:"secretId,omitempty"`
	// Related sitemap integration ID if applicable                           
	SitemapIntegrationID                               *string                `json:"sitemapIntegrationId,omitempty"`
	// Related skillset ID if applicable                                      
	SkillsetID                                         *string                `json:"skillsetId,omitempty"`
	// Related Slack integration ID if applicable                             
	SlackIntegrationID                                 *string                `json:"slackIntegrationId,omitempty"`
	// Related support integration ID if applicable                           
	SupportIntegrationID                               *string                `json:"supportIntegrationId,omitempty"`
	// Related task ID if applicable                                          
	TaskID                                             *string                `json:"taskId,omitempty"`
	// Related Telegram integration ID if applicable                          
	TelegramIntegrationID                              *string                `json:"telegramIntegrationId,omitempty"`
	// Related trigger integration ID if applicable                           
	TriggerIntegrationID                               *string                `json:"triggerIntegrationId,omitempty"`
	// Related Twilio integration ID if applicable                            
	TwilioIntegrationID                                *string                `json:"twilioIntegrationId,omitempty"`
	// The type of event (e.g., 'conversation.create')                        
	Type                                               string                 `json:"type"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// Related WhatsApp integration ID if applicable                          
	WhatsappIntegrationID                              *string                `json:"whatsappIntegrationId,omitempty"`
	// Related widget integration ID if applicable                            
	WidgetIntegrationID                                *string                `json:"widgetIntegrationId,omitempty"`
}

type EventLogsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type EventLogsListResponse struct {
	Items []EventLogsListResponseItem `json:"items"`
}

// Instance list properties
type EventLogsListResponseItem struct {
	// Related ability ID if applicable                                       
	AbilityID                                          *string                `json:"abilityId,omitempty"`
	// Related blueprint ID if applicable                                     
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// Related bot ID if applicable                                           
	BotID                                              *string                `json:"botId,omitempty"`
	// Related contact ID if applicable                                       
	ContactID                                          *string                `json:"contactId,omitempty"`
	// Related conversation ID if applicable                                  
	ConversationID                                     *string                `json:"conversationId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// Related dataset ID if applicable                                       
	DatasetID                                          *string                `json:"datasetId,omitempty"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// Related Discord integration ID if applicable                           
	DiscordIntegrationID                               *string                `json:"discordIntegrationId,omitempty"`
	// Related email integration ID if applicable                             
	EmailIntegrationID                                 *string                `json:"emailIntegrationId,omitempty"`
	// Related extract integration ID if applicable                           
	ExtractIntegrationID                               *string                `json:"extractIntegrationId,omitempty"`
	// Related file ID if applicable                                          
	FileID                                             *string                `json:"fileId,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Related MCP server integration ID if applicable                        
	McpserverIntegrationID                             *string                `json:"mcpserverIntegrationId,omitempty"`
	// Related Messenger integration ID if applicable                         
	MessengerIntegrationID                             *string                `json:"messengerIntegrationId,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// Related Notion integration ID if applicable                            
	NotionIntegrationID                                *string                `json:"notionIntegrationId,omitempty"`
	// Related portal ID if applicable                                        
	PortalID                                           *string                `json:"portalId,omitempty"`
	// Related record ID if applicable                                        
	RecordID                                           *string                `json:"recordId,omitempty"`
	// Related secret ID if applicable                                        
	SecretID                                           *string                `json:"secretId,omitempty"`
	// Related sitemap integration ID if applicable                           
	SitemapIntegrationID                               *string                `json:"sitemapIntegrationId,omitempty"`
	// Related skillset ID if applicable                                      
	SkillsetID                                         *string                `json:"skillsetId,omitempty"`
	// Related Slack integration ID if applicable                             
	SlackIntegrationID                                 *string                `json:"slackIntegrationId,omitempty"`
	// Related support integration ID if applicable                           
	SupportIntegrationID                               *string                `json:"supportIntegrationId,omitempty"`
	// Related task ID if applicable                                          
	TaskID                                             *string                `json:"taskId,omitempty"`
	// Related Telegram integration ID if applicable                          
	TelegramIntegrationID                              *string                `json:"telegramIntegrationId,omitempty"`
	// Related trigger integration ID if applicable                           
	TriggerIntegrationID                               *string                `json:"triggerIntegrationId,omitempty"`
	// Related Twilio integration ID if applicable                            
	TwilioIntegrationID                                *string                `json:"twilioIntegrationId,omitempty"`
	// The type of event (e.g., 'conversation.create')                        
	Type                                               string                 `json:"type"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// Related WhatsApp integration ID if applicable                          
	WhatsappIntegrationID                              *string                `json:"whatsappIntegrationId,omitempty"`
	// Related widget integration ID if applicable                            
	WidgetIntegrationID                                *string                `json:"widgetIntegrationId,omitempty"`
}

type EventLogsSubscribeRequest struct {
	// Number of recent historical events to replay before              
	// subscribing to live updates. When provided, the subscriber       
	// will first receive up to this many recent events that were       
	// logged before the subscription started. This is useful for       
	// catching up on events that may have occurred during              
	// connection setup.                                                
	HistoryLength                                                *int64 `json:"historyLength,omitempty"`
}

type FileDeleteParams struct {
	// The ID of the file to delete       
	FileID                         string `json:"fileId"`
}

type FileDeleteResponse struct {
	// The ID of the deleted file       
	ID                           string `json:"id"`
}

type FileDownloadParams struct {
	// The ID of the file to download       
	FileID                           string `json:"fileId"`
}

type FileDownloadResponse struct {
	// The URL to download the file       
	URL                            string `json:"url"`
}

type FileFetchParams struct {
	// The ID of the file to retrieve       
	FileID                           string `json:"fileId"`
}

// Blueprint properties
type FileFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The file visibility                                                    
	Visibility                                         *SecretVisibility      `json:"visibility,omitempty"`
}

type FileSyncParams struct {
	// The ID of the file to sync       
	FileID                       string `json:"fileId"`
}

type FileSyncResponse struct {
	// The ID of the file       
	ID                   string `json:"id"`
}

type FileUpdateParams struct {
	FileID string `json:"fileId"`
}

// Blueprint properties
type FileUpdateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                 
	BlueprintID                         *string                `json:"blueprintId,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The file visibility                                     
	Visibility                          *SecretVisibility      `json:"visibility,omitempty"`
}

type FileUpdateResponse struct {
	// The ID of the updated file       
	ID                           string `json:"id"`
}

type FileUploadParams struct {
	FileID string `json:"fileId"`
}

type FileUploadRequest struct {
	// The file to upload either as http: or data: URL                       
	//                                                                       
	// The file definition to upload                                         
	File                                              *FileUploadRequestFile `json:"file"`
}

// The file definition to upload
type FluffyFile struct {
	// The file name        
	Name            *string `json:"name,omitempty"`
	// The file size        
	Size            float64 `json:"size"`
	// The file type        
	Type            string  `json:"type"`
}

type FileUploadResponse struct {
	// The ID of the upload file                                               
	ID                                        string                           `json:"id"`
	// The request required to upload the file                                 
	UploadRequest                             *FileUploadResponseUploadRequest `json:"uploadRequest,omitempty"`
}

// The request required to upload the file
type FileUploadResponseUploadRequest struct {
	// The HTTP headers to use                       
	Headers                   map[string]interface{} `json:"headers"`
	// The HTTP method to use                        
	Method                    string                 `json:"method"`
	// The HTTP url to use                           
	URL                       string                 `json:"url"`
}

// Blueprint properties
type FileCreateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                 
	BlueprintID                         *string                `json:"blueprintId,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The file visibility                                     
	Visibility                          *SecretVisibility      `json:"visibility,omitempty"`
}

type FileCreateResponse struct {
	// The ID of the created file       
	ID                           string `json:"id"`
}

type FilesListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type FilesListResponse struct {
	Items []FilesListResponseItem `json:"items"`
}

// Blueprint properties
type FilesListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The file visibility                                                    
	Visibility                                         *SecretVisibility      `json:"visibility,omitempty"`
}

type DiscordIntegrationDeleteParams struct {
	// The ID of the Discord integration       
	DiscordIntegrationID                string `json:"discordIntegrationId"`
}

type DiscordIntegrationDeleteResponse struct {
	// The ID of the deleted Discord integration       
	ID                                          string `json:"id"`
}

type DiscordIntegrationFetchParams struct {
	// The ID of the Discord integration to retrieve       
	DiscordIntegrationID                            string `json:"discordIntegrationId"`
}

// Blueprint properties
type DiscordIntegrationFetchResponse struct {
	// The Discord application ID                                             
	AppID                                              *string                `json:"appId,omitempty"`
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                          
	BotID                                              *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                            
	ContactCollection                                  *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The Discord command handle                                             
	Handle                                             *string                `json:"handle,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The chat session duration                                              
	SessionDuration                                    *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type DiscordIntegrationSetupParams struct {
	// The ID of the Discord integration       
	DiscordIntegrationID                string `json:"discordIntegrationId"`
}

type DiscordIntegrationSetupResponse struct {
	// The ID of the setup Discord integration       
	ID                                        string `json:"id"`
}

type DiscordIntegrationUpdateParams struct {
	// The ID of the Discord integration       
	DiscordIntegrationID                string `json:"discordIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type DiscordIntegrationUpdateRequest struct {
	// The Discord application ID                                          
	AppID                                           *string                `json:"appId,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// The Discord bot token                                               
	BotToken                                        *string                `json:"botToken,omitempty"`
	// Weather to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// The Discord command handle                                          
	Handle                                          *string                `json:"handle,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The Discord public key                                              
	PublicKey                                       *string                `json:"publicKey,omitempty"`
	// The chat session duration                                           
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type DiscordIntegrationUpdateResponse struct {
	// The ID of the Discord Integration       
	ID                                  string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type DiscordIntegrationCreateRequest struct {
	// The Discord application ID                                          
	AppID                                           *string                `json:"appId,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// The Discord bot token                                               
	BotToken                                        *string                `json:"botToken,omitempty"`
	// Weather to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// The Discord command handle                                          
	Handle                                          *string                `json:"handle,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The Discord public key                                              
	PublicKey                                       *string                `json:"publicKey,omitempty"`
	// The chat session duration                                           
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type DiscordIntegrationCreateResponse struct {
	// The ID of the Discord Integration       
	ID                                  string `json:"id"`
}

type DiscordIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type DiscordIntegrationsListResponse struct {
	Items []DiscordIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type DiscordIntegrationsListResponseItem struct {
	// The Discord application ID                                             
	AppID                                              *string                `json:"appId,omitempty"`
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                          
	BotID                                              *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                            
	ContactCollection                                  *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The Discord command handle                                             
	Handle                                             *string                `json:"handle,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The chat session duration                                              
	SessionDuration                                    *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type EmailIntegrationDeleteParams struct {
	// The ID of the Email integration       
	EmailIntegrationID                string `json:"emailIntegrationId"`
}

type EmailIntegrationDeleteResponse struct {
	// The ID of the deleted Email integration       
	ID                                        string `json:"id"`
}

type EmailIntegrationFetchParams struct {
	// The ID of the Email integration to retrieve       
	EmailIntegrationID                            string `json:"emailIntegrationId"`
}

// Blueprint properties
type EmailIntegrationFetchResponse struct {
	// Newline-separated list of email patterns allowed to send messages to this integration                       
	AllowedSenderEmails                                                                     *string                `json:"allowedSenderEmails,omitempty"`
	// Weather the bot supports attachments                                                                        
	Attachments                                                                             *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                     
	BlueprintID                                                                             *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                               
	BotID                                                                                   *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                                                                 
	ContactCollection                                                                       *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                                            
	CreatedAt                                                                               float64                `json:"createdAt"`
	// The associated description                                                                                  
	Description                                                                             *string                `json:"description,omitempty"`
	// The instance ID                                                                                             
	ID                                                                                      string                 `json:"id"`
	// Meta data information                                                                                       
	Meta                                                                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                         
	Name                                                                                    *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                                                      
	SessionDuration                                                                         *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                                                            
	UpdatedAt                                                                               float64                `json:"updatedAt"`
}

type EmailIntegrationSetupParams struct {
	// The ID of the Email integration       
	EmailIntegrationID                string `json:"emailIntegrationId"`
}

type EmailIntegrationSetupResponse struct {
	// The ID of the Email Integration       
	ID                                string `json:"id"`
}

type EmailIntegrationUpdateParams struct {
	// The ID of the Email integration       
	EmailIntegrationID                string `json:"emailIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type EmailIntegrationUpdateRequest struct {
	// Newline-separated list of email patterns allowed to send messages to this integration                       
	AllowedSenderEmails                                                                     *string                `json:"allowedSenderEmails,omitempty"`
	// Weather the bot supports attachments                                                                        
	Attachments                                                                             *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                     
	BlueprintID                                                                             *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                               
	BotID                                                                                   *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                                                                 
	ContactCollection                                                                       *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                                                                  
	Description                                                                             *string                `json:"description,omitempty"`
	// Meta data information                                                                                       
	Meta                                                                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                         
	Name                                                                                    *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                                                      
	SessionDuration                                                                         *float64               `json:"sessionDuration,omitempty"`
}

type EmailIntegrationUpdateResponse struct {
	// The ID of the Email Integration       
	ID                                string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type EmailIntegrationCreateRequest struct {
	// Newline-separated list of email patterns allowed to send messages to this integration                       
	AllowedSenderEmails                                                                     *string                `json:"allowedSenderEmails,omitempty"`
	// Weather the bot supports attachments                                                                        
	Attachments                                                                             *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                     
	BlueprintID                                                                             *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                               
	BotID                                                                                   *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                                                                 
	ContactCollection                                                                       *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                                                                  
	Description                                                                             *string                `json:"description,omitempty"`
	// Meta data information                                                                                       
	Meta                                                                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                         
	Name                                                                                    *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                                                      
	SessionDuration                                                                         *float64               `json:"sessionDuration,omitempty"`
}

type EmailIntegrationCreateResponse struct {
	// The ID of the Email Integration       
	ID                                string `json:"id"`
}

type EmailIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type EmailIntegrationsListResponse struct {
	Items []EmailIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type EmailIntegrationsListResponseItem struct {
	// Newline-separated list of email patterns allowed to send messages to this integration                       
	AllowedSenderEmails                                                                     *string                `json:"allowedSenderEmails,omitempty"`
	// Weather the bot supports attachments                                                                        
	Attachments                                                                             *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                     
	BlueprintID                                                                             *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                               
	BotID                                                                                   *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                                                                 
	ContactCollection                                                                       *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                                            
	CreatedAt                                                                               float64                `json:"createdAt"`
	// The associated description                                                                                  
	Description                                                                             *string                `json:"description,omitempty"`
	// The instance ID                                                                                             
	ID                                                                                      string                 `json:"id"`
	// Meta data information                                                                                       
	Meta                                                                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                         
	Name                                                                                    *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                                                      
	SessionDuration                                                                         *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                                                            
	UpdatedAt                                                                               float64                `json:"updatedAt"`
}

type ExtractIntegrationDeleteParams struct {
	// The ID of the Extract integration       
	ExtractIntegrationID                string `json:"extractIntegrationId"`
}

type ExtractIntegrationDeleteResponse struct {
	// The ID of the deleted Extract integration       
	ID                                          string `json:"id"`
}

type ExtractIntegrationFetchParams struct {
	// The ID of the Extract integration to retrieve       
	ExtractIntegrationID                            string `json:"extractIntegrationId"`
}

// Blueprint properties
type ExtractIntegrationFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the Bot to use                                               
	BotID                                              string                 `json:"botId"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// Optional webhook to receive the extracted data                         
	Request                                            *string                `json:"request,omitempty"`
	// The configured extraction schema                                       
	Schema                                             map[string]interface{} `json:"schema,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type ExtractIntegrationUpdateParams struct {
	// The ID of the Extract integration       
	ExtractIntegrationID                string `json:"extractIntegrationId"`
}

// Blueprint properties
type ExtractIntegrationUpdateRequest struct {
	// The ID of the blueprint                                              
	BlueprintID                                      *string                `json:"blueprintId,omitempty"`
	// The ID of the Bot to use                                             
	BotID                                            *string                `json:"botId,omitempty"`
	// The associated description                                           
	Description                                      *string                `json:"description,omitempty"`
	// Meta data information                                                
	Meta                                             map[string]interface{} `json:"meta,omitempty"`
	// The language model to use for data extraction                        
	Model                                            *string                `json:"model,omitempty"`
	// The associated name                                                  
	Name                                             *string                `json:"name,omitempty"`
	// Optional webhook to receive the extracted data                       
	Request                                          *string                `json:"request,omitempty"`
	// The configured extraction schema                                     
	Schema                                           map[string]interface{} `json:"schema,omitempty"`
}

type ExtractIntegrationUpdateResponse struct {
	// The ID of the Extract Integration       
	ID                                  string `json:"id"`
}

// Blueprint properties
type ExtractIntegrationCreateRequest struct {
	// The ID of the blueprint                                              
	BlueprintID                                      *string                `json:"blueprintId,omitempty"`
	// The ID of the Bot to use                                             
	BotID                                            *string                `json:"botId,omitempty"`
	// The associated description                                           
	Description                                      *string                `json:"description,omitempty"`
	// Meta data information                                                
	Meta                                             map[string]interface{} `json:"meta,omitempty"`
	// The language model to use for data extraction                        
	Model                                            *string                `json:"model,omitempty"`
	// The associated name                                                  
	Name                                             *string                `json:"name,omitempty"`
	// Optional webhook to receive the extracted data                       
	Request                                          *string                `json:"request,omitempty"`
	// The configured extraction schema                                     
	Schema                                           map[string]interface{} `json:"schema,omitempty"`
}

type ExtractIntegrationCreateResponse struct {
	// The ID of the Extract Integration       
	ID                                  string `json:"id"`
}

type ExtractIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ExtractIntegrationsListResponse struct {
	Items []ExtractIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type ExtractIntegrationsListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the Bot to use                                               
	BotID                                              string                 `json:"botId"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// Optional webhook to receive the extracted data                         
	Request                                            *string                `json:"request,omitempty"`
	// The configured extraction schema                                       
	Schema                                             map[string]interface{} `json:"schema,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type InstagramIntegrationDeleteParams struct {
	// The ID of the Instagram integration       
	InstagramIntegrationID                string `json:"instagramIntegrationId"`
}

type InstagramIntegrationDeleteResponse struct {
	// The ID of the deleted Instagram integration       
	ID                                            string `json:"id"`
}

type InstagramIntegrationFetchParams struct {
	// The ID of the Instagram integration to retrieve       
	InstagramIntegrationID                            string `json:"instagramIntegrationId"`
}

// Blueprint properties
type InstagramIntegrationFetchResponse struct {
	// The Instagram integration access token (returned as '********' if configured, null                       
	// otherwise)                                                                                               
	AccessToken                                                                          *string                `json:"accessToken,omitempty"`
	// Whether the bot supports attachments                                                                     
	Attachments                                                                          *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                  
	BlueprintID                                                                          *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                            
	BotID                                                                                *string                `json:"botId,omitempty"`
	// Whether to collect contacts                                                                              
	ContactCollection                                                                    *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                                         
	CreatedAt                                                                            float64                `json:"createdAt"`
	// The associated description                                                                               
	Description                                                                          *string                `json:"description,omitempty"`
	// The instance ID                                                                                          
	ID                                                                                   string                 `json:"id"`
	// Meta data information                                                                                    
	Meta                                                                                 map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                      
	Name                                                                                 *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                                                   
	SessionDuration                                                                      *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                                                         
	UpdatedAt                                                                            float64                `json:"updatedAt"`
	// The Instagram integration verify token                                                                   
	VerifyToken                                                                          string                 `json:"verifyToken"`
}

type InstagramIntegrationSetupParams struct {
	// The ID of the Instagram integration       
	InstagramIntegrationID                string `json:"instagramIntegrationId"`
}

type InstagramIntegrationSetupResponse struct {
	// The ID of the Instagram Integration       
	ID                                    string `json:"id"`
}

type InstagramIntegrationUpdateParams struct {
	// The ID of the Instagram integration       
	InstagramIntegrationID                string `json:"instagramIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type InstagramIntegrationUpdateRequest struct {
	// The Instagram integration access token                              
	AccessToken                                     *string                `json:"accessToken,omitempty"`
	// Whether the bot supports attachments                                
	Attachments                                     *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// Whether to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type InstagramIntegrationUpdateResponse struct {
	// The ID of the Instagram Integration       
	ID                                    string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type InstagramIntegrationCreateRequest struct {
	// The Instagram integration access token                              
	AccessToken                                     *string                `json:"accessToken,omitempty"`
	// Whether the bot supports attachments                                
	Attachments                                     *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// Whether to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type InstagramIntegrationCreateResponse struct {
	// The ID of the Instagram Integration       
	ID                                    string `json:"id"`
}

type InstagramIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type InstagramIntegrationsListResponse struct {
	Items []InstagramIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type InstagramIntegrationsListResponseItem struct {
	// The Instagram integration access token (returned as '********' if configured, null                       
	// otherwise)                                                                                               
	AccessToken                                                                          *string                `json:"accessToken,omitempty"`
	// Whether the bot supports attachments                                                                     
	Attachments                                                                          *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                  
	BlueprintID                                                                          *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                            
	BotID                                                                                *string                `json:"botId,omitempty"`
	// Whether to collect contacts                                                                              
	ContactCollection                                                                    *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                                         
	CreatedAt                                                                            float64                `json:"createdAt"`
	// The associated description                                                                               
	Description                                                                          *string                `json:"description,omitempty"`
	// The instance ID                                                                                          
	ID                                                                                   string                 `json:"id"`
	// Meta data information                                                                                    
	Meta                                                                                 map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                      
	Name                                                                                 *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                                                   
	SessionDuration                                                                      *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                                                         
	UpdatedAt                                                                            float64                `json:"updatedAt"`
	// The Instagram integration verify token                                                                   
	VerifyToken                                                                          string                 `json:"verifyToken"`
}

type MCPServerIntegrationDeleteParams struct {
	// The ID of the McpServer integration       
	McpserverIntegrationID                string `json:"mcpserverIntegrationId"`
}

type MCPServerIntegrationDeleteResponse struct {
	// The ID of the deleted McpServer integration       
	ID                                            string `json:"id"`
}

type MCPServerIntegrationFetchParams struct {
	// The ID of the McpServer integration to retrieve       
	McpserverIntegrationID                            string `json:"mcpserverIntegrationId"`
}

// Blueprint properties
type MCPServerIntegrationFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The ID of the skillset                                                 
	SkillsetID                                         *string                `json:"skillsetId,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type MCPServerIntegrationUpdateParams struct {
	// The ID of the McpServer integration       
	McpserverIntegrationID                string `json:"mcpserverIntegrationId"`
}

// Blueprint properties
type MCPServerIntegrationUpdateRequest struct {
	// The ID of the blueprint                          
	BlueprintID                  *string                `json:"blueprintId,omitempty"`
	// The associated description                       
	Description                  *string                `json:"description,omitempty"`
	// Meta data information                            
	Meta                         map[string]interface{} `json:"meta,omitempty"`
	// The associated name                              
	Name                         *string                `json:"name,omitempty"`
	// The ID of the skillset                           
	SkillsetID                   *string                `json:"skillsetId,omitempty"`
}

type MCPServerIntegrationUpdateResponse struct {
	// The ID of the McpServer Integration       
	ID                                    string `json:"id"`
}

// Blueprint properties
type MCPServerIntegrationCreateRequest struct {
	// The ID of the blueprint                          
	BlueprintID                  *string                `json:"blueprintId,omitempty"`
	// The associated description                       
	Description                  *string                `json:"description,omitempty"`
	// Meta data information                            
	Meta                         map[string]interface{} `json:"meta,omitempty"`
	// The associated name                              
	Name                         *string                `json:"name,omitempty"`
	// The ID of the skillset                           
	SkillsetID                   *string                `json:"skillsetId,omitempty"`
}

type MCPServerIntegrationCreateResponse struct {
	// The ID of the McpServer Integration       
	ID                                    string `json:"id"`
}

type MCPServerIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type MCPServerIntegrationsListResponse struct {
	Items []MCPServerIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type MCPServerIntegrationsListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The ID of the skillset                                                 
	SkillsetID                                         *string                `json:"skillsetId,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type MessengerIntegrationDeleteParams struct {
	// The ID of the Messenger integration       
	MessengerIntegrationID                string `json:"messengerIntegrationId"`
}

type MessengerIntegrationDeleteResponse struct {
	// The ID of the deleted Messenger integration       
	ID                                            string `json:"id"`
}

type MessengerIntegrationFetchParams struct {
	// The ID of the Messenger integration to retrieve       
	MessengerIntegrationID                            string `json:"messengerIntegrationId"`
}

// Blueprint properties
type MessengerIntegrationFetchResponse struct {
	// The Messenger integration access token (returned as '********' if configured, null                       
	// otherwise)                                                                                               
	AccessToken                                                                          *string                `json:"accessToken,omitempty"`
	// Weather the bot supports attachments                                                                     
	Attachments                                                                          *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                  
	BlueprintID                                                                          *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                            
	BotID                                                                                *string                `json:"botId,omitempty"`
	// The timestamp (ms) when the instance was created                                                         
	CreatedAt                                                                            float64                `json:"createdAt"`
	// The associated description                                                                               
	Description                                                                          *string                `json:"description,omitempty"`
	// The instance ID                                                                                          
	ID                                                                                   string                 `json:"id"`
	// Meta data information                                                                                    
	Meta                                                                                 map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                      
	Name                                                                                 *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                                                   
	SessionDuration                                                                      *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                                                         
	UpdatedAt                                                                            float64                `json:"updatedAt"`
	// The Messenger integration verify token                                                                   
	VerifyToken                                                                          string                 `json:"verifyToken"`
}

type MessengerIntegrationSetupParams struct {
	// The ID of the Messenger integration       
	MessengerIntegrationID                string `json:"messengerIntegrationId"`
}

type MessengerIntegrationSetupResponse struct {
	// The ID of the Messenger Integration       
	ID                                    string `json:"id"`
}

type MessengerIntegrationUpdateParams struct {
	// The ID of the Messenger integration       
	MessengerIntegrationID                string `json:"messengerIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type MessengerIntegrationUpdateRequest struct {
	// The Messenger integration access token                              
	AccessToken                                     *string                `json:"accessToken,omitempty"`
	// Weather the bot supports attachments                                
	Attachments                                     *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type MessengerIntegrationUpdateResponse struct {
	// The ID of the Messenger Integration       
	ID                                    string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type MessengerIntegrationCreateRequest struct {
	// The Messenger integration access token                              
	AccessToken                                     *string                `json:"accessToken,omitempty"`
	// Weather the bot supports attachments                                
	Attachments                                     *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type MessengerIntegrationCreateResponse struct {
	// The ID of the Messenger Integration       
	ID                                    string `json:"id"`
}

type MessengerIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type MessengerIntegrationsListResponse struct {
	Items []MessengerIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type MessengerIntegrationsListResponseItem struct {
	// The Messenger integration access token (returned as '********' if configured, null                       
	// otherwise)                                                                                               
	AccessToken                                                                          *string                `json:"accessToken,omitempty"`
	// Weather the bot supports attachments                                                                     
	Attachments                                                                          *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                  
	BlueprintID                                                                          *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                            
	BotID                                                                                *string                `json:"botId,omitempty"`
	// The timestamp (ms) when the instance was created                                                         
	CreatedAt                                                                            float64                `json:"createdAt"`
	// The associated description                                                                               
	Description                                                                          *string                `json:"description,omitempty"`
	// The instance ID                                                                                          
	ID                                                                                   string                 `json:"id"`
	// Meta data information                                                                                    
	Meta                                                                                 map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                      
	Name                                                                                 *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                                                   
	SessionDuration                                                                      *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                                                         
	UpdatedAt                                                                            float64                `json:"updatedAt"`
	// The Messenger integration verify token                                                                   
	VerifyToken                                                                          string                 `json:"verifyToken"`
}

type NotionIntegrationDeleteParams struct {
	// The ID of the Notion integration       
	NotionIntegrationID                string `json:"notionIntegrationId"`
}

type NotionIntegrationDeleteResponse struct {
	// The ID of the deleted Notion integration       
	ID                                         string `json:"id"`
}

type NotionIntegrationFetchParams struct {
	// The ID of the Notion integration to retrieve       
	NotionIntegrationID                            string `json:"notionIntegrationId"`
}

// Blueprint properties
type NotionIntegrationFetchResponse struct {
	// The ID of the blueprint                                                                           
	BlueprintID                                                                   *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                                                  
	CreatedAt                                                                     float64                `json:"createdAt"`
	// The ID of the dataset to sync into                                                                
	DatasetID                                                                     string                 `json:"datasetId"`
	// The associated description                                                                        
	Description                                                                   *string                `json:"description,omitempty"`
	// The time in milliseconds until records expire                                                     
	ExpiresIn                                                                     *float64               `json:"expiresIn,omitempty"`
	// The instance ID                                                                                   
	ID                                                                            string                 `json:"id"`
	// The timestamp of the last successful sync                                                         
	LastSyncedAt                                                                  *time.Time             `json:"lastSyncedAt,omitempty"`
	// Meta data information                                                                             
	Meta                                                                          map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                               
	Name                                                                          *string                `json:"name,omitempty"`
	// The sync schedule                                                                                 
	SyncSchedule                                                                  *string                `json:"syncSchedule,omitempty"`
	// The sync status of an integration                                                                 
	SyncStatus                                                                    *SyncStatus            `json:"syncStatus,omitempty"`
	// The Notion API token (returned as '********' if configured, null otherwise)                       
	Token                                                                         *string                `json:"token,omitempty"`
	// The timestamp (ms) when the instance was updated                                                  
	UpdatedAt                                                                     float64                `json:"updatedAt"`
}

type NotionIntegrationSyncParams struct {
	// The ID of the Notion integration       
	NotionIntegrationID                string `json:"notionIntegrationId"`
}

type NotionIntegrationSyncResponse struct {
	// The ID of the synced Notion integration       
	ID                                        string `json:"id"`
}

type NotionIntegrationUpdateParams struct {
	// The ID of the Notion integration       
	NotionIntegrationID                string `json:"notionIntegrationId"`
}

// Blueprint properties
type NotionIntegrationUpdateRequest struct {
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the dataset to sync into                                  
	DatasetID                                       *string                `json:"datasetId,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// The time in milliseconds until records expire                       
	ExpiresIn                                       *float64               `json:"expiresIn,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The sync schedule                                                   
	SyncSchedule                                    *string                `json:"syncSchedule,omitempty"`
	// The Notion API token                                                
	Token                                           *string                `json:"token,omitempty"`
}

type NotionIntegrationUpdateResponse struct {
	// The ID of the Notion Integration       
	ID                                 string `json:"id"`
}

// Blueprint properties
type NotionIntegrationCreateRequest struct {
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the dataset to sync into                                  
	DatasetID                                       *string                `json:"datasetId,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// The time in milliseconds until records expire                       
	ExpiresIn                                       *float64               `json:"expiresIn,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The sync schedule                                                   
	SyncSchedule                                    *string                `json:"syncSchedule,omitempty"`
	// The Notion API token                                                
	Token                                           *string                `json:"token,omitempty"`
}

type NotionIntegrationCreateResponse struct {
	// The ID of the Notion Integration       
	ID                                 string `json:"id"`
}

type NotionIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type NotionIntegrationsListResponse struct {
	Items []NotionIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type NotionIntegrationsListResponseItem struct {
	// The ID of the blueprint                                                                           
	BlueprintID                                                                   *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                                                  
	CreatedAt                                                                     float64                `json:"createdAt"`
	// The ID of the dataset to sync into                                                                
	DatasetID                                                                     string                 `json:"datasetId"`
	// The associated description                                                                        
	Description                                                                   *string                `json:"description,omitempty"`
	// The time in milliseconds until records expire                                                     
	ExpiresIn                                                                     *float64               `json:"expiresIn,omitempty"`
	// The instance ID                                                                                   
	ID                                                                            string                 `json:"id"`
	// The timestamp of the last successful sync                                                         
	LastSyncedAt                                                                  *time.Time             `json:"lastSyncedAt,omitempty"`
	// Meta data information                                                                             
	Meta                                                                          map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                               
	Name                                                                          *string                `json:"name,omitempty"`
	// The sync schedule                                                                                 
	SyncSchedule                                                                  *string                `json:"syncSchedule,omitempty"`
	// The sync status of an integration                                                                 
	SyncStatus                                                                    *SyncStatus            `json:"syncStatus,omitempty"`
	// The Notion API token (returned as '********' if configured, null otherwise)                       
	Token                                                                         *string                `json:"token,omitempty"`
	// The timestamp (ms) when the instance was updated                                                  
	UpdatedAt                                                                     float64                `json:"updatedAt"`
}

type SitemapIntegrationDeleteParams struct {
	// The ID of the Sitemap integration       
	SitemapIntegrationID                string `json:"sitemapIntegrationId"`
}

type SitemapIntegrationDeleteResponse struct {
	// The ID of the deleted Sitemap integration       
	ID                                          string `json:"id"`
}

type SitemapIntegrationFetchParams struct {
	// The ID of the Sitemap integration to retrieve       
	SitemapIntegrationID                            string `json:"sitemapIntegrationId"`
}

// Blueprint properties
type SitemapIntegrationFetchResponse struct {
	// The ID of the blueprint                                                                                       
	BlueprintID                                                                               *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                                                              
	CreatedAt                                                                                 float64                `json:"createdAt"`
	// The ID of the dataset used in the Sitemap integration                                                         
	DatasetID                                                                                 string                 `json:"datasetId"`
	// The associated description                                                                                    
	Description                                                                               *string                `json:"description,omitempty"`
	// Record expiry in milliseconds                                                                                 
	ExpiresIn                                                                                 *float64               `json:"expiresIn,omitempty"`
	// The glob rules to use for this Sitemap integration                                                            
	Glob                                                                                      *string                `json:"glob,omitempty"`
	// The instance ID                                                                                               
	ID                                                                                        string                 `json:"id"`
	// Indicates if the Sitemap integration should use JavaScript during the spidering process                       
	Javascript                                                                                *bool                  `json:"javascript,omitempty"`
	// The timestamp of the last successful sync                                                                     
	LastSyncedAt                                                                              *time.Time             `json:"lastSyncedAt,omitempty"`
	// Meta data information                                                                                         
	Meta                                                                                      map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                           
	Name                                                                                      *string                `json:"name,omitempty"`
	// The selector rules to use for this Sitemap integration                                                        
	Selectors                                                                                 *string                `json:"selectors,omitempty"`
	// The sync schedule to use for this Sitemap integration                                                         
	SyncSchedule                                                                              *string                `json:"syncSchedule,omitempty"`
	// The sync status of an integration                                                                             
	SyncStatus                                                                                *SyncStatus            `json:"syncStatus,omitempty"`
	// The timestamp (ms) when the instance was updated                                                              
	UpdatedAt                                                                                 float64                `json:"updatedAt"`
	// The URL to use for this Sitemap integration                                                                   
	URL                                                                                       *string                `json:"url,omitempty"`
}

type SitemapIntegrationSyncParams struct {
	// The ID of the Sitemap integration       
	SitemapIntegrationID                string `json:"sitemapIntegrationId"`
}

type SitemapIntegrationSyncResponse struct {
	// The ID of the Sitemap Integration       
	ID                                  string `json:"id"`
}

type SitemapIntegrationUpdateParams struct {
	// The ID of the Sitemap integration       
	SitemapIntegrationID                string `json:"sitemapIntegrationId"`
}

// Blueprint properties
type SitemapIntegrationUpdateRequest struct {
	// The ID of the blueprint                                                                                       
	BlueprintID                                                                               *string                `json:"blueprintId,omitempty"`
	// The ID of the dataset to use for this Sitemap integration                                                     
	DatasetID                                                                                 *string                `json:"datasetId,omitempty"`
	// The associated description                                                                                    
	Description                                                                               *string                `json:"description,omitempty"`
	// Record expiry in milliseconds                                                                                 
	ExpiresIn                                                                                 *float64               `json:"expiresIn,omitempty"`
	// The glob rules to use for this Sitemap integration                                                            
	Glob                                                                                      *string                `json:"glob,omitempty"`
	// Indicates if the Sitemap integration should use JavaScript during the spidering process                       
	Javascript                                                                                *bool                  `json:"javascript,omitempty"`
	// Meta data information                                                                                         
	Meta                                                                                      map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                           
	Name                                                                                      *string                `json:"name,omitempty"`
	// The selector rules to use for this Sitemap integration                                                        
	Selectors                                                                                 *string                `json:"selectors,omitempty"`
	// The sync schedule to use for this Sitemap integration                                                         
	SyncSchedule                                                                              *string                `json:"syncSchedule,omitempty"`
	// The URL to use for this Sitemap integration                                                                   
	URL                                                                                       *string                `json:"url,omitempty"`
}

type SitemapIntegrationUpdateResponse struct {
	// The ID of the Sitemap Integration       
	ID                                  string `json:"id"`
}

// Blueprint properties
type SitemapIntegrationCreateRequest struct {
	// The ID of the blueprint                                                                                       
	BlueprintID                                                                               *string                `json:"blueprintId,omitempty"`
	// The ID of the dataset to use for this Sitemap integration                                                     
	DatasetID                                                                                 *string                `json:"datasetId,omitempty"`
	// The associated description                                                                                    
	Description                                                                               *string                `json:"description,omitempty"`
	// Record expiry in milliseconds                                                                                 
	ExpiresIn                                                                                 *float64               `json:"expiresIn,omitempty"`
	// The glob rules to use for this Sitemap integration                                                            
	Glob                                                                                      *string                `json:"glob,omitempty"`
	// Indicates if the Sitemap integration should use JavaScript during the spidering process                       
	Javascript                                                                                *bool                  `json:"javascript,omitempty"`
	// Meta data information                                                                                         
	Meta                                                                                      map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                           
	Name                                                                                      *string                `json:"name,omitempty"`
	// The selector rules to use for this Sitemap integration                                                        
	Selectors                                                                                 *string                `json:"selectors,omitempty"`
	// The sync schedule to use for this Sitemap integration                                                         
	SyncSchedule                                                                              *string                `json:"syncSchedule,omitempty"`
	// The URL to use for this Sitemap integration                                                                   
	URL                                                                                       *string                `json:"url,omitempty"`
}

type SitemapIntegrationCreateResponse struct {
	// The ID of the Sitemap Integration       
	ID                                  string `json:"id"`
}

type SitemapIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type SitemapIntegrationsListResponse struct {
	Items []SitemapIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type SitemapIntegrationsListResponseItem struct {
	// The ID of the blueprint                                                                                       
	BlueprintID                                                                               *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                                                              
	CreatedAt                                                                                 float64                `json:"createdAt"`
	// The ID of the dataset used in the Sitemap integration                                                         
	DatasetID                                                                                 string                 `json:"datasetId"`
	// The associated description                                                                                    
	Description                                                                               *string                `json:"description,omitempty"`
	// Record expiry in milliseconds                                                                                 
	ExpiresIn                                                                                 *float64               `json:"expiresIn,omitempty"`
	// The glob rules to use for this Sitemap integration                                                            
	Glob                                                                                      *string                `json:"glob,omitempty"`
	// The instance ID                                                                                               
	ID                                                                                        string                 `json:"id"`
	// Indicates if the Sitemap integration should use JavaScript during the spidering process                       
	Javascript                                                                                *bool                  `json:"javascript,omitempty"`
	// The timestamp of the last successful sync                                                                     
	LastSyncedAt                                                                              *time.Time             `json:"lastSyncedAt,omitempty"`
	// Meta data information                                                                                         
	Meta                                                                                      map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                           
	Name                                                                                      *string                `json:"name,omitempty"`
	// The selector rules to use for this Sitemap integration                                                        
	Selectors                                                                                 *string                `json:"selectors,omitempty"`
	// The sync schedule to use for this Sitemap integration                                                         
	SyncSchedule                                                                              *string                `json:"syncSchedule,omitempty"`
	// The sync status of an integration                                                                             
	SyncStatus                                                                                *SyncStatus            `json:"syncStatus,omitempty"`
	// The timestamp (ms) when the instance was updated                                                              
	UpdatedAt                                                                                 float64                `json:"updatedAt"`
	// The URL to use for this Sitemap integration                                                                   
	URL                                                                                       *string                `json:"url,omitempty"`
}

type SlackIntegrationDeleteParams struct {
	// The ID of the Slack integration       
	SlackIntegrationID                string `json:"slackIntegrationId"`
}

type SlackIntegrationDeleteResponse struct {
	// The ID of the deleted Slack integration       
	ID                                        string `json:"id"`
}

type SlackIntegrationFetchParams struct {
	// The ID of the Slack integration to retrieve       
	SlackIntegrationID                            string `json:"slackIntegrationId"`
}

// Blueprint properties
type SlackIntegrationFetchResponse struct {
	// Configure automatic response behavior. Use '@all' to respond to all messages, '@agent                         
	// <instructions>' for agent-powered decisions, or custom instructions for lightweight LLM                       
	// filtering. Null/empty defaults to current behavior (DMs, mentions, threads only).                             
	AutoRespond                                                                               *string                `json:"autoRespond,omitempty"`
	// The ID of the blueprint                                                                                       
	BlueprintID                                                                               *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                                 
	BotID                                                                                     *string                `json:"botId,omitempty"`
	// The bot token (returned as '********' if configured, null otherwise)                                          
	BotToken                                                                                  *string                `json:"botToken,omitempty"`
	// Weather to collect contacts                                                                                   
	ContactCollection                                                                         *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                                              
	CreatedAt                                                                                 float64                `json:"createdAt"`
	// The associated description                                                                                    
	Description                                                                               *string                `json:"description,omitempty"`
	// The instance ID                                                                                               
	ID                                                                                        string                 `json:"id"`
	// Meta data information                                                                                         
	Meta                                                                                      map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                           
	Name                                                                                      *string                `json:"name,omitempty"`
	// Whether to enable ratings buttons feature                                                                     
	Ratings                                                                                   *bool                  `json:"ratings,omitempty"`
	// Whether to enable references feature                                                                          
	References                                                                                *bool                  `json:"references,omitempty"`
	// The session duration for the Slack integration                                                                
	SessionDuration                                                                           *float64               `json:"sessionDuration,omitempty"`
	// The signing secret (returned as '********' if configured, null otherwise)                                     
	SigningSecret                                                                             *string                `json:"signingSecret,omitempty"`
	// The timestamp (ms) when the instance was updated                                                              
	UpdatedAt                                                                                 float64                `json:"updatedAt"`
	// The user token (returned as '********' if configured, null otherwise)                                         
	UserToken                                                                                 *string                `json:"userToken,omitempty"`
	// The number of visible messages outside of the new thread                                                      
	VisibleMessages                                                                           *float64               `json:"visibleMessages,omitempty"`
}

type SlackIntegrationSetupParams struct {
	// The ID of the Slack integration       
	SlackIntegrationID                string `json:"slackIntegrationId"`
}

type SlackIntegrationSetupResponse struct {
	// The ID of the setup Slack integration       
	ID                                      string `json:"id"`
}

type SlackIntegrationUpdateParams struct {
	// The ID of the Slack integration       
	SlackIntegrationID                string `json:"slackIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type SlackIntegrationUpdateRequest struct {
	// Configure automatic response behavior. Use '@all' to respond to all messages, '@agent                         
	// <instructions>' for agent-powered decisions, or custom instructions for lightweight LLM                       
	// filtering. Null/empty defaults to current behavior (DMs, mentions, threads only).                             
	AutoRespond                                                                               *string                `json:"autoRespond,omitempty"`
	// The ID of the blueprint                                                                                       
	BlueprintID                                                                               *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                                 
	BotID                                                                                     *string                `json:"botId,omitempty"`
	// The bot token for the Slack integration                                                                       
	BotToken                                                                                  *string                `json:"botToken,omitempty"`
	// Weather to collect contacts                                                                                   
	ContactCollection                                                                         *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                                                                    
	Description                                                                               *string                `json:"description,omitempty"`
	// Meta data information                                                                                         
	Meta                                                                                      map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                           
	Name                                                                                      *string                `json:"name,omitempty"`
	// Whether to enable ratings buttons feature                                                                     
	Ratings                                                                                   *bool                  `json:"ratings,omitempty"`
	// Whether to enable references feature                                                                          
	References                                                                                *bool                  `json:"references,omitempty"`
	// The session duration for the Slack integration                                                                
	SessionDuration                                                                           *float64               `json:"sessionDuration,omitempty"`
	// The signing secret for the Slack integration                                                                  
	SigningSecret                                                                             *string                `json:"signingSecret,omitempty"`
	// The user token for the Slack integration                                                                      
	UserToken                                                                                 *string                `json:"userToken,omitempty"`
	// The number of visible messages outside of the new thread                                                      
	VisibleMessages                                                                           *float64               `json:"visibleMessages,omitempty"`
}

type SlackIntegrationUpdateResponse struct {
	// The ID of the Slack Integration       
	ID                                string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type SlackIntegrationCreateRequest struct {
	// Configure automatic response behavior. Use '@all' to respond to all messages, '@agent                         
	// <instructions>' for agent-powered decisions, or custom instructions for lightweight LLM                       
	// filtering. Null/empty defaults to current behavior (DMs, mentions, threads only).                             
	AutoRespond                                                                               *string                `json:"autoRespond,omitempty"`
	// The ID of the blueprint                                                                                       
	BlueprintID                                                                               *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                                 
	BotID                                                                                     *string                `json:"botId,omitempty"`
	// The bot token for the Slack integration                                                                       
	BotToken                                                                                  *string                `json:"botToken,omitempty"`
	// Weather to collect contacts                                                                                   
	ContactCollection                                                                         *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                                                                    
	Description                                                                               *string                `json:"description,omitempty"`
	// Meta data information                                                                                         
	Meta                                                                                      map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                           
	Name                                                                                      *string                `json:"name,omitempty"`
	// Whether to enable ratings buttons feature                                                                     
	Ratings                                                                                   *bool                  `json:"ratings,omitempty"`
	// Whether to enable references feature                                                                          
	References                                                                                *bool                  `json:"references,omitempty"`
	// The session duration for the Slack integration                                                                
	SessionDuration                                                                           *float64               `json:"sessionDuration,omitempty"`
	// The signing secret for the Slack integration                                                                  
	SigningSecret                                                                             *string                `json:"signingSecret,omitempty"`
	// The user token for the Slack integration                                                                      
	UserToken                                                                                 *string                `json:"userToken,omitempty"`
	// The number of visible messages outside of the new thread                                                      
	VisibleMessages                                                                           *float64               `json:"visibleMessages,omitempty"`
}

type SlackIntegrationCreateResponse struct {
	// The ID of the Slack Integration       
	ID                                string `json:"id"`
}

type SlackIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type SlackIntegrationsListResponse struct {
	Items []SlackIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type SlackIntegrationsListResponseItem struct {
	// Configure automatic response behavior. Use '@all' to respond to all messages, '@agent                         
	// <instructions>' for agent-powered decisions, or custom instructions for lightweight LLM                       
	// filtering. Null/empty defaults to current behavior (DMs, mentions, threads only).                             
	AutoRespond                                                                               *string                `json:"autoRespond,omitempty"`
	// The ID of the blueprint                                                                                       
	BlueprintID                                                                               *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                                 
	BotID                                                                                     *string                `json:"botId,omitempty"`
	// The bot token (returned as '********' if configured, null otherwise)                                          
	BotToken                                                                                  *string                `json:"botToken,omitempty"`
	// Weather to collect contacts                                                                                   
	ContactCollection                                                                         *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                                              
	CreatedAt                                                                                 float64                `json:"createdAt"`
	// The associated description                                                                                    
	Description                                                                               *string                `json:"description,omitempty"`
	// The instance ID                                                                                               
	ID                                                                                        string                 `json:"id"`
	// Meta data information                                                                                         
	Meta                                                                                      map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                           
	Name                                                                                      *string                `json:"name,omitempty"`
	// Whether to enable ratings buttons feature                                                                     
	Ratings                                                                                   *bool                  `json:"ratings,omitempty"`
	// Whether to enable references feature                                                                          
	References                                                                                *bool                  `json:"references,omitempty"`
	// The session duration for the Slack integration                                                                
	SessionDuration                                                                           *float64               `json:"sessionDuration,omitempty"`
	// The signing secret (returned as '********' if configured, null otherwise)                                     
	SigningSecret                                                                             *string                `json:"signingSecret,omitempty"`
	// The timestamp (ms) when the instance was updated                                                              
	UpdatedAt                                                                                 float64                `json:"updatedAt"`
	// The user token (returned as '********' if configured, null otherwise)                                         
	UserToken                                                                                 *string                `json:"userToken,omitempty"`
	// The number of visible messages outside of the new thread                                                      
	VisibleMessages                                                                           *float64               `json:"visibleMessages,omitempty"`
}

type SupportIntegrationDeleteParams struct {
	// The ID of the Support integration       
	SupportIntegrationID                string `json:"supportIntegrationId"`
}

type SupportIntegrationDeleteResponse struct {
	// The ID of the deleted Support integration       
	ID                                          string `json:"id"`
}

type SupportIntegrationFetchParams struct {
	// The ID of the Support integration to retrieve       
	SupportIntegrationID                            string `json:"supportIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type SupportIntegrationFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                          
	BotID                                              string                 `json:"botId"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The email to use                                                       
	Email                                              *string                `json:"email,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type SupportIntegrationUpdateParams struct {
	// The ID of the Support integration       
	SupportIntegrationID                string `json:"supportIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type SupportIntegrationUpdateRequest struct {
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// The email to use                                                    
	Email                                           *string                `json:"email,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
}

type SupportIntegrationUpdateResponse struct {
	// The ID of the Support Integration       
	ID                                  string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type SupportIntegrationCreateRequest struct {
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// The email to use                                                    
	Email                                           *string                `json:"email,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
}

type SupportIntegrationCreateResponse struct {
	// The ID of the Support Integration       
	ID                                  string `json:"id"`
}

type SupportIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type SupportIntegrationsListResponse struct {
	Items []SupportIntegrationsListResponseItem `json:"items"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type SupportIntegrationsListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                          
	BotID                                              string                 `json:"botId"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The email to use                                                       
	Email                                              *string                `json:"email,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type TelegramIntegrationDeleteParams struct {
	// The ID of the Telegram integration       
	TelegramIntegrationID                string `json:"telegramIntegrationId"`
}

type TelegramIntegrationDeleteResponse struct {
	// The ID of the deleted Telegram integration       
	ID                                           string `json:"id"`
}

type TelegramIntegrationFetchParams struct {
	// The ID of the Telegram integration to retrieve       
	TelegramIntegrationID                            string `json:"telegramIntegrationId"`
}

// Blueprint properties
type TelegramIntegrationFetchResponse struct {
	// Weather the bot supports attachments                                   
	Attachments                                        *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                          
	BotID                                              *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                            
	ContactCollection                                  *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                 
	SessionDuration                                    *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type TelegramIntegrationSetupParams struct {
	// The ID of the Telegram integration       
	TelegramIntegrationID                string `json:"telegramIntegrationId"`
}

type TelegramIntegrationSetupResponse struct {
	// The ID of the Telegram Integration       
	ID                                   string `json:"id"`
}

type TelegramIntegrationUpdateParams struct {
	// The ID of the Telegram integration       
	TelegramIntegrationID                string `json:"telegramIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type TelegramIntegrationUpdateRequest struct {
	// Weather the bot supports attachments                                
	Attachments                                     *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// The Telegram integration bot token                                  
	BotToken                                        *string                `json:"botToken,omitempty"`
	// Weather to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type TelegramIntegrationUpdateResponse struct {
	// The ID of the Telegram Integration       
	ID                                   string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type TelegramIntegrationCreateRequest struct {
	// Weather the bot supports attachments                                
	Attachments                                     *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// The Telegram integration bot token                                  
	BotToken                                        *string                `json:"botToken,omitempty"`
	// Weather to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type TelegramIntegrationCreateResponse struct {
	// The ID of the Telegram Integration       
	ID                                   string `json:"id"`
}

type TelegramIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type TelegramIntegrationsListResponse struct {
	Items []TelegramIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type TelegramIntegrationsListResponseItem struct {
	// Weather the bot supports attachments                                   
	Attachments                                        *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                          
	BotID                                              *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                            
	ContactCollection                                  *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                 
	SessionDuration                                    *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type TriggerIntegrationDeleteParams struct {
	// The ID of the Trigger integration       
	TriggerIntegrationID                string `json:"triggerIntegrationId"`
}

type TriggerIntegrationDeleteResponse struct {
	// The ID of the deleted Trigger integration       
	ID                                          string `json:"id"`
}

type TriggerIntegrationFetchParams struct {
	// The ID of the Trigger integration to retrieve       
	TriggerIntegrationID                            string `json:"triggerIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type TriggerIntegrationFetchResponse struct {
	// When enabled the integration requires authentication                       
	Authenticate                                           *bool                  `json:"authenticate,omitempty"`
	// The ID of the blueprint                                                    
	BlueprintID                                            *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                              
	BotID                                                  *string                `json:"botId,omitempty"`
	// The timestamp (ms) when the instance was created                           
	CreatedAt                                              float64                `json:"createdAt"`
	// The associated description                                                 
	Description                                            *string                `json:"description,omitempty"`
	// The instance ID                                                            
	ID                                                     string                 `json:"id"`
	// Meta data information                                                      
	Meta                                                   map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                        
	Name                                                   *string                `json:"name,omitempty"`
	// The Trigger integration secret                                             
	Secret                                                 string                 `json:"secret"`
	// The session duration (in milliseconds)                                     
	SessionDuration                                        *float64               `json:"sessionDuration,omitempty"`
	// The schedule                                                               
	TriggerSchedule                                        *Schedule              `json:"triggerSchedule,omitempty"`
	// The timestamp (ms) when the instance was updated                           
	UpdatedAt                                              float64                `json:"updatedAt"`
}

type TriggerIntegrationInvokeParams struct {
	// The ID of the Trigger integration       
	TriggerIntegrationID                string `json:"triggerIntegrationId"`
}

type TriggerIntegrationInvokeResponse struct {
	// The ID of the trigged Trigger integration       
	ID                                          string `json:"id"`
}

type TriggerIntegrationSetupParams struct {
	// The ID of the Trigger integration       
	TriggerIntegrationID                string `json:"triggerIntegrationId"`
}

type TriggerIntegrationSetupResponse struct {
	// The ID of the Trigger Integration       
	ID                                  string `json:"id"`
}

type TriggerIntegrationUpdateParams struct {
	// The ID of the Trigger integration       
	TriggerIntegrationID                string `json:"triggerIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type TriggerIntegrationUpdateRequest struct {
	// When enabled the integration requires authentication                       
	Authenticate                                           *bool                  `json:"authenticate,omitempty"`
	// The ID of the blueprint                                                    
	BlueprintID                                            *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                              
	BotID                                                  *string                `json:"botId,omitempty"`
	// The associated description                                                 
	Description                                            *string                `json:"description,omitempty"`
	// Meta data information                                                      
	Meta                                                   map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                        
	Name                                                   *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                     
	SessionDuration                                        *float64               `json:"sessionDuration,omitempty"`
	// The schedule                                                               
	TriggerSchedule                                        *Schedule              `json:"triggerSchedule,omitempty"`
}

type TriggerIntegrationUpdateResponse struct {
	// The ID of the Trigger Integration       
	ID                                  string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type TriggerIntegrationCreateRequest struct {
	// When enabled the integration requires authentication                       
	Authenticate                                           *bool                  `json:"authenticate,omitempty"`
	// The ID of the blueprint                                                    
	BlueprintID                                            *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                              
	BotID                                                  *string                `json:"botId,omitempty"`
	// The associated description                                                 
	Description                                            *string                `json:"description,omitempty"`
	// Meta data information                                                      
	Meta                                                   map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                        
	Name                                                   *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                     
	SessionDuration                                        *float64               `json:"sessionDuration,omitempty"`
	// The schedule                                                               
	TriggerSchedule                                        *Schedule              `json:"triggerSchedule,omitempty"`
}

type TriggerIntegrationCreateResponse struct {
	// The ID of the Trigger Integration       
	ID                                  string `json:"id"`
}

type TriggerIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type TriggerIntegrationsListResponse struct {
	Items []TriggerIntegrationsListResponseItem `json:"items"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type TriggerIntegrationsListResponseItem struct {
	// When enabled the integration requires authentication                       
	Authenticate                                           *bool                  `json:"authenticate,omitempty"`
	// The ID of the blueprint                                                    
	BlueprintID                                            *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                              
	BotID                                                  *string                `json:"botId,omitempty"`
	// The timestamp (ms) when the instance was created                           
	CreatedAt                                              float64                `json:"createdAt"`
	// The associated description                                                 
	Description                                            *string                `json:"description,omitempty"`
	// The instance ID                                                            
	ID                                                     string                 `json:"id"`
	// Meta data information                                                      
	Meta                                                   map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                        
	Name                                                   *string                `json:"name,omitempty"`
	// The Trigger integration secret                                             
	Secret                                                 string                 `json:"secret"`
	// The session duration (in milliseconds)                                     
	SessionDuration                                        *float64               `json:"sessionDuration,omitempty"`
	// The schedule                                                               
	TriggerSchedule                                        *Schedule              `json:"triggerSchedule,omitempty"`
	// The timestamp (ms) when the instance was updated                           
	UpdatedAt                                              float64                `json:"updatedAt"`
}

type TwilioIntegrationDeleteParams struct {
	// The ID of the Twilio integration       
	TwilioIntegrationID                string `json:"twilioIntegrationId"`
}

type TwilioIntegrationDeleteResponse struct {
	// The ID of the deleted Twilio integration       
	ID                                         string `json:"id"`
}

type TwilioIntegrationFetchParams struct {
	// The ID of the Twilio integration to retrieve       
	TwilioIntegrationID                            string `json:"twilioIntegrationId"`
}

// Blueprint properties
type TwilioIntegrationFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                          
	BotID                                              *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                            
	ContactCollection                                  *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                 
	SessionDuration                                    *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type TwilioIntegrationSetupParams struct {
	// The ID of the Twilio integration       
	TwilioIntegrationID                string `json:"twilioIntegrationId"`
}

type TwilioIntegrationSetupResponse struct {
	// The ID of the Twilio Integration       
	ID                                 string `json:"id"`
}

type TwilioIntegrationUpdateParams struct {
	// The ID of the Twilio integration       
	TwilioIntegrationID                string `json:"twilioIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type TwilioIntegrationUpdateRequest struct {
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type TwilioIntegrationUpdateResponse struct {
	// The ID of the Twilio Integration       
	ID                                 string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type TwilioIntegrationCreateRequest struct {
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type TwilioIntegrationCreateResponse struct {
	// The ID of the Twilio Integration       
	ID                                 string `json:"id"`
}

type TwilioIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type TwilioIntegrationsListResponse struct {
	Items []TwilioIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type TwilioIntegrationsListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                          
	BotID                                              *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                            
	ContactCollection                                  *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The session duration (in milliseconds)                                 
	SessionDuration                                    *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type WhatsAppIntegrationDeleteParams struct {
	// The ID of the WhatsApp integration       
	WhatsappIntegrationID                string `json:"whatsappIntegrationId"`
}

type WhatsAppIntegrationDeleteResponse struct {
	// The ID of the deleted WhatsApp integration       
	ID                                           string `json:"id"`
}

type WhatsAppIntegrationFetchParams struct {
	// The ID of the WhatsApp integration to retrieve       
	WhatsappIntegrationID                            string `json:"whatsappIntegrationId"`
}

// Blueprint properties
type WhatsAppIntegrationFetchResponse struct {
	// The WhatsApp integration access token (returned as '********' if configured, null                       
	// otherwise)                                                                                              
	AccessToken                                                                         *string                `json:"accessToken,omitempty"`
	// Weather the bot supports attachments                                                                    
	Attachments                                                                         *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                 
	BlueprintID                                                                         *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                           
	BotID                                                                               *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                                                             
	ContactCollection                                                                   *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                                        
	CreatedAt                                                                           float64                `json:"createdAt"`
	// The associated description                                                                              
	Description                                                                         *string                `json:"description,omitempty"`
	// The instance ID                                                                                         
	ID                                                                                  string                 `json:"id"`
	// Meta data information                                                                                   
	Meta                                                                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                     
	Name                                                                                *string                `json:"name,omitempty"`
	// The WhatsApp integration phone number ID                                                                
	PhoneNumberID                                                                       *string                `json:"phoneNumberId,omitempty"`
	// The session duration (in milliseconds)                                                                  
	SessionDuration                                                                     *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                                                        
	UpdatedAt                                                                           float64                `json:"updatedAt"`
	// The WhatsApp integration verify token                                                                   
	VerifyToken                                                                         string                 `json:"verifyToken"`
}

type WhatsAppIntegrationSetupParams struct {
	// The ID of the WhatsApp integration       
	WhatsappIntegrationID                string `json:"whatsappIntegrationId"`
}

type WhatsAppIntegrationSetupResponse struct {
	// The ID of the WhatsApp Integration       
	ID                                   string `json:"id"`
}

type WhatsAppIntegrationUpdateParams struct {
	// The ID of the WhatsApp integration       
	WhatsappIntegrationID                string `json:"whatsappIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type WhatsAppIntegrationUpdateRequest struct {
	// The WhatsApp integration access token                               
	AccessToken                                     *string                `json:"accessToken,omitempty"`
	// Weather the bot supports attachments                                
	Attachments                                     *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The WhatsApp integration phone number ID                            
	PhoneNumberID                                   *string                `json:"phoneNumberId,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type WhatsAppIntegrationUpdateResponse struct {
	// The ID of the WhatsApp Integration       
	ID                                   string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type WhatsAppIntegrationCreateRequest struct {
	// The WhatsApp integration access token                               
	AccessToken                                     *string                `json:"accessToken,omitempty"`
	// Weather the bot supports attachments                                
	Attachments                                     *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                             
	BlueprintID                                     *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                       
	BotID                                           *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                         
	ContactCollection                               *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                          
	Description                                     *string                `json:"description,omitempty"`
	// Meta data information                                               
	Meta                                            map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                 
	Name                                            *string                `json:"name,omitempty"`
	// The WhatsApp integration phone number ID                            
	PhoneNumberID                                   *string                `json:"phoneNumberId,omitempty"`
	// The session duration (in milliseconds)                              
	SessionDuration                                 *float64               `json:"sessionDuration,omitempty"`
}

type WhatsAppIntegrationCreateResponse struct {
	// The ID of the WhatsApp Integration       
	ID                                   string `json:"id"`
}

type WhatsAppIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type WhatsAppIntegrationsListResponse struct {
	Items []WhatsAppIntegrationsListResponseItem `json:"items"`
}

// Blueprint properties
type WhatsAppIntegrationsListResponseItem struct {
	// The WhatsApp integration access token (returned as '********' if configured, null                       
	// otherwise)                                                                                              
	AccessToken                                                                         *string                `json:"accessToken,omitempty"`
	// Weather the bot supports attachments                                                                    
	Attachments                                                                         *bool                  `json:"attachments,omitempty"`
	// The ID of the blueprint                                                                                 
	BlueprintID                                                                         *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                                           
	BotID                                                                               *string                `json:"botId,omitempty"`
	// Weather to collect contacts                                                                             
	ContactCollection                                                                   *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                                        
	CreatedAt                                                                           float64                `json:"createdAt"`
	// The associated description                                                                              
	Description                                                                         *string                `json:"description,omitempty"`
	// The instance ID                                                                                         
	ID                                                                                  string                 `json:"id"`
	// Meta data information                                                                                   
	Meta                                                                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                     
	Name                                                                                *string                `json:"name,omitempty"`
	// The WhatsApp integration phone number ID                                                                
	PhoneNumberID                                                                       *string                `json:"phoneNumberId,omitempty"`
	// The session duration (in milliseconds)                                                                  
	SessionDuration                                                                     *float64               `json:"sessionDuration,omitempty"`
	// The timestamp (ms) when the instance was updated                                                        
	UpdatedAt                                                                           float64                `json:"updatedAt"`
	// The WhatsApp integration verify token                                                                   
	VerifyToken                                                                         string                 `json:"verifyToken"`
}

type WidgetIntegrationDeleteParams struct {
	// The ID of the Widget integration       
	WidgetIntegrationID                string `json:"widgetIntegrationId"`
}

type WidgetIntegrationDeleteResponse struct {
	// The ID of the deleted Widget integration       
	ID                                         string `json:"id"`
}

type WidgetIntegrationFetchParams struct {
	// The ID of the Widget integration to retrieve       
	WidgetIntegrationID                            string `json:"widgetIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type WidgetIntegrationFetchResponse struct {
	// Whether the Widget integration supports attachments                                         
	Attachments                                                             *bool                  `json:"attachments,omitempty"`
	// Whether the Widget integration auto scrolls                                                 
	AutoScroll                                                              *bool                  `json:"autoScroll,omitempty"`
	// The ID of the blueprint                                                                     
	BlueprintID                                                             *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                               
	BotID                                                                   *string                `json:"botId,omitempty"`
	// Whether the Widget integration supports carousels                                           
	Carousel                                                                *bool                  `json:"carousel,omitempty"`
	// Whether the Widget integration collects contacts                                            
	ContactCollection                                                       *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                            
	CreatedAt                                                               float64                `json:"createdAt"`
	// The associated description                                                                  
	Description                                                             *string                `json:"description,omitempty"`
	// Controls whether the Widget allows exporting the current conversation                       
	ExportConversation                                                      *bool                  `json:"exportConversation,omitempty"`
	// Whether the Widget integration supports forms                                               
	From                                                                    *bool                  `json:"from,omitempty"`
	// The instance ID                                                                             
	ID                                                                      string                 `json:"id"`
	// The initial message of the Widget integration                                               
	Initial                                                                 *string                `json:"initial,omitempty"`
	// The intro of the Widget integration                                                         
	Intro                                                                   *string                `json:"intro,omitempty"`
	// The language of the Widget integration                                                      
	Language                                                                *string                `json:"language,omitempty"`
	// The default layout of the Widget integration                                                
	Layout                                                                  *string                `json:"layout,omitempty"`
	// Whether the Widget integration supports math                                                
	Math                                                                    *bool                  `json:"math,omitempty"`
	// Controls whether the Widget allows maximizing the conversation                              
	Maximize                                                                *bool                  `json:"maximize,omitempty"`
	// Controls whether the Widget allows peeking at the initial messages                          
	MessagePeek                                                             *bool                  `json:"messagePeek,omitempty"`
	// Meta data information                                                                       
	Meta                                                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                         
	Name                                                                    *string                `json:"name,omitempty"`
	// The origin URLs of the Widget integration                                                   
	Origin                                                                  *string                `json:"origin,omitempty"`
	// The input placeholder of the Widget integration                                             
	Placeholder                                                             *string                `json:"placeholder,omitempty"`
	// The plugins of the Widget integration                                                       
	Plugins                                                                 *string                `json:"plugins,omitempty"`
	// Whether the Widget integration displays powered by                                          
	PoweredBy                                                               *bool                  `json:"poweredBy,omitempty"`
	// Controls whether the Widget allows restarting the conversation                              
	RestartConversation                                                     *bool                  `json:"restartConversation,omitempty"`
	// The session duration of the Widget integration                                              
	SessionDuration                                                         *float64               `json:"sessionDuration,omitempty"`
	// Whether the Widget integration starts first                                                 
	StartFirst                                                              *bool                  `json:"startFirst,omitempty"`
	// Whether the Widget integration is streaming                                                 
	Stream                                                                  *bool                  `json:"stream,omitempty"`
	// The theme of the Widget integration                                                         
	Theme                                                                   *string                `json:"theme,omitempty"`
	// The title of the Widget integration                                                         
	Title                                                                   *string                `json:"title,omitempty"`
	// Whether the Widget integration has tools                                                    
	Tools                                                                   *bool                  `json:"tools,omitempty"`
	// Whether the Widget integration unfurls links                                                
	Unfurl                                                                  *bool                  `json:"unfurl,omitempty"`
	// The timestamp (ms) when the instance was updated                                            
	UpdatedAt                                                               float64                `json:"updatedAt"`
	// Whether the Widget integration is verbose                                                   
	Verbose                                                                 *bool                  `json:"verbose,omitempty"`
	// Whether the Widget integration supports voice input                                         
	VoiceIn                                                                 *bool                  `json:"voiceIn,omitempty"`
	// Whether the Widget integration supports voice output                                        
	VoiceOut                                                                *bool                  `json:"voiceOut,omitempty"`
}

type WidgetIntegrationSetupParams struct {
	// The ID of the Widget integration       
	WidgetIntegrationID                string `json:"widgetIntegrationId"`
}

type WidgetIntegrationSetupResponse struct {
	// The ID of the Widget integration       
	ID                                 string `json:"id"`
}

type WidgetIntegrationUpdateParams struct {
	// The ID of the Widget integration       
	WidgetIntegrationID                string `json:"widgetIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type WidgetIntegrationUpdateRequest struct {
	// Whether the Widget integration supports attachments                                         
	Attachments                                                             *bool                  `json:"attachments,omitempty"`
	// Whether the Widget integration auto scrolls                                                 
	AutoScroll                                                              *bool                  `json:"autoScroll,omitempty"`
	// The ID of the blueprint                                                                     
	BlueprintID                                                             *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                               
	BotID                                                                   *string                `json:"botId,omitempty"`
	// Whether the Widget integration supports carousels                                           
	Carousel                                                                *bool                  `json:"carousel,omitempty"`
	// Whether the Widget integration collects contacts                                            
	ContactCollection                                                       *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                                                  
	Description                                                             *string                `json:"description,omitempty"`
	// Controls whether the Widget allows exporting the current conversation                       
	ExportConversation                                                      *bool                  `json:"exportConversation,omitempty"`
	// Whether the Widget integration supports forms                                               
	Form                                                                    *bool                  `json:"form,omitempty"`
	// The initial message of the Widget integration                                               
	Initial                                                                 *string                `json:"initial,omitempty"`
	// The intro of the Widget integration                                                         
	Intro                                                                   *string                `json:"intro,omitempty"`
	// The language of the Widget integration                                                      
	Language                                                                *string                `json:"language,omitempty"`
	// The default layout of the Widget integration                                                
	Layout                                                                  *string                `json:"layout,omitempty"`
	// Whether the Widget integration supports math                                                
	Math                                                                    *bool                  `json:"math,omitempty"`
	// Controls whether the Widget allows maximizing the conversation                              
	Maximize                                                                *bool                  `json:"maximize,omitempty"`
	// Controls whether the Widget allows peeking at the initial messages                          
	MessagePeek                                                             *bool                  `json:"messagePeek,omitempty"`
	// Meta data information                                                                       
	Meta                                                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                         
	Name                                                                    *string                `json:"name,omitempty"`
	// The origin URLs of the Widget integration                                                   
	Origin                                                                  *string                `json:"origin,omitempty"`
	// The input placeholder of the Widget integration                                             
	Placeholder                                                             *string                `json:"placeholder,omitempty"`
	// The plugins of the Widget integration                                                       
	Plugins                                                                 *string                `json:"plugins,omitempty"`
	// Whether the Widget integration displays powered by                                          
	PoweredBy                                                               *bool                  `json:"poweredBy,omitempty"`
	// Controls whether the Widget allows restarting the conversation                              
	RestartConversation                                                     *bool                  `json:"restartConversation,omitempty"`
	// The session duration of the Widget integration                                              
	SessionDuration                                                         *float64               `json:"sessionDuration,omitempty"`
	// Whether the Widget integration starts first                                                 
	StartFirst                                                              *bool                  `json:"startFirst,omitempty"`
	// Whether the Widget integration is streaming                                                 
	Stream                                                                  *bool                  `json:"stream,omitempty"`
	// The theme of the Widget integration                                                         
	Theme                                                                   *string                `json:"theme,omitempty"`
	// The title of the Widget integration                                                         
	Title                                                                   *string                `json:"title,omitempty"`
	// Whether the Widget integration has tools                                                    
	Tools                                                                   *bool                  `json:"tools,omitempty"`
	// Whether the Widget integration unfurls links                                                
	Unfurl                                                                  *bool                  `json:"unfurl,omitempty"`
	// Whether the Widget integration is verbose                                                   
	Verbose                                                                 *bool                  `json:"verbose,omitempty"`
	// Controls whether the Widget allows voice input                                              
	VoiceIn                                                                 *bool                  `json:"voiceIn,omitempty"`
	// Controls whether the Widget allows voice output                                             
	VoiceOut                                                                *bool                  `json:"voiceOut,omitempty"`
}

type WidgetIntegrationUpdateResponse struct {
	// The ID of the Widget Integration       
	ID                                 string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type WidgetIntegrationCreateRequest struct {
	// Weather the Widget integration supports attachments                                         
	Attachments                                                             *bool                  `json:"attachments,omitempty"`
	// Whether the Widget integration auto scrolls                                                 
	AutoScroll                                                              *bool                  `json:"autoScroll,omitempty"`
	// The ID of the blueprint                                                                     
	BlueprintID                                                             *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                               
	BotID                                                                   *string                `json:"botId,omitempty"`
	// Weather the Widget integration supports carousels                                           
	Carousel                                                                *bool                  `json:"carousel,omitempty"`
	// Whether the Widget integration collects contacts                                            
	ContactCollection                                                       *bool                  `json:"contactCollection,omitempty"`
	// The associated description                                                                  
	Description                                                             *string                `json:"description,omitempty"`
	// Controls whether the Widget allows exporting the current conversation                       
	ExportConversation                                                      *bool                  `json:"exportConversation,omitempty"`
	// Weather the Widget integration supports forms                                               
	Form                                                                    *bool                  `json:"form,omitempty"`
	// The initial message of the Widget integration                                               
	Initial                                                                 *string                `json:"initial,omitempty"`
	// The intro of the Widget integration                                                         
	Intro                                                                   *string                `json:"intro,omitempty"`
	// The language of the Widget integration                                                      
	Language                                                                *string                `json:"language,omitempty"`
	// The default layout of the Widget integration                                                
	Layout                                                                  *string                `json:"layout,omitempty"`
	// Weather the Widget integration supports math                                                
	Math                                                                    *bool                  `json:"math,omitempty"`
	// Controls whether the Widget allows maximizing the conversation                              
	Maximize                                                                *bool                  `json:"maximize,omitempty"`
	// Controls whether the Widget allows peeking at the initial messages                          
	MessagePeek                                                             *bool                  `json:"messagePeek,omitempty"`
	// Meta data information                                                                       
	Meta                                                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                         
	Name                                                                    *string                `json:"name,omitempty"`
	// The origin URLs of the Widget integration                                                   
	Origin                                                                  *string                `json:"origin,omitempty"`
	// The input placeholder of the Widget integration                                             
	Placeholder                                                             *string                `json:"placeholder,omitempty"`
	// The plugins of the Widget integration                                                       
	Plugins                                                                 *string                `json:"plugins,omitempty"`
	// Whether the Widget integration displays powered by                                          
	PoweredBy                                                               *bool                  `json:"poweredBy,omitempty"`
	// Controls whether the Widget allows restarting the conversation                              
	RestartConversation                                                     *bool                  `json:"restartConversation,omitempty"`
	// The session duration of the Widget integration                                              
	SessionDuration                                                         *float64               `json:"sessionDuration,omitempty"`
	// Whether the Widget integration starts first                                                 
	StartFirst                                                              *bool                  `json:"startFirst,omitempty"`
	// Whether the Widget integration is streaming                                                 
	Stream                                                                  *bool                  `json:"stream,omitempty"`
	// The theme of the Widget integration                                                         
	Theme                                                                   *string                `json:"theme,omitempty"`
	// The title of the Widget integration                                                         
	Title                                                                   *string                `json:"title,omitempty"`
	// Whether the Widget integration has tools                                                    
	Tools                                                                   *bool                  `json:"tools,omitempty"`
	// Whether the Widget integration unfurls links                                                
	Unfurl                                                                  *bool                  `json:"unfurl,omitempty"`
	// Whether the Widget integration is verbose                                                   
	Verbose                                                                 *bool                  `json:"verbose,omitempty"`
	// Controls whether the Widget allows voice input                                              
	VoiceIn                                                                 *bool                  `json:"voiceIn,omitempty"`
	// Controls whether the Widget allows voice output                                             
	VoiceOut                                                                *bool                  `json:"voiceOut,omitempty"`
}

type WidgetIntegrationCreateResponse struct {
	// The ID of the Widget Integration       
	ID                                 string `json:"id"`
}

type WidgetIntegrationsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type WidgetIntegrationsListResponse struct {
	Items []WidgetIntegrationsListResponseItem `json:"items"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type WidgetIntegrationsListResponseItem struct {
	// Weather the Widget integration supports attachments                                         
	Attachments                                                             *bool                  `json:"attachments,omitempty"`
	// Whether the Widget integration auto scrolls                                                 
	AutoScroll                                                              *bool                  `json:"autoScroll,omitempty"`
	// The ID of the blueprint                                                                     
	BlueprintID                                                             *string                `json:"blueprintId,omitempty"`
	// The ID of the bot this configuration is using                                               
	BotID                                                                   *string                `json:"botId,omitempty"`
	// Weather the Widget integration supports carousels                                           
	Carousel                                                                *bool                  `json:"carousel,omitempty"`
	// Whether the Widget integration collects contacts                                            
	ContactCollection                                                       *bool                  `json:"contactCollection,omitempty"`
	// The timestamp (ms) when the instance was created                                            
	CreatedAt                                                               float64                `json:"createdAt"`
	// The associated description                                                                  
	Description                                                             *string                `json:"description,omitempty"`
	// Controls whether the Widget allows exporting the current conversation                       
	ExportConversation                                                      *bool                  `json:"exportConversation,omitempty"`
	// Weather the Widget integration supports forms                                               
	Form                                                                    *bool                  `json:"form,omitempty"`
	// The instance ID                                                                             
	ID                                                                      string                 `json:"id"`
	// The initial message of the Widget integration                                               
	Initial                                                                 *string                `json:"initial,omitempty"`
	// The intro of the Widget integration                                                         
	Intro                                                                   *string                `json:"intro,omitempty"`
	// The language of the Widget integration                                                      
	Language                                                                *string                `json:"language,omitempty"`
	// The default layout of the Widget integration                                                
	Layout                                                                  *string                `json:"layout,omitempty"`
	// Weather the Widget integration supports math                                                
	Math                                                                    *bool                  `json:"math,omitempty"`
	// Controls whether the Widget allows maximizing the conversation                              
	Maximize                                                                *bool                  `json:"maximize,omitempty"`
	// Controls whether the Widget allows peeking at the initial messages                          
	MessagePeek                                                             *bool                  `json:"messagePeek,omitempty"`
	// Meta data information                                                                       
	Meta                                                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                         
	Name                                                                    *string                `json:"name,omitempty"`
	// The origin URLs of the Widget integration                                                   
	Origin                                                                  *string                `json:"origin,omitempty"`
	// The input placeholder of the Widget integration                                             
	Placeholder                                                             *string                `json:"placeholder,omitempty"`
	// The plugins of the Widget integration                                                       
	Plugins                                                                 *string                `json:"plugins,omitempty"`
	// Whether the Widget integration displays powered by                                          
	PoweredBy                                                               *bool                  `json:"poweredBy,omitempty"`
	// Controls whether the Widget allows restarting the conversation                              
	RestartConversation                                                     *bool                  `json:"restartConversation,omitempty"`
	// The session duration of the Widget integration                                              
	SessionDuration                                                         *float64               `json:"sessionDuration,omitempty"`
	// Whether the Widget integration starts first                                                 
	StartFirst                                                              *bool                  `json:"startFirst,omitempty"`
	// Whether the Widget integration is streaming                                                 
	Stream                                                                  *bool                  `json:"stream,omitempty"`
	// The theme of the Widget integration                                                         
	Theme                                                                   *string                `json:"theme,omitempty"`
	// The title of the Widget integration                                                         
	Title                                                                   *string                `json:"title,omitempty"`
	// Whether the Widget integration has tools                                                    
	Tools                                                                   *bool                  `json:"tools,omitempty"`
	// Whether the Widget integration unfurls links                                                
	Unfurl                                                                  *bool                  `json:"unfurl,omitempty"`
	// The timestamp (ms) when the instance was updated                                            
	UpdatedAt                                                               float64                `json:"updatedAt"`
	// Whether the Widget integration is verbose                                                   
	Verbose                                                                 *bool                  `json:"verbose,omitempty"`
	// Whether the Widget integration supports voice input                                         
	VoiceIn                                                                 *bool                  `json:"voiceIn,omitempty"`
	// Whether the Widget integration supports voice output                                        
	VoiceOut                                                                *bool                  `json:"voiceOut,omitempty"`
}

type MagicFromPromptGenerateParams struct {
	// The ID of the prompt to use for generation       
	PromptID                                     string `json:"promptId"`
}

type MagicFromPromptGenerateRequest struct {
	// Optional language model to use for generation                       
	Model                                           *string                `json:"model,omitempty"`
	// Additional properties to pass to the prompt                         
	Props                                           map[string]interface{} `json:"props,omitempty"`
	// The text to use as input                                            
	Text                                            string                 `json:"text"`
}

type MagicFromPromptGenerateResponse struct {
	// The input text                                        
	Text                string                               `json:"text"`
	// Usage information                                     
	Usage               MagicFromPromptGenerateResponseUsage `json:"usage"`
}

// Usage information
type MagicFromPromptGenerateResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

type MagicPromptsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type MagicPromptsListResponse struct {
	Items []MagicPromptsListResponseItem `json:"items"`
}

// Instance list properties
type MagicPromptsListResponseItem struct {
	// The alias of the item                                                  
	Alias                                              string                 `json:"alias"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type MemoryDeleteParams struct {
	// The ID of the memory to delete       
	MemoryID                         string `json:"memoryId"`
}

type MemoryDeleteResponse struct {
	// The ID of the deleted memory       
	ID                             string `json:"id"`
}

type MemoryFetchParams struct {
	// The ID of the memory to retrieve       
	MemoryID                           string `json:"memoryId"`
}

// Instance list properties
type MemoryFetchResponse struct {
	// The bot associated with the memory                                     
	BotID                                              *string                `json:"botId,omitempty"`
	// The contact associated with the memory                                 
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The text of the memory                                                 
	Text                                               *string                `json:"text,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type MemoryUpdateParams struct {
	MemoryID string `json:"memoryId"`
}

// Instance crud properties
type MemoryUpdateRequest struct {
	// The bot associated with the memory                           
	BotID                                    *string                `json:"botId,omitempty"`
	// The contact associated with the memory                       
	ContactID                                *string                `json:"contactId,omitempty"`
	// The associated description                                   
	Description                              *string                `json:"description,omitempty"`
	// Meta data information                                        
	Meta                                     map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                          
	Name                                     *string                `json:"name,omitempty"`
	// The text of the memory                                       
	Text                                     *string                `json:"text,omitempty"`
}

type MemoryUpdateResponse struct {
	// The ID of the updated memory       
	ID                             string `json:"id"`
}

// Instance crud properties
type MemoryCreateRequest struct {
	// The bot associated with the memory                           
	BotID                                    *string                `json:"botId,omitempty"`
	// The contact associated with the memory                       
	ContactID                                *string                `json:"contactId,omitempty"`
	// The associated description                                   
	Description                              *string                `json:"description,omitempty"`
	// Meta data information                                        
	Meta                                     map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                          
	Name                                     *string                `json:"name,omitempty"`
	// The text of the memory                                       
	Text                                     *string                `json:"text,omitempty"`
}

type MemoryCreateResponse struct {
	// The ID of the created memory       
	ID                             string `json:"id"`
}

type MemoriesExportParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type MemoriesExportResponse struct {
	Items []MemoriesExportResponseItem `json:"items"`
}

// Instance list properties
type MemoriesExportResponseItem struct {
	// The bot associated with the memory                                     
	BotID                                              *string                `json:"botId,omitempty"`
	// The contact associated with the memory                                 
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The text of the memory                                                 
	Text                                               *string                `json:"text,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type MemoriesListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type MemoriesListResponse struct {
	Items []MemoriesListResponseItem `json:"items"`
}

// Instance list properties
type MemoriesListResponseItem struct {
	// The bot associated with the memory                                     
	BotID                                              *string                `json:"botId,omitempty"`
	// The contact associated with the memory                                 
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The text of the memory                                                 
	Text                                               *string                `json:"text,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type MemorySearchRequest struct {
	// The ID of the bot to filter memories by            
	BotID                                         *string `json:"botId,omitempty"`
	// The ID of the contact to filter memories by        
	ContactID                                     *string `json:"contactId,omitempty"`
	// The keyword/phrase to search for                   
	Search                                        string  `json:"search"`
}

type MemorySearchResponse struct {
	// An array of memories matching the search query                           
	Items                                            []MemorySearchResponseItem `json:"items"`
}

type MemorySearchResponseItem struct {
	ID   string                 `json:"id"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Text string                 `json:"text"`
}

type PartnerUserDeleteParams struct {
	// The ID of the user to delete       
	UserID                         string `json:"userId"`
}

type PartnerUserDeleteResponse struct {
	// The ID of the deleted user       
	ID                           string `json:"id"`
}

type PartnerUserFetchParams struct {
	// The ID of the partner user to retrieve       
	UserID                                   string `json:"userId"`
}

// Instance list properties
type PartnerUserFetchResponse struct {
	// The timestamp (ms) when the instance was created                                
	CreatedAt                                          float64                         `json:"createdAt"`
	// The associated description                                                      
	Description                                        *string                         `json:"description,omitempty"`
	// The email of the partner user                                                   
	Email                                              *string                         `json:"email,omitempty"`
	// The instance ID                                                                 
	ID                                                 string                          `json:"id"`
	// The image of the partner user                                                   
	Image                                              *string                         `json:"image,omitempty"`
	// Limits information                                                              
	Limits                                             *PartnerUserFetchResponseLimits `json:"limits,omitempty"`
	// Meta data information                                                           
	Meta                                               map[string]interface{}          `json:"meta,omitempty"`
	// The associated name                                                             
	Name                                               *string                         `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                                
	UpdatedAt                                          float64                         `json:"updatedAt"`
}

// Limits information
type PartnerUserFetchResponseLimits struct {
	// The conversations limit                
	Conversations             *float64        `json:"conversations,omitempty"`
	// The database limits                    
	Database                  *PurpleDatabase `json:"database,omitempty"`
	// The messages limit                     
	Messages                  *float64        `json:"messages,omitempty"`
	// The tokens limit                       
	Tokens                    *float64        `json:"tokens,omitempty"`
}

// The database limits
type PurpleDatabase struct {
	// The abilities limit         
	Abilities             *float64 `json:"abilities,omitempty"`
	// The datasets limit          
	Datasets              *float64 `json:"datasets,omitempty"`
	// The files limit             
	Files                 *float64 `json:"files,omitempty"`
	// The records limit           
	Records               *float64 `json:"records,omitempty"`
	// The skillsets limit         
	Skillsets             *float64 `json:"skillsets,omitempty"`
}

type PartnerUserTokenDeleteParams struct {
	// The ID of the user token to delete       
	TokenID                              string `json:"tokenId"`
	// The ID of the user                       
	UserID                               string `json:"userId"`
}

type PartnerUserTokenDeleteResponse struct {
	// The ID of the deleted user token       
	ID                                 string `json:"id"`
}

type PartnerUserTokenCreateParams struct {
	// The ID of the user       
	UserID               string `json:"userId"`
}

type PartnerUserTokenCreateResponse struct {
	// The timestamp for when the user token was created (in milliseconds)        
	CreatedAt                                                             float64 `json:"createdAt"`
	// The ID of the created user token                                           
	ID                                                                    string  `json:"id"`
	// The token of the created user token                                        
	Token                                                                 string  `json:"token"`
}

type PartnerUserTokensListParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
	// The ID of the user                      
	UserID                             string  `json:"userId"`
}

type PartnerUserTokensListResponse struct {
	Items []PartnerUserTokensListResponseItem `json:"items"`
}

// Instance list properties
type PartnerUserTokensListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PartnerUserUpdateParams struct {
	// The ID of the partner user       
	UserID                       string `json:"userId"`
}

// Instance crud properties
type PartnerUserUpdateRequest struct {
	// The associated description                                   
	Description                     *string                         `json:"description,omitempty"`
	// The email of the partner user                                
	Email                           *string                         `json:"email,omitempty"`
	// The image of the partner user                                
	Image                           *string                         `json:"image,omitempty"`
	// Limits information                                           
	Limits                          *PartnerUserUpdateRequestLimits `json:"limits,omitempty"`
	// Meta data information                                        
	Meta                            map[string]interface{}          `json:"meta,omitempty"`
	// The associated name                                          
	Name                            *string                         `json:"name,omitempty"`
}

// Limits information
type PartnerUserUpdateRequestLimits struct {
	// The conversations limit                
	Conversations             *float64        `json:"conversations,omitempty"`
	// The database limits                    
	Database                  *FluffyDatabase `json:"database,omitempty"`
	// The messages limit                     
	Messages                  *float64        `json:"messages,omitempty"`
	// The tokens limit                       
	Tokens                    *float64        `json:"tokens,omitempty"`
}

// The database limits
type FluffyDatabase struct {
	// The abilities limit         
	Abilities             *float64 `json:"abilities,omitempty"`
	// The datasets limit          
	Datasets              *float64 `json:"datasets,omitempty"`
	// The files limit             
	Files                 *float64 `json:"files,omitempty"`
	// The records limit           
	Records               *float64 `json:"records,omitempty"`
	// The skillsets limit         
	Skillsets             *float64 `json:"skillsets,omitempty"`
}

type PartnerUserUpdateResponse struct {
	// The ID of the updated partner user       
	ID                                   string `json:"id"`
}

// Instance crud properties
type PartnerUserCreateRequest struct {
	// The associated description                                   
	Description                     *string                         `json:"description,omitempty"`
	// The email of the partner user                                
	Email                           *string                         `json:"email,omitempty"`
	// The image of the partner user                                
	Image                           *string                         `json:"image,omitempty"`
	// Limits information                                           
	Limits                          *PartnerUserCreateRequestLimits `json:"limits,omitempty"`
	// Meta data information                                        
	Meta                            map[string]interface{}          `json:"meta,omitempty"`
	// The associated name                                          
	Name                            *string                         `json:"name,omitempty"`
}

// Limits information
type PartnerUserCreateRequestLimits struct {
	// The conversations limit                   
	Conversations             *float64           `json:"conversations,omitempty"`
	// The database limits                       
	Database                  *TentacledDatabase `json:"database,omitempty"`
	// The messages limit                        
	Messages                  *float64           `json:"messages,omitempty"`
	// The tokens limit                          
	Tokens                    *float64           `json:"tokens,omitempty"`
}

// The database limits
type TentacledDatabase struct {
	// The abilities limit         
	Abilities             *float64 `json:"abilities,omitempty"`
	// The datasets limit          
	Datasets              *float64 `json:"datasets,omitempty"`
	// The files limit             
	Files                 *float64 `json:"files,omitempty"`
	// The records limit           
	Records               *float64 `json:"records,omitempty"`
	// The skillsets limit         
	Skillsets             *float64 `json:"skillsets,omitempty"`
}

type PartnerUserCreateResponse struct {
	// The ID of the created user       
	ID                           string `json:"id"`
}

type PartnerUsersListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PartnerUsersListResponse struct {
	Items []PartnerUsersListResponseItem `json:"items"`
}

// Instance list properties
type PartnerUsersListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The email of the partner user                                          
	Email                                              *string                `json:"email,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The image of the partner user                                          
	Image                                              *string                `json:"image,omitempty"`
	// Limits information                                                     
	Limits                                             *ItemLimits            `json:"limits,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

// Limits information
type ItemLimits struct {
	// The conversations limit                
	Conversations             *float64        `json:"conversations,omitempty"`
	// The database limits                    
	Database                  *StickyDatabase `json:"database,omitempty"`
	// The messages limit                     
	Messages                  *float64        `json:"messages,omitempty"`
	// The tokens limit                       
	Tokens                    *float64        `json:"tokens,omitempty"`
}

// The database limits
type StickyDatabase struct {
	// The abilities limit         
	Abilities             *float64 `json:"abilities,omitempty"`
	// The datasets limit          
	Datasets              *float64 `json:"datasets,omitempty"`
	// The files limit             
	Files                 *float64 `json:"files,omitempty"`
	// The records limit           
	Records               *float64 `json:"records,omitempty"`
	// The skillsets limit         
	Skillsets             *float64 `json:"skillsets,omitempty"`
}

type PlatformAbilitiesListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PlatformAbilitiesListResponse struct {
	Items []PlatformAbilitiesListResponseItem `json:"items"`
}

// Instance list properties
type PlatformAbilitiesListResponseItem struct {
	// The ID of the bot associated with the ability                                                               
	Bot                                                                                     *string                `json:"bot,omitempty"`
	Commentary                                                                              *string                `json:"commentary,omitempty"`
	// The timestamp (ms) when the instance was created                                                            
	CreatedAt                                                                               float64                `json:"createdAt"`
	// The associated description                                                                                  
	Description                                                                             *string                `json:"description,omitempty"`
	// The ID of the file associated with the ability                                                              
	File                                                                                    *string                `json:"file,omitempty"`
	Icon                                                                                    string                 `json:"icon"`
	// The instance ID                                                                                             
	ID                                                                                      string                 `json:"id"`
	Instruction                                                                             string                 `json:"instruction"`
	// Meta data information                                                                                       
	Meta                                                                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                                                         
	Name                                                                                    *string                `json:"name,omitempty"`
	// The provider of the ability                                                                                 
	Provider                                                                                *string                `json:"provider,omitempty"`
	// A JSON Schema object type definition (https://json-schema.org/). Represents an object                       
	// schema with properties and validation rules.                                                                
	Schema                                                                                  Schema                 `json:"schema"`
	// The ID of the secret associated with the ability                                                            
	Secret                                                                                  *string                `json:"secret,omitempty"`
	Setup                                                                                   *string                `json:"setup,omitempty"`
	// The ID of the space associated with the ability                                                             
	Space                                                                                   *string                `json:"space,omitempty"`
	Tags                                                                                    []string               `json:"tags,omitempty"`
	// The timestamp (ms) when the instance was updated                                                            
	UpdatedAt                                                                               float64                `json:"updatedAt"`
}

// A JSON Schema object type definition (https://json-schema.org/). Represents an object
// schema with properties and validation rules.
type Schema struct {
	// The schema description                                  
	Description                         *string                `json:"description,omitempty"`
	// Object property definitions                             
	Properties                          map[string]interface{} `json:"properties"`
	// Required property names                                 
	Required                            []string               `json:"required,omitempty"`
	// The schema title                                        
	Title                               *string                `json:"title,omitempty"`
	// The schema type, must be "object"                       
	Type                                ParametersType         `json:"type"`
}

type PlatformActionsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PlatformActionsListResponse struct {
	Items []PlatformActionsListResponseItem `json:"items"`
}

// Instance list properties
type PlatformActionsListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The description of the action                                          
	Description                                        string                 `json:"description"`
	// Example demonstrating the action usage                                 
	Examples                                           []string               `json:"examples"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformDocFetchParams struct {
	// The ID of the doc to fetch (e.g., "datasets", "skillsets")       
	DocID                                                        string `json:"docId"`
}

// Instance list properties
type PlatformDocFetchResponse struct {
	// The category of the manual                                             
	Category                                           *string                `json:"category,omitempty"`
	// The markdown content of the doc                                        
	Content                                            string                 `json:"content"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The display order index                                                
	Index                                              *float64               `json:"index,omitempty"`
	// The URL to the official documentation page                             
	Link                                               *string                `json:"link,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// Tags associated with the doc                                           
	Tags                                               []string               `json:"tags,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformDocsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PlatformDocsListResponse struct {
	Items []PlatformDocsListResponseItem `json:"items"`
}

// Instance list properties
type PlatformDocsListResponseItem struct {
	// The category of the doc                                                
	Category                                           *string                `json:"category,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        string                 `json:"description"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The display order index                                                
	Index                                              float64                `json:"index"`
	// The URL to the official documentation page                             
	Link                                               string                 `json:"link"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// Tags associated with the doc                                           
	Tags                                               []string               `json:"tags"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformDocsSearchRequest struct {
	// The search query to find relevant docs                            
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type PlatformDocsSearchResponse struct {
	Items []PlatformDocsSearchResponseItem `json:"items"`
}

// Instance list properties
type PlatformDocsSearchResponseItem struct {
	// The category of the doc                                                 
	Category                                            *string                `json:"category,omitempty"`
	// The timestamp (ms) when the instance was created                        
	CreatedAt                                           float64                `json:"createdAt"`
	// The associated description                                              
	Description                                         string                 `json:"description"`
	// An excerpt from the most relevant part of the doc                       
	Excerpt                                             string                 `json:"excerpt"`
	// The instance ID                                                         
	ID                                                  string                 `json:"id"`
	// The display order index                                                 
	Index                                               float64                `json:"index"`
	// The URL to the official documentation page                              
	Link                                                string                 `json:"link"`
	// Meta data information                                                   
	Meta                                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                     
	Name                                                string                 `json:"name"`
	// The similarity score of the search result                               
	Score                                               float64                `json:"score"`
	// Tags associated with the doc                                            
	Tags                                                []string               `json:"tags"`
	// The timestamp (ms) when the instance was updated                        
	UpdatedAt                                           float64                `json:"updatedAt"`
}

type PlatformExampleCloneParams struct {
	// The ID (slug) of the example to clone       
	ExampleID                               string `json:"exampleId"`
}

type PlatformExampleCloneResponse struct {
	// A map of resource types to arrays of created resources                      
	Resources                                                map[string][]Resource `json:"resources"`
}

type Resource struct {
	// The description of the resource              
	Description                             *string `json:"description,omitempty"`
	// The unique identifier of the resource        
	ID                                      string  `json:"id"`
	// The name of the resource                     
	Name                                    *string `json:"name,omitempty"`
}

type PlatformExampleFetchParams struct {
	// The ID (slug) of the example       
	ExampleID                      string `json:"exampleId"`
}

type PlatformExampleFetchResponse struct {
	// The full configuration details of the example                                 
	Config                                          map[string]interface{}           `json:"config"`
	// The creation timestamp                                                        
	CreatedAt                                       *float64                         `json:"createdAt,omitempty"`
	// The description of the example                                                
	Description                                     string                           `json:"description"`
	// The ID (slug) of the example                                                  
	ID                                              string                           `json:"id"`
	// The URL to the official example page                                          
	Link                                            string                           `json:"link"`
	// The name of the example                                                       
	Name                                            string                           `json:"name"`
	// Tags associated with the example                                              
	Tags                                            []string                         `json:"tags,omitempty"`
	// The type of the example                                                       
	Type                                            PlatformExampleFetchResponseType `json:"type"`
	// The last update timestamp                                                     
	UpdatedAt                                       *float64                         `json:"updatedAt,omitempty"`
}

type PlatformExamplesListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PlatformExamplesListResponse struct {
	Items []PlatformExamplesListResponseItem `json:"items"`
}

// Instance list properties
type PlatformExamplesListResponseItem struct {
	// The timestamp (ms) when the instance was created                                 
	CreatedAt                                          float64                          `json:"createdAt"`
	// The associated description                                                       
	Description                                        string                           `json:"description"`
	// The instance ID                                                                  
	ID                                                 string                           `json:"id"`
	// The URL to the official example page                                             
	Link                                               string                           `json:"link"`
	// Meta data information                                                            
	Meta                                               map[string]interface{}           `json:"meta,omitempty"`
	// The associated name                                                              
	Name                                               string                           `json:"name"`
	// Tags associated with the example                                                 
	Tags                                               []string                         `json:"tags,omitempty"`
	// The type of the example                                                          
	Type                                               PlatformExampleFetchResponseType `json:"type"`
	// The timestamp (ms) when the instance was updated                                 
	UpdatedAt                                          float64                          `json:"updatedAt"`
}

type PlatformExamplesSearchRequest struct {
	// The search query to find relevant examples                        
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type PlatformExamplesSearchResponse struct {
	Items []PlatformExamplesSearchResponseItem `json:"items"`
}

// Instance list properties
type PlatformExamplesSearchResponseItem struct {
	// The timestamp (ms) when the instance was created                                 
	CreatedAt                                          float64                          `json:"createdAt"`
	// The associated description                                                       
	Description                                        string                           `json:"description"`
	// The instance ID                                                                  
	ID                                                 string                           `json:"id"`
	// The URL to the official example page                                             
	Link                                               string                           `json:"link"`
	// Meta data information                                                            
	Meta                                               map[string]interface{}           `json:"meta,omitempty"`
	// The associated name                                                              
	Name                                               string                           `json:"name"`
	// Tags associated with the example                                                 
	Tags                                               []string                         `json:"tags,omitempty"`
	// The type of the example                                                          
	Type                                               PlatformExampleFetchResponseType `json:"type"`
	// The timestamp (ms) when the instance was updated                                 
	UpdatedAt                                          float64                          `json:"updatedAt"`
}

type PlatformGuideFetchParams struct {
	// The ID of the guide to fetch       
	GuideID                        string `json:"guideId"`
}

// Instance list properties
type PlatformGuideFetchResponse struct {
	// The category of the guide                                              
	Category                                           *string                `json:"category,omitempty"`
	// The markdown content of the guide                                      
	Content                                            string                 `json:"content"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The display order index                                                
	Index                                              *float64               `json:"index,omitempty"`
	// The URL to the official guide page                                     
	Link                                               *string                `json:"link,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// Tags associated with the guide                                         
	Tags                                               []string               `json:"tags,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformGuidesListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PlatformGuidesListResponse struct {
	Items []PlatformGuidesListResponseItem `json:"items"`
}

// Instance list properties
type PlatformGuidesListResponseItem struct {
	// The category of the guide                                              
	Category                                           *string                `json:"category,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        string                 `json:"description"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The display order index                                                
	Index                                              float64                `json:"index"`
	// The URL to the official guide page                                     
	Link                                               string                 `json:"link"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// Tags associated with the guide                                         
	Tags                                               []string               `json:"tags"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformGuidesSearchRequest struct {
	// The search query to find relevant guides                          
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type PlatformGuidesSearchResponse struct {
	Items []PlatformGuidesSearchResponseItem `json:"items"`
}

// Instance list properties
type PlatformGuidesSearchResponseItem struct {
	// The category of the guide                                                 
	Category                                              *string                `json:"category,omitempty"`
	// The timestamp (ms) when the instance was created                          
	CreatedAt                                             float64                `json:"createdAt"`
	// The associated description                                                
	Description                                           string                 `json:"description"`
	// An excerpt from the most relevant part of the guide                       
	Excerpt                                               string                 `json:"excerpt"`
	// The instance ID                                                           
	ID                                                    string                 `json:"id"`
	// The display order index                                                   
	Index                                                 float64                `json:"index"`
	// The URL to the official guide page                                        
	Link                                                  string                 `json:"link"`
	// Meta data information                                                     
	Meta                                                  map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                       
	Name                                                  string                 `json:"name"`
	// The similarity score of the search result                                 
	Score                                                 float64                `json:"score"`
	// Tags associated with the guide                                            
	Tags                                                  []string               `json:"tags"`
	// The timestamp (ms) when the instance was updated                          
	UpdatedAt                                             float64                `json:"updatedAt"`
}

type PlatformManualFetchParams struct {
	// The ID of the manual to fetch (e.g., "datasets", "skillsets")       
	ManualID                                                        string `json:"manualId"`
}

// Instance list properties
type PlatformManualFetchResponse struct {
	// The category of the manual                                             
	Category                                           *string                `json:"category,omitempty"`
	// The markdown content of the manual                                     
	Content                                            string                 `json:"content"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The display order index                                                
	Index                                              *float64               `json:"index,omitempty"`
	// The URL to the official documentation page                             
	Link                                               *string                `json:"link,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// Tags associated with the manual                                        
	Tags                                               []string               `json:"tags,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformManualsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PlatformManualsListResponse struct {
	Items []PlatformManualsListResponseItem `json:"items"`
}

// Instance list properties
type PlatformManualsListResponseItem struct {
	// The category of the manual                                             
	Category                                           *string                `json:"category,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        string                 `json:"description"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The display order index                                                
	Index                                              float64                `json:"index"`
	// The URL to the official documentation page                             
	Link                                               string                 `json:"link"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// Tags associated with the manual                                        
	Tags                                               []string               `json:"tags"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformManualsSearchRequest struct {
	// The search query to find relevant manuals                         
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type PlatformManualsSearchResponse struct {
	Items []PlatformManualsSearchResponseItem `json:"items"`
}

// Instance list properties
type PlatformManualsSearchResponseItem struct {
	// The category of the manual                                                 
	Category                                               *string                `json:"category,omitempty"`
	// The timestamp (ms) when the instance was created                           
	CreatedAt                                              float64                `json:"createdAt"`
	// The associated description                                                 
	Description                                            string                 `json:"description"`
	// An excerpt from the most relevant part of the manual                       
	Excerpt                                                string                 `json:"excerpt"`
	// The instance ID                                                            
	ID                                                     string                 `json:"id"`
	// The display order index                                                    
	Index                                                  float64                `json:"index"`
	// The URL to the official documentation page                                 
	Link                                                   string                 `json:"link"`
	// Meta data information                                                      
	Meta                                                   map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                        
	Name                                                   string                 `json:"name"`
	// The similarity score of the search result                                  
	Score                                                  float64                `json:"score"`
	// Tags associated with the manual                                            
	Tags                                                   []string               `json:"tags"`
	// The timestamp (ms) when the instance was updated                           
	UpdatedAt                                              float64                `json:"updatedAt"`
}

type PlatformModelsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PlatformModelsListResponse struct {
	Items []PlatformModelsListResponseItem `json:"items"`
}

// Instance list properties
type PlatformModelsListResponseItem struct {
	// The timestamp (ms) when the instance was created                          
	CreatedAt                                             float64                `json:"createdAt"`
	// The associated description                                                
	Description                                           *string                `json:"description,omitempty"`
	// The model of the model                                                    
	Family                                                string                 `json:"family"`
	// The instance ID                                                           
	ID                                                    string                 `json:"id"`
	// The maximum number of tokens the model can accept                         
	MaxInputTokens                                        float64                `json:"maxInputTokens"`
	// The maximum number of tokens the model can generate                       
	MaxOutputTokens                                       float64                `json:"maxOutputTokens"`
	// The maximum number of tokens the model can use                            
	MaxTokens                                             float64                `json:"maxTokens"`
	// Meta data information                                                     
	Meta                                                  map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                       
	Name                                                  *string                `json:"name,omitempty"`
	// The backstory of the model                                                
	Provider                                              string                 `json:"provider"`
	// The timestamp (ms) when the instance was updated                          
	UpdatedAt                                             float64                `json:"updatedAt"`
}

type PlatformSecretsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PlatformSecretsListResponse struct {
	Items []PlatformSecretsListResponseItem `json:"items"`
}

// Instance list properties
type PlatformSecretsListResponseItem struct {
	Commentary                                         *string                `json:"commentary,omitempty"`
	Config                                             map[string]interface{} `json:"config,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	Icon                                               *string                `json:"icon,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The kind of the secret                                                 
	Kind                                               *SecretKind            `json:"kind,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	Setup                                              *string                `json:"setup,omitempty"`
	Tags                                               []string               `json:"tags,omitempty"`
	// The type of the secret                                                 
	Type                                               SecretType             `json:"type"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformTutorialFetchParams struct {
	// The ID of the tutorial to fetch (e.g., "how-to-get-started-with-chatbotkit")       
	TutorialID                                                                     string `json:"tutorialId"`
}

// Instance list properties
type PlatformTutorialFetchResponse struct {
	// The category of the tutorial                                           
	Category                                           *string                `json:"category,omitempty"`
	// The markdown content of the tutorial                                   
	Content                                            string                 `json:"content"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The display order index                                                
	Index                                              *float64               `json:"index,omitempty"`
	// The URL to the official tutorial page                                  
	Link                                               *string                `json:"link,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// Tags associated with the tutorial                                      
	Tags                                               []string               `json:"tags,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformTutorialsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PlatformTutorialsListResponse struct {
	Items []PlatformTutorialsListResponseItem `json:"items"`
}

// Instance list properties
type PlatformTutorialsListResponseItem struct {
	// The category of the tutorial                                           
	Category                                           *string                `json:"category,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        string                 `json:"description"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The display order index                                                
	Index                                              float64                `json:"index"`
	// The URL to the official tutorial page                                  
	Link                                               string                 `json:"link"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// Tags associated with the tutorial                                      
	Tags                                               []string               `json:"tags"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PlatformTutorialsSearchRequest struct {
	// The search query to find relevant tutorials                       
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type PlatformTutorialsSearchResponse struct {
	Items []PlatformTutorialsSearchResponseItem `json:"items"`
}

// Instance list properties
type PlatformTutorialsSearchResponseItem struct {
	// The category of the tutorial                                                 
	Category                                                 *string                `json:"category,omitempty"`
	// The timestamp (ms) when the instance was created                             
	CreatedAt                                                float64                `json:"createdAt"`
	// The associated description                                                   
	Description                                              string                 `json:"description"`
	// An excerpt from the most relevant part of the tutorial                       
	Excerpt                                                  string                 `json:"excerpt"`
	// The instance ID                                                              
	ID                                                       string                 `json:"id"`
	// The display order index                                                      
	Index                                                    float64                `json:"index"`
	// The URL to the official tutorial page                                        
	Link                                                     string                 `json:"link"`
	// Meta data information                                                        
	Meta                                                     map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                          
	Name                                                     string                 `json:"name"`
	// The similarity score of the search result                                    
	Score                                                    float64                `json:"score"`
	// Tags associated with the tutorial                                            
	Tags                                                     []string               `json:"tags"`
	// The timestamp (ms) when the instance was updated                             
	UpdatedAt                                                float64                `json:"updatedAt"`
}

type PolicyDeleteParams struct {
	// The ID of the policy to delete       
	ID                               string `json:"id"`
}

type PolicyDeleteResponse struct {
	// The ID of the deleted policy       
	ID                             string `json:"id"`
}

type PolicyFetchParams struct {
	// The ID of the policy to fetch       
	ID                              string `json:"id"`
}

// Blueprint properties
type PolicyFetchResponse struct {
	// The ID of the blueprint                                
	BlueprintID                        *string                `json:"blueprintId,omitempty"`
	// The policy configuration as JSON                       
	Config                             map[string]interface{} `json:"config,omitempty"`
	// The associated description                             
	Description                        *string                `json:"description,omitempty"`
	// Meta data information                                  
	Meta                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                    
	Name                               *string                `json:"name,omitempty"`
	// The policy type                                        
	Type                               PolicyType             `json:"type"`
}

type PolicyUpdateParams struct {
	// The ID of the policy to update       
	ID                               string `json:"id"`
}

// Blueprint properties
type PolicyUpdateRequest struct {
	// The ID of the blueprint                                
	BlueprintID                        *string                `json:"blueprintId,omitempty"`
	// The policy configuration as JSON                       
	Config                             map[string]interface{} `json:"config,omitempty"`
	// The associated description                             
	Description                        *string                `json:"description,omitempty"`
	// Meta data information                                  
	Meta                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                    
	Name                               *string                `json:"name,omitempty"`
	// The policy type                                        
	Type                               *PolicyType            `json:"type,omitempty"`
}

type PolicyUpdateResponse struct {
	// The ID of the updated policy       
	ID                             string `json:"id"`
}

// Blueprint properties
type PolicyCreateRequest struct {
	// The ID of the blueprint                                
	BlueprintID                        *string                `json:"blueprintId,omitempty"`
	// The policy configuration as JSON                       
	Config                             map[string]interface{} `json:"config,omitempty"`
	// The associated description                             
	Description                        *string                `json:"description,omitempty"`
	// Meta data information                                  
	Meta                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                    
	Name                               *string                `json:"name,omitempty"`
	// The policy type                                        
	Type                               PolicyType             `json:"type"`
}

type PolicyCreateResponse struct {
	// The ID of the created policy       
	ID                             string `json:"id"`
}

type PoliciesListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PoliciesListResponse struct {
	Items []PoliciesListResponseItem `json:"items"`
}

// Blueprint properties
type PoliciesListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The policy configuration as JSON                                       
	Config                                             map[string]interface{} `json:"config,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The policy type                                                        
	Type                                               PolicyType             `json:"type"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PortalDeleteParams struct {
	// The ID of the portal to delete       
	PortalID                         string `json:"portalId"`
}

type PortalDeleteResponse struct {
	// The ID of the deleted portal       
	ID                             string `json:"id"`
}

type PortalFetchParams struct {
	// The ID of the portal to retrieve       
	PortalID                           string `json:"portalId"`
}

// Blueprint properties
type PortalFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The config of the portal                                               
	Config                                             map[string]interface{} `json:"config,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The slug of the portal                                                 
	Slug                                               *string                `json:"slug,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type PortalUpdateParams struct {
	PortalID string `json:"portalId"`
}

// Blueprint properties
type PortalUpdateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                 
	BlueprintID                         *string                `json:"blueprintId,omitempty"`
	// The config for the portal                               
	Config                              map[string]interface{} `json:"config,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The slug for the portal                                 
	Slug                                *string                `json:"slug,omitempty"`
}

type PortalUpdateResponse struct {
	// The ID of the updated portal       
	ID                             string `json:"id"`
}

// Blueprint properties
type PortalCreateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                 
	BlueprintID                         *string                `json:"blueprintId,omitempty"`
	// The config of the portal                                
	Config                              map[string]interface{} `json:"config,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The slug of the portal                                  
	Slug                                *string                `json:"slug,omitempty"`
}

type PortalCreateResponse struct {
	// The ID of the created portal       
	ID                             string `json:"id"`
}

type PortalsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type PortalsListResponse struct {
	Items []PortalsListResponseItem `json:"items"`
}

// Blueprint properties
type PortalsListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The config of the portal                                               
	Config                                             map[string]interface{} `json:"config,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The slug of the portal                                                 
	Slug                                               *string                `json:"slug,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type SecretAuthenticateParams struct {
	// The ID of the secret to authenticate       
	SecretID                               string `json:"secretId"`
}

type SecretAuthenticateResponse struct {
	// The ID of the secret to authenticate       
	ID                                     string `json:"id"`
	// The URL to authenticate the secret         
	URL                                    string `json:"url"`
}

type SecretDeleteParams struct {
	// The ID of the secret to delete       
	SecretID                         string `json:"secretId"`
}

type SecretDeleteResponse struct {
	// The ID of the deleted secret       
	ID                             string `json:"id"`
}

type SecretFetchParams struct {
	// The ID of the secret to retrieve       
	SecretID                           string `json:"secretId"`
}

// Blueprint properties
type SecretFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The config of the secret                                               
	Config                                             map[string]interface{} `json:"config,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The kind of the secret                                                 
	Kind                                               *SecretKind            `json:"kind,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The type of the secret                                                 
	Type                                               *SecretType            `json:"type,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The visibility of the secret                                           
	Visibility                                         *SecretVisibility      `json:"visibility,omitempty"`
}

type SecretRevokeParams struct {
	SecretID string `json:"secretId"`
}

type SecretRevokeResponse struct {
	// The ID of the revoked secret       
	ID                             string `json:"id"`
}

type SecretUpdateParams struct {
	SecretID string `json:"secretId"`
}

// Blueprint properties
type SecretUpdateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                 
	BlueprintID                         *string                `json:"blueprintId,omitempty"`
	// The config of the secret                                
	Config                              map[string]interface{} `json:"config,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// The kind of the secret                                  
	Kind                                *SecretKind            `json:"kind,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The type of the secret                                  
	Type                                *SecretType            `json:"type,omitempty"`
	// The value of the secret                                 
	Value                               *string                `json:"value,omitempty"`
	// The visibility of the secret                            
	Visibility                          *SecretVisibility      `json:"visibility,omitempty"`
}

type SecretUpdateResponse struct {
	// The ID of the updated secret       
	ID                             string `json:"id"`
}

type SecretVerifyParams struct {
	// The ID of the secret to be verified       
	SecretID                              string `json:"secretId"`
}

type SecretVerifyResponse struct {
	Action                          *SecretVerifyResponseAction `json:"action,omitempty"`
	// The ID of the verified secret                            
	ID                              string                      `json:"id"`
	// The status of the secret                                 
	Status                          Status                      `json:"status"`
}

// The action to take next
type SecretVerifyResponseAction struct {
	// The type of action to take                   
	Type                                 ActionType `json:"type"`
	// The URL to authenticate the secret           
	URL                                  string     `json:"url"`
}

// Blueprint properties
type SecretCreateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                 
	BlueprintID                         *string                `json:"blueprintId,omitempty"`
	// The config of the secret                                
	Config                              map[string]interface{} `json:"config,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// The kind of the secret                                  
	Kind                                *SecretKind            `json:"kind,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The type of the secret                                  
	Type                                *SecretType            `json:"type,omitempty"`
	// The value of the secret                                 
	Value                               *string                `json:"value,omitempty"`
	// The visibility of the secret                            
	Visibility                          *SecretVisibility      `json:"visibility,omitempty"`
}

type SecretCreateResponse struct {
	// The ID of the created secret       
	ID                             string `json:"id"`
}

type SecretsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type SecretsListResponse struct {
	Items []SecretsListResponseItem `json:"items"`
}

// Blueprint properties
type SecretsListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The config of the secret                                               
	Config                                             map[string]interface{} `json:"config,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The kind of the secret                                                 
	Kind                                               *SecretKind            `json:"kind,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The type of the secret                                                 
	Type                                               *SecretType            `json:"type,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The visibility of the secret                                           
	Visibility                                         *SecretVisibility      `json:"visibility,omitempty"`
}

type SkillsetAbilityDeleteParams struct {
	// The ID of the ability to delete       
	AbilityID                         string `json:"abilityId"`
	// The ID of the skillset                
	SkillsetID                        string `json:"skillsetId"`
}

type SkillsetAbilityDeleteResponse struct {
	// The ID of the deleted ability       
	ID                              string `json:"id"`
}

type SkillsetAbilityExecuteParams struct {
	// The ID of the ability to execute                    
	AbilityID                                       string `json:"abilityId"`
	// The ID of the skillset containing the ability       
	SkillsetID                                      string `json:"skillsetId"`
}

type SkillsetAbilityExecuteRequest struct {
	// The ID of the contact to associate with the execution                 
	ContactID                                                        *string `json:"contactId,omitempty"`
	// The input to process with the ability. This can be structured         
	// text such as JSON or YAML for precise parameter control, or           
	// unstructured natural language text. When unstructured text is         
	// provided, the system will automatically detect and extract the        
	// relevant parameters from the input.                                   
	Input                                                            *string `json:"input,omitempty"`
}

type SkillsetAbilityExecuteResponse struct {
	// Error message if execution failed                                          
	Error                                 *string                                 `json:"error,omitempty"`
	// Messages generated during execution                                        
	Messages                              []SkillsetAbilityExecuteResponseMessage `json:"messages,omitempty"`
	// The result of the ability execution                                        
	Result                                interface{}                             `json:"result"`
	// Usage information                                                          
	Usage                                 SkillsetAbilityExecuteResponseUsage     `json:"usage"`
}

// A message in the conversation
type SkillsetAbilityExecuteResponseMessage struct {
	// Meta data information                         
	Meta                      map[string]interface{} `json:"meta,omitempty"`
	// The text of the message                       
	Text                      string                 `json:"text"`
	// The type of the message                       
	Type                      MessageType            `json:"type"`
}

// Usage information
type SkillsetAbilityExecuteResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

type SkillsetAbilityFetchParams struct {
	// The ID of the ability to retrieve       
	AbilityID                           string `json:"abilityId"`
	// The ID of the skillset                  
	SkillsetID                          string `json:"skillsetId"`
}

// Blueprint properties
type SkillsetAbilityFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot associated with the ability                          
	BotID                                              *string                `json:"botId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        string                 `json:"description"`
	// The ID of the file associated with the ability                         
	FileID                                             *string                `json:"fileId,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// The instruction of the skillset ability                                
	Instruction                                        string                 `json:"instruction"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// The ID of the secret associated with the ability                       
	SecretID                                           *string                `json:"secretId,omitempty"`
	// The ID of the space associated with the ability                        
	SpaceID                                            *string                `json:"spaceId,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type SkillsetAbilityUpdateParams struct {
	AbilityID  string `json:"abilityId"`
	SkillsetID string `json:"skillsetId"`
}

// Blueprint properties
type SkillsetAbilityUpdateRequest struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot associated with the ability                          
	BotID                                              *string                `json:"botId,omitempty"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The ID of the file associated with the ability                         
	FileID                                             *string                `json:"fileId,omitempty"`
	// The text to update the ability with                                    
	Instruction                                        *string                `json:"instruction,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The ID of the secret associated with the ability                       
	SecretID                                           *string                `json:"secretId,omitempty"`
	// The ID of the space associated with the ability                        
	SpaceID                                            *string                `json:"spaceId,omitempty"`
}

type SkillsetAbilityUpdateResponse struct {
	// The ID of the updated ability       
	ID                              string `json:"id"`
}

type SkillsetAbilityCreateParams struct {
	SkillsetID string `json:"skillsetId"`
}

// Blueprint properties
type SkillsetAbilityCreateRequest struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot associated with the ability                          
	BotID                                              *string                `json:"botId,omitempty"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The ID of the file associated with the ability                         
	FileID                                             *string                `json:"fileId,omitempty"`
	// The instruction of the ability                                         
	Instruction                                        *string                `json:"instruction,omitempty"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The ID of the secret associated with the ability                       
	SecretID                                           *string                `json:"secretId,omitempty"`
	// The ID of the space associated with the ability                        
	SpaceID                                            *string                `json:"spaceId,omitempty"`
}

type SkillsetAbilityCreateResponse struct {
	// The ID of the created ability       
	ID                              string `json:"id"`
}

type SkillsetAbilitiesExportParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The ID of the skillset to export        
	SkillsetID                         string  `json:"skillsetId"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type SkillsetAbilitiesExportResponse struct {
	Items []SkillsetAbilitiesExportResponseItem `json:"items"`
}

// Blueprint properties
type SkillsetAbilitiesExportResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot associated with the ability                          
	BotID                                              *string                `json:"botId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        string                 `json:"description"`
	// The ID of the file associated with the ability                         
	FileID                                             *string                `json:"fileId,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	Instruction                                        string                 `json:"instruction"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// The ID of the secret associated with the ability                       
	SecretID                                           *string                `json:"secretId,omitempty"`
	// The ID of the space associated with the ability                        
	SpaceID                                            *string                `json:"spaceId,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type SkillsetAbilitiesListParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The ID of the skillset                  
	SkillsetID                         string  `json:"skillsetId"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type SkillsetAbilitiesListResponse struct {
	Items []SkillsetAbilitiesListResponseItem `json:"items"`
}

// Blueprint properties
type SkillsetAbilitiesListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The ID of the bot associated with the ability                          
	BotID                                              *string                `json:"botId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        string                 `json:"description"`
	// The ID of the file associated with the ability                         
	FileID                                             *string                `json:"fileId,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	Instruction                                        string                 `json:"instruction"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               string                 `json:"name"`
	// The ID of the secret associated with the ability                       
	SecretID                                           *string                `json:"secretId,omitempty"`
	// The ID of the space associated with the ability                        
	SpaceID                                            *string                `json:"spaceId,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type SkillsetDeleteParams struct {
	// The ID of the skillset to delete       
	SkillsetID                         string `json:"skillsetId"`
}

type SkillsetDeleteResponse struct {
	// The ID of the deleted skillset       
	ID                               string `json:"id"`
}

type SkillsetFetchParams struct {
	// The ID of the skillset to retrieve       
	SkillsetID                           string `json:"skillsetId"`
}

// Blueprint properties
type SkillsetFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The skillset visibility                                                
	Visibility                                         *SecretVisibility      `json:"visibility,omitempty"`
}

type SkillsetUpdateParams struct {
	SkillsetID string `json:"skillsetId"`
}

// Blueprint properties
type SkillsetUpdateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                 
	BlueprintID                         *string                `json:"blueprintId,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The skillset visibility                                 
	Visibility                          *SecretVisibility      `json:"visibility,omitempty"`
}

type SkillsetUpdateResponse struct {
	// The ID of the updated skillset       
	ID                               string `json:"id"`
}

// Blueprint properties
type SkillsetCreateRequest struct {
	// The unique alias for the instance                       
	Alias                               *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                 
	BlueprintID                         *string                `json:"blueprintId,omitempty"`
	// The associated description                              
	Description                         *string                `json:"description,omitempty"`
	// Meta data information                                   
	Meta                                map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                     
	Name                                *string                `json:"name,omitempty"`
	// The skillset visibility                                 
	Visibility                          *SecretVisibility      `json:"visibility,omitempty"`
}

type SkillsetCreateResponse struct {
	// The ID of the created skillset       
	ID                               string `json:"id"`
}

type SkillsetsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type SkillsetsListResponse struct {
	Items []SkillsetsListResponseItem `json:"items"`
}

// Blueprint properties
type SkillsetsListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
	// The skillset visibility                                                
	Visibility                                         *SecretVisibility      `json:"visibility,omitempty"`
}

type SpaceFetchParams struct {
	// The ID of the space to retrieve       
	SpaceID                           string `json:"spaceId"`
}

// Blueprint properties
type SpaceFetchResponse struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The contact associated with the space                                  
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type SpaceUpdateParams struct {
	SpaceID string `json:"spaceId"`
}

// Blueprint properties
type SpaceUpdateRequest struct {
	// The unique alias for the instance                           
	Alias                                   *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                     
	BlueprintID                             *string                `json:"blueprintId,omitempty"`
	// The contact associated with the space                       
	ContactID                               *string                `json:"contactId,omitempty"`
	// The associated description                                  
	Description                             *string                `json:"description,omitempty"`
	// Meta data information                                       
	Meta                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                         
	Name                                    *string                `json:"name,omitempty"`
}

type SpaceUpdateResponse struct {
	// The ID of the updated space       
	ID                            string `json:"id"`
}

// Blueprint properties
type SpaceCreateRequest struct {
	// The unique alias for the instance                           
	Alias                                   *string                `json:"alias,omitempty"`
	// The ID of the blueprint                                     
	BlueprintID                             *string                `json:"blueprintId,omitempty"`
	// The contact associated with the space                       
	ContactID                               *string                `json:"contactId,omitempty"`
	// The associated description                                  
	Description                             *string                `json:"description,omitempty"`
	// Meta data information                                       
	Meta                                    map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                         
	Name                                    *string                `json:"name,omitempty"`
}

type SpaceCreateResponse struct {
	// The ID of the created space       
	ID                            string `json:"id"`
}

type SpacesExportParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type SpacesExportResponse struct {
	Items []SpacesExportResponseItem `json:"items"`
}

// Blueprint properties
type SpacesExportResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The contact associated with the space                                  
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type SpacesListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type SpacesListResponse struct {
	Items []SpacesListResponseItem `json:"items"`
}

// Blueprint properties
type SpacesListResponseItem struct {
	// The ID of the blueprint                                                
	BlueprintID                                        *string                `json:"blueprintId,omitempty"`
	// The contact associated with the space                                  
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type TaskDeleteParams struct {
	// The ID of the task to delete       
	TaskID                         string `json:"taskId"`
}

type TaskDeleteResponse struct {
	// The ID of the deleted task       
	ID                           string `json:"id"`
}

type TaskFetchParams struct {
	// The ID of the task to retrieve       
	TaskID                           string `json:"taskId"`
}

// Instance list properties
type TaskFetchResponse struct {
	// The bot associated with the task                                       
	BotID                                              *string                `json:"botId,omitempty"`
	// The contact associated with the task                                   
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The task execution outcome                                             
	Outcome                                            *TaskOutcome           `json:"outcome,omitempty"`
	// The schedule of the task                                               
	Schedule                                           *string                `json:"schedule,omitempty"`
	// The task execution status                                              
	Status                                             *TaskStatus            `json:"status,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type TaskTriggerParams struct {
	TaskID string `json:"taskId"`
}

type TaskTriggerResponse struct {
	// The ID of the triggered task       
	ID                             string `json:"id"`
}

type TaskUpdateParams struct {
	TaskID string `json:"taskId"`
}

// Instance crud properties
type TaskUpdateRequest struct {
	// The bot associated with the task                                     
	BotID                                            *string                `json:"botId,omitempty"`
	// The contact associated with the task                                 
	ContactID                                        *string                `json:"contactId,omitempty"`
	// The associated description                                           
	Description                                      *string                `json:"description,omitempty"`
	// Meta data information                                                
	Meta                                             map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                  
	Name                                             *string                `json:"name,omitempty"`
	// The schedule of the task                                             
	Schedule                                         *string                `json:"schedule,omitempty"`
	// The session duration of the Widget integration                       
	SessionDuration                                  *float64               `json:"sessionDuration,omitempty"`
}

type TaskUpdateResponse struct {
	// The ID of the updated task       
	ID                           string `json:"id"`
}

// Instance crud properties
type TaskCreateRequest struct {
	// The bot associated with the task                                     
	BotID                                            *string                `json:"botId,omitempty"`
	// The contact associated with the task                                 
	ContactID                                        *string                `json:"contactId,omitempty"`
	// The associated description                                           
	Description                                      *string                `json:"description,omitempty"`
	// Meta data information                                                
	Meta                                             map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                  
	Name                                             *string                `json:"name,omitempty"`
	// The schedule of the task                                             
	Schedule                                         *string                `json:"schedule,omitempty"`
	// The session duration of the Widget integration                       
	SessionDuration                                  *float64               `json:"sessionDuration,omitempty"`
}

type TaskCreateResponse struct {
	// The ID of the created task       
	ID                           string `json:"id"`
}

type TasksExportParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type TasksExportResponse struct {
	Items []TasksExportResponseItem `json:"items"`
}

// Instance list properties
type TasksExportResponseItem struct {
	// The bot associated with the task                                       
	BotID                                              *string                `json:"botId,omitempty"`
	// The contact associated with the task                                   
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The schedule of the task                                               
	Schedule                                           *string                `json:"schedule,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type TasksListParams struct {
	// Filter by associated bot                                                 
	BotID                                                     *string           `json:"botId,omitempty"`
	// Filter by associated contact                                             
	ContactID                                                 *string           `json:"contactId,omitempty"`
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// Filter by task status                                                    
	Status                                                    *TaskStatus       `json:"status,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type TasksListResponse struct {
	Items []TasksListResponseItem `json:"items"`
}

// Instance list properties
type TasksListResponseItem struct {
	// The bot associated with the task                                       
	BotID                                              *string                `json:"botId,omitempty"`
	// The contact associated with the task                                   
	ContactID                                          *string                `json:"contactId,omitempty"`
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The task execution outcome                                             
	Outcome                                            *TaskOutcome           `json:"outcome,omitempty"`
	// The schedule of the task                                               
	Schedule                                           *string                `json:"schedule,omitempty"`
	// The task execution status                                              
	Status                                             *TaskStatus            `json:"status,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type TeamsListParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type TeamsListResponse struct {
	Items []TeamsListResponseItem `json:"items"`
}

// Instance list properties
type TeamsListResponseItem struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

type UsageFetchResponse struct {
	// The number of conversations the user has created                           
	Conversations                                      float64                    `json:"conversations"`
	// Database usage information                                                 
	Database                                           UsageFetchResponseDatabase `json:"database"`
	// The number of messages the user has sent                                   
	Messages                                           float64                    `json:"messages"`
	// The number of tokens the user has used                                     
	Tokens                                             float64                    `json:"tokens"`
}

// Database usage information
type UsageFetchResponseDatabase struct {
	// The number of abilities the user has created        
	Abilities                                      float64 `json:"abilities"`
	// The number of datasets the user has created         
	Datasets                                       float64 `json:"datasets"`
	// The number of files the user has created            
	Files                                          float64 `json:"files"`
	// The number of records the user has created          
	Records                                        float64 `json:"records"`
	// The number of skillsets the user has created        
	Skillsets                                      float64 `json:"skillsets"`
	// The number of users the user has created            
	Users                                          float64 `json:"users"`
}

type UsageSeriesFetchResponse struct {
	// The number of conversations the user has created                                  
	Conversations                                      []Conversation                    `json:"conversations"`
	// The number of messages the user has created                                       
	Messages                                           []UsageSeriesFetchResponseMessage `json:"messages"`
	// The number of tokens the user has used                                            
	Tokens                                             []TokenElement                    `json:"tokens"`
}

type Conversation struct {
	// The date of the data point                                 
	Date                                                  float64 `json:"date"`
	// The total number of conversations the user has used        
	Total                                                 float64 `json:"total"`
}

type UsageSeriesFetchResponseMessage struct {
	// The date of the data point                            
	Date                                             float64 `json:"date"`
	// The total number of messages the user has used        
	Total                                            float64 `json:"total"`
}

type TokenElement struct {
	// The date of the data point                          
	Date                                           float64 `json:"date"`
	// The total number of tokens the user has used        
	Total                                          float64 `json:"total"`
}

// A message in the conversation
type Message struct {
	// Meta data information                         
	Meta                      map[string]interface{} `json:"meta,omitempty"`
	// The text of the message                       
	Text                      string                 `json:"text"`
	// The type of the message                       
	Type                      MessageType            `json:"type"`
}

// Extracted entity from the message
type Entity struct {
	// Start offset                                   
	Begin                          float64            `json:"begin"`
	// End offset                                     
	End                            float64            `json:"end"`
	Replacement                    *EntityReplacement `json:"replacement,omitempty"`
	// The text value of the entity                   
	Text                           string             `json:"text"`
	// The entity type                                
	Type                           string             `json:"type"`
}

type EntityReplacement struct {
	// Start offset                             
	Begin                               float64 `json:"begin"`
	// End offset                               
	End                                 float64 `json:"end"`
	// The text value of the replacement        
	Text                                string  `json:"text"`
}

type DatasetFilterClass struct {
	Eq  *Eq      `json:"$eq"`
	Ne  *Eq      `json:"$ne"`
	Gt  *float64 `json:"$gt,omitempty"`
	Gte *float64 `json:"$gte,omitempty"`
	Lt  *float64 `json:"$lt,omitempty"`
	LTE *float64 `json:"$lte,omitempty"`
}

// Usage information
type Usage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

// Information about why the completion ended
type CompleteEnd struct {
	// The reason why the completion ended               
	Reason                                CompleteReason `json:"reason"`
}

// Execution limits to control conversation processing bounds
type ExecutionLimits struct {
	// Maximum number of function/tool calls. Controls how many total function calls can be made       
	// during the conversation.                                                                        
	Calls                                                                                       *int64 `json:"calls,omitempty"`
	// Maximum number of model continuations. Controls how many times the model can continue           
	// generating after reaching a stop condition.                                                     
	Continuations                                                                               *int64 `json:"continuations,omitempty"`
	// Maximum number of agentic iterations. Controls how many times the model can iterate             
	// through tool calls and responses.                                                               
	Iterations                                                                                  *int64 `json:"iterations,omitempty"`
}

// Limits information
type Limits struct {
	// The conversations limit                
	Conversations             *float64        `json:"conversations,omitempty"`
	// The database limits                    
	Database                  *LimitsDatabase `json:"database,omitempty"`
	// The messages limit                     
	Messages                  *float64        `json:"messages,omitempty"`
	// The tokens limit                       
	Tokens                    *float64        `json:"tokens,omitempty"`
}

// The database limits
type LimitsDatabase struct {
	// The abilities limit         
	Abilities             *float64 `json:"abilities,omitempty"`
	// The datasets limit          
	Datasets              *float64 `json:"datasets,omitempty"`
	// The files limit             
	Files                 *float64 `json:"files,omitempty"`
	// The records limit           
	Records               *float64 `json:"records,omitempty"`
	// The skillsets limit         
	Skillsets             *float64 `json:"skillsets,omitempty"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type BotRef struct {
	// The ID of the bot this configuration is using        
	BotID                                           *string `json:"botId,omitempty"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type BotConfig struct {
	// The backstory this configuration is using                 
	Backstory                                            *string `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using         
	DatasetID                                            *string `json:"datasetId,omitempty"`
	// A model definition                                        
	Model                                                *string `json:"model,omitempty"`
	// The moderation flag for this configuration                
	Moderation                                           *bool   `json:"moderation,omitempty"`
	// The privacy flag for this configuration                   
	Privacy                                              *bool   `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using        
	SkillsetID                                           *string `json:"skillsetId,omitempty"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type BotRefOrConfig struct {
	// The ID of the bot this configuration is using             
	BotID                                                *string `json:"botId,omitempty"`
	// The backstory this configuration is using                 
	Backstory                                            *string `json:"backstory,omitempty"`
	// The id of the dataset this configuration is using         
	DatasetID                                            *string `json:"datasetId,omitempty"`
	// A model definition                                        
	Model                                                *string `json:"model,omitempty"`
	// The moderation flag for this configuration                
	Moderation                                           *bool   `json:"moderation,omitempty"`
	// The privacy flag for this configuration                   
	Privacy                                              *bool   `json:"privacy,omitempty"`
	// The id of the skillset this configuration is using        
	SkillsetID                                           *string `json:"skillsetId,omitempty"`
}

// Blueprint properties
type BlueprintProps struct {
	// The ID of the blueprint        
	BlueprintID               *string `json:"blueprintId,omitempty"`
}

// Instance reference properties
type InstanceRefProperties struct {
	// The unique alias for the instance        
	Alias                               *string `json:"alias,omitempty"`
}

// Instance list properties
type InstanceMetaProps struct {
	// The timestamp (ms) when the instance was created        
	CreatedAt                                          float64 `json:"createdAt"`
	// The instance ID                                         
	ID                                                 string  `json:"id"`
	// The timestamp (ms) when the instance was updated        
	UpdatedAt                                          float64 `json:"updatedAt"`
}

// Instance crud properties
type InstanceCRUDProps struct {
	// The associated description                       
	Description                  *string                `json:"description,omitempty"`
	// Meta data information                            
	Meta                         map[string]interface{} `json:"meta,omitempty"`
	// The associated name                              
	Name                         *string                `json:"name,omitempty"`
}

// Instance list properties
type InstanceListProps struct {
	// The timestamp (ms) when the instance was created                       
	CreatedAt                                          float64                `json:"createdAt"`
	// The associated description                                             
	Description                                        *string                `json:"description,omitempty"`
	// The instance ID                                                        
	ID                                                 string                 `json:"id"`
	// Meta data information                                                  
	Meta                                               map[string]interface{} `json:"meta,omitempty"`
	// The associated name                                                    
	Name                                               *string                `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                       
	UpdatedAt                                          float64                `json:"updatedAt"`
}

// A JSON Schema object type definition (https://json-schema.org/). Represents an object
// schema with properties and validation rules.
type JSONSchemaObject struct {
	// The schema description                                  
	Description                         *string                `json:"description,omitempty"`
	// Object property definitions                             
	Properties                          map[string]interface{} `json:"properties"`
	// Required property names                                 
	Required                            []string               `json:"required,omitempty"`
	// The schema title                                        
	Title                               *string                `json:"title,omitempty"`
	// The schema type, must be "object"                       
	Type                                ParametersType         `json:"type"`
}

// An array of functions to be added to the conversation
type FunctionsDefinitionElement struct {
	// Configuration for when this function should be automatically called                                 
	Call                                                                     *FunctionsDefinitionCall      `json:"call,omitempty"`
	// The description of the function                                                                     
	Description                                                              string                        `json:"description"`
	// The name of the function (must be a valid JS identifier, max 64 chars)                              
	Name                                                                     string                        `json:"name"`
	// JSON Schema definition for the function parameters                                                  
	Parameters                                                               FunctionsDefinitionParameters `json:"parameters"`
	// The result of the function execution                                                                
	Result                                                                   *FunctionsDefinitionResult    `json:"result,omitempty"`
}

// Configuration for when this function should be automatically called
type FunctionsDefinitionCall struct {
	// If true, this function will be force-called at the end of the conversation        
	End                                                                            *bool `json:"end,omitempty"`
	// If true, this function will be force-called at the start of the conversation      
	Start                                                                          *bool `json:"start,omitempty"`
}

// JSON Schema definition for the function parameters
type FunctionsDefinitionParameters struct {
	// Object property definitions                             
	Properties                          map[string]interface{} `json:"properties"`
	// Required property names                                 
	Required                            []string               `json:"required,omitempty"`
	// The schema type, must be "object"                       
	Type                                ParametersType         `json:"type"`
}

// The result of the function execution
type FunctionsDefinitionResult struct {
	// The data returned by the function (can be any type)            
	Data                                                  interface{} `json:"data"`
	// The channel for streaming function results                     
	Channel                                               *string     `json:"channel,omitempty"`
}

// Extensions to enhance the bot's capabilities
type ExtensionsDefinition struct {
	// Additional backstory for the bot                                                
	Backstory                                           *string                        `json:"backstory,omitempty"`
	// Inline datasets to provide additional context                                   
	Datasets                                            []ExtensionsDefinitionDataset  `json:"datasets,omitempty"`
	// Feature flags to enable specific bot capabilities                               
	Features                                            []ExtensionsDefinitionFeature  `json:"features,omitempty"`
	// Inline skillsets to provide additional abilities                                
	Skillsets                                           []ExtensionsDefinitionSkillset `json:"skillsets,omitempty"`
}

type ExtensionsDefinitionDataset struct {
	// The description of the dataset                  
	Description                      *string           `json:"description,omitempty"`
	// The name of the dataset                         
	Name                             *string           `json:"name,omitempty"`
	// The records in the dataset                      
	Records                          []HilariousRecord `json:"records"`
}

type HilariousRecord struct {
	// Additional metadata for the record                       
	Meta                                 map[string]interface{} `json:"meta,omitempty"`
	// The text content of the record                           
	Text                                 string                 `json:"text"`
}

type ExtensionsDefinitionFeature struct {
	// The name of the feature to enable                                    
	Name                                             string                 `json:"name"`
	// Optional configuration options for the feature                       
	Options                                          map[string]interface{} `json:"options,omitempty"`
}

type ExtensionsDefinitionSkillset struct {
	// The abilities in the skillset                     
	Abilities                         []HilariousAbility `json:"abilities"`
	// The description of the skillset                   
	Description                       *string            `json:"description,omitempty"`
	// The name of the skillset                          
	Name                              *string            `json:"name,omitempty"`
}

type HilariousAbility struct {
	// The description of the ability                            
	Description                           string                 `json:"description"`
	// The instruction for the ability                           
	Instruction                           string                 `json:"instruction"`
	// Additional metadata for the ability                       
	Meta                                  map[string]interface{} `json:"meta,omitempty"`
	// The name of the ability                                   
	Name                                  string                 `json:"name"`
	// Optional secret ID for the ability                        
	SecretID                              *string                `json:"secretId,omitempty"`
}

// An item in the streaming completion response
type CompleteStreamingResponseItem struct {
	// The data for the event                                         
	//                                                                
	// A message in the conversation                                  
	Data                            Data                              `json:"data"`
	// The type of event                                              
	Type                            CompleteStreamingResponseItemType `json:"type"`
}

// The data for the event
//
// A message in the conversation
type Data struct {
	// The error message                                      
	Message                            *string                `json:"message,omitempty"`
	// The token generated                                    
	Token                              *string                `json:"token,omitempty"`
	// Meta data information                                  
	Meta                               map[string]interface{} `json:"meta,omitempty"`
	// The text of the message                                
	Text                               *string                `json:"text,omitempty"`
	// The type of the message                                
	Type                               *MessageType           `json:"type,omitempty"`
	// The number of input tokens used                        
	InputTokensUsed                    *float64               `json:"inputTokensUsed,omitempty"`
	// The model used                                         
	Model                              *string                `json:"model,omitempty"`
	// The number of output tokens used                       
	OutputTokensUsed                   *float64               `json:"outputTokensUsed,omitempty"`
}

// The order of the paginated items
type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)

// The blueprint visibility
//
// The bot visibility
//
// The dataset visibility
//
// The file visibility
//
// The visibility of the secret
//
// The skillset visibility
type SecretVisibility string

const (
	Private   SecretVisibility = "private"
	Protected SecretVisibility = "protected"
	Public    SecretVisibility = "public"
)

// The type of the message
type MessageType string

const (
	Backstory           MessageType = "backstory"
	Bot                 MessageType = "bot"
	Context             MessageType = "context"
	Instruction         MessageType = "instruction"
	MessageTypeActivity MessageType = "activity"
	Reasoning           MessageType = "reasoning"
	User                MessageType = "user"
)

// The type of action to take
type ActionType string

const (
	Authenticate ActionType = "authenticate"
)

// The status of the secret
type Status string

const (
	Authenticated   Status = "authenticated"
	Unauthenticated Status = "unauthenticated"
)

// The task execution outcome
type TaskOutcome string

const (
	Failure            TaskOutcome = "failure"
	Success            TaskOutcome = "success"
	TaskOutcomePending TaskOutcome = "pending"
)

// The task execution status
//
// Filter by task status
type TaskStatus string

const (
	Idle    TaskStatus = "idle"
	Running TaskStatus = "running"
)

// The schema type, must be "object"
type ParametersType string

const (
	Object ParametersType = "object"
)

// The reason why the completion ended
type CompleteReason string

const (
	CompleteReasonActivity CompleteReason = "activity"
	CompleteReasonError    CompleteReason = "error"
	Iteration              CompleteReason = "iteration"
	Length                 CompleteReason = "length"
	Stop                   CompleteReason = "stop"
)

// The dataset file attachment type
type DatasetFileAttachmentType string

const (
	Source DatasetFileAttachmentType = "source"
)

// The sync status of an integration
type SyncStatus string

const (
	SyncStatusError   SyncStatus = "error"
	SyncStatusPending SyncStatus = "pending"
	Synced            SyncStatus = "synced"
)

// The schedule
type Schedule string

const (
	Daily         Schedule = "daily"
	Halfhourly    Schedule = "halfhourly"
	Hourly        Schedule = "hourly"
	Monthly       Schedule = "monthly"
	Quarterhourly Schedule = "quarterhourly"
	ScheduleNever Schedule = "never"
	Weekly        Schedule = "weekly"
)

// The type of the example
type PlatformExampleFetchResponseType string

const (
	Blueprint   PlatformExampleFetchResponseType = "blueprint"
	Discord     PlatformExampleFetchResponseType = "discord"
	Email       PlatformExampleFetchResponseType = "email"
	Messenger   PlatformExampleFetchResponseType = "messenger"
	Project     PlatformExampleFetchResponseType = "project"
	Slack       PlatformExampleFetchResponseType = "slack"
	Telegram    PlatformExampleFetchResponseType = "telegram"
	Twilio      PlatformExampleFetchResponseType = "twilio"
	TypeTrigger PlatformExampleFetchResponseType = "trigger"
	Whatsapp    PlatformExampleFetchResponseType = "whatsapp"
	Widget      PlatformExampleFetchResponseType = "widget"
)

// The kind of the secret
type SecretKind string

const (
	Personal SecretKind = "personal"
	Shared   SecretKind = "shared"
)

// The type of the secret
type SecretType string

const (
	Basic     SecretType = "basic"
	Bearer    SecretType = "bearer"
	Oauth     SecretType = "oauth"
	Plain     SecretType = "plain"
	Reference SecretType = "reference"
	Template  SecretType = "template"
)

// The policy type
type PolicyType string

const (
	Retention PolicyType = "retention"
)

// The type of the trigger
type Trigger string

const (
	Automatic    Trigger = "automatic"
	TriggerNever Trigger = "never"
)

// The type of event
type CompleteStreamingResponseItemType string

const (
	CompleteBegin              CompleteStreamingResponseItemType = "completeBegin"
	ReasoningToken             CompleteStreamingResponseItemType = "reasoningToken"
	Token                      CompleteStreamingResponseItemType = "token"
	TypeCompleteEnd            CompleteStreamingResponseItemType = "completeEnd"
	TypeError                  CompleteStreamingResponseItemType = "error"
	TypeMessage                CompleteStreamingResponseItemType = "message"
	TypeUsage                  CompleteStreamingResponseItemType = "usage"
	WaitForChannelMessageBegin CompleteStreamingResponseItemType = "waitForChannelMessageBegin"
	WaitForChannelMessageEnd   CompleteStreamingResponseItemType = "waitForChannelMessageEnd"
)

type ConversationAttachmentUploadRequestFile struct {
	PurpleFile *PurpleFile
	String     *string
}

func (x *ConversationAttachmentUploadRequestFile) UnmarshalJSON(data []byte) error {
	x.PurpleFile = nil
	var c PurpleFile
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleFile = &c
	}
	return nil
}

func (x *ConversationAttachmentUploadRequestFile) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleFile != nil, x.PurpleFile, false, nil, false, nil, false)
}

// The contact ID to associate with this conversation
type ConversationCompleteRequestContactID struct {
	PurpleContactID *PurpleContactID
	String          *string
}

func (x *ConversationCompleteRequestContactID) UnmarshalJSON(data []byte) error {
	x.PurpleContactID = nil
	var c PurpleContactID
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleContactID = &c
	}
	return nil
}

func (x *ConversationCompleteRequestContactID) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleContactID != nil, x.PurpleContactID, false, nil, false, nil, false)
}

// The contact ID to associate with this conversation
type ConversationDispatchRequestContactID struct {
	FluffyContactID *FluffyContactID
	String          *string
}

func (x *ConversationDispatchRequestContactID) UnmarshalJSON(data []byte) error {
	x.FluffyContactID = nil
	var c FluffyContactID
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyContactID = &c
	}
	return nil
}

func (x *ConversationDispatchRequestContactID) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyContactID != nil, x.FluffyContactID, false, nil, false, nil, false)
}

type FilterValue struct {
	Bool        *bool
	Double      *float64
	FilterClass *FilterClass
	String      *string
}

func (x *FilterValue) UnmarshalJSON(data []byte) error {
	x.FilterClass = nil
	var c FilterClass
	object, err := unmarshalUnion(data, nil, &x.Double, &x.Bool, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FilterClass = &c
	}
	return nil
}

func (x *FilterValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, x.Bool, x.String, false, nil, x.FilterClass != nil, x.FilterClass, false, nil, false, nil, false)
}

type Eq struct {
	Bool   *bool
	Double *float64
	String *string
}

func (x *Eq) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, &x.Double, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Eq) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

type FileUploadRequestFile struct {
	FluffyFile *FluffyFile
	String     *string
}

func (x *FileUploadRequestFile) UnmarshalJSON(data []byte) error {
	x.FluffyFile = nil
	var c FluffyFile
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyFile = &c
	}
	return nil
}

func (x *FileUploadRequestFile) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyFile != nil, x.FluffyFile, false, nil, false, nil, false)
}

type DatasetFilterValue struct {
	Bool               *bool
	DatasetFilterClass *DatasetFilterClass
	Double             *float64
	String             *string
}

func (x *DatasetFilterValue) UnmarshalJSON(data []byte) error {
	x.DatasetFilterClass = nil
	var c DatasetFilterClass
	object, err := unmarshalUnion(data, nil, &x.Double, &x.Bool, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.DatasetFilterClass = &c
	}
	return nil
}

func (x *DatasetFilterValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, x.Bool, x.String, false, nil, x.DatasetFilterClass != nil, x.DatasetFilterClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
			*pi = nil
	}
	if pf != nil {
			*pf = nil
	}
	if pb != nil {
			*pb = nil
	}
	if ps != nil {
			*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
			return false, err
	}

	switch v := tok.(type) {
	case json.Number:
			if pi != nil {
					i, err := v.Int64()
					if err == nil {
							*pi = &i
							return false, nil
					}
			}
			if pf != nil {
					f, err := v.Float64()
					if err == nil {
							*pf = &f
							return false, nil
					}
					return false, errors.New("Unparsable number")
			}
			return false, errors.New("Union does not contain number")
	case float64:
			return false, errors.New("Decoder should not return float64")
	case bool:
			if pb != nil {
					*pb = &v
					return false, nil
			}
			return false, errors.New("Union does not contain bool")
	case string:
			if haveEnum {
					return false, json.Unmarshal(data, pe)
			}
			if ps != nil {
					*ps = &v
					return false, nil
			}
			return false, errors.New("Union does not contain string")
	case nil:
			if nullable {
					return false, nil
			}
			return false, errors.New("Union does not contain null")
	case json.Delim:
			if v == '{' {
					if haveObject {
							return true, json.Unmarshal(data, pc)
					}
					if haveMap {
							return false, json.Unmarshal(data, pm)
					}
					return false, errors.New("Union does not contain object")
			}
			if v == '[' {
					if haveArray {
							return false, json.Unmarshal(data, pa)
					}
					return false, errors.New("Union does not contain array")
			}
			return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")
}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
			return json.Marshal(*pi)
	}
	if pf != nil {
			return json.Marshal(*pf)
	}
	if pb != nil {
			return json.Marshal(*pb)
	}
	if ps != nil {
			return json.Marshal(*ps)
	}
	if haveArray {
			return json.Marshal(pa)
	}
	if haveObject {
			return json.Marshal(pc)
	}
	if haveMap {
			return json.Marshal(pm)
	}
	if haveEnum {
			return json.Marshal(pe)
	}
	if nullable {
			return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
