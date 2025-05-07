package domain

// Actor representa la información de un usuario.
type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

// Repo representa la información de un repositorio.
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Commit representa los datos de un commit dentro de un evento PushEvent.
type Commit struct {
	SHA    string `json:"sha"`
	Author struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	URL      string `json:"url"`
}

type User struct {
	Login          string `json:"login"`
	ID             int    `json:"id"`
	NodeID         string `json:"node_id"`
	AvatarURL      string `json:"avatar_url"`
	GravatarID     string `json:"gravatar_id"`
	URL            string `json:"url"`
	HTMLURL        string `json:"html_url"`
	FollowersURL   string `json:"followers_url"`
	FollowingURL   string `json:"following_url"`
	GistsURL       string `json:"gists_url"`
	StarredURL     string `json:"starred_url"`
	Subscriptions  string `json:"subscriptions_url"`
	Organizations  string `json:"organizations_url"`
	ReposURL       string `json:"repos_url"`
	EventsURL      string `json:"events_url"`
	ReceivedEvents string `json:"received_events_url"`
	Type           string `json:"type"`
	UserViewType   string `json:"user_view_type"`
	SiteAdmin      bool   `json:"site_admin"`
}

// Label representa una etiqueta asociada a un issue.
type Label struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}

// Issue representa un issue en GitHub.
type Issue struct {
	URL           string  `json:"url"`
	RepositoryURL string  `json:"repository_url"`
	LabelsURL     string  `json:"labels_url"`
	CommentsURL   string  `json:"comments_url"`
	EventsURL     string  `json:"events_url"`
	HTMLURL       string  `json:"html_url"`
	ID            int     `json:"id"`
	NodeID        string  `json:"node_id"`
	Number        int     `json:"number"`
	Title         string  `json:"title"`
	User          User    `json:"user"`
	Labels        []Label `json:"labels"`
	State         string  `json:"state"`
	Locked        bool    `json:"locked"`
	Comments      int     `json:"comments"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	ClosedAt      *string `json:"closed_at,omitempty"`
	Body          string  `json:"body"`
}

// Comment representa un comentario dentro de un issue.
type Comment struct {
	URL       string `json:"url"`
	HTMLURL   string `json:"html_url"`
	IssueURL  string `json:"issue_url"`
	ID        int    `json:"id"`
	NodeID    string `json:"node_id"`
	User      User   `json:"user"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Payload representa los datos específicos del evento.
type Payload struct {
	Action  string  `json:"action"`
	Issue   Issue   `json:"issue"`
	Comment Comment `json:"comment"`
	Size    int     `json:"size"`
}

// Event representa un evento de GitHub.
type Event struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Actor     Actor   `json:"actor"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	Public    bool    `json:"public"`
	CreatedAt string  `json:"created_at"`
}
