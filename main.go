package main

import (
	"net/url"
)

// The response sent by the snowflake api with the [token] field.
type TokenResponse struct {
	// The token data sent by the api.
	Token string `json:"token"`
}

// The roast type of response sent by the snowflake api.
type RoastResponse struct {
	// The roast content sent by the api.
	Roast string `json:"roast"`
}

// The response sent by the snowflake with one data string field.
type DataResponse struct {
	// The data content sent by the api.
	Data string `json:"data"`
}

// The response sent by the memer api of the snowflake api.
type MemeResponse struct {
	// Boolean stating is that meme a video or not.
	IsVideo bool `json:"isVideo"`
	// Boolean stating is the meme nsfw or not.
	NSFW bool `json:"nsfw"`
	// Unix timestamp when the meme was created at.
	CreatedAt string `json:"createdAt"`
	// Not documented.
	URL string `json:"url"`
	// The subreddit of the meme.
	Subreddit string `json:"subreddit"`
	// The title of the meme post.
	Title string `json:"title"`
	// Not documented.
	Link string `json:"link"`
	// The ratings of the meme with upvotes, downvotes and comments.
	Ratings struct {
		// The number of upvotes for the meme.
		Upvotes uint32 `json:"upvote"`
		// The number of downvotes for the meme.
		Downvotes uint32 `json:"downvote"`
		// The number of comments for the meme.
		Comments uint32 `json:"comments"`
	} `json:"ratings"`
}

// Data structure represeting pokemon info.
type PokemonInfo struct {
	// The name of the pokemon.
	Name string `json:"name"`
	// The pokemon id which represents it.
	ID uint64 `json:"id"`
	// Base experience of the pokemon.
	BaseXP float64 `json:"baseExperience"`
	// Height of the pokemon.
	Height float64 `json:"height"`
	// Weight of the pokemon.
	Weight float64 `json:"weight"`
	// The type of the pokemon.
	Type string `json:"type"`
	// An array of moves which the pokemon can use.
	Moves []string `json:"moves"`
	// Not documented.
	Stats interface{} `json:"stats"`
	// The image url of the pokemon.
	Image string `json:"image"`
}

// Data structure representing the deno module information.
type DenoModuleInfo struct {
	// The registry name of the module. 'Deno'.
	Registry string `json:"registry"`
	// The icon of the registry. 'https://www.iconfinder.com/data/icons/logos-brands-5/24/deno-512.png'.
	Icon string `json:"icon"`
	// The url to the registry. 'https://deno.land'.
	URL string `json:"url"`
	// The module information.
	Module struct {
		// The name of the module.
		Name string `json:"name"`
		// The url to the module in the registry.
		URL string `json:"url"`
		// The description of the module.
		Description string `json:"description"`
		// The version of the module.
		Version string `json:"version"`
		// The stars of the module.
		Stars string `json:"stars"`
		// The github url of the module.
		Github string `json:"github"`
		// The unix timetstamp when the module was created at.
		CreatedAt string `json:"createdAt"`
		// The information of the developer of the module.
		Developer struct {
			// The name of the developer.
			Name string `json:"name"`
			// The url to the profile of the developer.
			URL string `json:"url"`
		} `json:"developer"`
	} `json:"module"`
}

// Data structure representing the npm module information.
type NPMModuleInfo struct {
	// The registry name of the module. 'npmjs'.
	Registry string `json:"registry"`
	// The icon of the registry. 'https://pbs.twimg.com/profile_images/1285630920263966721/Uk6O1QGC_400x400.jpg'.
	Icon string `json:"icon"`
	// The url to the registry. 'https://npmjs.com'.
	URL string `json:"url"`
	// The runkit of the module.
	Runkit string `json:"runkit"`
	// The module information.
	Module struct {
		// The name of the module.
		Name string `json:"name"`
		// The url to the module in the registry.
		URL string `json:"url"`
		// The description of the module.
		Description string `json:"description"`
		// The version of the module.
		Version string `json:"version"`
		// The main file of the module.
		Main string `json:"main"`
		// The license of the module.
		License string `json:"license"`
		// The author name of the module.
		Author string `json:"author"`
		// The maintainers of the module
		Maintainers []string `json:"maintainers"`
		// The dependencies of the module required.
		Dependencies []string `json:"dependencies"`
		// The repository details of the module.
		Repository struct {
			// The type of the repository.
			Type string `json:"type"`
			// The url to the repository.
			URL string `json:"url"`
		} `json:"repository"`
		// The banner url of the module.
		Banner string `json:"banner"`
	} `json:"module"`
}

