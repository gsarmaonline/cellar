package core

type (
	App struct {
		ResourceObject

		Services    []*Service    `json:"services,omitempty"`
		Databases   []*Database   `json:"databases,omitempty"`
		Caches      []*Cache      `json:"caches,omitempty"`
		Workerpools []*Workerpool `json:"workerpools,omitempty"`

		ShareInstance bool `json:"share_instance,omitempty"` // Whether to share instances across services
	}
)

func DefaultApp() *App {
	return &App{
		ResourceObject: ResourceObject{
			Name: "webapp",
		},
		ShareInstance: true,
		Services: []*Service{
			{
				ResourceObject: ResourceObject{
					Name: "web_service",
				},
				ExecutorType: ExecutorTypeGolang,
			},
		},
		Databases: []*Database{
			{
				ResourceObject: ResourceObject{
					Name: "web_database",
				},
				DatabaseType: DatabaseTypeSqlite,
			},
		},
	}
}

func (w *App) Deploy() error {
	for _, db := range w.Databases {
		if err := db.Deploy(); err != nil {
			return err
		}
	}
	for _, cache := range w.Caches {
		if err := cache.Deploy(); err != nil {
			return err
		}
	}
	for _, service := range w.Services {
		if err := service.Deploy(); err != nil {
			return err
		}
	}
	for _, workerpool := range w.Workerpools {
		if err := workerpool.Deploy(); err != nil {
			return err
		}
	}
	return nil
}
