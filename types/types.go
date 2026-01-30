// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    graphqlRequest, err := UnmarshalGraphqlRequest(bytes)
//    bytes, err = graphqlRequest.Marshal()
//
//    graphqlResponse, err := UnmarshalGraphqlResponse(bytes)
//    bytes, err = graphqlResponse.Marshal()
//
//    listPlatformReportsParams, err := UnmarshalListPlatformReportsParams(bytes)
//    bytes, err = listPlatformReportsParams.Marshal()
//
//    listPlatformReportsResponse, err := UnmarshalListPlatformReportsResponse(bytes)
//    bytes, err = listPlatformReportsResponse.Marshal()
//
//    generateReportParams, err := UnmarshalGenerateReportParams(bytes)
//    bytes, err = generateReportParams.Marshal()
//
//    generateReportRequest, err := UnmarshalGenerateReportRequest(bytes)
//    bytes, err = generateReportRequest.Marshal()
//
//    generateReportResponse, err := UnmarshalGenerateReportResponse(bytes)
//    bytes, err = generateReportResponse.Marshal()
//
//    generateReportsRequest, err := UnmarshalGenerateReportsRequest(bytes)
//    bytes, err = generateReportsRequest.Marshal()
//
//    generateReportsResponse, err := UnmarshalGenerateReportsResponse(bytes)
//    bytes, err = generateReportsResponse.Marshal()
//
//    cloneBlueprintParams, err := UnmarshalCloneBlueprintParams(bytes)
//    bytes, err = cloneBlueprintParams.Marshal()
//
//    cloneBlueprintRequest, err := UnmarshalCloneBlueprintRequest(bytes)
//    bytes, err = cloneBlueprintRequest.Marshal()
//
//    cloneBlueprintResponse, err := UnmarshalCloneBlueprintResponse(bytes)
//    bytes, err = cloneBlueprintResponse.Marshal()
//
//    deleteBlueprintParams, err := UnmarshalDeleteBlueprintParams(bytes)
//    bytes, err = deleteBlueprintParams.Marshal()
//
//    deleteBlueprintRequest, err := UnmarshalDeleteBlueprintRequest(bytes)
//    bytes, err = deleteBlueprintRequest.Marshal()
//
//    deleteBlueprintResponse, err := UnmarshalDeleteBlueprintResponse(bytes)
//    bytes, err = deleteBlueprintResponse.Marshal()
//
//    fetchBlueprintParams, err := UnmarshalFetchBlueprintParams(bytes)
//    bytes, err = fetchBlueprintParams.Marshal()
//
//    fetchBlueprintResponse, err := UnmarshalFetchBlueprintResponse(bytes)
//    bytes, err = fetchBlueprintResponse.Marshal()
//
//    listBlueprintResourcesParams, err := UnmarshalListBlueprintResourcesParams(bytes)
//    bytes, err = listBlueprintResourcesParams.Marshal()
//
//    listBlueprintResourcesResponse, err := UnmarshalListBlueprintResourcesResponse(bytes)
//    bytes, err = listBlueprintResourcesResponse.Marshal()
//
//    updateBlueprintParams, err := UnmarshalUpdateBlueprintParams(bytes)
//    bytes, err = updateBlueprintParams.Marshal()
//
//    updateBlueprintRequest, err := UnmarshalUpdateBlueprintRequest(bytes)
//    bytes, err = updateBlueprintRequest.Marshal()
//
//    updateBlueprintResponse, err := UnmarshalUpdateBlueprintResponse(bytes)
//    bytes, err = updateBlueprintResponse.Marshal()
//
//    createBlueprintRequest, err := UnmarshalCreateBlueprintRequest(bytes)
//    bytes, err = createBlueprintRequest.Marshal()
//
//    createBlueprintResponse, err := UnmarshalCreateBlueprintResponse(bytes)
//    bytes, err = createBlueprintResponse.Marshal()
//
//    listBlueprintsParams, err := UnmarshalListBlueprintsParams(bytes)
//    bytes, err = listBlueprintsParams.Marshal()
//
//    listBlueprintsResponse, err := UnmarshalListBlueprintsResponse(bytes)
//    bytes, err = listBlueprintsResponse.Marshal()
//
//    cloneBotParams, err := UnmarshalCloneBotParams(bytes)
//    bytes, err = cloneBotParams.Marshal()
//
//    cloneBotRequest, err := UnmarshalCloneBotRequest(bytes)
//    bytes, err = cloneBotRequest.Marshal()
//
//    cloneBotResponse, err := UnmarshalCloneBotResponse(bytes)
//    bytes, err = cloneBotResponse.Marshal()
//
//    deleteBotParams, err := UnmarshalDeleteBotParams(bytes)
//    bytes, err = deleteBotParams.Marshal()
//
//    deleteBotRequest, err := UnmarshalDeleteBotRequest(bytes)
//    bytes, err = deleteBotRequest.Marshal()
//
//    deleteBotResponse, err := UnmarshalDeleteBotResponse(bytes)
//    bytes, err = deleteBotResponse.Marshal()
//
//    downvoteBotParams, err := UnmarshalDownvoteBotParams(bytes)
//    bytes, err = downvoteBotParams.Marshal()
//
//    downvoteBotRequest, err := UnmarshalDownvoteBotRequest(bytes)
//    bytes, err = downvoteBotRequest.Marshal()
//
//    downvoteBotResponse, err := UnmarshalDownvoteBotResponse(bytes)
//    bytes, err = downvoteBotResponse.Marshal()
//
//    fetchBotParams, err := UnmarshalFetchBotParams(bytes)
//    bytes, err = fetchBotParams.Marshal()
//
//    fetchBotResponse, err := UnmarshalFetchBotResponse(bytes)
//    bytes, err = fetchBotResponse.Marshal()
//
//    searchBotMemoryParams, err := UnmarshalSearchBotMemoryParams(bytes)
//    bytes, err = searchBotMemoryParams.Marshal()
//
//    searchBotMemoryRequest, err := UnmarshalSearchBotMemoryRequest(bytes)
//    bytes, err = searchBotMemoryRequest.Marshal()
//
//    searchBotMemoryResponse, err := UnmarshalSearchBotMemoryResponse(bytes)
//    bytes, err = searchBotMemoryResponse.Marshal()
//
//    createBotSessionParams, err := UnmarshalCreateBotSessionParams(bytes)
//    bytes, err = createBotSessionParams.Marshal()
//
//    createBotSessionRequest, err := UnmarshalCreateBotSessionRequest(bytes)
//    bytes, err = createBotSessionRequest.Marshal()
//
//    createBotSessionResponse, err := UnmarshalCreateBotSessionResponse(bytes)
//    bytes, err = createBotSessionResponse.Marshal()
//
//    updateBotParams, err := UnmarshalUpdateBotParams(bytes)
//    bytes, err = updateBotParams.Marshal()
//
//    updateBotRequest, err := UnmarshalUpdateBotRequest(bytes)
//    bytes, err = updateBotRequest.Marshal()
//
//    updateBotResponse, err := UnmarshalUpdateBotResponse(bytes)
//    bytes, err = updateBotResponse.Marshal()
//
//    upvoteBotParams, err := UnmarshalUpvoteBotParams(bytes)
//    bytes, err = upvoteBotParams.Marshal()
//
//    upvoteBotRequest, err := UnmarshalUpvoteBotRequest(bytes)
//    bytes, err = upvoteBotRequest.Marshal()
//
//    upvoteBotResponse, err := UnmarshalUpvoteBotResponse(bytes)
//    bytes, err = upvoteBotResponse.Marshal()
//
//    fetchBotUsageParams, err := UnmarshalFetchBotUsageParams(bytes)
//    bytes, err = fetchBotUsageParams.Marshal()
//
//    fetchBotUsageResponse, err := UnmarshalFetchBotUsageResponse(bytes)
//    bytes, err = fetchBotUsageResponse.Marshal()
//
//    createBotRequest, err := UnmarshalCreateBotRequest(bytes)
//    bytes, err = createBotRequest.Marshal()
//
//    createBotResponse, err := UnmarshalCreateBotResponse(bytes)
//    bytes, err = createBotResponse.Marshal()
//
//    listBotsParams, err := UnmarshalListBotsParams(bytes)
//    bytes, err = listBotsParams.Marshal()
//
//    listBotsResponse, err := UnmarshalListBotsResponse(bytes)
//    bytes, err = listBotsResponse.Marshal()
//
//    publishChannelMessageParams, err := UnmarshalPublishChannelMessageParams(bytes)
//    bytes, err = publishChannelMessageParams.Marshal()
//
//    publishChannelMessageRequest, err := UnmarshalPublishChannelMessageRequest(bytes)
//    bytes, err = publishChannelMessageRequest.Marshal()
//
//    publishChannelMessageResponse, err := UnmarshalPublishChannelMessageResponse(bytes)
//    bytes, err = publishChannelMessageResponse.Marshal()
//
//    subscribeChannelMessagesParams, err := UnmarshalSubscribeChannelMessagesParams(bytes)
//    bytes, err = subscribeChannelMessagesParams.Marshal()
//
//    subscribeChannelMessagesRequest, err := UnmarshalSubscribeChannelMessagesRequest(bytes)
//    bytes, err = subscribeChannelMessagesRequest.Marshal()
//
//    listContactConversationsParams, err := UnmarshalListContactConversationsParams(bytes)
//    bytes, err = listContactConversationsParams.Marshal()
//
//    listContactConversationsResponse, err := UnmarshalListContactConversationsResponse(bytes)
//    bytes, err = listContactConversationsResponse.Marshal()
//
//    deleteContactParams, err := UnmarshalDeleteContactParams(bytes)
//    bytes, err = deleteContactParams.Marshal()
//
//    deleteContactRequest, err := UnmarshalDeleteContactRequest(bytes)
//    bytes, err = deleteContactRequest.Marshal()
//
//    deleteContactResponse, err := UnmarshalDeleteContactResponse(bytes)
//    bytes, err = deleteContactResponse.Marshal()
//
//    fetchContactParams, err := UnmarshalFetchContactParams(bytes)
//    bytes, err = fetchContactParams.Marshal()
//
//    fetchContactResponse, err := UnmarshalFetchContactResponse(bytes)
//    bytes, err = fetchContactResponse.Marshal()
//
//    listContactMemoriesParams, err := UnmarshalListContactMemoriesParams(bytes)
//    bytes, err = listContactMemoriesParams.Marshal()
//
//    listContactMemoriesResponse, err := UnmarshalListContactMemoriesResponse(bytes)
//    bytes, err = listContactMemoriesResponse.Marshal()
//
//    searchContactMemoryParams, err := UnmarshalSearchContactMemoryParams(bytes)
//    bytes, err = searchContactMemoryParams.Marshal()
//
//    searchContactMemoryRequest, err := UnmarshalSearchContactMemoryRequest(bytes)
//    bytes, err = searchContactMemoryRequest.Marshal()
//
//    searchContactMemoryResponse, err := UnmarshalSearchContactMemoryResponse(bytes)
//    bytes, err = searchContactMemoryResponse.Marshal()
//
//    authenticateContactSecretParams, err := UnmarshalAuthenticateContactSecretParams(bytes)
//    bytes, err = authenticateContactSecretParams.Marshal()
//
//    authenticateContactSecretRequest, err := UnmarshalAuthenticateContactSecretRequest(bytes)
//    bytes, err = authenticateContactSecretRequest.Marshal()
//
//    authenticateContactSecretResponse, err := UnmarshalAuthenticateContactSecretResponse(bytes)
//    bytes, err = authenticateContactSecretResponse.Marshal()
//
//    revokeContactSecretParams, err := UnmarshalRevokeContactSecretParams(bytes)
//    bytes, err = revokeContactSecretParams.Marshal()
//
//    revokeContactSecretRequest, err := UnmarshalRevokeContactSecretRequest(bytes)
//    bytes, err = revokeContactSecretRequest.Marshal()
//
//    revokeContactSecretResponse, err := UnmarshalRevokeContactSecretResponse(bytes)
//    bytes, err = revokeContactSecretResponse.Marshal()
//
//    verifyContactSecretParams, err := UnmarshalVerifyContactSecretParams(bytes)
//    bytes, err = verifyContactSecretParams.Marshal()
//
//    verifyContactSecretRequest, err := UnmarshalVerifyContactSecretRequest(bytes)
//    bytes, err = verifyContactSecretRequest.Marshal()
//
//    verifyContactSecretResponse, err := UnmarshalVerifyContactSecretResponse(bytes)
//    bytes, err = verifyContactSecretResponse.Marshal()
//
//    listContactSecretsParams, err := UnmarshalListContactSecretsParams(bytes)
//    bytes, err = listContactSecretsParams.Marshal()
//
//    listContactSecretsResponse, err := UnmarshalListContactSecretsResponse(bytes)
//    bytes, err = listContactSecretsResponse.Marshal()
//
//    listContactSpacesParams, err := UnmarshalListContactSpacesParams(bytes)
//    bytes, err = listContactSpacesParams.Marshal()
//
//    listContactSpacesResponse, err := UnmarshalListContactSpacesResponse(bytes)
//    bytes, err = listContactSpacesResponse.Marshal()
//
//    listContactTasksParams, err := UnmarshalListContactTasksParams(bytes)
//    bytes, err = listContactTasksParams.Marshal()
//
//    listContactTasksResponse, err := UnmarshalListContactTasksResponse(bytes)
//    bytes, err = listContactTasksResponse.Marshal()
//
//    updateContactParams, err := UnmarshalUpdateContactParams(bytes)
//    bytes, err = updateContactParams.Marshal()
//
//    updateContactRequest, err := UnmarshalUpdateContactRequest(bytes)
//    bytes, err = updateContactRequest.Marshal()
//
//    updateContactResponse, err := UnmarshalUpdateContactResponse(bytes)
//    bytes, err = updateContactResponse.Marshal()
//
//    createContactRequest, err := UnmarshalCreateContactRequest(bytes)
//    bytes, err = createContactRequest.Marshal()
//
//    createContactResponse, err := UnmarshalCreateContactResponse(bytes)
//    bytes, err = createContactResponse.Marshal()
//
//    ensureContactRequest, err := UnmarshalEnsureContactRequest(bytes)
//    bytes, err = ensureContactRequest.Marshal()
//
//    ensureContactResponse, err := UnmarshalEnsureContactResponse(bytes)
//    bytes, err = ensureContactResponse.Marshal()
//
//    exportContactsParams, err := UnmarshalExportContactsParams(bytes)
//    bytes, err = exportContactsParams.Marshal()
//
//    exportContactsResponse, err := UnmarshalExportContactsResponse(bytes)
//    bytes, err = exportContactsResponse.Marshal()
//
//    listContactsParams, err := UnmarshalListContactsParams(bytes)
//    bytes, err = listContactsParams.Marshal()
//
//    listContactsResponse, err := UnmarshalListContactsResponse(bytes)
//    bytes, err = listContactsResponse.Marshal()
//
//    uploadConversationAttachmentParams, err := UnmarshalUploadConversationAttachmentParams(bytes)
//    bytes, err = uploadConversationAttachmentParams.Marshal()
//
//    uploadConversationAttachmentRequest, err := UnmarshalUploadConversationAttachmentRequest(bytes)
//    bytes, err = uploadConversationAttachmentRequest.Marshal()
//
//    uploadConversationAttachmentResponse, err := UnmarshalUploadConversationAttachmentResponse(bytes)
//    bytes, err = uploadConversationAttachmentResponse.Marshal()
//
//    completeConversationMessageParams, err := UnmarshalCompleteConversationMessageParams(bytes)
//    bytes, err = completeConversationMessageParams.Marshal()
//
//    completeConversationMessageRequest, err := UnmarshalCompleteConversationMessageRequest(bytes)
//    bytes, err = completeConversationMessageRequest.Marshal()
//
//    completeConversationMessageResponse, err := UnmarshalCompleteConversationMessageResponse(bytes)
//    bytes, err = completeConversationMessageResponse.Marshal()
//
//    upsertConversationContactParams, err := UnmarshalUpsertConversationContactParams(bytes)
//    bytes, err = upsertConversationContactParams.Marshal()
//
//    upsertConversationContactRequest, err := UnmarshalUpsertConversationContactRequest(bytes)
//    bytes, err = upsertConversationContactRequest.Marshal()
//
//    upsertConversationContactResponse, err := UnmarshalUpsertConversationContactResponse(bytes)
//    bytes, err = upsertConversationContactResponse.Marshal()
//
//    deleteConversationParams, err := UnmarshalDeleteConversationParams(bytes)
//    bytes, err = deleteConversationParams.Marshal()
//
//    deleteConversationRequest, err := UnmarshalDeleteConversationRequest(bytes)
//    bytes, err = deleteConversationRequest.Marshal()
//
//    deleteConversationResponse, err := UnmarshalDeleteConversationResponse(bytes)
//    bytes, err = deleteConversationResponse.Marshal()
//
//    dispatchStatefulConversationRequest, err := UnmarshalDispatchStatefulConversationRequest(bytes)
//    bytes, err = dispatchStatefulConversationRequest.Marshal()
//
//    dispatchStatefulConversationResponse, err := UnmarshalDispatchStatefulConversationResponse(bytes)
//    bytes, err = dispatchStatefulConversationResponse.Marshal()
//
//    downvoteConversationParams, err := UnmarshalDownvoteConversationParams(bytes)
//    bytes, err = downvoteConversationParams.Marshal()
//
//    downvoteConversationRequest, err := UnmarshalDownvoteConversationRequest(bytes)
//    bytes, err = downvoteConversationRequest.Marshal()
//
//    downvoteConversationResponse, err := UnmarshalDownvoteConversationResponse(bytes)
//    bytes, err = downvoteConversationResponse.Marshal()
//
//    fetchConversationParams, err := UnmarshalFetchConversationParams(bytes)
//    bytes, err = fetchConversationParams.Marshal()
//
//    fetchConversationResponse, err := UnmarshalFetchConversationResponse(bytes)
//    bytes, err = fetchConversationResponse.Marshal()
//
//    deleteConversationMessageParams, err := UnmarshalDeleteConversationMessageParams(bytes)
//    bytes, err = deleteConversationMessageParams.Marshal()
//
//    deleteConversationMessageRequest, err := UnmarshalDeleteConversationMessageRequest(bytes)
//    bytes, err = deleteConversationMessageRequest.Marshal()
//
//    deleteConversationMessageResponse, err := UnmarshalDeleteConversationMessageResponse(bytes)
//    bytes, err = deleteConversationMessageResponse.Marshal()
//
//    downvoteConversationMessageParams, err := UnmarshalDownvoteConversationMessageParams(bytes)
//    bytes, err = downvoteConversationMessageParams.Marshal()
//
//    downvoteConversationMessageRequest, err := UnmarshalDownvoteConversationMessageRequest(bytes)
//    bytes, err = downvoteConversationMessageRequest.Marshal()
//
//    downvoteConversationMessageResponse, err := UnmarshalDownvoteConversationMessageResponse(bytes)
//    bytes, err = downvoteConversationMessageResponse.Marshal()
//
//    fetchConversationMessageParams, err := UnmarshalFetchConversationMessageParams(bytes)
//    bytes, err = fetchConversationMessageParams.Marshal()
//
//    fetchConversationMessageResponse, err := UnmarshalFetchConversationMessageResponse(bytes)
//    bytes, err = fetchConversationMessageResponse.Marshal()
//
//    synthesizeConversationMessageParams, err := UnmarshalSynthesizeConversationMessageParams(bytes)
//    bytes, err = synthesizeConversationMessageParams.Marshal()
//
//    synthesizeConversationMessageRequest, err := UnmarshalSynthesizeConversationMessageRequest(bytes)
//    bytes, err = synthesizeConversationMessageRequest.Marshal()
//
//    synthesizeConversationMessageResponse, err := UnmarshalSynthesizeConversationMessageResponse(bytes)
//    bytes, err = synthesizeConversationMessageResponse.Marshal()
//
//    updateConversationMessageParams, err := UnmarshalUpdateConversationMessageParams(bytes)
//    bytes, err = updateConversationMessageParams.Marshal()
//
//    updateConversationMessageRequest, err := UnmarshalUpdateConversationMessageRequest(bytes)
//    bytes, err = updateConversationMessageRequest.Marshal()
//
//    updateConversationMessageResponse, err := UnmarshalUpdateConversationMessageResponse(bytes)
//    bytes, err = updateConversationMessageResponse.Marshal()
//
//    upvoteConversationMessageParams, err := UnmarshalUpvoteConversationMessageParams(bytes)
//    bytes, err = upvoteConversationMessageParams.Marshal()
//
//    upvoteConversationMessageRequest, err := UnmarshalUpvoteConversationMessageRequest(bytes)
//    bytes, err = upvoteConversationMessageRequest.Marshal()
//
//    upvoteConversationMessageResponse, err := UnmarshalUpvoteConversationMessageResponse(bytes)
//    bytes, err = upvoteConversationMessageResponse.Marshal()
//
//    createConversationMessageParams, err := UnmarshalCreateConversationMessageParams(bytes)
//    bytes, err = createConversationMessageParams.Marshal()
//
//    createConversationMessageRequest, err := UnmarshalCreateConversationMessageRequest(bytes)
//    bytes, err = createConversationMessageRequest.Marshal()
//
//    createConversationMessageResponse, err := UnmarshalCreateConversationMessageResponse(bytes)
//    bytes, err = createConversationMessageResponse.Marshal()
//
//    listConversationMessagesParams, err := UnmarshalListConversationMessagesParams(bytes)
//    bytes, err = listConversationMessagesParams.Marshal()
//
//    listConversationMessagesResponse, err := UnmarshalListConversationMessagesResponse(bytes)
//    bytes, err = listConversationMessagesResponse.Marshal()
//
//    receiveConversationMessageParams, err := UnmarshalReceiveConversationMessageParams(bytes)
//    bytes, err = receiveConversationMessageParams.Marshal()
//
//    receiveConversationMessageRequest, err := UnmarshalReceiveConversationMessageRequest(bytes)
//    bytes, err = receiveConversationMessageRequest.Marshal()
//
//    receiveConversationMessageResponse, err := UnmarshalReceiveConversationMessageResponse(bytes)
//    bytes, err = receiveConversationMessageResponse.Marshal()
//
//    sendConversationMessageParams, err := UnmarshalSendConversationMessageParams(bytes)
//    bytes, err = sendConversationMessageParams.Marshal()
//
//    sendConversationMessageRequest, err := UnmarshalSendConversationMessageRequest(bytes)
//    bytes, err = sendConversationMessageRequest.Marshal()
//
//    sendConversationMessageResponse, err := UnmarshalSendConversationMessageResponse(bytes)
//    bytes, err = sendConversationMessageResponse.Marshal()
//
//    createConversationSessionParams, err := UnmarshalCreateConversationSessionParams(bytes)
//    bytes, err = createConversationSessionParams.Marshal()
//
//    createConversationSessionRequest, err := UnmarshalCreateConversationSessionRequest(bytes)
//    bytes, err = createConversationSessionRequest.Marshal()
//
//    createConversationSessionResponse, err := UnmarshalCreateConversationSessionResponse(bytes)
//    bytes, err = createConversationSessionResponse.Marshal()
//
//    updateConversationParams, err := UnmarshalUpdateConversationParams(bytes)
//    bytes, err = updateConversationParams.Marshal()
//
//    updateConversationRequest, err := UnmarshalUpdateConversationRequest(bytes)
//    bytes, err = updateConversationRequest.Marshal()
//
//    updateConversationResponse, err := UnmarshalUpdateConversationResponse(bytes)
//    bytes, err = updateConversationResponse.Marshal()
//
//    upvoteConversationParams, err := UnmarshalUpvoteConversationParams(bytes)
//    bytes, err = upvoteConversationParams.Marshal()
//
//    upvoteConversationRequest, err := UnmarshalUpvoteConversationRequest(bytes)
//    bytes, err = upvoteConversationRequest.Marshal()
//
//    upvoteConversationResponse, err := UnmarshalUpvoteConversationResponse(bytes)
//    bytes, err = upvoteConversationResponse.Marshal()
//
//    fetchConversationUsageParams, err := UnmarshalFetchConversationUsageParams(bytes)
//    bytes, err = fetchConversationUsageParams.Marshal()
//
//    fetchConversationUsageResponse, err := UnmarshalFetchConversationUsageResponse(bytes)
//    bytes, err = fetchConversationUsageResponse.Marshal()
//
//    completeConversationRequest, err := UnmarshalCompleteConversationRequest(bytes)
//    bytes, err = completeConversationRequest.Marshal()
//
//    completeConversationResponse, err := UnmarshalCompleteConversationResponse(bytes)
//    bytes, err = completeConversationResponse.Marshal()
//
//    createConversationRequest, err := UnmarshalCreateConversationRequest(bytes)
//    bytes, err = createConversationRequest.Marshal()
//
//    createConversationResponse, err := UnmarshalCreateConversationResponse(bytes)
//    bytes, err = createConversationResponse.Marshal()
//
//    dispatchConversationRequest, err := UnmarshalDispatchConversationRequest(bytes)
//    bytes, err = dispatchConversationRequest.Marshal()
//
//    dispatchConversationResponse, err := UnmarshalDispatchConversationResponse(bytes)
//    bytes, err = dispatchConversationResponse.Marshal()
//
//    exportConversationsParams, err := UnmarshalExportConversationsParams(bytes)
//    bytes, err = exportConversationsParams.Marshal()
//
//    exportConversationsResponse, err := UnmarshalExportConversationsResponse(bytes)
//    bytes, err = exportConversationsResponse.Marshal()
//
//    listConversationsParams, err := UnmarshalListConversationsParams(bytes)
//    bytes, err = listConversationsParams.Marshal()
//
//    listConversationsResponse, err := UnmarshalListConversationsResponse(bytes)
//    bytes, err = listConversationsResponse.Marshal()
//
//    deleteDatasetParams, err := UnmarshalDeleteDatasetParams(bytes)
//    bytes, err = deleteDatasetParams.Marshal()
//
//    deleteDatasetRequest, err := UnmarshalDeleteDatasetRequest(bytes)
//    bytes, err = deleteDatasetRequest.Marshal()
//
//    deleteDatasetResponse, err := UnmarshalDeleteDatasetResponse(bytes)
//    bytes, err = deleteDatasetResponse.Marshal()
//
//    fetchDatasetParams, err := UnmarshalFetchDatasetParams(bytes)
//    bytes, err = fetchDatasetParams.Marshal()
//
//    fetchDatasetResponse, err := UnmarshalFetchDatasetResponse(bytes)
//    bytes, err = fetchDatasetResponse.Marshal()
//
//    attachDatasetFileParams, err := UnmarshalAttachDatasetFileParams(bytes)
//    bytes, err = attachDatasetFileParams.Marshal()
//
//    attachDatasetFileRequest, err := UnmarshalAttachDatasetFileRequest(bytes)
//    bytes, err = attachDatasetFileRequest.Marshal()
//
//    attachDatasetFileResponse, err := UnmarshalAttachDatasetFileResponse(bytes)
//    bytes, err = attachDatasetFileResponse.Marshal()
//
//    detachDatasetFileParams, err := UnmarshalDetachDatasetFileParams(bytes)
//    bytes, err = detachDatasetFileParams.Marshal()
//
//    detachDatasetFileRequest, err := UnmarshalDetachDatasetFileRequest(bytes)
//    bytes, err = detachDatasetFileRequest.Marshal()
//
//    detachDatasetFileResponse, err := UnmarshalDetachDatasetFileResponse(bytes)
//    bytes, err = detachDatasetFileResponse.Marshal()
//
//    syncDatasetFileParams, err := UnmarshalSyncDatasetFileParams(bytes)
//    bytes, err = syncDatasetFileParams.Marshal()
//
//    syncDatasetFileRequest, err := UnmarshalSyncDatasetFileRequest(bytes)
//    bytes, err = syncDatasetFileRequest.Marshal()
//
//    syncDatasetFileResponse, err := UnmarshalSyncDatasetFileResponse(bytes)
//    bytes, err = syncDatasetFileResponse.Marshal()
//
//    listDatasetFilesParams, err := UnmarshalListDatasetFilesParams(bytes)
//    bytes, err = listDatasetFilesParams.Marshal()
//
//    listDatasetFilesResponse, err := UnmarshalListDatasetFilesResponse(bytes)
//    bytes, err = listDatasetFilesResponse.Marshal()
//
//    deleteDatasetRecordParams, err := UnmarshalDeleteDatasetRecordParams(bytes)
//    bytes, err = deleteDatasetRecordParams.Marshal()
//
//    deleteDatasetRecordRequest, err := UnmarshalDeleteDatasetRecordRequest(bytes)
//    bytes, err = deleteDatasetRecordRequest.Marshal()
//
//    deleteDatasetRecordResponse, err := UnmarshalDeleteDatasetRecordResponse(bytes)
//    bytes, err = deleteDatasetRecordResponse.Marshal()
//
//    fetchDatasetRecordParams, err := UnmarshalFetchDatasetRecordParams(bytes)
//    bytes, err = fetchDatasetRecordParams.Marshal()
//
//    fetchDatasetRecordResponse, err := UnmarshalFetchDatasetRecordResponse(bytes)
//    bytes, err = fetchDatasetRecordResponse.Marshal()
//
//    updateDatasetRecordParams, err := UnmarshalUpdateDatasetRecordParams(bytes)
//    bytes, err = updateDatasetRecordParams.Marshal()
//
//    updateDatasetRecordRequest, err := UnmarshalUpdateDatasetRecordRequest(bytes)
//    bytes, err = updateDatasetRecordRequest.Marshal()
//
//    updateDatasetRecordResponse, err := UnmarshalUpdateDatasetRecordResponse(bytes)
//    bytes, err = updateDatasetRecordResponse.Marshal()
//
//    createDatasetRecordParams, err := UnmarshalCreateDatasetRecordParams(bytes)
//    bytes, err = createDatasetRecordParams.Marshal()
//
//    createDatasetRecordRequest, err := UnmarshalCreateDatasetRecordRequest(bytes)
//    bytes, err = createDatasetRecordRequest.Marshal()
//
//    createDatasetRecordResponse, err := UnmarshalCreateDatasetRecordResponse(bytes)
//    bytes, err = createDatasetRecordResponse.Marshal()
//
//    exportDatasetRecordsParams, err := UnmarshalExportDatasetRecordsParams(bytes)
//    bytes, err = exportDatasetRecordsParams.Marshal()
//
//    exportDatasetRecordsResponse, err := UnmarshalExportDatasetRecordsResponse(bytes)
//    bytes, err = exportDatasetRecordsResponse.Marshal()
//
//    listDatasetRecordsParams, err := UnmarshalListDatasetRecordsParams(bytes)
//    bytes, err = listDatasetRecordsParams.Marshal()
//
//    listDatasetRecordsResponse, err := UnmarshalListDatasetRecordsResponse(bytes)
//    bytes, err = listDatasetRecordsResponse.Marshal()
//
//    searchDatasetParams, err := UnmarshalSearchDatasetParams(bytes)
//    bytes, err = searchDatasetParams.Marshal()
//
//    searchDatasetRequest, err := UnmarshalSearchDatasetRequest(bytes)
//    bytes, err = searchDatasetRequest.Marshal()
//
//    searchDatasetResponse, err := UnmarshalSearchDatasetResponse(bytes)
//    bytes, err = searchDatasetResponse.Marshal()
//
//    updateDatasetParams, err := UnmarshalUpdateDatasetParams(bytes)
//    bytes, err = updateDatasetParams.Marshal()
//
//    updateDatasetRequest, err := UnmarshalUpdateDatasetRequest(bytes)
//    bytes, err = updateDatasetRequest.Marshal()
//
//    updateDatasetResponse, err := UnmarshalUpdateDatasetResponse(bytes)
//    bytes, err = updateDatasetResponse.Marshal()
//
//    createDatasetRequest, err := UnmarshalCreateDatasetRequest(bytes)
//    bytes, err = createDatasetRequest.Marshal()
//
//    createDatasetResponse, err := UnmarshalCreateDatasetResponse(bytes)
//    bytes, err = createDatasetResponse.Marshal()
//
//    listDatasetsParams, err := UnmarshalListDatasetsParams(bytes)
//    bytes, err = listDatasetsParams.Marshal()
//
//    listDatasetsResponse, err := UnmarshalListDatasetsResponse(bytes)
//    bytes, err = listDatasetsResponse.Marshal()
//
//    exportEventLogsParams, err := UnmarshalExportEventLogsParams(bytes)
//    bytes, err = exportEventLogsParams.Marshal()
//
//    exportEventLogsResponse, err := UnmarshalExportEventLogsResponse(bytes)
//    bytes, err = exportEventLogsResponse.Marshal()
//
//    listEventLogsParams, err := UnmarshalListEventLogsParams(bytes)
//    bytes, err = listEventLogsParams.Marshal()
//
//    listEventLogsResponse, err := UnmarshalListEventLogsResponse(bytes)
//    bytes, err = listEventLogsResponse.Marshal()
//
//    subscribeEventLogsRequest, err := UnmarshalSubscribeEventLogsRequest(bytes)
//    bytes, err = subscribeEventLogsRequest.Marshal()
//
//    deleteFileParams, err := UnmarshalDeleteFileParams(bytes)
//    bytes, err = deleteFileParams.Marshal()
//
//    deleteFileRequest, err := UnmarshalDeleteFileRequest(bytes)
//    bytes, err = deleteFileRequest.Marshal()
//
//    deleteFileResponse, err := UnmarshalDeleteFileResponse(bytes)
//    bytes, err = deleteFileResponse.Marshal()
//
//    downloadFileParams, err := UnmarshalDownloadFileParams(bytes)
//    bytes, err = downloadFileParams.Marshal()
//
//    downloadFileResponse, err := UnmarshalDownloadFileResponse(bytes)
//    bytes, err = downloadFileResponse.Marshal()
//
//    fetchFileParams, err := UnmarshalFetchFileParams(bytes)
//    bytes, err = fetchFileParams.Marshal()
//
//    fetchFileResponse, err := UnmarshalFetchFileResponse(bytes)
//    bytes, err = fetchFileResponse.Marshal()
//
//    syncFileParams, err := UnmarshalSyncFileParams(bytes)
//    bytes, err = syncFileParams.Marshal()
//
//    syncFileRequest, err := UnmarshalSyncFileRequest(bytes)
//    bytes, err = syncFileRequest.Marshal()
//
//    syncFileResponse, err := UnmarshalSyncFileResponse(bytes)
//    bytes, err = syncFileResponse.Marshal()
//
//    updateFileParams, err := UnmarshalUpdateFileParams(bytes)
//    bytes, err = updateFileParams.Marshal()
//
//    updateFileRequest, err := UnmarshalUpdateFileRequest(bytes)
//    bytes, err = updateFileRequest.Marshal()
//
//    updateFileResponse, err := UnmarshalUpdateFileResponse(bytes)
//    bytes, err = updateFileResponse.Marshal()
//
//    uploadFileParams, err := UnmarshalUploadFileParams(bytes)
//    bytes, err = uploadFileParams.Marshal()
//
//    uploadFileRequest, err := UnmarshalUploadFileRequest(bytes)
//    bytes, err = uploadFileRequest.Marshal()
//
//    uploadFileResponse, err := UnmarshalUploadFileResponse(bytes)
//    bytes, err = uploadFileResponse.Marshal()
//
//    createFileRequest, err := UnmarshalCreateFileRequest(bytes)
//    bytes, err = createFileRequest.Marshal()
//
//    createFileResponse, err := UnmarshalCreateFileResponse(bytes)
//    bytes, err = createFileResponse.Marshal()
//
//    listFilesParams, err := UnmarshalListFilesParams(bytes)
//    bytes, err = listFilesParams.Marshal()
//
//    listFilesResponse, err := UnmarshalListFilesResponse(bytes)
//    bytes, err = listFilesResponse.Marshal()
//
//    deleteDiscordIntegrationParams, err := UnmarshalDeleteDiscordIntegrationParams(bytes)
//    bytes, err = deleteDiscordIntegrationParams.Marshal()
//
//    deleteDiscordIntegrationRequest, err := UnmarshalDeleteDiscordIntegrationRequest(bytes)
//    bytes, err = deleteDiscordIntegrationRequest.Marshal()
//
//    deleteDiscordIntegrationResponse, err := UnmarshalDeleteDiscordIntegrationResponse(bytes)
//    bytes, err = deleteDiscordIntegrationResponse.Marshal()
//
//    fetchDiscordIntegrationParams, err := UnmarshalFetchDiscordIntegrationParams(bytes)
//    bytes, err = fetchDiscordIntegrationParams.Marshal()
//
//    fetchDiscordIntegrationResponse, err := UnmarshalFetchDiscordIntegrationResponse(bytes)
//    bytes, err = fetchDiscordIntegrationResponse.Marshal()
//
//    setupDiscordIntegrationParams, err := UnmarshalSetupDiscordIntegrationParams(bytes)
//    bytes, err = setupDiscordIntegrationParams.Marshal()
//
//    setupDiscordIntegrationRequest, err := UnmarshalSetupDiscordIntegrationRequest(bytes)
//    bytes, err = setupDiscordIntegrationRequest.Marshal()
//
//    setupDiscordIntegrationResponse, err := UnmarshalSetupDiscordIntegrationResponse(bytes)
//    bytes, err = setupDiscordIntegrationResponse.Marshal()
//
//    updateDiscordIntegrationParams, err := UnmarshalUpdateDiscordIntegrationParams(bytes)
//    bytes, err = updateDiscordIntegrationParams.Marshal()
//
//    updateDiscordIntegrationRequest, err := UnmarshalUpdateDiscordIntegrationRequest(bytes)
//    bytes, err = updateDiscordIntegrationRequest.Marshal()
//
//    updateDiscordIntegrationResponse, err := UnmarshalUpdateDiscordIntegrationResponse(bytes)
//    bytes, err = updateDiscordIntegrationResponse.Marshal()
//
//    createDiscordIntegrationRequest, err := UnmarshalCreateDiscordIntegrationRequest(bytes)
//    bytes, err = createDiscordIntegrationRequest.Marshal()
//
//    createDiscordIntegrationResponse, err := UnmarshalCreateDiscordIntegrationResponse(bytes)
//    bytes, err = createDiscordIntegrationResponse.Marshal()
//
//    listDiscordIntegrationsParams, err := UnmarshalListDiscordIntegrationsParams(bytes)
//    bytes, err = listDiscordIntegrationsParams.Marshal()
//
//    listDiscordIntegrationsResponse, err := UnmarshalListDiscordIntegrationsResponse(bytes)
//    bytes, err = listDiscordIntegrationsResponse.Marshal()
//
//    deleteEmailIntegrationParams, err := UnmarshalDeleteEmailIntegrationParams(bytes)
//    bytes, err = deleteEmailIntegrationParams.Marshal()
//
//    deleteEmailIntegrationRequest, err := UnmarshalDeleteEmailIntegrationRequest(bytes)
//    bytes, err = deleteEmailIntegrationRequest.Marshal()
//
//    deleteEmailIntegrationResponse, err := UnmarshalDeleteEmailIntegrationResponse(bytes)
//    bytes, err = deleteEmailIntegrationResponse.Marshal()
//
//    fetchEmailIntegrationParams, err := UnmarshalFetchEmailIntegrationParams(bytes)
//    bytes, err = fetchEmailIntegrationParams.Marshal()
//
//    fetchEmailIntegrationResponse, err := UnmarshalFetchEmailIntegrationResponse(bytes)
//    bytes, err = fetchEmailIntegrationResponse.Marshal()
//
//    setupEmailIntegrationParams, err := UnmarshalSetupEmailIntegrationParams(bytes)
//    bytes, err = setupEmailIntegrationParams.Marshal()
//
//    setupEmailIntegrationRequest, err := UnmarshalSetupEmailIntegrationRequest(bytes)
//    bytes, err = setupEmailIntegrationRequest.Marshal()
//
//    setupEmailIntegrationResponse, err := UnmarshalSetupEmailIntegrationResponse(bytes)
//    bytes, err = setupEmailIntegrationResponse.Marshal()
//
//    updateEmailIntegrationParams, err := UnmarshalUpdateEmailIntegrationParams(bytes)
//    bytes, err = updateEmailIntegrationParams.Marshal()
//
//    updateEmailIntegrationRequest, err := UnmarshalUpdateEmailIntegrationRequest(bytes)
//    bytes, err = updateEmailIntegrationRequest.Marshal()
//
//    updateEmailIntegrationResponse, err := UnmarshalUpdateEmailIntegrationResponse(bytes)
//    bytes, err = updateEmailIntegrationResponse.Marshal()
//
//    createEmailIntegrationRequest, err := UnmarshalCreateEmailIntegrationRequest(bytes)
//    bytes, err = createEmailIntegrationRequest.Marshal()
//
//    createEmailIntegrationResponse, err := UnmarshalCreateEmailIntegrationResponse(bytes)
//    bytes, err = createEmailIntegrationResponse.Marshal()
//
//    listEmailIntegrationsParams, err := UnmarshalListEmailIntegrationsParams(bytes)
//    bytes, err = listEmailIntegrationsParams.Marshal()
//
//    listEmailIntegrationsResponse, err := UnmarshalListEmailIntegrationsResponse(bytes)
//    bytes, err = listEmailIntegrationsResponse.Marshal()
//
//    deleteExtractIntegrationParams, err := UnmarshalDeleteExtractIntegrationParams(bytes)
//    bytes, err = deleteExtractIntegrationParams.Marshal()
//
//    deleteExtractIntegrationRequest, err := UnmarshalDeleteExtractIntegrationRequest(bytes)
//    bytes, err = deleteExtractIntegrationRequest.Marshal()
//
//    deleteExtractIntegrationResponse, err := UnmarshalDeleteExtractIntegrationResponse(bytes)
//    bytes, err = deleteExtractIntegrationResponse.Marshal()
//
//    fetchExtractIntegrationParams, err := UnmarshalFetchExtractIntegrationParams(bytes)
//    bytes, err = fetchExtractIntegrationParams.Marshal()
//
//    fetchExtractIntegrationResponse, err := UnmarshalFetchExtractIntegrationResponse(bytes)
//    bytes, err = fetchExtractIntegrationResponse.Marshal()
//
//    updateExtractIntegrationParams, err := UnmarshalUpdateExtractIntegrationParams(bytes)
//    bytes, err = updateExtractIntegrationParams.Marshal()
//
//    updateExtractIntegrationRequest, err := UnmarshalUpdateExtractIntegrationRequest(bytes)
//    bytes, err = updateExtractIntegrationRequest.Marshal()
//
//    updateExtractIntegrationResponse, err := UnmarshalUpdateExtractIntegrationResponse(bytes)
//    bytes, err = updateExtractIntegrationResponse.Marshal()
//
//    createExtractIntegrationRequest, err := UnmarshalCreateExtractIntegrationRequest(bytes)
//    bytes, err = createExtractIntegrationRequest.Marshal()
//
//    createExtractIntegrationResponse, err := UnmarshalCreateExtractIntegrationResponse(bytes)
//    bytes, err = createExtractIntegrationResponse.Marshal()
//
//    listExtractIntegrationsParams, err := UnmarshalListExtractIntegrationsParams(bytes)
//    bytes, err = listExtractIntegrationsParams.Marshal()
//
//    listExtractIntegrationsResponse, err := UnmarshalListExtractIntegrationsResponse(bytes)
//    bytes, err = listExtractIntegrationsResponse.Marshal()
//
//    deleteInstagramIntegrationParams, err := UnmarshalDeleteInstagramIntegrationParams(bytes)
//    bytes, err = deleteInstagramIntegrationParams.Marshal()
//
//    deleteInstagramIntegrationRequest, err := UnmarshalDeleteInstagramIntegrationRequest(bytes)
//    bytes, err = deleteInstagramIntegrationRequest.Marshal()
//
//    deleteInstagramIntegrationResponse, err := UnmarshalDeleteInstagramIntegrationResponse(bytes)
//    bytes, err = deleteInstagramIntegrationResponse.Marshal()
//
//    fetchInstagramIntegrationParams, err := UnmarshalFetchInstagramIntegrationParams(bytes)
//    bytes, err = fetchInstagramIntegrationParams.Marshal()
//
//    fetchInstagramIntegrationResponse, err := UnmarshalFetchInstagramIntegrationResponse(bytes)
//    bytes, err = fetchInstagramIntegrationResponse.Marshal()
//
//    setupInstagramIntegrationParams, err := UnmarshalSetupInstagramIntegrationParams(bytes)
//    bytes, err = setupInstagramIntegrationParams.Marshal()
//
//    setupInstagramIntegrationRequest, err := UnmarshalSetupInstagramIntegrationRequest(bytes)
//    bytes, err = setupInstagramIntegrationRequest.Marshal()
//
//    setupInstagramIntegrationResponse, err := UnmarshalSetupInstagramIntegrationResponse(bytes)
//    bytes, err = setupInstagramIntegrationResponse.Marshal()
//
//    updateInstagramIntegrationParams, err := UnmarshalUpdateInstagramIntegrationParams(bytes)
//    bytes, err = updateInstagramIntegrationParams.Marshal()
//
//    updateInstagramIntegrationRequest, err := UnmarshalUpdateInstagramIntegrationRequest(bytes)
//    bytes, err = updateInstagramIntegrationRequest.Marshal()
//
//    updateInstagramIntegrationResponse, err := UnmarshalUpdateInstagramIntegrationResponse(bytes)
//    bytes, err = updateInstagramIntegrationResponse.Marshal()
//
//    createInstagramIntegrationRequest, err := UnmarshalCreateInstagramIntegrationRequest(bytes)
//    bytes, err = createInstagramIntegrationRequest.Marshal()
//
//    createInstagramIntegrationResponse, err := UnmarshalCreateInstagramIntegrationResponse(bytes)
//    bytes, err = createInstagramIntegrationResponse.Marshal()
//
//    listInstagramIntegrationsParams, err := UnmarshalListInstagramIntegrationsParams(bytes)
//    bytes, err = listInstagramIntegrationsParams.Marshal()
//
//    listInstagramIntegrationsResponse, err := UnmarshalListInstagramIntegrationsResponse(bytes)
//    bytes, err = listInstagramIntegrationsResponse.Marshal()
//
//    deleteMCPServerIntegrationParams, err := UnmarshalDeleteMCPServerIntegrationParams(bytes)
//    bytes, err = deleteMCPServerIntegrationParams.Marshal()
//
//    deleteMCPServerIntegrationRequest, err := UnmarshalDeleteMCPServerIntegrationRequest(bytes)
//    bytes, err = deleteMCPServerIntegrationRequest.Marshal()
//
//    deleteMCPServerIntegrationResponse, err := UnmarshalDeleteMCPServerIntegrationResponse(bytes)
//    bytes, err = deleteMCPServerIntegrationResponse.Marshal()
//
//    fetchMCPServerIntegrationParams, err := UnmarshalFetchMCPServerIntegrationParams(bytes)
//    bytes, err = fetchMCPServerIntegrationParams.Marshal()
//
//    fetchMCPServerIntegrationResponse, err := UnmarshalFetchMCPServerIntegrationResponse(bytes)
//    bytes, err = fetchMCPServerIntegrationResponse.Marshal()
//
//    updateMCPServerIntegrationParams, err := UnmarshalUpdateMCPServerIntegrationParams(bytes)
//    bytes, err = updateMCPServerIntegrationParams.Marshal()
//
//    updateMCPServerIntegrationRequest, err := UnmarshalUpdateMCPServerIntegrationRequest(bytes)
//    bytes, err = updateMCPServerIntegrationRequest.Marshal()
//
//    updateMCPServerIntegrationResponse, err := UnmarshalUpdateMCPServerIntegrationResponse(bytes)
//    bytes, err = updateMCPServerIntegrationResponse.Marshal()
//
//    createMCPServerIntegrationRequest, err := UnmarshalCreateMCPServerIntegrationRequest(bytes)
//    bytes, err = createMCPServerIntegrationRequest.Marshal()
//
//    createMCPServerIntegrationResponse, err := UnmarshalCreateMCPServerIntegrationResponse(bytes)
//    bytes, err = createMCPServerIntegrationResponse.Marshal()
//
//    listMCPServerIntegrationsParams, err := UnmarshalListMCPServerIntegrationsParams(bytes)
//    bytes, err = listMCPServerIntegrationsParams.Marshal()
//
//    listMCPServerIntegrationsResponse, err := UnmarshalListMCPServerIntegrationsResponse(bytes)
//    bytes, err = listMCPServerIntegrationsResponse.Marshal()
//
//    deleteMessengerIntegrationParams, err := UnmarshalDeleteMessengerIntegrationParams(bytes)
//    bytes, err = deleteMessengerIntegrationParams.Marshal()
//
//    deleteMessengerIntegrationRequest, err := UnmarshalDeleteMessengerIntegrationRequest(bytes)
//    bytes, err = deleteMessengerIntegrationRequest.Marshal()
//
//    deleteMessengerIntegrationResponse, err := UnmarshalDeleteMessengerIntegrationResponse(bytes)
//    bytes, err = deleteMessengerIntegrationResponse.Marshal()
//
//    fetchMessengerIntegrationParams, err := UnmarshalFetchMessengerIntegrationParams(bytes)
//    bytes, err = fetchMessengerIntegrationParams.Marshal()
//
//    fetchMessengerIntegrationResponse, err := UnmarshalFetchMessengerIntegrationResponse(bytes)
//    bytes, err = fetchMessengerIntegrationResponse.Marshal()
//
//    setupMessengerIntegrationParams, err := UnmarshalSetupMessengerIntegrationParams(bytes)
//    bytes, err = setupMessengerIntegrationParams.Marshal()
//
//    setupMessengerIntegrationRequest, err := UnmarshalSetupMessengerIntegrationRequest(bytes)
//    bytes, err = setupMessengerIntegrationRequest.Marshal()
//
//    setupMessengerIntegrationResponse, err := UnmarshalSetupMessengerIntegrationResponse(bytes)
//    bytes, err = setupMessengerIntegrationResponse.Marshal()
//
//    updateMessengerIntegrationParams, err := UnmarshalUpdateMessengerIntegrationParams(bytes)
//    bytes, err = updateMessengerIntegrationParams.Marshal()
//
//    updateMessengerIntegrationRequest, err := UnmarshalUpdateMessengerIntegrationRequest(bytes)
//    bytes, err = updateMessengerIntegrationRequest.Marshal()
//
//    updateMessengerIntegrationResponse, err := UnmarshalUpdateMessengerIntegrationResponse(bytes)
//    bytes, err = updateMessengerIntegrationResponse.Marshal()
//
//    createMessengerIntegrationRequest, err := UnmarshalCreateMessengerIntegrationRequest(bytes)
//    bytes, err = createMessengerIntegrationRequest.Marshal()
//
//    createMessengerIntegrationResponse, err := UnmarshalCreateMessengerIntegrationResponse(bytes)
//    bytes, err = createMessengerIntegrationResponse.Marshal()
//
//    listMessengerIntegrationsParams, err := UnmarshalListMessengerIntegrationsParams(bytes)
//    bytes, err = listMessengerIntegrationsParams.Marshal()
//
//    listMessengerIntegrationsResponse, err := UnmarshalListMessengerIntegrationsResponse(bytes)
//    bytes, err = listMessengerIntegrationsResponse.Marshal()
//
//    deleteNotionIntegrationParams, err := UnmarshalDeleteNotionIntegrationParams(bytes)
//    bytes, err = deleteNotionIntegrationParams.Marshal()
//
//    deleteNotionIntegrationRequest, err := UnmarshalDeleteNotionIntegrationRequest(bytes)
//    bytes, err = deleteNotionIntegrationRequest.Marshal()
//
//    deleteNotionIntegrationResponse, err := UnmarshalDeleteNotionIntegrationResponse(bytes)
//    bytes, err = deleteNotionIntegrationResponse.Marshal()
//
//    fetchNotionIntegrationParams, err := UnmarshalFetchNotionIntegrationParams(bytes)
//    bytes, err = fetchNotionIntegrationParams.Marshal()
//
//    fetchNotionIntegrationResponse, err := UnmarshalFetchNotionIntegrationResponse(bytes)
//    bytes, err = fetchNotionIntegrationResponse.Marshal()
//
//    syncNotionIntegrationParams, err := UnmarshalSyncNotionIntegrationParams(bytes)
//    bytes, err = syncNotionIntegrationParams.Marshal()
//
//    syncNotionIntegrationRequest, err := UnmarshalSyncNotionIntegrationRequest(bytes)
//    bytes, err = syncNotionIntegrationRequest.Marshal()
//
//    syncNotionIntegrationResponse, err := UnmarshalSyncNotionIntegrationResponse(bytes)
//    bytes, err = syncNotionIntegrationResponse.Marshal()
//
//    updateNotionIntegrationParams, err := UnmarshalUpdateNotionIntegrationParams(bytes)
//    bytes, err = updateNotionIntegrationParams.Marshal()
//
//    updateNotionIntegrationRequest, err := UnmarshalUpdateNotionIntegrationRequest(bytes)
//    bytes, err = updateNotionIntegrationRequest.Marshal()
//
//    updateNotionIntegrationResponse, err := UnmarshalUpdateNotionIntegrationResponse(bytes)
//    bytes, err = updateNotionIntegrationResponse.Marshal()
//
//    createNotionIntegrationRequest, err := UnmarshalCreateNotionIntegrationRequest(bytes)
//    bytes, err = createNotionIntegrationRequest.Marshal()
//
//    createNotionIntegrationResponse, err := UnmarshalCreateNotionIntegrationResponse(bytes)
//    bytes, err = createNotionIntegrationResponse.Marshal()
//
//    listNotionIntegrationsParams, err := UnmarshalListNotionIntegrationsParams(bytes)
//    bytes, err = listNotionIntegrationsParams.Marshal()
//
//    listNotionIntegrationsResponse, err := UnmarshalListNotionIntegrationsResponse(bytes)
//    bytes, err = listNotionIntegrationsResponse.Marshal()
//
//    deleteSitemapIntegrationParams, err := UnmarshalDeleteSitemapIntegrationParams(bytes)
//    bytes, err = deleteSitemapIntegrationParams.Marshal()
//
//    deleteSitemapIntegrationRequest, err := UnmarshalDeleteSitemapIntegrationRequest(bytes)
//    bytes, err = deleteSitemapIntegrationRequest.Marshal()
//
//    deleteSitemapIntegrationResponse, err := UnmarshalDeleteSitemapIntegrationResponse(bytes)
//    bytes, err = deleteSitemapIntegrationResponse.Marshal()
//
//    fetchSitemapIntegrationParams, err := UnmarshalFetchSitemapIntegrationParams(bytes)
//    bytes, err = fetchSitemapIntegrationParams.Marshal()
//
//    fetchSitemapIntegrationResponse, err := UnmarshalFetchSitemapIntegrationResponse(bytes)
//    bytes, err = fetchSitemapIntegrationResponse.Marshal()
//
//    syncSitemapIntegrationParams, err := UnmarshalSyncSitemapIntegrationParams(bytes)
//    bytes, err = syncSitemapIntegrationParams.Marshal()
//
//    syncSitemapIntegrationRequest, err := UnmarshalSyncSitemapIntegrationRequest(bytes)
//    bytes, err = syncSitemapIntegrationRequest.Marshal()
//
//    syncSitemapIntegrationResponse, err := UnmarshalSyncSitemapIntegrationResponse(bytes)
//    bytes, err = syncSitemapIntegrationResponse.Marshal()
//
//    updateSitemapIntegrationParams, err := UnmarshalUpdateSitemapIntegrationParams(bytes)
//    bytes, err = updateSitemapIntegrationParams.Marshal()
//
//    updateSitemapIntegrationRequest, err := UnmarshalUpdateSitemapIntegrationRequest(bytes)
//    bytes, err = updateSitemapIntegrationRequest.Marshal()
//
//    updateSitemapIntegrationResponse, err := UnmarshalUpdateSitemapIntegrationResponse(bytes)
//    bytes, err = updateSitemapIntegrationResponse.Marshal()
//
//    createSitemapIntegrationRequest, err := UnmarshalCreateSitemapIntegrationRequest(bytes)
//    bytes, err = createSitemapIntegrationRequest.Marshal()
//
//    createSitemapIntegrationResponse, err := UnmarshalCreateSitemapIntegrationResponse(bytes)
//    bytes, err = createSitemapIntegrationResponse.Marshal()
//
//    listSitemapIntegrationsParams, err := UnmarshalListSitemapIntegrationsParams(bytes)
//    bytes, err = listSitemapIntegrationsParams.Marshal()
//
//    listSitemapIntegrationsResponse, err := UnmarshalListSitemapIntegrationsResponse(bytes)
//    bytes, err = listSitemapIntegrationsResponse.Marshal()
//
//    deleteSlackIntegrationParams, err := UnmarshalDeleteSlackIntegrationParams(bytes)
//    bytes, err = deleteSlackIntegrationParams.Marshal()
//
//    deleteSlackIntegrationRequest, err := UnmarshalDeleteSlackIntegrationRequest(bytes)
//    bytes, err = deleteSlackIntegrationRequest.Marshal()
//
//    deleteSlackIntegrationResponse, err := UnmarshalDeleteSlackIntegrationResponse(bytes)
//    bytes, err = deleteSlackIntegrationResponse.Marshal()
//
//    fetchSlackIntegrationParams, err := UnmarshalFetchSlackIntegrationParams(bytes)
//    bytes, err = fetchSlackIntegrationParams.Marshal()
//
//    fetchSlackIntegrationResponse, err := UnmarshalFetchSlackIntegrationResponse(bytes)
//    bytes, err = fetchSlackIntegrationResponse.Marshal()
//
//    setupSlackIntegrationParams, err := UnmarshalSetupSlackIntegrationParams(bytes)
//    bytes, err = setupSlackIntegrationParams.Marshal()
//
//    setupSlackIntegrationRequest, err := UnmarshalSetupSlackIntegrationRequest(bytes)
//    bytes, err = setupSlackIntegrationRequest.Marshal()
//
//    setupSlackIntegrationResponse, err := UnmarshalSetupSlackIntegrationResponse(bytes)
//    bytes, err = setupSlackIntegrationResponse.Marshal()
//
//    updateSlackIntegrationParams, err := UnmarshalUpdateSlackIntegrationParams(bytes)
//    bytes, err = updateSlackIntegrationParams.Marshal()
//
//    updateSlackIntegrationRequest, err := UnmarshalUpdateSlackIntegrationRequest(bytes)
//    bytes, err = updateSlackIntegrationRequest.Marshal()
//
//    updateSlackIntegrationResponse, err := UnmarshalUpdateSlackIntegrationResponse(bytes)
//    bytes, err = updateSlackIntegrationResponse.Marshal()
//
//    createSlackIntegrationRequest, err := UnmarshalCreateSlackIntegrationRequest(bytes)
//    bytes, err = createSlackIntegrationRequest.Marshal()
//
//    createSlackIntegrationResponse, err := UnmarshalCreateSlackIntegrationResponse(bytes)
//    bytes, err = createSlackIntegrationResponse.Marshal()
//
//    listSlackIntegrationsParams, err := UnmarshalListSlackIntegrationsParams(bytes)
//    bytes, err = listSlackIntegrationsParams.Marshal()
//
//    listSlackIntegrationsResponse, err := UnmarshalListSlackIntegrationsResponse(bytes)
//    bytes, err = listSlackIntegrationsResponse.Marshal()
//
//    deleteSupportIntegrationParams, err := UnmarshalDeleteSupportIntegrationParams(bytes)
//    bytes, err = deleteSupportIntegrationParams.Marshal()
//
//    deleteSupportIntegrationRequest, err := UnmarshalDeleteSupportIntegrationRequest(bytes)
//    bytes, err = deleteSupportIntegrationRequest.Marshal()
//
//    deleteSupportIntegrationResponse, err := UnmarshalDeleteSupportIntegrationResponse(bytes)
//    bytes, err = deleteSupportIntegrationResponse.Marshal()
//
//    fetchSupportIntegrationParams, err := UnmarshalFetchSupportIntegrationParams(bytes)
//    bytes, err = fetchSupportIntegrationParams.Marshal()
//
//    fetchSupportIntegrationResponse, err := UnmarshalFetchSupportIntegrationResponse(bytes)
//    bytes, err = fetchSupportIntegrationResponse.Marshal()
//
//    updateSupportIntegrationParams, err := UnmarshalUpdateSupportIntegrationParams(bytes)
//    bytes, err = updateSupportIntegrationParams.Marshal()
//
//    updateSupportIntegrationRequest, err := UnmarshalUpdateSupportIntegrationRequest(bytes)
//    bytes, err = updateSupportIntegrationRequest.Marshal()
//
//    updateSupportIntegrationResponse, err := UnmarshalUpdateSupportIntegrationResponse(bytes)
//    bytes, err = updateSupportIntegrationResponse.Marshal()
//
//    createSupportIntegrationRequest, err := UnmarshalCreateSupportIntegrationRequest(bytes)
//    bytes, err = createSupportIntegrationRequest.Marshal()
//
//    createSupportIntegrationResponse, err := UnmarshalCreateSupportIntegrationResponse(bytes)
//    bytes, err = createSupportIntegrationResponse.Marshal()
//
//    listSupportIntegrationsParams, err := UnmarshalListSupportIntegrationsParams(bytes)
//    bytes, err = listSupportIntegrationsParams.Marshal()
//
//    listSupportIntegrationsResponse, err := UnmarshalListSupportIntegrationsResponse(bytes)
//    bytes, err = listSupportIntegrationsResponse.Marshal()
//
//    deleteTelegramIntegrationParams, err := UnmarshalDeleteTelegramIntegrationParams(bytes)
//    bytes, err = deleteTelegramIntegrationParams.Marshal()
//
//    deleteTelegramIntegrationRequest, err := UnmarshalDeleteTelegramIntegrationRequest(bytes)
//    bytes, err = deleteTelegramIntegrationRequest.Marshal()
//
//    deleteTelegramIntegrationResponse, err := UnmarshalDeleteTelegramIntegrationResponse(bytes)
//    bytes, err = deleteTelegramIntegrationResponse.Marshal()
//
//    fetchTelegramIntegrationParams, err := UnmarshalFetchTelegramIntegrationParams(bytes)
//    bytes, err = fetchTelegramIntegrationParams.Marshal()
//
//    fetchTelegramIntegrationResponse, err := UnmarshalFetchTelegramIntegrationResponse(bytes)
//    bytes, err = fetchTelegramIntegrationResponse.Marshal()
//
//    setupTelegramIntegrationParams, err := UnmarshalSetupTelegramIntegrationParams(bytes)
//    bytes, err = setupTelegramIntegrationParams.Marshal()
//
//    setupTelegramIntegrationRequest, err := UnmarshalSetupTelegramIntegrationRequest(bytes)
//    bytes, err = setupTelegramIntegrationRequest.Marshal()
//
//    setupTelegramIntegrationResponse, err := UnmarshalSetupTelegramIntegrationResponse(bytes)
//    bytes, err = setupTelegramIntegrationResponse.Marshal()
//
//    updateTelegramIntegrationParams, err := UnmarshalUpdateTelegramIntegrationParams(bytes)
//    bytes, err = updateTelegramIntegrationParams.Marshal()
//
//    updateTelegramIntegrationRequest, err := UnmarshalUpdateTelegramIntegrationRequest(bytes)
//    bytes, err = updateTelegramIntegrationRequest.Marshal()
//
//    updateTelegramIntegrationResponse, err := UnmarshalUpdateTelegramIntegrationResponse(bytes)
//    bytes, err = updateTelegramIntegrationResponse.Marshal()
//
//    createTelegramIntegrationRequest, err := UnmarshalCreateTelegramIntegrationRequest(bytes)
//    bytes, err = createTelegramIntegrationRequest.Marshal()
//
//    createTelegramIntegrationResponse, err := UnmarshalCreateTelegramIntegrationResponse(bytes)
//    bytes, err = createTelegramIntegrationResponse.Marshal()
//
//    listTelegramIntegrationsParams, err := UnmarshalListTelegramIntegrationsParams(bytes)
//    bytes, err = listTelegramIntegrationsParams.Marshal()
//
//    listTelegramIntegrationsResponse, err := UnmarshalListTelegramIntegrationsResponse(bytes)
//    bytes, err = listTelegramIntegrationsResponse.Marshal()
//
//    deleteTriggerIntegrationParams, err := UnmarshalDeleteTriggerIntegrationParams(bytes)
//    bytes, err = deleteTriggerIntegrationParams.Marshal()
//
//    deleteTriggerIntegrationRequest, err := UnmarshalDeleteTriggerIntegrationRequest(bytes)
//    bytes, err = deleteTriggerIntegrationRequest.Marshal()
//
//    deleteTriggerIntegrationResponse, err := UnmarshalDeleteTriggerIntegrationResponse(bytes)
//    bytes, err = deleteTriggerIntegrationResponse.Marshal()
//
//    fetchTriggerIntegrationParams, err := UnmarshalFetchTriggerIntegrationParams(bytes)
//    bytes, err = fetchTriggerIntegrationParams.Marshal()
//
//    fetchTriggerIntegrationResponse, err := UnmarshalFetchTriggerIntegrationResponse(bytes)
//    bytes, err = fetchTriggerIntegrationResponse.Marshal()
//
//    invokeTriggerIntegrationParams, err := UnmarshalInvokeTriggerIntegrationParams(bytes)
//    bytes, err = invokeTriggerIntegrationParams.Marshal()
//
//    invokeTriggerIntegrationRequest, err := UnmarshalInvokeTriggerIntegrationRequest(bytes)
//    bytes, err = invokeTriggerIntegrationRequest.Marshal()
//
//    invokeTriggerIntegrationResponse, err := UnmarshalInvokeTriggerIntegrationResponse(bytes)
//    bytes, err = invokeTriggerIntegrationResponse.Marshal()
//
//    setupTriggerIntegrationParams, err := UnmarshalSetupTriggerIntegrationParams(bytes)
//    bytes, err = setupTriggerIntegrationParams.Marshal()
//
//    setupTriggerIntegrationRequest, err := UnmarshalSetupTriggerIntegrationRequest(bytes)
//    bytes, err = setupTriggerIntegrationRequest.Marshal()
//
//    setupTriggerIntegrationResponse, err := UnmarshalSetupTriggerIntegrationResponse(bytes)
//    bytes, err = setupTriggerIntegrationResponse.Marshal()
//
//    updateTriggerIntegrationParams, err := UnmarshalUpdateTriggerIntegrationParams(bytes)
//    bytes, err = updateTriggerIntegrationParams.Marshal()
//
//    updateTriggerIntegrationRequest, err := UnmarshalUpdateTriggerIntegrationRequest(bytes)
//    bytes, err = updateTriggerIntegrationRequest.Marshal()
//
//    updateTriggerIntegrationResponse, err := UnmarshalUpdateTriggerIntegrationResponse(bytes)
//    bytes, err = updateTriggerIntegrationResponse.Marshal()
//
//    createTriggerIntegrationRequest, err := UnmarshalCreateTriggerIntegrationRequest(bytes)
//    bytes, err = createTriggerIntegrationRequest.Marshal()
//
//    createTriggerIntegrationResponse, err := UnmarshalCreateTriggerIntegrationResponse(bytes)
//    bytes, err = createTriggerIntegrationResponse.Marshal()
//
//    listTriggerIntegrationsParams, err := UnmarshalListTriggerIntegrationsParams(bytes)
//    bytes, err = listTriggerIntegrationsParams.Marshal()
//
//    listTriggerIntegrationsResponse, err := UnmarshalListTriggerIntegrationsResponse(bytes)
//    bytes, err = listTriggerIntegrationsResponse.Marshal()
//
//    deleteTwilioIntegrationParams, err := UnmarshalDeleteTwilioIntegrationParams(bytes)
//    bytes, err = deleteTwilioIntegrationParams.Marshal()
//
//    deleteTwilioIntegrationRequest, err := UnmarshalDeleteTwilioIntegrationRequest(bytes)
//    bytes, err = deleteTwilioIntegrationRequest.Marshal()
//
//    deleteTwilioIntegrationResponse, err := UnmarshalDeleteTwilioIntegrationResponse(bytes)
//    bytes, err = deleteTwilioIntegrationResponse.Marshal()
//
//    fetchTwilioIntegrationParams, err := UnmarshalFetchTwilioIntegrationParams(bytes)
//    bytes, err = fetchTwilioIntegrationParams.Marshal()
//
//    fetchTwilioIntegrationResponse, err := UnmarshalFetchTwilioIntegrationResponse(bytes)
//    bytes, err = fetchTwilioIntegrationResponse.Marshal()
//
//    setupTwilioIntegrationParams, err := UnmarshalSetupTwilioIntegrationParams(bytes)
//    bytes, err = setupTwilioIntegrationParams.Marshal()
//
//    setupTwilioIntegrationRequest, err := UnmarshalSetupTwilioIntegrationRequest(bytes)
//    bytes, err = setupTwilioIntegrationRequest.Marshal()
//
//    setupTwilioIntegrationResponse, err := UnmarshalSetupTwilioIntegrationResponse(bytes)
//    bytes, err = setupTwilioIntegrationResponse.Marshal()
//
//    updateTwilioIntegrationParams, err := UnmarshalUpdateTwilioIntegrationParams(bytes)
//    bytes, err = updateTwilioIntegrationParams.Marshal()
//
//    updateTwilioIntegrationRequest, err := UnmarshalUpdateTwilioIntegrationRequest(bytes)
//    bytes, err = updateTwilioIntegrationRequest.Marshal()
//
//    updateTwilioIntegrationResponse, err := UnmarshalUpdateTwilioIntegrationResponse(bytes)
//    bytes, err = updateTwilioIntegrationResponse.Marshal()
//
//    createTwilioIntegrationRequest, err := UnmarshalCreateTwilioIntegrationRequest(bytes)
//    bytes, err = createTwilioIntegrationRequest.Marshal()
//
//    createTwilioIntegrationResponse, err := UnmarshalCreateTwilioIntegrationResponse(bytes)
//    bytes, err = createTwilioIntegrationResponse.Marshal()
//
//    listTwilioIntegrationsParams, err := UnmarshalListTwilioIntegrationsParams(bytes)
//    bytes, err = listTwilioIntegrationsParams.Marshal()
//
//    listTwilioIntegrationsResponse, err := UnmarshalListTwilioIntegrationsResponse(bytes)
//    bytes, err = listTwilioIntegrationsResponse.Marshal()
//
//    deleteWhatsAppIntegrationParams, err := UnmarshalDeleteWhatsAppIntegrationParams(bytes)
//    bytes, err = deleteWhatsAppIntegrationParams.Marshal()
//
//    deleteWhatsAppIntegrationRequest, err := UnmarshalDeleteWhatsAppIntegrationRequest(bytes)
//    bytes, err = deleteWhatsAppIntegrationRequest.Marshal()
//
//    deleteWhatsAppIntegrationResponse, err := UnmarshalDeleteWhatsAppIntegrationResponse(bytes)
//    bytes, err = deleteWhatsAppIntegrationResponse.Marshal()
//
//    fetchWhatsAppIntegrationParams, err := UnmarshalFetchWhatsAppIntegrationParams(bytes)
//    bytes, err = fetchWhatsAppIntegrationParams.Marshal()
//
//    fetchWhatsAppIntegrationResponse, err := UnmarshalFetchWhatsAppIntegrationResponse(bytes)
//    bytes, err = fetchWhatsAppIntegrationResponse.Marshal()
//
//    setupWhatsAppIntegrationParams, err := UnmarshalSetupWhatsAppIntegrationParams(bytes)
//    bytes, err = setupWhatsAppIntegrationParams.Marshal()
//
//    setupWhatsAppIntegrationRequest, err := UnmarshalSetupWhatsAppIntegrationRequest(bytes)
//    bytes, err = setupWhatsAppIntegrationRequest.Marshal()
//
//    setupWhatsAppIntegrationResponse, err := UnmarshalSetupWhatsAppIntegrationResponse(bytes)
//    bytes, err = setupWhatsAppIntegrationResponse.Marshal()
//
//    updateWhatsAppIntegrationParams, err := UnmarshalUpdateWhatsAppIntegrationParams(bytes)
//    bytes, err = updateWhatsAppIntegrationParams.Marshal()
//
//    updateWhatsAppIntegrationRequest, err := UnmarshalUpdateWhatsAppIntegrationRequest(bytes)
//    bytes, err = updateWhatsAppIntegrationRequest.Marshal()
//
//    updateWhatsAppIntegrationResponse, err := UnmarshalUpdateWhatsAppIntegrationResponse(bytes)
//    bytes, err = updateWhatsAppIntegrationResponse.Marshal()
//
//    createWhatsAppIntegrationRequest, err := UnmarshalCreateWhatsAppIntegrationRequest(bytes)
//    bytes, err = createWhatsAppIntegrationRequest.Marshal()
//
//    createWhatsAppIntegrationResponse, err := UnmarshalCreateWhatsAppIntegrationResponse(bytes)
//    bytes, err = createWhatsAppIntegrationResponse.Marshal()
//
//    listWhatsAppIntegrationsParams, err := UnmarshalListWhatsAppIntegrationsParams(bytes)
//    bytes, err = listWhatsAppIntegrationsParams.Marshal()
//
//    listWhatsAppIntegrationsResponse, err := UnmarshalListWhatsAppIntegrationsResponse(bytes)
//    bytes, err = listWhatsAppIntegrationsResponse.Marshal()
//
//    deleteWidgetIntegrationParams, err := UnmarshalDeleteWidgetIntegrationParams(bytes)
//    bytes, err = deleteWidgetIntegrationParams.Marshal()
//
//    deleteWidgetIntegrationRequest, err := UnmarshalDeleteWidgetIntegrationRequest(bytes)
//    bytes, err = deleteWidgetIntegrationRequest.Marshal()
//
//    deleteWidgetIntegrationResponse, err := UnmarshalDeleteWidgetIntegrationResponse(bytes)
//    bytes, err = deleteWidgetIntegrationResponse.Marshal()
//
//    fetchWidgetIntegrationParams, err := UnmarshalFetchWidgetIntegrationParams(bytes)
//    bytes, err = fetchWidgetIntegrationParams.Marshal()
//
//    fetchWidgetIntegrationResponse, err := UnmarshalFetchWidgetIntegrationResponse(bytes)
//    bytes, err = fetchWidgetIntegrationResponse.Marshal()
//
//    setupWidgetIntegrationParams, err := UnmarshalSetupWidgetIntegrationParams(bytes)
//    bytes, err = setupWidgetIntegrationParams.Marshal()
//
//    setupWidgetIntegrationRequest, err := UnmarshalSetupWidgetIntegrationRequest(bytes)
//    bytes, err = setupWidgetIntegrationRequest.Marshal()
//
//    setupWidgetIntegrationResponse, err := UnmarshalSetupWidgetIntegrationResponse(bytes)
//    bytes, err = setupWidgetIntegrationResponse.Marshal()
//
//    updateWidgetIntegrationParams, err := UnmarshalUpdateWidgetIntegrationParams(bytes)
//    bytes, err = updateWidgetIntegrationParams.Marshal()
//
//    updateWidgetIntegrationRequest, err := UnmarshalUpdateWidgetIntegrationRequest(bytes)
//    bytes, err = updateWidgetIntegrationRequest.Marshal()
//
//    updateWidgetIntegrationResponse, err := UnmarshalUpdateWidgetIntegrationResponse(bytes)
//    bytes, err = updateWidgetIntegrationResponse.Marshal()
//
//    createWidgetIntegrationRequest, err := UnmarshalCreateWidgetIntegrationRequest(bytes)
//    bytes, err = createWidgetIntegrationRequest.Marshal()
//
//    createWidgetIntegrationResponse, err := UnmarshalCreateWidgetIntegrationResponse(bytes)
//    bytes, err = createWidgetIntegrationResponse.Marshal()
//
//    listWidgetIntegrationsParams, err := UnmarshalListWidgetIntegrationsParams(bytes)
//    bytes, err = listWidgetIntegrationsParams.Marshal()
//
//    listWidgetIntegrationsResponse, err := UnmarshalListWidgetIntegrationsResponse(bytes)
//    bytes, err = listWidgetIntegrationsResponse.Marshal()
//
//    generateMagicFromPromptParams, err := UnmarshalGenerateMagicFromPromptParams(bytes)
//    bytes, err = generateMagicFromPromptParams.Marshal()
//
//    generateMagicFromPromptRequest, err := UnmarshalGenerateMagicFromPromptRequest(bytes)
//    bytes, err = generateMagicFromPromptRequest.Marshal()
//
//    generateMagicFromPromptResponse, err := UnmarshalGenerateMagicFromPromptResponse(bytes)
//    bytes, err = generateMagicFromPromptResponse.Marshal()
//
//    listMagicPromptsParams, err := UnmarshalListMagicPromptsParams(bytes)
//    bytes, err = listMagicPromptsParams.Marshal()
//
//    listMagicPromptsResponse, err := UnmarshalListMagicPromptsResponse(bytes)
//    bytes, err = listMagicPromptsResponse.Marshal()
//
//    deleteMemoryParams, err := UnmarshalDeleteMemoryParams(bytes)
//    bytes, err = deleteMemoryParams.Marshal()
//
//    deleteMemoryRequest, err := UnmarshalDeleteMemoryRequest(bytes)
//    bytes, err = deleteMemoryRequest.Marshal()
//
//    deleteMemoryResponse, err := UnmarshalDeleteMemoryResponse(bytes)
//    bytes, err = deleteMemoryResponse.Marshal()
//
//    fetchMemoryParams, err := UnmarshalFetchMemoryParams(bytes)
//    bytes, err = fetchMemoryParams.Marshal()
//
//    fetchMemoryResponse, err := UnmarshalFetchMemoryResponse(bytes)
//    bytes, err = fetchMemoryResponse.Marshal()
//
//    updateMemoryParams, err := UnmarshalUpdateMemoryParams(bytes)
//    bytes, err = updateMemoryParams.Marshal()
//
//    updateMemoryRequest, err := UnmarshalUpdateMemoryRequest(bytes)
//    bytes, err = updateMemoryRequest.Marshal()
//
//    updateMemoryResponse, err := UnmarshalUpdateMemoryResponse(bytes)
//    bytes, err = updateMemoryResponse.Marshal()
//
//    createMemoryRequest, err := UnmarshalCreateMemoryRequest(bytes)
//    bytes, err = createMemoryRequest.Marshal()
//
//    createMemoryResponse, err := UnmarshalCreateMemoryResponse(bytes)
//    bytes, err = createMemoryResponse.Marshal()
//
//    exportMemoriesParams, err := UnmarshalExportMemoriesParams(bytes)
//    bytes, err = exportMemoriesParams.Marshal()
//
//    exportMemoriesResponse, err := UnmarshalExportMemoriesResponse(bytes)
//    bytes, err = exportMemoriesResponse.Marshal()
//
//    listMemoriesParams, err := UnmarshalListMemoriesParams(bytes)
//    bytes, err = listMemoriesParams.Marshal()
//
//    listMemoriesResponse, err := UnmarshalListMemoriesResponse(bytes)
//    bytes, err = listMemoriesResponse.Marshal()
//
//    searchMemoryRequest, err := UnmarshalSearchMemoryRequest(bytes)
//    bytes, err = searchMemoryRequest.Marshal()
//
//    searchMemoryResponse, err := UnmarshalSearchMemoryResponse(bytes)
//    bytes, err = searchMemoryResponse.Marshal()
//
//    deletePartnerUserParams, err := UnmarshalDeletePartnerUserParams(bytes)
//    bytes, err = deletePartnerUserParams.Marshal()
//
//    deletePartnerUserRequest, err := UnmarshalDeletePartnerUserRequest(bytes)
//    bytes, err = deletePartnerUserRequest.Marshal()
//
//    deletePartnerUserResponse, err := UnmarshalDeletePartnerUserResponse(bytes)
//    bytes, err = deletePartnerUserResponse.Marshal()
//
//    fetchPartnerUserParams, err := UnmarshalFetchPartnerUserParams(bytes)
//    bytes, err = fetchPartnerUserParams.Marshal()
//
//    fetchPartnerUserResponse, err := UnmarshalFetchPartnerUserResponse(bytes)
//    bytes, err = fetchPartnerUserResponse.Marshal()
//
//    deletePartnerUserTokenParams, err := UnmarshalDeletePartnerUserTokenParams(bytes)
//    bytes, err = deletePartnerUserTokenParams.Marshal()
//
//    deletePartnerUserTokenRequest, err := UnmarshalDeletePartnerUserTokenRequest(bytes)
//    bytes, err = deletePartnerUserTokenRequest.Marshal()
//
//    deletePartnerUserTokenResponse, err := UnmarshalDeletePartnerUserTokenResponse(bytes)
//    bytes, err = deletePartnerUserTokenResponse.Marshal()
//
//    createPartnerUserTokenParams, err := UnmarshalCreatePartnerUserTokenParams(bytes)
//    bytes, err = createPartnerUserTokenParams.Marshal()
//
//    createPartnerUserTokenRequest, err := UnmarshalCreatePartnerUserTokenRequest(bytes)
//    bytes, err = createPartnerUserTokenRequest.Marshal()
//
//    createPartnerUserTokenResponse, err := UnmarshalCreatePartnerUserTokenResponse(bytes)
//    bytes, err = createPartnerUserTokenResponse.Marshal()
//
//    listPartnerUserTokensParams, err := UnmarshalListPartnerUserTokensParams(bytes)
//    bytes, err = listPartnerUserTokensParams.Marshal()
//
//    listPartnerUserTokensResponse, err := UnmarshalListPartnerUserTokensResponse(bytes)
//    bytes, err = listPartnerUserTokensResponse.Marshal()
//
//    updatePartnerUserParams, err := UnmarshalUpdatePartnerUserParams(bytes)
//    bytes, err = updatePartnerUserParams.Marshal()
//
//    updatePartnerUserRequest, err := UnmarshalUpdatePartnerUserRequest(bytes)
//    bytes, err = updatePartnerUserRequest.Marshal()
//
//    updatePartnerUserResponse, err := UnmarshalUpdatePartnerUserResponse(bytes)
//    bytes, err = updatePartnerUserResponse.Marshal()
//
//    createPartnerUserRequest, err := UnmarshalCreatePartnerUserRequest(bytes)
//    bytes, err = createPartnerUserRequest.Marshal()
//
//    createPartnerUserResponse, err := UnmarshalCreatePartnerUserResponse(bytes)
//    bytes, err = createPartnerUserResponse.Marshal()
//
//    listPartnerUsersParams, err := UnmarshalListPartnerUsersParams(bytes)
//    bytes, err = listPartnerUsersParams.Marshal()
//
//    listPartnerUsersResponse, err := UnmarshalListPartnerUsersResponse(bytes)
//    bytes, err = listPartnerUsersResponse.Marshal()
//
//    listPlatformAbilitiesParams, err := UnmarshalListPlatformAbilitiesParams(bytes)
//    bytes, err = listPlatformAbilitiesParams.Marshal()
//
//    listPlatformAbilitiesResponse, err := UnmarshalListPlatformAbilitiesResponse(bytes)
//    bytes, err = listPlatformAbilitiesResponse.Marshal()
//
//    listPlatformActionsParams, err := UnmarshalListPlatformActionsParams(bytes)
//    bytes, err = listPlatformActionsParams.Marshal()
//
//    listPlatformActionsResponse, err := UnmarshalListPlatformActionsResponse(bytes)
//    bytes, err = listPlatformActionsResponse.Marshal()
//
//    fetchPlatformDocParams, err := UnmarshalFetchPlatformDocParams(bytes)
//    bytes, err = fetchPlatformDocParams.Marshal()
//
//    fetchPlatformDocResponse, err := UnmarshalFetchPlatformDocResponse(bytes)
//    bytes, err = fetchPlatformDocResponse.Marshal()
//
//    listPlatformDocsParams, err := UnmarshalListPlatformDocsParams(bytes)
//    bytes, err = listPlatformDocsParams.Marshal()
//
//    listPlatformDocsResponse, err := UnmarshalListPlatformDocsResponse(bytes)
//    bytes, err = listPlatformDocsResponse.Marshal()
//
//    searchPlatformDocsRequest, err := UnmarshalSearchPlatformDocsRequest(bytes)
//    bytes, err = searchPlatformDocsRequest.Marshal()
//
//    searchPlatformDocsResponse, err := UnmarshalSearchPlatformDocsResponse(bytes)
//    bytes, err = searchPlatformDocsResponse.Marshal()
//
//    clonePlatformExampleParams, err := UnmarshalClonePlatformExampleParams(bytes)
//    bytes, err = clonePlatformExampleParams.Marshal()
//
//    clonePlatformExampleRequest, err := UnmarshalClonePlatformExampleRequest(bytes)
//    bytes, err = clonePlatformExampleRequest.Marshal()
//
//    clonePlatformExampleResponse, err := UnmarshalClonePlatformExampleResponse(bytes)
//    bytes, err = clonePlatformExampleResponse.Marshal()
//
//    fetchPlatformExampleParams, err := UnmarshalFetchPlatformExampleParams(bytes)
//    bytes, err = fetchPlatformExampleParams.Marshal()
//
//    fetchPlatformExampleResponse, err := UnmarshalFetchPlatformExampleResponse(bytes)
//    bytes, err = fetchPlatformExampleResponse.Marshal()
//
//    listPlatformExamplesParams, err := UnmarshalListPlatformExamplesParams(bytes)
//    bytes, err = listPlatformExamplesParams.Marshal()
//
//    listPlatformExamplesResponse, err := UnmarshalListPlatformExamplesResponse(bytes)
//    bytes, err = listPlatformExamplesResponse.Marshal()
//
//    searchPlatformExamplesRequest, err := UnmarshalSearchPlatformExamplesRequest(bytes)
//    bytes, err = searchPlatformExamplesRequest.Marshal()
//
//    searchPlatformExamplesResponse, err := UnmarshalSearchPlatformExamplesResponse(bytes)
//    bytes, err = searchPlatformExamplesResponse.Marshal()
//
//    fetchPlatformGuideParams, err := UnmarshalFetchPlatformGuideParams(bytes)
//    bytes, err = fetchPlatformGuideParams.Marshal()
//
//    fetchPlatformGuideResponse, err := UnmarshalFetchPlatformGuideResponse(bytes)
//    bytes, err = fetchPlatformGuideResponse.Marshal()
//
//    listPlatformGuidesParams, err := UnmarshalListPlatformGuidesParams(bytes)
//    bytes, err = listPlatformGuidesParams.Marshal()
//
//    listPlatformGuidesResponse, err := UnmarshalListPlatformGuidesResponse(bytes)
//    bytes, err = listPlatformGuidesResponse.Marshal()
//
//    searchPlatformGuidesRequest, err := UnmarshalSearchPlatformGuidesRequest(bytes)
//    bytes, err = searchPlatformGuidesRequest.Marshal()
//
//    searchPlatformGuidesResponse, err := UnmarshalSearchPlatformGuidesResponse(bytes)
//    bytes, err = searchPlatformGuidesResponse.Marshal()
//
//    fetchPlatformManualParams, err := UnmarshalFetchPlatformManualParams(bytes)
//    bytes, err = fetchPlatformManualParams.Marshal()
//
//    fetchPlatformManualResponse, err := UnmarshalFetchPlatformManualResponse(bytes)
//    bytes, err = fetchPlatformManualResponse.Marshal()
//
//    listPlatformManualsParams, err := UnmarshalListPlatformManualsParams(bytes)
//    bytes, err = listPlatformManualsParams.Marshal()
//
//    listPlatformManualsResponse, err := UnmarshalListPlatformManualsResponse(bytes)
//    bytes, err = listPlatformManualsResponse.Marshal()
//
//    searchPlatformManualsRequest, err := UnmarshalSearchPlatformManualsRequest(bytes)
//    bytes, err = searchPlatformManualsRequest.Marshal()
//
//    searchPlatformManualsResponse, err := UnmarshalSearchPlatformManualsResponse(bytes)
//    bytes, err = searchPlatformManualsResponse.Marshal()
//
//    listPlatformModelsParams, err := UnmarshalListPlatformModelsParams(bytes)
//    bytes, err = listPlatformModelsParams.Marshal()
//
//    listPlatformModelsResponse, err := UnmarshalListPlatformModelsResponse(bytes)
//    bytes, err = listPlatformModelsResponse.Marshal()
//
//    listPlatformSecretsParams, err := UnmarshalListPlatformSecretsParams(bytes)
//    bytes, err = listPlatformSecretsParams.Marshal()
//
//    listPlatformSecretsResponse, err := UnmarshalListPlatformSecretsResponse(bytes)
//    bytes, err = listPlatformSecretsResponse.Marshal()
//
//    fetchPlatformTutorialParams, err := UnmarshalFetchPlatformTutorialParams(bytes)
//    bytes, err = fetchPlatformTutorialParams.Marshal()
//
//    fetchPlatformTutorialResponse, err := UnmarshalFetchPlatformTutorialResponse(bytes)
//    bytes, err = fetchPlatformTutorialResponse.Marshal()
//
//    listPlatformTutorialsParams, err := UnmarshalListPlatformTutorialsParams(bytes)
//    bytes, err = listPlatformTutorialsParams.Marshal()
//
//    listPlatformTutorialsResponse, err := UnmarshalListPlatformTutorialsResponse(bytes)
//    bytes, err = listPlatformTutorialsResponse.Marshal()
//
//    searchPlatformTutorialsRequest, err := UnmarshalSearchPlatformTutorialsRequest(bytes)
//    bytes, err = searchPlatformTutorialsRequest.Marshal()
//
//    searchPlatformTutorialsResponse, err := UnmarshalSearchPlatformTutorialsResponse(bytes)
//    bytes, err = searchPlatformTutorialsResponse.Marshal()
//
//    deletePolicyParams, err := UnmarshalDeletePolicyParams(bytes)
//    bytes, err = deletePolicyParams.Marshal()
//
//    deletePolicyRequest, err := UnmarshalDeletePolicyRequest(bytes)
//    bytes, err = deletePolicyRequest.Marshal()
//
//    deletePolicyResponse, err := UnmarshalDeletePolicyResponse(bytes)
//    bytes, err = deletePolicyResponse.Marshal()
//
//    fetchPolicyParams, err := UnmarshalFetchPolicyParams(bytes)
//    bytes, err = fetchPolicyParams.Marshal()
//
//    fetchPolicyResponse, err := UnmarshalFetchPolicyResponse(bytes)
//    bytes, err = fetchPolicyResponse.Marshal()
//
//    updatePolicyParams, err := UnmarshalUpdatePolicyParams(bytes)
//    bytes, err = updatePolicyParams.Marshal()
//
//    updatePolicyRequest, err := UnmarshalUpdatePolicyRequest(bytes)
//    bytes, err = updatePolicyRequest.Marshal()
//
//    updatePolicyResponse, err := UnmarshalUpdatePolicyResponse(bytes)
//    bytes, err = updatePolicyResponse.Marshal()
//
//    createPolicyRequest, err := UnmarshalCreatePolicyRequest(bytes)
//    bytes, err = createPolicyRequest.Marshal()
//
//    createPolicyResponse, err := UnmarshalCreatePolicyResponse(bytes)
//    bytes, err = createPolicyResponse.Marshal()
//
//    listPoliciesParams, err := UnmarshalListPoliciesParams(bytes)
//    bytes, err = listPoliciesParams.Marshal()
//
//    listPoliciesResponse, err := UnmarshalListPoliciesResponse(bytes)
//    bytes, err = listPoliciesResponse.Marshal()
//
//    deletePortalParams, err := UnmarshalDeletePortalParams(bytes)
//    bytes, err = deletePortalParams.Marshal()
//
//    deletePortalRequest, err := UnmarshalDeletePortalRequest(bytes)
//    bytes, err = deletePortalRequest.Marshal()
//
//    deletePortalResponse, err := UnmarshalDeletePortalResponse(bytes)
//    bytes, err = deletePortalResponse.Marshal()
//
//    fetchPortalParams, err := UnmarshalFetchPortalParams(bytes)
//    bytes, err = fetchPortalParams.Marshal()
//
//    fetchPortalResponse, err := UnmarshalFetchPortalResponse(bytes)
//    bytes, err = fetchPortalResponse.Marshal()
//
//    updatePortalParams, err := UnmarshalUpdatePortalParams(bytes)
//    bytes, err = updatePortalParams.Marshal()
//
//    updatePortalRequest, err := UnmarshalUpdatePortalRequest(bytes)
//    bytes, err = updatePortalRequest.Marshal()
//
//    updatePortalResponse, err := UnmarshalUpdatePortalResponse(bytes)
//    bytes, err = updatePortalResponse.Marshal()
//
//    createPortalRequest, err := UnmarshalCreatePortalRequest(bytes)
//    bytes, err = createPortalRequest.Marshal()
//
//    createPortalResponse, err := UnmarshalCreatePortalResponse(bytes)
//    bytes, err = createPortalResponse.Marshal()
//
//    listPortalsParams, err := UnmarshalListPortalsParams(bytes)
//    bytes, err = listPortalsParams.Marshal()
//
//    listPortalsResponse, err := UnmarshalListPortalsResponse(bytes)
//    bytes, err = listPortalsResponse.Marshal()
//
//    authenticateSecretParams, err := UnmarshalAuthenticateSecretParams(bytes)
//    bytes, err = authenticateSecretParams.Marshal()
//
//    authenticateSecretRequest, err := UnmarshalAuthenticateSecretRequest(bytes)
//    bytes, err = authenticateSecretRequest.Marshal()
//
//    authenticateSecretResponse, err := UnmarshalAuthenticateSecretResponse(bytes)
//    bytes, err = authenticateSecretResponse.Marshal()
//
//    deleteSecretParams, err := UnmarshalDeleteSecretParams(bytes)
//    bytes, err = deleteSecretParams.Marshal()
//
//    deleteSecretRequest, err := UnmarshalDeleteSecretRequest(bytes)
//    bytes, err = deleteSecretRequest.Marshal()
//
//    deleteSecretResponse, err := UnmarshalDeleteSecretResponse(bytes)
//    bytes, err = deleteSecretResponse.Marshal()
//
//    fetchSecretParams, err := UnmarshalFetchSecretParams(bytes)
//    bytes, err = fetchSecretParams.Marshal()
//
//    fetchSecretResponse, err := UnmarshalFetchSecretResponse(bytes)
//    bytes, err = fetchSecretResponse.Marshal()
//
//    revokeSecretParams, err := UnmarshalRevokeSecretParams(bytes)
//    bytes, err = revokeSecretParams.Marshal()
//
//    revokeSecretRequest, err := UnmarshalRevokeSecretRequest(bytes)
//    bytes, err = revokeSecretRequest.Marshal()
//
//    revokeSecretResponse, err := UnmarshalRevokeSecretResponse(bytes)
//    bytes, err = revokeSecretResponse.Marshal()
//
//    updateSecretParams, err := UnmarshalUpdateSecretParams(bytes)
//    bytes, err = updateSecretParams.Marshal()
//
//    updateSecretRequest, err := UnmarshalUpdateSecretRequest(bytes)
//    bytes, err = updateSecretRequest.Marshal()
//
//    updateSecretResponse, err := UnmarshalUpdateSecretResponse(bytes)
//    bytes, err = updateSecretResponse.Marshal()
//
//    verifySecretParams, err := UnmarshalVerifySecretParams(bytes)
//    bytes, err = verifySecretParams.Marshal()
//
//    verifySecretRequest, err := UnmarshalVerifySecretRequest(bytes)
//    bytes, err = verifySecretRequest.Marshal()
//
//    verifySecretResponse, err := UnmarshalVerifySecretResponse(bytes)
//    bytes, err = verifySecretResponse.Marshal()
//
//    createSecretRequest, err := UnmarshalCreateSecretRequest(bytes)
//    bytes, err = createSecretRequest.Marshal()
//
//    createSecretResponse, err := UnmarshalCreateSecretResponse(bytes)
//    bytes, err = createSecretResponse.Marshal()
//
//    listSecretsParams, err := UnmarshalListSecretsParams(bytes)
//    bytes, err = listSecretsParams.Marshal()
//
//    listSecretsResponse, err := UnmarshalListSecretsResponse(bytes)
//    bytes, err = listSecretsResponse.Marshal()
//
//    deleteSkillsetAbilityParams, err := UnmarshalDeleteSkillsetAbilityParams(bytes)
//    bytes, err = deleteSkillsetAbilityParams.Marshal()
//
//    deleteSkillsetAbilityRequest, err := UnmarshalDeleteSkillsetAbilityRequest(bytes)
//    bytes, err = deleteSkillsetAbilityRequest.Marshal()
//
//    deleteSkillsetAbilityResponse, err := UnmarshalDeleteSkillsetAbilityResponse(bytes)
//    bytes, err = deleteSkillsetAbilityResponse.Marshal()
//
//    executeSkillsetAbilityParams, err := UnmarshalExecuteSkillsetAbilityParams(bytes)
//    bytes, err = executeSkillsetAbilityParams.Marshal()
//
//    executeSkillsetAbilityRequest, err := UnmarshalExecuteSkillsetAbilityRequest(bytes)
//    bytes, err = executeSkillsetAbilityRequest.Marshal()
//
//    executeSkillsetAbilityResponse, err := UnmarshalExecuteSkillsetAbilityResponse(bytes)
//    bytes, err = executeSkillsetAbilityResponse.Marshal()
//
//    fetchSkillsetAbilityParams, err := UnmarshalFetchSkillsetAbilityParams(bytes)
//    bytes, err = fetchSkillsetAbilityParams.Marshal()
//
//    fetchSkillsetAbilityResponse, err := UnmarshalFetchSkillsetAbilityResponse(bytes)
//    bytes, err = fetchSkillsetAbilityResponse.Marshal()
//
//    updateSkillsetAbilityParams, err := UnmarshalUpdateSkillsetAbilityParams(bytes)
//    bytes, err = updateSkillsetAbilityParams.Marshal()
//
//    updateSkillsetAbilityRequest, err := UnmarshalUpdateSkillsetAbilityRequest(bytes)
//    bytes, err = updateSkillsetAbilityRequest.Marshal()
//
//    updateSkillsetAbilityResponse, err := UnmarshalUpdateSkillsetAbilityResponse(bytes)
//    bytes, err = updateSkillsetAbilityResponse.Marshal()
//
//    createSkillsetAbilityParams, err := UnmarshalCreateSkillsetAbilityParams(bytes)
//    bytes, err = createSkillsetAbilityParams.Marshal()
//
//    createSkillsetAbilityRequest, err := UnmarshalCreateSkillsetAbilityRequest(bytes)
//    bytes, err = createSkillsetAbilityRequest.Marshal()
//
//    createSkillsetAbilityResponse, err := UnmarshalCreateSkillsetAbilityResponse(bytes)
//    bytes, err = createSkillsetAbilityResponse.Marshal()
//
//    exportSkillsetAbilitiesParams, err := UnmarshalExportSkillsetAbilitiesParams(bytes)
//    bytes, err = exportSkillsetAbilitiesParams.Marshal()
//
//    exportSkillsetAbilitiesResponse, err := UnmarshalExportSkillsetAbilitiesResponse(bytes)
//    bytes, err = exportSkillsetAbilitiesResponse.Marshal()
//
//    listSkillsetAbilitiesParams, err := UnmarshalListSkillsetAbilitiesParams(bytes)
//    bytes, err = listSkillsetAbilitiesParams.Marshal()
//
//    listSkillsetAbilitiesResponse, err := UnmarshalListSkillsetAbilitiesResponse(bytes)
//    bytes, err = listSkillsetAbilitiesResponse.Marshal()
//
//    deleteSkillsetParams, err := UnmarshalDeleteSkillsetParams(bytes)
//    bytes, err = deleteSkillsetParams.Marshal()
//
//    deleteSkillsetRequest, err := UnmarshalDeleteSkillsetRequest(bytes)
//    bytes, err = deleteSkillsetRequest.Marshal()
//
//    deleteSkillsetResponse, err := UnmarshalDeleteSkillsetResponse(bytes)
//    bytes, err = deleteSkillsetResponse.Marshal()
//
//    fetchSkillsetParams, err := UnmarshalFetchSkillsetParams(bytes)
//    bytes, err = fetchSkillsetParams.Marshal()
//
//    fetchSkillsetResponse, err := UnmarshalFetchSkillsetResponse(bytes)
//    bytes, err = fetchSkillsetResponse.Marshal()
//
//    updateSkillsetParams, err := UnmarshalUpdateSkillsetParams(bytes)
//    bytes, err = updateSkillsetParams.Marshal()
//
//    updateSkillsetRequest, err := UnmarshalUpdateSkillsetRequest(bytes)
//    bytes, err = updateSkillsetRequest.Marshal()
//
//    updateSkillsetResponse, err := UnmarshalUpdateSkillsetResponse(bytes)
//    bytes, err = updateSkillsetResponse.Marshal()
//
//    createSkillsetRequest, err := UnmarshalCreateSkillsetRequest(bytes)
//    bytes, err = createSkillsetRequest.Marshal()
//
//    createSkillsetResponse, err := UnmarshalCreateSkillsetResponse(bytes)
//    bytes, err = createSkillsetResponse.Marshal()
//
//    listSkillsetsParams, err := UnmarshalListSkillsetsParams(bytes)
//    bytes, err = listSkillsetsParams.Marshal()
//
//    listSkillsetsResponse, err := UnmarshalListSkillsetsResponse(bytes)
//    bytes, err = listSkillsetsResponse.Marshal()
//
//    fetchSpaceParams, err := UnmarshalFetchSpaceParams(bytes)
//    bytes, err = fetchSpaceParams.Marshal()
//
//    fetchSpaceResponse, err := UnmarshalFetchSpaceResponse(bytes)
//    bytes, err = fetchSpaceResponse.Marshal()
//
//    updateSpaceParams, err := UnmarshalUpdateSpaceParams(bytes)
//    bytes, err = updateSpaceParams.Marshal()
//
//    updateSpaceRequest, err := UnmarshalUpdateSpaceRequest(bytes)
//    bytes, err = updateSpaceRequest.Marshal()
//
//    updateSpaceResponse, err := UnmarshalUpdateSpaceResponse(bytes)
//    bytes, err = updateSpaceResponse.Marshal()
//
//    createSpaceRequest, err := UnmarshalCreateSpaceRequest(bytes)
//    bytes, err = createSpaceRequest.Marshal()
//
//    createSpaceResponse, err := UnmarshalCreateSpaceResponse(bytes)
//    bytes, err = createSpaceResponse.Marshal()
//
//    exportSpacesParams, err := UnmarshalExportSpacesParams(bytes)
//    bytes, err = exportSpacesParams.Marshal()
//
//    exportSpacesResponse, err := UnmarshalExportSpacesResponse(bytes)
//    bytes, err = exportSpacesResponse.Marshal()
//
//    listSpacesParams, err := UnmarshalListSpacesParams(bytes)
//    bytes, err = listSpacesParams.Marshal()
//
//    listSpacesResponse, err := UnmarshalListSpacesResponse(bytes)
//    bytes, err = listSpacesResponse.Marshal()
//
//    deleteTaskParams, err := UnmarshalDeleteTaskParams(bytes)
//    bytes, err = deleteTaskParams.Marshal()
//
//    deleteTaskRequest, err := UnmarshalDeleteTaskRequest(bytes)
//    bytes, err = deleteTaskRequest.Marshal()
//
//    deleteTaskResponse, err := UnmarshalDeleteTaskResponse(bytes)
//    bytes, err = deleteTaskResponse.Marshal()
//
//    fetchTaskParams, err := UnmarshalFetchTaskParams(bytes)
//    bytes, err = fetchTaskParams.Marshal()
//
//    fetchTaskResponse, err := UnmarshalFetchTaskResponse(bytes)
//    bytes, err = fetchTaskResponse.Marshal()
//
//    triggerTaskParams, err := UnmarshalTriggerTaskParams(bytes)
//    bytes, err = triggerTaskParams.Marshal()
//
//    triggerTaskRequest, err := UnmarshalTriggerTaskRequest(bytes)
//    bytes, err = triggerTaskRequest.Marshal()
//
//    triggerTaskResponse, err := UnmarshalTriggerTaskResponse(bytes)
//    bytes, err = triggerTaskResponse.Marshal()
//
//    updateTaskParams, err := UnmarshalUpdateTaskParams(bytes)
//    bytes, err = updateTaskParams.Marshal()
//
//    updateTaskRequest, err := UnmarshalUpdateTaskRequest(bytes)
//    bytes, err = updateTaskRequest.Marshal()
//
//    updateTaskResponse, err := UnmarshalUpdateTaskResponse(bytes)
//    bytes, err = updateTaskResponse.Marshal()
//
//    createTaskRequest, err := UnmarshalCreateTaskRequest(bytes)
//    bytes, err = createTaskRequest.Marshal()
//
//    createTaskResponse, err := UnmarshalCreateTaskResponse(bytes)
//    bytes, err = createTaskResponse.Marshal()
//
//    exportTasksParams, err := UnmarshalExportTasksParams(bytes)
//    bytes, err = exportTasksParams.Marshal()
//
//    exportTasksResponse, err := UnmarshalExportTasksResponse(bytes)
//    bytes, err = exportTasksResponse.Marshal()
//
//    listTasksParams, err := UnmarshalListTasksParams(bytes)
//    bytes, err = listTasksParams.Marshal()
//
//    listTasksResponse, err := UnmarshalListTasksResponse(bytes)
//    bytes, err = listTasksResponse.Marshal()
//
//    listTeamsParams, err := UnmarshalListTeamsParams(bytes)
//    bytes, err = listTeamsParams.Marshal()
//
//    listTeamsResponse, err := UnmarshalListTeamsResponse(bytes)
//    bytes, err = listTeamsResponse.Marshal()
//
//    fetchUsageResponse, err := UnmarshalFetchUsageResponse(bytes)
//    bytes, err = fetchUsageResponse.Marshal()
//
//    fetchUsageSeriesResponse, err := UnmarshalFetchUsageSeriesResponse(bytes)
//    bytes, err = fetchUsageSeriesResponse.Marshal()
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