// Data structure representing the pypi module information.
type PypiModuleInfo struct {
	// The registry name of the module. 'pypi'.
	Registry string `json:"registry"`
	// The icon of the registry. 'https://jollyremedy.files.wordpress.com/2018/02/pypi.png?w=656'.
	Icon string `json:"icon"`
	// The url to the registry. 'https://pypi.org'.
	URL string `json:"url"`
	// The module information.
	Module struct {
		// The name of the module.
		Name string `json:"name"`
		// The url to the module in the registry.
		URL string `json:"url"`
		// The description of the module.
		Description string `json:"description"`
		// The version of the module.
		Version string `json:"version"`
		// The author of the module.
		Author string `json:"author"`
		// The unix timestamp when the module was last updated at.
		UpdatedAt string `json:"updatedAt"`
		// The documentation url of the module.
		Documentation string `json:"documentation"`
		// The homepage url of the module.
		Homepage string `json:"homepage"`
	} `json:"module"`
}

// Data structure representing the discord token information.
type DiscordTokenInfo struct {
	// The type of the discord token.
	Type string `json:"type"`
	// The discord token.
	Token string `json:"token"`
	// The id of the app.
	ID string `json:"id"`
	// The username of the app.
	Username string `json:"username"`
	// The discriminator of the app.
	Discriminator uint32 `json:"discriminator"`
	// The avatar if of the app.
	Avatar string `json:"avatar"`
	// The avatar url of the app.
	AvatarURL string `json:"avatarURL"`
	// Not Documented.
	SnowflakeInfo interface{} `json:"snowflakeInfo"`
}

// Data structure representing the api stats.
type ApiStats struct {
	// The total requests received by the api.
	TotalRequests uint32 `json:"total_requests"`
	// The total free users of the api.
	FreeUsers uint32 `json:"free_users"`
	// The total pro users of the api.
	ProUsers uint32 `json:"pro_users"`
	// The total users of the api.
	TotalUsers uint32 `json:"total_users"`
	// The total banned users of the api.
	BannedUsers uint32 `json:"banned_users"`
	// The os used by the api.
	OS string `json:"os"`
	// The processor info which is used by the api.
	Processor struct {
		// The model of the processor.
		Model string `json:"model"`
		// Not Documented.
		Count uint32 `json:"count"`
	} `json:"processor"`
	// The memory stats of the api.
	Memory struct {
		// The heap total of the memory.
		HeapTotal float64 `json:"heap_total"`
		// The heap which is been used by the memory.
		HeapUsed float64 `json:"heap_used"`
		// Not Documented.
		RSS float64 `json:"rss"`
		// Not Documented.
		AB float64 `json:"ab"`
		// Not Documented.
		External float64 `json:"external"`
	} `json:"memory"`
}

// Data structure representing the api profile info.
type ProfileInfo struct {
	// The discord user if of the profile user.
	User string `json:"user"`
	// Boolean stating is the user has subscribed for premium or not.
	IsPro bool `json:"pro"`
	// Number of ratelimits occured.
	Ratelimits uint32 `json:"ratelimits"`
	// Number of requests received by the api from the current user.
	Requests string `json:"requests"`
	// The timestamp when the token was created at.
	TokenCreatedTimestamp uint64 `json:"tokenCreatedTimestamp"`
	// Not Documented.
	CreatedTimestamp uint64 `json:"createdTimestamp"`
}

// Data structure representing the github statistics.
type GithubStats struct {
	// The name of the github user.
	Name string `json:"name"`
	// The avatar url of the github user.
	Avatar string `json:"avatar"`
	// The follower count of the user.
	Followers uint32 `json:"followers"`
	// The repository count of the user.
	Repos uint32 `json:"repos"`
	// The pull request count of the user.
	PullRequests uint32 `json:"pullRequests"`
	// The issues count created by the user.
	Issues uint32 `json:"issues"`
	// The number of npm downloads of the user's modules.
	NPMDownloads uint32 `json:"npmDownloads"`
}

// Data structure representing the Youtube video information.
type YoutubeVideo struct {
	// The youtube video id.
	ID string `json:"id"`
	// The title of the youtube video
	Title string `json:"title"`
	// The author of the youtube video.
	Author string `json:"author"`
	// The url to the youtube video.
	URL string `json:"url"`
	// The timestamp when the video was published.
	PublishedAt uint64 `json:"publishedAt"`
	// The thumbnail url of the video.
	Thumbnail string `json:"thumbnail"`
}

// Data structure representing a Youtube Channel.
type YoutubeChannel struct {
	// The channel basic information.
	Channel struct {
		// The name of the channel
		Name string `json:"name"`
		// The url to the youtube channel.
		URL string `json:"url"`
	} `json:"channel"`
	// The videos of the youtube channel.
	Videos []YoutubeVideo `json:"videos"`
}

// The main client to access the Snowflake api.
type Client struct {
	// The token required for the client.
	Token string
}

// Creates a new client. A very shortand method.
func NewClient(token string) Client {
	return Client{token}
}

// Creates a default chat option object with the message field only.
func NewDefaultChatOptions(message string) ChatOptions {
	return ChatOptions{"Chatbot", "female", 1, message}
}

