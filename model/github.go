package model

type GithubRelease struct {
	Url             string               `json:"url,omitempty"`
	HtmlUrl         string               `json:"html_url,omitempty"`
	AssetsUrl       string               `json:"assets_url,omitempty"`
	UploadUrl       string               `json:"upload_url,omitempty"`
	TarballUrl      string               `json:"tarball_url,omitempty"`
	ZipballUrl      string               `json:"zipball_url,omitempty"`
	Id              int                  `json:"id,omitempty"`
	NodeId          string               `json:"node_id,omitempty"`
	TagName         string               `json:"tag_name,omitempty"`
	TargetCommitish string               `json:"target_commitish,omitempty"`
	Name            string               `json:"name,omitempty"`
	Body            string               `json:"body,omitempty"`
	Draft           bool                 `json:"draft,omitempty"`
	Prerelease      bool                 `json:"prerelease,omitempty"`
	CreatedAt       string               `json:"created_at,omitempty"`
	PublishedAt     string               `json:"published_at,omitempty"`
	Author          GithubUser           `json:"author"`
	Assets          []GithubReleaseAsset `json:"assets"`
}

type GithubUser struct {
	Login             string `json:"login,omitempty"`
	Id                int    `json:"id,omitempty"`
	NodeId            string `json:"node_id,omitempty"`
	AvatarUrl         string `json:"avatar_url,omitempty"`
	GravatarId        string `json:"gravatar_id,omitempty"`
	Url               string `json:"url,omitempty"`
	HtmlUrl           string `json:"html_url,omitempty"`
	FollowersUrl      string `json:"followers_url,omitempty"`
	FollowingUrl      string `json:"following_url,omitempty"`
	GistsUrl          string `json:"gists_url,omitempty"`
	StarredUrl        string `json:"starred_url,omitempty"`
	SubscriptionsUrl  string `json:"subscriptions_url,omitempty"`
	OrganizationsUrl  string `json:"organizations_url,omitempty"`
	ReposUrl          string `json:"repos_url,omitempty"`
	EventsUrl         string `json:"events_url,omitempty"`
	ReceivedEventsUrl string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}

type GithubReleaseAsset struct {
	Url                string     `json:"url,omitempty"`
	BrowserDownloadUrl string     `json:"browser_download_url,omitempty"`
	Id                 int        `json:"id,omitempty"`
	NodeId             string     `json:"node_id,omitempty"`
	Name               string     `json:"name,omitempty"`
	Label              string     `json:"label,omitempty"`
	State              string     `json:"state,omitempty"`
	ContentType        string     `json:"content_type,omitempty"`
	Size               int        `json:"size,omitempty"`
	DownloadCount      int        `json:"download_count,omitempty"`
	CreatedAt          string     `json:"created_at,omitempty"`
	UpdatedAt          string     `json:"updated_at,omitempty"`
	Uploader           GithubUser `json:"uploader"`
}