func UnmarshalListPlatformReportsParams(data []byte) (ListPlatformReportsParams, error) {
	var r ListPlatformReportsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformReportsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformReportsResponse(data []byte) (ListPlatformReportsResponse, error) {
	var r ListPlatformReportsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformReportsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalGenerateReportParams(data []byte) (GenerateReportParams, error) {
	var r GenerateReportParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GenerateReportParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GenerateReportRequest map[string]interface{}

func UnmarshalGenerateReportRequest(data []byte) (GenerateReportRequest, error) {
	var r GenerateReportRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GenerateReportRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GenerateReportResponse map[string]interface{}

func UnmarshalGenerateReportResponse(data []byte) (GenerateReportResponse, error) {
	var r GenerateReportResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GenerateReportResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GenerateReportsRequest map[string]map[string]interface{}

func UnmarshalGenerateReportsRequest(data []byte) (GenerateReportsRequest, error) {
	var r GenerateReportsRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GenerateReportsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GenerateReportsResponse map[string]GenerateReportsResponseValue

func UnmarshalGenerateReportsResponse(data []byte) (GenerateReportsResponse, error) {
	var r GenerateReportsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GenerateReportsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCloneBlueprintParams(data []byte) (CloneBlueprintParams, error) {
	var r CloneBlueprintParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloneBlueprintParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CloneBlueprintRequest map[string]interface{}

func UnmarshalCloneBlueprintRequest(data []byte) (CloneBlueprintRequest, error) {
	var r CloneBlueprintRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloneBlueprintRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCloneBlueprintResponse(data []byte) (CloneBlueprintResponse, error) {
	var r CloneBlueprintResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloneBlueprintResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteBlueprintParams(data []byte) (DeleteBlueprintParams, error) {
	var r DeleteBlueprintParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteBlueprintParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteBlueprintRequest(data []byte) (DeleteBlueprintRequest, error) {
	var r DeleteBlueprintRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteBlueprintRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteBlueprintResponse(data []byte) (DeleteBlueprintResponse, error) {
	var r DeleteBlueprintResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteBlueprintResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchBlueprintParams(data []byte) (FetchBlueprintParams, error) {
	var r FetchBlueprintParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchBlueprintParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchBlueprintResponse(data []byte) (FetchBlueprintResponse, error) {
	var r FetchBlueprintResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchBlueprintResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListBlueprintResourcesParams(data []byte) (ListBlueprintResourcesParams, error) {
	var r ListBlueprintResourcesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListBlueprintResourcesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListBlueprintResourcesResponse(data []byte) (ListBlueprintResourcesResponse, error) {
	var r ListBlueprintResourcesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListBlueprintResourcesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateBlueprintParams(data []byte) (UpdateBlueprintParams, error) {
	var r UpdateBlueprintParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateBlueprintParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateBlueprintRequest(data []byte) (UpdateBlueprintRequest, error) {
	var r UpdateBlueprintRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateBlueprintRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateBlueprintResponse(data []byte) (UpdateBlueprintResponse, error) {
	var r UpdateBlueprintResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateBlueprintResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateBlueprintRequest(data []byte) (CreateBlueprintRequest, error) {
	var r CreateBlueprintRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateBlueprintRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateBlueprintResponse(data []byte) (CreateBlueprintResponse, error) {
	var r CreateBlueprintResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateBlueprintResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListBlueprintsParams(data []byte) (ListBlueprintsParams, error) {
	var r ListBlueprintsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListBlueprintsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListBlueprintsResponse(data []byte) (ListBlueprintsResponse, error) {
	var r ListBlueprintsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListBlueprintsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCloneBotParams(data []byte) (CloneBotParams, error) {
	var r CloneBotParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloneBotParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CloneBotRequest interface{}

func UnmarshalCloneBotRequest(data []byte) (CloneBotRequest, error) {
	var r CloneBotRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloneBotRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCloneBotResponse(data []byte) (CloneBotResponse, error) {
	var r CloneBotResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloneBotResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteBotParams(data []byte) (DeleteBotParams, error) {
	var r DeleteBotParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteBotParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteBotRequest map[string]interface{}

func UnmarshalDeleteBotRequest(data []byte) (DeleteBotRequest, error) {
	var r DeleteBotRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteBotRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteBotResponse(data []byte) (DeleteBotResponse, error) {
	var r DeleteBotResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteBotResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownvoteBotParams(data []byte) (DownvoteBotParams, error) {
	var r DownvoteBotParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownvoteBotParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownvoteBotRequest(data []byte) (DownvoteBotRequest, error) {
	var r DownvoteBotRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownvoteBotRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownvoteBotResponse(data []byte) (DownvoteBotResponse, error) {
	var r DownvoteBotResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownvoteBotResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchBotParams(data []byte) (FetchBotParams, error) {
	var r FetchBotParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchBotParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchBotResponse(data []byte) (FetchBotResponse, error) {
	var r FetchBotResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchBotResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchBotMemoryParams(data []byte) (SearchBotMemoryParams, error) {
	var r SearchBotMemoryParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchBotMemoryParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchBotMemoryRequest(data []byte) (SearchBotMemoryRequest, error) {
	var r SearchBotMemoryRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchBotMemoryRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchBotMemoryResponse(data []byte) (SearchBotMemoryResponse, error) {
	var r SearchBotMemoryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchBotMemoryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateBotSessionParams(data []byte) (CreateBotSessionParams, error) {
	var r CreateBotSessionParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateBotSessionParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateBotSessionRequest(data []byte) (CreateBotSessionRequest, error) {
	var r CreateBotSessionRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateBotSessionRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateBotSessionResponse(data []byte) (CreateBotSessionResponse, error) {
	var r CreateBotSessionResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateBotSessionResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateBotParams(data []byte) (UpdateBotParams, error) {
	var r UpdateBotParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateBotParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateBotRequest(data []byte) (UpdateBotRequest, error) {
	var r UpdateBotRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateBotRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateBotResponse(data []byte) (UpdateBotResponse, error) {
	var r UpdateBotResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateBotResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpvoteBotParams(data []byte) (UpvoteBotParams, error) {
	var r UpvoteBotParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpvoteBotParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpvoteBotRequest(data []byte) (UpvoteBotRequest, error) {
	var r UpvoteBotRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpvoteBotRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpvoteBotResponse(data []byte) (UpvoteBotResponse, error) {
	var r UpvoteBotResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpvoteBotResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchBotUsageParams(data []byte) (FetchBotUsageParams, error) {
	var r FetchBotUsageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchBotUsageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchBotUsageResponse(data []byte) (FetchBotUsageResponse, error) {
	var r FetchBotUsageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchBotUsageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateBotRequest(data []byte) (CreateBotRequest, error) {
	var r CreateBotRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateBotRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateBotResponse(data []byte) (CreateBotResponse, error) {
	var r CreateBotResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateBotResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListBotsParams(data []byte) (ListBotsParams, error) {
	var r ListBotsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListBotsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListBotsResponse(data []byte) (ListBotsResponse, error) {
	var r ListBotsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListBotsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPublishChannelMessageParams(data []byte) (PublishChannelMessageParams, error) {
	var r PublishChannelMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PublishChannelMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPublishChannelMessageRequest(data []byte) (PublishChannelMessageRequest, error) {
	var r PublishChannelMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PublishChannelMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPublishChannelMessageResponse(data []byte) (PublishChannelMessageResponse, error) {
	var r PublishChannelMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PublishChannelMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSubscribeChannelMessagesParams(data []byte) (SubscribeChannelMessagesParams, error) {
	var r SubscribeChannelMessagesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SubscribeChannelMessagesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSubscribeChannelMessagesRequest(data []byte) (SubscribeChannelMessagesRequest, error) {
	var r SubscribeChannelMessagesRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SubscribeChannelMessagesRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactConversationsParams(data []byte) (ListContactConversationsParams, error) {
	var r ListContactConversationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactConversationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactConversationsResponse(data []byte) (ListContactConversationsResponse, error) {
	var r ListContactConversationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactConversationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteContactParams(data []byte) (DeleteContactParams, error) {
	var r DeleteContactParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteContactParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteContactRequest map[string]interface{}

func UnmarshalDeleteContactRequest(data []byte) (DeleteContactRequest, error) {
	var r DeleteContactRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteContactRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteContactResponse(data []byte) (DeleteContactResponse, error) {
	var r DeleteContactResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteContactResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchContactParams(data []byte) (FetchContactParams, error) {
	var r FetchContactParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchContactParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchContactResponse(data []byte) (FetchContactResponse, error) {
	var r FetchContactResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchContactResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactMemoriesParams(data []byte) (ListContactMemoriesParams, error) {
	var r ListContactMemoriesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactMemoriesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactMemoriesResponse(data []byte) (ListContactMemoriesResponse, error) {
	var r ListContactMemoriesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactMemoriesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchContactMemoryParams(data []byte) (SearchContactMemoryParams, error) {
	var r SearchContactMemoryParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchContactMemoryParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchContactMemoryRequest(data []byte) (SearchContactMemoryRequest, error) {
	var r SearchContactMemoryRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchContactMemoryRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchContactMemoryResponse(data []byte) (SearchContactMemoryResponse, error) {
	var r SearchContactMemoryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchContactMemoryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAuthenticateContactSecretParams(data []byte) (AuthenticateContactSecretParams, error) {
	var r AuthenticateContactSecretParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthenticateContactSecretParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuthenticateContactSecretRequest map[string]interface{}

func UnmarshalAuthenticateContactSecretRequest(data []byte) (AuthenticateContactSecretRequest, error) {
	var r AuthenticateContactSecretRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthenticateContactSecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAuthenticateContactSecretResponse(data []byte) (AuthenticateContactSecretResponse, error) {
	var r AuthenticateContactSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthenticateContactSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRevokeContactSecretParams(data []byte) (RevokeContactSecretParams, error) {
	var r RevokeContactSecretParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RevokeContactSecretParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RevokeContactSecretRequest map[string]interface{}

func UnmarshalRevokeContactSecretRequest(data []byte) (RevokeContactSecretRequest, error) {
	var r RevokeContactSecretRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RevokeContactSecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRevokeContactSecretResponse(data []byte) (RevokeContactSecretResponse, error) {
	var r RevokeContactSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RevokeContactSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalVerifyContactSecretParams(data []byte) (VerifyContactSecretParams, error) {
	var r VerifyContactSecretParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VerifyContactSecretParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type VerifyContactSecretRequest map[string]interface{}

func UnmarshalVerifyContactSecretRequest(data []byte) (VerifyContactSecretRequest, error) {
	var r VerifyContactSecretRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VerifyContactSecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalVerifyContactSecretResponse(data []byte) (VerifyContactSecretResponse, error) {
	var r VerifyContactSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VerifyContactSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactSecretsParams(data []byte) (ListContactSecretsParams, error) {
	var r ListContactSecretsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactSecretsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactSecretsResponse(data []byte) (ListContactSecretsResponse, error) {
	var r ListContactSecretsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactSecretsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactSpacesParams(data []byte) (ListContactSpacesParams, error) {
	var r ListContactSpacesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactSpacesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactSpacesResponse(data []byte) (ListContactSpacesResponse, error) {
	var r ListContactSpacesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactSpacesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactTasksParams(data []byte) (ListContactTasksParams, error) {
	var r ListContactTasksParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactTasksParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactTasksResponse(data []byte) (ListContactTasksResponse, error) {
	var r ListContactTasksResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactTasksResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateContactParams(data []byte) (UpdateContactParams, error) {
	var r UpdateContactParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateContactParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateContactRequest(data []byte) (UpdateContactRequest, error) {
	var r UpdateContactRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateContactRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateContactResponse(data []byte) (UpdateContactResponse, error) {
	var r UpdateContactResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateContactResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateContactRequest(data []byte) (CreateContactRequest, error) {
	var r CreateContactRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateContactRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateContactResponse(data []byte) (CreateContactResponse, error) {
	var r CreateContactResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateContactResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEnsureContactRequest(data []byte) (EnsureContactRequest, error) {
	var r EnsureContactRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EnsureContactRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEnsureContactResponse(data []byte) (EnsureContactResponse, error) {
	var r EnsureContactResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EnsureContactResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportContactsParams(data []byte) (ExportContactsParams, error) {
	var r ExportContactsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportContactsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportContactsResponse(data []byte) (ExportContactsResponse, error) {
	var r ExportContactsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportContactsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactsParams(data []byte) (ListContactsParams, error) {
	var r ListContactsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListContactsResponse(data []byte) (ListContactsResponse, error) {
	var r ListContactsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListContactsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUploadConversationAttachmentParams(data []byte) (UploadConversationAttachmentParams, error) {
	var r UploadConversationAttachmentParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UploadConversationAttachmentParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUploadConversationAttachmentRequest(data []byte) (UploadConversationAttachmentRequest, error) {
	var r UploadConversationAttachmentRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UploadConversationAttachmentRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUploadConversationAttachmentResponse(data []byte) (UploadConversationAttachmentResponse, error) {
	var r UploadConversationAttachmentResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UploadConversationAttachmentResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCompleteConversationMessageParams(data []byte) (CompleteConversationMessageParams, error) {
	var r CompleteConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompleteConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCompleteConversationMessageRequest(data []byte) (CompleteConversationMessageRequest, error) {
	var r CompleteConversationMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompleteConversationMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCompleteConversationMessageResponse(data []byte) (CompleteConversationMessageResponse, error) {
	var r CompleteConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompleteConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpsertConversationContactParams(data []byte) (UpsertConversationContactParams, error) {
	var r UpsertConversationContactParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpsertConversationContactParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpsertConversationContactRequest(data []byte) (UpsertConversationContactRequest, error) {
	var r UpsertConversationContactRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpsertConversationContactRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpsertConversationContactResponse(data []byte) (UpsertConversationContactResponse, error) {
	var r UpsertConversationContactResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpsertConversationContactResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteConversationParams(data []byte) (DeleteConversationParams, error) {
	var r DeleteConversationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteConversationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteConversationRequest map[string]interface{}

func UnmarshalDeleteConversationRequest(data []byte) (DeleteConversationRequest, error) {
	var r DeleteConversationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteConversationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteConversationResponse(data []byte) (DeleteConversationResponse, error) {
	var r DeleteConversationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteConversationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDispatchStatefulConversationRequest(data []byte) (DispatchStatefulConversationRequest, error) {
	var r DispatchStatefulConversationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DispatchStatefulConversationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDispatchStatefulConversationResponse(data []byte) (DispatchStatefulConversationResponse, error) {
	var r DispatchStatefulConversationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DispatchStatefulConversationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownvoteConversationParams(data []byte) (DownvoteConversationParams, error) {
	var r DownvoteConversationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownvoteConversationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownvoteConversationRequest(data []byte) (DownvoteConversationRequest, error) {
	var r DownvoteConversationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownvoteConversationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownvoteConversationResponse(data []byte) (DownvoteConversationResponse, error) {
	var r DownvoteConversationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownvoteConversationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchConversationParams(data []byte) (FetchConversationParams, error) {
	var r FetchConversationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchConversationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchConversationResponse(data []byte) (FetchConversationResponse, error) {
	var r FetchConversationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchConversationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteConversationMessageParams(data []byte) (DeleteConversationMessageParams, error) {
	var r DeleteConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteConversationMessageRequest map[string]interface{}

func UnmarshalDeleteConversationMessageRequest(data []byte) (DeleteConversationMessageRequest, error) {
	var r DeleteConversationMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteConversationMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteConversationMessageResponse(data []byte) (DeleteConversationMessageResponse, error) {
	var r DeleteConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownvoteConversationMessageParams(data []byte) (DownvoteConversationMessageParams, error) {
	var r DownvoteConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownvoteConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownvoteConversationMessageRequest(data []byte) (DownvoteConversationMessageRequest, error) {
	var r DownvoteConversationMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownvoteConversationMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownvoteConversationMessageResponse(data []byte) (DownvoteConversationMessageResponse, error) {
	var r DownvoteConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownvoteConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchConversationMessageParams(data []byte) (FetchConversationMessageParams, error) {
	var r FetchConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchConversationMessageResponse(data []byte) (FetchConversationMessageResponse, error) {
	var r FetchConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSynthesizeConversationMessageParams(data []byte) (SynthesizeConversationMessageParams, error) {
	var r SynthesizeConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SynthesizeConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SynthesizeConversationMessageRequest map[string]interface{}

func UnmarshalSynthesizeConversationMessageRequest(data []byte) (SynthesizeConversationMessageRequest, error) {
	var r SynthesizeConversationMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SynthesizeConversationMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSynthesizeConversationMessageResponse(data []byte) (SynthesizeConversationMessageResponse, error) {
	var r SynthesizeConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SynthesizeConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateConversationMessageParams(data []byte) (UpdateConversationMessageParams, error) {
	var r UpdateConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateConversationMessageRequest(data []byte) (UpdateConversationMessageRequest, error) {
	var r UpdateConversationMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateConversationMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateConversationMessageResponse(data []byte) (UpdateConversationMessageResponse, error) {
	var r UpdateConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpvoteConversationMessageParams(data []byte) (UpvoteConversationMessageParams, error) {
	var r UpvoteConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpvoteConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpvoteConversationMessageRequest(data []byte) (UpvoteConversationMessageRequest, error) {
	var r UpvoteConversationMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpvoteConversationMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpvoteConversationMessageResponse(data []byte) (UpvoteConversationMessageResponse, error) {
	var r UpvoteConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpvoteConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateConversationMessageParams(data []byte) (CreateConversationMessageParams, error) {
	var r CreateConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateConversationMessageRequest(data []byte) (CreateConversationMessageRequest, error) {
	var r CreateConversationMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateConversationMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateConversationMessageResponse(data []byte) (CreateConversationMessageResponse, error) {
	var r CreateConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListConversationMessagesParams(data []byte) (ListConversationMessagesParams, error) {
	var r ListConversationMessagesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListConversationMessagesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListConversationMessagesResponse(data []byte) (ListConversationMessagesResponse, error) {
	var r ListConversationMessagesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListConversationMessagesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalReceiveConversationMessageParams(data []byte) (ReceiveConversationMessageParams, error) {
	var r ReceiveConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReceiveConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalReceiveConversationMessageRequest(data []byte) (ReceiveConversationMessageRequest, error) {
	var r ReceiveConversationMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReceiveConversationMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalReceiveConversationMessageResponse(data []byte) (ReceiveConversationMessageResponse, error) {
	var r ReceiveConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ReceiveConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSendConversationMessageParams(data []byte) (SendConversationMessageParams, error) {
	var r SendConversationMessageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SendConversationMessageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSendConversationMessageRequest(data []byte) (SendConversationMessageRequest, error) {
	var r SendConversationMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SendConversationMessageRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSendConversationMessageResponse(data []byte) (SendConversationMessageResponse, error) {
	var r SendConversationMessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SendConversationMessageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateConversationSessionParams(data []byte) (CreateConversationSessionParams, error) {
	var r CreateConversationSessionParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateConversationSessionParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateConversationSessionRequest(data []byte) (CreateConversationSessionRequest, error) {
	var r CreateConversationSessionRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateConversationSessionRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateConversationSessionResponse(data []byte) (CreateConversationSessionResponse, error) {
	var r CreateConversationSessionResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateConversationSessionResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateConversationParams(data []byte) (UpdateConversationParams, error) {
	var r UpdateConversationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateConversationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateConversationRequest(data []byte) (UpdateConversationRequest, error) {
	var r UpdateConversationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateConversationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateConversationResponse(data []byte) (UpdateConversationResponse, error) {
	var r UpdateConversationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateConversationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpvoteConversationParams(data []byte) (UpvoteConversationParams, error) {
	var r UpvoteConversationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpvoteConversationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpvoteConversationRequest(data []byte) (UpvoteConversationRequest, error) {
	var r UpvoteConversationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpvoteConversationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpvoteConversationResponse(data []byte) (UpvoteConversationResponse, error) {
	var r UpvoteConversationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpvoteConversationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchConversationUsageParams(data []byte) (FetchConversationUsageParams, error) {
	var r FetchConversationUsageParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchConversationUsageParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchConversationUsageResponse(data []byte) (FetchConversationUsageResponse, error) {
	var r FetchConversationUsageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchConversationUsageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCompleteConversationRequest(data []byte) (CompleteConversationRequest, error) {
	var r CompleteConversationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompleteConversationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCompleteConversationResponse(data []byte) (CompleteConversationResponse, error) {
	var r CompleteConversationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompleteConversationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateConversationRequest(data []byte) (CreateConversationRequest, error) {
	var r CreateConversationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateConversationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateConversationResponse(data []byte) (CreateConversationResponse, error) {
	var r CreateConversationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateConversationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDispatchConversationRequest(data []byte) (DispatchConversationRequest, error) {
	var r DispatchConversationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DispatchConversationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDispatchConversationResponse(data []byte) (DispatchConversationResponse, error) {
	var r DispatchConversationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DispatchConversationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportConversationsParams(data []byte) (ExportConversationsParams, error) {
	var r ExportConversationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportConversationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportConversationsResponse(data []byte) (ExportConversationsResponse, error) {
	var r ExportConversationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportConversationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListConversationsParams(data []byte) (ListConversationsParams, error) {
	var r ListConversationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListConversationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListConversationsResponse(data []byte) (ListConversationsResponse, error) {
	var r ListConversationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListConversationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteDatasetParams(data []byte) (DeleteDatasetParams, error) {
	var r DeleteDatasetParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteDatasetParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteDatasetRequest map[string]interface{}

func UnmarshalDeleteDatasetRequest(data []byte) (DeleteDatasetRequest, error) {
	var r DeleteDatasetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteDatasetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteDatasetResponse(data []byte) (DeleteDatasetResponse, error) {
	var r DeleteDatasetResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteDatasetResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchDatasetParams(data []byte) (FetchDatasetParams, error) {
	var r FetchDatasetParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchDatasetParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchDatasetResponse(data []byte) (FetchDatasetResponse, error) {
	var r FetchDatasetResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchDatasetResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAttachDatasetFileParams(data []byte) (AttachDatasetFileParams, error) {
	var r AttachDatasetFileParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AttachDatasetFileParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAttachDatasetFileRequest(data []byte) (AttachDatasetFileRequest, error) {
	var r AttachDatasetFileRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AttachDatasetFileRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAttachDatasetFileResponse(data []byte) (AttachDatasetFileResponse, error) {
	var r AttachDatasetFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AttachDatasetFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDetachDatasetFileParams(data []byte) (DetachDatasetFileParams, error) {
	var r DetachDatasetFileParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DetachDatasetFileParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDetachDatasetFileRequest(data []byte) (DetachDatasetFileRequest, error) {
	var r DetachDatasetFileRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DetachDatasetFileRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDetachDatasetFileResponse(data []byte) (DetachDatasetFileResponse, error) {
	var r DetachDatasetFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DetachDatasetFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSyncDatasetFileParams(data []byte) (SyncDatasetFileParams, error) {
	var r SyncDatasetFileParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncDatasetFileParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SyncDatasetFileRequest map[string]interface{}

func UnmarshalSyncDatasetFileRequest(data []byte) (SyncDatasetFileRequest, error) {
	var r SyncDatasetFileRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncDatasetFileRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSyncDatasetFileResponse(data []byte) (SyncDatasetFileResponse, error) {
	var r SyncDatasetFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncDatasetFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListDatasetFilesParams(data []byte) (ListDatasetFilesParams, error) {
	var r ListDatasetFilesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListDatasetFilesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListDatasetFilesResponse(data []byte) (ListDatasetFilesResponse, error) {
	var r ListDatasetFilesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListDatasetFilesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteDatasetRecordParams(data []byte) (DeleteDatasetRecordParams, error) {
	var r DeleteDatasetRecordParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteDatasetRecordParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteDatasetRecordRequest map[string]interface{}

func UnmarshalDeleteDatasetRecordRequest(data []byte) (DeleteDatasetRecordRequest, error) {
	var r DeleteDatasetRecordRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteDatasetRecordRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteDatasetRecordResponse(data []byte) (DeleteDatasetRecordResponse, error) {
	var r DeleteDatasetRecordResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteDatasetRecordResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchDatasetRecordParams(data []byte) (FetchDatasetRecordParams, error) {
	var r FetchDatasetRecordParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchDatasetRecordParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchDatasetRecordResponse(data []byte) (FetchDatasetRecordResponse, error) {
	var r FetchDatasetRecordResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchDatasetRecordResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateDatasetRecordParams(data []byte) (UpdateDatasetRecordParams, error) {
	var r UpdateDatasetRecordParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateDatasetRecordParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateDatasetRecordRequest(data []byte) (UpdateDatasetRecordRequest, error) {
	var r UpdateDatasetRecordRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateDatasetRecordRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateDatasetRecordResponse(data []byte) (UpdateDatasetRecordResponse, error) {
	var r UpdateDatasetRecordResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateDatasetRecordResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateDatasetRecordParams(data []byte) (CreateDatasetRecordParams, error) {
	var r CreateDatasetRecordParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateDatasetRecordParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateDatasetRecordRequest(data []byte) (CreateDatasetRecordRequest, error) {
	var r CreateDatasetRecordRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateDatasetRecordRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateDatasetRecordResponse(data []byte) (CreateDatasetRecordResponse, error) {
	var r CreateDatasetRecordResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateDatasetRecordResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportDatasetRecordsParams(data []byte) (ExportDatasetRecordsParams, error) {
	var r ExportDatasetRecordsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportDatasetRecordsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportDatasetRecordsResponse(data []byte) (ExportDatasetRecordsResponse, error) {
	var r ExportDatasetRecordsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportDatasetRecordsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListDatasetRecordsParams(data []byte) (ListDatasetRecordsParams, error) {
	var r ListDatasetRecordsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListDatasetRecordsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListDatasetRecordsResponse(data []byte) (ListDatasetRecordsResponse, error) {
	var r ListDatasetRecordsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListDatasetRecordsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchDatasetParams(data []byte) (SearchDatasetParams, error) {
	var r SearchDatasetParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchDatasetParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchDatasetRequest(data []byte) (SearchDatasetRequest, error) {
	var r SearchDatasetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchDatasetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchDatasetResponse(data []byte) (SearchDatasetResponse, error) {
	var r SearchDatasetResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchDatasetResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateDatasetParams(data []byte) (UpdateDatasetParams, error) {
	var r UpdateDatasetParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateDatasetParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateDatasetRequest(data []byte) (UpdateDatasetRequest, error) {
	var r UpdateDatasetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateDatasetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateDatasetResponse(data []byte) (UpdateDatasetResponse, error) {
	var r UpdateDatasetResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateDatasetResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateDatasetRequest(data []byte) (CreateDatasetRequest, error) {
	var r CreateDatasetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateDatasetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateDatasetResponse(data []byte) (CreateDatasetResponse, error) {
	var r CreateDatasetResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateDatasetResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListDatasetsParams(data []byte) (ListDatasetsParams, error) {
	var r ListDatasetsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListDatasetsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListDatasetsResponse(data []byte) (ListDatasetsResponse, error) {
	var r ListDatasetsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListDatasetsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportEventLogsParams(data []byte) (ExportEventLogsParams, error) {
	var r ExportEventLogsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportEventLogsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportEventLogsResponse(data []byte) (ExportEventLogsResponse, error) {
	var r ExportEventLogsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportEventLogsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListEventLogsParams(data []byte) (ListEventLogsParams, error) {
	var r ListEventLogsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListEventLogsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListEventLogsResponse(data []byte) (ListEventLogsResponse, error) {
	var r ListEventLogsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListEventLogsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSubscribeEventLogsRequest(data []byte) (SubscribeEventLogsRequest, error) {
	var r SubscribeEventLogsRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SubscribeEventLogsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteFileParams(data []byte) (DeleteFileParams, error) {
	var r DeleteFileParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteFileParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteFileRequest map[string]interface{}

func UnmarshalDeleteFileRequest(data []byte) (DeleteFileRequest, error) {
	var r DeleteFileRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteFileRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteFileResponse(data []byte) (DeleteFileResponse, error) {
	var r DeleteFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownloadFileParams(data []byte) (DownloadFileParams, error) {
	var r DownloadFileParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownloadFileParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDownloadFileResponse(data []byte) (DownloadFileResponse, error) {
	var r DownloadFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DownloadFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchFileParams(data []byte) (FetchFileParams, error) {
	var r FetchFileParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchFileParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchFileResponse(data []byte) (FetchFileResponse, error) {
	var r FetchFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSyncFileParams(data []byte) (SyncFileParams, error) {
	var r SyncFileParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncFileParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SyncFileRequest map[string]interface{}

func UnmarshalSyncFileRequest(data []byte) (SyncFileRequest, error) {
	var r SyncFileRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncFileRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSyncFileResponse(data []byte) (SyncFileResponse, error) {
	var r SyncFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateFileParams(data []byte) (UpdateFileParams, error) {
	var r UpdateFileParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateFileParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateFileRequest(data []byte) (UpdateFileRequest, error) {
	var r UpdateFileRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateFileRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateFileResponse(data []byte) (UpdateFileResponse, error) {
	var r UpdateFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUploadFileParams(data []byte) (UploadFileParams, error) {
	var r UploadFileParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UploadFileParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUploadFileRequest(data []byte) (UploadFileRequest, error) {
	var r UploadFileRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UploadFileRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUploadFileResponse(data []byte) (UploadFileResponse, error) {
	var r UploadFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UploadFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateFileRequest(data []byte) (CreateFileRequest, error) {
	var r CreateFileRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateFileRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateFileResponse(data []byte) (CreateFileResponse, error) {
	var r CreateFileResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateFileResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListFilesParams(data []byte) (ListFilesParams, error) {
	var r ListFilesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListFilesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListFilesResponse(data []byte) (ListFilesResponse, error) {
	var r ListFilesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListFilesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteDiscordIntegrationParams(data []byte) (DeleteDiscordIntegrationParams, error) {
	var r DeleteDiscordIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteDiscordIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteDiscordIntegrationRequest map[string]interface{}

func UnmarshalDeleteDiscordIntegrationRequest(data []byte) (DeleteDiscordIntegrationRequest, error) {
	var r DeleteDiscordIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteDiscordIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteDiscordIntegrationResponse(data []byte) (DeleteDiscordIntegrationResponse, error) {
	var r DeleteDiscordIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteDiscordIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchDiscordIntegrationParams(data []byte) (FetchDiscordIntegrationParams, error) {
	var r FetchDiscordIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchDiscordIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchDiscordIntegrationResponse(data []byte) (FetchDiscordIntegrationResponse, error) {
	var r FetchDiscordIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchDiscordIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupDiscordIntegrationParams(data []byte) (SetupDiscordIntegrationParams, error) {
	var r SetupDiscordIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupDiscordIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupDiscordIntegrationRequest map[string]interface{}

func UnmarshalSetupDiscordIntegrationRequest(data []byte) (SetupDiscordIntegrationRequest, error) {
	var r SetupDiscordIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupDiscordIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupDiscordIntegrationResponse(data []byte) (SetupDiscordIntegrationResponse, error) {
	var r SetupDiscordIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupDiscordIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateDiscordIntegrationParams(data []byte) (UpdateDiscordIntegrationParams, error) {
	var r UpdateDiscordIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateDiscordIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateDiscordIntegrationRequest(data []byte) (UpdateDiscordIntegrationRequest, error) {
	var r UpdateDiscordIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateDiscordIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateDiscordIntegrationResponse(data []byte) (UpdateDiscordIntegrationResponse, error) {
	var r UpdateDiscordIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateDiscordIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateDiscordIntegrationRequest(data []byte) (CreateDiscordIntegrationRequest, error) {
	var r CreateDiscordIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateDiscordIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateDiscordIntegrationResponse(data []byte) (CreateDiscordIntegrationResponse, error) {
	var r CreateDiscordIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateDiscordIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListDiscordIntegrationsParams(data []byte) (ListDiscordIntegrationsParams, error) {
	var r ListDiscordIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListDiscordIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListDiscordIntegrationsResponse(data []byte) (ListDiscordIntegrationsResponse, error) {
	var r ListDiscordIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListDiscordIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteEmailIntegrationParams(data []byte) (DeleteEmailIntegrationParams, error) {
	var r DeleteEmailIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteEmailIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteEmailIntegrationRequest map[string]interface{}

func UnmarshalDeleteEmailIntegrationRequest(data []byte) (DeleteEmailIntegrationRequest, error) {
	var r DeleteEmailIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteEmailIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteEmailIntegrationResponse(data []byte) (DeleteEmailIntegrationResponse, error) {
	var r DeleteEmailIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteEmailIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchEmailIntegrationParams(data []byte) (FetchEmailIntegrationParams, error) {
	var r FetchEmailIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchEmailIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchEmailIntegrationResponse(data []byte) (FetchEmailIntegrationResponse, error) {
	var r FetchEmailIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchEmailIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupEmailIntegrationParams(data []byte) (SetupEmailIntegrationParams, error) {
	var r SetupEmailIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupEmailIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupEmailIntegrationRequest map[string]interface{}

func UnmarshalSetupEmailIntegrationRequest(data []byte) (SetupEmailIntegrationRequest, error) {
	var r SetupEmailIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupEmailIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupEmailIntegrationResponse(data []byte) (SetupEmailIntegrationResponse, error) {
	var r SetupEmailIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupEmailIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateEmailIntegrationParams(data []byte) (UpdateEmailIntegrationParams, error) {
	var r UpdateEmailIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateEmailIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateEmailIntegrationRequest(data []byte) (UpdateEmailIntegrationRequest, error) {
	var r UpdateEmailIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateEmailIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateEmailIntegrationResponse(data []byte) (UpdateEmailIntegrationResponse, error) {
	var r UpdateEmailIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateEmailIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateEmailIntegrationRequest(data []byte) (CreateEmailIntegrationRequest, error) {
	var r CreateEmailIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateEmailIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateEmailIntegrationResponse(data []byte) (CreateEmailIntegrationResponse, error) {
	var r CreateEmailIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateEmailIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListEmailIntegrationsParams(data []byte) (ListEmailIntegrationsParams, error) {
	var r ListEmailIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListEmailIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListEmailIntegrationsResponse(data []byte) (ListEmailIntegrationsResponse, error) {
	var r ListEmailIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListEmailIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteExtractIntegrationParams(data []byte) (DeleteExtractIntegrationParams, error) {
	var r DeleteExtractIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteExtractIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteExtractIntegrationRequest map[string]interface{}

func UnmarshalDeleteExtractIntegrationRequest(data []byte) (DeleteExtractIntegrationRequest, error) {
	var r DeleteExtractIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteExtractIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteExtractIntegrationResponse(data []byte) (DeleteExtractIntegrationResponse, error) {
	var r DeleteExtractIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteExtractIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchExtractIntegrationParams(data []byte) (FetchExtractIntegrationParams, error) {
	var r FetchExtractIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchExtractIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchExtractIntegrationResponse(data []byte) (FetchExtractIntegrationResponse, error) {
	var r FetchExtractIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchExtractIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateExtractIntegrationParams(data []byte) (UpdateExtractIntegrationParams, error) {
	var r UpdateExtractIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateExtractIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateExtractIntegrationRequest(data []byte) (UpdateExtractIntegrationRequest, error) {
	var r UpdateExtractIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateExtractIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateExtractIntegrationResponse(data []byte) (UpdateExtractIntegrationResponse, error) {
	var r UpdateExtractIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateExtractIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateExtractIntegrationRequest(data []byte) (CreateExtractIntegrationRequest, error) {
	var r CreateExtractIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateExtractIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateExtractIntegrationResponse(data []byte) (CreateExtractIntegrationResponse, error) {
	var r CreateExtractIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateExtractIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListExtractIntegrationsParams(data []byte) (ListExtractIntegrationsParams, error) {
	var r ListExtractIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListExtractIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListExtractIntegrationsResponse(data []byte) (ListExtractIntegrationsResponse, error) {
	var r ListExtractIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListExtractIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteInstagramIntegrationParams(data []byte) (DeleteInstagramIntegrationParams, error) {
	var r DeleteInstagramIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteInstagramIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteInstagramIntegrationRequest map[string]interface{}

func UnmarshalDeleteInstagramIntegrationRequest(data []byte) (DeleteInstagramIntegrationRequest, error) {
	var r DeleteInstagramIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteInstagramIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteInstagramIntegrationResponse(data []byte) (DeleteInstagramIntegrationResponse, error) {
	var r DeleteInstagramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteInstagramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchInstagramIntegrationParams(data []byte) (FetchInstagramIntegrationParams, error) {
	var r FetchInstagramIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchInstagramIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchInstagramIntegrationResponse(data []byte) (FetchInstagramIntegrationResponse, error) {
	var r FetchInstagramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchInstagramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupInstagramIntegrationParams(data []byte) (SetupInstagramIntegrationParams, error) {
	var r SetupInstagramIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupInstagramIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupInstagramIntegrationRequest map[string]interface{}

func UnmarshalSetupInstagramIntegrationRequest(data []byte) (SetupInstagramIntegrationRequest, error) {
	var r SetupInstagramIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupInstagramIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupInstagramIntegrationResponse(data []byte) (SetupInstagramIntegrationResponse, error) {
	var r SetupInstagramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupInstagramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateInstagramIntegrationParams(data []byte) (UpdateInstagramIntegrationParams, error) {
	var r UpdateInstagramIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateInstagramIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateInstagramIntegrationRequest(data []byte) (UpdateInstagramIntegrationRequest, error) {
	var r UpdateInstagramIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateInstagramIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateInstagramIntegrationResponse(data []byte) (UpdateInstagramIntegrationResponse, error) {
	var r UpdateInstagramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateInstagramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateInstagramIntegrationRequest(data []byte) (CreateInstagramIntegrationRequest, error) {
	var r CreateInstagramIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateInstagramIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateInstagramIntegrationResponse(data []byte) (CreateInstagramIntegrationResponse, error) {
	var r CreateInstagramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateInstagramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListInstagramIntegrationsParams(data []byte) (ListInstagramIntegrationsParams, error) {
	var r ListInstagramIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListInstagramIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListInstagramIntegrationsResponse(data []byte) (ListInstagramIntegrationsResponse, error) {
	var r ListInstagramIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListInstagramIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteMCPServerIntegrationParams(data []byte) (DeleteMCPServerIntegrationParams, error) {
	var r DeleteMCPServerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteMCPServerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteMCPServerIntegrationRequest map[string]interface{}

func UnmarshalDeleteMCPServerIntegrationRequest(data []byte) (DeleteMCPServerIntegrationRequest, error) {
	var r DeleteMCPServerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteMCPServerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteMCPServerIntegrationResponse(data []byte) (DeleteMCPServerIntegrationResponse, error) {
	var r DeleteMCPServerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteMCPServerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchMCPServerIntegrationParams(data []byte) (FetchMCPServerIntegrationParams, error) {
	var r FetchMCPServerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchMCPServerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchMCPServerIntegrationResponse(data []byte) (FetchMCPServerIntegrationResponse, error) {
	var r FetchMCPServerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchMCPServerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateMCPServerIntegrationParams(data []byte) (UpdateMCPServerIntegrationParams, error) {
	var r UpdateMCPServerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateMCPServerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateMCPServerIntegrationRequest(data []byte) (UpdateMCPServerIntegrationRequest, error) {
	var r UpdateMCPServerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateMCPServerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateMCPServerIntegrationResponse(data []byte) (UpdateMCPServerIntegrationResponse, error) {
	var r UpdateMCPServerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateMCPServerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateMCPServerIntegrationRequest(data []byte) (CreateMCPServerIntegrationRequest, error) {
	var r CreateMCPServerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateMCPServerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateMCPServerIntegrationResponse(data []byte) (CreateMCPServerIntegrationResponse, error) {
	var r CreateMCPServerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateMCPServerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListMCPServerIntegrationsParams(data []byte) (ListMCPServerIntegrationsParams, error) {
	var r ListMCPServerIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListMCPServerIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListMCPServerIntegrationsResponse(data []byte) (ListMCPServerIntegrationsResponse, error) {
	var r ListMCPServerIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListMCPServerIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteMessengerIntegrationParams(data []byte) (DeleteMessengerIntegrationParams, error) {
	var r DeleteMessengerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteMessengerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteMessengerIntegrationRequest map[string]interface{}

func UnmarshalDeleteMessengerIntegrationRequest(data []byte) (DeleteMessengerIntegrationRequest, error) {
	var r DeleteMessengerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteMessengerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteMessengerIntegrationResponse(data []byte) (DeleteMessengerIntegrationResponse, error) {
	var r DeleteMessengerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteMessengerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchMessengerIntegrationParams(data []byte) (FetchMessengerIntegrationParams, error) {
	var r FetchMessengerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchMessengerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchMessengerIntegrationResponse(data []byte) (FetchMessengerIntegrationResponse, error) {
	var r FetchMessengerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchMessengerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupMessengerIntegrationParams(data []byte) (SetupMessengerIntegrationParams, error) {
	var r SetupMessengerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupMessengerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupMessengerIntegrationRequest map[string]interface{}

func UnmarshalSetupMessengerIntegrationRequest(data []byte) (SetupMessengerIntegrationRequest, error) {
	var r SetupMessengerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupMessengerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupMessengerIntegrationResponse(data []byte) (SetupMessengerIntegrationResponse, error) {
	var r SetupMessengerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupMessengerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateMessengerIntegrationParams(data []byte) (UpdateMessengerIntegrationParams, error) {
	var r UpdateMessengerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateMessengerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateMessengerIntegrationRequest(data []byte) (UpdateMessengerIntegrationRequest, error) {
	var r UpdateMessengerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateMessengerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateMessengerIntegrationResponse(data []byte) (UpdateMessengerIntegrationResponse, error) {
	var r UpdateMessengerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateMessengerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateMessengerIntegrationRequest(data []byte) (CreateMessengerIntegrationRequest, error) {
	var r CreateMessengerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateMessengerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateMessengerIntegrationResponse(data []byte) (CreateMessengerIntegrationResponse, error) {
	var r CreateMessengerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateMessengerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListMessengerIntegrationsParams(data []byte) (ListMessengerIntegrationsParams, error) {
	var r ListMessengerIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListMessengerIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListMessengerIntegrationsResponse(data []byte) (ListMessengerIntegrationsResponse, error) {
	var r ListMessengerIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListMessengerIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteNotionIntegrationParams(data []byte) (DeleteNotionIntegrationParams, error) {
	var r DeleteNotionIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteNotionIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteNotionIntegrationRequest map[string]interface{}

func UnmarshalDeleteNotionIntegrationRequest(data []byte) (DeleteNotionIntegrationRequest, error) {
	var r DeleteNotionIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteNotionIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteNotionIntegrationResponse(data []byte) (DeleteNotionIntegrationResponse, error) {
	var r DeleteNotionIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteNotionIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchNotionIntegrationParams(data []byte) (FetchNotionIntegrationParams, error) {
	var r FetchNotionIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchNotionIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchNotionIntegrationResponse(data []byte) (FetchNotionIntegrationResponse, error) {
	var r FetchNotionIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchNotionIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSyncNotionIntegrationParams(data []byte) (SyncNotionIntegrationParams, error) {
	var r SyncNotionIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncNotionIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SyncNotionIntegrationRequest map[string]interface{}

func UnmarshalSyncNotionIntegrationRequest(data []byte) (SyncNotionIntegrationRequest, error) {
	var r SyncNotionIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncNotionIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSyncNotionIntegrationResponse(data []byte) (SyncNotionIntegrationResponse, error) {
	var r SyncNotionIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncNotionIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateNotionIntegrationParams(data []byte) (UpdateNotionIntegrationParams, error) {
	var r UpdateNotionIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateNotionIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateNotionIntegrationRequest(data []byte) (UpdateNotionIntegrationRequest, error) {
	var r UpdateNotionIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateNotionIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateNotionIntegrationResponse(data []byte) (UpdateNotionIntegrationResponse, error) {
	var r UpdateNotionIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateNotionIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateNotionIntegrationRequest(data []byte) (CreateNotionIntegrationRequest, error) {
	var r CreateNotionIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateNotionIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateNotionIntegrationResponse(data []byte) (CreateNotionIntegrationResponse, error) {
	var r CreateNotionIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateNotionIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListNotionIntegrationsParams(data []byte) (ListNotionIntegrationsParams, error) {
	var r ListNotionIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListNotionIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListNotionIntegrationsResponse(data []byte) (ListNotionIntegrationsResponse, error) {
	var r ListNotionIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListNotionIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSitemapIntegrationParams(data []byte) (DeleteSitemapIntegrationParams, error) {
	var r DeleteSitemapIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSitemapIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteSitemapIntegrationRequest map[string]interface{}

func UnmarshalDeleteSitemapIntegrationRequest(data []byte) (DeleteSitemapIntegrationRequest, error) {
	var r DeleteSitemapIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSitemapIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSitemapIntegrationResponse(data []byte) (DeleteSitemapIntegrationResponse, error) {
	var r DeleteSitemapIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSitemapIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSitemapIntegrationParams(data []byte) (FetchSitemapIntegrationParams, error) {
	var r FetchSitemapIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSitemapIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSitemapIntegrationResponse(data []byte) (FetchSitemapIntegrationResponse, error) {
	var r FetchSitemapIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSitemapIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSyncSitemapIntegrationParams(data []byte) (SyncSitemapIntegrationParams, error) {
	var r SyncSitemapIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncSitemapIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SyncSitemapIntegrationRequest map[string]interface{}

func UnmarshalSyncSitemapIntegrationRequest(data []byte) (SyncSitemapIntegrationRequest, error) {
	var r SyncSitemapIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncSitemapIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSyncSitemapIntegrationResponse(data []byte) (SyncSitemapIntegrationResponse, error) {
	var r SyncSitemapIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SyncSitemapIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSitemapIntegrationParams(data []byte) (UpdateSitemapIntegrationParams, error) {
	var r UpdateSitemapIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSitemapIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSitemapIntegrationRequest(data []byte) (UpdateSitemapIntegrationRequest, error) {
	var r UpdateSitemapIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSitemapIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSitemapIntegrationResponse(data []byte) (UpdateSitemapIntegrationResponse, error) {
	var r UpdateSitemapIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSitemapIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSitemapIntegrationRequest(data []byte) (CreateSitemapIntegrationRequest, error) {
	var r CreateSitemapIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSitemapIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSitemapIntegrationResponse(data []byte) (CreateSitemapIntegrationResponse, error) {
	var r CreateSitemapIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSitemapIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSitemapIntegrationsParams(data []byte) (ListSitemapIntegrationsParams, error) {
	var r ListSitemapIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSitemapIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSitemapIntegrationsResponse(data []byte) (ListSitemapIntegrationsResponse, error) {
	var r ListSitemapIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSitemapIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSlackIntegrationParams(data []byte) (DeleteSlackIntegrationParams, error) {
	var r DeleteSlackIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSlackIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteSlackIntegrationRequest map[string]interface{}

func UnmarshalDeleteSlackIntegrationRequest(data []byte) (DeleteSlackIntegrationRequest, error) {
	var r DeleteSlackIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSlackIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSlackIntegrationResponse(data []byte) (DeleteSlackIntegrationResponse, error) {
	var r DeleteSlackIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSlackIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSlackIntegrationParams(data []byte) (FetchSlackIntegrationParams, error) {
	var r FetchSlackIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSlackIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSlackIntegrationResponse(data []byte) (FetchSlackIntegrationResponse, error) {
	var r FetchSlackIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSlackIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupSlackIntegrationParams(data []byte) (SetupSlackIntegrationParams, error) {
	var r SetupSlackIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupSlackIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupSlackIntegrationRequest map[string]interface{}

func UnmarshalSetupSlackIntegrationRequest(data []byte) (SetupSlackIntegrationRequest, error) {
	var r SetupSlackIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupSlackIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupSlackIntegrationResponse(data []byte) (SetupSlackIntegrationResponse, error) {
	var r SetupSlackIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupSlackIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSlackIntegrationParams(data []byte) (UpdateSlackIntegrationParams, error) {
	var r UpdateSlackIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSlackIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSlackIntegrationRequest(data []byte) (UpdateSlackIntegrationRequest, error) {
	var r UpdateSlackIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSlackIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSlackIntegrationResponse(data []byte) (UpdateSlackIntegrationResponse, error) {
	var r UpdateSlackIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSlackIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSlackIntegrationRequest(data []byte) (CreateSlackIntegrationRequest, error) {
	var r CreateSlackIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSlackIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSlackIntegrationResponse(data []byte) (CreateSlackIntegrationResponse, error) {
	var r CreateSlackIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSlackIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSlackIntegrationsParams(data []byte) (ListSlackIntegrationsParams, error) {
	var r ListSlackIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSlackIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSlackIntegrationsResponse(data []byte) (ListSlackIntegrationsResponse, error) {
	var r ListSlackIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSlackIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSupportIntegrationParams(data []byte) (DeleteSupportIntegrationParams, error) {
	var r DeleteSupportIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSupportIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteSupportIntegrationRequest map[string]interface{}

func UnmarshalDeleteSupportIntegrationRequest(data []byte) (DeleteSupportIntegrationRequest, error) {
	var r DeleteSupportIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSupportIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSupportIntegrationResponse(data []byte) (DeleteSupportIntegrationResponse, error) {
	var r DeleteSupportIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSupportIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSupportIntegrationParams(data []byte) (FetchSupportIntegrationParams, error) {
	var r FetchSupportIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSupportIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSupportIntegrationResponse(data []byte) (FetchSupportIntegrationResponse, error) {
	var r FetchSupportIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSupportIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSupportIntegrationParams(data []byte) (UpdateSupportIntegrationParams, error) {
	var r UpdateSupportIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSupportIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSupportIntegrationRequest(data []byte) (UpdateSupportIntegrationRequest, error) {
	var r UpdateSupportIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSupportIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSupportIntegrationResponse(data []byte) (UpdateSupportIntegrationResponse, error) {
	var r UpdateSupportIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSupportIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSupportIntegrationRequest(data []byte) (CreateSupportIntegrationRequest, error) {
	var r CreateSupportIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSupportIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSupportIntegrationResponse(data []byte) (CreateSupportIntegrationResponse, error) {
	var r CreateSupportIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSupportIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSupportIntegrationsParams(data []byte) (ListSupportIntegrationsParams, error) {
	var r ListSupportIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSupportIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSupportIntegrationsResponse(data []byte) (ListSupportIntegrationsResponse, error) {
	var r ListSupportIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSupportIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteTelegramIntegrationParams(data []byte) (DeleteTelegramIntegrationParams, error) {
	var r DeleteTelegramIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTelegramIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteTelegramIntegrationRequest map[string]interface{}

func UnmarshalDeleteTelegramIntegrationRequest(data []byte) (DeleteTelegramIntegrationRequest, error) {
	var r DeleteTelegramIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTelegramIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteTelegramIntegrationResponse(data []byte) (DeleteTelegramIntegrationResponse, error) {
	var r DeleteTelegramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTelegramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchTelegramIntegrationParams(data []byte) (FetchTelegramIntegrationParams, error) {
	var r FetchTelegramIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchTelegramIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchTelegramIntegrationResponse(data []byte) (FetchTelegramIntegrationResponse, error) {
	var r FetchTelegramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchTelegramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupTelegramIntegrationParams(data []byte) (SetupTelegramIntegrationParams, error) {
	var r SetupTelegramIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupTelegramIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupTelegramIntegrationRequest map[string]interface{}

func UnmarshalSetupTelegramIntegrationRequest(data []byte) (SetupTelegramIntegrationRequest, error) {
	var r SetupTelegramIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupTelegramIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupTelegramIntegrationResponse(data []byte) (SetupTelegramIntegrationResponse, error) {
	var r SetupTelegramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupTelegramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTelegramIntegrationParams(data []byte) (UpdateTelegramIntegrationParams, error) {
	var r UpdateTelegramIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTelegramIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTelegramIntegrationRequest(data []byte) (UpdateTelegramIntegrationRequest, error) {
	var r UpdateTelegramIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTelegramIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTelegramIntegrationResponse(data []byte) (UpdateTelegramIntegrationResponse, error) {
	var r UpdateTelegramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTelegramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateTelegramIntegrationRequest(data []byte) (CreateTelegramIntegrationRequest, error) {
	var r CreateTelegramIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateTelegramIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateTelegramIntegrationResponse(data []byte) (CreateTelegramIntegrationResponse, error) {
	var r CreateTelegramIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateTelegramIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTelegramIntegrationsParams(data []byte) (ListTelegramIntegrationsParams, error) {
	var r ListTelegramIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTelegramIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTelegramIntegrationsResponse(data []byte) (ListTelegramIntegrationsResponse, error) {
	var r ListTelegramIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTelegramIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteTriggerIntegrationParams(data []byte) (DeleteTriggerIntegrationParams, error) {
	var r DeleteTriggerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTriggerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteTriggerIntegrationRequest map[string]interface{}

func UnmarshalDeleteTriggerIntegrationRequest(data []byte) (DeleteTriggerIntegrationRequest, error) {
	var r DeleteTriggerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTriggerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteTriggerIntegrationResponse(data []byte) (DeleteTriggerIntegrationResponse, error) {
	var r DeleteTriggerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTriggerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchTriggerIntegrationParams(data []byte) (FetchTriggerIntegrationParams, error) {
	var r FetchTriggerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchTriggerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchTriggerIntegrationResponse(data []byte) (FetchTriggerIntegrationResponse, error) {
	var r FetchTriggerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchTriggerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInvokeTriggerIntegrationParams(data []byte) (InvokeTriggerIntegrationParams, error) {
	var r InvokeTriggerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InvokeTriggerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InvokeTriggerIntegrationRequest map[string]interface{}

func UnmarshalInvokeTriggerIntegrationRequest(data []byte) (InvokeTriggerIntegrationRequest, error) {
	var r InvokeTriggerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InvokeTriggerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInvokeTriggerIntegrationResponse(data []byte) (InvokeTriggerIntegrationResponse, error) {
	var r InvokeTriggerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InvokeTriggerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupTriggerIntegrationParams(data []byte) (SetupTriggerIntegrationParams, error) {
	var r SetupTriggerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupTriggerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupTriggerIntegrationRequest map[string]interface{}

func UnmarshalSetupTriggerIntegrationRequest(data []byte) (SetupTriggerIntegrationRequest, error) {
	var r SetupTriggerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupTriggerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupTriggerIntegrationResponse(data []byte) (SetupTriggerIntegrationResponse, error) {
	var r SetupTriggerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupTriggerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTriggerIntegrationParams(data []byte) (UpdateTriggerIntegrationParams, error) {
	var r UpdateTriggerIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTriggerIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTriggerIntegrationRequest(data []byte) (UpdateTriggerIntegrationRequest, error) {
	var r UpdateTriggerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTriggerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTriggerIntegrationResponse(data []byte) (UpdateTriggerIntegrationResponse, error) {
	var r UpdateTriggerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTriggerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateTriggerIntegrationRequest(data []byte) (CreateTriggerIntegrationRequest, error) {
	var r CreateTriggerIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateTriggerIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateTriggerIntegrationResponse(data []byte) (CreateTriggerIntegrationResponse, error) {
	var r CreateTriggerIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateTriggerIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTriggerIntegrationsParams(data []byte) (ListTriggerIntegrationsParams, error) {
	var r ListTriggerIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTriggerIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTriggerIntegrationsResponse(data []byte) (ListTriggerIntegrationsResponse, error) {
	var r ListTriggerIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTriggerIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteTwilioIntegrationParams(data []byte) (DeleteTwilioIntegrationParams, error) {
	var r DeleteTwilioIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTwilioIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteTwilioIntegrationRequest map[string]interface{}

func UnmarshalDeleteTwilioIntegrationRequest(data []byte) (DeleteTwilioIntegrationRequest, error) {
	var r DeleteTwilioIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTwilioIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteTwilioIntegrationResponse(data []byte) (DeleteTwilioIntegrationResponse, error) {
	var r DeleteTwilioIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTwilioIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchTwilioIntegrationParams(data []byte) (FetchTwilioIntegrationParams, error) {
	var r FetchTwilioIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchTwilioIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchTwilioIntegrationResponse(data []byte) (FetchTwilioIntegrationResponse, error) {
	var r FetchTwilioIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchTwilioIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupTwilioIntegrationParams(data []byte) (SetupTwilioIntegrationParams, error) {
	var r SetupTwilioIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupTwilioIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupTwilioIntegrationRequest map[string]interface{}

func UnmarshalSetupTwilioIntegrationRequest(data []byte) (SetupTwilioIntegrationRequest, error) {
	var r SetupTwilioIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupTwilioIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupTwilioIntegrationResponse(data []byte) (SetupTwilioIntegrationResponse, error) {
	var r SetupTwilioIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupTwilioIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTwilioIntegrationParams(data []byte) (UpdateTwilioIntegrationParams, error) {
	var r UpdateTwilioIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTwilioIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTwilioIntegrationRequest(data []byte) (UpdateTwilioIntegrationRequest, error) {
	var r UpdateTwilioIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTwilioIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTwilioIntegrationResponse(data []byte) (UpdateTwilioIntegrationResponse, error) {
	var r UpdateTwilioIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTwilioIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateTwilioIntegrationRequest(data []byte) (CreateTwilioIntegrationRequest, error) {
	var r CreateTwilioIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateTwilioIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateTwilioIntegrationResponse(data []byte) (CreateTwilioIntegrationResponse, error) {
	var r CreateTwilioIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateTwilioIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTwilioIntegrationsParams(data []byte) (ListTwilioIntegrationsParams, error) {
	var r ListTwilioIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTwilioIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTwilioIntegrationsResponse(data []byte) (ListTwilioIntegrationsResponse, error) {
	var r ListTwilioIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTwilioIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteWhatsAppIntegrationParams(data []byte) (DeleteWhatsAppIntegrationParams, error) {
	var r DeleteWhatsAppIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteWhatsAppIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteWhatsAppIntegrationRequest map[string]interface{}

func UnmarshalDeleteWhatsAppIntegrationRequest(data []byte) (DeleteWhatsAppIntegrationRequest, error) {
	var r DeleteWhatsAppIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteWhatsAppIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteWhatsAppIntegrationResponse(data []byte) (DeleteWhatsAppIntegrationResponse, error) {
	var r DeleteWhatsAppIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteWhatsAppIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchWhatsAppIntegrationParams(data []byte) (FetchWhatsAppIntegrationParams, error) {
	var r FetchWhatsAppIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchWhatsAppIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchWhatsAppIntegrationResponse(data []byte) (FetchWhatsAppIntegrationResponse, error) {
	var r FetchWhatsAppIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchWhatsAppIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupWhatsAppIntegrationParams(data []byte) (SetupWhatsAppIntegrationParams, error) {
	var r SetupWhatsAppIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupWhatsAppIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupWhatsAppIntegrationRequest map[string]interface{}

func UnmarshalSetupWhatsAppIntegrationRequest(data []byte) (SetupWhatsAppIntegrationRequest, error) {
	var r SetupWhatsAppIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupWhatsAppIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupWhatsAppIntegrationResponse(data []byte) (SetupWhatsAppIntegrationResponse, error) {
	var r SetupWhatsAppIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupWhatsAppIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateWhatsAppIntegrationParams(data []byte) (UpdateWhatsAppIntegrationParams, error) {
	var r UpdateWhatsAppIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateWhatsAppIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateWhatsAppIntegrationRequest(data []byte) (UpdateWhatsAppIntegrationRequest, error) {
	var r UpdateWhatsAppIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateWhatsAppIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateWhatsAppIntegrationResponse(data []byte) (UpdateWhatsAppIntegrationResponse, error) {
	var r UpdateWhatsAppIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateWhatsAppIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateWhatsAppIntegrationRequest(data []byte) (CreateWhatsAppIntegrationRequest, error) {
	var r CreateWhatsAppIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateWhatsAppIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateWhatsAppIntegrationResponse(data []byte) (CreateWhatsAppIntegrationResponse, error) {
	var r CreateWhatsAppIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateWhatsAppIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListWhatsAppIntegrationsParams(data []byte) (ListWhatsAppIntegrationsParams, error) {
	var r ListWhatsAppIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListWhatsAppIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListWhatsAppIntegrationsResponse(data []byte) (ListWhatsAppIntegrationsResponse, error) {
	var r ListWhatsAppIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListWhatsAppIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteWidgetIntegrationParams(data []byte) (DeleteWidgetIntegrationParams, error) {
	var r DeleteWidgetIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteWidgetIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteWidgetIntegrationRequest map[string]interface{}

func UnmarshalDeleteWidgetIntegrationRequest(data []byte) (DeleteWidgetIntegrationRequest, error) {
	var r DeleteWidgetIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteWidgetIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteWidgetIntegrationResponse(data []byte) (DeleteWidgetIntegrationResponse, error) {
	var r DeleteWidgetIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteWidgetIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchWidgetIntegrationParams(data []byte) (FetchWidgetIntegrationParams, error) {
	var r FetchWidgetIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchWidgetIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchWidgetIntegrationResponse(data []byte) (FetchWidgetIntegrationResponse, error) {
	var r FetchWidgetIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchWidgetIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupWidgetIntegrationParams(data []byte) (SetupWidgetIntegrationParams, error) {
	var r SetupWidgetIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupWidgetIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SetupWidgetIntegrationRequest map[string]interface{}

func UnmarshalSetupWidgetIntegrationRequest(data []byte) (SetupWidgetIntegrationRequest, error) {
	var r SetupWidgetIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupWidgetIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSetupWidgetIntegrationResponse(data []byte) (SetupWidgetIntegrationResponse, error) {
	var r SetupWidgetIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SetupWidgetIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateWidgetIntegrationParams(data []byte) (UpdateWidgetIntegrationParams, error) {
	var r UpdateWidgetIntegrationParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateWidgetIntegrationParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateWidgetIntegrationRequest(data []byte) (UpdateWidgetIntegrationRequest, error) {
	var r UpdateWidgetIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateWidgetIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateWidgetIntegrationResponse(data []byte) (UpdateWidgetIntegrationResponse, error) {
	var r UpdateWidgetIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateWidgetIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateWidgetIntegrationRequest(data []byte) (CreateWidgetIntegrationRequest, error) {
	var r CreateWidgetIntegrationRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateWidgetIntegrationRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateWidgetIntegrationResponse(data []byte) (CreateWidgetIntegrationResponse, error) {
	var r CreateWidgetIntegrationResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateWidgetIntegrationResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListWidgetIntegrationsParams(data []byte) (ListWidgetIntegrationsParams, error) {
	var r ListWidgetIntegrationsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListWidgetIntegrationsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListWidgetIntegrationsResponse(data []byte) (ListWidgetIntegrationsResponse, error) {
	var r ListWidgetIntegrationsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListWidgetIntegrationsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalGenerateMagicFromPromptParams(data []byte) (GenerateMagicFromPromptParams, error) {
	var r GenerateMagicFromPromptParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GenerateMagicFromPromptParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalGenerateMagicFromPromptRequest(data []byte) (GenerateMagicFromPromptRequest, error) {
	var r GenerateMagicFromPromptRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GenerateMagicFromPromptRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalGenerateMagicFromPromptResponse(data []byte) (GenerateMagicFromPromptResponse, error) {
	var r GenerateMagicFromPromptResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GenerateMagicFromPromptResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListMagicPromptsParams(data []byte) (ListMagicPromptsParams, error) {
	var r ListMagicPromptsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListMagicPromptsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListMagicPromptsResponse(data []byte) (ListMagicPromptsResponse, error) {
	var r ListMagicPromptsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListMagicPromptsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteMemoryParams(data []byte) (DeleteMemoryParams, error) {
	var r DeleteMemoryParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteMemoryParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteMemoryRequest map[string]interface{}

func UnmarshalDeleteMemoryRequest(data []byte) (DeleteMemoryRequest, error) {
	var r DeleteMemoryRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteMemoryRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteMemoryResponse(data []byte) (DeleteMemoryResponse, error) {
	var r DeleteMemoryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteMemoryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchMemoryParams(data []byte) (FetchMemoryParams, error) {
	var r FetchMemoryParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchMemoryParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchMemoryResponse(data []byte) (FetchMemoryResponse, error) {
	var r FetchMemoryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchMemoryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateMemoryParams(data []byte) (UpdateMemoryParams, error) {
	var r UpdateMemoryParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateMemoryParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateMemoryRequest(data []byte) (UpdateMemoryRequest, error) {
	var r UpdateMemoryRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateMemoryRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateMemoryResponse(data []byte) (UpdateMemoryResponse, error) {
	var r UpdateMemoryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateMemoryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateMemoryRequest(data []byte) (CreateMemoryRequest, error) {
	var r CreateMemoryRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateMemoryRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateMemoryResponse(data []byte) (CreateMemoryResponse, error) {
	var r CreateMemoryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateMemoryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportMemoriesParams(data []byte) (ExportMemoriesParams, error) {
	var r ExportMemoriesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportMemoriesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportMemoriesResponse(data []byte) (ExportMemoriesResponse, error) {
	var r ExportMemoriesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportMemoriesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListMemoriesParams(data []byte) (ListMemoriesParams, error) {
	var r ListMemoriesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListMemoriesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListMemoriesResponse(data []byte) (ListMemoriesResponse, error) {
	var r ListMemoriesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListMemoriesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchMemoryRequest(data []byte) (SearchMemoryRequest, error) {
	var r SearchMemoryRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchMemoryRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchMemoryResponse(data []byte) (SearchMemoryResponse, error) {
	var r SearchMemoryResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchMemoryResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeletePartnerUserParams(data []byte) (DeletePartnerUserParams, error) {
	var r DeletePartnerUserParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePartnerUserParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeletePartnerUserRequest map[string]interface{}

func UnmarshalDeletePartnerUserRequest(data []byte) (DeletePartnerUserRequest, error) {
	var r DeletePartnerUserRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePartnerUserRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeletePartnerUserResponse(data []byte) (DeletePartnerUserResponse, error) {
	var r DeletePartnerUserResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePartnerUserResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPartnerUserParams(data []byte) (FetchPartnerUserParams, error) {
	var r FetchPartnerUserParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPartnerUserParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPartnerUserResponse(data []byte) (FetchPartnerUserResponse, error) {
	var r FetchPartnerUserResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPartnerUserResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeletePartnerUserTokenParams(data []byte) (DeletePartnerUserTokenParams, error) {
	var r DeletePartnerUserTokenParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePartnerUserTokenParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeletePartnerUserTokenRequest map[string]interface{}

func UnmarshalDeletePartnerUserTokenRequest(data []byte) (DeletePartnerUserTokenRequest, error) {
	var r DeletePartnerUserTokenRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePartnerUserTokenRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeletePartnerUserTokenResponse(data []byte) (DeletePartnerUserTokenResponse, error) {
	var r DeletePartnerUserTokenResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePartnerUserTokenResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreatePartnerUserTokenParams(data []byte) (CreatePartnerUserTokenParams, error) {
	var r CreatePartnerUserTokenParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePartnerUserTokenParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreatePartnerUserTokenRequest map[string]interface{}

func UnmarshalCreatePartnerUserTokenRequest(data []byte) (CreatePartnerUserTokenRequest, error) {
	var r CreatePartnerUserTokenRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePartnerUserTokenRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreatePartnerUserTokenResponse(data []byte) (CreatePartnerUserTokenResponse, error) {
	var r CreatePartnerUserTokenResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePartnerUserTokenResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPartnerUserTokensParams(data []byte) (ListPartnerUserTokensParams, error) {
	var r ListPartnerUserTokensParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPartnerUserTokensParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPartnerUserTokensResponse(data []byte) (ListPartnerUserTokensResponse, error) {
	var r ListPartnerUserTokensResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPartnerUserTokensResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdatePartnerUserParams(data []byte) (UpdatePartnerUserParams, error) {
	var r UpdatePartnerUserParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdatePartnerUserParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdatePartnerUserRequest(data []byte) (UpdatePartnerUserRequest, error) {
	var r UpdatePartnerUserRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdatePartnerUserRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdatePartnerUserResponse(data []byte) (UpdatePartnerUserResponse, error) {
	var r UpdatePartnerUserResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdatePartnerUserResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreatePartnerUserRequest(data []byte) (CreatePartnerUserRequest, error) {
	var r CreatePartnerUserRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePartnerUserRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreatePartnerUserResponse(data []byte) (CreatePartnerUserResponse, error) {
	var r CreatePartnerUserResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePartnerUserResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPartnerUsersParams(data []byte) (ListPartnerUsersParams, error) {
	var r ListPartnerUsersParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPartnerUsersParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPartnerUsersResponse(data []byte) (ListPartnerUsersResponse, error) {
	var r ListPartnerUsersResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPartnerUsersResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformAbilitiesParams(data []byte) (ListPlatformAbilitiesParams, error) {
	var r ListPlatformAbilitiesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformAbilitiesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformAbilitiesResponse(data []byte) (ListPlatformAbilitiesResponse, error) {
	var r ListPlatformAbilitiesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformAbilitiesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformActionsParams(data []byte) (ListPlatformActionsParams, error) {
	var r ListPlatformActionsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformActionsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformActionsResponse(data []byte) (ListPlatformActionsResponse, error) {
	var r ListPlatformActionsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformActionsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformDocParams(data []byte) (FetchPlatformDocParams, error) {
	var r FetchPlatformDocParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformDocParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformDocResponse(data []byte) (FetchPlatformDocResponse, error) {
	var r FetchPlatformDocResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformDocResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformDocsParams(data []byte) (ListPlatformDocsParams, error) {
	var r ListPlatformDocsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformDocsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformDocsResponse(data []byte) (ListPlatformDocsResponse, error) {
	var r ListPlatformDocsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformDocsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformDocsRequest(data []byte) (SearchPlatformDocsRequest, error) {
	var r SearchPlatformDocsRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformDocsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformDocsResponse(data []byte) (SearchPlatformDocsResponse, error) {
	var r SearchPlatformDocsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformDocsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalClonePlatformExampleParams(data []byte) (ClonePlatformExampleParams, error) {
	var r ClonePlatformExampleParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ClonePlatformExampleParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ClonePlatformExampleRequest map[string]interface{}

func UnmarshalClonePlatformExampleRequest(data []byte) (ClonePlatformExampleRequest, error) {
	var r ClonePlatformExampleRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ClonePlatformExampleRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalClonePlatformExampleResponse(data []byte) (ClonePlatformExampleResponse, error) {
	var r ClonePlatformExampleResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ClonePlatformExampleResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformExampleParams(data []byte) (FetchPlatformExampleParams, error) {
	var r FetchPlatformExampleParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformExampleParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformExampleResponse(data []byte) (FetchPlatformExampleResponse, error) {
	var r FetchPlatformExampleResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformExampleResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformExamplesParams(data []byte) (ListPlatformExamplesParams, error) {
	var r ListPlatformExamplesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformExamplesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformExamplesResponse(data []byte) (ListPlatformExamplesResponse, error) {
	var r ListPlatformExamplesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformExamplesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformExamplesRequest(data []byte) (SearchPlatformExamplesRequest, error) {
	var r SearchPlatformExamplesRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformExamplesRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformExamplesResponse(data []byte) (SearchPlatformExamplesResponse, error) {
	var r SearchPlatformExamplesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformExamplesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformGuideParams(data []byte) (FetchPlatformGuideParams, error) {
	var r FetchPlatformGuideParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformGuideParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformGuideResponse(data []byte) (FetchPlatformGuideResponse, error) {
	var r FetchPlatformGuideResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformGuideResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformGuidesParams(data []byte) (ListPlatformGuidesParams, error) {
	var r ListPlatformGuidesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformGuidesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformGuidesResponse(data []byte) (ListPlatformGuidesResponse, error) {
	var r ListPlatformGuidesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformGuidesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformGuidesRequest(data []byte) (SearchPlatformGuidesRequest, error) {
	var r SearchPlatformGuidesRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformGuidesRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformGuidesResponse(data []byte) (SearchPlatformGuidesResponse, error) {
	var r SearchPlatformGuidesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformGuidesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformManualParams(data []byte) (FetchPlatformManualParams, error) {
	var r FetchPlatformManualParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformManualParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformManualResponse(data []byte) (FetchPlatformManualResponse, error) {
	var r FetchPlatformManualResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformManualResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformManualsParams(data []byte) (ListPlatformManualsParams, error) {
	var r ListPlatformManualsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformManualsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformManualsResponse(data []byte) (ListPlatformManualsResponse, error) {
	var r ListPlatformManualsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformManualsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformManualsRequest(data []byte) (SearchPlatformManualsRequest, error) {
	var r SearchPlatformManualsRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformManualsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformManualsResponse(data []byte) (SearchPlatformManualsResponse, error) {
	var r SearchPlatformManualsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformManualsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformModelsParams(data []byte) (ListPlatformModelsParams, error) {
	var r ListPlatformModelsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformModelsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformModelsResponse(data []byte) (ListPlatformModelsResponse, error) {
	var r ListPlatformModelsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformModelsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformSecretsParams(data []byte) (ListPlatformSecretsParams, error) {
	var r ListPlatformSecretsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformSecretsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformSecretsResponse(data []byte) (ListPlatformSecretsResponse, error) {
	var r ListPlatformSecretsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformSecretsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformTutorialParams(data []byte) (FetchPlatformTutorialParams, error) {
	var r FetchPlatformTutorialParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformTutorialParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPlatformTutorialResponse(data []byte) (FetchPlatformTutorialResponse, error) {
	var r FetchPlatformTutorialResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPlatformTutorialResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformTutorialsParams(data []byte) (ListPlatformTutorialsParams, error) {
	var r ListPlatformTutorialsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformTutorialsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPlatformTutorialsResponse(data []byte) (ListPlatformTutorialsResponse, error) {
	var r ListPlatformTutorialsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPlatformTutorialsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformTutorialsRequest(data []byte) (SearchPlatformTutorialsRequest, error) {
	var r SearchPlatformTutorialsRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformTutorialsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchPlatformTutorialsResponse(data []byte) (SearchPlatformTutorialsResponse, error) {
	var r SearchPlatformTutorialsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchPlatformTutorialsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeletePolicyParams(data []byte) (DeletePolicyParams, error) {
	var r DeletePolicyParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePolicyParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeletePolicyRequest map[string]interface{}

func UnmarshalDeletePolicyRequest(data []byte) (DeletePolicyRequest, error) {
	var r DeletePolicyRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePolicyRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeletePolicyResponse(data []byte) (DeletePolicyResponse, error) {
	var r DeletePolicyResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePolicyResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPolicyParams(data []byte) (FetchPolicyParams, error) {
	var r FetchPolicyParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPolicyParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPolicyResponse(data []byte) (FetchPolicyResponse, error) {
	var r FetchPolicyResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPolicyResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdatePolicyParams(data []byte) (UpdatePolicyParams, error) {
	var r UpdatePolicyParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdatePolicyParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdatePolicyRequest(data []byte) (UpdatePolicyRequest, error) {
	var r UpdatePolicyRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdatePolicyRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdatePolicyResponse(data []byte) (UpdatePolicyResponse, error) {
	var r UpdatePolicyResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdatePolicyResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreatePolicyRequest(data []byte) (CreatePolicyRequest, error) {
	var r CreatePolicyRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePolicyRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreatePolicyResponse(data []byte) (CreatePolicyResponse, error) {
	var r CreatePolicyResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePolicyResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPoliciesParams(data []byte) (ListPoliciesParams, error) {
	var r ListPoliciesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPoliciesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPoliciesResponse(data []byte) (ListPoliciesResponse, error) {
	var r ListPoliciesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPoliciesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeletePortalParams(data []byte) (DeletePortalParams, error) {
	var r DeletePortalParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePortalParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeletePortalRequest map[string]interface{}

func UnmarshalDeletePortalRequest(data []byte) (DeletePortalRequest, error) {
	var r DeletePortalRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePortalRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeletePortalResponse(data []byte) (DeletePortalResponse, error) {
	var r DeletePortalResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeletePortalResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPortalParams(data []byte) (FetchPortalParams, error) {
	var r FetchPortalParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPortalParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchPortalResponse(data []byte) (FetchPortalResponse, error) {
	var r FetchPortalResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchPortalResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdatePortalParams(data []byte) (UpdatePortalParams, error) {
	var r UpdatePortalParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdatePortalParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdatePortalRequest(data []byte) (UpdatePortalRequest, error) {
	var r UpdatePortalRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdatePortalRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdatePortalResponse(data []byte) (UpdatePortalResponse, error) {
	var r UpdatePortalResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdatePortalResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreatePortalRequest(data []byte) (CreatePortalRequest, error) {
	var r CreatePortalRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePortalRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreatePortalResponse(data []byte) (CreatePortalResponse, error) {
	var r CreatePortalResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePortalResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPortalsParams(data []byte) (ListPortalsParams, error) {
	var r ListPortalsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPortalsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListPortalsResponse(data []byte) (ListPortalsResponse, error) {
	var r ListPortalsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListPortalsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAuthenticateSecretParams(data []byte) (AuthenticateSecretParams, error) {
	var r AuthenticateSecretParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthenticateSecretParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuthenticateSecretRequest map[string]interface{}

func UnmarshalAuthenticateSecretRequest(data []byte) (AuthenticateSecretRequest, error) {
	var r AuthenticateSecretRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthenticateSecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAuthenticateSecretResponse(data []byte) (AuthenticateSecretResponse, error) {
	var r AuthenticateSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthenticateSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSecretParams(data []byte) (DeleteSecretParams, error) {
	var r DeleteSecretParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSecretParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteSecretRequest map[string]interface{}

func UnmarshalDeleteSecretRequest(data []byte) (DeleteSecretRequest, error) {
	var r DeleteSecretRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSecretResponse(data []byte) (DeleteSecretResponse, error) {
	var r DeleteSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSecretParams(data []byte) (FetchSecretParams, error) {
	var r FetchSecretParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSecretParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSecretResponse(data []byte) (FetchSecretResponse, error) {
	var r FetchSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRevokeSecretParams(data []byte) (RevokeSecretParams, error) {
	var r RevokeSecretParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RevokeSecretParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RevokeSecretRequest map[string]interface{}

func UnmarshalRevokeSecretRequest(data []byte) (RevokeSecretRequest, error) {
	var r RevokeSecretRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RevokeSecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRevokeSecretResponse(data []byte) (RevokeSecretResponse, error) {
	var r RevokeSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RevokeSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSecretParams(data []byte) (UpdateSecretParams, error) {
	var r UpdateSecretParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSecretParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSecretRequest(data []byte) (UpdateSecretRequest, error) {
	var r UpdateSecretRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSecretResponse(data []byte) (UpdateSecretResponse, error) {
	var r UpdateSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalVerifySecretParams(data []byte) (VerifySecretParams, error) {
	var r VerifySecretParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VerifySecretParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type VerifySecretRequest map[string]interface{}

func UnmarshalVerifySecretRequest(data []byte) (VerifySecretRequest, error) {
	var r VerifySecretRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VerifySecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalVerifySecretResponse(data []byte) (VerifySecretResponse, error) {
	var r VerifySecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VerifySecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSecretRequest(data []byte) (CreateSecretRequest, error) {
	var r CreateSecretRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSecretResponse(data []byte) (CreateSecretResponse, error) {
	var r CreateSecretResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSecretResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSecretsParams(data []byte) (ListSecretsParams, error) {
	var r ListSecretsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSecretsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSecretsResponse(data []byte) (ListSecretsResponse, error) {
	var r ListSecretsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSecretsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSkillsetAbilityParams(data []byte) (DeleteSkillsetAbilityParams, error) {
	var r DeleteSkillsetAbilityParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSkillsetAbilityParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteSkillsetAbilityRequest map[string]interface{}

func UnmarshalDeleteSkillsetAbilityRequest(data []byte) (DeleteSkillsetAbilityRequest, error) {
	var r DeleteSkillsetAbilityRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSkillsetAbilityRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSkillsetAbilityResponse(data []byte) (DeleteSkillsetAbilityResponse, error) {
	var r DeleteSkillsetAbilityResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSkillsetAbilityResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExecuteSkillsetAbilityParams(data []byte) (ExecuteSkillsetAbilityParams, error) {
	var r ExecuteSkillsetAbilityParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExecuteSkillsetAbilityParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExecuteSkillsetAbilityRequest(data []byte) (ExecuteSkillsetAbilityRequest, error) {
	var r ExecuteSkillsetAbilityRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExecuteSkillsetAbilityRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExecuteSkillsetAbilityResponse(data []byte) (ExecuteSkillsetAbilityResponse, error) {
	var r ExecuteSkillsetAbilityResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExecuteSkillsetAbilityResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSkillsetAbilityParams(data []byte) (FetchSkillsetAbilityParams, error) {
	var r FetchSkillsetAbilityParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSkillsetAbilityParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSkillsetAbilityResponse(data []byte) (FetchSkillsetAbilityResponse, error) {
	var r FetchSkillsetAbilityResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSkillsetAbilityResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSkillsetAbilityParams(data []byte) (UpdateSkillsetAbilityParams, error) {
	var r UpdateSkillsetAbilityParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSkillsetAbilityParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSkillsetAbilityRequest(data []byte) (UpdateSkillsetAbilityRequest, error) {
	var r UpdateSkillsetAbilityRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSkillsetAbilityRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSkillsetAbilityResponse(data []byte) (UpdateSkillsetAbilityResponse, error) {
	var r UpdateSkillsetAbilityResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSkillsetAbilityResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSkillsetAbilityParams(data []byte) (CreateSkillsetAbilityParams, error) {
	var r CreateSkillsetAbilityParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSkillsetAbilityParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSkillsetAbilityRequest(data []byte) (CreateSkillsetAbilityRequest, error) {
	var r CreateSkillsetAbilityRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSkillsetAbilityRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSkillsetAbilityResponse(data []byte) (CreateSkillsetAbilityResponse, error) {
	var r CreateSkillsetAbilityResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSkillsetAbilityResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportSkillsetAbilitiesParams(data []byte) (ExportSkillsetAbilitiesParams, error) {
	var r ExportSkillsetAbilitiesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportSkillsetAbilitiesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportSkillsetAbilitiesResponse(data []byte) (ExportSkillsetAbilitiesResponse, error) {
	var r ExportSkillsetAbilitiesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportSkillsetAbilitiesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSkillsetAbilitiesParams(data []byte) (ListSkillsetAbilitiesParams, error) {
	var r ListSkillsetAbilitiesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSkillsetAbilitiesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSkillsetAbilitiesResponse(data []byte) (ListSkillsetAbilitiesResponse, error) {
	var r ListSkillsetAbilitiesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSkillsetAbilitiesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSkillsetParams(data []byte) (DeleteSkillsetParams, error) {
	var r DeleteSkillsetParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSkillsetParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteSkillsetRequest map[string]interface{}

func UnmarshalDeleteSkillsetRequest(data []byte) (DeleteSkillsetRequest, error) {
	var r DeleteSkillsetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSkillsetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteSkillsetResponse(data []byte) (DeleteSkillsetResponse, error) {
	var r DeleteSkillsetResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteSkillsetResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSkillsetParams(data []byte) (FetchSkillsetParams, error) {
	var r FetchSkillsetParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSkillsetParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSkillsetResponse(data []byte) (FetchSkillsetResponse, error) {
	var r FetchSkillsetResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSkillsetResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSkillsetParams(data []byte) (UpdateSkillsetParams, error) {
	var r UpdateSkillsetParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSkillsetParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSkillsetRequest(data []byte) (UpdateSkillsetRequest, error) {
	var r UpdateSkillsetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSkillsetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSkillsetResponse(data []byte) (UpdateSkillsetResponse, error) {
	var r UpdateSkillsetResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSkillsetResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSkillsetRequest(data []byte) (CreateSkillsetRequest, error) {
	var r CreateSkillsetRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSkillsetRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSkillsetResponse(data []byte) (CreateSkillsetResponse, error) {
	var r CreateSkillsetResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSkillsetResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSkillsetsParams(data []byte) (ListSkillsetsParams, error) {
	var r ListSkillsetsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSkillsetsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSkillsetsResponse(data []byte) (ListSkillsetsResponse, error) {
	var r ListSkillsetsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSkillsetsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSpaceParams(data []byte) (FetchSpaceParams, error) {
	var r FetchSpaceParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSpaceParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchSpaceResponse(data []byte) (FetchSpaceResponse, error) {
	var r FetchSpaceResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchSpaceResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSpaceParams(data []byte) (UpdateSpaceParams, error) {
	var r UpdateSpaceParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSpaceParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSpaceRequest(data []byte) (UpdateSpaceRequest, error) {
	var r UpdateSpaceRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSpaceRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateSpaceResponse(data []byte) (UpdateSpaceResponse, error) {
	var r UpdateSpaceResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateSpaceResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSpaceRequest(data []byte) (CreateSpaceRequest, error) {
	var r CreateSpaceRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSpaceRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateSpaceResponse(data []byte) (CreateSpaceResponse, error) {
	var r CreateSpaceResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateSpaceResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportSpacesParams(data []byte) (ExportSpacesParams, error) {
	var r ExportSpacesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportSpacesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportSpacesResponse(data []byte) (ExportSpacesResponse, error) {
	var r ExportSpacesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportSpacesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSpacesParams(data []byte) (ListSpacesParams, error) {
	var r ListSpacesParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSpacesParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListSpacesResponse(data []byte) (ListSpacesResponse, error) {
	var r ListSpacesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListSpacesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteTaskParams(data []byte) (DeleteTaskParams, error) {
	var r DeleteTaskParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTaskParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeleteTaskRequest map[string]interface{}

func UnmarshalDeleteTaskRequest(data []byte) (DeleteTaskRequest, error) {
	var r DeleteTaskRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTaskRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDeleteTaskResponse(data []byte) (DeleteTaskResponse, error) {
	var r DeleteTaskResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeleteTaskResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchTaskParams(data []byte) (FetchTaskParams, error) {
	var r FetchTaskParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchTaskParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchTaskResponse(data []byte) (FetchTaskResponse, error) {
	var r FetchTaskResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchTaskResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerTaskParams(data []byte) (TriggerTaskParams, error) {
	var r TriggerTaskParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerTaskParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TriggerTaskRequest map[string]interface{}

func UnmarshalTriggerTaskRequest(data []byte) (TriggerTaskRequest, error) {
	var r TriggerTaskRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerTaskRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTriggerTaskResponse(data []byte) (TriggerTaskResponse, error) {
	var r TriggerTaskResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TriggerTaskResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTaskParams(data []byte) (UpdateTaskParams, error) {
	var r UpdateTaskParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTaskParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTaskRequest(data []byte) (UpdateTaskRequest, error) {
	var r UpdateTaskRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTaskRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUpdateTaskResponse(data []byte) (UpdateTaskResponse, error) {
	var r UpdateTaskResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UpdateTaskResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateTaskRequest(data []byte) (CreateTaskRequest, error) {
	var r CreateTaskRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateTaskRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCreateTaskResponse(data []byte) (CreateTaskResponse, error) {
	var r CreateTaskResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateTaskResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportTasksParams(data []byte) (ExportTasksParams, error) {
	var r ExportTasksParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportTasksParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExportTasksResponse(data []byte) (ExportTasksResponse, error) {
	var r ExportTasksResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExportTasksResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTasksParams(data []byte) (ListTasksParams, error) {
	var r ListTasksParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTasksParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTasksResponse(data []byte) (ListTasksResponse, error) {
	var r ListTasksResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTasksResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTeamsParams(data []byte) (ListTeamsParams, error) {
	var r ListTeamsParams
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTeamsParams) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalListTeamsResponse(data []byte) (ListTeamsResponse, error) {
	var r ListTeamsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ListTeamsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchUsageResponse(data []byte) (FetchUsageResponse, error) {
	var r FetchUsageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchUsageResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFetchUsageSeriesResponse(data []byte) (FetchUsageSeriesResponse, error) {
	var r FetchUsageSeriesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FetchUsageSeriesResponse) Marshal() ([]byte, error) {
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

type ListPlatformReportsParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type ListPlatformReportsResponse struct {
	Items []ListPlatformReportsResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformReportsResponseItem struct {
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

type GenerateReportParams struct {
	// The ID of the report to generate       
	ReportID                           string `json:"reportId"`
}

// Successful report output data
type GenerateReportsResponseValue struct {
	// Error message if report generation failed        
	Error                                       *string `json:"error,omitempty"`
}

type CloneBlueprintParams struct {
	// The ID of the blueprint to clone       
	BlueprintID                        string `json:"blueprintId"`
}

type CloneBlueprintResponse struct {
	// The ID of the cloned blueprint                                
	ID                                        string                 `json:"id"`
	// A map of the resources that were cloned                       
	Resources                                 map[string]interface{} `json:"resources"`
}

type DeleteBlueprintParams struct {
	// The ID of the blueprint to delete       
	BlueprintID                         string `json:"blueprintId"`
}

type DeleteBlueprintRequest struct {
	// If true, deletes all resources associated with the blueprint. If false or omitted, only      
	// the blueprint is deleted.                                                                    
	DeleteResources                                                                           *bool `json:"deleteResources,omitempty"`
}

type DeleteBlueprintResponse struct {
	// The ID of the deleted blueprint       
	ID                                string `json:"id"`
}

type FetchBlueprintParams struct {
	// The ID of the blueprint to retrieve       
	BlueprintID                           string `json:"blueprintId"`
}

// Instance list properties
type FetchBlueprintResponse struct {
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

type ListBlueprintResourcesParams struct {
	// The ID of the blueprint to clone       
	BlueprintID                        string `json:"blueprintId"`
}

type ListBlueprintResourcesResponse struct {
	// The ID of the blueprint                       
	ID                        string                 `json:"id"`
	// A map of the resources                        
	Resources                 map[string]interface{} `json:"resources"`
}

type UpdateBlueprintParams struct {
	BlueprintID string `json:"blueprintId"`
}

// Instance crud properties
type UpdateBlueprintRequest struct {
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

type UpdateBlueprintResponse struct {
	// The ID of the updated blueprint       
	ID                                string `json:"id"`
}

// Instance crud properties
type CreateBlueprintRequest struct {
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

type CreateBlueprintResponse struct {
	// The ID of the created blueprint       
	ID                                string `json:"id"`
}

type ListBlueprintsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListBlueprintsResponse struct {
	Items []ListBlueprintsResponseItem `json:"items"`
}

// Instance list properties
type ListBlueprintsResponseItem struct {
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

type CloneBotParams struct {
	BotID string `json:"botId"`
}

type CloneBotResponse struct {
	// The ID of the cloned bot       
	ID                         string `json:"id"`
}

type DeleteBotParams struct {
	// The ID of the bot to delete       
	BotID                         string `json:"botId"`
}

type DeleteBotResponse struct {
	// The ID of the deleted bot       
	ID                          string `json:"id"`
}

type DownvoteBotParams struct {
	// The ID of the bot       
	BotID               string `json:"botId"`
}

type DownvoteBotRequest struct {
	// The reason for the downvote        
	Reason                        *string `json:"reason,omitempty"`
	// The value of the downvote          
	Value                         *int64  `json:"value,omitempty"`
}

type DownvoteBotResponse struct {
	// The bot ID of the downvoted bot       
	ID                                string `json:"id"`
}

type FetchBotParams struct {
	// The ID of the bot to retrieve       
	BotID                           string `json:"botId"`
}

// Blueprint properties
type FetchBotResponse struct {
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

type SearchBotMemoryParams struct {
	// The ID of the bot to search memories for       
	BotID                                      string `json:"botId"`
}

type SearchBotMemoryRequest struct {
	// The keyword/phrase to search for       
	Search                             string `json:"search"`
}

type SearchBotMemoryResponse struct {
	// An array of memories matching the search query                              
	Items                                            []SearchBotMemoryResponseItem `json:"items"`
}

type SearchBotMemoryResponseItem struct {
	ID   string                 `json:"id"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Text string                 `json:"text"`
}

type CreateBotSessionParams struct {
	// The ID of the bot for this session       
	BotID                                string `json:"botId"`
}

type CreateBotSessionRequest struct {
	// The maximum amount of time this session will stay open                                  
	DurationInSeconds                                         *float64                         `json:"durationInSeconds,omitempty"`
	// An array of messages to be included in the conversation                                 
	Messages                                                  []CreateBotSessionRequestMessage `json:"messages,omitempty"`
	// Meta data information                                                                   
	Meta                                                      map[string]interface{}           `json:"meta,omitempty"`
}

type CreateBotSessionRequestMessage struct {
	// The text of the message            
	Text                      string      `json:"text"`
	// The type of the message            
	Type                      MessageType `json:"type"`
}

type CreateBotSessionResponse struct {
	// The ID of the conversation                                                         
	ConversationID                                      string                            `json:"conversationId"`
	// The time the token will expire in milliseconds                                     
	ExpiresAt                                           float64                           `json:"expiresAt"`
	// The ID of the bot                                                                  
	ID                                                  string                            `json:"id"`
	// An array of messages included in the conversation                                  
	Messages                                            []CreateBotSessionResponseMessage `json:"messages,omitempty"`
	// The token for this conversation                                                    
	Token                                               string                            `json:"token"`
}

type CreateBotSessionResponseMessage struct {
	// The text of the message            
	Text                      string      `json:"text"`
	// The type of the message            
	Type                      MessageType `json:"type"`
}

type UpdateBotParams struct {
	BotID string `json:"botId"`
}

// Blueprint properties
type UpdateBotRequest struct {
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

type UpdateBotResponse struct {
	// The ID of the updated bot       
	ID                          string `json:"id"`
}

type UpvoteBotParams struct {
	// The ID of the bot       
	BotID               string `json:"botId"`
}

type UpvoteBotRequest struct {
	// The reason for the upvote        
	Reason                      *string `json:"reason,omitempty"`
	// The value of the upvote          
	Value                       *int64  `json:"value,omitempty"`
}

type UpvoteBotResponse struct {
	// The ID of the upvoted bot       
	ID                          string `json:"id"`
}

type FetchBotUsageParams struct {
	// The ID of the bot                                     
	BotID                                         string     `json:"botId"`
	// Start date for the period (ISO 8601 format)           
	From                                          *time.Time `json:"from,omitempty"`
	// End date for the period (ISO 8601 format)             
	To                                            *time.Time `json:"to,omitempty"`
}

type FetchBotUsageResponse struct {
	// Total number of conversations          
	Conversations                      *int64 `json:"conversations,omitempty"`
	// Total number of messages               
	Messages                           *int64 `json:"messages,omitempty"`
	// Total number of BASE tokens used       
	Tokens                             *int64 `json:"tokens,omitempty"`
}

// Blueprint properties
type CreateBotRequest struct {
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

type CreateBotResponse struct {
	// The ID of the created bot       
	ID                          string `json:"id"`
}

type ListBotsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListBotsResponse struct {
	Items []ListBotsResponseItem `json:"items"`
}

// Blueprint properties
type ListBotsResponseItem struct {
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

type PublishChannelMessageParams struct {
	// The ID of the channel to publish to (minimum 16 characters)       
	ChannelID                                                     string `json:"channelId"`
}

type PublishChannelMessageRequest struct {
	// The message to publish to the channel                       
	Message                                 map[string]interface{} `json:"message"`
}

type PublishChannelMessageResponse struct {
	// The ID of the channel the message was published to       
	ID                                                   string `json:"id"`
}

type SubscribeChannelMessagesParams struct {
	// The ID of the channel to subscribe to (minimum 16 characters)       
	ChannelID                                                       string `json:"channelId"`
}

type SubscribeChannelMessagesRequest struct {
	// Number of historical messages to replay from the channel       
	// before subscribing to live updates. When provided, the         
	// subscriber will first receive up to this many recent           
	// messages that were published before the subscription           
	// started. This is useful for catching up on messages that       
	// may have been published during connection setup.               
	HistoryLength                                              *int64 `json:"historyLength,omitempty"`
}

type ListContactConversationsParams struct {
	// The ID of the contact to list conversations for        
	ContactID                                         string  `json:"contactId"`
	// The cursor to use for pagination                       
	Cursor                                            *string `json:"cursor,omitempty"`
	// The order of the paginated items                       
	Order                                             *Order  `json:"order,omitempty"`
	// The number of items to retrieve                        
	Take                                              *int64  `json:"take,omitempty"`
}

type ListContactConversationsResponse struct {
	Items []ListContactConversationsResponseItem `json:"items"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ListContactConversationsResponseItem struct {
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

type DeleteContactParams struct {
	// The ID of the contact to delete       
	ContactID                         string `json:"contactId"`
}

type DeleteContactResponse struct {
	// The ID of the deleted contact       
	ID                              string `json:"id"`
}

type FetchContactParams struct {
	// The ID of the contact to retrieve       
	ContactID                           string `json:"contactId"`
}

// Instance list properties
type FetchContactResponse struct {
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

type ListContactMemoriesParams struct {
	// The ID of the contact to list memories for        
	ContactID                                    string  `json:"contactId"`
	// The cursor to use for pagination                  
	Cursor                                       *string `json:"cursor,omitempty"`
	// The order of the paginated items                  
	Order                                        *Order  `json:"order,omitempty"`
	// The number of items to retrieve                   
	Take                                         *int64  `json:"take,omitempty"`
}

type ListContactMemoriesResponse struct {
	Items []ListContactMemoriesResponseItem `json:"items"`
}

// Instance list properties
type ListContactMemoriesResponseItem struct {
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

type SearchContactMemoryParams struct {
	// The ID of the contact to search memories for       
	ContactID                                      string `json:"contactId"`
}

type SearchContactMemoryRequest struct {
	// The keyword/phrase to search for       
	Search                             string `json:"search"`
}

type SearchContactMemoryResponse struct {
	// An array of memories matching the search query                                  
	Items                                            []SearchContactMemoryResponseItem `json:"items"`
}

type SearchContactMemoryResponseItem struct {
	ID   string                 `json:"id"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Text string                 `json:"text"`
}

type AuthenticateContactSecretParams struct {
	// The ID of the contact the secret belongs to       
	ContactID                                     string `json:"contactId"`
	// The ID of the secret to authenticate              
	SecretID                                      string `json:"secretId"`
}

type AuthenticateContactSecretResponse struct {
	// The ID of the secret to authenticate       
	ID                                     string `json:"id"`
	// The URL to authenticate the secret         
	URL                                    string `json:"url"`
}

type RevokeContactSecretParams struct {
	// The ID of the contact the secret belongs to       
	ContactID                                     string `json:"contactId"`
	// The ID of the secret to be revoked                
	SecretID                                      string `json:"secretId"`
}

type RevokeContactSecretResponse struct {
	// The ID of the revoked secret       
	ID                             string `json:"id"`
}

type VerifyContactSecretParams struct {
	// The ID of the contact the secret belongs to       
	ContactID                                     string `json:"contactId"`
	// The ID of the secret to be verified               
	SecretID                                      string `json:"secretId"`
}

type VerifyContactSecretResponse struct {
	Action                          *VerifyContactSecretResponseAction `json:"action,omitempty"`
	// The ID of the verified secret                                   
	ID                              string                             `json:"id"`
	// The status of the secret                                        
	Status                          Status                             `json:"status"`
}

// The action to take next
type VerifyContactSecretResponseAction struct {
	// The type of action to take                   
	Type                                 ActionType `json:"type"`
	// The URL to authenticate the secret           
	URL                                  string     `json:"url"`
}

type ListContactSecretsParams struct {
	// The ID of the contact to list secrets for        
	ContactID                                   string  `json:"contactId"`
	// The cursor to use for pagination                 
	Cursor                                      *string `json:"cursor,omitempty"`
	// The order of the paginated items                 
	Order                                       *Order  `json:"order,omitempty"`
	// The number of items to retrieve                  
	Take                                        *int64  `json:"take,omitempty"`
}

type ListContactSecretsResponse struct {
	Items []ListContactSecretsResponseItem `json:"items"`
}

// Instance list properties
type ListContactSecretsResponseItem struct {
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

type ListContactSpacesParams struct {
	// The ID of the contact to list spaces for        
	ContactID                                  string  `json:"contactId"`
	// The cursor to use for pagination                
	Cursor                                     *string `json:"cursor,omitempty"`
	// The order of the paginated items                
	Order                                      *Order  `json:"order,omitempty"`
	// The number of items to retrieve                 
	Take                                       *int64  `json:"take,omitempty"`
}

type ListContactSpacesResponse struct {
	Items []ListContactSpacesResponseItem `json:"items"`
}

// Instance list properties
type ListContactSpacesResponseItem struct {
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

type ListContactTasksParams struct {
	// The ID of the contact to list tasks for        
	ContactID                                 string  `json:"contactId"`
	// The cursor to use for pagination               
	Cursor                                    *string `json:"cursor,omitempty"`
	// The order of the paginated items               
	Order                                     *Order  `json:"order,omitempty"`
	// The number of items to retrieve                
	Take                                      *int64  `json:"take,omitempty"`
}

type ListContactTasksResponse struct {
	Items []ListContactTasksResponseItem `json:"items"`
}

// Instance list properties
type ListContactTasksResponseItem struct {
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

type UpdateContactParams struct {
	ContactID string `json:"contactId"`
}

// Instance crud properties
type UpdateContactRequest struct {
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

type UpdateContactResponse struct {
	// The ID of the updated contact       
	ID                              string `json:"id"`
}

// Instance crud properties
type CreateContactRequest struct {
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

type CreateContactResponse struct {
	// The ID of the created contact       
	ID                              string `json:"id"`
}

// Instance crud properties
type EnsureContactRequest struct {
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

type EnsureContactResponse struct {
	// The ID of the ensured contact       
	ID                              string `json:"id"`
}

type ExportContactsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ExportContactsResponse struct {
	Items []ExportContactsResponseItem `json:"items"`
}

// Instance list properties
type ExportContactsResponseItem struct {
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

type ListContactsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListContactsResponse struct {
	Items []ListContactsResponseItem `json:"items"`
}

// Instance list properties
type ListContactsResponseItem struct {
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

type UploadConversationAttachmentParams struct {
	ConversationID string `json:"conversationId"`
}

type UploadConversationAttachmentRequest struct {
	// The file to upload either as http: or data: URL                                         
	//                                                                                         
	// The file definition to upload                                                           
	File                                              *UploadConversationAttachmentRequestFile `json:"file"`
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

type UploadConversationAttachmentResponse struct {
	// The ID of the upload file                                                                 
	ID                                        string                                             `json:"id"`
	// The name of the uploaded file                                                             
	Name                                      *string                                            `json:"name,omitempty"`
	// The request required to upload the file                                                   
	UploadRequest                             *UploadConversationAttachmentResponseUploadRequest `json:"uploadRequest,omitempty"`
}

// The request required to upload the file
type UploadConversationAttachmentResponseUploadRequest struct {
	// The HTTP headers to use                       
	Headers                   map[string]interface{} `json:"headers"`
	// The HTTP method to use                        
	Method                    string                 `json:"method"`
	// The HTTP url to use                           
	URL                       string                 `json:"url"`
}

type CompleteConversationMessageParams struct {
	// The ID of the conversation to receive message from       
	ConversationID                                       string `json:"conversationId"`
}

type CompleteConversationMessageRequest struct {
	// Known entities                                                                                          
	Entities                                                     []CompleteConversationMessageRequestEntity    `json:"entities,omitempty"`
	// Extensions to enhance the bot's capabilities                                                            
	Extensions                                                   *CompleteConversationMessageRequestExtensions `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                                   
	Functions                                                    []CompleteConversationMessageRequestFunction  `json:"functions,omitempty"`
	// Execution limits to control conversation processing bounds                                              
	Limits                                                       *CompleteConversationMessageRequestLimits     `json:"limits,omitempty"`
	// The text of the message to send                                                                         
	Text                                                         string                                        `json:"text"`
}

// Extracted entity from the message
type CompleteConversationMessageRequestEntity struct {
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
type CompleteConversationMessageRequestExtensions struct {
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

type CompleteConversationMessageRequestFunction struct {
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
type CompleteConversationMessageRequestLimits struct {
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

type CompleteConversationMessageResponse struct {
	// Information about why the completion ended                                         
	End                                          CompleteConversationMessageResponseEnd   `json:"end"`
	// The ID of the created message                                                      
	ID                                           string                                   `json:"id"`
	// The text of the message received                                                   
	Text                                         string                                   `json:"text"`
	// Usage information                                                                  
	Usage                                        CompleteConversationMessageResponseUsage `json:"usage"`
}

// Information about why the completion ended
type CompleteConversationMessageResponseEnd struct {
	// The reason why the completion ended               
	Reason                                CompleteReason `json:"reason"`
}

// Usage information
type CompleteConversationMessageResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

type UpsertConversationContactParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

// Instance crud properties
type UpsertConversationContactRequest struct {
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

type UpsertConversationContactResponse struct {
	// The ID of the created contact       
	ID                              string `json:"id"`
}

type DeleteConversationParams struct {
	// The ID of the conversation to delete       
	ConversationID                         string `json:"conversationId"`
}

type DeleteConversationResponse struct {
	// The ID of the deleted conversation       
	ID                                   string `json:"id"`
}

type DispatchStatefulConversationRequest struct {
	// A unique ID to deduplicate dispatch requests                                                             
	ChannelID                                                    *string                                        `json:"channelId,omitempty"`
	// Known entities                                                                                           
	Entities                                                     []DispatchStatefulConversationRequestEntity    `json:"entities,omitempty"`
	// Extensions to enhance the bot's capabilities                                                             
	Extensions                                                   *DispatchStatefulConversationRequestExtensions `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                                    
	Functions                                                    []DispatchStatefulConversationRequestFunction  `json:"functions,omitempty"`
	// Execution limits to control conversation processing bounds                                               
	Limits                                                       *DispatchStatefulConversationRequestLimits     `json:"limits,omitempty"`
	// The text of the message to send                                                                          
	Text                                                         string                                         `json:"text"`
}

// Extracted entity from the message
type DispatchStatefulConversationRequestEntity struct {
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
type DispatchStatefulConversationRequestExtensions struct {
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

type DispatchStatefulConversationRequestFunction struct {
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
type DispatchStatefulConversationRequestLimits struct {
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

type DispatchStatefulConversationResponse struct {
	// The channel ID to subscribe to for completion events       
	ChannelID                                              string `json:"channelId"`
}

type DownvoteConversationParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

type DownvoteConversationRequest struct {
	// The reason for the downvote        
	Reason                        *string `json:"reason,omitempty"`
	// The value of the downvote          
	Value                         *int64  `json:"value,omitempty"`
}

type DownvoteConversationResponse struct {
	// The conversation ID of the downvoted conversation       
	ID                                                  string `json:"id"`
}

type FetchConversationParams struct {
	// The ID of the conversation to retrieve       
	ConversationID                           string `json:"conversationId"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type FetchConversationResponse struct {
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

type DeleteConversationMessageParams struct {
	// The ID of the conversation containing the message       
	ConversationID                                      string `json:"conversationId"`
	// The ID of the message to delete                         
	MessageID                                           string `json:"messageId"`
}

type DeleteConversationMessageResponse struct {
	// The ID of the deleted message       
	ID                              string `json:"id"`
}

type DownvoteConversationMessageParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
	// The ID of the message            
	MessageID                    string `json:"messageId"`
}

type DownvoteConversationMessageRequest struct {
	// The reason for the downvote        
	Reason                        *string `json:"reason,omitempty"`
	// The value of the downvote          
	Value                         *int64  `json:"value,omitempty"`
}

type DownvoteConversationMessageResponse struct {
	// The ID of the downvoted message       
	ID                                string `json:"id"`
}

type FetchConversationMessageParams struct {
	// The ID of the conversation containing the message       
	ConversationID                                      string `json:"conversationId"`
	// The ID of the message to retrieve                       
	MessageID                                           string `json:"messageId"`
}

// Instance list properties
type FetchConversationMessageResponse struct {
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

type SynthesizeConversationMessageParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
	// The ID of the message            
	MessageID                    string `json:"messageId"`
}

type SynthesizeConversationMessageResponse struct {
	// The ID of the synthesized message       
	ID                                  string `json:"id"`
}

type UpdateConversationMessageParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
	// The ID of the message            
	MessageID                    string `json:"messageId"`
}

// Instance crud properties
type UpdateConversationMessageRequest struct {
	// The associated description                                              
	Description                       *string                                  `json:"description,omitempty"`
	// Known entities                                                          
	Entities                          []UpdateConversationMessageRequestEntity `json:"entities,omitempty"`
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
type UpdateConversationMessageRequestEntity struct {
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

type UpdateConversationMessageResponse struct {
	// The ID of the updated message       
	ID                              string `json:"id"`
}

type UpvoteConversationMessageParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
	// The ID of the message            
	MessageID                    string `json:"messageId"`
}

type UpvoteConversationMessageRequest struct {
	// The reason for the upvote        
	Reason                      *string `json:"reason,omitempty"`
	// The value of the upvote          
	Value                       *int64  `json:"value,omitempty"`
}

type UpvoteConversationMessageResponse struct {
	// The ID of the upvoted message       
	ID                              string `json:"id"`
}

type CreateConversationMessageParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

// Instance crud properties
type CreateConversationMessageRequest struct {
	// The associated description                                         
	Description                  *string                                  `json:"description,omitempty"`
	// Known entities                                                     
	Entities                     []CreateConversationMessageRequestEntity `json:"entities,omitempty"`
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
type CreateConversationMessageRequestEntity struct {
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

type CreateConversationMessageResponse struct {
	// Extracted entities from the message                                          
	Entities                              []CreateConversationMessageResponseEntity `json:"entities"`
	// The ID of the created message                                                
	ID                                    string                                    `json:"id"`
}

// Extracted entity from the message
type CreateConversationMessageResponseEntity struct {
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

type ListConversationMessagesParams struct {
	// The ID of the conversation to list messages for        
	ConversationID                                    string  `json:"conversationId"`
	// The cursor to use for pagination                       
	Cursor                                            *string `json:"cursor,omitempty"`
	// The order of the paginated items                       
	Order                                             *Order  `json:"order,omitempty"`
	// The number of items to retrieve                        
	Take                                              *int64  `json:"take,omitempty"`
}

type ListConversationMessagesResponse struct {
	Items []ListConversationMessagesResponseItem `json:"items"`
}

// Instance list properties
type ListConversationMessagesResponseItem struct {
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

type ReceiveConversationMessageParams struct {
	// The ID of the conversation to receive message from       
	ConversationID                                       string `json:"conversationId"`
}

type ReceiveConversationMessageRequest struct {
	// Extensions to enhance the bot's capabilities                                                      
	Extensions                                              *ReceiveConversationMessageRequestExtensions `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                             
	Functions                                               []ReceiveConversationMessageRequestFunction  `json:"functions,omitempty"`
}

// Extensions to enhance the bot's capabilities
type ReceiveConversationMessageRequestExtensions struct {
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

type ReceiveConversationMessageRequestFunction struct {
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

type ReceiveConversationMessageResponse struct {
	// The ID of the created message                                           
	ID                                 string                                  `json:"id"`
	// The text of the message received                                        
	Text                               string                                  `json:"text"`
	// Usage information                                                       
	Usage                              ReceiveConversationMessageResponseUsage `json:"usage"`
}

// Usage information
type ReceiveConversationMessageResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

type SendConversationMessageParams struct {
	// The ID of the conversation to send the message to       
	ConversationID                                      string `json:"conversationId"`
}

type SendConversationMessageRequest struct {
	// Known entities                                                                                 
	Entities                                                []SendConversationMessageRequestEntity    `json:"entities,omitempty"`
	// Extensions to enhance the bot's capabilities                                                   
	Extensions                                              *SendConversationMessageRequestExtensions `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                          
	Functions                                               []SendConversationMessageRequestFunction  `json:"functions,omitempty"`
	// The text of the message to send                                                                
	Text                                                    string                                    `json:"text"`
}

// Extracted entity from the message
type SendConversationMessageRequestEntity struct {
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
type SendConversationMessageRequestExtensions struct {
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

type SendConversationMessageRequestFunction struct {
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

type SendConversationMessageResponse struct {
	// Extracted entities from the message                                        
	Entities                              []SendConversationMessageResponseEntity `json:"entities"`
	// The ID of the sent message                                                 
	ID                                    string                                  `json:"id"`
}

// Extracted entity from the message
type SendConversationMessageResponseEntity struct {
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

type CreateConversationSessionParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

type CreateConversationSessionRequest struct {
	// The maximum amount of time this session will stay open         
	DurationInSeconds                                        *float64 `json:"durationInSeconds,omitempty"`
}

type CreateConversationSessionResponse struct {
	// The time the token will expire in milliseconds        
	ExpiresAt                                        float64 `json:"expiresAt"`
	// The ID of the conversation                            
	ID                                               string  `json:"id"`
	// The token for this conversation                       
	Token                                            string  `json:"token"`
}

type UpdateConversationParams struct {
	ConversationID string `json:"conversationId"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type UpdateConversationRequest struct {
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

type UpdateConversationResponse struct {
	// The ID of the updated conversation       
	ID                                   string `json:"id"`
}

type UpvoteConversationParams struct {
	// The ID of the conversation       
	ConversationID               string `json:"conversationId"`
}

type UpvoteConversationRequest struct {
	// The reason for the upvote        
	Reason                      *string `json:"reason,omitempty"`
	// The value of the upvote          
	Value                       *int64  `json:"value,omitempty"`
}

type UpvoteConversationResponse struct {
	// The ID of the upvoted conversation       
	ID                                   string `json:"id"`
}

type FetchConversationUsageParams struct {
	// The ID of the conversation                            
	ConversationID                                string     `json:"conversationId"`
	// Start date for the period (ISO 8601 format)           
	From                                          *time.Time `json:"from,omitempty"`
	// End date for the period (ISO 8601 format)             
	To                                            *time.Time `json:"to,omitempty"`
}

type FetchConversationUsageResponse struct {
	// Total number of messages               
	Messages                           *int64 `json:"messages,omitempty"`
	// Total number of BASE tokens used       
	Tokens                             *int64 `json:"tokens,omitempty"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type CompleteConversationRequest struct {
	// An array of attachments to be added to the conversation                                           
	Attachments                                                  []CompleteConversationRequestAttachment `json:"attachments,omitempty"`
	// The contact ID to associate with this conversation                                                
	ContactID                                                    *CompleteConversationRequestContactID   `json:"contactId"`
	// Extensions to enhance the bot's capabilities                                                      
	Extensions                                                   *CompleteConversationRequestExtensions  `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                             
	Functions                                                    []CompleteConversationRequestFunction   `json:"functions,omitempty"`
	// Execution limits to control conversation processing bounds                                        
	Limits                                                       *CompleteConversationRequestLimits      `json:"limits,omitempty"`
	// An array of messages to be added to the conversation                                              
	Messages                                                     []CompleteConversationRequestMessage    `json:"messages"`
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

type CompleteConversationRequestAttachment struct {
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
type CompleteConversationRequestExtensions struct {
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

type CompleteConversationRequestFunction struct {
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
type CompleteConversationRequestLimits struct {
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
type CompleteConversationRequestMessage struct {
	// Meta data information                         
	Meta                      map[string]interface{} `json:"meta,omitempty"`
	// The text of the message                       
	Text                      string                 `json:"text"`
	// The type of the message                       
	Type                      MessageType            `json:"type"`
}

type CompleteConversationResponse struct {
	// Information about why the completion ended                                  
	End                                          CompleteConversationResponseEnd   `json:"end"`
	// The text of the message received                                            
	Text                                         string                            `json:"text"`
	// Usage information                                                           
	Usage                                        CompleteConversationResponseUsage `json:"usage"`
}

// Information about why the completion ended
type CompleteConversationResponseEnd struct {
	// The reason why the completion ended               
	Reason                                CompleteReason `json:"reason"`
}

// Usage information
type CompleteConversationResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type CreateConversationRequest struct {
	// The contact id assigned to this conversation                                           
	ContactID                                              *string                            `json:"contactId,omitempty"`
	// The associated description                                                             
	Description                                            *string                            `json:"description,omitempty"`
	// An array of messages to be added to the conversation                                   
	Messages                                               []CreateConversationRequestMessage `json:"messages,omitempty"`
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

type CreateConversationRequestMessage struct {
	// The text of the message            
	Text                      string      `json:"text"`
	// The type of the message            
	Type                      MessageType `json:"type"`
}

type CreateConversationResponse struct {
	// The ID of the created conversation                                                   
	ID                                                  string                              `json:"id"`
	// An array of messages included in the conversation                                    
	Messages                                            []CreateConversationResponseMessage `json:"messages,omitempty"`
}

type CreateConversationResponseMessage struct {
	// The text of the message            
	Text                      string      `json:"text"`
	// The type of the message            
	Type                      MessageType `json:"type"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type DispatchConversationRequest struct {
	// An array of attachments to be added to the conversation                                           
	Attachments                                                  []DispatchConversationRequestAttachment `json:"attachments,omitempty"`
	// A unique channel ID to subscribe to for completion events                                         
	ChannelID                                                    *string                                 `json:"channelId,omitempty"`
	// The contact ID to associate with this conversation                                                
	ContactID                                                    *DispatchConversationRequestContactID   `json:"contactId"`
	// Extensions to enhance the bot's capabilities                                                      
	Extensions                                                   *DispatchConversationRequestExtensions  `json:"extensions,omitempty"`
	// An array of functions to be added to the conversation                                             
	Functions                                                    []DispatchConversationRequestFunction   `json:"functions,omitempty"`
	// Execution limits to control conversation processing bounds                                        
	Limits                                                       *DispatchConversationRequestLimits      `json:"limits,omitempty"`
	// An array of messages to be added to the conversation                                              
	Messages                                                     []DispatchConversationRequestMessage    `json:"messages"`
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

type DispatchConversationRequestAttachment struct {
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
type DispatchConversationRequestExtensions struct {
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

type DispatchConversationRequestFunction struct {
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
type DispatchConversationRequestLimits struct {
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
type DispatchConversationRequestMessage struct {
	// Meta data information                         
	Meta                      map[string]interface{} `json:"meta,omitempty"`
	// The text of the message                       
	Text                      string                 `json:"text"`
	// The type of the message                       
	Type                      MessageType            `json:"type"`
}

type DispatchConversationResponse struct {
	// The channel ID to subscribe to for completion events       
	ChannelID                                              string `json:"channelId"`
}

type ExportConversationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ExportConversationsResponse struct {
	Items []ExportConversationsResponseItem `json:"items"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ExportConversationsResponseItem struct {
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

type ListConversationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListConversationsResponse struct {
	Items []ListConversationsResponseItem `json:"items"`
}

// A bot configuration or reference
//
// A bot configuration that can be applied without a dedicated bot instance.
type ListConversationsResponseItem struct {
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

type DeleteDatasetParams struct {
	// The ID of the dataset to delete       
	DatasetID                         string `json:"datasetId"`
}

type DeleteDatasetResponse struct {
	// The ID of the deleted dataset       
	ID                              string `json:"id"`
}

type FetchDatasetParams struct {
	// The ID of the dataset to retrieve       
	DatasetID                           string `json:"datasetId"`
}

// Blueprint properties
type FetchDatasetResponse struct {
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

type AttachDatasetFileParams struct {
	// The ID of the dataset       
	DatasetID               string `json:"datasetId"`
	// The ID of the file          
	FileID                  string `json:"fileId"`
}

type AttachDatasetFileRequest struct {
	// The dataset file attachment type                           
	Type                               *DatasetFileAttachmentType `json:"type,omitempty"`
}

type AttachDatasetFileResponse struct {
	// The ID of the dataset file       
	ID                           string `json:"id"`
}

type DetachDatasetFileParams struct {
	// The ID of the dataset       
	DatasetID               string `json:"datasetId"`
	// The ID of the file          
	FileID                  string `json:"fileId"`
}

type DetachDatasetFileRequest struct {
	// Delete records associated with the file      
	DeleteRecords                             *bool `json:"deleteRecords,omitempty"`
}

type DetachDatasetFileResponse struct {
	// The ID of the dataset file       
	ID                           string `json:"id"`
}

type SyncDatasetFileParams struct {
	// The ID of the dataset       
	DatasetID               string `json:"datasetId"`
	// The ID of the file          
	FileID                  string `json:"fileId"`
}

type SyncDatasetFileResponse struct {
	// The ID of the dataset file       
	ID                           string `json:"id"`
}

type ListDatasetFilesParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The ID of the dataset                   
	DatasetID                          string  `json:"datasetId"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type ListDatasetFilesResponse struct {
	Items []ListDatasetFilesResponseItem `json:"items"`
}

// Instance list properties
type ListDatasetFilesResponseItem struct {
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

type DeleteDatasetRecordParams struct {
	// The ID of the dataset                
	DatasetID                        string `json:"datasetId"`
	// The ID of the record to delete       
	RecordID                         string `json:"recordId"`
}

type DeleteDatasetRecordResponse struct {
	// The ID of the deleted record       
	ID                             string `json:"id"`
}

type FetchDatasetRecordParams struct {
	// The ID of the dataset                  
	DatasetID                          string `json:"datasetId"`
	// The ID of the record to retrieve       
	RecordID                           string `json:"recordId"`
}

// Instance list properties
type FetchDatasetRecordResponse struct {
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

type UpdateDatasetRecordParams struct {
	DatasetID string `json:"datasetId"`
	RecordID  string `json:"recordId"`
}

type UpdateDatasetRecordRequest struct {
	// Meta data information                                      
	Meta                                   map[string]interface{} `json:"meta,omitempty"`
	// The source to update the record with                       
	Source                                 *string                `json:"source,omitempty"`
	// The text to update the record with                         
	Text                                   *string                `json:"text,omitempty"`
}

type UpdateDatasetRecordResponse struct {
	// The ID of the updated record       
	ID                             string `json:"id"`
}

type CreateDatasetRecordParams struct {
	DatasetID string `json:"datasetId"`
}

type CreateDatasetRecordRequest struct {
	// Meta data information                          
	Meta                       map[string]interface{} `json:"meta,omitempty"`
	// The source of the record                       
	Source                     *string                `json:"source,omitempty"`
	// The text of the record                         
	Text                       string                 `json:"text"`
}

type CreateDatasetRecordResponse struct {
	// The ID of the created record       
	ID                             string `json:"id"`
}

type ExportDatasetRecordsParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The ID of the dataset to export         
	DatasetID                          string  `json:"datasetId"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type ExportDatasetRecordsResponse struct {
	Items []ExportDatasetRecordsResponseItem `json:"items"`
}

// Instance list properties
type ExportDatasetRecordsResponseItem struct {
	// The timestamp (ms) when the instance was created        
	CreatedAt                                          float64 `json:"createdAt"`
	// The instance ID                                         
	ID                                                 string  `json:"id"`
	Source                                             *string `json:"source,omitempty"`
	Text                                               string  `json:"text"`
	// The timestamp (ms) when the instance was updated        
	UpdatedAt                                          float64 `json:"updatedAt"`
}

type ListDatasetRecordsParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The ID of the dataset                   
	DatasetID                          string  `json:"datasetId"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type ListDatasetRecordsResponse struct {
	Items []ListDatasetRecordsResponseItem `json:"items"`
}

// Instance list properties
type ListDatasetRecordsResponseItem struct {
	// The timestamp (ms) when the instance was created        
	CreatedAt                                          float64 `json:"createdAt"`
	// The instance ID                                         
	ID                                                 string  `json:"id"`
	Source                                             *string `json:"source,omitempty"`
	Text                                               string  `json:"text"`
	// The timestamp (ms) when the instance was updated        
	UpdatedAt                                          float64 `json:"updatedAt"`
}

type SearchDatasetParams struct {
	// The ID of the dataset to search       
	DatasetID                         string `json:"datasetId"`
}

type SearchDatasetRequest struct {
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

type SearchDatasetResponse struct {
	// The ID of the dataset that was searched                                    
	ID                                              string                        `json:"id"`
	// An array of records matching the search query                              
	Records                                         []SearchDatasetResponseRecord `json:"records"`
}

type SearchDatasetResponseRecord struct {
	ID     string                 `json:"id"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
	Score  float64                `json:"score"`
	Source *string                `json:"source,omitempty"`
	Text   string                 `json:"text"`
}

type UpdateDatasetParams struct {
	DatasetID string `json:"datasetId"`
}

// Blueprint properties
type UpdateDatasetRequest struct {
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

type UpdateDatasetResponse struct {
	// The ID of the updated dataset       
	ID                              string `json:"id"`
}

// Blueprint properties
type CreateDatasetRequest struct {
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

type CreateDatasetResponse struct {
	// The ID of the created dataset       
	ID                              string `json:"id"`
}

type ListDatasetsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListDatasetsResponse struct {
	Items []ListDatasetsResponseItem `json:"items"`
}

// Blueprint properties
type ListDatasetsResponseItem struct {
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

type ExportEventLogsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ExportEventLogsResponse struct {
	Items []ExportEventLogsResponseItem `json:"items"`
}

// Instance list properties
type ExportEventLogsResponseItem struct {
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

type ListEventLogsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListEventLogsResponse struct {
	Items []ListEventLogsResponseItem `json:"items"`
}

// Instance list properties
type ListEventLogsResponseItem struct {
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

type SubscribeEventLogsRequest struct {
	// Number of recent historical events to replay before              
	// subscribing to live updates. When provided, the subscriber       
	// will first receive up to this many recent events that were       
	// logged before the subscription started. This is useful for       
	// catching up on events that may have occurred during              
	// connection setup.                                                
	HistoryLength                                                *int64 `json:"historyLength,omitempty"`
}

type DeleteFileParams struct {
	// The ID of the file to delete       
	FileID                         string `json:"fileId"`
}

type DeleteFileResponse struct {
	// The ID of the deleted file       
	ID                           string `json:"id"`
}

type DownloadFileParams struct {
	// The ID of the file to download       
	FileID                           string `json:"fileId"`
}

type DownloadFileResponse struct {
	// The URL to download the file       
	URL                            string `json:"url"`
}

type FetchFileParams struct {
	// The ID of the file to retrieve       
	FileID                           string `json:"fileId"`
}

// Blueprint properties
type FetchFileResponse struct {
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

type SyncFileParams struct {
	// The ID of the file to sync       
	FileID                       string `json:"fileId"`
}

type SyncFileResponse struct {
	// The ID of the file       
	ID                   string `json:"id"`
}

type UpdateFileParams struct {
	FileID string `json:"fileId"`
}

// Blueprint properties
type UpdateFileRequest struct {
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

type UpdateFileResponse struct {
	// The ID of the updated file       
	ID                           string `json:"id"`
}

type UploadFileParams struct {
	FileID string `json:"fileId"`
}

type UploadFileRequest struct {
	// The file to upload either as http: or data: URL                       
	//                                                                       
	// The file definition to upload                                         
	File                                              *UploadFileRequestFile `json:"file"`
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

type UploadFileResponse struct {
	// The ID of the upload file                                               
	ID                                        string                           `json:"id"`
	// The request required to upload the file                                 
	UploadRequest                             *UploadFileResponseUploadRequest `json:"uploadRequest,omitempty"`
}

// The request required to upload the file
type UploadFileResponseUploadRequest struct {
	// The HTTP headers to use                       
	Headers                   map[string]interface{} `json:"headers"`
	// The HTTP method to use                        
	Method                    string                 `json:"method"`
	// The HTTP url to use                           
	URL                       string                 `json:"url"`
}

// Blueprint properties
type CreateFileRequest struct {
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

type CreateFileResponse struct {
	// The ID of the created file       
	ID                           string `json:"id"`
}

type ListFilesParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListFilesResponse struct {
	Items []ListFilesResponseItem `json:"items"`
}

// Blueprint properties
type ListFilesResponseItem struct {
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

type DeleteDiscordIntegrationParams struct {
	// The ID of the Discord integration       
	DiscordIntegrationID                string `json:"discordIntegrationId"`
}

type DeleteDiscordIntegrationResponse struct {
	// The ID of the deleted Discord integration       
	ID                                          string `json:"id"`
}

type FetchDiscordIntegrationParams struct {
	// The ID of the Discord integration to retrieve       
	DiscordIntegrationID                            string `json:"discordIntegrationId"`
}

// Blueprint properties
type FetchDiscordIntegrationResponse struct {
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

type SetupDiscordIntegrationParams struct {
	// The ID of the Discord integration       
	DiscordIntegrationID                string `json:"discordIntegrationId"`
}

type SetupDiscordIntegrationResponse struct {
	// The ID of the setup Discord integration       
	ID                                        string `json:"id"`
}

type UpdateDiscordIntegrationParams struct {
	// The ID of the Discord integration       
	DiscordIntegrationID                string `json:"discordIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateDiscordIntegrationRequest struct {
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

type UpdateDiscordIntegrationResponse struct {
	// The ID of the Discord Integration       
	ID                                  string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateDiscordIntegrationRequest struct {
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

type CreateDiscordIntegrationResponse struct {
	// The ID of the Discord Integration       
	ID                                  string `json:"id"`
}

type ListDiscordIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListDiscordIntegrationsResponse struct {
	Items []ListDiscordIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListDiscordIntegrationsResponseItem struct {
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

type DeleteEmailIntegrationParams struct {
	// The ID of the Email integration       
	EmailIntegrationID                string `json:"emailIntegrationId"`
}

type DeleteEmailIntegrationResponse struct {
	// The ID of the deleted Email integration       
	ID                                        string `json:"id"`
}

type FetchEmailIntegrationParams struct {
	// The ID of the Email integration to retrieve       
	EmailIntegrationID                            string `json:"emailIntegrationId"`
}

// Blueprint properties
type FetchEmailIntegrationResponse struct {
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

type SetupEmailIntegrationParams struct {
	// The ID of the Email integration       
	EmailIntegrationID                string `json:"emailIntegrationId"`
}

type SetupEmailIntegrationResponse struct {
	// The ID of the Email Integration       
	ID                                string `json:"id"`
}

type UpdateEmailIntegrationParams struct {
	// The ID of the Email integration       
	EmailIntegrationID                string `json:"emailIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateEmailIntegrationRequest struct {
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

type UpdateEmailIntegrationResponse struct {
	// The ID of the Email Integration       
	ID                                string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateEmailIntegrationRequest struct {
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

type CreateEmailIntegrationResponse struct {
	// The ID of the Email Integration       
	ID                                string `json:"id"`
}

type ListEmailIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListEmailIntegrationsResponse struct {
	Items []ListEmailIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListEmailIntegrationsResponseItem struct {
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

type DeleteExtractIntegrationParams struct {
	// The ID of the Extract integration       
	ExtractIntegrationID                string `json:"extractIntegrationId"`
}

type DeleteExtractIntegrationResponse struct {
	// The ID of the deleted Extract integration       
	ID                                          string `json:"id"`
}

type FetchExtractIntegrationParams struct {
	// The ID of the Extract integration to retrieve       
	ExtractIntegrationID                            string `json:"extractIntegrationId"`
}

// Blueprint properties
type FetchExtractIntegrationResponse struct {
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

type UpdateExtractIntegrationParams struct {
	// The ID of the Extract integration       
	ExtractIntegrationID                string `json:"extractIntegrationId"`
}

// Blueprint properties
type UpdateExtractIntegrationRequest struct {
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

type UpdateExtractIntegrationResponse struct {
	// The ID of the Extract Integration       
	ID                                  string `json:"id"`
}

// Blueprint properties
type CreateExtractIntegrationRequest struct {
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

type CreateExtractIntegrationResponse struct {
	// The ID of the Extract Integration       
	ID                                  string `json:"id"`
}

type ListExtractIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListExtractIntegrationsResponse struct {
	Items []ListExtractIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListExtractIntegrationsResponseItem struct {
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

type DeleteInstagramIntegrationParams struct {
	// The ID of the Instagram integration       
	InstagramIntegrationID                string `json:"instagramIntegrationId"`
}

type DeleteInstagramIntegrationResponse struct {
	// The ID of the deleted Instagram integration       
	ID                                            string `json:"id"`
}

type FetchInstagramIntegrationParams struct {
	// The ID of the Instagram integration to retrieve       
	InstagramIntegrationID                            string `json:"instagramIntegrationId"`
}

// Blueprint properties
type FetchInstagramIntegrationResponse struct {
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

type SetupInstagramIntegrationParams struct {
	// The ID of the Instagram integration       
	InstagramIntegrationID                string `json:"instagramIntegrationId"`
}

type SetupInstagramIntegrationResponse struct {
	// The ID of the Instagram Integration       
	ID                                    string `json:"id"`
}

type UpdateInstagramIntegrationParams struct {
	// The ID of the Instagram integration       
	InstagramIntegrationID                string `json:"instagramIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateInstagramIntegrationRequest struct {
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

type UpdateInstagramIntegrationResponse struct {
	// The ID of the Instagram Integration       
	ID                                    string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateInstagramIntegrationRequest struct {
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

type CreateInstagramIntegrationResponse struct {
	// The ID of the Instagram Integration       
	ID                                    string `json:"id"`
}

type ListInstagramIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListInstagramIntegrationsResponse struct {
	Items []ListInstagramIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListInstagramIntegrationsResponseItem struct {
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

type DeleteMCPServerIntegrationParams struct {
	// The ID of the McpServer integration       
	McpserverIntegrationID                string `json:"mcpserverIntegrationId"`
}

type DeleteMCPServerIntegrationResponse struct {
	// The ID of the deleted McpServer integration       
	ID                                            string `json:"id"`
}

type FetchMCPServerIntegrationParams struct {
	// The ID of the McpServer integration to retrieve       
	McpserverIntegrationID                            string `json:"mcpserverIntegrationId"`
}

// Blueprint properties
type FetchMCPServerIntegrationResponse struct {
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

type UpdateMCPServerIntegrationParams struct {
	// The ID of the McpServer integration       
	McpserverIntegrationID                string `json:"mcpserverIntegrationId"`
}

// Blueprint properties
type UpdateMCPServerIntegrationRequest struct {
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

type UpdateMCPServerIntegrationResponse struct {
	// The ID of the McpServer Integration       
	ID                                    string `json:"id"`
}

// Blueprint properties
type CreateMCPServerIntegrationRequest struct {
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

type CreateMCPServerIntegrationResponse struct {
	// The ID of the McpServer Integration       
	ID                                    string `json:"id"`
}

type ListMCPServerIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListMCPServerIntegrationsResponse struct {
	Items []ListMCPServerIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListMCPServerIntegrationsResponseItem struct {
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

type DeleteMessengerIntegrationParams struct {
	// The ID of the Messenger integration       
	MessengerIntegrationID                string `json:"messengerIntegrationId"`
}

type DeleteMessengerIntegrationResponse struct {
	// The ID of the deleted Messenger integration       
	ID                                            string `json:"id"`
}

type FetchMessengerIntegrationParams struct {
	// The ID of the Messenger integration to retrieve       
	MessengerIntegrationID                            string `json:"messengerIntegrationId"`
}

// Blueprint properties
type FetchMessengerIntegrationResponse struct {
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

type SetupMessengerIntegrationParams struct {
	// The ID of the Messenger integration       
	MessengerIntegrationID                string `json:"messengerIntegrationId"`
}

type SetupMessengerIntegrationResponse struct {
	// The ID of the Messenger Integration       
	ID                                    string `json:"id"`
}

type UpdateMessengerIntegrationParams struct {
	// The ID of the Messenger integration       
	MessengerIntegrationID                string `json:"messengerIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateMessengerIntegrationRequest struct {
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

type UpdateMessengerIntegrationResponse struct {
	// The ID of the Messenger Integration       
	ID                                    string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateMessengerIntegrationRequest struct {
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

type CreateMessengerIntegrationResponse struct {
	// The ID of the Messenger Integration       
	ID                                    string `json:"id"`
}

type ListMessengerIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListMessengerIntegrationsResponse struct {
	Items []ListMessengerIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListMessengerIntegrationsResponseItem struct {
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

type DeleteNotionIntegrationParams struct {
	// The ID of the Notion integration       
	NotionIntegrationID                string `json:"notionIntegrationId"`
}

type DeleteNotionIntegrationResponse struct {
	// The ID of the deleted Notion integration       
	ID                                         string `json:"id"`
}

type FetchNotionIntegrationParams struct {
	// The ID of the Notion integration to retrieve       
	NotionIntegrationID                            string `json:"notionIntegrationId"`
}

// Blueprint properties
type FetchNotionIntegrationResponse struct {
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

type SyncNotionIntegrationParams struct {
	// The ID of the Notion integration       
	NotionIntegrationID                string `json:"notionIntegrationId"`
}

type SyncNotionIntegrationResponse struct {
	// The ID of the synced Notion integration       
	ID                                        string `json:"id"`
}

type UpdateNotionIntegrationParams struct {
	// The ID of the Notion integration       
	NotionIntegrationID                string `json:"notionIntegrationId"`
}

// Blueprint properties
type UpdateNotionIntegrationRequest struct {
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

type UpdateNotionIntegrationResponse struct {
	// The ID of the Notion Integration       
	ID                                 string `json:"id"`
}

// Blueprint properties
type CreateNotionIntegrationRequest struct {
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

type CreateNotionIntegrationResponse struct {
	// The ID of the Notion Integration       
	ID                                 string `json:"id"`
}

type ListNotionIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListNotionIntegrationsResponse struct {
	Items []ListNotionIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListNotionIntegrationsResponseItem struct {
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

type DeleteSitemapIntegrationParams struct {
	// The ID of the Sitemap integration       
	SitemapIntegrationID                string `json:"sitemapIntegrationId"`
}

type DeleteSitemapIntegrationResponse struct {
	// The ID of the deleted Sitemap integration       
	ID                                          string `json:"id"`
}

type FetchSitemapIntegrationParams struct {
	// The ID of the Sitemap integration to retrieve       
	SitemapIntegrationID                            string `json:"sitemapIntegrationId"`
}

// Blueprint properties
type FetchSitemapIntegrationResponse struct {
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

type SyncSitemapIntegrationParams struct {
	// The ID of the Sitemap integration       
	SitemapIntegrationID                string `json:"sitemapIntegrationId"`
}

type SyncSitemapIntegrationResponse struct {
	// The ID of the Sitemap Integration       
	ID                                  string `json:"id"`
}

type UpdateSitemapIntegrationParams struct {
	// The ID of the Sitemap integration       
	SitemapIntegrationID                string `json:"sitemapIntegrationId"`
}

// Blueprint properties
type UpdateSitemapIntegrationRequest struct {
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

type UpdateSitemapIntegrationResponse struct {
	// The ID of the Sitemap Integration       
	ID                                  string `json:"id"`
}

// Blueprint properties
type CreateSitemapIntegrationRequest struct {
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

type CreateSitemapIntegrationResponse struct {
	// The ID of the Sitemap Integration       
	ID                                  string `json:"id"`
}

type ListSitemapIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListSitemapIntegrationsResponse struct {
	Items []ListSitemapIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListSitemapIntegrationsResponseItem struct {
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

type DeleteSlackIntegrationParams struct {
	// The ID of the Slack integration       
	SlackIntegrationID                string `json:"slackIntegrationId"`
}

type DeleteSlackIntegrationResponse struct {
	// The ID of the deleted Slack integration       
	ID                                        string `json:"id"`
}

type FetchSlackIntegrationParams struct {
	// The ID of the Slack integration to retrieve       
	SlackIntegrationID                            string `json:"slackIntegrationId"`
}

// Blueprint properties
type FetchSlackIntegrationResponse struct {
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

type SetupSlackIntegrationParams struct {
	// The ID of the Slack integration       
	SlackIntegrationID                string `json:"slackIntegrationId"`
}

type SetupSlackIntegrationResponse struct {
	// The ID of the setup Slack integration       
	ID                                      string `json:"id"`
}

type UpdateSlackIntegrationParams struct {
	// The ID of the Slack integration       
	SlackIntegrationID                string `json:"slackIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateSlackIntegrationRequest struct {
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

type UpdateSlackIntegrationResponse struct {
	// The ID of the Slack Integration       
	ID                                string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateSlackIntegrationRequest struct {
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

type CreateSlackIntegrationResponse struct {
	// The ID of the Slack Integration       
	ID                                string `json:"id"`
}

type ListSlackIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListSlackIntegrationsResponse struct {
	Items []ListSlackIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListSlackIntegrationsResponseItem struct {
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

type DeleteSupportIntegrationParams struct {
	// The ID of the Support integration       
	SupportIntegrationID                string `json:"supportIntegrationId"`
}

type DeleteSupportIntegrationResponse struct {
	// The ID of the deleted Support integration       
	ID                                          string `json:"id"`
}

type FetchSupportIntegrationParams struct {
	// The ID of the Support integration to retrieve       
	SupportIntegrationID                            string `json:"supportIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type FetchSupportIntegrationResponse struct {
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

type UpdateSupportIntegrationParams struct {
	// The ID of the Support integration       
	SupportIntegrationID                string `json:"supportIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateSupportIntegrationRequest struct {
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

type UpdateSupportIntegrationResponse struct {
	// The ID of the Support Integration       
	ID                                  string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateSupportIntegrationRequest struct {
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

type CreateSupportIntegrationResponse struct {
	// The ID of the Support Integration       
	ID                                  string `json:"id"`
}

type ListSupportIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListSupportIntegrationsResponse struct {
	Items []ListSupportIntegrationsResponseItem `json:"items"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type ListSupportIntegrationsResponseItem struct {
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

type DeleteTelegramIntegrationParams struct {
	// The ID of the Telegram integration       
	TelegramIntegrationID                string `json:"telegramIntegrationId"`
}

type DeleteTelegramIntegrationResponse struct {
	// The ID of the deleted Telegram integration       
	ID                                           string `json:"id"`
}

type FetchTelegramIntegrationParams struct {
	// The ID of the Telegram integration to retrieve       
	TelegramIntegrationID                            string `json:"telegramIntegrationId"`
}

// Blueprint properties
type FetchTelegramIntegrationResponse struct {
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

type SetupTelegramIntegrationParams struct {
	// The ID of the Telegram integration       
	TelegramIntegrationID                string `json:"telegramIntegrationId"`
}

type SetupTelegramIntegrationResponse struct {
	// The ID of the Telegram Integration       
	ID                                   string `json:"id"`
}

type UpdateTelegramIntegrationParams struct {
	// The ID of the Telegram integration       
	TelegramIntegrationID                string `json:"telegramIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateTelegramIntegrationRequest struct {
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

type UpdateTelegramIntegrationResponse struct {
	// The ID of the Telegram Integration       
	ID                                   string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateTelegramIntegrationRequest struct {
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

type CreateTelegramIntegrationResponse struct {
	// The ID of the Telegram Integration       
	ID                                   string `json:"id"`
}

type ListTelegramIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListTelegramIntegrationsResponse struct {
	Items []ListTelegramIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListTelegramIntegrationsResponseItem struct {
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

type DeleteTriggerIntegrationParams struct {
	// The ID of the Trigger integration       
	TriggerIntegrationID                string `json:"triggerIntegrationId"`
}

type DeleteTriggerIntegrationResponse struct {
	// The ID of the deleted Trigger integration       
	ID                                          string `json:"id"`
}

type FetchTriggerIntegrationParams struct {
	// The ID of the Trigger integration to retrieve       
	TriggerIntegrationID                            string `json:"triggerIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type FetchTriggerIntegrationResponse struct {
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

type InvokeTriggerIntegrationParams struct {
	// The ID of the Trigger integration       
	TriggerIntegrationID                string `json:"triggerIntegrationId"`
}

type InvokeTriggerIntegrationResponse struct {
	// The ID of the trigged Trigger integration       
	ID                                          string `json:"id"`
}

type SetupTriggerIntegrationParams struct {
	// The ID of the Trigger integration       
	TriggerIntegrationID                string `json:"triggerIntegrationId"`
}

type SetupTriggerIntegrationResponse struct {
	// The ID of the Trigger Integration       
	ID                                  string `json:"id"`
}

type UpdateTriggerIntegrationParams struct {
	// The ID of the Trigger integration       
	TriggerIntegrationID                string `json:"triggerIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateTriggerIntegrationRequest struct {
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

type UpdateTriggerIntegrationResponse struct {
	// The ID of the Trigger Integration       
	ID                                  string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateTriggerIntegrationRequest struct {
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

type CreateTriggerIntegrationResponse struct {
	// The ID of the Trigger Integration       
	ID                                  string `json:"id"`
}

type ListTriggerIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListTriggerIntegrationsResponse struct {
	Items []ListTriggerIntegrationsResponseItem `json:"items"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type ListTriggerIntegrationsResponseItem struct {
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

type DeleteTwilioIntegrationParams struct {
	// The ID of the Twilio integration       
	TwilioIntegrationID                string `json:"twilioIntegrationId"`
}

type DeleteTwilioIntegrationResponse struct {
	// The ID of the deleted Twilio integration       
	ID                                         string `json:"id"`
}

type FetchTwilioIntegrationParams struct {
	// The ID of the Twilio integration to retrieve       
	TwilioIntegrationID                            string `json:"twilioIntegrationId"`
}

// Blueprint properties
type FetchTwilioIntegrationResponse struct {
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

type SetupTwilioIntegrationParams struct {
	// The ID of the Twilio integration       
	TwilioIntegrationID                string `json:"twilioIntegrationId"`
}

type SetupTwilioIntegrationResponse struct {
	// The ID of the Twilio Integration       
	ID                                 string `json:"id"`
}

type UpdateTwilioIntegrationParams struct {
	// The ID of the Twilio integration       
	TwilioIntegrationID                string `json:"twilioIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateTwilioIntegrationRequest struct {
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

type UpdateTwilioIntegrationResponse struct {
	// The ID of the Twilio Integration       
	ID                                 string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateTwilioIntegrationRequest struct {
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

type CreateTwilioIntegrationResponse struct {
	// The ID of the Twilio Integration       
	ID                                 string `json:"id"`
}

type ListTwilioIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListTwilioIntegrationsResponse struct {
	Items []ListTwilioIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListTwilioIntegrationsResponseItem struct {
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

type DeleteWhatsAppIntegrationParams struct {
	// The ID of the WhatsApp integration       
	WhatsappIntegrationID                string `json:"whatsappIntegrationId"`
}

type DeleteWhatsAppIntegrationResponse struct {
	// The ID of the deleted WhatsApp integration       
	ID                                           string `json:"id"`
}

type FetchWhatsAppIntegrationParams struct {
	// The ID of the WhatsApp integration to retrieve       
	WhatsappIntegrationID                            string `json:"whatsappIntegrationId"`
}

// Blueprint properties
type FetchWhatsAppIntegrationResponse struct {
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

type SetupWhatsAppIntegrationParams struct {
	// The ID of the WhatsApp integration       
	WhatsappIntegrationID                string `json:"whatsappIntegrationId"`
}

type SetupWhatsAppIntegrationResponse struct {
	// The ID of the WhatsApp Integration       
	ID                                   string `json:"id"`
}

type UpdateWhatsAppIntegrationParams struct {
	// The ID of the WhatsApp integration       
	WhatsappIntegrationID                string `json:"whatsappIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateWhatsAppIntegrationRequest struct {
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

type UpdateWhatsAppIntegrationResponse struct {
	// The ID of the WhatsApp Integration       
	ID                                   string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateWhatsAppIntegrationRequest struct {
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

type CreateWhatsAppIntegrationResponse struct {
	// The ID of the WhatsApp Integration       
	ID                                   string `json:"id"`
}

type ListWhatsAppIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListWhatsAppIntegrationsResponse struct {
	Items []ListWhatsAppIntegrationsResponseItem `json:"items"`
}

// Blueprint properties
type ListWhatsAppIntegrationsResponseItem struct {
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

type DeleteWidgetIntegrationParams struct {
	// The ID of the Widget integration       
	WidgetIntegrationID                string `json:"widgetIntegrationId"`
}

type DeleteWidgetIntegrationResponse struct {
	// The ID of the deleted Widget integration       
	ID                                         string `json:"id"`
}

type FetchWidgetIntegrationParams struct {
	// The ID of the Widget integration to retrieve       
	WidgetIntegrationID                            string `json:"widgetIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type FetchWidgetIntegrationResponse struct {
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

type SetupWidgetIntegrationParams struct {
	// The ID of the Widget integration       
	WidgetIntegrationID                string `json:"widgetIntegrationId"`
}

type SetupWidgetIntegrationResponse struct {
	// The ID of the Widget integration       
	ID                                 string `json:"id"`
}

type UpdateWidgetIntegrationParams struct {
	// The ID of the Widget integration       
	WidgetIntegrationID                string `json:"widgetIntegrationId"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type UpdateWidgetIntegrationRequest struct {
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

type UpdateWidgetIntegrationResponse struct {
	// The ID of the Widget Integration       
	ID                                 string `json:"id"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type CreateWidgetIntegrationRequest struct {
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

type CreateWidgetIntegrationResponse struct {
	// The ID of the Widget Integration       
	ID                                 string `json:"id"`
}

type ListWidgetIntegrationsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListWidgetIntegrationsResponse struct {
	Items []ListWidgetIntegrationsResponseItem `json:"items"`
}

// A bot configuration that can be applied without a dedicated bot instance.
type ListWidgetIntegrationsResponseItem struct {
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

type GenerateMagicFromPromptParams struct {
	// The ID of the prompt to use for generation       
	PromptID                                     string `json:"promptId"`
}

type GenerateMagicFromPromptRequest struct {
	// Optional language model to use for generation                       
	Model                                           *string                `json:"model,omitempty"`
	// Additional properties to pass to the prompt                         
	Props                                           map[string]interface{} `json:"props,omitempty"`
	// The text to use as input                                            
	Text                                            string                 `json:"text"`
}

type GenerateMagicFromPromptResponse struct {
	// The input text                                        
	Text                string                               `json:"text"`
	// Usage information                                     
	Usage               GenerateMagicFromPromptResponseUsage `json:"usage"`
}

// Usage information
type GenerateMagicFromPromptResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

type ListMagicPromptsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListMagicPromptsResponse struct {
	Items []ListMagicPromptsResponseItem `json:"items"`
}

// Instance list properties
type ListMagicPromptsResponseItem struct {
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

type DeleteMemoryParams struct {
	// The ID of the memory to delete       
	MemoryID                         string `json:"memoryId"`
}

type DeleteMemoryResponse struct {
	// The ID of the deleted memory       
	ID                             string `json:"id"`
}

type FetchMemoryParams struct {
	// The ID of the memory to retrieve       
	MemoryID                           string `json:"memoryId"`
}

// Instance list properties
type FetchMemoryResponse struct {
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

type UpdateMemoryParams struct {
	MemoryID string `json:"memoryId"`
}

// Instance crud properties
type UpdateMemoryRequest struct {
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

type UpdateMemoryResponse struct {
	// The ID of the updated memory       
	ID                             string `json:"id"`
}

// Instance crud properties
type CreateMemoryRequest struct {
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

type CreateMemoryResponse struct {
	// The ID of the created memory       
	ID                             string `json:"id"`
}

type ExportMemoriesParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ExportMemoriesResponse struct {
	Items []ExportMemoriesResponseItem `json:"items"`
}

// Instance list properties
type ExportMemoriesResponseItem struct {
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

type ListMemoriesParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListMemoriesResponse struct {
	Items []ListMemoriesResponseItem `json:"items"`
}

// Instance list properties
type ListMemoriesResponseItem struct {
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

type SearchMemoryRequest struct {
	// The ID of the bot to filter memories by            
	BotID                                         *string `json:"botId,omitempty"`
	// The ID of the contact to filter memories by        
	ContactID                                     *string `json:"contactId,omitempty"`
	// The keyword/phrase to search for                   
	Search                                        string  `json:"search"`
}

type SearchMemoryResponse struct {
	// An array of memories matching the search query                           
	Items                                            []SearchMemoryResponseItem `json:"items"`
}

type SearchMemoryResponseItem struct {
	ID   string                 `json:"id"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Text string                 `json:"text"`
}

type DeletePartnerUserParams struct {
	// The ID of the user to delete       
	UserID                         string `json:"userId"`
}

type DeletePartnerUserResponse struct {
	// The ID of the deleted user       
	ID                           string `json:"id"`
}

type FetchPartnerUserParams struct {
	// The ID of the partner user to retrieve       
	UserID                                   string `json:"userId"`
}

// Instance list properties
type FetchPartnerUserResponse struct {
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
	Limits                                             *FetchPartnerUserResponseLimits `json:"limits,omitempty"`
	// Meta data information                                                           
	Meta                                               map[string]interface{}          `json:"meta,omitempty"`
	// The associated name                                                             
	Name                                               *string                         `json:"name,omitempty"`
	// The timestamp (ms) when the instance was updated                                
	UpdatedAt                                          float64                         `json:"updatedAt"`
}

// Limits information
type FetchPartnerUserResponseLimits struct {
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

type DeletePartnerUserTokenParams struct {
	// The ID of the user token to delete       
	TokenID                              string `json:"tokenId"`
	// The ID of the user                       
	UserID                               string `json:"userId"`
}

type DeletePartnerUserTokenResponse struct {
	// The ID of the deleted user token       
	ID                                 string `json:"id"`
}

type CreatePartnerUserTokenParams struct {
	// The ID of the user       
	UserID               string `json:"userId"`
}

type CreatePartnerUserTokenResponse struct {
	// The timestamp for when the user token was created (in milliseconds)        
	CreatedAt                                                             float64 `json:"createdAt"`
	// The ID of the created user token                                           
	ID                                                                    string  `json:"id"`
	// The token of the created user token                                        
	Token                                                                 string  `json:"token"`
}

type ListPartnerUserTokensParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
	// The ID of the user                      
	UserID                             string  `json:"userId"`
}

type ListPartnerUserTokensResponse struct {
	Items []ListPartnerUserTokensResponseItem `json:"items"`
}

// Instance list properties
type ListPartnerUserTokensResponseItem struct {
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

type UpdatePartnerUserParams struct {
	// The ID of the partner user       
	UserID                       string `json:"userId"`
}

// Instance crud properties
type UpdatePartnerUserRequest struct {
	// The associated description                                   
	Description                     *string                         `json:"description,omitempty"`
	// The email of the partner user                                
	Email                           *string                         `json:"email,omitempty"`
	// The image of the partner user                                
	Image                           *string                         `json:"image,omitempty"`
	// Limits information                                           
	Limits                          *UpdatePartnerUserRequestLimits `json:"limits,omitempty"`
	// Meta data information                                        
	Meta                            map[string]interface{}          `json:"meta,omitempty"`
	// The associated name                                          
	Name                            *string                         `json:"name,omitempty"`
}

// Limits information
type UpdatePartnerUserRequestLimits struct {
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

type UpdatePartnerUserResponse struct {
	// The ID of the updated partner user       
	ID                                   string `json:"id"`
}

// Instance crud properties
type CreatePartnerUserRequest struct {
	// The associated description                                   
	Description                     *string                         `json:"description,omitempty"`
	// The email of the partner user                                
	Email                           *string                         `json:"email,omitempty"`
	// The image of the partner user                                
	Image                           *string                         `json:"image,omitempty"`
	// Limits information                                           
	Limits                          *CreatePartnerUserRequestLimits `json:"limits,omitempty"`
	// Meta data information                                        
	Meta                            map[string]interface{}          `json:"meta,omitempty"`
	// The associated name                                          
	Name                            *string                         `json:"name,omitempty"`
}

// Limits information
type CreatePartnerUserRequestLimits struct {
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

type CreatePartnerUserResponse struct {
	// The ID of the created user       
	ID                           string `json:"id"`
}

type ListPartnerUsersParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPartnerUsersResponse struct {
	Items []ListPartnerUsersResponseItem `json:"items"`
}

// Instance list properties
type ListPartnerUsersResponseItem struct {
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

type ListPlatformAbilitiesParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPlatformAbilitiesResponse struct {
	Items []ListPlatformAbilitiesResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformAbilitiesResponseItem struct {
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

type ListPlatformActionsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPlatformActionsResponse struct {
	Items []ListPlatformActionsResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformActionsResponseItem struct {
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

type FetchPlatformDocParams struct {
	// The ID of the doc to fetch (e.g., "datasets", "skillsets")       
	DocID                                                        string `json:"docId"`
}

// Instance list properties
type FetchPlatformDocResponse struct {
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

type ListPlatformDocsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPlatformDocsResponse struct {
	Items []ListPlatformDocsResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformDocsResponseItem struct {
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

type SearchPlatformDocsRequest struct {
	// The search query to find relevant docs                            
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type SearchPlatformDocsResponse struct {
	Items []SearchPlatformDocsResponseItem `json:"items"`
}

// Instance list properties
type SearchPlatformDocsResponseItem struct {
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

type ClonePlatformExampleParams struct {
	// The ID (slug) of the example to clone       
	ExampleID                               string `json:"exampleId"`
}

type ClonePlatformExampleResponse struct {
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

type FetchPlatformExampleParams struct {
	// The ID (slug) of the example       
	ExampleID                      string `json:"exampleId"`
}

type FetchPlatformExampleResponse struct {
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
	Type                                            FetchPlatformExampleResponseType `json:"type"`
	// The last update timestamp                                                     
	UpdatedAt                                       *float64                         `json:"updatedAt,omitempty"`
}

type ListPlatformExamplesParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPlatformExamplesResponse struct {
	Items []ListPlatformExamplesResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformExamplesResponseItem struct {
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
	Type                                               FetchPlatformExampleResponseType `json:"type"`
	// The timestamp (ms) when the instance was updated                                 
	UpdatedAt                                          float64                          `json:"updatedAt"`
}

type SearchPlatformExamplesRequest struct {
	// The search query to find relevant examples                        
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type SearchPlatformExamplesResponse struct {
	Items []SearchPlatformExamplesResponseItem `json:"items"`
}

// Instance list properties
type SearchPlatformExamplesResponseItem struct {
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
	Type                                               FetchPlatformExampleResponseType `json:"type"`
	// The timestamp (ms) when the instance was updated                                 
	UpdatedAt                                          float64                          `json:"updatedAt"`
}

type FetchPlatformGuideParams struct {
	// The ID of the guide to fetch       
	GuideID                        string `json:"guideId"`
}

// Instance list properties
type FetchPlatformGuideResponse struct {
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

type ListPlatformGuidesParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPlatformGuidesResponse struct {
	Items []ListPlatformGuidesResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformGuidesResponseItem struct {
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

type SearchPlatformGuidesRequest struct {
	// The search query to find relevant guides                          
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type SearchPlatformGuidesResponse struct {
	Items []SearchPlatformGuidesResponseItem `json:"items"`
}

// Instance list properties
type SearchPlatformGuidesResponseItem struct {
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

type FetchPlatformManualParams struct {
	// The ID of the manual to fetch (e.g., "datasets", "skillsets")       
	ManualID                                                        string `json:"manualId"`
}

// Instance list properties
type FetchPlatformManualResponse struct {
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

type ListPlatformManualsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPlatformManualsResponse struct {
	Items []ListPlatformManualsResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformManualsResponseItem struct {
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

type SearchPlatformManualsRequest struct {
	// The search query to find relevant manuals                         
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type SearchPlatformManualsResponse struct {
	Items []SearchPlatformManualsResponseItem `json:"items"`
}

// Instance list properties
type SearchPlatformManualsResponseItem struct {
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

type ListPlatformModelsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPlatformModelsResponse struct {
	Items []ListPlatformModelsResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformModelsResponseItem struct {
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

type ListPlatformSecretsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPlatformSecretsResponse struct {
	Items []ListPlatformSecretsResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformSecretsResponseItem struct {
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

type FetchPlatformTutorialParams struct {
	// The ID of the tutorial to fetch (e.g., "how-to-get-started-with-chatbotkit")       
	TutorialID                                                                     string `json:"tutorialId"`
}

// Instance list properties
type FetchPlatformTutorialResponse struct {
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

type ListPlatformTutorialsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPlatformTutorialsResponse struct {
	Items []ListPlatformTutorialsResponseItem `json:"items"`
}

// Instance list properties
type ListPlatformTutorialsResponseItem struct {
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

type SearchPlatformTutorialsRequest struct {
	// The search query to find relevant tutorials                       
	Search                                                        string `json:"search"`
	// The maximum number of results to return (1-100, default 10)       
	Take                                                          *int64 `json:"take,omitempty"`
}

type SearchPlatformTutorialsResponse struct {
	Items []SearchPlatformTutorialsResponseItem `json:"items"`
}

// Instance list properties
type SearchPlatformTutorialsResponseItem struct {
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

type DeletePolicyParams struct {
	// The ID of the policy to delete       
	ID                               string `json:"id"`
}

type DeletePolicyResponse struct {
	// The ID of the deleted policy       
	ID                             string `json:"id"`
}

type FetchPolicyParams struct {
	// The ID of the policy to fetch       
	ID                              string `json:"id"`
}

// Blueprint properties
type FetchPolicyResponse struct {
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

type UpdatePolicyParams struct {
	// The ID of the policy to update       
	ID                               string `json:"id"`
}

// Blueprint properties
type UpdatePolicyRequest struct {
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

type UpdatePolicyResponse struct {
	// The ID of the updated policy       
	ID                             string `json:"id"`
}

// Blueprint properties
type CreatePolicyRequest struct {
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

type CreatePolicyResponse struct {
	// The ID of the created policy       
	ID                             string `json:"id"`
}

type ListPoliciesParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPoliciesResponse struct {
	Items []ListPoliciesResponseItem `json:"items"`
}

// Blueprint properties
type ListPoliciesResponseItem struct {
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

type DeletePortalParams struct {
	// The ID of the portal to delete       
	PortalID                         string `json:"portalId"`
}

type DeletePortalResponse struct {
	// The ID of the deleted portal       
	ID                             string `json:"id"`
}

type FetchPortalParams struct {
	// The ID of the portal to retrieve       
	PortalID                           string `json:"portalId"`
}

// Blueprint properties
type FetchPortalResponse struct {
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

type UpdatePortalParams struct {
	PortalID string `json:"portalId"`
}

// Blueprint properties
type UpdatePortalRequest struct {
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

type UpdatePortalResponse struct {
	// The ID of the updated portal       
	ID                             string `json:"id"`
}

// Blueprint properties
type CreatePortalRequest struct {
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

type CreatePortalResponse struct {
	// The ID of the created portal       
	ID                             string `json:"id"`
}

type ListPortalsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListPortalsResponse struct {
	Items []ListPortalsResponseItem `json:"items"`
}

// Blueprint properties
type ListPortalsResponseItem struct {
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

type AuthenticateSecretParams struct {
	// The ID of the secret to authenticate       
	SecretID                               string `json:"secretId"`
}

type AuthenticateSecretResponse struct {
	// The ID of the secret to authenticate       
	ID                                     string `json:"id"`
	// The URL to authenticate the secret         
	URL                                    string `json:"url"`
}

type DeleteSecretParams struct {
	// The ID of the secret to delete       
	SecretID                         string `json:"secretId"`
}

type DeleteSecretResponse struct {
	// The ID of the deleted secret       
	ID                             string `json:"id"`
}

type FetchSecretParams struct {
	// The ID of the secret to retrieve       
	SecretID                           string `json:"secretId"`
}

// Blueprint properties
type FetchSecretResponse struct {
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

type RevokeSecretParams struct {
	SecretID string `json:"secretId"`
}

type RevokeSecretResponse struct {
	// The ID of the revoked secret       
	ID                             string `json:"id"`
}

type UpdateSecretParams struct {
	SecretID string `json:"secretId"`
}

// Blueprint properties
type UpdateSecretRequest struct {
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

type UpdateSecretResponse struct {
	// The ID of the updated secret       
	ID                             string `json:"id"`
}

type VerifySecretParams struct {
	// The ID of the secret to be verified       
	SecretID                              string `json:"secretId"`
}

type VerifySecretResponse struct {
	Action                          *VerifySecretResponseAction `json:"action,omitempty"`
	// The ID of the verified secret                            
	ID                              string                      `json:"id"`
	// The status of the secret                                 
	Status                          Status                      `json:"status"`
}

// The action to take next
type VerifySecretResponseAction struct {
	// The type of action to take                   
	Type                                 ActionType `json:"type"`
	// The URL to authenticate the secret           
	URL                                  string     `json:"url"`
}

// Blueprint properties
type CreateSecretRequest struct {
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

type CreateSecretResponse struct {
	// The ID of the created secret       
	ID                             string `json:"id"`
}

type ListSecretsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListSecretsResponse struct {
	Items []ListSecretsResponseItem `json:"items"`
}

// Blueprint properties
type ListSecretsResponseItem struct {
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

type DeleteSkillsetAbilityParams struct {
	// The ID of the ability to delete       
	AbilityID                         string `json:"abilityId"`
	// The ID of the skillset                
	SkillsetID                        string `json:"skillsetId"`
}

type DeleteSkillsetAbilityResponse struct {
	// The ID of the deleted ability       
	ID                              string `json:"id"`
}

type ExecuteSkillsetAbilityParams struct {
	// The ID of the ability to execute                    
	AbilityID                                       string `json:"abilityId"`
	// The ID of the skillset containing the ability       
	SkillsetID                                      string `json:"skillsetId"`
}

type ExecuteSkillsetAbilityRequest struct {
	// The ID of the contact to associate with the execution                 
	ContactID                                                        *string `json:"contactId,omitempty"`
	// The input to process with the ability. This can be structured         
	// text such as JSON or YAML for precise parameter control, or           
	// unstructured natural language text. When unstructured text is         
	// provided, the system will automatically detect and extract the        
	// relevant parameters from the input.                                   
	Input                                                            *string `json:"input,omitempty"`
}

type ExecuteSkillsetAbilityResponse struct {
	// Error message if execution failed                                          
	Error                                 *string                                 `json:"error,omitempty"`
	// Messages generated during execution                                        
	Messages                              []ExecuteSkillsetAbilityResponseMessage `json:"messages,omitempty"`
	// The result of the ability execution                                        
	Result                                interface{}                             `json:"result"`
	// Usage information                                                          
	Usage                                 ExecuteSkillsetAbilityResponseUsage     `json:"usage"`
}

// A message in the conversation
type ExecuteSkillsetAbilityResponseMessage struct {
	// Meta data information                         
	Meta                      map[string]interface{} `json:"meta,omitempty"`
	// The text of the message                       
	Text                      string                 `json:"text"`
	// The type of the message                       
	Type                      MessageType            `json:"type"`
}

// Usage information
type ExecuteSkillsetAbilityResponseUsage struct {
	// The tokens used in this exchange        
	Token                              float64 `json:"token"`
}

type FetchSkillsetAbilityParams struct {
	// The ID of the ability to retrieve       
	AbilityID                           string `json:"abilityId"`
	// The ID of the skillset                  
	SkillsetID                          string `json:"skillsetId"`
}

// Blueprint properties
type FetchSkillsetAbilityResponse struct {
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

type UpdateSkillsetAbilityParams struct {
	AbilityID  string `json:"abilityId"`
	SkillsetID string `json:"skillsetId"`
}

// Blueprint properties
type UpdateSkillsetAbilityRequest struct {
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

type UpdateSkillsetAbilityResponse struct {
	// The ID of the updated ability       
	ID                              string `json:"id"`
}

type CreateSkillsetAbilityParams struct {
	SkillsetID string `json:"skillsetId"`
}

// Blueprint properties
type CreateSkillsetAbilityRequest struct {
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

type CreateSkillsetAbilityResponse struct {
	// The ID of the created ability       
	ID                              string `json:"id"`
}

type ExportSkillsetAbilitiesParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The ID of the skillset to export        
	SkillsetID                         string  `json:"skillsetId"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type ExportSkillsetAbilitiesResponse struct {
	Items []ExportSkillsetAbilitiesResponseItem `json:"items"`
}

// Blueprint properties
type ExportSkillsetAbilitiesResponseItem struct {
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

type ListSkillsetAbilitiesParams struct {
	// The cursor to use for pagination        
	Cursor                             *string `json:"cursor,omitempty"`
	// The order of the paginated items        
	Order                              *Order  `json:"order,omitempty"`
	// The ID of the skillset                  
	SkillsetID                         string  `json:"skillsetId"`
	// The number of items to retrieve         
	Take                               *int64  `json:"take,omitempty"`
}

type ListSkillsetAbilitiesResponse struct {
	Items []ListSkillsetAbilitiesResponseItem `json:"items"`
}

// Blueprint properties
type ListSkillsetAbilitiesResponseItem struct {
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

type DeleteSkillsetParams struct {
	// The ID of the skillset to delete       
	SkillsetID                         string `json:"skillsetId"`
}

type DeleteSkillsetResponse struct {
	// The ID of the deleted skillset       
	ID                               string `json:"id"`
}

type FetchSkillsetParams struct {
	// The ID of the skillset to retrieve       
	SkillsetID                           string `json:"skillsetId"`
}

// Blueprint properties
type FetchSkillsetResponse struct {
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

type UpdateSkillsetParams struct {
	SkillsetID string `json:"skillsetId"`
}

// Blueprint properties
type UpdateSkillsetRequest struct {
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

type UpdateSkillsetResponse struct {
	// The ID of the updated skillset       
	ID                               string `json:"id"`
}

// Blueprint properties
type CreateSkillsetRequest struct {
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

type CreateSkillsetResponse struct {
	// The ID of the created skillset       
	ID                               string `json:"id"`
}

type ListSkillsetsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListSkillsetsResponse struct {
	Items []ListSkillsetsResponseItem `json:"items"`
}

// Blueprint properties
type ListSkillsetsResponseItem struct {
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

type FetchSpaceParams struct {
	// The ID of the space to retrieve       
	SpaceID                           string `json:"spaceId"`
}

// Blueprint properties
type FetchSpaceResponse struct {
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

type UpdateSpaceParams struct {
	SpaceID string `json:"spaceId"`
}

// Blueprint properties
type UpdateSpaceRequest struct {
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

type UpdateSpaceResponse struct {
	// The ID of the updated space       
	ID                            string `json:"id"`
}

// Blueprint properties
type CreateSpaceRequest struct {
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

type CreateSpaceResponse struct {
	// The ID of the created space       
	ID                            string `json:"id"`
}

type ExportSpacesParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ExportSpacesResponse struct {
	Items []ExportSpacesResponseItem `json:"items"`
}

// Blueprint properties
type ExportSpacesResponseItem struct {
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

type ListSpacesParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListSpacesResponse struct {
	Items []ListSpacesResponseItem `json:"items"`
}

// Blueprint properties
type ListSpacesResponseItem struct {
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

type DeleteTaskParams struct {
	// The ID of the task to delete       
	TaskID                         string `json:"taskId"`
}

type DeleteTaskResponse struct {
	// The ID of the deleted task       
	ID                           string `json:"id"`
}

type FetchTaskParams struct {
	// The ID of the task to retrieve       
	TaskID                           string `json:"taskId"`
}

// Instance list properties
type FetchTaskResponse struct {
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

type TriggerTaskParams struct {
	TaskID string `json:"taskId"`
}

type TriggerTaskResponse struct {
	// The ID of the triggered task       
	ID                             string `json:"id"`
}

type UpdateTaskParams struct {
	TaskID string `json:"taskId"`
}

// Instance crud properties
type UpdateTaskRequest struct {
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

type UpdateTaskResponse struct {
	// The ID of the updated task       
	ID                           string `json:"id"`
}

// Instance crud properties
type CreateTaskRequest struct {
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

type CreateTaskResponse struct {
	// The ID of the created task       
	ID                           string `json:"id"`
}

type ExportTasksParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ExportTasksResponse struct {
	Items []ExportTasksResponseItem `json:"items"`
}

// Instance list properties
type ExportTasksResponseItem struct {
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

type ListTasksParams struct {
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

type ListTasksResponse struct {
	Items []ListTasksResponseItem `json:"items"`
}

// Instance list properties
type ListTasksResponseItem struct {
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

type ListTeamsParams struct {
	// The cursor to use for pagination                                         
	Cursor                                                    *string           `json:"cursor,omitempty"`
	// Key-value pairs to filter the partner users by metadata                  
	Meta                                                      map[string]string `json:"meta,omitempty"`
	// The order of the paginated items                                         
	Order                                                     *Order            `json:"order,omitempty"`
	// The number of items to retrieve                                          
	Take                                                      *int64            `json:"take,omitempty"`
}

type ListTeamsResponse struct {
	Items []ListTeamsResponseItem `json:"items"`
}

// Instance list properties
type ListTeamsResponseItem struct {
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

type FetchUsageResponse struct {
	// The number of conversations the user has created                           
	Conversations                                      float64                    `json:"conversations"`
	// Database usage information                                                 
	Database                                           FetchUsageResponseDatabase `json:"database"`
	// The number of messages the user has sent                                   
	Messages                                           float64                    `json:"messages"`
	// The number of tokens the user has used                                     
	Tokens                                             float64                    `json:"tokens"`
}

// Database usage information
type FetchUsageResponseDatabase struct {
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

type FetchUsageSeriesResponse struct {
	// The number of conversations the user has created                                  
	Conversations                                      []Conversation                    `json:"conversations"`
	// The number of messages the user has created                                       
	Messages                                           []FetchUsageSeriesResponseMessage `json:"messages"`
	// The number of tokens the user has used                                            
	Tokens                                             []TokenElement                    `json:"tokens"`
}

type Conversation struct {
	// The date of the data point                                 
	Date                                                  float64 `json:"date"`
	// The total number of conversations the user has used        
	Total                                                 float64 `json:"total"`
}

type FetchUsageSeriesResponseMessage struct {
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
type FetchPlatformExampleResponseType string

const (
	Blueprint   FetchPlatformExampleResponseType = "blueprint"
	Discord     FetchPlatformExampleResponseType = "discord"
	Email       FetchPlatformExampleResponseType = "email"
	Messenger   FetchPlatformExampleResponseType = "messenger"
	Project     FetchPlatformExampleResponseType = "project"
	Slack       FetchPlatformExampleResponseType = "slack"
	Telegram    FetchPlatformExampleResponseType = "telegram"
	Twilio      FetchPlatformExampleResponseType = "twilio"
	TypeTrigger FetchPlatformExampleResponseType = "trigger"
	Whatsapp    FetchPlatformExampleResponseType = "whatsapp"
	Widget      FetchPlatformExampleResponseType = "widget"
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

type UploadConversationAttachmentRequestFile struct {
	PurpleFile *PurpleFile
	String     *string
}

func (x *UploadConversationAttachmentRequestFile) UnmarshalJSON(data []byte) error {
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

func (x *UploadConversationAttachmentRequestFile) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleFile != nil, x.PurpleFile, false, nil, false, nil, false)
}

// The contact ID to associate with this conversation
type CompleteConversationRequestContactID struct {
	PurpleContactID *PurpleContactID
	String          *string
}

func (x *CompleteConversationRequestContactID) UnmarshalJSON(data []byte) error {
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

func (x *CompleteConversationRequestContactID) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleContactID != nil, x.PurpleContactID, false, nil, false, nil, false)
}

// The contact ID to associate with this conversation
type DispatchConversationRequestContactID struct {
	FluffyContactID *FluffyContactID
	String          *string
}

func (x *DispatchConversationRequestContactID) UnmarshalJSON(data []byte) error {
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

func (x *DispatchConversationRequestContactID) MarshalJSON() ([]byte, error) {
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

type UploadFileRequestFile struct {
	FluffyFile *FluffyFile
	String     *string
}

func (x *UploadFileRequestFile) UnmarshalJSON(data []byte) error {
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

func (x *UploadFileRequestFile) MarshalJSON() ([]byte, error) {
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