// Get a fake discord bot token from the snowflake api. (You can basically generate it yourself).
func (self *Client) GetFakeDiscordBotToken() (TokenResponse, error) {
	var response TokenResponse
	err := self.Fetch("GET", "/api/token", &response)
	return response, err
}

// Get a random meme from the api.
func (self *Client) GetMeme() (MemeResponse, error) {
	var response MemeResponse
	err := self.Fetch("GET", "/api/meme", &response)
	return response, err
}

// Get a random meme from a paticular subreddit from the api.
func (self *Client) GetMemeSbr(sbr string) (MemeResponse, error) {
	var response MemeResponse
	err := self.Fetch("GET", "/api/meme?sbr="+sbr, &response)
	return response, err
}

// Get an random image of a type which will be one of (dog, cat, duck, fox).
func (self *Client) GetImage(t string) ([]byte, error) {
	return self.FetchBinary("GET", "/api/"+t)
}

// Get a random dog image.
func (self *Client) GetDogImage() ([]byte, error) {
	return self.GetImage("dog")
}

// Get a random cat image.
func (self *Client) GetCatImage() ([]byte, error) {
	return self.GetImage("cat")
}

// Get a random dog image.
func (self *Client) GetFoxImage() ([]byte, error) {
	return self.GetImage("fox")
}

// Get a random duck image.
func (self *Client) GetDuckImage() ([]byte, error) {
	return self.GetImage("duck")
}

// Get a random roast from the api.
func (self *Client) GetRoast() (RoastResponse, error) {
	var response RoastResponse
	err := self.Fetch("GET", "/api/roast", &response)
	return response, err
}

// Get the pokemon info by its name.
func (self *Client) GetPokemon(name string) (PokemonInfo, error) {
	var response PokemonInfo
	err := self.Fetch("GET", "/api/pokemon?name="+url.QueryEscape(name), &response)
	return response, err
}

// Encode the readable string to morsecode using the api.
func (self *Client) EncodeMorse(text string) (DataResponse, error) {
	var response DataResponse
	err := self.Fetch("GET", "/api/morse/encode?text="+url.QueryEscape(text), &response)
	return response, err
}

// Decode the morse code to readable string using the api.
func (self *Client) DecodeMorse(text string) (DataResponse, error) {
	var response DataResponse
	err := self.Fetch("GET", "/api/morse/decode?text="+url.QueryEscape(text), &response)
	return response, err
}

// Encode the readable string to base64 using the api.
func (self *Client) EncodeBase64(text string) (DataResponse, error) {
	var response DataResponse
	err := self.Fetch("GET", "/api/base64/encode?data="+url.QueryEscape(text), &response)
	return response, err
}

// Decode the base64 to readable string using the api.
func (self *Client) DecodeBase64(text string) (DataResponse, error) {
	var response DataResponse
	err := self.Fetch("GET", "/api/base64/decode?data="+url.QueryEscape(text), &response)
	return response, err
}

// Get the deno module info by its name.
func (self *Client) GetDenoModuleInfo(name string) (DenoModuleInfo, error) {
	var response DenoModuleInfo
	err := self.Fetch("GET", "/api/registry/deno?module="+name, &response)
	return response, err
}

// Get the npm module info by its name.
func (self *Client) GetNPMModuleInfo(name string) (NPMModuleInfo, error) {
	var response NPMModuleInfo
	err := self.Fetch("GET", "/api/registry/npm?module="+name, &response)
	return response, err
}

// Get the pypi module info by its name.
func (self *Client) GetPypiModuleInfo(name string) (PypiModuleInfo, error) {
	var response PypiModuleInfo
	err := self.Fetch("GET", "/api/registry/pypi?module="+name, &response)
	return response, err
}

// Get the info of a discord token.
func (self *Client) GetDiscordTokenInfo(token string) (DiscordTokenInfo, error) {
	var response DiscordTokenInfo
	err := self.Fetch("GET", "/api/tokeninfo?token="+token, &response)
	return response, err
}

// Get the info of the api stats.
func (self *Client) GetApiStats() (ApiStats, error) {
	var response ApiStats
	err := self.Fetch("GET", "/api/stats", &response)
	return response, err
}

// Get the profile info of the current user.
func (self *Client) GetProfileInfo() (ProfileInfo, error) {
	var response ProfileInfo
	err := self.Fetch("GET", "/api/me", &response)
	return response, err
}

// Get the github stats info by username.
func (self *Client) GetGithubStats(username string) (GithubStats, error) {
	var response GithubStats
	err := self.Fetch("GET", "/api/githubstats?username="+username, &response)
	return response, err
}

// Get the youtube channel info by its id.
func (self *Client) GetYoutubeChannelInfo(id string) (YoutubeChannel, error) {
	var response YoutubeChannel
	err := self.Fetch("GET", "/api/youtube?channel="+id, &response)
	return response, err
}
