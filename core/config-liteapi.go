package config

import liteapi "github.com/liteapi-travel/go-sdk/v2"

func ApiClient() *liteapi.APIClient {
	configuration := liteapi.NewConfiguration()
	configuration.AddDefaultHeader("X-API-KEY", "YOUR_API_KEY")
	apiClient := liteapi.NewAPIClient(configuration)
	return apiClient
}
