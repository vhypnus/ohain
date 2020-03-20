package base

const (
	ROOT = "/data"

	REPO_MAP[string]Repository = {
		"":
	}
)


type Repository struct {

	name string

	path string

	url string

	modelpath string

	apipath string
} 

func NewRepository(url string) {

}

func NewRepository(name string) {

}

func (r *Repository) merge() {

}


func (r *Repository) sync(branch string) (bool,error) {

}

func (r *Repository) diff(src string,target string) {

}


func (r *Repository) api(branch string) string {

}

func (r *Repository) model(branch string) string {

}