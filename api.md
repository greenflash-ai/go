# Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateMessageParams">CreateMessageParams</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#MessageItemParam">MessageItemParam</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#SystemPromptUnionParam">SystemPromptUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateMessageResponse">CreateMessageResponse</a>

Methods:

- <code title="post /messages">client.Messages.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#MessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#MessageNewParams">MessageNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateMessageResponse">CreateMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Interactions

Response Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#GetInteractionAnalyticsResponse">GetInteractionAnalyticsResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ListInteractionsResponse">ListInteractionsResponse</a>

Methods:

- <code title="get /interactions">client.Interactions.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#InteractionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#InteractionListParams">InteractionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ListInteractionsResponse">ListInteractionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /interactions/{interactionId}/analytics">client.Interactions.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#InteractionService.GetInteractionAnalytics">GetInteractionAnalytics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, interactionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#InteractionGetInteractionAnalyticsParams">InteractionGetInteractionAnalyticsParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#GetInteractionAnalyticsResponse">GetInteractionAnalyticsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateUserParams">CreateUserParams</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UpdateUserParams">UpdateUserParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateUserResponse">CreateUserResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#GetUserAnalyticsResponse">GetUserAnalyticsResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ListUsersResponse">ListUsersResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#Participant">Participant</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UpdateUserResponse">UpdateUserResponse</a>

Methods:

- <code title="post /users">client.Users.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UserNewParams">UserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateUserResponse">CreateUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/{userId}">client.Users.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UserUpdateParams">UserUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UpdateUserResponse">UpdateUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users">client.Users.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UserListParams">UserListParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ListUsersResponse">ListUsersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users/{userId}/analytics">client.Users.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UserService.GetUserAnalytics">GetUserAnalytics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UserGetUserAnalyticsParams">UserGetUserAnalyticsParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#GetUserAnalyticsResponse">GetUserAnalyticsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Ratings

Params Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#LogRatingParams">LogRatingParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#LogRatingResponse">LogRatingResponse</a>

Methods:

- <code title="post /ratings">client.Ratings.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#RatingService.Log">Log</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#RatingLogParams">RatingLogParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#LogRatingResponse">LogRatingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Organizations

Params Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateOrganizationParams">CreateOrganizationParams</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UpdateOrganizationParams">UpdateOrganizationParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateOrganizationResponse">CreateOrganizationResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#GetOrganizationAnalyticsResponse">GetOrganizationAnalyticsResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ListOrganizationsResponse">ListOrganizationsResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#TenantOrganization">TenantOrganization</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UpdateOrganizationResponse">UpdateOrganizationResponse</a>

Methods:

- <code title="post /organizations">client.Organizations.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#OrganizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#OrganizationNewParams">OrganizationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateOrganizationResponse">CreateOrganizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /organizations/{organizationId}">client.Organizations.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#OrganizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#OrganizationUpdateParams">OrganizationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UpdateOrganizationResponse">UpdateOrganizationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /organizations">client.Organizations.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#OrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#OrganizationListParams">OrganizationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ListOrganizationsResponse">ListOrganizationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /organizations/{organizationId}/analytics">client.Organizations.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#OrganizationService.GetOrganizationAnalytics">GetOrganizationAnalytics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#OrganizationGetOrganizationAnalyticsParams">OrganizationGetOrganizationAnalyticsParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#GetOrganizationAnalyticsResponse">GetOrganizationAnalyticsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Prompts

Params Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ComponentInputParam">ComponentInputParam</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ComponentUpdateParam">ComponentUpdateParam</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreatePromptParams">CreatePromptParams</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UpdatePromptParams">UpdatePromptParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreatePromptResponse">CreatePromptResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#DeletePromptResponse">DeletePromptResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#GetPromptResponse">GetPromptResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ListPromptsResponse">ListPromptsResponse</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#Prompt">Prompt</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#PromptComponent">PromptComponent</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#SlimPrompt">SlimPrompt</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#SlimPromptComponent">SlimPromptComponent</a>
- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UpdatePromptResponse">UpdatePromptResponse</a>

Methods:

- <code title="post /prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#PromptService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#PromptNewParams">PromptNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreatePromptResponse">CreatePromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /prompts/{id}">client.Prompts.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#PromptService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#PromptUpdateParams">PromptUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#UpdatePromptResponse">UpdatePromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#PromptService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#PromptListParams">PromptListParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#ListPromptsResponse">ListPromptsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /prompts/{id}">client.Prompts.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#PromptService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#DeletePromptResponse">DeletePromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /prompts/{id}">client.Prompts.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#PromptService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#GetPromptResponse">GetPromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Events

Params Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateEventParams">CreateEventParams</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateEventResponse">CreateEventResponse</a>

Methods:

- <code title="post /events">client.Events.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#EventService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#EventNewParams">EventNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/greenflash-ai/go">greenflashpublicapi</a>.<a href="https://pkg.go.dev/github.com/greenflash-ai/go#CreateEventResponse">CreateEventResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
