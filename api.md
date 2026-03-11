# Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateMessageParams">CreateMessageParams</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#MessageItemParam">MessageItemParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#SystemPromptUnionParam">SystemPromptUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateMessageResponse">CreateMessageResponse</a>

Methods:

- <code title="post /messages">client.Messages.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#MessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#MessageNewParams">MessageNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateMessageResponse">CreateMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Interactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#GetInteractionAnalyticsResponse">GetInteractionAnalyticsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ListInteractionsResponse">ListInteractionsResponse</a>

Methods:

- <code title="get /interactions">client.Interactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#InteractionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#InteractionListParams">InteractionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ListInteractionsResponse">ListInteractionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /interactions/{interactionId}/analytics">client.Interactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#InteractionService.GetInteractionAnalytics">GetInteractionAnalytics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, interactionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#InteractionGetInteractionAnalyticsParams">InteractionGetInteractionAnalyticsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#GetInteractionAnalyticsResponse">GetInteractionAnalyticsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateUserParams">CreateUserParams</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UpdateUserParams">UpdateUserParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateUserResponse">CreateUserResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#GetUserAnalyticsResponse">GetUserAnalyticsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ListUsersResponse">ListUsersResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#Participant">Participant</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UpdateUserResponse">UpdateUserResponse</a>

Methods:

- <code title="post /users">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UserNewParams">UserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateUserResponse">CreateUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/{userId}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UserUpdateParams">UserUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UpdateUserResponse">UpdateUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UserListParams">UserListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ListUsersResponse">ListUsersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users/{userId}/analytics">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UserService.GetUserAnalytics">GetUserAnalytics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UserGetUserAnalyticsParams">UserGetUserAnalyticsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#GetUserAnalyticsResponse">GetUserAnalyticsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Ratings

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#LogRatingParams">LogRatingParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#LogRatingResponse">LogRatingResponse</a>

Methods:

- <code title="post /ratings">client.Ratings.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#RatingService.Log">Log</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#RatingLogParams">RatingLogParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#LogRatingResponse">LogRatingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Organizations

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateOrganizationParams">CreateOrganizationParams</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UpdateOrganizationParams">UpdateOrganizationParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateOrganizationResponse">CreateOrganizationResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#GetOrganizationAnalyticsResponse">GetOrganizationAnalyticsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ListOrganizationsResponse">ListOrganizationsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#TenantOrganization">TenantOrganization</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UpdateOrganizationResponse">UpdateOrganizationResponse</a>

Methods:

- <code title="post /organizations">client.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#OrganizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#OrganizationNewParams">OrganizationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateOrganizationResponse">CreateOrganizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /organizations/{organizationId}">client.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#OrganizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#OrganizationUpdateParams">OrganizationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UpdateOrganizationResponse">UpdateOrganizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /organizations">client.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#OrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#OrganizationListParams">OrganizationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ListOrganizationsResponse">ListOrganizationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /organizations/{organizationId}/analytics">client.Organizations.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#OrganizationService.GetOrganizationAnalytics">GetOrganizationAnalytics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#OrganizationGetOrganizationAnalyticsParams">OrganizationGetOrganizationAnalyticsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#GetOrganizationAnalyticsResponse">GetOrganizationAnalyticsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Prompts

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ComponentInputParam">ComponentInputParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ComponentUpdateParam">ComponentUpdateParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreatePromptParams">CreatePromptParams</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UpdatePromptParams">UpdatePromptParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreatePromptResponse">CreatePromptResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#DeletePromptResponse">DeletePromptResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#GetPromptResponse">GetPromptResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ListPromptsResponse">ListPromptsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#Prompt">Prompt</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#PromptComponent">PromptComponent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#SlimPrompt">SlimPrompt</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#SlimPromptComponent">SlimPromptComponent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UpdatePromptResponse">UpdatePromptResponse</a>

Methods:

- <code title="post /prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#PromptService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#PromptNewParams">PromptNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreatePromptResponse">CreatePromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /prompts/{id}">client.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#PromptService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#PromptUpdateParams">PromptUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#UpdatePromptResponse">UpdatePromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#PromptService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#PromptListParams">PromptListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#ListPromptsResponse">ListPromptsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /prompts/{id}">client.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#PromptService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#DeletePromptResponse">DeletePromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /prompts/{id}">client.Prompts.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#PromptService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#GetPromptResponse">GetPromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Events

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateEventParams">CreateEventParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateEventResponse">CreateEventResponse</a>

Methods:

- <code title="post /events">client.Events.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#EventService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#EventNewParams">EventNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/greenflash-public-api-go#CreateEventResponse">CreateEventResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
